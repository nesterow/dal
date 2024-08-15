//@ts-ignore
import { encode } from "https://deno.land/x/msgpack@v1.2/mod.ts";

const Query = {
  "db": "database.sqlite",
  "commands": [
    {"method": "In", "args": ["data"]},
    {
        "method": "Find",
        "args": [{
            "a": 1,
            "b": {
                "$gt": 2,
            },
        }]
    },
  ],
};

const encoded: Uint8Array = encode(Query);
//@ts-ignore
Deno.writeFileSync("proto_test.msgpack", encoded);