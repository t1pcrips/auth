LOCAL_BIN := $(CURDIR)/bin

install golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0


lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... -- config .golangci.reference.yml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generation-protoc:
	mkdir -p $(CURDIR)/pkg/auth_v1
	protoc --proto_path grpc/auth/v1 \
	--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	grpc/auth/v1/auth.proto

generate:
	make generation-protoc;

build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/service-linux cmd/server/main.go

copy-to-server:
	scp ./bin/service-linux root@5.159.100.165:

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/server/test-server:v0.0.1 .
	docker login -u token -p CRgAAAAAMQ3VM9YXtVQHnwYf8Lx9aDq-6VyC9Ry4 cr.selcloud.ru/server
	docker push cr.selcloud.ru/server/test-server:v0.0.1