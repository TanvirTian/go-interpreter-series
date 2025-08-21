package lexer 

import (
	"strconv"
	"unicode"

	"interpreter/constant"
	"interpreter/token"
) 

type Lexer struct {
	text 		string
	pos 		int 
	currentChar rune 
}

func NewLexer(text string) *Lexer {
	runes := []rune(text)
	l := &Lexer{text: text, pos: 0}
	if len(runes) > 0 {
		l.currentChar = runes[0]
	} else {
		l.currentChar = 0 
	}
	return l 
}

func (l *Lexer) Error() {
	panic("Character Bhul")
}

func (l *Lexer) cursor() {
	l.pos++
	if l.pos > len([]rune(l.text))-1 {
		l.currentChar = 0 
	} else {
		l.currentChar = []rune(l.text)[l.pos]
	}
}

func (l *Lexer) SkipWhiteSpace() {
	for l.currentChar != 0 && unicode.IsSpace(l.currentChar) {
		l.cursor()
	}
}

func (l *Lexer) Integer() int {
	result := ""
	for l.currentChar != 0 && unicode.IsDigit(l.currentChar) {
		result += string(l.currentChar)
		l.cursor()
	}
	val, _ := strconv.Atoi(result)
	return val 
}


func (l *Lexer) GetNextToken() token.Token {
	for l.currentChar != 0 {
		if unicode.IsSpace(l.currentChar) {
			l.SkipWhiteSpace()
			continue
		}

		if unicode.IsDigit(l.currentChar) {
			val := l.Integer()
			return token.Token{Type: constant.INTEGER, Value: val}
		}


		switch l.currentChar {
		case '+':
			l.cursor()
			return token.Token{Type: constant.JOG, Value: "+"}

		case '-':
			l.cursor()
			return token.Token{Type: constant.BIYOG, Value: "-"}
		case '*':
			l.cursor()
			return token.Token{Type: constant.GUN, Value: "*"}	
		case '/':
			l.cursor()
			return token.Token{Type: constant.BHAG, Value: "/"}	
		case '(':
			l.cursor()
			return token.Token{Type: constant.LPAREN, Value: "("}	
		case ')':
			l.cursor()
			return token.Token{Type: constant.RPAREN, Value: ")"}	
		}

		 l.Error()
	}

	return token.Token{Type: constant.EOF, Value: nil}
}