package lexer

import (
	"fmt"
)

type Token struct {
	Name  string
	Value string
}

type Tokenizer interface {
	Test(runes *[]rune) (*Token, int, error)
}

func Tokenize(input []byte) ([]Token, error) {
	var tokens []Token
	runes := []rune(string(input))
	tokenizers := []Tokenizer{Punctuator{}, Name{}, IntValue{}, FloatValue{}, Comma{}, Comment{}, LineTerminator{}, WhiteSpace{}, UnicodeBOM{}, SourceCharacter{}}

	var tokenizingError error
	for len(runes) > 0 && tokenizingError == nil {
		for _, tokenizer := range tokenizers {
			token, size, err := tokenizer.Test(&runes)
			if err != nil {
				tokenizingError = err
				break
			}
			if token == nil {
				continue
			}

			tokens = append(tokens, *token)
			runes = runes[size:]
			fmt.Printf("%s: %q\n", token.Name, token.Value)
			break
		}
	}

	return tokens, tokenizingError
}
