
.PHONY: proto
proto:
	@./protoc gateway/**/proto/*.proto --go_out=. --go-grpc_out=.

.PHONY: gateway
gateway:
	go run gateway/main.go

.PHONY: auth
auth:
	go run auth/main.go

.PHONY: product
product:
	go run product/main.go

.PHONY: order
order:
	go run order/main.go