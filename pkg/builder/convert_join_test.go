package builder

import (
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	j := Join{
		For: "artist a",
		Do: Find{
			"a.impl": "t.impl",
		},
		As: "LEFT",
	}
	ctx := CommonDialect{
		TableAlias: "t",
	}
	result, vals := j.Convert(ctx)
	fmt.Println("Join:", vals)
	if result == `LEFT JOIN artist a ON a.impl = t.impl` {
		return
	}
	t.Errorf(`Expected "LEFT JOIN artist a ON a.impl = t.impl", got %s`, result)
}

func TestConvertJoin(t *testing.T) {
	joins := []interface{}{
		`{"$for": "artist a", "$do": {"a.impl": "t.impl"}, "$as": "LEFT"}`,
		Join{
			For: "artist a",
			Do: Find{
				"a.impl": "t.impl",
			},
		},
	}
	ctx := CommonDialect{
		TableAlias: "t",
	}
	result, vals := convertJoin(ctx, joins...)
	fmt.Println("Join:", vals)
	if result[1] != `JOIN artist a ON a.impl = t.impl` {
		t.Errorf(`Expected "JOIN artist a ON a.impl = t.impl", got %s`, result[1])
	}

	if result[0] != `LEFT JOIN artist a ON a.impl = t.impl` {
		t.Errorf(`Expected "LEFT JOIN artist a ON a.impl = t.impl", got %s`, result[0])
	}

}
func TestConvertMap(t *testing.T) {
	joins := []interface{}{
		Map{"$for": "artist a", "$do": Map{"a.impl": "t.impl"}, "$as": "LEFT"},
	}
	ctx := CommonDialect{
		TableAlias: "t",
	}
	result, vals := convertJoin(ctx, joins...)
	fmt.Println("Join:", vals)
	if result[0] != `LEFT JOIN artist a ON a.impl = t.impl` {
		t.Errorf(`Expected "LEFT JOIN artist a ON a.impl = t.impl", got %s`, result[0])
	}

}
