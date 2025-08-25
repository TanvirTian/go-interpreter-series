package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"interpreter/lexer"
	"interpreter/interpreter"
	"interpreter/parser"
	
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		lexer := lexer.NewLexer(text)
		parser := parser.NewParser(lexer)
		interpreter := interpreter.NewInterpreter(parser)
		result := interpreter.Interpret()
		fmt.Println(result)
	}
}
