package filters

var FilterRegistry = map[string]Filter{
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

func Convert(ctx Context, json interface{}) string {
	for _, t := range FilterRegistry {
		value := t.FromJSON(json).ToSQLPart(ctx)
		if value != "" {
			return value
		}
	}
	return ""
}
