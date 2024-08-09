package builder

import (
	"fmt"
	"strings"
)

func ConvertUpsert(keys []string) string {
	set := make([]string, 0)
	for _, key := range keys {
		set = append(set, fmt.Sprintf("%s = EXCLUDED.%s", key, key))
	}
	return fmt.Sprintf(
		"UPDATE SET %s",
		strings.Join(set, ", "),
	)
}
