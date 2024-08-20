module github.com/nesterow/dal/pkg/builder

go 1.22.6

require (
	github.com/nesterow/dal/pkg/adapter v0.0.0
	github.com/nesterow/dal/pkg/filters v0.0.0-00010101000000-000000000000
	github.com/nesterow/dal/pkg/utils v0.0.0
)

replace github.com/nesterow/dal/pkg/utils => ../utils

require github.com/pkg/errors v0.9.1 // indirect

replace github.com/nesterow/dal/pkg/filters => ../filters

replace github.com/nesterow/dal/pkg/adapter => ../adapter
