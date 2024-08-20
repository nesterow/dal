module pkg/tests

go 1.22.6

require pkg/builder v0.0.0

replace pkg/builder v0.0.0 => ../builder

replace pkg/utils v0.0.0 => ../utils

require pkg/adapter v0.0.0

require pkg/proto v0.0.0

require (
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
)

replace pkg/adapter v0.0.0 => ../adapter

replace pkg/proto v0.0.0 => ../proto

replace pkg/filters v0.0.0 => ../filters

require (
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/pkg/errors v0.9.1 // indirect
	pkg/filters v0.0.0 // indirect
	pkg/utils v0.0.0 // indirect
)
