package main

import (
	"github.com/nesterow/dal/pkg/facade"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	server := facade.SQLiteServer{}
	server.Init()
	server.Serve()
}
