package filters

import (
	"fmt"
	"strings"

	"pkg/utils"
)

type In struct {
	In []interface{} `json:"$in"`
}

func (f In) FromJSON(data interface{}) IFilter {
	return FromJson[In](data)
}

func (f In) ToSQLPart(ctx Dialect) (string, Values) {
	if f.In == nil {
		return "", nil
	}

	name := ctx.GetFieldName()
	values := utils.Map(f.In, ctx.NormalizeValue)
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
	return fmt.Sprintf("%s IN (%v)", name, value), returnValues
}
