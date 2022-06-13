# Makefile

build:
	go build -o cmd/script/ cmd/script/main.go

run: build
	./cmd/script/main

########################################################################################################################

test-unit:
	go test -race -v `go list ./... | grep -v ./internal/tools/handlers/errors`

test-unit-cover:
	go test -race -v -coverprofile ./docs/reports/tests/unit/cover.out `go list ./... | grep -v ./internal/tools/handlers/errors`

test-unit-cover-silent:
	go test -race -coverprofile ./docs/reports/tests/unit/cover.out `go list ./... | grep -v ./internal/tools/handlers/errors`

test-unit-cover-all:
	go test -race -v -coverprofile ./docs/reports/tests/unit/cover-all.out ./...

test-unit-cover-all-silent:
	go test -race -coverprofile ./docs/reports/tests/unit/cover-all.out ./...

test-unit-cover-report:
	go tool cover -html=./docs/reports/tests/unit/cover.out

test-unit-cover-all-report:
	go tool cover -html=./docs/reports/tests/unit/cover-all.out

########################################################################################################################
