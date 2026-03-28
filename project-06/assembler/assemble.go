package main

import (
	"fmt"
	"os"
	"strconv"

	parser "github.com/devnchill/Nand2Tetris/project-06/assembler/parser"
	symboltable "github.com/devnchill/Nand2Tetris/project-06/assembler/symbolTable"
	translator "github.com/devnchill/Nand2Tetris/project-06/assembler/translator"
)

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
	first := true
	for p.Advance() {
		commandType, err := p.GetCommandType()
		if err != nil {
			fmt.Printf("Invalid Insturction detected \n")
			fmt.Printf("%s was the currentCommand\n", p.CurrentCommand)
			panic(err)
		}
		fmt.Printf("Current Instruction -> %s\n", p.CurrentCommand)
		fmt.Printf("Current Instruction type -> %s\n", parser.GetCommandTypeInString(commandType))

		if commandType == parser.LCommand {
			fmt.Println("Lable found skipping ...")
			continue
		}

		if !first {
			f.WriteString("\n")
		}

		if commandType == parser.ACommand {
			symbol, _ := p.GetSymbol(commandType)
			fmt.Printf("[ symbol ] -> %s\n", symbol)
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
			fmt.Printf("[ address ] -> %s\n", strconv.Itoa(address))
			binaryValue := fmt.Sprintf("%016b", address)
			f.WriteString(binaryValue)
		} else if commandType == parser.CCommand {
			var dest, comp, jmp string
			if p.HasDest() {
				fmt.Printf("Dest -> %s\n", p.GetDest())
				destInBinary, err := t.TranslateDest(p.GetDest())
				if err != nil {
					fmt.Println(err)
					break
				}
				dest = destInBinary
				fmt.Printf("Dest in binary -> %s\n", destInBinary)
			} else {
				dest = "000"
			}
			fmt.Printf("Comp -> %s\n", p.GetComp())
			comp, err := t.TranslateComp(p.GetComp())
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Printf("Comp in binary -> %s\n", comp)

			if p.HasJump() {
				fmt.Printf("Jump -> %s\n", p.GetJump())
				jmp, err = t.TranslateJump(p.GetJump())
				if err != nil {
					fmt.Println(err)
					break
				}
				fmt.Printf("Jump in binary -> %s\n", jmp)

			} else {
				jmp = "000"
			}
			f.WriteString("111" + comp + dest + jmp)
		}
		first = false
	}
}
