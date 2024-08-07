package filters

import (
	"fmt"
)

type Eq struct {
	Eq interface{} `json:"$eq"`
}

func (f Eq) FromJSON(data interface{}) Filter {
	return FromJson[Eq](data)
}

func (f Eq) ToSQLPart(ctx Context) string {
	if f.Eq == nil {
		return ""
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Eq)
	if value == "NULL" {
		return fmt.Sprintf("%s IS NULL", name)
	}
	return fmt.Sprintf("%s = %v", name, value)
}
