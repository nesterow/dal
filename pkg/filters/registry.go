package filters

import (
	"reflect"
)

var FilterRegistry = map[string]IFilter{
	"And":        &And{},
	"Or":         &Or{},
	"Eq":         &Eq{},
	"Ne":         &Ne{},
	"Gt":         &Gt{},
	"Gte":        &Gte{},
	"Lt":         &Lt{},
	"Lte":        &Lte{},
	"In":         &In{},
	"Nin":        &NotIn{},
	"Between":    &Between{},
	"NotBetween": &NotBetween{},
	"Glob":       &Glob{},
	"Like":       &Like{},
	"NotLike":    &NotLike{},
}

func Convert(ctx Dialect, data interface{}) (string, []interface{}) {
	for _, impl := range FilterRegistry {
		filter := impl.FromJSON(data)
		if reflect.DeepEqual(impl, filter) {
			continue
		}
		sfmt, values := filter.ToSQLPart(ctx)
		if sfmt != "" {
			return sfmt, values
		}
	}
	return Eq{Eq: data}.ToSQLPart(ctx)
}
