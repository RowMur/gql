package lexer

type WhiteSpace struct{}

func (ws WhiteSpace) Test(runes *[]rune) (*Token, int, error) {
	space := rune(0x20)
	tab := rune(0x09)
	nextRune := (*runes)[0]
	name := "WhiteSpace"

	if nextRune == tab {
		return &Token{Name: name, Value: string(nextRune)}, 1, nil
	} else if nextRune == space {
		return &Token{Name: name, Value: string(nextRune)}, 1, nil
	}

	return nil, 0, nil
}
