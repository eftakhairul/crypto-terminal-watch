GOOS=
GOARCH=

export PATH=$(shell printenv PATH):${GOPATH}/build

default: build

rebuild: build_clean
	GOOS=$(GOOS) GOARCH=$(GOARCH) ${GOPATH}/bin/godep go build -v -o ./build/ctw .

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) ${GOPATH}/bin/godep go build -v -o ./build/ctw .

install: 
	go get github.com/tools/godep
	${GOPATH}/bin/godep go build ./… 

test: unit_test integration_test

