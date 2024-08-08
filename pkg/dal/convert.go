package dal

import (
	"fmt"
	"strings"

	filters "l12.xyz/dal/filters"
)

func CovertFind(find filters.Find, ctx filters.Context) string {
	expressions := []string{}
	for key, value := range find {
		values, err := filters.Convert(ctx.New(map[string]string{
			"FieldName": key,
		}), value)
		expressions = append(expressions, values)
		fmt.Println(err)

	}
	return strings.Join(expressions, " AND ")
}
