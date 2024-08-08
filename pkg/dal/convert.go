package dal

import (
	"strings"

	filters "l12.xyz/dal/filters"
)

func CovertFind(find Find, ctx Context) string {
	expressions := []string{}
	for key, value := range find {
		context := ctx.New(CtxOpts{
			"FieldName": key,
		})
		values, _ := filters.Convert(context, value)
		expressions = append(expressions, values)
	}
	return strings.Join(expressions, " AND ")
}
