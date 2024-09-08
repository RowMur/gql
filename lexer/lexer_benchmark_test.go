package lexer_test

import (
	"testing"

	"github.com/RowMur/gql/lexer"
)

const gql = `
# This is a query for the Star Wars API
query Query {
  allFilms(q: "star wars", first: 10, someVar: 12345.6789e25) {
    films {
      __typename
      title # I want to get the title of the film
      director
      releaseDate
      speciesConnection {
        species {
          name
          classification
          homeworld {
            name
          }
        }
      }
      ... on Film {
        episodeID
      }
    }
  }
}
`

func BenchmarkLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexer.Tokenize([]byte(gql))
	}
}
