# DAL [WIP]

Data Accees Layer for SQL databases written in Go.

Mongodb inspired query interface:

```typescript
const query = Db.In("users")
  .Find({
    fullname: { $glob: "*son" },
  })
  .Query();

// Result:
console.log(users)[
  ({ id: 25, fullname: "John Menson" }, { id: 76, fullname: "John Johnson" })
];
```
