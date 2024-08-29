package main

import (
	_ "github.com/mattn/go-sqlite3"
	"l12.xyz/x/dal/pkg/facade"
)

func main() {
	server := facade.SQLiteServer{}
	server.Init()
	server.Serve()
}
