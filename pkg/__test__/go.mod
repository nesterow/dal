module github.com/nesterow/dal/pkg/tests

go 1.22.6

require github.com/nesterow/dal/pkg/builder v0.0.0

replace github.com/nesterow/dal/pkg/builder v0.0.0 => ../builder

replace github.com/nesterow/dal/pkg/utils v0.0.0 => ../utils

require github.com/nesterow/dal/pkg/adapter v0.0.0

require github.com/nesterow/dal/pkg/proto v0.0.0

require (
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
)

replace github.com/nesterow/dal/pkg/adapter v0.0.0 => ../adapter

replace github.com/nesterow/dal/pkg/proto v0.0.0 => ../proto

replace github.com/nesterow/dal/pkg/filters v0.0.0 => ../filters

require (
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/pkg/errors v0.9.1 // indirect
	github.com/nesterow/dal/pkg/filters v0.0.0 // indirect
	github.com/nesterow/dal/pkg/utils v0.0.0 // indirect
)
