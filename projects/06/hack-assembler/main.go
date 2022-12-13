package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// initialize symbol table
func initializeSymbolTable() map[string]int {
	symbolTable := make(map[string]int)
	for i := 0; i < 16; i++ {
		symbolTable["R"+strconv.Itoa(i)] = i
	}
	symbolTable["SCREEN"] = 16384
	symbolTable["KBD"] = 24576
	symbolTable["SP"] = 0
	symbolTable["LCL"] = 1
	symbolTable["ARG"] = 2
	symbolTable["THIS"] = 3
	symbolTable["THAT"] = 4

	return symbolTable
}

func initializeCompTable() map[string]string {
	compTable := make(map[string]string)
	// comp
	compTable["0"] = "0101010"
	compTable["1"] = "0111111"
	compTable["-1"] = "0111010"
	compTable["D"] = "0001100"
	compTable["A"] = "0110000"
	compTable["!D"] = "0001101"
	compTable["!A"] = "0110001"
	compTable["-D"] = "0001111"
	compTable["-A"] = "0110011"
	compTable["D+1"] = "0011111"
	compTable["A+1"] = "0110111"
	compTable["D-1"] = "0001110"
	compTable["A-1"] = "0110010"
	compTable["D+A"] = "0000010"
	compTable["D-A"] = "0010011"
	compTable["A-D"] = "0000111"
	compTable["D&A"] = "0000000"
	compTable["D|A"] = "0010101"
	compTable["M"] = "1110000"
	compTable["!M"] = "1110001"
	compTable["-M"] = "1110011"
	compTable["M+1"] = "1110111"
	compTable["M-1"] = "1110010"
	compTable["D+M"] = "1000010"
	compTable["D-M"] = "1010011"
	compTable["M-D"] = "1000111"
	compTable["D&M"] = "1000000"
	compTable["D|M"] = "1010101"

	return compTable
}

func initializeDestTable() map[string]string {
	destTable := make(map[string]string)

	// dest
	destTable["null"] = "000"
	destTable["M"] = "001"
	destTable["D"] = "010"
	destTable["MD"] = "011"
	destTable["A"] = "100"
	destTable["AM"] = "101"
	destTable["AD"] = "110"
	destTable["AMD"] = "111"


	return destTable
}
func initializeJmpTable() map[string]string {
	jmpTable := make(map[string]string)
	
	// jump
	jmpTable["null"] = "000"
	jmpTable["JGT"] = "001"
	jmpTable["JEQ"] = "010"
	jmpTable["JGE"] = "011"
	jmpTable["JLT"] = "100"
	jmpTable["JNE"] = "101"
	jmpTable["JLE"] = "110"
	jmpTable["JMP"] = "111"

	return jmpTable
}

func main() {
	args := os.Args
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Printf("error: opening file: %v\n", err)
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	symbolTable := initializeSymbolTable()
	destTable := initializeDestTable()
	compTable := initializeCompTable()
	jumpTable := initializeJmpTable()


	buf := bytes.Buffer{}
	// first pass
	n := 0
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "//") {
			text = text[:strings.Index(text, "//")]
		}
		text = strings.TrimSpace(text)
		if len(text) == 0 {
			continue
		}
		fmt.Fprintln(&buf, text)
		if text[0] == '(' && text[len(text)-1] == ')' {
			// this line will be ignored, do not increment n
			symbolTable[text[1:len(text)-1]] = n
		} else {
			n++
		}
	}

	n = 16
	output, err := os.Create(args[2])
	if err != nil {
		fmt.Printf("error: creating file: %v\n", err)
		panic(err)
	}

	writer := bufio.NewWriter(output)
	trimmedScanner := bufio.NewScanner(&buf)

	// second pass
	for trimmedScanner.Scan() {
		instruction := "0000000000000000"
		text := trimmedScanner.Text()

		if text[0] == '(' {
			continue
		}
		if text[0] == '@' {
			// A instruction
			value := text[1:]
			if num, err := strconv.Atoi(value); err == nil {
				// value is a number
				instruction = "0" + fmt.Sprintf("%015b", num)
			} else {
				address, exist := symbolTable[value]
				if exist {
					instruction = "0" + fmt.Sprintf("%015b", address)
				} else {
					symbolTable[value] = n
					instruction = "0" + fmt.Sprintf("%015b", n)
					n++
				}
			}
		} else {
			// C instruction
			var dest, comp, jump string
			if strings.Contains(text, "=") {
				equalsSignIndex := strings.Index(text, "=")
				dest = text[:equalsSignIndex]
				text = text[equalsSignIndex+1:]
			} else {
				dest = "null"
			}

			if strings.Contains(text, ";") {
				semicolonIndex := strings.Index(text, ";")
				comp = text[:semicolonIndex]
				jump = text[semicolonIndex+1:]
			} else {
				comp = text
				jump = "null"
			}

			instruction = "111" + compTable[comp] + destTable[dest] + jumpTable[jump]
		}
		writer.WriteString(instruction + "\n")
	}
	writer.Flush()
}
