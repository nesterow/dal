module github.com/nesterow/dal/pkg/filters

go 1.22.6

require github.com/pkg/errors v0.9.1 // indirect

require (
	github.com/nesterow/dal/pkg/adapter v0.0.0-20240820175837-f06ad4a34238
	github.com/nesterow/dal/pkg/utils v0.0.0-20240820175837-f06ad4a34238
)

replace github.com/nesterow/dal/pkg/utils => ../utils

replace github.com/nesterow/dal/pkg/adapter => ../adapter
