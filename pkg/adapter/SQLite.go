package adapter

import (
	"strconv"
	"strings"

	utils "l12.xyz/dal/utils"
)

type SQLite struct {
	TableName  string
	TableAlias string
	FieldName  string
}

func (c SQLite) New(opts DialectOpts) Dialect {
	tn := opts["TableName"]
	if tn == "" {
		tn = c.TableName
	}
	ta := opts["TableAlias"]
	if ta == "" {
		ta = c.TableAlias
	}
	fn := opts["FieldName"]
	if fn == "" {
		fn = c.FieldName
	}
	return SQLite{
		TableName:  tn,
		TableAlias: ta,
		FieldName:  fn,
	}
}

func (c SQLite) GetTableName() string {
	return c.TableName
}

func (c SQLite) GetFieldName() string {
	if strings.Contains(c.FieldName, ".") {
		return c.FieldName
	}
	if c.TableAlias != "" {
		return c.TableAlias + "." + c.FieldName
	}
	return c.FieldName
}

func (c SQLite) GetColumnName(key string) string {
	if strings.Contains(key, ".") {
		return key
	}
	if c.TableAlias != "" {
		return c.TableAlias + "." + key
	}
	return key
}

func (c SQLite) NormalizeValue(value interface{}) interface{} {
	str, ok := value.(string)
	if utils.IsSQLFunction(str) {
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
	return "'" + utils.EscapeSingleQuote(string(val)) + "'"
}
