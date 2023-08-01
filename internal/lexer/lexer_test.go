package lexer

import "testing"

func lexToken(input string) (T, string) {
	tokens := Tokenizer(input).Tokens
	if len(tokens) > 0 {
		return tokens[0].Kind, tokens[0].DecodedText(input)
	}
	return TEof, ""
}

func TestTokens(t *testing.T) {
	expected := []struct {
		content string
		text    string
		token   T
	}{
		{"ident", "<ident-token>", TIdent},
		{" ", "<whitespace-token>", TWhiteSpace},
		{"[", "<[-token>", TOpenBraceket},
		{"]", "<]-token>", TCloseBracket},
		{"", "<eof-token>", TEof},
	}
	for _, expr := range expected {
		t.Run(expr.content, func(t *testing.T) {
			tok, _ := lexToken(expr.content)
			if tok != expr.token {
				t.Fatalf("%s != %s", tok, expr.token)
			}
			if tok.String() != expr.text {
				t.Fatalf("%s != %s", tok.String(), expr.text)
			}
		})
	}
}
