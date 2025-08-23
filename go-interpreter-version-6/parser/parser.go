package parser 

import (
	"interpreter/token"
	"interpreter/lexer"
	"interpreter/ast"
)


type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{lexer: lexer, currentToken: lexer.GetNextToken()}
}

func (p *Parser) Eat(tokenType string) {
	if p.currentToken.Type == tokenType {
		p.currentToken = p.lexer.GetNextToken()
	} else {
		panic("Invalid syntax")
	}
}

// func (p *Parser) Factor() ast.AST {
// 	tok := p.currentToken
// 	if tok.Type == token.INTEGER {
// 		p.Eat(token.INTEGER)
// 		return ast.Num{token.Token: tok, Value: token.Value.(int)}
// 	} else if tok.Type == token.LPAREN {
// 		p.Eat(token.LPAREN)
// 		node := p.Expressionession()
// 		p.Eat(token.RPAREN)
// 		return node
// 	}
// 	panic("Syntax Bhul")
// }

func (p *Parser) Factor() ast.AST {
    tok := p.currentToken
    if tok.Type == token.INTEGER {
        p.Eat(token.INTEGER)
        return ast.Num{Token: tok, Value: tok.Value.(int)}
    } else if tok.Type == token.LPAREN {
        p.Eat(token.LPAREN)
        node := p.Expression()
        p.Eat(token.RPAREN)
        return node
    }
    panic("Syntax Bhul")
}


func (p *Parser) term() ast.AST {
	node := p.Factor()
	for p.currentToken.Type == token.GUN || p.currentToken.Type == token.BHAG {
		tok := p.currentToken
		if tok.Type == token.GUN {
			p.Eat(token.GUN)
		} else if tok.Type == token.BHAG {
			p.Eat(token.BHAG)
		}
		node = ast.BinOp{Left: node, Op: tok, Right: p.Factor()}
	}
	return node
}

func (p *Parser) Expression() ast.AST {
	node := p.term()
	for p.currentToken.Type == token.JOG || p.currentToken.Type == token.BIYOG {
		tok := p.currentToken
		if tok.Type == token.JOG {
			p.Eat(token.JOG)
		} else if tok.Type == token.BIYOG {
			p.Eat(token.BIYOG)
		}
		node = ast.BinOp{Left: node, Op: tok, Right: p.term()}
	}
	return node
}

func (p *Parser) Parse() ast.AST {
	return p.Expression()
}
