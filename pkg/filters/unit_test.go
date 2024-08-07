package filters

import (
	"testing"
)

func TestEq(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result := Convert(ctx, `{"$eq": "NULL"}`)
	resultMap := Convert(ctx, map[string]any{"$eq": "NULL"})
	if result != `t.test IS NULL` {
		t.Errorf("Expected t.test IS NULL, got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}

func TestNe(t *testing.T) {
	ctx := SQLiteContext{
		FieldName: "test",
	}
	result := Convert(ctx, `{"$ne": "1"}`)
	resultMap := Convert(ctx, map[string]any{"$ne": "1"})
	if result != `test != '1'` {
		t.Errorf("Expected test != '1', got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}

func TestBetween(t *testing.T) {
	ctx := SQLiteContext{
		FieldName: "test",
	}
	result := Convert(ctx, `{"$between": ["1", "5"]}`)
	resultMap := Convert(ctx, map[string]any{"$between": []string{"1", "5"}})
	if result != `test BETWEEN '1' AND '5'` {
		t.Errorf("Expected test BETWEEN '1' AND '5', got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}

func TestGlob(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result := Convert(ctx, `{"$glob": "*son"}`)
	resultMap := Convert(ctx, map[string]any{"$glob": "*son"})
	if result != `t.test GLOB '*son'` {
		t.Errorf("Expected t.test GLOB '*son', got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}

func TestIn(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result := Convert(ctx, `{"$in": [1, 2, 3]}`)
	resultMap := Convert(ctx, map[string]any{"$in": []int{1, 2, 3}})
	if result != `t.test IN (1, 2, 3)` {
		t.Errorf("Expected t.test IN (1, 2, 3), got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}
