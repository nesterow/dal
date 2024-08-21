import { createRequire } from "node:module";
const require = createRequire(import.meta.url);
type SQLite = {
  InitSQLite: (pragmas: Buffer) => void;
  Handle: (input: Buffer) => Buffer;
};
export default require("../build/Release/dal.node") as SQLite;
