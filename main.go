package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/RowMur/gql/editor"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("endpoint not set.")
	}

	endpoint := args[len(args)-1]
	if endpoint == "" {
		panic("endpoint not set.")
	}

	editor := editor.NewEditor()
	content := editor.Run()

	if *content == "" {
		return
	}

	reqBody := map[string]string{
		"query": *content,
	}

	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}
	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonReqBody))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var body interface{}
	json.NewDecoder(res.Body).Decode(&body)
	fmt.Printf("%+v\n", body)
}
