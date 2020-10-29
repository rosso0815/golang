package factory

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestVertex(t *testing.T) {
	fmt.Println("### TestVertex start")
	x := NewVertex(0, 0)
	x.Add(1, 1)
	fmt.Println("x=", x)
	fmt.Println("### TestVertex done")
}
func TestSum(t *testing.T) {
	fmt.Println("### TestSum start")
	fmt.Println("My favorite number is", rand.Intn(100))

	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	fmt.Println("### TestSum done")
}
