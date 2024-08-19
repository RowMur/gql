package lexer

type IntValue struct{}

func (iv IntValue) Test(runes *[]rune) (*Token, int, error) {
	token, size, _ := IntegerPart{}.Test(runes)
	if token == nil {
		return nil, 0, nil
	}

	remainingRunes := (*runes)[size:]
	nextRune := remainingRunes[0]
	dotRune := rune(0x2E)
	if nextRune == dotRune {
		return nil, 0, nil
	}

	underscoreRune := rune(0x5F)
	letterToken, _, _ := Letter{}.Test(&remainingRunes)
	digitToken, _, _ := Digit{}.Test(&remainingRunes)
	if letterToken != nil || nextRune == underscoreRune || digitToken != nil {
		return nil, 0, nil
	}

	return &Token{Name: "IntValue", Value: token.Value}, size, nil
}

type IntegerPart struct{}

func (ip IntegerPart) Test(runes *[]rune) (*Token, int, error) {
	remaingRunes := *runes
	runesInIntegerPart := []rune{}

	token, _, _ := NegativeSign{}.Test(&remaingRunes)
	if token != nil {
		nextRune := remaingRunes[0]
		runesInIntegerPart = append(runesInIntegerPart, nextRune)
		remaingRunes = remaingRunes[1:]
	}

	nextRune := remaingRunes[0]
	zeroRune := rune(0x30)
	if nextRune == zeroRune {
		runesInIntegerPart = append(runesInIntegerPart, nextRune)
		return &Token{Name: "IntegerPart", Value: string(runesInIntegerPart)}, len(runesInIntegerPart), nil
	} else {
		token, _, _ := NonZeroDigit{}.Test(&remaingRunes)
		if token == nil {
			return nil, 0, nil
		}

		runesInIntegerPart = append(runesInIntegerPart, nextRune)
		remaingRunes = remaingRunes[1:]
	}

	for len(remaingRunes) > 0 {
		token, _, _ := Digit{}.Test(&remaingRunes)
		if token == nil {
			break
		}

		nextRune := remaingRunes[0]
		runesInIntegerPart = append(runesInIntegerPart, nextRune)
		remaingRunes = remaingRunes[1:]
	}

	return &Token{Name: "IntegerPart", Value: string(runesInIntegerPart)}, len(runesInIntegerPart), nil
}

type NegativeSign struct{}

func (ns NegativeSign) Test(runes *[]rune) (*Token, int, error) {
	nextRune := (*runes)[0]
	negativeSign := rune(0x2D)

	if nextRune != negativeSign {
		return nil, 0, nil
	}

	return &Token{Name: "NegativeSign", Value: string(nextRune)}, 1, nil
}

type NonZeroDigit struct{}

func (nzd NonZeroDigit) Test(runes *[]rune) (*Token, int, error) {
	nextRune := (*runes)[0]
	zeroRune := rune(0x30)

	if nextRune == zeroRune {
		return nil, 0, nil
	}

	DigitToken, _, _ := Digit{}.Test(runes)
	if DigitToken == nil {
		return nil, 0, nil
	}

	return &Token{Name: "NonZeroDigit", Value: string(nextRune)}, 1, nil
}
