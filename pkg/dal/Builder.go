package dal

type SQLParts struct {
	operation string
	selectExp string
	fromExp   string
	fiterExp  string
	joinExp   []string
	groupExp  string
	orderExp  string
	limitExp  string
	updateExp string
	upsertExp string
}

type Builder struct {
	parts SQLParts
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) In(selectExp string) *Builder {
	b.parts.selectExp = selectExp
	return b
}

func (b *Builder) Find(fromExp string) *Builder {
	b.parts.fromExp = fromExp
	return b
}

func (b *Builder) Join(fiterExp string) *Builder {
	b.parts.fiterExp = fiterExp
	return b
}
