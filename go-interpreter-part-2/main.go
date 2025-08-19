package main 

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"interpreter/interpreter"
)






func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		interpreter := interpreter.NewInterpreter(text) 
		result := interpreter.Expression()
		fmt.Println(result)
	}
}