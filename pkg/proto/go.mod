module github.com/nesterow/dal/pkg/proto

go 1.22.6

replace github.com/nesterow/dal/pkg/builder => ../builder
replace github.com/nesterow/dal/pkg/adapter => ../adapter
replace github.com/nesterow/dal/pkg/filters => ../filters
replace github.com/nesterow/dal/pkg/utils => ../utils

require (
	github.com/nesterow/dal/pkg/adapter v0.0.0-20240820175837-f06ad4a34238
	github.com/nesterow/dal/pkg/builder v0.0.0-20240820175837-f06ad4a34238
	github.com/nesterow/dal/pkg/filters v0.0.0-20240820175837-f06ad4a34238
	github.com/nesterow/dal/pkg/utils v0.0.0-20240820175837-f06ad4a34238
	github.com/tinylib/msgp v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
)

require github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
