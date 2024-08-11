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

func (f NotIn) ToSQLPart(ctx Dialect) string {
	if f.NotIn == nil {
		return ""
	}

	name := ctx.GetFieldName()
	values := utils.Map(f.NotIn, ctx.NormalizeValue)
	data := make([]string, len(values))
	for i, v := range values {
		data[i] = fmt.Sprintf("%v", v)
	}
	value := strings.Join(data, ", ")
	return fmt.Sprintf("%s NOT IN (%v)", name, value)
}
