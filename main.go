package main

import (
	"fmt"
	"os"

	"github.com/nonzzz/bbcode/internal/lexer"
)

func main() {

	buf, _ := os.ReadFile("./case/content.txt")
	s := string(buf)
	toks := lexer.Tokenizer(s).Tokens
	for _, tok := range toks {
		fmt.Println(tok.DecodedText(s))
	}

}
