package filters

import (
	"fmt"

	"l12.xyz/dal/utils"
)

type Between struct {
	Between []interface{} `json:"$between"`
}

func (f Between) FromJSON(data interface{}) Filter {
	return FromJson[Between](data)
}

func (f Between) ToSQLPart(ctx Context) string {
	if f.Between == nil {
		return ""
	}
	name := ctx.GetFieldName()
	values := utils.Map(f.Between, ctx.NormalizeValue)
	condition := fmt.Sprintf("%v AND %v", values[0], values[1])
	return fmt.Sprintf("%s BETWEEN %v", name, condition)
}
