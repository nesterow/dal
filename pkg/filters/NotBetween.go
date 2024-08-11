package filters

import (
	"fmt"

	"l12.xyz/dal/utils"
)

type NotBetween struct {
	NotBetween []interface{} `json:"$nbetween"`
}

func (f NotBetween) FromJSON(data interface{}) IFilter {
	return FromJson[NotBetween](data)
}

func (f NotBetween) ToSQLPart(ctx Dialect) string {
	if f.NotBetween == nil {
		return ""
	}
	name := ctx.GetFieldName()
	values := utils.Map(f.NotBetween, ctx.NormalizeValue)
	condition := fmt.Sprintf("%v AND %v", values[0], values[1])
	return fmt.Sprintf("%s NOT BETWEEN %v", name, condition)
}
