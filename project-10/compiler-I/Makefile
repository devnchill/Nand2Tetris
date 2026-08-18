BIN := build/jack
SRC := $(shell find . -type f -name "*.go")

main:build

build:$(SRC)
	go build -o build/jack

run:$(BIN)
	./build/jack

$(BIN):$(SRC)
	go build -o build/jack
