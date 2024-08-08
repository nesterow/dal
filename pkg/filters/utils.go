package filters

import (
	"encoding/json"
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
