module github.com/nesterow/dal/pkg/builder

go 1.22.6

replace github.com/nesterow/dal/pkg/utils => ../utils

replace github.com/nesterow/dal/pkg/filters => ../filters

replace github.com/nesterow/dal/pkg/adapter => ../adapter

require (
	github.com/nesterow/dal/pkg/adapter v0.0.0-20240820175837-f06ad4a34238
	github.com/nesterow/dal/pkg/filters v0.0.0-20240820175837-f06ad4a34238
	github.com/nesterow/dal/pkg/utils v0.0.0-20240820175837-f06ad4a34238
)

require github.com/pkg/errors v0.9.1 // indirect
