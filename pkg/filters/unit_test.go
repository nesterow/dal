package filters

import (
	"testing"

	adapter "l12.xyz/dal/adapter"
)

type SQLiteContext = adapter.SQLiteContext

func TestEq(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result, _ := Convert(ctx, `{"$eq": "NULL"}`)
	resultMap, _ := Convert(ctx, Filter{"$eq": "NULL"})
	if result != `t.test IS NULL` {
		t.Errorf("Expected t.test IS NULL, got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}

func TestGte(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result, _ := Convert(ctx, `{"$gte": 1}`)
	resultMap, _ := Convert(ctx, Filter{"$gte": 1})
	if result != `t.test >= 1` {
		t.Errorf("Expected t.test >= 1, got %s", result)
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
	result, _ := Convert(ctx, `{"$ne": "1"}`)
	resultMap, _ := Convert(ctx, Filter{"$ne": "1"})
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
	result, _ := Convert(ctx, `{"$between": ["1", "5"]}`)
	resultMap, _ := Convert(ctx, Filter{"$between": []string{"1", "5"}})
	if result != `test BETWEEN '1' AND '5'` {
		t.Errorf("Expected test BETWEEN '1' AND '5', got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}

func TestNotBetween(t *testing.T) {
	ctx := SQLiteContext{
		FieldName: "test",
	}
	result, _ := Convert(ctx, `{"$nbetween": ["1", "5"]}`)
	resultMap, _ := Convert(ctx, Filter{"$nbetween": []string{"1", "5"}})
	if result != `test NOT BETWEEN '1' AND '5'` {
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
	result, _ := Convert(ctx, `{"$glob": "*son"}`)
	resultMap, _ := Convert(ctx, Filter{"$glob": "*son"})
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
	result, _ := Convert(ctx, `{"$in": [1, 2, 3]}`)
	resultMap, _ := Convert(ctx, Filter{"$in": []int{1, 2, 3}})
	if result != `t.test IN (1, 2, 3)` {
		t.Errorf("Expected t.test IN (1, 2, 3), got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}

func TestNotIn(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result, _ := Convert(ctx, `{"$nin": [1, 2, 3]}`)
	resultMap, _ := Convert(ctx, Filter{"$nin": []int{1, 2, 3}})
	if result != `t.test NOT IN (1, 2, 3)` {
		t.Errorf("Expected t.test NOT IN (1, 2, 3), got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}

func TestLike(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result, _ := Convert(ctx, `{"$like": "199_"}`)
	resultMap, _ := Convert(ctx, Filter{"$like": "199_"})
	if result != `t.test LIKE '199_' ESCAPE '\'` {
		t.Errorf("Expected t.test LIKE '199_' ESCAPE '\\', got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}

func TestAnd(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result, _ := Convert(ctx, `{"$and": ["a > 0", "b < 10"]}`)
	resultMap, _ := Convert(ctx, Filter{"$and": []string{"a > 0", "b < 10"}})
	if result != `(a > 0 AND b < 10)` {
		t.Errorf("Expected (a > 0 AND b < 10), got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}

func TestOr(t *testing.T) {
	ctx := SQLiteContext{
		TableAlias: "t",
		FieldName:  "test",
	}
	result, _ := Convert(ctx, `{"$or": ["a = 0", "b < 10"]}`)
	resultMap, _ := Convert(ctx, Filter{"$or": []string{"a = 0", "b < 10"}})
	if result != `(a = 0 OR b < 10)` {
		t.Errorf("Expected (a > 0 OR b < 10), got %s", result)
	}
	if resultMap != result {
		t.Log(resultMap)
		t.Errorf("Expected resultMap to be equal to result")
	}
}
