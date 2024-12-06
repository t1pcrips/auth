include local.env

LOCAL_BIN := $(CURDIR)/bin
LOCAL_MIGRATION_DSN="host=localhost port=54321 dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=disable"
LOCAL_MIGRATION_DIR = ./migrations

install golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... -- config .golangci.reference.yml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

get-deps-for-protoc:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generation-protoc:
	mkdir -p pkg/user_v1
	protoc --proto_path grpc/user_v1  --proto_path vendor.protogen \
            --go_out=pkg/user_v1 --go_opt=paths=source_relative \
            --plugin=protoc-gen-go=bin/protoc-gen-go \
            --go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
            --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
            --validate_out=lang=go:pkg/user_v1 --validate_opt=paths=source_relative \
            --plugin=protoc-gen-validate=bin/protoc-gen-validate \
            --grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
            --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
            grpc/user_v1/user.proto

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
        			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
        			mkdir -p  vendor.protogen/google/ &&\
        			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
        			rm -rf vendor.protogen/googleapis ;\
        fi

migrate-new:
	mkdir -p migrations
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) create ${NAME} sql

migrate-status:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) status -v

migrate-up:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) up -v

migrate-down:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) down -v

migrate-reset:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) reset -v