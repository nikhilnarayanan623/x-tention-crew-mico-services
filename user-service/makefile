build:
	go build -o ./build/bin/api ./cmd/api

run: 
	go run ./cmd/api

proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/proto/*.proto

wire:
	cd pkg/di && wire

docker-build:
	docker build -t nikhil382/x-tention-crew-user-service .

swag:
	swag init -g pkg/api/server.go -o ./cmd/api/docs