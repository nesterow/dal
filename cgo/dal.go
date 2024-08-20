package main

// #include <stdlib.h>
import "C"
import (
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"l12.xyz/dal/facade"
)

//export InitSQLite
func InitSQLite(pragmas *C.char) {
	str := C.GoString(pragmas)
	pragmasArray := strings.Split(str, ";")
	facade.InitSQLite(pragmasArray)
}

//export HandleQuery
func HandleQuery(input *C.char) []byte {
	var in, out []byte
	in = []byte(C.GoString(input))
	facade.HandleQuery(&in, &out)
	return out
}

func main() {}
