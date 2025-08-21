package token 

import "fmt"

type Token struct {
	Type  string 
	Value interface{}
}

func (t Token) String() string {
	return fmt.Sprintf("Token %s, %s", t.Type, t.Value)
}