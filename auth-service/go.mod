module github.com/Fox216540/shop/auth-service

go 1.24.0

require (
	github.com/Fox216540/shop/proto v1.0.0
	github.com/caarlos0/env/v10 v10.0.0
	github.com/golang-jwt/jwt/v5 v5.3.0
	github.com/google/uuid v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/v9 v9.16.0
	github.com/rs/zerolog v1.34.0
	google.golang.org/grpc v1.77.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/net v0.47.0 // indirect
	golang.org/x/sys v0.38.0 // indirect
	golang.org/x/text v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251111163417-95abcf5c77ba // indirect
	google.golang.org/protobuf v1.36.10 // indirect
)

replace github.com/Fox216540/shop/proto => ../proto
