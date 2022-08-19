start:
	go run cmd/main.go

start-server:
	go run cmd/main.go -protocol=server

start-client:
	go run cmd/main.go -protocol=client

test-unit:
	go test ./tests/unit/... -coverpkg=./internal/...

test-unit-detail:
	go test -v -race ./tests/unit/... -coverpkg=./internal/...

test-coverage-with-validation:
	go test ./tests/... -v -coverpkg=./internal/...

test-coverage:
	go test --cover -covermode=count ./tests/...  -coverpkg=./internal/... -coverprofile ./coverage.out
	go tool cover -func ./coverage.out
	go tool cover -html=coverage.out -o coverage.html
	rm -f coverage.out

test-integration:
	go test ./tests/integration/... -coverpkg=./internal/...

test-cucumber: 
	godog run ./tests/integration/

gen-grpc: 
	rm -f ./pkg/grpc/*.go
	protoc --proto_path=proto --go_out=pkg/grpc --go_opt=paths=source_relative\
		--go-grpc_out=pkg/grpc --go-grpc_opt=paths=source_relative\
		./proto/*.proto

migrate-up:
	migrate -path migration -database "postgresql://username:123@localhost:5432/test?sslmode=disable" -verbose up
	
migrate-down:
	migrate -path migration -database "postgresql://username:123@localhost:5432/test?sslmode=disable" -verbose down