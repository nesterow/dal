module github.com/nesterow/dal/pkg/adapter

go 1.22.6

replace github.com/nesterow/dal/pkg/utils v0.0.0 => ../utils

replace github.com/nesterow/dal/pkg/filters v0.0.0 => ../filters

require github.com/nesterow/dal/pkg/utils v0.0.0

require github.com/pkg/errors v0.9.1 // indirect
