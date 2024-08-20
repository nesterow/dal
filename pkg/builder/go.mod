module github.com/nesterow/dal/pkg/builder

go 1.22.6

require github.com/nesterow/dal/pkg/utils v0.0.0

replace github.com/nesterow/dal/pkg/utils v0.0.0 => ../utils

require github.com/nesterow/dal/pkg/filters v0.0.0

require github.com/pkg/errors v0.9.1 // indirect

replace github.com/nesterow/dal/pkg/filters v0.0.0 => ../filters

require github.com/nesterow/dal/pkg/adapter v0.0.0

replace github.com/nesterow/dal/pkg/adapter v0.0.0 => ../adapter
