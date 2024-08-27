build-lib:
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o ../lib/libdal-arm64.dylib -buildmode=c-shared ./dal.go
	for arch in amd64 arm64; do\
		CC=x86_64-unknown-linux-gnu-gcc CGO_ENABLED=1 GOOS=linux GOARCH=$${arch} go build -o ../lib/libdal-$${arch}.so -buildmode=c-shared ./dal.go;\
	done