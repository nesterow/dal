package filters

import (
	"fmt"
	"testing"

	"github.com/nesterow/dal/pkg/adapter"
)

type SQLiteContext = adapter.CommonDialect

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
	result, vals := Convert(ctx, `{"$gte": 1}`)
	resultMap, _ := Convert(ctx, Filter{"$gte": 1})
	if vals[0].(float64) != 1 {
		t.Errorf("Expected 1, got %v", vals[0])
	}
	if result != `t.test >= ?` {
		t.Errorf("Expected t.test >= ?, got %s", result)
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
	if result != `test != ?` {
		t.Errorf("Expected test != ?, got %s", result)
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
	result, vals := Convert(ctx, `{"$between": ["1", "5"]}`)
	fmt.Println(vals)
	resultMap, _ := Convert(ctx, Filter{"$between": []string{"1", "5"}})
	if result != `test BETWEEN ? AND ?` {
		t.Errorf("Expected test BETWEEN ? AND ?, got %s", result)
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
	if result != `test NOT BETWEEN ? AND ?` {
		t.Errorf("Expected test NOT BETWEEN ? AND ?, got %s", result)
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
	result, vals := Convert(ctx, `{"$glob": "*son"}`)
	resultMap, _ := Convert(ctx, Filter{"$glob": "*son"})
	if vals[0].(string) != "*son" {
		t.Errorf("Expected *son, got %v", vals[0])
	}
	if result != `t.test GLOB ?` {
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
	if result != `t.test IN (?, ?, ?)` {
		t.Errorf("Expected t.test IN (?, ?, ?), got %s", result)
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
	result, vals := Convert(ctx, `{"$nin": [1, 2, 3]}`)
	resultMap, _ := Convert(ctx, Filter{"$nin": []int{1, 2, 3}})
	if vals[1].(float64) != 2 {
		t.Errorf("Expected 1, got %v", vals[1])
	}
	if result != `t.test NOT IN (?, ?, ?)` {
		t.Errorf("Expected t.test NOT IN (?, ?, ?), got %s", result)
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
	result, vals := Convert(ctx, `{"$like": "199_"}`)
	resultMap, _ := Convert(ctx, Filter{"$like": "199_"})
	if vals[0].(string) != "199_" {
		t.Errorf("Expected 199_, got %v", vals[0])
	}
	if result != `t.test LIKE ? ESCAPE '\'` {
		t.Errorf("Expected t.test LIKE ? ESCAPE '\\', got %s", result)
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
