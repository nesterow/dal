import Builder from "./Builder";
import Binding from "./Binding";
import { encodeRequest, decodeRows, decodeResponse } from "./Protocol";
import type { ExecResult } from "./Protocol";

type Options = {
  database: string;
};

/**
 * Allows to use SQLite databases in a NodeJS process.
 * It is less memory-efficient than a seaparate server, and uses absolute path for database name.
 */
export default class CBuilder<
  I extends abstract new (...args: any) => any,
> extends Builder<I> {
  constructor(opts: Options) {
    super({ database: opts.database, url: "" });
  }
  /**
   * Not really an iterator, since addonn allocates memory for all rows
   * but returns an iterator
   */
  async *Rows<T = InstanceType<I>>(): AsyncGenerator<T> {
    this.formatRequest();
    const req = Buffer.from(encodeRequest(this.request));
    const response = Binding.Handle(req);
    const rows = decodeRows(response);
    for (const row of rows) {
      if (this.headerRow === null) {
        this.headerRow = row.r;
        continue;
      }
      yield this.formatRow(row.r);
    }
  }
  async Query<T = InstanceType<I>>(): Promise<T[]> {
    const rows = this.Rows();
    const result: T[] = [];
    for await (const row of rows) {
      result.push(row);
    }
    return result;
  }
  async Exec(): Promise<ExecResult> {
    this.formatRequest();
    const req = Buffer.from(encodeRequest(this.request));
    const response = Binding.Handle(req);
    return decodeResponse(response);
  }
}
