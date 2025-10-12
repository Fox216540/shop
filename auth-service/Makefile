PROTO_DIR=api/proto
OUT_DIR=$(PROTO_DIR)
PROTO_FILES=$(wildcard $(PROTO_DIR)/*.proto)

proto:
	@mkdir -p $(OUT_DIR)
	protoc -I=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		$(PROTO_FILES)
	@echo "✅ Proto files generated in $(OUT_DIR)"
