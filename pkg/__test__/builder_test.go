package tests

import (
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"l12.xyz/dal/adapter"
	"l12.xyz/dal/builder"
)

func TestBuilderBasic(t *testing.T) {
	a := adapter.DBAdapter{Type: "sqlite3"}
	db, err := a.Open("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, name BLOB)")
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	insert, values := builder.New(adapter.SQLite{}).In("test t").Insert([]builder.Map{
		{"name": "a"},
		{"name": 'b'},
	}).Sql()
	fmt.Println(insert, values)
	_, err = db.Exec(insert, values...)
	if err != nil {
		t.Fatalf("failed to insert data: %v", err)
	}

	expr, _ := builder.New(adapter.SQLite{}).In("test t").Find(builder.Find{"name": builder.Is{"$in": []interface{}{"a", 'b'}}}).Sql()
	fmt.Println(expr)
	rows, err := a.Query(adapter.Query{
		Db:         "file::memory:?cache=shared",
		Expression: expr,
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
