package main

import "C"
import (
	"strings"

	"github.com/nesterow/dal/pkg/facade"

	_ "github.com/mattn/go-sqlite3"
)

//export InitSQLite
func InitSQLite(pragmas string) {
	pragmasArray := strings.Split(pragmas, ";")
	facade.InitSQLite(pragmasArray)
}

//export HandleQuery
func HandleQuery(input []byte) []byte {
	var out []byte
	facade.HandleQuery(&input, &out)
	return out
}

func main() {}
