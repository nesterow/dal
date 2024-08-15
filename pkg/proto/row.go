package proto

import (
	"bytes"
)

//go:generate msgp

/**
* In most cases we need streaming responses for the table data
**/
type Row struct {
	Data []interface{} `msg:"r"`
}

func MarshalRow[T interface{}](columns []T) ([]byte, error) {
	s := make([]interface{}, len(columns))
	for i, v := range columns {
		s[i] = v
	}
	row := Row{Data: s}
	return row.MarshalMsg(nil)
}

func UnmarshalRows(input []byte) []Row {
	tag := []byte{0x81, 0xa1, 0x72}
	result := [][]byte{}
	buf := []byte{}
	count := 0
	for count < len(input) {
		if input[count] != 0x81 {
			buf = append(buf, input[count])
			count += 1
			continue
		}
		seq := input[count : len(tag)+count]
		if bytes.Equal(seq, tag) {
			result = append(result, append(tag, buf...))
			buf = []byte{}
		} else {
			buf = append(buf, seq...)
		}
		count += len(tag)
	}
	result = append(result, append(tag, buf...))
	rows := make([]Row, len(result))
	for i, r := range result {
		rows[i].UnmarshalMsg(r)
	}
	return rows[1:]
}
