package builder

import (
	"testing"
)

func TestConvertSort(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result, err := ConvertSort(ctx, Map{
		"a": -1,
		"c": "desc",
		"b": 1,
		"d": nil,
	})
	if err != nil {
		t.Error(err)
	}
	if result != `ORDER BY t.a DESC, t.b ASC, t.c DESC, t.d` {
		t.Errorf("Expected ORDER BY t.a DESC, t.b ASC, t.c DESC, t.d, got %s", result)
	}
}
