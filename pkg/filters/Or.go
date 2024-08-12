package filters

import (
	"fmt"
	"strings"
)

type Or struct {
	Or []string `json:"$or"`
}

func (f Or) ToSQLPart(ctx Dialect) (string, Values) {
	if f.Or == nil {
		return "", nil
	}
	value := strings.Join(f.Or, " OR ")
	return fmt.Sprintf("(%s)", value), nil
}

func (a Or) FromJSON(data interface{}) IFilter {
	return FromJson[Or](data)
}
