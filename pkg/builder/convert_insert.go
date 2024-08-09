package builder

import (
	"fmt"
	"strings"
)

type InsertData struct {
	Statement string
	Values    []interface{}
}

func ConvertInsert(ctx Context, inserts []Map) (InsertData, error) {
	keys := AggregateKeys(inserts)
	placeholder := make([]string, 0)
	for range keys {
		placeholder = append(placeholder, "?")
	}

	values := make([]interface{}, 0)
	for _, insert := range inserts {
		vals := make([]interface{}, 0)
		for _, key := range keys {
			vals = append(vals, insert[key])
		}
		values = append(values, vals)
	}

	sfmt := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)", ctx.GetTableName(),
		strings.Join(keys, ","),
		strings.Join(placeholder, ","),
	)
	return InsertData{
		Statement: sfmt,
		Values:    values,
	}, nil
}
