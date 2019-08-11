dev:
	go generate ./...
	go build -o demo main.go
	./demo

install:
	GOPROXY=https://mirrors.aliyun.com/goproxy/ GO111MODULE=on go mod vendor -v