package lexer

import (
	"regexp"
	"strings"
)

type StringValue struct{}

func (sv StringValue) Test(runes *[]rune) (*Token, int, error) {
	if strings.HasPrefix(string(*runes), "\"\"\"") {
		remainingRunes := (*runes)[3:]
		blockStringCharToken, _, _ := BlockStringCharacter{}.Test(&remainingRunes)
		if blockStringCharToken == nil {
			if strings.HasPrefix(string(remainingRunes), "\"\"\"") {
				return &Token{Name: "StringValue", Value: "\"\"\"\"\"\""}, 6, nil
			} else {
				return nil, 0, nil
			}
		}

		runesInString := []rune("\"\"\"" + blockStringCharToken.Value)
		for len(remainingRunes) > 0 {
			blockStringCharToken, size, _ := BlockStringCharacter{}.Test(&remainingRunes)
			if blockStringCharToken == nil {
				break
			}

			runesInString = append(runesInString, []rune(blockStringCharToken.Value)...)
			remainingRunes = remainingRunes[size:]
		}

		if strings.HasPrefix(string(remainingRunes), "\"\"\"") {
			return &Token{Name: "StringValue", Value: string(runesInString) + "\"\"\""}, 3 + len(runesInString) + 3, nil
		}

		return nil, 0, nil
	}

	if strings.HasPrefix(string(*runes), "\"\"") {
		return &Token{Name: "StringValue", Value: "\"\""}, 2, nil
	}

	if (*runes)[0] != '"' {
		return nil, 0, nil
	}

	remainingRunes := (*runes)[1:]
	runesInString := []rune{'"'}
	for len(remainingRunes) > 0 {
		stringCharacterToken, size, _ := StringCharacter{}.Test(&remainingRunes)
		if stringCharacterToken == nil {
			break
		}

		runesInString = append(runesInString, []rune(stringCharacterToken.Value)...)
		remainingRunes = remainingRunes[size:]
	}

	if remainingRunes[0] == '"' {
		runesInString = append(runesInString, '"')
		return &Token{Name: "StringValue", Value: string(runesInString)}, len(runesInString) + 1, nil
	}

	return nil, 0, nil
}

type StringCharacter struct{}

func (sc StringCharacter) Test(runes *[]rune) (*Token, int, error) {
	nextRune := (*runes)[0]

	lineTerminatorToken, _, _ := LineTerminator{}.Test(runes)
	if lineTerminatorToken != nil {
		return nil, 0, nil
	}

	if nextRune == '"' {
		return nil, 0, nil
	}

	if nextRune == '\\' {
		remainingRunes := (*runes)[1:]
		escapedCharacterToken, lenOfEscapedToken, _ := EscapedCharacter{}.Test(&remainingRunes)
		if escapedCharacterToken != nil {
			return &Token{Name: "StringCharacter", Value: string(nextRune) + escapedCharacterToken.Value}, 1 + lenOfEscapedToken, nil
		}

		if remainingRunes[0] == 'u' {
			remainingRunes = remainingRunes[1:]
			escapedUnicodeToken, lenOfEscapedToken, _ := EscapedUnicode{}.Test(&remainingRunes)
			if escapedUnicodeToken != nil {
				return &Token{Name: "StringCharacter", Value: string(nextRune) + escapedUnicodeToken.Value}, 2 + lenOfEscapedToken, nil
			}
		}
	}

	sourceCharacterToken, _, _ := SourceCharacter{}.Test(runes)
	if sourceCharacterToken == nil {
		return nil, 0, nil
	}

	return &Token{Name: "StringCharacter", Value: sourceCharacterToken.Value}, 1, nil
}

type EscapedUnicode struct{}

const escapedUnicodeRegex = "/[0-9A-Fa-f]{4}/"

func (eu EscapedUnicode) Test(runes *[]rune) (*Token, int, error) {
	regexp := regexp.MustCompile(escapedUnicodeRegex)
	res := regexp.Find([]byte(string(*runes)))
	if res == nil {
		return nil, 0, nil
	}

	return &Token{Name: "EscapedUnicode", Value: string(res)}, len(string(res)), nil
}

type EscapedCharacter struct{}

const escapedCharacterRegex = "[\"\\\\/bfnrt]$"

func (ec EscapedCharacter) Test(runes *[]rune) (*Token, int, error) {
	regexp := regexp.MustCompile(escapedCharacterRegex)
	res := regexp.Find([]byte(string(*runes)))
	if res == nil {
		return nil, 0, nil
	}

	return &Token{Name: "EscapedCharacter", Value: string(res)}, len(string(res)), nil
}

type BlockStringCharacter struct{}

func (bsc BlockStringCharacter) Test(runes *[]rune) (*Token, int, error) {
	sourceCharacterToken, _, _ := SourceCharacter{}.Test(runes)
	if sourceCharacterToken == nil {
		return nil, 0, nil
	}

	remaingRunesAsString := string(*runes)
	if strings.HasPrefix(remaingRunesAsString, "\"\"\"") {
		return nil, 0, nil
	}
	if strings.HasPrefix(remaingRunesAsString, "\\\"\"\"") {
		return nil, 0, nil
	}

	return &Token{Name: "BlockStringCharacter", Value: sourceCharacterToken.Value}, 1, nil
}
