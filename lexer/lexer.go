package lexer

import (
	"fmt"
	"slices"
)

type Token struct {
	Name  string
	Value string
}

type Tokenizer interface {
	Test(runes *[]rune) (*Token, int, error)
}

var lexicalTokens = []string{
	"Punctuator",
	"Name",
	"IntValue",
	"FloatValue",
	"StringValue",
}

func Tokenize(input []byte) ([]Token, error) {
	var tokens []Token
	runes := []rune(string(input))
	tokenizers := []Tokenizer{Punctuator{}, Name{}, IntValue{}, FloatValue{}, StringValue{}, Comma{}, Comment{}, LineTerminator{}, WhiteSpace{}, UnicodeBOM{}, SourceCharacter{}}

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

			isLexicalToken := slices.Index[[]string, string](lexicalTokens, token.Name)
			if isLexicalToken != -1 {
				tokens = append(tokens, *token)
				fmt.Printf("%s: %q\n", token.Name, token.Value)
			}

			runes = runes[size:]
			break
		}
	}

	return tokens, tokenizingError
}
