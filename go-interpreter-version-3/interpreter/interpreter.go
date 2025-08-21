package interpreter 

import (
	"strconv"

	"interpreter/constant"
	"interpreter/lexer"
	"interpreter/token"
)

type Interpreter struct {
	lexer 		 *lexer.Lexer
	currentToken token.Token 
}

func NewInterpreter(lexer *lexer.Lexer) *Interpreter {
	return &Interpreter{
		lexer: 		  lexer,
		currentToken: lexer.GetNextToken(),
	}
}

func (i *Interpreter) Error() {
	panic("Syntax Bhul")
}

func (i *Interpreter) Eat(tokenType string) {
	if i.currentToken.Type == tokenType {
		i.currentToken = i.lexer.GetNextToken()
	} else {
		i.Error()
	}
}


func (i *Interpreter) Factor() int {
	token := i.currentToken
	i.Eat(constant.INTEGER)
	val, _ := strconv.Atoi(token.Value)
	return val 
}

func (i *Interpreter) Expression() int {
	result := i.Factor()

	for i.currentToken.Type == constant.GUN || i.currentToken.Type == constant.BHAG {
		token := i.currentToken
		if token.Type == constant.GUN {
			i.Eat(constant.GUN)
			result *= i.Factor()
		} else if token.Type == constant.BHAG {
			i.Eat(constant.BHAG)
			result /= i.Factor()
		}
	}
	return result
}
