package filters

import (
	"fmt"
)

type Gt struct {
	Gt interface{} `json:"$gt"`
}

func (f Gt) FromJSON(data interface{}) IFilter {
	return FromJson[Gt](data)
}

func (f Gt) ToSQLPart(ctx Context) string {
	if f.Gt == nil {
		return ""
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Gt)
	return fmt.Sprintf("%s > %v", name, value)
}
