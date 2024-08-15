package adapter

import (
	"strconv"
	"strings"

	utils "l12.xyz/dal/utils"
)

/*
CommonDialect is a simple implementation of the Dialect interface.
Should be usable for most SQL databases.
*/
type CommonDialect struct {
	TableName  string
	TableAlias string
	FieldName  string
}

func (c CommonDialect) New(opts DialectOpts) Dialect {
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
	return CommonDialect{
		TableName:  tn,
		TableAlias: ta,
		FieldName:  fn,
	}
}

func (c CommonDialect) GetTableName() string {
	return c.TableName
}

func (c CommonDialect) GetFieldName() string {
	if strings.Contains(c.FieldName, ".") {
		return c.FieldName
	}
	if c.TableAlias != "" {
		return c.TableAlias + "." + c.FieldName
	}
	return c.FieldName
}

func (c CommonDialect) GetColumnName(key string) string {
	if strings.Contains(key, ".") {
		return key
	}
	if c.TableAlias != "" {
		return c.TableAlias + "." + key
	}
	return key
}

func (c CommonDialect) NormalizeValue(value interface{}) interface{} {
	str, isStr := value.(string)
	if !isStr {
		return value
	}
	if str == "?" {
		return str
	}
	if utils.IsSQLFunction(str) {
		return str
	}
	if strings.Contains(str, ".") {
		_, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return value
		}
	}
	val, err := utils.EscapeSQL(str)
	if err != nil {
		return str
	}
	return string(val)
}