package builder

import (
	"fmt"
	"strings"
)

type UpdateData struct {
	Statement string
	Values    []interface{}
}

func ConvertUpdate(ctx Context, updates Map) (UpdateData, error) {
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
	}, nil
}
