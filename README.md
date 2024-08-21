# DAL

**Data Access Layer**
DAL is a proxy layer for SQL databases with a MongoDB inspired query interface.
It can be used as a Go or NodeJS package (requires compiler) to create your own proxy and apply custom middlewares.

_Notes:_

- This project is still in early alpha. You need to build it yourself and use at your own risk.
- At the time only SQLite is implemented, however, other drivers might work.

_Use cases:_

- For IOT networks when MySQL/PG are too heavy.
- If you need a layer between your application and the database (i.e. for caching).
- If you want a MongoDB-like query interface for your SQL.

## Usage

### Server

The most efficient way to use DAL is to run the server as a standalone service.

Build:

```bash
go build -o server
```

Run:

```bash
export SQLITE_DIRECTORY=/opt/data

./server

2024/08/21 22:01:54 Starting server on port 8118
2024/08/21 22:01:54 Using directory: /opt/data
```

### Client

Install:

```bash
pnpm add git+git@github.com:nesterow/dal.git
```

### Query Interface

| Method                                       | Description                            | SQL                                             |
| -------------------------------------------- | -------------------------------------- | ----------------------------------------------- |
| `In(table: string)`                          | Select table                           | `SELECT * FROM table`                           |
| `Find(filter: object)`                       | Filter rows                            | `SELECT * FROM table WHERE filter`              |
| `Fields(fields: string[])`                   | Select fields                          | `SELECT fields, FROM table`                     |
| `Sort(sort)`                                 | Sort rows                              | `SELECT * FROM table ORDER BY sort`             |
| `Limit(limit: number)`                       | Limit rows                             | `SELECT * FROM table LIMIT limit`               |
| `Offset(offset: number)`                     | Offset rows                            | `SELECT * FROM table OFFSET offset`             |
| `Join({ $for: "t_2", $do: { "t.a": "b" } })` | Join tables                            | `SELECT * FROM table t JOIN t_2 ON t.a = b`     |
| `Insert({name: "J"}, {name: "B"})`           | Insert row                             | `INSERT INTO table (name,) VALUES ('J', 'B')`   |
| `Set({name: "Julian"})`                      | Update row (Find(filter).Set({}))      | `UPDATE table SET name = 'Julian' WHERE filter` |
| `Delete()`                                   | Delete row (Find(filter).Delete())     | `DELETE FROM table WHERE filter`                |
| `As(DTO)`                                    | Map rows to a DTO                      | `SELECT * FROM table`                           |
| `Rows()`                                     | Get rows iterator                      | `SELECT * FROM table`                           |
| `Exec()`                                     | Execute query (update, insert, delete) | `SQL RESULT`                                    |
| `Query()`                                    | Query database                         | DTO array                                       |
| `Tx()`                                       | Run in trasaction                      |                                                 |

### Filters

| Filter                        | Description            | SQL                            |
| ----------------------------- | ---------------------- | ------------------------------ |
| `{id: 1, num: 2}`             | Equals, default filter | `WHERE id = 1 AND num = 2`     |
| `{id: { $eq: 1 }}`            | Equals, explicit       | `WHERE id = 1`                 |
| `{id: { $gt: 1 }}`            | Greater than           | `WHERE id > 1`                 |
| `{id: { $gte: 1 }}`           | Greater than or equal  | `WHERE id >= 1`                |
| `{id: { $lt: 1 }}`            | Less than              | `WHERE id < 1`                 |
| `{id: { $lte: 1 }}`           | Less than or equal     | `WHERE id <= 1`                |
| `{id: { $ne: 1 }}`            | Not equal              | `WHERE id != 1`                |
| `{id: { $in: [1, 2] }}`       | In                     | `WHERE id IN (1, 2)`           |
| `{id: { $nin: [1, 2] }}`      | Not in                 | `WHERE id NOT IN (1, 2)`       |
| `{id: { $like: "a" }}`        | Like                   | `WHERE id LIKE '%a%'`          |
| `{id: { $nlike: "a" }}`       | Not like               | `WHERE id NOT LIKE '%a%'`      |
| `{id: { $between: [1, 2] }}`  | Between                | `WHERE id BETWEEN 1 AND 2`     |
| `{id: { $nbetween: [1, 2] }}` | Not between            | `WHERE id NOT BETWEEN 1 AND 2` |
| `{id: { $glob: "\*son" }}`    | Glob                   | `WHERE id GLOB '*son'`         |

### Example

```typescript
import { DAL } from "@nesterow/dal";

class UserDTO {
  id: number = 0;
  name: string = "";
  data: string = "";
  age: number | undefined;
}

const db = new DAL({
  database: "test.sqlite",
  url: "http://localhost:8111",
});

const rows = db
  .In("test t")
  .Find({
    name: { $glob: "*son" },
    age: { $gte: 18 },
  }) // SELECT * FROM test t WHERE name GLOB '*son' AND age >= 18
  .As(UserDTO) // Map every row to DTO
  .Rows(); // Get iterator

for await (const row of rows) {
  console.log(row); // Jason, Jackson
}
```
