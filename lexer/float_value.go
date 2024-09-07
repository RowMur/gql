package lexer

type FloatValue struct{}

func (fv FloatValue) Test(runes *[]rune) (*Token, int, error) {
	_, lengthOfIntegerPart, _ := IntegerPart{}.Test(runes)
	if lengthOfIntegerPart == 0 {
		return nil, 0, nil
	}

	remaingRunes := (*runes)[lengthOfIntegerPart:]
	runesInFloatValue := (*runes)[:lengthOfIntegerPart]

	exponentToken, lengthOfExponentPart, _ := ExponentPart{}.Test(&remaingRunes)
	if exponentToken != nil {
		runesInFloatValue = append(runesInFloatValue, []rune(exponentToken.Value)...)
		remaingRunes = remaingRunes[lengthOfExponentPart:]
		nextRune := remaingRunes[0]

		token, _, _ := Digit{}.Test(&remaingRunes)
		if token != nil {
			return nil, 0, nil
		}

		dotRune := rune(0x2E)
		if nextRune == dotRune {
			return nil, 0, nil
		}

		token, _, _ = Letter{}.Test(&remaingRunes)
		if token != nil {
			return nil, 0, nil
		}

		underscoreToken := rune(0x5F)
		if nextRune == underscoreToken {
			return nil, 0, nil
		}

		return &Token{Name: "FloatValue", Value: string(runesInFloatValue)}, len(runesInFloatValue), nil
	}

	fractionalToken, lengthOfFractionalPart, _ := FractionalPart{}.Test(&remaingRunes)
	if fractionalToken == nil {
		return nil, 0, nil
	}

	runesInFloatValue = append(runesInFloatValue, []rune(fractionalToken.Value)...)
	remaingRunes = remaingRunes[lengthOfFractionalPart:]
	if len(remaingRunes) == 0 {
		return &Token{Name: "FloatValue", Value: string(runesInFloatValue)}, len(runesInFloatValue), nil
	}
	nextRune := remaingRunes[0]

	exponentToken, lengthOfExponentPart, _ = ExponentPart{}.Test(&remaingRunes)
	if exponentToken == nil {
		token, _, _ := Digit{}.Test(&remaingRunes)
		if token != nil {
			return nil, 0, nil
		}

		dotRune := rune(0x2E)
		if nextRune == dotRune {
			return nil, 0, nil
		}

		token, _, _ = Letter{}.Test(&remaingRunes)
		if token != nil {
			return nil, 0, nil
		}

		underscoreToken := rune(0x5F)
		if nextRune == underscoreToken {
			return nil, 0, nil
		}
		return &Token{Name: "FloatValue", Value: string(runesInFloatValue)}, len(runesInFloatValue), nil
	} else {
		runesInFloatValue = append(runesInFloatValue, []rune(exponentToken.Value)...)
		remaingRunes = remaingRunes[lengthOfExponentPart:]
	}

	token, _, _ := Digit{}.Test(&remaingRunes)
	if token != nil {
		return nil, 0, nil
	}

	dotRune := rune(0x2E)
	if nextRune == dotRune {
		return nil, 0, nil
	}

	token, _, _ = Letter{}.Test(&remaingRunes)
	if token != nil {
		return nil, 0, nil
	}

	underscoreToken := rune(0x5F)
	if nextRune == underscoreToken {
		return nil, 0, nil
	}

	return &Token{Name: "FloatValue", Value: string(runesInFloatValue)}, len(runesInFloatValue), nil
}

type FractionalPart struct{}

func (fp FractionalPart) Test(runes *[]rune) (*Token, int, error) {
	nextRune := (*runes)[0]
	dotRune := rune(0x2E)
	if nextRune != dotRune {
		return nil, 0, nil
	}

	runesInFractionalPart := []rune{nextRune}
	remaingRunes := (*runes)[1:]

	for len(remaingRunes) > 0 {
		token, _, _ := Digit{}.Test(&remaingRunes)
		if token == nil {
			break
		}

		nextRune := remaingRunes[0]
		runesInFractionalPart = append(runesInFractionalPart, nextRune)
		remaingRunes = remaingRunes[1:]
	}

	return &Token{Name: "FractionalPart", Value: string(runesInFractionalPart)}, len(runesInFractionalPart), nil
}

type ExponentPart struct{}

func (ep ExponentPart) Test(runes *[]rune) (*Token, int, error) {
	nextRune := (*runes)[0]
	token, _, _ := ExponentIndicator{}.Test(runes)
	if token == nil {
		return nil, 0, nil
	}

	runesInExponentPart := []rune{nextRune}
	remaingRunes := (*runes)[1:]

	token, _, _ = Sign{}.Test(&remaingRunes)
	if token != nil {
		nextRune := remaingRunes[0]
		runesInExponentPart = append(runesInExponentPart, nextRune)
		remaingRunes = remaingRunes[1:]
	}

	for len(remaingRunes) > 0 {
		token, _, _ := Digit{}.Test(&remaingRunes)
		if token == nil {
			break
		}

		nextRune := remaingRunes[0]
		runesInExponentPart = append(runesInExponentPart, nextRune)
		remaingRunes = remaingRunes[1:]
	}

	return &Token{Name: "ExponentPart", Value: string(runesInExponentPart)}, len(runesInExponentPart), nil
}

type ExponentIndicator struct{}

func (ei ExponentIndicator) Test(runes *[]rune) (*Token, int, error) {
	lowerE := rune(0x65)
	upperE := rune(0x45)

	nextRune := (*runes)[0]
	if nextRune != lowerE && nextRune != upperE {
		return nil, 0, nil
	}

	return &Token{Name: "ExponentIndicator", Value: string(nextRune)}, 1, nil
}

type Sign struct{}

func (s Sign) Test(runes *[]rune) (*Token, int, error) {
	plusSign := rune(0x2B)
	minusSign := rune(0x2D)

	nextRune := (*runes)[0]
	if nextRune != plusSign && nextRune != minusSign {
		return nil, 0, nil
	}

	return &Token{Name: "Sign", Value: string(nextRune)}, 1, nil
}
