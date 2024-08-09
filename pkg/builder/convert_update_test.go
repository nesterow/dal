package builder

import (
	"testing"
)

func TestConvertUpdate(t *testing.T) {
	ctx := SQLiteContext{
		TableName:  "test",
		TableAlias: "t",
		FieldName:  "test",
	}
	result, err := ConvertUpdate(ctx, Map{
		"c": nil,
		"a": 1,
		"b": 2,
	})
	if err != nil {
		t.Error(err)
	}
	if result.Statement != `UPDATE test SET a = ?,b = ?,c = ?` {
		t.Errorf(`Expected "UPDATE test SET a = ?,b = ?,c = ?", got %s`, result.Statement)
	}
}
