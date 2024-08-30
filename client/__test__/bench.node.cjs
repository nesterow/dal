const fs = require("fs");
const dal = require("../../build/Release/dal.node");

const Mb = (num) => Math.round(num / 1024 / 1024);

class Stats {
  constructor() {
    this.calls = 0;
    this.avg_rss = 0;
    this.avg_heapTotal = 0;
    this.avg_heapUsed = 0;
    this.avg_external = 0;
  }
  add(mem) {
    this.calls++;
    this.avg_rss += mem.rss;
    this.avg_heapTotal += mem.heapTotal;
    this.avg_heapUsed += mem.heapUsed;
    this.avg_external += mem.external;
  }
  avg() {
    const n = this.calls;
    this.avg_rss /= n;
    this.avg_heapTotal /= n;
    this.avg_heapUsed /= n;
    this.avg_external /= n;
  }
  print() {
    console.log(`
AVERAGE:
rss: ${Mb(this.avg_rss)} Mb
external: ${Mb(this.avg_external)} Mb
buffers: ${Mb(this.avg_heapUsed)} Mb
total: ${Mb(this.avg_heapTotal)} Mb`);
  }
}

const stats = new Stats();
let prevMem = process.memoryUsage();
stats.add(prevMem);

function MEM(when = "") {
  const mem = process.memoryUsage();
  stats.add(mem);
  console.log(`
${when}
rss: ${Mb(mem.rss)} Mb [delta> ${mem.rss - prevMem.rss}] 
external: ${Mb(mem.external)} Mb [delta> ${mem.external - prevMem.external}] 
buffers: ${Mb(mem.heapUsed)} Mb [delta> ${mem.heapUsed - prevMem.heapUsed}]
total: ${Mb(mem.heapTotal)} Mb [delta> ${mem.heapTotal - prevMem.heapTotal}]`);
}

console.time("Time to end");
MEM("START");

const buf = fs.readFileSync("./pkg/__test__/proto_test.msgpack");

const Iterator = dal.rowIterator(buf);
MEM("AFTER INIT");

let dataTransferedBytes = 0;
for (let i = 0; i < 100000000; i++) {
  const b = Iterator.next();
  if (b.length === 0) {
    break;
  }
  dataTransferedBytes += b.length;
  if (i % 1000000 === 0) {
    MEM(`ITERATION ${i}`);
  }
}

MEM("AFTER ITERATION");

Iterator.cleanup();
MEM("AFTER CLEANUP");

console.log("\nData transfered: ", Mb(dataTransferedBytes), "Mb");
console.timeEnd("Time to end");

setTimeout(() => {
  MEM("AFTER SOME TIME");
  stats.avg();
  stats.print();
}, 30000);
