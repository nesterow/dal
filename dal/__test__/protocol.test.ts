import { test, expect } from "bun:test";
import { DAL } from ".."

const options = {
    database: "test.sqlite",
    url: "http://localhost:8111",
}


test("Rows iter",  async () => {
    const dal = new DAL(options);
    const rows = dal
        .In("test t")
        .Find({
            id: 1,
        })
        .Rows();
    for await (const row of rows) {
        // console.log(row);
        //@ts-ignore
        expect(row.r.length).toBe(3);
    }
    expect(true).toBe(true);
});
