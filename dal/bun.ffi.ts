import { dlopen, FFIType, suffix, ptr } from "bun:ffi";
import { join } from "node:path";
import fs from "node:fs";

// TODO: build a shared library compatible with cffi
const libname = `libdal-arm64.${suffix}`;
const libpath = join(__dirname, '..', 'lib', libname);

const {
    symbols: {
        InitSQLite,
        CreateRowIterator,
        NextRow,
        GetLen,
        FreeIter,
    },
  } = dlopen(
    libpath,
    {
      InitSQLite: {
        args: [ FFIType.cstring ],
        returns: FFIType.void,
      },
      CreateRowIterator: {
        args: [ FFIType.cstring ],
        returns: FFIType.i32,
      },
      NextRow: {
        args: [FFIType.i32],
        returns: FFIType.cstring,
      },
      GetLen: {
        args: [FFIType.i32],
        returns: FFIType.i32,
      },
      FreeIter : {
        args: [FFIType.i32],
        returns: FFIType.void,
      },
    },
  );
