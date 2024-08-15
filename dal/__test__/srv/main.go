package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

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
	defer os.Remove("test.sqlite")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove("test.sqlite")
	}()
	db := adapter.DBAdapter{
		Type: "sqlite3",
	}
	mock(db)
	queryHandler := server.QueryHandler(db)
	mux := http.NewServeMux()
	mux.Handle("/", queryHandler)
	fmt.Println("Server running on port 8111")
	http.ListenAndServe(":8111", mux)
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
