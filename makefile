dev-init:
	cp -i .env.example .env

run:
	go run ./cmd/slogger/main.go ${c}

build:
	CGO_ENABLED=0 GOOS=linux go build -v -a -o ./bin/slogger ./cmd/slogger/main.go \
		&& chmod +x ./bin/slogger

grpc-generate:
	protoc --go_out=./internal/api/grpc/gen/ \
		--go-grpc_out=./internal/api/grpc/gen/ \
			./internal/api/grpc/proto/*.proto