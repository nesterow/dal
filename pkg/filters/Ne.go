package filters

type Ne struct {
	Ne interface{} `json:"$ne"`
}

func (f Ne) FromJSON(data interface{}) IFilter {
	return FromJson[Ne](data)
}

func (f Ne) ToSQLPart(ctx Dialect) (string, Values) {
	if f.Ne == nil {
		return "", nil
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Ne)
	return FmtCompare("!=", name, value)
}
