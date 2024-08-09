package builder

import (
	"sort"
	"strings"
)

func aggregateSortedKeys(maps []Map) []string {
	set := make(map[string]int)
	keys := make([]string, 0)
	for _, item := range maps {
		for k := range item {
			if set[k] == 1 {
				continue
			}
			keys = append(keys, k)
			set[k] = 1
		}
	}
	set = nil
	sort.Strings(keys)
	return keys
}

func getTableAlias(tableName string) (string, string) {
	if !strings.Contains(tableName, " ") {
		return tableName, ""
	}
	if strings.Contains(strings.ToLower(tableName), " as ") {
		data := strings.Split(strings.ToLower(tableName), " as ")
		return data[0], data[1]
	}
	data := strings.Split(tableName, " ")
	return data[0], data[1]
}

func unspace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
