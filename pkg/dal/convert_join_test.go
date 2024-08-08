package dal

import (
	"testing"

	f "l12.xyz/dal/filters"
)

func TestJoin(t *testing.T) {
	j := Join{
		For: "artist a",
		Do: Find{
			"a.impl": "t.impl",
		},
		As: "LEFT",
	}
	ctx := f.SQLiteContext{
		TableAlias: "t",
	}
	result := j.Convert(ctx)
	if result == `LEFT JOIN artist a ON a.impl = t.impl` {
		return
	}
	t.Errorf(`Expected "LEFT JOIN artist a ON a.impl = t.impl", got %s`, result)
}
