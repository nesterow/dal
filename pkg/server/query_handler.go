package server

import (
	"io"
	"net/http"
	"reflect"

	"l12.xyz/dal/adapter"
	"l12.xyz/dal/proto"
)

/**
* QueryHandler is a http.Handler that reads a proto.Request from the request body,
* parses it into a query, executes the query on the provided db and writes the
* result to the response body.
* - The request body is expected to be in msgpack format (proto.Request).
* - The response body is written in msgpack format.
* - The respose is a stream of rows (proto.Row), where the first row is the column names.
* - The columns are sorted alphabetically, so it is client's responsibility to match them and sort as needed.
**/
func QueryHandler(db adapter.DBAdapter) http.Handler {
	dialect := adapter.GetDialect(db.Type)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyReader, err := r.GetBody()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		body, err := io.ReadAll(bodyReader)
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

		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		w.Header().Set("Content-Type", "application/x-msgpack")

		columns, _ := rows.Columns()
		types, _ := rows.ColumnTypes()
		cols, _ := proto.MarshalRow(columns)
		w.Write(cols)

		for rows.Next() {
			data := make([]interface{}, len(columns))
			for i := range data {
				typ := reflect.New(types[i].ScanType()).Interface()
				data[i] = &typ
			}
			rows.Scan(data...)
			cols, _ := proto.MarshalRow(data)
			w.Write(cols)
		}
	})
}
