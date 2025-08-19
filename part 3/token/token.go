package token 

import "fmt"

type Token struct {
	Type  string
	Value string
}

func (t Token) String() string {
	return fmt.Sprintf("%s, %s", t.Type, t.Value)
}
