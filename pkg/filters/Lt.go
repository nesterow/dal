package filters

import (
	"fmt"
)

type Lt struct {
	Lt interface{} `json:"$lt"`
}

func (f Lt) FromJSON(data interface{}) Filter {
	return FromJson[Lt](data)
}

func (f Lt) ToSQLPart(ctx Context) string {
	if f.Lt == nil {
		return ""
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Lt)
	return fmt.Sprintf("%s < %v", name, value)
}
