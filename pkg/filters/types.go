package filters

import "l12.xyz/dal/adapter"

type DialectOpts = adapter.DialectOpts
type Dialect = adapter.Dialect

type IFilter interface {
	ToSQLPart(ctx Dialect) string
	FromJSON(interface{}) IFilter
}

type Find map[string]interface{}

// Filter{ "$eq": 1 }
type Filter map[string]interface{}
