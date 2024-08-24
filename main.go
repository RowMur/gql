package main

import (
	"fmt"

	"github.com/RowMur/gql/editor"
)

func main() {
	editor := editor.NewEditor()
	content := editor.Run()

	fmt.Println(*content)
}
