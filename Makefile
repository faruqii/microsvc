PROTO_DIR=proto
PRODUCT_OUTPUT_DIR=productSvc/internal/protobuff
PRODUCT_PROTO_FILE=$(PROTO_DIR)/product.proto

ORDER_OUTPUT_DIR=orderSvc/internal/protobuff
ORDER_PROTO_FILE=$(PROTO_DIR)/order.proto

PROTOC=protoc
PROTOC_GEN_GO=protoc-gen-go
PROTOC_GEN_GO_GRPC=protoc-gen-go-grpc

.PHONY: all product-proto clean

all: product-proto

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

clean:
	@echo "Cleaning up generated files..."
	rm -f $(OUTPUT_DIR)/*.pb.go
