include local.env

LOCAL_BIN := $(CURDIR)/bin
LOCAL_MIGRATION_DSN="host=localhost port=54321 dbname=${POSTGRES_DB} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=disable"
LOCAL_MIGRATION_DIR = ./migrations

install golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... -- config .golangci.reference.yml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.35.1
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@v2.50.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.1.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

get-deps-for-protoc:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generation-protoc-user:
	mkdir -p pkg/swagger
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
            --openapiv2_out=pkg/swagger \
            --openapiv2_opt=allow_merge=true \
            --openapiv2_opt=merge_file_name=grpc \
            --plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
            grpc/user_v1/user.proto

generation-protoc-auth:
	mkdir -p pkg/auth_v1
	protoc --proto_path grpc/auth_v1  --proto_path vendor.protogen \
            --go_out=pkg/auth_v1 --go_opt=paths=source_relative \
            --plugin=protoc-gen-go=bin/protoc-gen-go \
            --go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
            --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
            --validate_out=lang=go:pkg/auth_v1 --validate_opt=paths=source_relative \
            --plugin=protoc-gen-validate=bin/protoc-gen-validate \
            --grpc-gateway_out=pkg/auth_v1 --grpc-gateway_opt=paths=source_relative \
            --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
			grpc/auth_v1/auth.proto

generation-protoc-access:
	mkdir -p pkg/access_v1
	protoc --proto_path grpc/access_v1  --proto_path vendor.protogen \
            --go_out=pkg/access_v1 --go_opt=paths=source_relative \
            --plugin=protoc-gen-go=bin/protoc-gen-go \
            --go-grpc_out=pkg/access_v1 --go-grpc_opt=paths=source_relative \
            --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
            --validate_out=lang=go:pkg/access_v1 --validate_opt=paths=source_relative \
            --plugin=protoc-gen-validate=bin/protoc-gen-validate \
            --grpc-gateway_out=pkg/access_v1 --grpc-gateway_opt=paths=source_relative \
            --plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
			grpc/access_v1/access.proto

generate-statik:
	$(LOCAL_BIN)/statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'

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
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
        			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
        			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
        			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
        			rm -rf vendor.protogen/openapiv2 ;\
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

opensl-cert:
	openssl genrsa -out tls/ca.key 4096
	openssl req -new -x509 -key tls/ca.key -sha256 -subj "/C=RU/ST=LO/O=t1p, Inc." -days 365 -out tls/ca.cert
	openssl genrsa -out tls/service.key 4096
	openssl req -new -key tls/service.key -out tls/service.csr -config certificate.conf
	openssl x509 -req -in tls/service.csr -CA tls/ca.cert -CAkey tls/ca.key -CAcreateserial \
        		-out tls/service.pem -days 365 -sha256 -extfile certificate.conf -extensions req_ext

openssl-secretKey:
	openssl genrsa -out secrets/access_secret.key 2048
	openssl genrsa -out secrets/refresh_secret.key 2048

grpc-load-test:
	ghz \
		--proto grpc/user_v1/user.proto \
		--call user_v1.User.Get \
		--data '{"id": 1}' \
		--rps 100 \
		--total 3000 \
		--insecure \
		localhost:50051
