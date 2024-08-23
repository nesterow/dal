import { test, expect } from "bun:test";
import DAL from "../Builder";

const options = {
  database: "test.sqlite",
  url: "http://localhost:8111",
};

class DTO {
  id: number = 0;
  name: string = "";
  data: string = "";
  age: number | undefined;
}

test("Rows iter, no format", async () => {
  const dal = new DAL(options);
  const rows = dal
    .In("test t")
    .Find({
      id: 1,
    })
    .Rows<any[]>();
  for await (const row of rows) {
    //console.log(row);
    expect(row.length).toBe(3);
  }
  expect(true).toBe(true);
});

test("Rows iter, format", async () => {
  const dal = new DAL(options);
  const rows = dal.In("test t").Find({}).As(DTO).Rows();
  for await (const row of rows) {
    //console.log(row);
    //expect(row.id).toBe(1);
  }
  expect(true).toBe(true);
});

test("Query format", async () => {
  const dal = new DAL(options);
  const rows = await dal
    .In("test t")
    .Find({
      id: 1,
    })
    .As(DTO)
    .Query();
  for (const row of rows) {
    expect(row.id).toBeDefined();
    expect(row.age).toBeUndefined();
  }
  expect(true).toBe(true);
});

test("Query raw", async () => {
  const dal = new DAL(options);
  const rows = await dal.Raw("SELECT * FROM test WHERE id = 1").As(DTO).Query();
  for (const row of rows) {
    expect(row.id).toBeDefined();
    expect(row.age).toBeUndefined();
  }
  expect(true).toBe(true);
});
