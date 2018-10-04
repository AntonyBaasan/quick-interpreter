package lexer

import (
	"../token"
	"fmt"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct{
		expecteddType token.TokenType
		expectedLiteral string
	}{
		{token.	ASSIGN, "="},
	}

	l := New(input)

	fmt.Println(tests)
	fmt.Println(l.input)
}
