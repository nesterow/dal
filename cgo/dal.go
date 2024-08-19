package main

// #include <stdlib.h>
import "C"
import (
	"strings"
	"unsafe"

	"l12.xyz/dal/facade"
)

//export InitSQLite
func InitSQLite(pragmas *C.char) {
	str := C.GoString(pragmas)
	pragmasArray := strings.Split(str, ";")
	facade.InitSQLite(pragmasArray)
}

//export HandleQuery
func HandleQuery(input *C.char) *C.char {
	var in, out []byte
	inPtr := unsafe.Pointer(input)
	defer C.free(inPtr)

	in = C.GoBytes(inPtr, C.int(len(C.GoString(input))))
	facade.HandleQuery(&in, &out)
	output := C.CString(string(out))
	return output
}

func main() {}
