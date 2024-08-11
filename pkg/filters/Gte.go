package filters

import (
	"fmt"
)

type Gte struct {
	Gte interface{} `json:"$gte"`
}

func (f Gte) FromJSON(data interface{}) IFilter {
	return FromJson[Gte](data)
}

func (f Gte) ToSQLPart(ctx Dialect) string {
	if f.Gte == nil {
		return ""
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Gte)
	return fmt.Sprintf("%s >= %v", name, value)
}
