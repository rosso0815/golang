package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	const src = `
// This is scanned code.
if a > 10 {
	someParsable = text
}
1:"1222":1
2:"11":333
`
	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

}
