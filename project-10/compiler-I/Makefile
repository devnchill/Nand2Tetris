BIN := build/jack
SRC := $(shell find . -type f -name "*.go")
main:$(BIN)
	./build/jack

build:$(BIN)

$(BIN):$(SRC)
	go build -o build/jack
