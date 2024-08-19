package lexer

type UnicodeBOM struct{}

func (ub UnicodeBOM) Test(runes *[]rune) (*Token, int, error) {
	BOM := rune(0xFEFF)
	nextRune := (*runes)[0]
	name := "UnicodeBOM"

	if nextRune == BOM {
		return &Token{Name: name, Value: string(nextRune)}, 1, nil
	}

	return nil, 0, nil
}
