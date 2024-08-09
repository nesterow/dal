package builder

import (
	"testing"
)

func TestConvertConflict(t *testing.T) {
	ctx := SQLiteContext{
		TableName:  "test",
		TableAlias: "t",
		FieldName:  "test",
	}
	result := ConvertConflict(ctx, "a", "b", "tb.c")

	if result != `ON CONFLICT (t.a,t.b,tb.c) DO` {
		t.Errorf(`Expected "ON CONFLICT (t.a,t.b,tb.c) DO", got %s`, result)
	}
}
