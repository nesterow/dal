package dal

import (
	"testing"
)

func TestConvertFieldsBool(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result, err := ConvertFields(ctx, []Map{
		{"test": true},
		{"test2": false},
	})
	if err != nil {
		t.Error(err)
	}
	if result != `test` {
		t.Errorf("Expected test, got %s", result)
	}
}

func TestConvertFieldsInt(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result, err := ConvertFields(ctx, []Map{
		{"test": 0},
		{"test2": 1},
	})
	if err != nil {
		t.Error(err)
	}
	if result != `test2` {
		t.Errorf("Expected test, got %s", result)
	}
}

func TestConvertFieldsStr(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result, err := ConvertFields(ctx, []Map{
		{"t.test": "Test"},
		{"SUM(t.test, t.int)": "Sum"},
	})
	if err != nil {
		t.Error(err)
	}
	if result != `t.test AS Test, SUM(t.test, t.int) AS Sum` {
		t.Errorf("Expected test, got %s", result)
	}
}
