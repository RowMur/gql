package lexer

import (
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
	tokens := []Token{}
	runes := []rune(string(input))
	tokenizers := []Tokenizer{Punctuator{}, Name{}, IntValue{}, FloatValue{}, StringValue{}, Comma{}, Comment{}, LineTerminator{}, WhiteSpace{}, UnicodeBOM{}}

	var tokenizingError error
	currentRow := 1
	currentCol := 1
	var lineTerminatorIndex = 7
	for len(runes) > 0 && tokenizingError == nil {
		for i, tokenizer := range tokenizers {
			token, size, err := tokenizer.Test(&runes)
			if err != nil {
				tokenizingError = err
				break
			}
			if token == nil {
				if i == len(tokenizers)-1 {
					return nil, UnexpectedTokenError{string(runes[0]), currentRow, currentCol}
				}
				continue
			}

			if i == lineTerminatorIndex {
				currentRow++
				currentCol = 1
			} else {
				currentCol += size
			}

			isLexicalToken := slices.Index[[]string, string](lexicalTokens, token.Name)
			if isLexicalToken != -1 {
				tokens = append(tokens, *token)
			}

			runes = runes[size:]
			break
		}
	}

	return tokens, tokenizingError
}
