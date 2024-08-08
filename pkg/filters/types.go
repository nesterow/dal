package filters

type Context interface {
	New(opts map[string]string) Context
	GetFieldName() string
	NormalizeValue(interface{}) interface{}
}

type IFilter interface {
	ToSQLPart(ctx Context) string
	FromJSON(interface{}) IFilter
}

type Find map[string]interface{}

// Filter{ "$eq": 1 }
type Filter map[string]interface{}
