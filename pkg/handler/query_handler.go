package handler

import (
	"io"
	"net/http"
	"reflect"

	"pkg/adapter"
	"pkg/proto"
)

/*
QueryHandler is a http.Handler that reads a proto.Request from the request body,
parses it into a query, executes the query on the provided db and writes the
result to the response body.
- The request body is expected to be in msgpack format (proto.Request).
- The response body is written in msgpack format.
- The respose is a stream of rows (proto.Row), where the first row is the column names.
- The columns are sorted alphabetically, so it is client's responsibility to match them and sort as needed.
*/
func QueryHandler(db adapter.DBAdapter) http.Handler {
	dialect := adapter.GetDialect(db.Type)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		req := proto.Request{}
		req.UnmarshalMsg(body)

		query, err := req.Parse(dialect)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if query.Exec {
			result, err := db.Exec(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/x-msgpack")
			ra, _ := result.RowsAffected()
			la, _ := result.LastInsertId()
			res := proto.Response{
				Id:           0,
				RowsAffected: ra,
				LastInsertId: la,
			}
			out, _ := res.MarshalMsg(nil)
			w.Write(out)
			return
		}

		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Content-Type", "application/x-msgpack")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "expected http.ResponseWriter to be an http.Flusher", http.StatusInternalServerError)
			return
		}
		columns, _ := rows.Columns()
		types, _ := rows.ColumnTypes()
		cols, _ := proto.MarshalRow(columns)
		w.Write(cols)
		flusher.Flush()

		for rows.Next() {
			data := make([]interface{}, len(columns))
			for i := range data {
				typ := reflect.New(types[i].ScanType()).Interface()
				data[i] = &typ
			}
			rows.Scan(data...)
			cols, _ := proto.MarshalRow(data)
			w.Write(cols)
			flusher.Flush()
		}
	})
}
