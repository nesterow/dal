package filters

import (
	"fmt"

	"github.com/nesterow/dal/pkg/utils"
)

type NotBetween struct {
	NotBetween []interface{} `json:"$nbetween"`
}

func (f NotBetween) FromJSON(data interface{}) IFilter {
	return FromJson[NotBetween](data)
}

func (f NotBetween) ToSQLPart(ctx Dialect) (string, Values) {
	if f.NotBetween == nil {
		return "", nil
	}
	name := ctx.GetFieldName()
	values := utils.Map(f.NotBetween, ctx.NormalizeValue)
	placeholders := utils.Map(values, ValueOrPlaceholder)
	condition := fmt.Sprintf("%s AND %s", placeholders[0], placeholders[1])
	return fmt.Sprintf("%s NOT BETWEEN %v", name, condition), values
}
