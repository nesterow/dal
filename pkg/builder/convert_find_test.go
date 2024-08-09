package builder

import (
	"testing"
)

func TestConvertFind(t *testing.T) {
	find := Find{
		"impl": "1",
		"exp": Is{
			"$gt": 1,
		},
	}
	ctx := SQLiteContext{
		TableAlias: "t",
	}
	result := covertFind(ctx, find)
	if result == `t.exp > 1 AND t.impl = '1'` {
		return
	}
	if result == `t.impl = '1' AND t.exp > 1` {
		return
	}
	t.Errorf(`Expected "t.impl = '1' AND t.exp = 1", got %s`, result)
}

func TestConvertFindAnd(t *testing.T) {
	find := Find{
		"$and": Find{
			"a": Is{
				"$gt": 1,
			},
			"b": Is{
				"$lt": 10,
			},
		},
	}
	ctx := SQLiteContext{
		TableAlias: "t",
	}
	result := covertFind(ctx, find)
	if result == `(t.a > 1 AND t.b < 10)` {
		return
	}
	if result == `(t.b < 10 AND t.a > 1)` {
		return
	}
	t.Errorf(`Expected "(t.b < 10 AND t.a > 1)", got %s`, result)
}

func TestConvertFindOr(t *testing.T) {
	find := Query{
		"$or": Query{
			"a": Is{
				"$gt": 1,
			},
			"b": Is{
				"$lt": 10,
			},
		},
	}
	ctx := SQLiteContext{
		TableAlias: "t",
	}
	result := covertFind(ctx, find)
	if result == `(t.a > 1 OR t.b < 10)` {
		return
	}
	if result == `(t.b < 10 OR t.a > 1)` {
		return
	}
	t.Errorf(`Expected "(t.b < 10 OR t.a > 1)", got %s`, result)
}
