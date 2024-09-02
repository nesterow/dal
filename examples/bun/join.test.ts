import { describe, expect, test } from "bun:test";
import path from "path";
import DAL from "@nesterow/dal/client/libdal";

// in this case we need to use absolute path
const DATABASE_PATH = path.join(import.meta.dir, "..", "data", "chinook.db");

const db = new DAL({
  database: DATABASE_PATH,
});

describe("Query Interface", () => {
  test(".Join [album < artist < tracks (find tracks for all artists whose names start with 'A')]", async () => {
    class Album {
      TrackId = 0;
      TrackName = "";
      ArtistName = "";
      AlbumTitle = "";
    }
    const items = db
      .In("albums al")
      .Join(
        {
          $for: "artists ar",
          $do: {
            "al.ArtistId": "ar.ArtistId",
          },
        },
        {
          $for: "tracks tr",
          $do: {
            "al.AlbumId": "tr.AlbumId",
          },
        },
      )
      .Find({
        "ar.Name": { $glob: "A*" },
      })
      .Fields({
        "tr.TrackId" : "TrackId",
        "tr.Name" : "TrackName",
        "ar.Name" : "ArtistName",
        "al.Title" : "AlbumTitle",
      })
      .Limit(10)
      .As(Album)
      .Rows();

    for await (const item of items) {
      console.log(item);
    }

    expect(true).toBe(true);
  });
});
