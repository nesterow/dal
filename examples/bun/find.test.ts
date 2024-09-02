import { describe, expect, test } from "bun:test";
import path from "path";
import DAL from "@nesterow/dal/client/native";

// in this case we need to use absolute path
const DATABASE_PATH = path.join(import.meta.dir, "..", "data", "chinook.db");

const db = new DAL({
  database: DATABASE_PATH,
});

describe("Query Interface", () => {
  test(".Find [find 10 artists whose names start with 'A']", async () => {
    const items = db
      .In("artists")
      .Find({
        name: { $glob: "A*" },
      })
      .Limit(10)
      .Rows();

    for await (const result of items) {
      const [item, error] = result;
      console.log(item);
    }

    expect(true).toBe(true);
  });

  test(".Find.As [find 5 artists whose names start with 'B'; Represent each row as an Artist object]", async () => {
    class Artist {
      ArtistId = 0;
      Name = "";
    }

    const items = db
      .In("artists")
      .Find({
        name: { $glob: "B*" },
      })
      .As(Artist)
      .Limit(5)
      .Rows();

    for await (const result of items) {
      const [item, error] = result;
      console.log(123, item);
    }

    const all_rows = await db
      .In("artists")
      .Find({
        name: { $glob: "B*" },
      })
      .As(Artist)
      .Limit(5)
      .Query();

    expect(true).toBe(true);
  });
});
