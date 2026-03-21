SRC := main.go assemble.go

main:$(SRC)
	mkdir -p build
	go build -o build/assembler $(SRC)

clean:
	rm -rf build
