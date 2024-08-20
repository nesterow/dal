package builder

import (
	adapter "github.com/nesterow/dal/pkg/adapter"
	filters "github.com/nesterow/dal/pkg/filters"
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
