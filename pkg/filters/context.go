package filters

import (
	"slices"
	"strconv"
	"strings"
	"unicode"

	utils "l12.xyz/dal/utils"
)

type SQLiteContext struct {
	TableAlias string
	FieldName  string
}

func (c SQLiteContext) New(opts map[string]string) Context {
	ta := opts["TableAlias"]
	if ta == "" {
		ta = c.TableAlias
	}
	fn := opts["FieldName"]
	if fn == "" {
		fn = c.FieldName
	}
	return SQLiteContext{
		TableAlias: ta,
		FieldName:  fn,
	}
}

func (c SQLiteContext) GetFieldName() string {
	if strings.Contains(c.FieldName, ".") {
		return c.FieldName
	}
	if c.TableAlias != "" {
		return c.TableAlias + "." + c.FieldName
	}
	return c.FieldName
}

func (c SQLiteContext) NormalizeValue(value interface{}) interface{} {
	str, ok := value.(string)
	if isSQLFunction(str) {
		return str
	}
	if strings.Contains(str, ".") {
		_, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return value
		}
	}
	if !ok {
		return value
	}
	val, err := utils.EscapeSQL(str)
	if err != nil {
		return str
	}
	return "'" + escapeSingleQuote(string(val)) + "'"
}

func isSQLFunction(str string) bool {
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

func escapeSingleQuote(str string) string {
	return strings.ReplaceAll(str, "'", "''")
}
