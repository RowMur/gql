package lexer

import (
	"fmt"
)

type Token struct {
	Name  string
	Value string
}

type Tokenizer interface {
	Test(runes *[]rune) (*Token, int)
}

func Tokenize(input []byte) []Token {
	var tokens []Token
	runes := []rune(string(input))
	tokenizers := []Tokenizer{SourceCharacter{}}

	for len(runes) > 0 {
		for _, tokenizer := range tokenizers {
			token, size := tokenizer.Test(&runes)
			if token == nil {
				continue
			}

			tokens = append(tokens, *token)
			runes = runes[size:]
			fmt.Printf("%s: %q\n", token.Name, token.Value)
		}
	}

	return tokens
}
