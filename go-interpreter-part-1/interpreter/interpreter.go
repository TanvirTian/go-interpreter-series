package interpreter 

import (
	"strconv"
	"unicode"
	
	"interpreter/constant"
	"interpreter/token"
	
)

//interpreter struct
type Interpreter struct {
	text         string 
	pos          int 
	currentToken token.Token 
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

func (i *Interpreter) Error() {
	panic("Syntax Bhul")
}

func (i *Interpreter) cursor() {
	i.pos++
	if i.pos > len(i.text)-1 {
		i.currentChar = 0 //end of input
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
			return token.Token{Type: constant.INTEGER, Value: i.Integer()}
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
	return token.Token{Type: constant.EOF, Value: nil}
}

func (i *Interpreter) Eat(tokenType string) {
	if i.currentToken.Type == tokenType {
		i.currentToken = i.GetNextToken()
	} else {
		i.Error()
	}
}

func (i *Interpreter) Expression() int {
	//first token
	i.currentToken  =i.GetNextToken()

	left := i.currentToken
	i.Eat(constant.INTEGER)

	op := i.currentToken
	if op.Type == constant.JOG {
		i.Eat(constant.JOG)
	} else {
		i.Eat(constant.BIYOG)
	}

	right := i.currentToken
	i.Eat(constant.INTEGER)

	//calculate result
	if op.Type == constant.JOG {
		return left.Value.(int) + right.Value.(int)
	} else {
		return left.Value.(int) - right.Value.(int)
	}
}
