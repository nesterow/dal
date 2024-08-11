package filters

import (
	"fmt"
	"strings"

	"l12.xyz/dal/utils"
)

type In struct {
	In []interface{} `json:"$in"`
}

func (f In) FromJSON(data interface{}) IFilter {
	return FromJson[In](data)
}

func (f In) ToSQLPart(ctx Dialect) string {
	if f.In == nil {
		return ""
	}

	name := ctx.GetFieldName()
	values := utils.Map(f.In, ctx.NormalizeValue)
	data := make([]string, len(values))
	for i, v := range values {
		data[i] = fmt.Sprintf("%v", v)
	}
	value := strings.Join(data, ", ")
	return fmt.Sprintf("%s IN (%v)", name, value)
}
