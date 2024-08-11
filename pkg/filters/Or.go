package filters

import (
	"fmt"
	"strings"
)

type Or struct {
	Or []string `json:"$or"`
}

func (f Or) ToSQLPart(ctx Dialect) string {
	if f.Or == nil {
		return ""
	}
	value := strings.Join(f.Or, " OR ")
	return fmt.Sprintf("(%s)", value)
}

func (a Or) FromJSON(data interface{}) IFilter {
	return FromJson[Or](data)
}
