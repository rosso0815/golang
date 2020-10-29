package main

import (
	"fmt"

	c "concurrency/myconcurrency"
)

const i = 1

func main() {

	fmt.Println("Hallo Rosso")

	t := c.Sum(i, i)

	fmt.Printf("%v + %v = %v ", i, i, t)
}
