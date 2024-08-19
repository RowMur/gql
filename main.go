package main

import (
	"fmt"
	"os"

	"github.com/RowMur/gql/lexer"
)

const filePath = "./example-gql/query.gql"

func main() {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	_, err = lexer.Tokenize(file)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
}
