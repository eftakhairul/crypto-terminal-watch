GOOS=
GOARCH=

export PATH=$(shell printenv PATH):${GOPATH}/build

default: build

rebuild: build_clean
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o ./build/proxy ./src/main/*.go

build: vendor_get
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o ./build/proxy ./src/main/*.go

build_clean:
	rm -rf ./build/proxy

vendor_get: 
	go get https://github.com/tools/godep

test: unit_test integration_test

