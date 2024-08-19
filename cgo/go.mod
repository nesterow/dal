module srv

go 1.22.6

replace l12.xyz/dal/filters v0.0.0 => ../pkg/filters

replace l12.xyz/dal/builder v0.0.0 => ../pkg/builder

replace l12.xyz/dal/handler v0.0.0 => ../pkg/handler

require l12.xyz/dal/adapter v0.0.0 // indirect

replace l12.xyz/dal/adapter v0.0.0 => ../pkg/adapter

replace l12.xyz/dal/utils v0.0.0 => ../pkg/utils

replace l12.xyz/dal/proto v0.0.0 => ../pkg/proto

require l12.xyz/dal/facade v0.0.0

require l12.xyz/dal/handler v0.0.0 // indirect

replace l12.xyz/dal/facade v0.0.0 => ../pkg/facade

require (
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tinylib/msgp v1.2.0 // indirect
	l12.xyz/dal/builder v0.0.0 // indirect
	l12.xyz/dal/filters v0.0.0 // indirect
	l12.xyz/dal/proto v0.0.0 // indirect
	l12.xyz/dal/utils v0.0.0 // indirect
)
