package adapter

type Query struct {
	Db         string        `json:"db"`
	Expression string        `json:"expr"`
	Data       []interface{} `json:"data"`
}

type DialectOpts map[string]string

/**
* Dialect interface provides general utilities for normalizing values for particular DB.
**/
type Dialect interface {
	New(opts DialectOpts) Dialect
	GetTableName() string
	GetFieldName() string
	GetColumnName(key string) string
	NormalizeValue(interface{}) interface{}
}
