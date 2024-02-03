all: build

build:
	go build -o kabbase cmd/api/*.go
	go build -o kbase-cli cmd/cli/*.go

linuxbuild:
	GOOS=linux GOARCH=amd64 go build -o kabbase.linux cmd/api/*.go
	GOOS=linux GOARCH=amd64 go build -o kbase-cli.linux cmd/cli/*.go

clean:
	rm -rf kabbase
	rm -rf kbase-cli
	rm -rf kabbase.*
	rm -rf kbase-cli.linux

test:
