package filters

import "fmt"

type Like struct {
	Like interface{} `json:"$like"`
}

func (f Like) FromJSON(data interface{}) IFilter {
	return FromJson[Like](data)
}

func (f Like) ToSQLPart(ctx Dialect) (string, Values) {
	if f.Like == nil {
		return "", nil
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Like)
	return fmt.Sprintf("%s LIKE ? ESCAPE '\\'", name), Values{value}
}
