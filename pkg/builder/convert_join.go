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

func (j Join) Convert(ctx Dialect) string {
	if j.For == "" {
		return ""
	}
	filter := covertFind(ctx, j.Do)
	var as string = ""
	if j.As != "" {
		as = fmt.Sprintf("%s ", j.As)
	}
	return as + fmt.Sprintf("JOIN %s ON %s", j.For, filter)
}

func convertJoin(ctx Dialect, joins ...interface{}) []string {
	var result []string
	for _, join := range joins {
		jstr, ok := join.(string)
		if ok {
			jjson := Join{}
			err := json.Unmarshal([]byte(jstr), &jjson)
			if err == nil {
				result = append(result, jjson.Convert(ctx))
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
				result = append(result, jjson.Convert(ctx))
			}
			continue
		}
		j, ok := join.(Join)
		if !ok {
			continue
		}
		result = append(result, j.Convert(ctx))
	}
	return result
}
