const fs = require("fs");
const dal = require("../../build/Release/dal.node");
//dal.InitSQLite(Buffer.from([]));
const buf = fs.readFileSync("./pkg/__test__/proto_test.msgpack");
//console.log(dal.Handle(buf));
const iter = dal.RowIterator(buf);
console.log(iter);
for (let i = 0; i < 10; i++) {
  const b = iter.next();
  if (b.length === 0) {
    break;
  }
  console.log(b.toString());
}
iter.free();