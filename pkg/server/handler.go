package server

import (
	"fmt"
	"io"
	"net/http"

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

		fmt.Println(query, "QueryHandler")
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/x-msgpack")
		defer rows.Close()
		for rows.Next() {
			row := []byte{}
			rows.Scan(row)
			w.Write(row)
		}
	})
}
