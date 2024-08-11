package builder

import (
	"fmt"
	"strings"
)

type UpdateData struct {
	Statement string
	Upsert    string
	UpsertExp string
	Values    []interface{}
}

func convertUpdate(ctx Dialect, updates Map) UpdateData {
	keys := aggregateSortedKeys([]Map{updates})
	set := make([]string, 0)
	values := make([]interface{}, 0)
	for _, key := range keys {
		set = append(set, fmt.Sprintf("%s = ?", key))
		values = append(values, updates[key])
	}
	sfmt := fmt.Sprintf(
		"UPDATE %s SET %s", ctx.GetTableName(),
		strings.Join(set, ","),
	)
	return UpdateData{
		Statement: sfmt,
		Values:    values,
	}
}
