package dal

import (
	"fmt"
	"strings"

	filters "l12.xyz/dal/filters"
)

func CovertFind(find Find, ctx Context) string {
	return covert_find(find, ctx, "")
}

func covert_find(find Find, ctx Context, join string) string {
	if join == "" {
		join = " AND "
	}
	expressions := []string{}
	for key, value := range find {
		if strings.Contains(key, "$and") {
			v := covert_find(value.(Find), ctx, "")
			expressions = append(expressions, fmt.Sprintf("(%s)", v))
			continue
		}
		if strings.Contains(key, "$or") {
			v := covert_find(value.(Find), ctx, " OR ")
			expressions = append(expressions, fmt.Sprintf("(%s)", v))
			continue
		}
		context := ctx.New(CtxOpts{
			"FieldName": key,
		})
		values, _ := filters.Convert(context, value)
		expressions = append(expressions, values)
	}
	return strings.Join(expressions, join)
}
