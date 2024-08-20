package builder

import (
	"fmt"
	"strings"

	filters "pkg/filters"
)

type Values = []interface{}

func covertFind(ctx Dialect, find Find) (string, Values) {
	return covert_find(ctx, find, "")
}

func covert_find(ctx Dialect, find Find, join string) (string, Values) {
	if join == "" {
		join = " AND "
	}
	keys := aggregateSortedKeys([]Map{find})
	expressions := []string{}
	values := Values{}
	for _, key := range keys {
		value := find[key]
		if strings.Contains(key, "$and") {
			exp, vals := covert_find(ctx, value.(Find), "")
			values = append(values, vals...)
			expressions = append(expressions, fmt.Sprintf("(%s)", exp))
			continue
		}
		if strings.Contains(key, "$or") {
			exp, vals := covert_find(ctx, value.(Find), " OR ")
			values = append(values, vals...)
			expressions = append(expressions, fmt.Sprintf("(%s)", exp))
			continue
		}
		context := ctx.New(DialectOpts{
			"FieldName": key,
		})
		expr, vals := filters.Convert(context, value)
		values = append(values, vals...)
		expressions = append(expressions, expr)
	}
	return strings.Join(expressions, join), values
}
