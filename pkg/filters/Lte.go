package filters

type Lte struct {
	Lte interface{} `json:"$lte"`
}

func (f Lte) FromJSON(data interface{}) IFilter {
	return FromJson[Lte](data)
}

func (f Lte) ToSQLPart(ctx Dialect) (string, Values) {
	if f.Lte == nil {
		return "", nil
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Lte)
	return FmtCompare("<=", name, value)
}
