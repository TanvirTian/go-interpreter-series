package ast  

import (
	"interpreter/token"
)

type AST interface{}

type BinOp struct {
	Left  AST
	Op    token.Token
	Right AST
}

type Num struct {
	Token token.Token
	Value int
}

type UnaryOp struct {
	Token      token.Token 
	Op 	       token.Token 
	Expression AST
}