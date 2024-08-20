module github.com/nesterow/dal/pkg/builder

go 1.22.6

require (
	pkg/adapter v0.0.0
	pkg/filters v0.0.0-00010101000000-000000000000
	pkg/utils v0.0.0
)

replace pkg/utils => ../utils

require github.com/pkg/errors v0.9.1 // indirect

replace pkg/filters => ../filters

replace pkg/adapter => ../adapter
