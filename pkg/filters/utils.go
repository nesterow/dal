package filters

import (
	"encoding/json"
)

func FromJson[T Filter](data interface{}) *T {
	var t T
	str, ok := data.(string)
	if ok {
		err := json.Unmarshal([]byte(str), &t)
		if err != nil {
			return nil
		}
	}
	m, ok := data.(map[string]interface{})
	if ok {
		s, err := json.Marshal(m)
		if err != nil {
			return nil
		}

		e := json.Unmarshal(s, &t)
		if e != nil {
			return nil
		}
	}
	return &t
}
