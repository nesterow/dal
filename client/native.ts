import Builder from "./Builder";
import Napi from "./napi";
import { encodeRequest, decodeRows, decodeResponse } from "./Protocol";
import type { ExecResult, IError } from "./Protocol";

//@ts-ignore
const Binding = Napi.default ?? Napi;

type Options = {
  database: string;
};

/**
 * Allows to use SQLite databases in a NodeJS process.
 */
export default class C <
  I extends abstract new (...args: any) => any,
> extends Builder<I> {
  constructor(opts: Options) {
    super({ database: opts.database, url: "" });
  }
  async *Rows<T = InstanceType<I>>(): AsyncGenerator<[T, IError]> {
    this.formatRequest();
    const req = Buffer.from(encodeRequest(this.request));
    const iter = Binding.rowIterator(req);
    for (;;) {
      const data = iter.next() as Buffer;
      const [_, error] = decodeResponse(data);
      if (error) {
        yield [{} as T, error];
        iter.cleanup();
        return;
      }
      const [rows, err] = decodeRows(data);
      if (err || !rows || rows.length === 0) {
        iter.cleanup();
        return;
      }
      for (const row of rows) {
        if (this.headerRow === null) {
          this.headerRow = row.r;
          continue;
        }
        yield [this.formatRow(row.r), null];
      }
    }
  }
  async Exec(): Promise<[ExecResult, IError]> {
    this.formatRequest();
    const req = Buffer.from(encodeRequest(this.request));
    const iter = Binding.rowIterator(req);
    const response = iter.next();
    return decodeResponse(response)!;
  }
}
