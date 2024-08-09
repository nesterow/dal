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
	result := convertConflict(ctx, "a", "b", "tb.c")

	if result != `ON CONFLICT (t.a,t.b,tb.c)` {
		t.Errorf(`Expected "ON CONFLICT (t.a,t.b,tb.c)", got %s`, result)
	}
}
