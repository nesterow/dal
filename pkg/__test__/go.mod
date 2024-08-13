module l12.xyz/dal/tests

go 1.22.6

require l12.xyz/dal/builder v0.0.0

replace l12.xyz/dal/builder v0.0.0 => ../builder

replace l12.xyz/dal/utils v0.0.0 => ../utils

require l12.xyz/dal/adapter v0.0.0

require l12.xyz/dal/proto v0.0.0

require (
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
)

replace l12.xyz/dal/adapter v0.0.0 => ../adapter

replace l12.xyz/dal/proto v0.0.0 => ../proto

replace l12.xyz/dal/filters v0.0.0 => ../filters

require (
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/pkg/errors v0.9.1 // indirect
	l12.xyz/dal/filters v0.0.0 // indirect
	l12.xyz/dal/utils v0.0.0 // indirect
)
