package tests

import (
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"l12.xyz/dal/adapter"
	"l12.xyz/dal/proto"
)

func TestProtoMessagePack(t *testing.T) {
	message, err := os.ReadFile("proto_test.msgpack")
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}
	req := proto.Request{}
	req.UnmarshalMsg(message)
	query, err := req.Parse(adapter.CommonDialect{})
	if err != nil {
		t.Fatalf("failed to parse query: %v", err)
	}
	db := "database.sqlite"
	if query.Db != db {
		t.Fatalf("expected db %s, got %s", db, query.Db)
	}
	expr := "SELECT * FROM data WHERE a = ? AND b > ?"
	if query.Expression != expr {
		t.Fatalf("expected expression %s, got %s", expr, query.Expression)
	}
	//fmt.Println(q)
}
