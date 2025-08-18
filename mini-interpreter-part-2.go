package main 

import (
	"fmt"
	"bufio"
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

type Token struct {
	Type  string 
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("Token %s, %s", t.Type, t.Value)
}

type Interpreter struct {
	text         string 
	pos          int
	currentChar  rune 
	currentToken Token 
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

func (i *Interpreter) error()  {
	panic("Syntax Bhul")
}

func (i *Interpreter) cursor() {
	i.pos++ 
	if i.pos > len(i.text)-1 { 
		i.currentChar = 0
	} else {
		i.currentChar = rune(i.text[i.pos])
	}
} 

func (i *Interpreter) skipWhiteSpace() {
	for i.currentChar != 0 && unicode.IsSpace(i.currentChar) {
		i.cursor()
	}
}

func (i *Interpreter) integer() int {
	result := ""
	for i.currentChar != 0 && unicode.IsDigit(i.currentChar) {
		result += string(i.currentChar)
		i.cursor()
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
			return Token{Type: INTEGER, Value: strconv.Itoa(i.integer())}
		}

		if i.currentChar == '+' {
			i.cursor()
			return Token{Type: JOG, Value: "+"}
		}

		if i.currentChar == '-' {
			i.cursor()
			return Token{Type: BIYOG, Value: "-"}
		}

		i.error()
	}

	return Token{Type: EOF, Value: ""}
}


func (i *Interpreter) eat(tokenType string) {
	if i.currentToken.Type == tokenType{
		i.currentToken = i.getNextToken()
	} else {
		i.error()
	}
}

func (i *Interpreter) term() int {
	token := i.currentToken
	i.eat(INTEGER)
	val, _ := strconv.Atoi(token.Value)
	return val 
}

func (i *Interpreter) expression() int {
	i.currentToken = i.getNextToken()
	result := i.term()

	for i.currentToken.Type == JOG || i.currentToken.Type == BIYOG {
		token := i.currentToken
		if token.Type == JOG {
			i.eat(JOG)
			result += i.term()
		} else if token.Type == BIYOG {
			i.eat(BIYOG)
			result -= i.term()
		}
	}

	return result
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
		result := interpreter.expression()
		fmt.Println(result)
	}
}