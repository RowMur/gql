package lexer

type Letter struct{}

func (l Letter) Test(runes *[]rune) (*Token, int, error) {
	nextRune := (*runes)[0]

	if nextRune >= rune(0x41) && nextRune <= rune(0x5A) {
		return &Token{Name: "Letter", Value: string(nextRune)}, 1, nil
	} else if nextRune >= rune(0x61) && nextRune <= rune(0x7A) {
		return &Token{Name: "Letter", Value: string(nextRune)}, 1, nil
	}

	return nil, 0, nil
}
