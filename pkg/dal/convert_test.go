package dal

import (
	"testing"

	f "l12.xyz/dal/filters"
)

func TestConvertFind(t *testing.T) {
	find := f.Find{
		"test":  "1",
		"test2": "2",
		"test3": f.Filter{
			"$ne": 1,
		},
	}
	ctx := f.SQLiteContext{
		TableAlias: "t",
	}
	result := CovertFind(find, ctx)
	if result != `t.test = '1'` {
		t.Errorf("Expected t.test = '1', got %s", result)
	}
}
