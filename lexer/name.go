package lexer

type Name struct{}

func (n Name) Test(runes *[]rune) (*Token, int, error) {
	underscore := rune(0x5F)
	letter := Letter{}
	digit := Digit{}

	firstRune := (*runes)[0]
	token, _, _ := letter.Test(runes)
	if token == nil && firstRune != underscore {
		return nil, 0, nil
	}

	runesInName := []rune{firstRune}
	remainingRunes := (*runes)[1:]
	for len(remainingRunes) > 0 {
		letterToken, _, _ := letter.Test(&remainingRunes)
		digitToken, _, _ := digit.Test(&remainingRunes)

		if letterToken == nil && digitToken == nil && remainingRunes[0] != underscore {
			break
		}

		runesInName = append(runesInName, remainingRunes[0])
		remainingRunes = remainingRunes[1:]
	}

	return &Token{Name: "Name", Value: string(runesInName)}, len(runesInName), nil
}
