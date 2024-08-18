package lexer

type UnicodeBOM struct{}

func (ub UnicodeBOM) Test(runes *[]rune) (*Token, int) {
	BOM := rune(0xFEFF)
	nextRune := (*runes)[0]
	name := "UnicodeBOM"

	if nextRune == BOM {
		return &Token{Name: name, Value: string(nextRune)}, 1
	}

	return nil, 0
}
