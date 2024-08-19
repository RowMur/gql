package lexer

import "fmt"

type LineTerminator struct{}

func (lt LineTerminator) Test(runes *[]rune) (*Token, int, error) {
	// New line OKAY
	// Carriage return, !New line NOT OKAY
	// Carriage return, New line OKAY

	LF := rune(0x0A)
	CR := rune(0x0D)
	nextRune := (*runes)[0]
	name := "LineTerminator"

	if nextRune == LF {
		return &Token{Name: name, Value: string(nextRune)}, 1, nil
	}

	if nextRune != CR {
		return nil, 0, nil
	}

	nextNextRune := (*runes)[1]
	if len(*runes) == 1 || nextNextRune != LF {
		return nil, 0, fmt.Errorf("carriage return without New line")
	}

	return &Token{Name: name, Value: string(nextRune) + string(nextNextRune)}, 2, nil
}
