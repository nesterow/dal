package filters

var FilterRegistry = map[string]Filter{
	"Eq":         &Eq{},
	"Ne":         &Ne{},
	"Gt":         &Gt{},
	"Gte":        &Gte{},
	"Lt":         &Lt{},
	"Lte":        &Lte{},
	"In":         &In{},
	"Between":    &Between{},
	"NotBetween": &NotBetween{},
	"Glob":       &Glob{},
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
