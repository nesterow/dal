package builder

import (
	"fmt"
	"strings"
)

type InsertData struct {
	Statement string
	Values    []interface{}
}

func convertInsert(ctx Dialect, inserts []Map) (InsertData, error) {
	keys := aggregateSortedKeys(inserts)
	posEnum := make([]string, 0)
	for range keys {
		posEnum = append(posEnum, "?")
	}

	placeholder := strings.Join(posEnum, ",")
	positional := []string{}
	values := make([]interface{}, 0)
	for _, insert := range inserts {
		vals := make([]interface{}, 0)
		for _, key := range keys {
			vals = append(vals, insert[key])
		}
		values = append(values, vals...)
		positional = append(
			positional,
			fmt.Sprintf("(%s)", placeholder),
		)
	}

	sfmt := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES %s",
		ctx.GetTableName(),
		strings.Join(keys, ","),
		strings.Join(positional, ","),
	)
	return InsertData{
		Statement: sfmt,
		Values:    values,
	}, nil
}
