package main

import (
	"fmt"
	"os"
	"strconv"

	parser "github.com/devnchill/Nand2Tetris/project-06/assembler/parser"
	symboltable "github.com/devnchill/Nand2Tetris/project-06/assembler/symbolTable"
	translator "github.com/devnchill/Nand2Tetris/project-06/assembler/translator"
)

func main() {
	filePath := os.Args[1]
	p, err := parser.NewParser(filePath)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	p2, err := parser.NewParser(filePath)
	t := translator.NewTranslator()
	table := symboltable.NewSymbolTable()
	firstPass(p, table)
	secondPass(p2, table, &t)
}

func firstPass(p *parser.Parser, table *symboltable.SymbolTable) {
	romAddress := 0
	for p.Advance() {
		commandType, err := p.GetCommandType()
		if err != nil {
			panic(err)
		}
		if commandType == parser.LCommand {
			symbol, err := p.GetSymbol(commandType)
			if err != nil {
				panic(err)
			}
			table.AddEntry(symbol, romAddress)
		} else {
			romAddress++
		}
	}
}
func secondPass(p *parser.Parser, table *symboltable.SymbolTable, t *translator.Translator) {
	ramAddress := 16
	fmt.Println("Creating Prog.hack")
	f, err := os.Create("Prog.hack")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for p.Advance() {
		commandType, err := p.GetCommandType()
		if err != nil {
			fmt.Printf("Invalid Insturction detected \n")
			fmt.Printf("%s was the currentCommand\n", p.CurrentCommand)
			panic(err)
		}
		if commandType == parser.LCommand {
			continue
		}
		if commandType == parser.ACommand {
			symbol, _ := p.GetSymbol(commandType)
			var address int
			if val, err := strconv.Atoi(symbol); err == nil {
				address = val
			} else {
				if !table.Contains(symbol) {
					table.AddEntry(symbol, ramAddress)
					ramAddress++
				}
				address = table.GetAddress(symbol)
			}
			binaryValue := fmt.Sprintf("%016b", address)
			f.WriteString(binaryValue + "\n")
		} else if commandType == parser.CCommand {
			f.WriteString("111")
			if p.HasDest() {
				fmt.Printf("dest -> %s\n", p.GetDest())
				destInBinary, err := t.TranslateDest(p.GetDest())
				f.WriteString(destInBinary)
				if err != nil {
					fmt.Println(err)
					break
				}
				fmt.Printf("dest in binary -> %s\n", destInBinary)
			} else {
				f.WriteString("000")
			}
			fmt.Printf("comp -> %s\n", p.GetComp())
			compInBinary, err := t.TranslateComp(p.GetComp())
			f.WriteString(compInBinary)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("comp in binary -> %s\n", compInBinary)

			if p.HasJump() {
				fmt.Printf("jump -> %s\n", p.GetJump())
				jumpInBinary, err := t.TranslateJump(p.GetJump())
				f.WriteString(jumpInBinary)
				if err != nil {
					fmt.Println(err)
					break
				}
				fmt.Printf("jump in binary -> %s\n", jumpInBinary)

			} else {
				f.WriteString("000")
			}
			f.WriteString("\n")
		}
	}
}
