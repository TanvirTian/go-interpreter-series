package interpreter 

import (
	"interpreter/constant"
	"interpreter/lexer"
	"interpreter/token"
)


type Interpreter struct {
	lexer        *lexer.Lexer 
	currentToken  token.Token 
}

func NewInterpreter(lexer *lexer.Lexer) *Interpreter {
	return &Interpreter{lexer: lexer, currentToken: lexer.GetNextToken()}
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
	if token.Type == constant.INTEGER {
		i.Eat(constant.INTEGER)
		return token.Value.(int)
	} else if token.Type ==  constant.LPAREN {
		i.Eat(constant.LPAREN)
		result := i.Expression()
		i.Eat(constant.RPAREN)
		return result
	}
	i.Error()
	return 0 
}


func (i *Interpreter) term() int {
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


func (i *Interpreter) Expression() int {
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
