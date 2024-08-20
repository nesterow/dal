module pkg/facade

go 1.22.6

replace github.com/nesterow/dal/pkg/filters => ../filters

replace github.com/nesterow/dal/pkg/builder => ../builder

require (
	github.com/nesterow/dal/pkg/adapter v0.0.0
	github.com/nesterow/dal/pkg/handler v0.0.0-00010101000000-000000000000
	github.com/nesterow/dal/pkg/proto v0.0.0
)

replace github.com/nesterow/dal/pkg/adapter => ../adapter

replace github.com/nesterow/dal/pkg/utils => ../utils

require (
	github.com/nesterow/dal/pkg/builder v0.0.0 // indirect
	github.com/nesterow/dal/pkg/filters v0.0.0 // indirect
	github.com/nesterow/dal/pkg/utils v0.0.0 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
)

replace github.com/nesterow/dal/pkg/proto => ../proto

replace github.com/nesterow/dal/pkg/handler => ../handler
