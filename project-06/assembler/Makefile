SRC := main.go assemble.go
BIN := build/assembler

main:$(SRC)
	mkdir -p build
	go build -o build/assembler $(SRC)

run:main 
	./$(BIN) ../test_programs/Add.asm
clean:
	rm -rf build
