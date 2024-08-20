module github.com/nesterow/dal/pkg/filters/v0.0.0

go 1.22.6

require github.com/pkg/errors v0.9.1 // indirect

require (
	github.com/nesterow/dal/pkg/adapter v0.0.0-00010101000000-000000000000
	github.com/nesterow/dal/pkg/utils v0.0.0-00010101000000-000000000000
)

replace github.com/nesterow/dal/pkg/utils => ../utils

replace github.com/nesterow/dal/pkg/adapter => ../adapter
