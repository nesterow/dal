package adapter

import "fmt"

var DIALECTS = map[string]Dialect{
	"sqlite3": CommonDialect{},
}

/**
 * Register a new dialect for a given driver name.
 * `driverName` is the valid name of the db driver (e.g. "sqlite3", "postgres").
 * `dialect` is an implementation of the Dialect interface.
**/
func RegisterDialect(driverName string, dialect Dialect) {
	DIALECTS[driverName] = dialect
}

func GetDialect(driverName string) Dialect {
	dialect, ok := DIALECTS[driverName]
	if !ok {
		panic(fmt.Errorf("db driver %s not found", driverName))
	}
	return dialect
}
