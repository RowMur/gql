package lexer

type Punctuator struct{}

func (p Punctuator) Test(runes *[]rune) (*Token, int, error) {
	validPunctuators := map[string][]rune{
		"exclamationMark": {0x21},
		"dollarSign":      {0x24},
		"ampersand":       {0x26},
		"leftParen":       {0x28},
		"rightParen":      {0x29},
		"ellipsis":        {0x2E, 0x2E, 0x2E},
		"colon":           {0x3A},
		"equal":           {0x3D},
		"at":              {0x40},
		"leftSquare":      {0x5B},
		"rightSquare":     {0x5D},
		"leftCurly":       {0x7B},
		"pipe":            {0x7C},
		"rightCurly":      {0x7D},
	}

	for _, punctuator := range validPunctuators {
		for i := 0; i < len(punctuator); i++ {
			if punctuator[i] != (*runes)[i] {
				break
			}

			if i == len(punctuator)-1 {
				return &Token{Name: "Punctuator", Value: string(punctuator)}, len(punctuator), nil
			}
		}
	}

	return nil, 0, nil
}
