FROM golang:1.23.3-alpine AS builder

COPY . /github.com/t1pcrips/auth/source/
WORKDIR /github.com/t1pcrips/auth/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/t1pcrips/auth/source/bin/crud_server .

CMD ["./crud_server"]