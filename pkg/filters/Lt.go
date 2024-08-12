package filters

type Lt struct {
	Lt interface{} `json:"$lt"`
}

func (f Lt) FromJSON(data interface{}) IFilter {
	return FromJson[Lt](data)
}

func (f Lt) ToSQLPart(ctx Dialect) (string, Values) {
	if f.Lt == nil {
		return "", nil
	}
	name := ctx.GetFieldName()
	value := ctx.NormalizeValue(f.Lt)
	return FmtCompare("<", name, value)
}
