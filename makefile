.PHONY: dev
dev:
	air

.PHONY: migration
create-migration:
	migrate create -ext=sql -dir=sql/migrations -seq $(name)

.PHONY: migrate
migrate:
	migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

.PHONY: generate-grpc
generate-grpc:
	protoc --go_out=. --go-grpc_out=. ./internal/infra/grpc/protofiles/*.proto