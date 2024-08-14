package proto

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	"l12.xyz/dal/adapter"
	"l12.xyz/dal/builder"
)

//go:generate msgp

type BuildCmd struct {
	Method string        `msg:"method"`
	Args   []interface{} `msg:"args"`
}

type Request struct {
	Id       uint32     `msg:"id"`
	Db       string     `msg:"db"`
	Commands []BuildCmd `msg:"commands"`
}

var allowedMethods = strings.Split(builder.BUILDER_CLIENT_METHODS, "|")

func (q *Request) Parse(dialect adapter.Dialect) (adapter.Query, error) {
	if q.Db == "" {
		return adapter.Query{}, fmt.Errorf("Request format: db url is required")
	}
	if len(q.Commands) == 0 {
		return adapter.Query{}, fmt.Errorf("Request format: commands are required")
	}
	b := builder.New(dialect)
	for _, cmd := range q.Commands {
		if !slices.Contains(allowedMethods, cmd.Method) {
			return adapter.Query{}, fmt.Errorf(
				"method %s is not allowed, available methods are %v",
				cmd.Method,
				allowedMethods,
			)
		}
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
