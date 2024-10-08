package builder

import (
	"testing"
)

func TestConvertFieldsBool(t *testing.T) {
	result, err := convertFields([]Map{
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
	result, err := convertFields([]Map{
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
	result, err := convertFields([]Map{
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
