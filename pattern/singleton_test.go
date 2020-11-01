package pattern

import (
	"fmt"
	"testing"
)

func TestSingleton(t *testing.T) {

	fmt.Println("@@@ TestSingleton")

	s1 := New()
	s1["this"] = "that"
	fmt.Println("This is ", s1["this"])

	s2 := New()
	fmt.Println("This is ", s2["this"])
}
