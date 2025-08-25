package token 

import "fmt"


const (
	INTEGER = "INTEGER"
	JOG    = "JOG"
	BIYOG   = "BIYOG"
	GUN     = "GUN"
	BHAG     = "BHAG"
	LPAREN  = "("
	RPAREN  = ")"
	EOF     = "EOF"
)


type Token struct {
	Type  string
	Value interface{}
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%s, %v)", t.Type, t.Value)
}

