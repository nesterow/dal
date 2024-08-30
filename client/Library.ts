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

let Library: SQLite;

if (process.isBun) {
  Library = require("./Bunding") as SQLite;
} else {
  Library = require("./Binding") as SQLite;
}

export default Library;
