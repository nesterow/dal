package builder

import "strings"

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
	updateExp string
	upsertExp string
}

type Builder struct {
	Parts      SQLParts
	TableName  string
	TableAlias string
	Ctx        Context
}

func New(ctx Context) *Builder {
	return &Builder{
		Parts: SQLParts{
			Operation: "SELECT",
			From:      "FROM",
		},
		Ctx: ctx,
	}
}

func (b *Builder) In(table string) *Builder {
	b.TableName, b.TableAlias = getTableAlias(table)
	b.Parts.FromExp = table
	b.Ctx = b.Ctx.New(CtxOpts{
		"TableName":  b.TableName,
		"TableAlias": b.TableAlias,
	})
	return b
}

func (b *Builder) Find(query Find) *Builder {
	b.Parts.FiterExp = covertFind(
		b.Ctx,
		query,
	)
	if b.Parts.Operation == "" {
		b.Parts.Operation = "SELECT"
	}
	if b.Parts.HavingExp == "" {
		b.Parts.HavingExp = "WHERE"
	}
	if b.Parts.FieldsExp == "" {
		b.Parts.FieldsExp = "*"
	}
	return b
}

func (b *Builder) Join(joins ...interface{}) *Builder {
	b.Parts.JoinExps = convertJoin(b.Ctx, joins...)
	return b
}

func (b *Builder) Sql() string {
	operation := b.Parts.Operation
	switch {
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
		}, " "))
	default:
		return ""
	}
}
