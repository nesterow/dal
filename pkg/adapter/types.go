package adapter

type CtxOpts map[string]string

type Context interface {
	New(opts CtxOpts) Context
	GetTableName() string
	GetFieldName() string
	NormalizeValue(interface{}) interface{}
}
