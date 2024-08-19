package lexer

type Comment struct{}

func (c Comment) Test(runes *[]rune) (*Token, int, error) {
	hashtag := rune(0x23)
	nextRune := (*runes)[0]

	if nextRune != hashtag {
		return nil, 0, nil
	}

	runesInComment := []rune{nextRune}

	runesAfterHash := (*runes)[1:]
	for len(runesAfterHash) > 0 {
		token, _, err := LineTerminator{}.Test(&runesAfterHash)
		if err != nil {
			return nil, 0, err
		}

		if token == nil {
			runesInComment = append(runesInComment, runesAfterHash[0])
			runesAfterHash = runesAfterHash[1:]
			continue
		}

		break
	}

	return &Token{Name: "Comment", Value: string(runesInComment)}, len(runesInComment), nil
}
