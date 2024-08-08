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

func Convert(ctx Context, data interface{}) (string, error) {
	for _, t := range FilterRegistry {
		filter := t.FromJSON(data)
		if reflect.DeepEqual(t, filter) {
			continue
		}
		value := filter.ToSQLPart(ctx)
		if value != "" {
			return value, nil
		}
	}
	value := Eq{Eq: data}.ToSQLPart(ctx)
	return value, nil
}
