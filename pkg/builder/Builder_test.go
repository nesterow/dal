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
	if db.Sql() != expect {
		t.Errorf(`Expected: "%s", Got: %s`, expect, db.Sql())
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
	if db.Sql() != expect {
		t.Errorf(`Expected: "%s", Got: %s`, expect, db.Sql())
	}
}
