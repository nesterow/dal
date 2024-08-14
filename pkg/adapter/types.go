package adapter

type Query struct {
	Db         string        `json:"db"`
	Expression string        `json:"expr"`
	Data       []interface{} `json:"data"`
}

type DialectOpts map[string]string

type Dialect interface {
	New(opts DialectOpts) Dialect
	GetTableName() string
	GetFieldName() string
	GetColumnName(key string) string
	NormalizeValue(interface{}) interface{}
}

var DIALECTS = map[string]Dialect{
	"sqlite3": SQLite{},
	"sqlite":  SQLite{},
}
