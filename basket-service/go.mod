module github.com/Fox216540/shop/basket-service

go 1.25.0

require (
	github.com/Fox216540/shop/proto v1.0.0
	github.com/caarlos0/env/v10 v10.0.0
	github.com/google/uuid v1.6.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.35.1
	github.com/shopspring/decimal v1.4.0
	google.golang.org/grpc v1.81.1
	google.golang.org/protobuf v1.36.11
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.31.1
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.10.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.22 // indirect
	golang.org/x/crypto v0.53.0 // indirect
	golang.org/x/net v0.55.0 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.38.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260608224507-4308a22a1bab // indirect
)

replace github.com/Fox216540/shop/proto => ../proto
