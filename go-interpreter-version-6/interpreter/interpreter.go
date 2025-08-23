package interpreter

import (
	"interpreter/token"
	"interpreter/parser"
	"interpreter/ast"
)

type Interpreter struct {
	parser *parser.Parser
}

func NewInterpreter(parser *parser.Parser) *Interpreter {
	return &Interpreter{parser: parser}
}

func (i *Interpreter) visit(node ast.AST) int {
	switch n := node.(type) {
	case ast.BinOp:
		if n.Op.Type == token.JOG {
			return i.visit(n.Left) + i.visit(n.Right)
		} else if n.Op.Type == token.BIYOG {
			return i.visit(n.Left) - i.visit(n.Right)
		} else if n.Op.Type == token.GUN {
			return i.visit(n.Left) * i.visit(n.Right)
		} else if n.Op.Type == token.BHAG {
			return i.visit(n.Left) / i.visit(n.Right)
		}
	case ast.Num:
		return n.Value
	}
	panic("No visit method")
}

func (i *Interpreter) Interpret() int {
	tree := i.parser.Parse()
	return i.visit(tree)
}
