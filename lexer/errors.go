package lexer

import "strconv"

type UnexpectedTokenError struct {
	Token string
	Col   int
	Row   int
}

func (e UnexpectedTokenError) Error() string {
	return "Unexpected token " + e.Token + " at " + strconv.Itoa(e.Col) + ":" + strconv.Itoa(e.Row)
}
