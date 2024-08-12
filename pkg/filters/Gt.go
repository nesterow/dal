package filters

type Gt struct {
	Gt interface{} `json:"$gt"`
}

func (f Gt) FromJSON(data interface{}) IFilter {
	return FromJson[Gt](data)
}

func (f Gt) ToSQLPart(ctx Dialect) (string, Values) {
	if f.Gt == nil {
		return "", nil
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Gt)
	return FmtCompare(">", name, value)
}
