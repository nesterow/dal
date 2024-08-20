module github.com/nesterow/dal

go 1.22.6

replace pkg/adapter => ./pkg/adapter

replace pkg/builder => ./pkg/builder

replace pkg/filters => ./pkg/filters

replace pkg/utils => ./pkg/utils

replace pkg/handler => ./pkg/handler

replace pkg/proto => ./pkg/proto

require (
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/nesterow/dal/pkg/facade v0.0.0
)

require (
	pkg/adapter v0.0.0 // indirect
	pkg/builder v0.0.0 // indirect
	pkg/filters v0.0.0 // indirect
	pkg/handler v0.0.0-00010101000000-000000000000 // indirect
	pkg/proto v0.0.0 // indirect
	pkg/utils v0.0.0 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
)
