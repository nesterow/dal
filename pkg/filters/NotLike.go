package filters

import "fmt"

type NotLike struct {
	NotLike interface{} `json:"$nlike"`
}

func (f NotLike) FromJSON(data interface{}) IFilter {
	return FromJson[NotLike](data)
}

func (f NotLike) ToSQLPart(ctx Dialect) string {
	if f.NotLike == nil {
		return ""
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.NotLike)
	return fmt.Sprintf("%s NOT LIKE %v ESCAPE '\\'", name, value)
}
