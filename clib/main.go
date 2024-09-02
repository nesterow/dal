package main

// #include <stdlib.h>
// #include <stdio.h>
import "C"

import (
	"strings"
	"unsafe"

	"l12.xyz/x/dal/pkg/facade"

	_ "github.com/mattn/go-sqlite3"
)

var (
	iterators = make(map[int]*facade.RowsIter)
	itersize  = make(map[int]C.int)
)

//export InitSQLite
func InitSQLite(params *C.char) {
	pragmas := C.GoString(params)
	pragmasArray := strings.Split(pragmas, ";")
	facade.InitSQLite(pragmasArray)
}

//export CreateRowIterator
func CreateRowIterator(data *C.char, size C.int) C.int {
	var it = &facade.RowsIter{}
	input := C.GoBytes(unsafe.Pointer(data), size)
	it.Exec(input)
	ptr := C.int(len(iterators))
	iterators[len(iterators)] = it
	return ptr
}

//export NextRow
func NextRow(itid C.int) unsafe.Pointer {
	it := iterators[int(itid)]
	if it.Result != nil {
		itersize[int(itid)] = C.int(len(it.Result))
		return C.CBytes(it.Result)
	}
	data := it.Next()
	if data == nil {
		return nil
	}
	itersize[int(itid)] = C.int(len(data))
	res := C.CBytes(data)
	return res
}

//export GetLen
func GetLen(idx C.int) C.int {
	return itersize[int(idx)]
}

//export Cleanup
func Cleanup(itid C.int) {
	it := iterators[int(itid)]
	it.Close()
	delete(iterators, int(itid))
	delete(itersize, int(itid))
}

//export Free
func Free(ptr unsafe.Pointer) {
	C.free(ptr)
}

func main() {}
