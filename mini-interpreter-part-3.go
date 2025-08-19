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
	GUN = "GUN"
	BHAG = "BHAG"
	EOF = "EOF"
)


type Token struct {
	Type  string
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("%s, %s", t.Type, t.Value)
}

type Lexer struct {
	text   		string 
	pos         int 
	currentChar rune 
}

func NewLexer(text string) *Lexer {
	l := &Lexer{text: text, pos: 0}
	if len(text) > 0 {
		l.currentChar = rune(text[0])
	} else {
		l.currentChar = 0
	}
	return l 
}

func (l *Lexer) error() {
	panic("Character Bhul")
}

func (l *Lexer) cursor() {
	l.pos++
	if l.pos > len(l.text)-1 {
		l.currentChar = 0 
	} else {
		l.currentChar = rune(l.text[l.pos])
	}
}

func (l *Lexer) skipWhiteSpace() {
	for l.currentChar != 0 && unicode.IsSpace(l.currentChar) {
		l.cursor()
	}
}

func (l *Lexer) integer() int {
	result := ""
	for l.currentChar != 0 && unicode.IsDigit(l.currentChar) {
		result += string(l.currentChar)
		l.cursor()
	}

	val, _ := strconv.Atoi(result)
	return val 
}

func (l *Lexer) getNextToken() Token {
	for l.currentChar != 0 {
		if unicode.IsSpace(l.currentChar) {
			l.skipWhiteSpace()
			continue 
		}

		if unicode.IsDigit(l.currentChar) {
			return Token{Type: INTEGER, Value: strconv.Itoa(l.integer())}
		}

		if l.currentChar == '*' {
			l.cursor()
			return Token{Type: GUN, Value: "*"}
		}

		if l.currentChar == '/' {
			l.cursor()
			return Token{Type: BHAG, Value: "/"}
		}

		l.error()
	}

	return Token{Type: EOF, Value: ""}
}

type Interpreter struct {
	lexer 		 *Lexer 
	currentToken Token 
}

func NewInterpreter(lexer *Lexer) *Interpreter {
	return &Interpreter{
		lexer: 		  lexer,
		currentToken: lexer.getNextToken(),
	}
}

func (i *Interpreter) error() {
	panic("Syntax Bhul")
}

func (i *Interpreter) eat(tokenType string) {
	if i.currentToken.Type == tokenType {
		i.currentToken = i.lexer.getNextToken()
	} else {
		i.error()
	}
}


func (i *Interpreter) factor() int {
	token := i.currentToken
	i.eat(INTEGER)
	val, _ := strconv.Atoi(token.Value)
	return val 
}

func (i *Interpreter) expression() int {
	result := i.factor()

	for i.currentToken.Type == GUN || i.currentToken.Type == BHAG {
		token := i.currentToken
		if token.Type == GUN {
			i.eat(GUN)
			result *= i.factor()
		} else if token.Type == BHAG {
			i.eat(BHAG)
			result /= i.factor()
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}

		lexer := NewLexer(text)
		interpreter := NewInterpreter(lexer)
		result := interpreter.expression()
		fmt.Println(result)
	}
}