package utils

import (
	"slices"
	"strings"
	"unicode"
)

func IsSQLFunction(str string) bool {
	stopChars := []string{" ", "_", "-", ".", "("}
	isUpper := false
	for _, char := range str {
		if slices.Contains(stopChars, string(char)) {
			break
		}
		if unicode.IsUpper(char) {
			isUpper = true
		} else {
			isUpper = false
			break
		}
	}
	return isUpper
}

func EscapeSingleQuote(str string) string {
	return strings.ReplaceAll(str, "'", "''")
}
