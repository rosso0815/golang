package main

import "fmt"

type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF

	literal_beg
	IDENT  // main
	INT    // 12345
	STRING // “abc”
	literal_end

	keyword_beg
	BREAK
	FUNC
	SELECT
	keyword_end
)

func (tok Token) IsLiteral() bool {
	return literal_beg < tok && tok < literal_end
}

func main() {
	fmt.Println("vim-go")

	var token Token

	token = INT

	fmt.Println("token=", token)

}
