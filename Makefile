start:
	go run cmd/main.go

test-unit:
	go test ./tests/unit/... -coverpkg=./internal/...

test-unit-detail:
	go test -v -race ./tests/unit/... -coverpkg=./internal/...

test-coverage-with-validation:
	go test ./tests/... -v -coverpkg=./internal/...

test-coverage:
	go test ./tests/...  -coverpkg=./internal/... -coverprofile ./coverage.out
	go tool cover -func ./coverage.out
	rm -f coverage.out

test-integration:
	go test -test.v -test.run ^TestFeatures$