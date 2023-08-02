package parse

import (
	"github.com/nonzzz/bbcode/internal/lexer"
)

type parser struct {
	input  string
	start  int
	end    int
	tokens []lexer.Token
}

func Parse(input string) {
	result := lexer.Tokenizer(input)
	p := &parser{
		input:  input,
		tokens: result.Tokens,
	}
	p.parse()
	//
}

func (p *parser) parse() {}

func (p *parser) eat(kind lexer.T) bool {
	if p.peek(kind) {
		p.advance()
		return true
	}
	return false
}

func (p *parser) advance() {}

func (p *parser) peek(kind lexer.T) bool {
	return kind == p.current().Kind
}

func (p *parser) at(pos int) lexer.Token {
	if pos < p.end {
		return p.tokens[pos]
	}
	if p.end < len(p.tokens) {
		return lexer.Token{
			Kind: lexer.TEof,
			Loc:  p.tokens[p.end].Loc,
		}
	}
	return lexer.Token{
		Kind: lexer.TEof,
		Loc:  lexer.Loc{Start: int32(len(p.input))},
	}
}

func (p *parser) current() lexer.Token {
	return p.at(p.start)
}

func (p *parser) decode() string {
	return p.current().DecodedText(p.input)
}
