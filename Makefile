PROTO_DIR=proto
PRODUCT_OUTPUT_DIR=productSvc/internal
PRODUCT_PROTO_FILE=$(PROTO_DIR)/product.proto

ORDER_OUTPUT_DIR=orderSvc/internal
ORDER_PROTO_FILE=$(PROTO_DIR)/order.proto

PROTOC=protoc
PROTOC_GEN_GO=protoc-gen-go
PROTOC_GEN_GO_GRPC=protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY=protoc-gen-grpc-gateway

.PHONY: all product-proto order-proto service-gateway clean

all: product-proto order-proto

product-proto:
	@echo "Generating Go files for $(PRODUCT_PROTO_FILE)..."
	@if [ ! -d "$(OUTPUT_DIR)" ]; then mkdir -p $(PRODUCT_OUTPUT_DIR); fi
	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(PRODUCT_OUTPUT_DIR) \
		--go-grpc_out=$(PRODUCT_OUTPUT_DIR) \
		$(PRODUCT_PROTO_FILE)

order-proto:
	@echo "Generating Go files for $(ORDER_PROTO_FILE)..."
	@if [ ! -d "$(OUTPUT_DIR)" ]; then mkdir -p $(ORDER_OUTPUT_DIR); fi
	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(ORDER_OUTPUT_DIR) \
		--go-grpc_out=$(ORDER_OUTPUT_DIR) \
		$(ORDER_PROTO_FILE)

# Add new target for service-gateway
service-gateway:
	@echo "Generating Go files and Gateway files for $(PRODUCT_PROTO_FILE) and $(ORDER_PROTO_FILE)..."
	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--proto_path=$(PROTO_DIR)/google \
		--go_out=service-gateway/product \
		--go-grpc_out=service-gateway/product \
		--grpc-gateway_out=logtostderr=true:service-gateway/product \
		$(PRODUCT_PROTO_FILE)

	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--proto_path=$(PROTO_DIR)/google \
		--go_out=service-gateway/order \
		--go-grpc_out=service-gateway/order \
		--grpc-gateway_out=logtostderr=true:service-gateway/order \
		$(ORDER_PROTO_FILE)


clean:
	@echo "Cleaning up generated files..."
	rm -f $(PRODUCT_OUTPUT_DIR)/*.pb.go
	rm -f $(ORDER_OUTPUT_DIR)/*.pb.go
