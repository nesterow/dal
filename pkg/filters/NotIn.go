package filters

import (
	"fmt"
	"strings"

	"l12.xyz/dal/utils"
)

type NotIn struct {
	NotIn []interface{} `json:"$nin"`
}

func (f NotIn) FromJSON(data interface{}) IFilter {
	return FromJson[NotIn](data)
}

func (f NotIn) ToSQLPart(ctx Dialect) (string, Values) {
	if f.NotIn == nil {
		return "", nil
	}

	name := ctx.GetFieldName()
	values := utils.Map(f.NotIn, ctx.NormalizeValue)
	returnValues := make(Values, 0)
	data := make([]string, len(values))
	for i, value := range values {
		val := ValueOrPlaceholder(value).(string)
		data[i] = val
		if val == "?" {
			returnValues = append(returnValues, value)
		}
	}
	value := strings.Join(data, ", ")
	return fmt.Sprintf("%s NOT IN (%v)", name, value), returnValues
}
