package builder

import (
	"encoding/json"
	"fmt"
)

type Join struct {
	For string `json:"$for"`
	Do  Find   `json:"$do"`
	As  string `json:"$as"`
}

func (j Join) Convert(ctx Dialect) (string, Values) {
	if j.For == "" {
		return "", nil
	}
	filter, values := covertFind(ctx, j.Do)
	var as string = ""
	if j.As != "" {
		as = fmt.Sprintf("%s ", j.As)
	}
	return as + fmt.Sprintf("JOIN %s ON %s", j.For, filter), values
}

func convertJoin(ctx Dialect, joins ...interface{}) ([]string, Values) {
	var result []string
	var values Values
	for _, join := range joins {
		jstr, ok := join.(string)
		if ok {
			jjson := Join{}
			err := json.Unmarshal([]byte(jstr), &jjson)
			if err == nil {
				r, vals := jjson.Convert(ctx)
				result = append(result, r)
				values = append(values, vals...)
			}
			continue
		}
		jmap, ok := join.(map[string]interface{})
		if ok {
			jjson := Join{}
			jstr, err := json.Marshal(jmap)
			if err != nil {
				continue
			}
			err = json.Unmarshal(jstr, &jjson)
			if err == nil {
				r, vals := jjson.Convert(ctx)
				result = append(result, r)
				values = append(values, vals...)
			}
			continue
		}
		j, ok := join.(Join)
		if !ok {
			continue
		}
		r, vals := j.Convert(ctx)
		result = append(result, r)
		values = append(values, vals...)
	}
	return result, values
}
