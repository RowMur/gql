package lexer

type Digit struct{}

func (d Digit) Test(runes *[]rune) (*Token, int, error) {
	if len(*runes) == 0 {
		return nil, 0, nil
	}

	nextRune := (*runes)[0]

	if nextRune >= rune(0x30) && nextRune <= rune(0x39) {
		return &Token{Name: "Digit", Value: string(nextRune)}, 1, nil
	}

	return nil, 0, nil
}
