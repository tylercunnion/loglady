all: test build

build:
	mkdir -p bin
	go build -o bin/loglady -v github.com/tylercunnion/loglady/cmd/loglady 

test:
	go test ./...
