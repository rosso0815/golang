package main

import "fmt"

// Sum add x y
func Sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("sum=", Sum(5, 5))
}
