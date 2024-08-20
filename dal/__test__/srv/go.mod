module srv

go 1.22.6

replace pkg/filters v0.0.0 => ../../../pkg/filters

replace pkg/builder v0.0.0 => ../../../pkg/builder

require pkg/adapter v0.0.0

replace pkg/adapter v0.0.0 => ../../../pkg/adapter

replace pkg/utils v0.0.0 => ../../../pkg/utils

replace pkg/proto v0.0.0 => ../../../pkg/proto

require pkg/handler v0.0.0

replace pkg/handler v0.0.0 => ../../../pkg/handler

require (
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
	pkg/builder v0.0.0 // indirect
	pkg/filters v0.0.0 // indirect
	pkg/proto v0.0.0 // indirect
	pkg/utils v0.0.0 // indirect
)
