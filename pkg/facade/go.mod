module github.com/nesterow/dal/pkg/facade

go 1.22.6

replace pkg/filters => ../filters

replace pkg/builder => ../builder

require (
	pkg/adapter v0.0.0
	pkg/handler v0.0.0-00010101000000-000000000000
	pkg/proto v0.0.0
)

replace pkg/adapter => ../adapter

replace pkg/utils => ../utils

require (
	pkg/builder v0.0.0 // indirect
	pkg/filters v0.0.0 // indirect
	pkg/utils v0.0.0 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
)

replace pkg/proto => ../proto

replace pkg/handler => ../handler
