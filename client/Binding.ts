import { createRequire } from "node:module";
const require = createRequire(import.meta.url);
type RowIterator = {
  next: () => Buffer;
  cleanup: () => void;
};
type SQLite = {
  initSQLite: (pragmas: Buffer) => void;
  rowIterator: (input: Buffer) => RowIterator;
};
export default require("../build/Release/dal.node") as SQLite;
