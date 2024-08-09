package builder

import (
	"fmt"
	"strings"
)

func convertFields(fields []Map) (string, error) {
	var expressions []string
	for _, fieldAssoc := range fields {
		keys := aggregateSortedKeys([]Map{fieldAssoc})
		for _, key := range keys {
			field := key
			as := fieldAssoc[key]
			asBool, ok := as.(bool)
			if ok {
				if asBool {
					expressions = append(expressions, field)
				}
				continue
			}
			asNum, ok := as.(int)
			if ok {
				if asNum == 1 {
					expressions = append(expressions, field)
				}
				continue
			}
			asStr, ok := as.(string)
			if ok {
				expressions = append(expressions, fmt.Sprintf("%s AS %s", field, asStr))
				continue
			}
			return "", fmt.Errorf("invalid field value: %v", as)
		}
	}
	return strings.Join(expressions, ", "), nil
}
