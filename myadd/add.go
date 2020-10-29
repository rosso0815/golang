package myadd

import "log"

// Add x ,y return x+y
func Add(x int, y int) int {
	log.Println("@@@ myadd -> Add")
	return x + y
}
