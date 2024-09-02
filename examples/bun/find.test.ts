import { describe, expect, test } from "bun:test";
import path from "path";
import DAL from "@nesterow/dal/client/libdal";

// in this case we need to use absolute path
const DATABASE_PATH = path.join(import.meta.dir, "..", "data", "chinook.db");

const db = new DAL({
  database: DATABASE_PATH,
});

describe("Query Interface", () => {
  test(".Find", async () => {
    const items = db
      .In("artists")
      .Find({
        name: { $glob: "A*" },
      })
      .Limit(10)
      .Rows();

    for await (const item of items) {
      console.log(item);
    }

    expect(true).toBe(true);
  });

  test(".Find.As", async () => {
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
      .Limit(1)
      .Rows();

    for await (const item of items) {
      console.log(item);
    }
    console.log("done");
    expect(true).toBe(true);
  });
});
