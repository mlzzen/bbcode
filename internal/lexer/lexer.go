package lexer

import (
	"unicode/utf8"
)

type T uint8

const endOfFile = -1

const (
	TEof T = iota
	TIdent
	TWhiteSpace
	TEqual
	TOpenBraceket
	TCloseBracket
	TBadToken
	TBold
	TItalic
	TUnderline
	TStrikethrough
	TFontSize
	TFontColor
	TCenterText
	TLeftAlignText
	TRightAlignText
	TQuote
	TSpoiler
	TLink
	TImage
	TList
	TListItem
	TCode
	TPreformatted
	TTables
	TTableRows
	TTableContentCells
	TYoutubeVideos
	TAttribute
)

var tokenToString = []string{
	"<eof-token>",
	"<ident-token>",
	"<whitespace-token>",
	"<=-token>",
	"<[-token>",
	"<]-token>",
	"<bad-token>",
}

var Keywords = map[string]T{
	"b":       TBold,
	"i":       TItalic,
	"u":       TUnderline,
	"s":       TStrikethrough,
	"size":    TFontSize,
	"color":   TFontColor,
	"center":  TCenterText,
	"left":    TLeftAlignText,
	"right":   TRightAlignText,
	"quote":   TQuote,
	"spoiler": TSpoiler,
	"url":     TLink,
	"img":     TImage,
	"ul":      TList,
	"ol":      TList,
	"list":    TList,
	"li":      TListItem,
	"code":    TCode,
	"pre":     TPreformatted,
	"table":   TTables,
	"tr":      TTableRows,
	"th":      TTableContentCells,
	"td":      TTableContentCells,
	"youtube": TYoutubeVideos,
}

var Attributes = map[string]T{
	"width":  TAttribute,
	"height": TAttribute,
	"*":      TAttribute,
}

func (t T) String() string {
	return tokenToString[t]
}

func isWhiteSpace(s rune) bool {
	return s == ' ' || s == '\t'
}

type Loc struct {
	Start int32
	Len   int32
}

func (loc Loc) End() int32 {
	return loc.Start + loc.Len
}

type Token struct {
	Kind T
	Loc  Loc
}

func (token Token) DecodedText(s string) string {
	raw := s[token.Loc.Start:token.Loc.End()]
	return raw
}

type lexer struct {
	source string
	pos    int
	cp     rune
	token  Token
}

type TokenizeResult struct {
	Tokens []Token
}

func Tokenizer(input string) TokenizeResult {
	l := &lexer{
		source: input,
	}
	var tokens []Token
	l.step()
	if l.cp == '\uFEFF' {
		l.step()
	}
	l.next()
	for l.token.Kind != TEof {
		tokens = append(tokens, l.token)
		l.next()
	}

	return TokenizeResult{
		Tokens: tokens,
	}
}

func (lexer *lexer) step() {
	cp, width := utf8.DecodeRuneInString(lexer.source[lexer.pos:])
	if width == 0 {
		cp = -1
	}

	lexer.cp = cp
	lexer.token.Loc.Len = int32(lexer.pos) - lexer.token.Loc.Start
	lexer.pos += width
}

func (lexer *lexer) next() {
	for {
		lexer.token = Token{Loc: Loc{Start: lexer.token.Loc.End()}}
		switch lexer.cp {
		case endOfFile:
			lexer.token.Kind = TEof
		case ' ', '\t':
			lexer.step()
			for {
				if !isWhiteSpace(lexer.cp) {
					break
				}
				lexer.step()
			}
			lexer.token.Kind = TWhiteSpace
		case '\r', '\n', '\f':
			if lexer.cp == '\r' {
				lexer.step()
			}
			lexer.step()
			continue
		case '[':
			lexer.step()
			lexer.token.Kind = TOpenBraceket
		case ']':
			lexer.step()
			lexer.token.Kind = TCloseBracket
		case '=':
			lexer.step()
			lexer.token.Kind = TEqual
		default:
			lexer.token.Kind = lexer.consumeIdent()
		}
		return
	}
}

func (lexer *lexer) consumeIdent() T {
	for {
		if lexer.cp == endOfFile || lexer.cp == '[' {
			break
		}
		lexer.step()
	}
	return TIdent
}
