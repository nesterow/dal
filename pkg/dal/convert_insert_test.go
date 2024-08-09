package dal

import (
	"fmt"
	"testing"

	filters "l12.xyz/dal/filters"
)

func TestConvertInsert(t *testing.T) {
	ctx := filters.SQLiteContext{
		TableName:  "test",
		TableAlias: "t",
	}
	insert := []Map{
		{"a": "1", "b": 2},
		{"b": 2, "a": "1", "c": 3},
	}
	result, _ := ConvertInsert(ctx, insert)

	if result.Statement != `INSERT INTO test (a,b,c) VALUES (?,?,?)` {
		t.Errorf(`Expected "INSERT INTO test (a,b,c) VALUES (?,?,?)", got %s`, result.Statement)
	}

	for _, r := range result.Values {
		fmt.Println(r)
	}

}
