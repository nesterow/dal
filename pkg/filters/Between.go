package filters

import (
	"fmt"

	"l12.xyz/x/dal/pkg/utils"
)

type Between struct {
	Between []interface{} `json:"$between"`
}

func (f Between) FromJSON(data interface{}) IFilter {
	return FromJson[Between](data)
}

func (f Between) ToSQLPart(ctx Dialect) (string, Values) {
	if f.Between == nil {
		return "", nil
	}
	name := ctx.GetFieldName()
	values := utils.Map(f.Between, ctx.NormalizeValue)
	placeholders := utils.Map(values, ValueOrPlaceholder)
	condition := fmt.Sprintf("%s AND %s", placeholders[0], placeholders[1])
	return fmt.Sprintf("%s BETWEEN %v", name, condition), values
}
