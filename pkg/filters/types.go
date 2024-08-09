package filters

import "l12.xyz/dal/adapter"

type CtxOpts = adapter.CtxOpts
type Context = adapter.Context

type IFilter interface {
	ToSQLPart(ctx Context) string
	FromJSON(interface{}) IFilter
}

type Find map[string]interface{}

// Filter{ "$eq": 1 }
type Filter map[string]interface{}
