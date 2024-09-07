package lexer_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/RowMur/gql/lexer"
)

func genericTokenizeTest(t *testing.T, gql string, want []lexer.Token, wantErr string) {
	t.Helper()

	res, err := lexer.Tokenize([]byte(gql))
	failedTestMessage := fmt.Sprintf("Tokenize(%q) = %v, %v, want %v, nil", gql, res, err, want)

	if err != nil && (wantErr == "" || err.Error() != wantErr) {
		t.Fatal(failedTestMessage)
	}

	if !reflect.DeepEqual(res, want) {
		t.Fatal(failedTestMessage)
	}
}

func TestTokenize(t *testing.T) {
	t.Run("Empty input", func(t *testing.T) {
		gql := ""
		want := []lexer.Token{}
		genericTokenizeTest(t, gql, want, "")
	})

	t.Run("Single tokens", func(t *testing.T) {
		t.Run("Punctuators", func(t *testing.T) {
			gql := "{}"
			want := []lexer.Token{
				{Name: "Punctuator", Value: "{"},
				{Name: "Punctuator", Value: "}"},
			}
			genericTokenizeTest(t, gql, want, "")
		})

		t.Run("Names", func(t *testing.T) {
			t.Run("Basic", func(t *testing.T) {
				gql := "name"
				want := []lexer.Token{
					{Name: "Name", Value: "name"},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("With numbers", func(t *testing.T) {
				gql := "name123"
				want := []lexer.Token{
					{Name: "Name", Value: "name123"},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("With underscores", func(t *testing.T) {
				gql := "_name_123"
				want := []lexer.Token{
					{Name: "Name", Value: "_name_123"},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("Name cannot start with a number", func(t *testing.T) {
				gql := "123name"
				genericTokenizeTest(t, gql, nil, "Unexpected token 1 at 1:1")
			})
		})

		t.Run("IntValue", func(t *testing.T) {
			t.Run("Basic", func(t *testing.T) {
				gql := "123"
				want := []lexer.Token{
					{Name: "IntValue", Value: "123"},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("Negative", func(t *testing.T) {
				gql := "-123"
				want := []lexer.Token{
					{Name: "IntValue", Value: "-123"},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("Negative with zero", func(t *testing.T) {
				gql := "-0"
				want := []lexer.Token{
					{Name: "IntValue", Value: "-0"},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("Cannot start with a zero", func(t *testing.T) {
				gql := "0123"
				genericTokenizeTest(t, gql, nil, "Unexpected token 0 at 1:1")
			})

			t.Run("Cannot be directly followed by a name", func(t *testing.T) {
				gql := "123name"
				genericTokenizeTest(t, gql, nil, "Unexpected token 1 at 1:1")
			})
		})

		t.Run("FloatValue", func(t *testing.T) {
			t.Run("Basic", func(t *testing.T) {
				gql := "123.456"
				want := []lexer.Token{
					{Name: "FloatValue", Value: "123.456"},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("Negative", func(t *testing.T) {
				gql := "-123.456"
				want := []lexer.Token{
					{Name: "FloatValue", Value: "-123.456"},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("Negative with zero", func(t *testing.T) {
				gql := "-0.456"
				want := []lexer.Token{
					{Name: "FloatValue", Value: "-0.456"},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("Cannot start with a zero", func(t *testing.T) {
				gql := "0123.456"
				genericTokenizeTest(t, gql, nil, "Unexpected token 0 at 1:1")
			})

			t.Run("Cannot be directly followed by a name", func(t *testing.T) {
				gql := "123.456name"
				genericTokenizeTest(t, gql, nil, "Unexpected token 1 at 1:1")
			})

			t.Run("exponents", func(t *testing.T) {
				t.Run("Basic", func(t *testing.T) {
					gql := "123.456e789"
					want := []lexer.Token{
						{Name: "FloatValue", Value: "123.456e789"},
					}
					genericTokenizeTest(t, gql, want, "")
				})

				t.Run("Negative", func(t *testing.T) {
					gql := "123.456e-789"
					want := []lexer.Token{
						{Name: "FloatValue", Value: "123.456e-789"},
					}
					genericTokenizeTest(t, gql, want, "")
				})

				t.Run("Negative with zero", func(t *testing.T) {
					gql := "123.456e-0"
					want := []lexer.Token{
						{Name: "FloatValue", Value: "123.456e-0"},
					}
					genericTokenizeTest(t, gql, want, "")
				})

				t.Run("Cannot start with a zero", func(t *testing.T) {
					gql := "0123.456e0123"
					genericTokenizeTest(t, gql, nil, "Unexpected token 0 at 1:1")
				})

				t.Run("Cannot be directly followed by a name", func(t *testing.T) {
					gql := "123.456e789name"
					genericTokenizeTest(t, gql, nil, "Unexpected token 1 at 1:1")
				})
			})
		})

		t.Run("String values", func(t *testing.T) {
			t.Run("Basic", func(t *testing.T) {
				gql := `"string"`
				want := []lexer.Token{
					{Name: "StringValue", Value: `"string"`},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("Empty", func(t *testing.T) {
				gql := `""`
				want := []lexer.Token{
					{Name: "StringValue", Value: `""`},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("Empty multi line", func(t *testing.T) {
				gql := `""""""`
				want := []lexer.Token{
					{Name: "StringValue", Value: `""""""`},
				}
				genericTokenizeTest(t, gql, want, ``)
			})

			t.Run("No closing", func(t *testing.T) {
				gql := `"string`
				genericTokenizeTest(t, gql, nil, "Unexpected token \" at 1:1")
			})

			t.Run("Multi line with no closing", func(t *testing.T) {
				gql := `"""string`
				genericTokenizeTest(t, gql, nil, "Unexpected token \" at 1:1")
			})

			t.Run("Multi line", func(t *testing.T) {
				gql := `"""string"""`
				want := []lexer.Token{
					{Name: "StringValue", Value: `"""string"""`},
				}
				genericTokenizeTest(t, gql, want, "")
			})

			t.Run("Multi line with escaped multi line", func(t *testing.T) {
				gql := `"""string\""""""`
				want := []lexer.Token{
					{Name: "StringValue", Value: `"""string\""""""`},
				}
				genericTokenizeTest(t, gql, want, "")
			})
		})
	})

	t.Run("Proper GQL", func(t *testing.T) {
		t.Run("Simple query", func(t *testing.T) {
			gql := `
				query {
					hero {
						name
					}
				}
			`
			want := []lexer.Token{
				{Name: "Name", Value: "query"},
				{Name: "Punctuator", Value: "{"},
				{Name: "Name", Value: "hero"},
				{Name: "Punctuator", Value: "{"},
				{Name: "Name", Value: "name"},
				{Name: "Punctuator", Value: "}"},
				{Name: "Punctuator", Value: "}"},
			}
			genericTokenizeTest(t, gql, want, "")
		})

		t.Run("Comments", func(t *testing.T) {
			gql := `
				# This is a comment
				query {
					# This is another comment
					hero {
						name
					}
				}
			`
			want := []lexer.Token{
				{Name: "Name", Value: "query"},
				{Name: "Punctuator", Value: "{"},
				{Name: "Name", Value: "hero"},
				{Name: "Punctuator", Value: "{"},
				{Name: "Name", Value: "name"},
				{Name: "Punctuator", Value: "}"},
				{Name: "Punctuator", Value: "}"},
			}
			genericTokenizeTest(t, gql, want, "")
		})

		t.Run("Variables", func(t *testing.T) {
			gql := `
				query ($id: ID!) {
					hero(id: $id) {
						name
					}
				}
			`
			want := []lexer.Token{
				{Name: "Name", Value: "query"},
				{Name: "Punctuator", Value: "("},
				{Name: "Punctuator", Value: "$"},
				{Name: "Name", Value: "id"},
				{Name: "Punctuator", Value: ":"},
				{Name: "Name", Value: "ID"},
				{Name: "Punctuator", Value: "!"},
				{Name: "Punctuator", Value: ")"},
				{Name: "Punctuator", Value: "{"},
				{Name: "Name", Value: "hero"},
				{Name: "Punctuator", Value: "("},
				{Name: "Name", Value: "id"},
				{Name: "Punctuator", Value: ":"},
				{Name: "Punctuator", Value: "$"},
				{Name: "Name", Value: "id"},
				{Name: "Punctuator", Value: ")"},
				{Name: "Punctuator", Value: "{"},
				{Name: "Name", Value: "name"},
				{Name: "Punctuator", Value: "}"},
				{Name: "Punctuator", Value: "}"},
			}
			genericTokenizeTest(t, gql, want, "")
		})
	})
}
