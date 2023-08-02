package lexer

import (
	"testing"

	"github.com/nonzzz/bbcode/internal/test"
)

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
		{"[", "<[-token>", TOpenBracket},
		{"]", "<]-token>", TCloseBracket},
		{"", "<eof-token>", TEof},
	}
	for _, expr := range expected {
		t.Run(expr.content, func(t *testing.T) {
			tok, _ := lexToken(expr.content)
			test.AssertEqual(t, tok, expr.token)
			test.AssertEqual(t, tok.String(), expr.text)
		})
	}
}

// func TestBlold(t *testing.T) {
// 	input := "I would like to [b]emphasize[/b] this"
// 	expected := []struct {
// 		content string
// 		text    string
// 		token   T
// 	}{
// 		{"I would like to ", "<ident-token>", TIdent},
// 		{"[b]", "<bold-token>", TBold},
// 		{"emphasize", "<ident-token>", TIdent},
// 		{"[/b]", "<bold-token>", TBold},
// 		{" ", "<whitespace-token>", TWhiteSpace},
// 		{"this", "<ident-token>", TIdent},
// 	}
// 	tokens := Tokenizer(input).Tokens
// 	for i, expr := range expected {
// 		t.Run(expr.content, func(t *testing.T) {
// 			tok := tokens[i].Kind
// 			test.AssertEqual(t, tok, expr.token)
// 			test.AssertEqual(t, tok.String(), expr.text)
// 		})
// 	}
// }

func TestKeywords(t *testing.T) {
	lexKeyWord := func(input string) string {
		t.Helper()
		tok, raw := lexToken(input)
		test.AssertEqual(t, input, raw)
		return tok.String()
	}
	test.AssertEqual(t, lexKeyWord("[b]"), "<bold-token>")
	test.AssertEqual(t, lexKeyWord("[/b]"), "<bold-token>")
	test.AssertEqual(t, lexKeyWord("[i]"), "<italic-token>")
	test.AssertEqual(t, lexKeyWord("[/i]"), "<italic-token>")
	test.AssertEqual(t, lexKeyWord("[u]"), "<underline-token>")
	test.AssertEqual(t, lexKeyWord("[/u]"), "<underline-token>")
	test.AssertEqual(t, lexKeyWord("[s]"), "<strikethrough-token>")
	test.AssertEqual(t, lexKeyWord("[/s]"), "<strikethrough-token>")
	test.AssertEqual(t, lexKeyWord("[size]"), "<font-size-token>")
	test.AssertEqual(t, lexKeyWord("[/size]"), "<font-size-token>")
	test.AssertEqual(t, lexKeyWord("[color]"), "<font-color-token>")
	test.AssertEqual(t, lexKeyWord("[/color]"), "<font-color-token>")
	test.AssertEqual(t, lexKeyWord("[center]"), "<center-text-token>")
	test.AssertEqual(t, lexKeyWord("[/center]"), "<center-text-token>")
	test.AssertEqual(t, lexKeyWord("[left]"), "<left-aligin-text-token>")
	test.AssertEqual(t, lexKeyWord("[/left]"), "<left-aligin-text-token>")
	test.AssertEqual(t, lexKeyWord("[right]"), "<right-align-text-token>")
	test.AssertEqual(t, lexKeyWord("[/right]"), "<right-align-text-token>")
	test.AssertEqual(t, lexKeyWord("[quote]"), "<quote-token>")
	test.AssertEqual(t, lexKeyWord("[/quote]"), "<quote-token>")
	test.AssertEqual(t, lexKeyWord("[spoiler]"), "<spoiler-token>")
	test.AssertEqual(t, lexKeyWord("[/spoiler]"), "<spoiler-token>")
	test.AssertEqual(t, lexKeyWord("[url]"), "<url-token>")
	test.AssertEqual(t, lexKeyWord("[/url]"), "<url-token>")
	test.AssertEqual(t, lexKeyWord("[img]"), "<img-token>")
	test.AssertEqual(t, lexKeyWord("[/img]"), "<img-token>")
	test.AssertEqual(t, lexKeyWord("[ul]"), "<list-token>")
	test.AssertEqual(t, lexKeyWord("[/ul]"), "<list-token>")
	test.AssertEqual(t, lexKeyWord("[ol]"), "<list-token>")
	test.AssertEqual(t, lexKeyWord("[/ol]"), "<list-token>")
	test.AssertEqual(t, lexKeyWord("[list]"), "<list-token>")
	test.AssertEqual(t, lexKeyWord("[/list]"), "<list-token>")
	test.AssertEqual(t, lexKeyWord("[li]"), "<list-item-token>")
	test.AssertEqual(t, lexKeyWord("[/li]"), "<list-item-token>")
	test.AssertEqual(t, lexKeyWord("[code]"), "<code-token>")
	test.AssertEqual(t, lexKeyWord("[/code]"), "<code-token>")
	test.AssertEqual(t, lexKeyWord("[pre]"), "<pre-token>")
	test.AssertEqual(t, lexKeyWord("[/pre]"), "<pre-token>")
	test.AssertEqual(t, lexKeyWord("[table]"), "<table-token>")
	test.AssertEqual(t, lexKeyWord("[/table]"), "<table-token>")
	test.AssertEqual(t, lexKeyWord("[tr]"), "<table-row-token>")
	test.AssertEqual(t, lexKeyWord("[/tr]"), "<table-row-token>")
	test.AssertEqual(t, lexKeyWord("[th]"), "<table-cell-token>")
	test.AssertEqual(t, lexKeyWord("[/th]"), "<table-cell-token>")
	test.AssertEqual(t, lexKeyWord("[td]"), "<table-cell-token>")
	test.AssertEqual(t, lexKeyWord("[/td]"), "<table-cell-token>")
	test.AssertEqual(t, lexKeyWord("[youtube]"), "<youtube-videos-token>")
	test.AssertEqual(t, lexKeyWord("[/youtube]"), "<youtube-videos-token>")
}

func TestKeywordWithUnnecessaryWhiteSpace(t *testing.T) {
	lexContenet := func(input string) string {
		t.Helper()
		toks := Tokenizer(input).Tokens
		test.AssertEqual(t, len(toks), 1)
		test.AssertEqual(t, toks[0].DecodedText(input), input)
		return toks[0].Kind.String()
	}

	test.AssertEqual(t, lexContenet("[/b align justify-all]"), "<bold-token>")
}

func TestKeyWordWithAttribues(t *testing.T) {
	input := "[b align justify-all] hello world! [/b]"
	toks := Tokenizer(input).Tokens
	test.AssertEqual(t, len(toks), 6)
}

func TestBOM(t *testing.T) {
	tok, _ := lexToken("\uFEFF")
	test.AssertEqual(t, tok, TEof)
}
