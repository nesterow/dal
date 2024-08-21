const fs = require("fs");
const dal = require("../../build/Release/dal.node");
dal.InitSQLite(Buffer.from([]));
const buf = fs.readFileSync("./pkg/__test__/proto_test.msgpack");
data = dal.Handle(buf);
console.log(data);
