module pkg/proto

go 1.22.6

replace pkg/builder => ../builder

require (
	pkg/adapter v0.0.0
	pkg/builder v0.0.0-00010101000000-000000000000
	github.com/tinylib/msgp v1.2.0
)

require github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect

replace pkg/utils => ../utils

replace pkg/adapter => ../adapter

replace pkg/filters => ../filters

require (
	pkg/filters v0.0.0-00010101000000-000000000000 // indirect
	pkg/utils v0.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)
