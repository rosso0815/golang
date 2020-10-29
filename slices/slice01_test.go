package slices

import (
	"fmt"
	"log"
	"testing"
)

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func TestSlice01(t *testing.T) {
	log.Println("@@@ testSlice01")

	var v []string

	printSlice(v)

	v = append(v, "hallo1")

	printSlice(v)

	v = append(v, "hallo2")
	v = append(v, "hallo3")
	printSlice(v)

}
