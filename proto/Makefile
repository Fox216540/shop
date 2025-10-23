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

gen-catalog-interservice:
	@make gen-common
	@mkdir -p catalog-service/gen
	protoc -I=. \
			--go_out=catalog-service/ \
			--go_opt=module=github.com/Fox216540/shop/proto/catalog-service \
			--go-grpc_out=catalog-service/ \
			--go-grpc_opt=module=github.com/Fox216540/shop/proto/catalog-service \
		catalog-service/interservice.proto
	@echo "✅ Catalog Interservice generated"

gen-catalog:
	@make gen-catalog-api gen-catalog-interservice
	@echo "✅ Catalog generated"


gen-auth-api:
	@make gen-common
	@mkdir -p auth-service/gen
	protoc -I=. \
			--go_out=auth-service/ \
			--go_opt=module=github.com/Fox216540/shop/proto/auth-service \
			--go-grpc_out=auth-service/ \
			--go-grpc_opt=module=github.com/Fox216540/shop/proto/auth-service \
		auth-service/api.proto
	@echo "✅ Auth Api proto generated"

gen-auth-interservice:
	@make gen-common
	@mkdir -p auth-service/gen
	protoc -I=. \
			--go_out=auth-service/ \
			--go_opt=module=github.com/Fox216540/shop/proto/auth-service \
			--go-grpc_out=auth-service/ \
			--go-grpc_opt=module=github.com/Fox216540/shop/proto/auth-service \
		auth-service/interservice.proto
	@echo "✅ Auth Interservice proto generated"

gen-auth:
	@make gen-auth-api gen-auth-interservice
	@echo "✅ Auth generated"

gen-order:
	@make gen-common
	@mkdir -p order-service/gen
	protoc -I=. \
			--go_out=order-service/ \
			--go_opt=module=github.com/Fox216540/shop/proto/order-service \
			--go-grpc_out=order-service/ \
			--go-grpc_opt=module=github.com/Fox216540/shop/proto/order-service \
		order-service/interservice.proto
	@echo "✅ Order proto generated"

gen-user-api:
	@make gen-common
	@mkdir -p user-service/gen
	protoc -I=. \
			--go_out=user-service/ \
			--go_opt=module=github.com/Fox216540/shop/proto/user-service \
			--go-grpc_out=user-service/ \
			--go-grpc_opt=module=github.com/Fox216540/shop/proto/user-service \
		user-service/api.proto
	@echo "✅ User Api proto generated"

gen-user-interservice:
	@make gen-common
	@mkdir -p user-service/gen
	protoc -I=. \
			--go_out=user-service/ \
			--go_opt=module=github.com/Fox216540/shop/proto/user-service \
			--go-grpc_out=user-service/ \
			--go-grpc_opt=module=github.com/Fox216540/shop/proto/user-service \
		user-service/interservice.proto
	@echo "✅ User Interservice proto generated"

gen-user:
	@make gen-user-api gen-user-interservice
	@echo "✅ User generated"


gen-all: gen-catalog gen-auth gen-order gen-user
