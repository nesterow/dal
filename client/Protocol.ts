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
  Error?: string;
}

export interface Row {
  r: unknown[];
}

export type IError = any;

export const METHODS =
  "Raw|In|Find|Select|Fields|Join|Group|Sort|Limit|Offset|Delete|Insert|Set|Update|OnConflict|DoUpdate|DoNothing|Tx".split(
    "|",
  );

export function encodeRequest(request: Request): Uint8Array {
  return encode(request);
}

export function decodeResponse(input: Uint8Array): [ExecResult, IError] {
  try {
    const res = decode(input) as {
      i: number;
      ra: number;
      li: number;
      e?: string;
    };
    const result = {
      Id: res.i,
      RowsAffected: res.ra,
      LastInsertId: res.li,
      Error: res.e,
    };
    return [result, result.Error];
  } catch (e) {
    return [{} as ExecResult, e];
  }
}

const ROW_TAG = [0x81, 0xa1, 0x72];

export function decodeRows(input: Uint8Array): [Row[], IError] {
  try {
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
    return [
      rows.map((row) => decode(new Uint8Array(row as number[]))) as Row[],
      null,
    ];
  } catch (e) {
    return [[], e];
  }
}

export async function* decodeRowsIterator(
  stream: ReadableStream<Uint8Array>,
): AsyncGenerator<[Row, IError]> {
  const reader = stream.getReader();
  for (;;) {
    const { value, done } = await reader.read();
    if (done) {
      break;
    }
    const [rows, err] = decodeRows(value);
    if (err) {
      yield [{} as Row, err];
      break;
    }
    for (const row of rows) {
      yield [row, null];
    }
  }
}
