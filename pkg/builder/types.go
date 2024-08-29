package builder

import (
	"l12.xyz/x/dal/pkg/adapter"
	"l12.xyz/x/dal/pkg/filters"
)

type RawSql = map[string]interface{}
type CommonDialect = adapter.CommonDialect
type Map = map[string]interface{}
type Fields = Map
type Find = filters.Find
type Query = filters.Find
type Filter = filters.Filter
type Is = filters.Filter
type Dialect = adapter.Dialect
type DialectOpts = adapter.DialectOpts
