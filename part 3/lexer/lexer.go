package lexer

import (
	"strconv"
	"unicode"

	"interpreter/constant"
	"interpreter/token"
)

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

func (l *Lexer) Error() {
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
			return token.Token{Type: constant.INTEGER, Value: strconv.Itoa(l.Integer())}
		}

		if l.currentChar == '*' {
			l.cursor()
			return token.Token{Type: constant.GUN, Value: "*"}
		}

		if l.currentChar == '/' {
			l.cursor()
			return token.Token{Type: constant.BHAG, Value: "/"}
		}

		l.Error()
	}

	return token.Token{Type: constant.EOF, Value: ""}
}