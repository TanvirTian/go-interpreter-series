package main 

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"interpreter/interpreter"
	"interpreter/lexer"	
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
		interpreter := interpreter.NewInterpreter(lexer)
		result := interpreter.Expression()
		fmt.Println(result)
	}
}