.PHONY: run test cover proto

run:    ; DATABASE_URL=postgres://pvz:pvz@localhost:5432/pvz?sslmode=disable JWT_SECRET=dev go run ./cmd/service

test:   ; go test ./...

cover:
	go test -tags=integration -coverpkg=./internal/... -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | tail -1

proto:  ; buf generate
