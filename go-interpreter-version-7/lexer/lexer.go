package lexer 

import (
	"strconv"
	"unicode"
	
	"interpreter/token"
)

type Lexer struct {
	text        string
	pos         int
	currentChar *rune
}

func NewLexer(text string) *Lexer {
	l := &Lexer{text: text, pos: 0}
	if len(text) > 0 {
		r := rune(text[0])
		l.currentChar = &r
	}
	return l
}

func (l *Lexer) cursor() {
	l.pos++
	if l.pos > len(l.text)-1 {
		l.currentChar = nil
	} else {
		r := rune(l.text[l.pos])
		l.currentChar = &r
	}
}

func (l *Lexer) SkipWhitespace() {
	for l.currentChar != nil && unicode.IsSpace(*l.currentChar) {
		l.cursor()
	}
}

func (l *Lexer) Integer() int {
	result := ""
	for l.currentChar != nil && unicode.IsDigit(*l.currentChar) {
		result += string(*l.currentChar)
		l.cursor()
	}
	val, _ := strconv.Atoi(result)
	return val
}

func (l *Lexer) GetNextToken() token.Token {
	for l.currentChar != nil {
		
		if unicode.IsSpace(*l.currentChar) {
			l.SkipWhitespace()
			continue
		}

		if unicode.IsDigit(*l.currentChar) {
			return token.Token{token.INTEGER, l.Integer()}
		}

		switch *l.currentChar {
		case '+':
			l.cursor()
			return token.Token{token.JOG, "+"}
		case '-':
			l.cursor()
			return token.Token{token.BIYOG, "-"}
		case '*':
			l.cursor()
			return token.Token{token.GUN, "*"}
		case '/':
			l.cursor()
			return token.Token{token.BHAG, "/"}
		case '(':
			l.cursor()
			return token.Token{token.LPAREN, "("}
		case ')':
			l.cursor()
			return token.Token{token.RPAREN, ")"}
		default:
			panic("Character Bhul: Abar Check Korun.")
		}
	}
	return token.Token{token.EOF, nil}
}

