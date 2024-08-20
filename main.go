package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/nesterow/dal/pkg/facade"
)

func main() {
	facade.Init()
	facade.Serve()
}
