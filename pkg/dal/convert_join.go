package dal

import "fmt"

type Join struct {
	For string `json:"$for"`
	Do  Find   `json:"$do"`
	As  string `json:"$as"`
}

func (j Join) Convert(ctx Context) string {
	if j.For == "" {
		return ""
	}
	filter := CovertFind(j.Do, ctx)
	return fmt.Sprintf("%s JOIN %s ON %s", j.As, j.For, filter)
}
