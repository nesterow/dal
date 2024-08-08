package filters

import "fmt"

type Ne struct {
	Ne interface{} `json:"$ne"`
}

func (f Ne) FromJSON(data interface{}) IFilter {
	return FromJson[Ne](data)
}

func (f Ne) ToSQLPart(ctx Context) string {
	if f.Ne == nil {
		return ""
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Ne)
	if value == "NULL" {
		return fmt.Sprintf("%s IS NOT NULL", name)
	}
	return fmt.Sprintf("%s != %v", name, value)
}
