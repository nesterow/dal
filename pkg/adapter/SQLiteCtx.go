package adapter

import (
	"strconv"
	"strings"

	utils "l12.xyz/dal/utils"
)

type SQLiteContext struct {
	TableName  string
	TableAlias string
	FieldName  string
}

func (c SQLiteContext) New(opts CtxOpts) Context {
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

func (c SQLiteContext) GetTableName() string {
	return c.TableName
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
