package main

//go:generate gopherjs build

// This is an experiment to see if gopherjs can reasonably generate js code from go source
// so that we can have a single-source solution for keys and addresses.
// Use "go generate" to build this.

import (
	"crypto/sha256"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Module.Get("exports").Set("hashit", hashit)
}

func hashit(s string) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}
