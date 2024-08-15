package server

import (
	"fmt"
	"io"
	"net/http"
	"reflect"

	"l12.xyz/dal/adapter"
	"l12.xyz/dal/proto"
)

func QueryHandler(db adapter.DBAdapter) http.Handler {
	dialect, ok := adapter.DIALECTS[db.Type]
	if !ok {
		panic(fmt.Errorf("dialect %s not found", db.Type))
	}
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
