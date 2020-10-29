package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("vim-go")

	r := strings.NewReader("Hallo Andreas")

	if _, err := io.Copy(os.Stderr, r); err != nil {
		log.Fatal(err)
	}
}
