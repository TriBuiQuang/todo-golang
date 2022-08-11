start:
	go run cmd/main.go

test-unit:
	go test ./tests/unit/... -coverpkg=./internal/...

test-coverage-with-validation:
	go test ./tests/... -v -coverpkg=./internal/...

test-coverage-clearly:
	go test ./tests/...  -coverpkg=./internal/... -coverprofile ./coverage.out
	go tool cover -func ./coverage.out