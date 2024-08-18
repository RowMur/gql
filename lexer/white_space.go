package lexer

type WhiteSpace struct{}

func (ws WhiteSpace) Test(runes *[]rune) (*Token, int) {
	space := rune(0x20)
	tab := rune(0x09)
	nextRune := (*runes)[0]
	name := "WhiteSpace"

	if nextRune == tab {
		return &Token{Name: name, Value: string(nextRune)}, 1
	} else if nextRune == space {
		return &Token{Name: name, Value: string(nextRune)}, 1
	}

	return nil, 0
}
