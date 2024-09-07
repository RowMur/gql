package lexer

import "strconv"

type UnexpectedTokenError struct {
	token string
	col   int
	row   int
}

func (e UnexpectedTokenError) Error() string {
	return "Unexpected token " + e.token + " at " + strconv.Itoa(e.col) + ":" + strconv.Itoa(e.row)
}
