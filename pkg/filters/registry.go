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

// RegisterFilter registers a new filter for a given name.
// `name` is the name of the filter, and `filter` is an empty instance (&reference) of the IFilter.
func RegisterFilter(name string, filter IFilter) {
	FilterRegistry[name] = filter
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
