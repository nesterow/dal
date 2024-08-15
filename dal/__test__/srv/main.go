package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"l12.xyz/dal/adapter"
	"l12.xyz/dal/server"
)

func mock(adapter adapter.DBAdapter) {
	db, _ := adapter.Open("test.sqlite")
	defer db.Close()
	db.Exec("CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name BLOB, data TEXT)")
	db.Exec("INSERT INTO test (name, data) VALUES (?,?)", "test", "y")
	db.Exec("INSERT INTO test (name, data) VALUES (?,?)", "tost", "x")
	db.Exec("INSERT INTO test (name, data) VALUES (?,?)", "foo", "a")
	db.Exec("INSERT INTO test (name, data) VALUES (?,?)", "bar", "b")
}

func main() {
	db := adapter.DBAdapter{
		Type: "sqlite3",
	}
	mock(db)
	queryHandler := server.QueryHandler(db)
	mux := http.NewServeMux()
	mux.Handle("/", queryHandler)
	fmt.Println("Server running on port 8111")
	http.ListenAndServe(":8111", mux)
}
