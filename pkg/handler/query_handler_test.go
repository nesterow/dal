package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"pkg/adapter"
	"pkg/proto"
)

func TestQueryHandler(t *testing.T) {
	adapter.RegisterDialect("sqlite3", adapter.CommonDialect{})
	a := adapter.DBAdapter{Type: "sqlite3"}
	db, err := a.Open("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, name BLOB, data TEXT)")
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
	_, err = db.Exec("INSERT INTO test (name, data) VALUES (?,?)", "test", "y")
	if err != nil {
		t.Fatalf("failed to insert data: %v", err)
	}
	data := proto.Request{
		Id: 0,
		Db: "file::memory:?cache=shared",
		Commands: []proto.BuilderMethod{
			{Method: "In", Args: []interface{}{"test t"}},
			{Method: "Find", Args: []interface{}{
				map[string]interface{}{"id": 1},
			}},
			{Method: "Fields", Args: []interface{}{
				map[string]interface{}{
					"id":   1,
					"name": "Name",
					"data": 1,
				},
			}},
		},
	}
	body, _ := data.MarshalMsg(nil)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := QueryHandler(a)
	handler.ServeHTTP(rr, req)
	res, _ := io.ReadAll(rr.Result().Body)
	result := proto.UnmarshalRows(res)
	fmt.Println(result)
}

func TestQueryHandlerInsert(t *testing.T) {
	adapter.RegisterDialect("sqlite3", adapter.CommonDialect{})
	a := adapter.DBAdapter{Type: "sqlite3"}
	db, err := a.Open("file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, name BLOB, data TEXT)")
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
	data := proto.Request{
		Id: 0,
		Db: "file::memory:?cache=shared",
		Commands: []proto.BuilderMethod{
			{Method: "In", Args: []interface{}{"test t"}},
			{Method: "Insert", Args: []interface{}{
				map[string]interface{}{"name": "test", "data": "y"},
			}},
		},
	}
	body, _ := data.MarshalMsg(nil)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := QueryHandler(a)
	handler.ServeHTTP(rr, req)
	res, _ := io.ReadAll(rr.Result().Body)
	result := proto.UnmarshalRows(res)
	fmt.Println(result)
}
