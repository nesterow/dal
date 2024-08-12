package filters

type Gte struct {
	Gte interface{} `json:"$gte"`
}

func (f Gte) FromJSON(data interface{}) IFilter {
	return FromJson[Gte](data)
}

func (f Gte) ToSQLPart(ctx Dialect) (string, Values) {
	if f.Gte == nil {
		return "", nil
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Gte)
	return FmtCompare(">=", name, value)
}
