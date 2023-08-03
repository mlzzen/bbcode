package lexer

var tokenToString = map[T]string{
	TEof:               "<eof-token>",
	TIdent:             "<ident-token>",
	TWhiteSpace:        "<whitespace-token>",
	TEqual:             "<=-token>",
	TOpenBracket:       "<[-token>",
	TCloseBracket:      "<]-token>",
	TBadToken:          "<bad-token>",
	TBold:              "<bold-token>",
	TItalic:            "<italic-token>",
	TUnderline:         "<underline-token>",
	TStrikethrough:     "<strikethrough-token>",
	TFontSize:          "<font-size-token>",
	TFontColor:         "<font-color-token>",
	TCenterText:        "<center-text-token>",
	TLeftAlignText:     "<left-aligin-text-token>",
	TRightAlignText:    "<right-align-text-token>",
	TQuote:             "<quote-token>",
	TSpoiler:           "<spoiler-token>",
	TLink:              "<url-token>",
	TImage:             "<img-token>",
	TList:              "<list-token>",
	TListItem:          "<list-item-token>",
	TCode:              "<code-token>",
	TPreformatted:      "<pre-token>",
	TTables:            "<table-token>",
	TTableRows:         "<table-row-token>",
	TTableContentCells: "<table-cell-token>",
	TYoutubeVideos:     "<youtube-videos-token>",
	TAttribute:         "<attribute-token>",
	TCollapse:          "<collapse-token>",
	TStyle:             "<style-token>",
}

var KnowElementAttrToString = map[T]string{
	TFontSize:       "font-size",
	TFontColor:      "color",
	TCenterText:     "text-align",
	TLeftAlignText:  "text-align",
	TRightAlignText: "text-align",
}
