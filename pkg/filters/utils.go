package filters

import (
	"encoding/json"
	"fmt"
	"strings"
)

func FromJson[T IFilter](data interface{}) *T {
	var t T
	str, ok := data.(string)
	if ok {
		err := json.Unmarshal([]byte(str), &t)
		if err != nil {
			return &t
		}
	}
	m, ok := data.(Filter)
	if ok {
		s, err := json.Marshal(m)
		if err != nil {
			return &t
		}

		e := json.Unmarshal(s, &t)
		if e != nil {
			return &t
		}
	}
	return &t
}

func ValueOrPlaceholder(value interface{}) interface{} {
	if value == nil {
		return "NULL"
	}
	val, ok := value.(string)
	if !ok {
		return "?"
	}
	if strings.Contains(val, ".") {
		return value
	}
	return "?"
}

func FmtCompare(operator string, a interface{}, b interface{}) (string, Values) {
	if ValueOrPlaceholder(b) == "?" {
		return fmt.Sprintf("%s %s ?", a, operator), Values{b}
	}
	return fmt.Sprintf("%s %s %s", a, operator, ValueOrPlaceholder(b)), nil
}
