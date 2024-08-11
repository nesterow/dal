package builder

import (
	adapter "l12.xyz/dal/adapter"
	filters "l12.xyz/dal/filters"
)

type Map = map[string]interface{}
type Fields = Map
type Find = filters.Find
type Query = filters.Find
type Filter = filters.Filter
type Is = filters.Filter
type Dialect = adapter.Dialect
type DialectOpts = adapter.DialectOpts
