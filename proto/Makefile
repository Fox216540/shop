PROTO_DIR=proto
OUT_DIR=proto

# Catalog API
CATALOG_API=catalog-service/api.proto common/types.proto

# Catalog Internal (для Order)
CATALOG_INTERNAL=catalog-service/internal.proto common/types.proto

# Order
#ORDER_PROTO=proto/order/order.proto proto/common/types.proto

gen-common:
	@mkdir -p common/gen
	protoc -I=. \
		--go_out=common/ \
        --go_opt=module=github.com/Fox216540/shop/proto/common \
        --go-grpc_out=common/ \
        --go-grpc_opt=module=github.com/Fox216540/shop/proto/common \
        common/types.proto
	@echo "✅ Common proto generated in common/gen/"

gen-catalog-api:
	@make gen-common
	@mkdir -p catalog-service/gen
	protoc -I=. \
            --go_out=catalog-service/ \
            --go_opt=module=github.com/Fox216540/shop/proto/catalog-service \
            --go-grpc_out=catalog-service/ \
            --go-grpc_opt=module=github.com/Fox216540/shop/proto/catalog-service \
		catalog-service/api.proto
	@echo "✅ Catalog API generated"

gen-catalog-internal:
	@make gen-common
	@mkdir -p catalog-service/gen
	protoc -I=. \
			--go_out=catalog-service/ \
			--go_opt=module=github.com/Fox216540/shop/proto/catalog-service \
			--go-grpc_out=catalog-service/ \
			--go-grpc_opt=module=github.com/Fox216540/shop/proto/catalog-service \
		catalog-service/internal.proto
	@echo "✅ Catalog Internal generated"

gen-catalog:
	@make gen-catalog-api gen-catalog-internal
	@echo "✅ Catalog generated"


gen-auth:
	@mkdir -p auth-service/gen
	protoc -I=. \
			--go_out=auth-service/ \
			--go_opt=module=github.com/Fox216540/shop/proto/auth-service \
			--go-grpc_out=auth-service/ \
			--go-grpc_opt=module=github.com/Fox216540/shop/proto/auth-service \
		auth-service/internal.proto
	@echo "✅ Auth proto generated"

gen-order:
	@mkdir -p order-service/gen
	protoc -I=. \
			--go_out=order-service/ \
			--go_opt=module=github.com/Fox216540/shop/proto/order-service \
			--go-grpc_out=order-service/ \
			--go-grpc_opt=module=github.com/Fox216540/shop/proto/order-service \
		order-service/internal.proto
	@echo "✅ Order proto generated"

gen-user:
	@mkdir -p user-service/gen
	protoc -I=. \
			--go_out=user-service/ \
			--go_opt=module=github.com/Fox216540/shop/proto/user-service \
			--go-grpc_out=user-service/ \
			--go-grpc_opt=module=github.com/Fox216540/shop/proto/user-service \
		user-service/api.proto
	@echo "✅ User proto generated"

gen-all: gen-catalog gen-auth gen-order gen-user
