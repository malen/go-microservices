
.PHONY: proto
proto:
	@./protoc gateway/**/proto/*.proto --go_out=. --go-grpc_out=.

.PHONY: gateway
gateway:
	go run gateway/main.go