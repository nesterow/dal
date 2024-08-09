package adapter

type CtxOpts map[string]string

type Context interface {
	New(opts CtxOpts) Context
	GetTableName() string
	GetFieldName() string
	GetColumnName(key string) string
	NormalizeValue(interface{}) interface{}
}
