package builder

import (
	"testing"
)

func TestBuilderFind(t *testing.T) {
	db := New(SQLiteContext{})
	db.In("table t").Find(Query{
		"field": "value",
		"a":     1,
	})
	expect := "SELECT * FROM table t WHERE t.a = 1 AND t.field = 'value'"
	result, _ := db.Sql()
	if result != expect {
		t.Errorf(`Expected: "%s", Got: %s`, expect, result)
	}
}

func TestBuilderFields(t *testing.T) {
	db := New(SQLiteContext{})
	db.In("table t")
	db.Find(Query{
		"field": "value",
		"a":     1,
	})
	db.Fields(Map{
		"t.field": "field",
		"t.a":     1,
	})
	expect := "SELECT t.a, t.field AS field FROM table t WHERE t.a = 1 AND t.field = 'value'"
	result, _ := db.Sql()
	if result != expect {
		t.Errorf(`Expected: "%s", Got: %s`, expect, result)
	}
}

func TestBuilderGroup(t *testing.T) {
	db := New(SQLiteContext{})
	db.In("table t")
	db.Find(Query{
		"field": Is{
			"$gt": 1,
		},
	})
	db.Fields(Map{
		"SUM(t.field)": "field",
	})
	db.Group("field")
	expect := "SELECT SUM(t.field) AS field FROM table t GROUP BY t.field HAVING t.field > 1"
	result, _ := db.Sql()
	if result != expect {
		t.Errorf(`Expected: "%s", Got: %s`, expect, result)
	}
}

func TestBuilderJoin(t *testing.T) {
	db := New(SQLiteContext{})
	db.In("table t")
	db.Find(Query{
		"field": "value",
		"a":     1,
	})
	db.Join(Join{
		For: "table2 t2",
		Do: Query{
			"t2.field": "t.field",
		},
	})
	expect := "SELECT * FROM table t JOIN table2 t2 ON t2.field = t.field WHERE t.a = 1 AND t.field = 'value'"
	result, _ := db.Sql()
	if result != expect {
		t.Errorf(`Expected: "%s", Got: %s`, expect, result)
	}
}
