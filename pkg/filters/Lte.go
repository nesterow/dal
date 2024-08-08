package filters

import (
	"fmt"
)

type Lte struct {
	Lte interface{} `json:"$lte"`
}

func (f Lte) FromJSON(data interface{}) IFilter {
	return FromJson[Lte](data)
}

func (f Lte) ToSQLPart(ctx Context) string {
	if f.Lte == nil {
		return ""
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Lte)
	return fmt.Sprintf("%s <= %v", name, value)
}
