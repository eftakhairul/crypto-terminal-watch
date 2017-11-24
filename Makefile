GOOS=
GOARCH=

export PATH=$(shell printenv PATH):${GOPATH}/build

default: build

rebuild: build_clean
	GOOS=$(GOOS) GOARCH=$(GOARCH) ${GOPATH}/bin/godep go build -v -o ./build/ctw .

build: build_clean
	GOOS=$(GOOS) GOARCH=$(GOARCH) ${GOPATH}/bin/godep go build -v -o ./build/ctw .

build_clean:
	rm -rf ./build/ctw

install: 
	go get github.com/tools/godep
	${GOPATH}/bin/godep go build ./… 

test:
	${GOPATH}/bin/godep go test -v crypto/coinmarketcap_test.go