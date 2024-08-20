package tests

import (
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nesterow/dal/pkg/adapter"
	"github.com/nesterow/dal/pkg/builder"
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

	inserts := []builder.Map{
		{"name": "a"},
		{"name": 'b'},
	}
	insert, values := builder.New(adapter.CommonDialect{}).In("test t").Insert(inserts...).Sql()

	_, err = db.Exec(insert, values...)
	if err != nil {
		t.Fatalf("failed to insert data: %v", err)
	}

	expr, values := builder.New(adapter.CommonDialect{}).In("test t").Find(builder.Find{"name": builder.Is{"$in": []interface{}{"a", 98}}}).Sql()
	fmt.Println(expr)
	rows, err := a.Query(adapter.Query{
		Db:         "file::memory:?cache=shared",
		Expression: expr,
		Data:       values,
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

func TestBuilderSet(t *testing.T) {
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

	inserts := []builder.Map{
		{"name": "a"},
		{"name": 'b'},
	}
	insert, values := builder.New(adapter.CommonDialect{}).In("test t").Insert(inserts...).Sql()

	_, err = db.Exec(insert, values...)
	if err != nil {
		t.Fatalf("failed to insert data: %v", err)
	}

	b := builder.New(adapter.CommonDialect{}).In("test t")
	b.Find(builder.Find{"id": builder.Is{"$eq": 2}})
	b.Set(builder.Map{"name": "c"})
	b.Tx()
	expr, values := b.Sql()
	fmt.Println(expr, values)
	_, err = a.Exec(adapter.Query{
		Db:          "file::memory:?cache=shared",
		Expression:  expr,
		Data:        values,
		Transaction: b.Transaction,
	})
	if err != nil {
		t.Fatalf("failed to query data: %v", err)
	}
	b = builder.New(adapter.CommonDialect{}).In("test t")
	b.Find(builder.Find{})
	expr, values = b.Sql()
	fmt.Println(expr, values)
	rows, _ := a.Query(adapter.Query{
		Db:          "file::memory:?cache=shared",
		Expression:  expr,
		Data:        values,
		Transaction: b.Transaction,
	})
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
