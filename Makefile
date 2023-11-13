GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --go-grpc_out=. --go_out=:. api/proto/pathfinder-service.proto

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: vendor
vendor: tidy
	@go mod vendor

.PHONY: test
test:
	@go test -v ./... -coverprofile coverage.out

.PHONY: bench
bench:
	@go test -bench . ./... -benchmem

.PHONY: dev
dev:
	@go build -o bin/pathfinder cmd/pathfinder/main.go
	@bin/pathfinder -dev
