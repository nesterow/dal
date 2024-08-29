package tests

import (
	"fmt"
	"testing"

	"l12.xyz/x/dal/pkg/adapter"

	_ "github.com/mattn/go-sqlite3"
)

func TestAdapterBasic(t *testing.T) {
	a := adapter.DBAdapter{Type: "sqlite3"}
	db, err := a.Open("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
	_, err = db.Exec("INSERT INTO test (name) VALUES (?)", "test")
	if err != nil {
		t.Fatalf("failed to insert data: %v", err)
	}
	rows, err := a.Query(adapter.Query{
		Db:         "file::memory:?cache=shared",
		Expression: "SELECT * FROM test",
		Data:       []interface{}{},
	})
	if err != nil {
		t.Fatalf("failed to query data: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			t.Fatalf("failed to scan row: %v", err)
		}
		fmt.Printf("id: %d, name: %s\n", id, name)
	}
}
