package builder

import (
	"fmt"
	"testing"
)

func TestConvertFind(t *testing.T) {
	find := Find{
		"impl": "1",
		"exp": Is{
			"$gt": 2,
		},
	}
	ctx := SQLiteContext{
		TableAlias: "t",
	}
	result, values := covertFind(ctx, find)
	if values[1] != "1" {
		t.Errorf("Expected '1', got %v", values[1])
	}
	if values[0].(float64) != 2 {
		t.Errorf("Expected 2, got %v", values[0])
	}
	if result == `t.exp > ? AND t.impl = ?` {
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
	result, values := covertFind(ctx, find)
	fmt.Println(values)
	if result == `(t.a > ? AND t.b < ?)` {
		return
	}
	t.Errorf(`Expected "(t.b < ? AND t.a > ?)", got %s`, result)
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
	result, values := covertFind(ctx, find)
	fmt.Println(values)
	if result == `(t.a > ? OR t.b < ?)` {
		return
	}
	t.Errorf(`Expected "(t.b < ? OR t.a > ?)", got %s`, result)
}
