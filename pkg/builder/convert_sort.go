package builder

import (
	"fmt"
	"strings"
)

func convertSort(ctx Dialect, sort Map) (string, error) {
	if sort == nil {
		return "", nil
	}
	keys := aggregateSortedKeys([]Map{sort})
	expressions := make([]string, 0)
	for _, key := range keys {
		name := ctx.GetColumnName(key)
		order := normalize_order(sort[key])
		if order != "" {
			order = " " + order
		}
		expressions = append(expressions, name+order)
	}
	return fmt.Sprintf("ORDER BY %s", strings.Join(expressions, ", ")), nil
}

func normalize_order(order interface{}) string {
	if order == nil {
		return ""
	}
	orderInt, ok := order.(int)
	if ok {
		if orderInt == 1 {
			return "ASC"
		}
		if orderInt == -1 {
			return "DESC"
		}
	}
	orderStr, ok := order.(string)
	if !ok {
		return ""
	}
	if orderStr == "" {
		return ""
	}
	if orderStr == "1" {
		return "ASC"
	}
	if orderStr == "-1" {
		return "DESC"
	}
	if strings.ToUpper(orderStr) == "ASC" {
		return "ASC"
	}
	if strings.ToUpper(orderStr) == "DESC" {
		return "DESC"
	}
	return ""
}
