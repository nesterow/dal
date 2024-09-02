package builder

import (
	"fmt"
	"strings"
)

const (
	BUILDER_VERSION        = "0.0.1"
	BUILDER_CLIENT_METHODS = "Raw|In|Find|Select|Fields|Join|Group|Sort|Limit|Offset|Delete|Insert|Set|Update|OnConflict|DoUpdate|DoNothing"
	BUILDER_SERVER_METHODS = "Sql"
)

type Builder struct {
	TableName   string
	TableAlias  string
	Parts       SQLParts
	Dialect     Dialect
	LastQuery   Find
	Transaction bool
	RawSql      RawSql
}

type SQLParts struct {
	Operation string
	From      string
	FieldsExp string
	FromExp   string
	HavingExp string
	FiterExp  string
	JoinExps  []string
	GroupExp  string
	OrderExp  string
	LimitExp  string
	OffsetExp string
	Values    []interface{}
	Insert    InsertData
	Update    UpdateData
}

func New(dialect Dialect) *Builder {
	return &Builder{
		Parts: SQLParts{
			Operation: "SELECT",
			From:      "FROM",
		},
		Dialect: dialect,
	}
}

func (b *Builder) Raw(sql map[string]interface{}) *Builder {
	b.Parts.Operation = "RAW"
	b.RawSql = sql
	return b
}

func (b *Builder) In(table string) *Builder {
	b.TableName, b.TableAlias = getTableAlias(table)
	b.Parts.FromExp = table
	b.Dialect = b.Dialect.New(DialectOpts{
		"TableName":  b.TableName,
		"TableAlias": b.TableAlias,
	})
	return b
}

func (b *Builder) Find(query Find) *Builder {
	b.LastQuery = query
	b.Parts.FiterExp, b.Parts.Values = covertFind(
		b.Dialect,
		query,
	)
	if len(b.Parts.Operation) == 0 {
		b.Parts.Operation = "SELECT"
	}
	if len(b.Parts.HavingExp) == 0 {
		b.Parts.HavingExp = "WHERE"
	}
	if len(b.Parts.FieldsExp) == 0 {
		b.Parts.FieldsExp = "*"
	}
	if len(b.Parts.FiterExp) == 0 {
		b.Parts.HavingExp = ""
	}
	return b
}

func (b *Builder) Select(fields ...Map) *Builder {
	fieldsExp, err := convertFields(fields)
	if err != nil {
		return b
	}
	b.Parts.FieldsExp = fieldsExp
	return b
}

func (b *Builder) Fields(fields ...Map) *Builder {
	return b.Select(fields...)
}

func (b *Builder) Join(joins ...interface{}) *Builder {
	exps, vals := convertJoin(b.Dialect, joins...)
	b.Parts.JoinExps = append(b.Parts.JoinExps, exps...)
	b.Parts.Values = append(b.Parts.Values, vals...)
	return b
}

func (b *Builder) Group(keys ...string) *Builder {
	b.Parts.HavingExp = "HAVING"
	b.Parts.GroupExp = convertGroup(b.Dialect, keys)
	return b
}

func (b *Builder) Sort(sort Map) *Builder {
	b.Parts.OrderExp, _ = convertSort(b.Dialect, sort)
	return b
}

func (b *Builder) Limit(limit int64) *Builder {
	b.Parts.LimitExp = fmt.Sprintf("LIMIT %d", limit)
	return b
}

func (b *Builder) Offset(offset int64) *Builder {
	b.Parts.OffsetExp = fmt.Sprintf("OFFSET %d", offset)
	return b
}

func (b *Builder) Delete() *Builder {
	b.Parts.Operation = "DELETE"
	return b
}

func (b *Builder) Insert(inserts ...Map) *Builder {
	insertData, _ := convertInsert(b.Dialect, inserts)
	b.Parts = SQLParts{
		Operation: "INSERT INTO",
		Insert:    insertData,
	}
	return b
}

func (b *Builder) Set(updates Map) *Builder {
	b.Dialect = b.Dialect.New(DialectOpts{
		"TableAlias": "",
	})
	updateData := convertUpdate(b.Dialect, updates)
	b.Find(b.LastQuery)
	b.Parts.Operation = "UPDATE"
	b.Parts.Update = updateData
	return b
}

func (b *Builder) Update(updates Map) *Builder {
	return b.Set(updates)
}

func (b *Builder) OnConflict(fields ...string) *Builder {
	if b.Parts.Operation == "UPDATE" {
		b.Parts.Update.Upsert = convertConflict(b.Dialect, fields...)
		b.Parts.Update.UpsertExp = "DO NOTHING"
	} else {
		panic("OnConflict is only available for UPDATE operation")
	}
	return b
}

func (b *Builder) DoUpdate(fields ...string) *Builder {
	if b.Parts.Operation == "UPDATE" {
		b.Parts.Update.UpsertExp = convertUpsert(fields)
	} else {
		panic("DoUpdate is only available for UPDATE operation")
	}
	return b
}

func (b *Builder) DoNothing() *Builder {
	if b.Parts.Operation == "UPDATE" {
		b.Parts.Update.UpsertExp = "DO NOTHING"
	} else {
		panic("DoNothing is only available for UPDATE operation")
	}
	return b
}

func (b *Builder) Tx() *Builder {
	b.Transaction = true
	return b
}

func (b *Builder) Sql() (string, []interface{}) {
	operation := b.Parts.Operation
	switch {
	case operation == "RAW":
		return b.RawSql["s"].(string), b.RawSql["v"].([]interface{})
	case operation == "SELECT" || operation == "SELECT DISTINCT":
		return unspace(strings.Join([]string{
			b.Parts.Operation,
			b.Parts.FieldsExp,
			b.Parts.From,
			b.Parts.FromExp,
			strings.Join(
				b.Parts.JoinExps,
				" ",
			),
			b.Parts.GroupExp,
			b.Parts.HavingExp,
			b.Parts.FiterExp,
			b.Parts.OrderExp,
			b.Parts.LimitExp,
			b.Parts.OffsetExp,
		}, " ")), b.Parts.Values
	case operation == "DELETE":
		return unspace(strings.Join([]string{
			b.Parts.Operation,
			b.Parts.From,
			b.Parts.FromExp,
			b.Parts.HavingExp,
			b.Parts.FiterExp,
			b.Parts.OrderExp,
			b.Parts.LimitExp,
			b.Parts.OffsetExp,
		}, " ")), b.Parts.Values
	case operation == "INSERT INTO":
		return b.Parts.Insert.Statement, b.Parts.Insert.Values
	case operation == "UPDATE":
		values := append(b.Parts.Update.Values, b.Parts.Values...)
		return unspace(strings.Join([]string{
			b.Parts.Update.Statement,
			b.Parts.HavingExp,
			b.Parts.FiterExp,
			b.Parts.OrderExp,
			b.Parts.LimitExp,
			b.Parts.OffsetExp,
			b.Parts.Update.Upsert,
			b.Parts.Update.UpsertExp,
		}, " ")), values
	default:
		return "", nil
	}
}
