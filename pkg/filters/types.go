package filters

type Context interface {
	GetFieldName() string
	NormalizeValue(interface{}) interface{}
}

type Filter interface {
	ToSQLPart(ctx Context) string
	FromJSON(interface{}) Filter
}
