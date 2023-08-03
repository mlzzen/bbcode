package parse

import (
	"fmt"
	"strings"

	"github.com/nonzzz/bbcode/internal/lexer"
)

type ElementNode struct {
	Type      lexer.T
	Text      string
	Attribues map[string]interface{}
	Children  []ElementNode
}

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
		end:    len(result.Tokens),
	}
	elements := p.parse()
	fmt.Println(elements, len(elements))
}

func (p *parser) parse() []ElementNode {
	var elements []ElementNode = make([]ElementNode, 0)
loop:
	for {
		switch p.current().Kind {
		case lexer.TEof:
			break loop
		case lexer.TWhiteSpace:
			p.advance()
			continue
		case lexer.TOpenBracket:
			elements = append(elements, p.makeNode(lexer.TOpenBracket))
			p.eat(lexer.TOpenBracket)
			continue
		case lexer.TIdent:
			elements = append(elements, p.makeNode(lexer.TIdent))
			p.eat(lexer.TIdent)
			continue
		case lexer.TBold, lexer.TItalic, lexer.TUnderline, lexer.TStrikethrough,
			lexer.TFontSize, lexer.TFontColor, lexer.TCenterText, lexer.TLeftAlignText,
			lexer.TRightAlignText, lexer.TQuote, lexer.TSpoiler,
			lexer.TLink, lexer.TImage, lexer.TList, lexer.TListItem,
			lexer.TCode, lexer.TPreformatted, lexer.TTables, lexer.TTableRows,
			lexer.TTableContentCells, lexer.TYoutubeVideos, lexer.TCollapse, lexer.TStyle:
			element := p.parseKnowElement()
			if element.Type != lexer.TBadToken {
				elements = append(elements, element)
			}
			continue
		default:
			break loop
		}
	}
	return elements
}

func (p *parser) eat(kind lexer.T) bool {
	if p.peek(kind) {
		p.advance()
		return true
	}
	return false
}

func (p *parser) advance() {
	if p.start < p.end {
		p.start++
	}
}

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

func (p *parser) parseKnowElement() ElementNode {
	expected := p.current().Kind
	bad := false
	element := p.makeNode(expected)
	p.advance()
knowElement:
	for {
		switch p.current().Kind {
		case expected:
			p.eat(expected)
			break knowElement
		case lexer.TEof:
			element.Type = lexer.TIdent
			break knowElement
		case lexer.TBadToken:
			bad = true
			p.advance()
			break knowElement
		case lexer.TAttribute:
			raw := p.decode()
			if expected == lexer.TFontSize || expected == lexer.TFontColor {
				if strings.HasPrefix(raw, "=") {
					element.Attribues[lexer.KnowElementAttrToString[expected]] = raw[1:]
				}
			}
			p.eat(lexer.TAttribute)
			continue
		case lexer.TIdent:
			element.Children = append(element.Children, p.makeNode(lexer.TIdent))
			p.eat(lexer.TIdent)
			continue
		default:
			element.Children = append(element.Children, p.parseKnowElement())
		}

	}
	if bad {
		return p.makeNode(lexer.TBadToken)
	}
	return element
}

func (p *parser) makeNode(kind lexer.T) ElementNode {
	return ElementNode{
		Type:      kind,
		Text:      p.decode(),
		Attribues: make(map[string]interface{}),
	}
}
