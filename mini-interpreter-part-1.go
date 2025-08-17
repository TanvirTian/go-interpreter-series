package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	INTEGER = "INTEGER"
	JOG = "JOG"
	BIYOG = "BIYOG"
	EOF = "EOF"

)

//token struct 
type Token struct {
	Type  string 
	Value interface{}
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%s, %v)", t.Type, t.Value)
}

//interpreter struct
type Interpreter struct {
	text         string 
	pos          int 
	currentToken Token 
	currentChar  rune
}

func NewInterpreter(text string) *Interpreter {
	i := &Interpreter{text: text, pos: 0}
	if len(text) > 0 {
		i.currentChar = rune(text[0])
	} else {
		i.currentChar = 0 
	}
	return i 
}

func (i *Interpreter) error() {
	panic("ERROR parsing input")
}

func (i *Interpreter) advance() {
	i.pos++
	if i.pos > len(i.text)-1 {
		i.currentChar = 0 //end of input
	} else {
		i.currentChar = rune(i.text[i.pos])
	}
}

func (i *Interpreter) skipWhiteSpace() {
	for i.currentChar != 0 && unicode.IsSpace(i.currentChar) {
		i.advance()
	}
}

func (i *Interpreter) integer() int {
	result := ""
	for i.currentChar != 0 && unicode.IsDigit(i.currentChar) {
		result += string(i.currentChar)
		i.advance()
	}

	val, _ := strconv.Atoi(result)
	return val 
}

func (i *Interpreter) getNextToken() Token {
	for i.currentChar != 0 {
		if unicode.IsSpace(i.currentChar) {
			i.skipWhiteSpace()
			continue
		}

		if unicode.IsDigit(i.currentChar) {
			return Token{Type: INTEGER, Value: i.integer()}
		}

		if i.currentChar == '+' {
			i.advance()
			return Token{Type: JOG, Value: "+"}
		}

		if i.currentChar == '-' {
			i.advance()
			return Token{Type: BIYOG, Value: "-"}
		}

		i.error()
	}
	return Token{Type: EOF, Value: nil}
}

func (i *Interpreter) eat(tokenType string) {
	if i.currentToken.Type == tokenType {
		i.currentToken = i.getNextToken()
	} else {
		i.error()
	}
}

func (i *Interpreter) expr() int {
	//first token
	i.currentToken  =i.getNextToken()

	left := i.currentToken
	i.eat(INTEGER)

	op := i.currentToken
	if op.Type == JOG {
		i.eat(JOG)
	} else {
		i.eat(BIYOG)
	}

	right := i.currentToken
	i.eat(INTEGER)

	//calculate result
	if op.Type == JOG {
		return left.Value.(int) + right.Value.(int)
	} else {
		return left.Value.(int) - right.Value.(int)
	}
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}

		interpreter := NewInterpreter(text)
		result := interpreter.expr()
		fmt.Println(result)
	}
}