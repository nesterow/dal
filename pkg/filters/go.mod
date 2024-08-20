module github.com/nesterow/dal/pkg/filters

go 1.22.6

require github.com/pkg/errors v0.9.1 // indirect

require github.com/nesterow/dal/pkg/utils v0.0.0

replace github.com/nesterow/dal/pkg/utils v0.0.0 => ../utils

require github.com/nesterow/dal/pkg/adapter v0.0.0

replace github.com/nesterow/dal/pkg/adapter v0.0.0 => ../adapter
