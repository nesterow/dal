package filters

import "github.com/nesterow/dal/pkg/adapter"

type DialectOpts = adapter.DialectOpts
type Dialect = adapter.Dialect
type Values = []interface{}
type IFilter interface {
	ToSQLPart(ctx Dialect) (string, Values)
	FromJSON(interface{}) IFilter
}

type Find map[string]interface{}

// Filter{ "$eq": 1 }
type Filter map[string]interface{}
