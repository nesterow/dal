package builder

import (
	"fmt"
	"strings"

	filters "l12.xyz/dal/filters"
)

func covertFind(ctx Dialect, find Find) string {
	return covert_find(ctx, find, "")
}

func covert_find(ctx Dialect, find Find, join string) string {
	if join == "" {
		join = " AND "
	}
	keys := aggregateSortedKeys([]Map{find})
	expressions := []string{}
	for _, key := range keys {
		value := find[key]
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
		context := ctx.New(DialectOpts{
			"FieldName": key,
		})
		values, _ := filters.Convert(context, value)
		expressions = append(expressions, values)
	}
	return strings.Join(expressions, join)
}
