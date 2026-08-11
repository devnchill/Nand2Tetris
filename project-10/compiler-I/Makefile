BIN := build/jack
SRC := $(shell find . -type f -name "*.go")
main:$(BIN)
	./build/jack
$BIN:$(SRC)
	go build  -o build/jack

