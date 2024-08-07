package filters

import (
	"fmt"
)

type And struct {
	And []interface{} `json:"$and"`
}

func (a And) ToSQLPart(ctx Context) string {

	fmt.Println(ctx, a)
	return ""
}

func (a And) FromJSON(data interface{}) Filter {
	return FromJson[And](data)
}
