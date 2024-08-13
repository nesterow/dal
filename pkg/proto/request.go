package proto

import (
	"fmt"
	"reflect"

	"l12.xyz/dal/adapter"
	"l12.xyz/dal/builder"
)

//go:generate msgp

type BuildCmd struct {
	Method string        `msg:"method"`
	Args   []interface{} `msg:"args"`
}

type Request struct {
	Db       string     `msg:"db"`
	Commands []BuildCmd `msg:"commands"`
}

func (q *Request) Parse(dialect adapter.Dialect) (adapter.Query, error) {
	b := builder.New(dialect)
	for _, cmd := range q.Commands {
		method := reflect.ValueOf(b).MethodByName(cmd.Method)
		if !method.IsValid() {
			return adapter.Query{}, fmt.Errorf("method %s not found", cmd.Method)
		}
		args := make([]reflect.Value, len(cmd.Args))
		for i, arg := range cmd.Args {
			args[i] = reflect.ValueOf(arg)
		}
		method.Call(args)
	}
	expr, data := b.Sql()
	return adapter.Query{
		Db:         q.Db,
		Expression: expr,
		Data:       data,
	}, nil
}
