package dal

import (
	"fmt"
	"strings"

	filters "l12.xyz/dal/filters"
)

func CovertFind(ctx Context, find Find) string {
	return covert_find(ctx, find, "")
}

func covert_find(ctx Context, find Find, join string) string {
	if join == "" {
		join = " AND "
	}
	expressions := []string{}
	for key, value := range find {
		if strings.Contains(key, "$and") {
			v := covert_find(ctx, value.(Find), "")
			expressions = append(expressions, fmt.Sprintf("(%s)", v))
			continue
		}
		if strings.Contains(key, "$or") {
			v := covert_find(ctx, value.(Find), " OR ")
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
