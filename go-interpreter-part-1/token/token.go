package token 

import "fmt"

//token struct 
type Token struct {
	Type  string 
	Value interface{}
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%s, %v)", t.Type, t.Value)
}