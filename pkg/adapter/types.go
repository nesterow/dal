package adapter

type Query struct {
	Db         string        `json:"db"`
	Expression string        `json:"expr"`
	Data       []interface{} `json:"data"`
}

type CtxOpts map[string]string

type Context interface {
	New(opts CtxOpts) Context
	GetTableName() string
	GetFieldName() string
	GetColumnName(key string) string
	NormalizeValue(interface{}) interface{}
}
