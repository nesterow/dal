package builder

import (
	"testing"
)

func TestConvertUpdate(t *testing.T) {
	ctx := CommonDialect{
		TableName:  "test",
		TableAlias: "t",
		FieldName:  "test",
	}
	result := convertUpdate(ctx, Map{
		"c": nil,
		"a": 1,
		"b": 2,
	})
	if result.Statement != `UPDATE test SET a = ?,b = ?,c = ?` {
		t.Errorf(`Expected "UPDATE test SET a = ?,b = ?,c = ?", got %s`, result.Statement)
	}
}
