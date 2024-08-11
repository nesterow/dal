package filters

import "fmt"

type Glob struct {
	Glob interface{} `json:"$glob"`
}

func (f Glob) FromJSON(data interface{}) IFilter {
	return FromJson[Glob](data)
}

func (f Glob) ToSQLPart(ctx Dialect) string {
	if f.Glob == nil {
		return ""
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Glob)
	return fmt.Sprintf("%s GLOB %v", name, value)
}
