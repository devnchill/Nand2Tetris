BIN:= build/vm 
SRC := $(shell find . -type f -name "*.go")

$(BIN):$(SRC)
	mkdir -p build
	go build -o $(BIN) .

main:$(BIN)

run: $(BIN)
	./build/vm
