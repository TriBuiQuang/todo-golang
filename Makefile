start:
	go run cmd/main.go

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