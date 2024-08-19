package lexer

type Comma struct{}

func (c Comma) Test(runes *[]rune) (*Token, int, error) {
	comma := rune(0x2C)
	nextRune := (*runes)[0]

	if nextRune != comma {
		return nil, 0, nil
	}

	return &Token{Name: "Comma", Value: string(nextRune)}, 1, nil
}
