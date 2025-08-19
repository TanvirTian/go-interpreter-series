package interpreter 

import (
	"strconv"
	"unicode"

	"interpreter/token"
	"interpreter/constant"
)


type Interpreter struct {
	text         string 
	pos          int
	currentChar  rune 
	currentToken token.Token 
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

func (i *Interpreter) Error()  {
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

func (i *Interpreter) SkipWhiteSpace() {
	for i.currentChar != 0 && unicode.IsSpace(i.currentChar) {
		i.cursor()
	}
}

func (i *Interpreter) Integer() int {
	result := ""
	for i.currentChar != 0 && unicode.IsDigit(i.currentChar) {
		result += string(i.currentChar)
		i.cursor()
	}

	val, _ := strconv.Atoi(result)
	return val 
}


func (i *Interpreter) GetNextToken() token.Token {
	for i.currentChar != 0 {
		if unicode.IsSpace(i.currentChar) {
			i.SkipWhiteSpace()
			continue 
		}

		if unicode.IsDigit(i.currentChar) {
			return token.Token{Type: constant.INTEGER, Value: strconv.Itoa(i.Integer())}
		}

		if i.currentChar == '+' {
			i.cursor()
			return token.Token{Type: constant.JOG, Value: "+"}
		}

		if i.currentChar == '-' {
			i.cursor()
			return token.Token{Type: constant.BIYOG, Value: "-"}
		}

		i.Error()
	}

	return token.Token{Type: constant.EOF, Value: ""}
}


func (i *Interpreter) Eat(tokenType string) {
	if i.currentToken.Type == tokenType{
		i.currentToken = i.GetNextToken()
	} else {
		i.Error()
	}
}

func (i *Interpreter) term() int {
	token := i.currentToken
	i.Eat(constant.INTEGER)
	val, _ := strconv.Atoi(token.Value)
	return val 
}

func (i *Interpreter) Expression() int {
	i.currentToken = i.GetNextToken()
	result := i.term()

	for i.currentToken.Type == constant.JOG || i.currentToken.Type == constant.BIYOG {
		token := i.currentToken
		if token.Type == constant.JOG {
			i.Eat(constant.JOG)
			result += i.term()
		} else if token.Type == constant.BIYOG {
			i.Eat(constant.BIYOG)
			result -= i.term()
		}
	}

	return result
}
