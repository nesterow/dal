import { createRequire } from "node:module";
const require = createRequire(import.meta.url);
type RowIterator = {
  next: () => Buffer;
  free: () => void;
};
type SQLite = {
  InitSQLite: (pragmas: Buffer) => void;
  RowIterator: (input: Buffer) => RowIterator;
};
export default require("../build/Release/dal.node") as SQLite;
