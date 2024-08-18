package lexer

import (
	"fmt"
	"unicode/utf8"
)

type SourceCharacter struct{}

func (sc SourceCharacter) Test(runes *[]rune) (*Token, int) {
	tab := rune(0x09)
	LF := rune(0x0A)
	CR := rune(0x0D)
	minRune := rune(0x020)
	maxRune, _ := utf8.DecodeRune([]byte{0xE2, 0x86, 0x92})

	nextRune := (*runes)[0]
	name := "SourceCharacter"

	if nextRune == tab {
		return &Token{Name: name, Value: "\t"}, 1
	} else if nextRune == LF {
		return &Token{Name: name, Value: "\n"}, 1
	} else if nextRune == CR {
		return &Token{Name: name, Value: "\r"}, 1
	} else if minRune <= nextRune && nextRune <= maxRune {
		return &Token{Name: name, Value: string(nextRune)}, 1
	} else {
		fmt.Printf("unknown token: %q\n", nextRune)
	}

	return nil, 1
}
