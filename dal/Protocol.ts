import { encode, decode } from "@msgpack/msgpack";

export interface Method {
  method: string;
  args: any;
}

export interface Request {
  id: number;
  db: string;
  commands: Method[];
}

export interface ExecResult {
  Id: number;
  RowsAffected: number;
  LastInsertId: number;
  Msg?: string;
}

export interface Row {
  r: unknown[];
}

export const METHODS =
  "Raw|In|Find|Select|Fields|Join|Group|Sort|Limit|Offset|Delete|Insert|Set|Update|OnConflict|DoUpdate|DoNothing|Tx".split(
    "|",
  );

export function encodeRequest(request: Request): Uint8Array {
  return encode(request);
}

export function decodeResponse(input: Uint8Array): ExecResult {
  const res = decode(input) as {
    i: number;
    ra: number;
    li: number;
    m?: string;
  };
  return {
    Id: res.i,
    RowsAffected: res.ra,
    LastInsertId: res.li,
    Msg: res.m,
  };
}

const ROW_TAG = [0x81, 0xa1, 0x72];

export function decodeRows(input: Uint8Array): Row[] {
  const rows = [];
  let count = 0;
  let buf = [];
  while (count < input.length) {
    if (input.at(count) != 0x81) {
      buf.push(input.at(count));
      count++;
      continue;
    }
    const [a, b, c] = ROW_TAG;
    const [aa, bb, cc] = input.slice(count, count + 4);
    if (aa == a && bb == b && cc == c) {
      rows.push([...ROW_TAG, ...buf]);
      buf = [];
      count += 3;
    } else {
      buf.push(input.at(count));
      count++;
    }
  }
  rows.push([...ROW_TAG, ...buf]);
  rows.shift();
  return rows.map((row) => decode(new Uint8Array(row as number[]))) as Row[];
}

export async function* decodeRowsIterator(
  stream: ReadableStream<Uint8Array>,
): AsyncGenerator<Row> {
  const reader = stream.getReader();
  let buf = new Uint8Array();
  for (;;) {
    const { value, done } = await reader.read();
    if (done) {
      break;
    }
    buf = new Uint8Array([...buf, ...value]);
    // the server flushes after each row
    // so we decode "complete" rows
    const rows = decodeRows(buf);
    for (const row of rows) {
      yield row;
    }
  }
}
