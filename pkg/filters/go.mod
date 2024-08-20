module github.com/nesterow/dal/pkg/filters

go 1.22.6

require github.com/pkg/errors v0.9.1 // indirect

require (
	pkg/adapter v0.0.0-00010101000000-000000000000
	pkg/utils v0.0.0-00010101000000-000000000000
)

replace pkg/utils => ../utils

replace pkg/adapter => ../adapter
