GOPATH := $(shell pwd)

all:
	rm -rf ./httpd
	go build -o ./httpd ./src/httpd.go
