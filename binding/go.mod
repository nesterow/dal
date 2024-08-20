module binding

go 1.22.6

replace pkg/filters v0.0.0 => ../pkg/filters

replace pkg/builder v0.0.0 => ../pkg/builder

replace pkg/handler v0.0.0 => ../pkg/handler

require pkg/adapter v0.0.0 // indirect

replace pkg/adapter v0.0.0 => ../pkg/adapter

replace pkg/utils v0.0.0 => ../pkg/utils

replace pkg/proto v0.0.0 => ../pkg/proto

require (
	github.com/mattn/go-sqlite3 v1.14.22
	pkg/facade v0.0.0
)

require pkg/handler v0.0.0 // indirect

replace pkg/facade v0.0.0 => ../pkg/facade

require (
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
	pkg/builder v0.0.0 // indirect
	pkg/filters v0.0.0 // indirect
	pkg/proto v0.0.0 // indirect
	pkg/utils v0.0.0 // indirect
)
