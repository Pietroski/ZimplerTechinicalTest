# Makefile

build:
	go build -o cmd/script/ cmd/script/main.go

run: build
	./cmd/script/main
