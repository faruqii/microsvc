PROTO_DIR=../proto
PRODUCT_OUTPUT_DIR=externals/
PRODUCT_PROTO_FILE=$(PROTO_DIR)/product.proto

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