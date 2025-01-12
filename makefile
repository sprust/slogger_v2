run:
	go run ./cmd/slogger/main.go ${c}

build:
	CGO_ENABLED=0 GOOS=linux go build -v -a -o ./bin/ ./cmd/slogger/main.go \
		&& chmod +x ./bin/main