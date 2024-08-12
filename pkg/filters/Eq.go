package filters

import (
	"fmt"
)

type Eq struct {
	Eq interface{} `json:"$eq"`
}

func (f Eq) FromJSON(data interface{}) IFilter {
	return FromJson[Eq](data)
}

func (f Eq) ToSQLPart(ctx Dialect) (string, Values) {
	if f.Eq == nil {
		return "", nil
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Eq)
	if value == "NULL" {
		return fmt.Sprintf("%s IS NULL", ValueOrPlaceholder(name)), Values{name}
	}
	return FmtCompare("=", name, value)
}
