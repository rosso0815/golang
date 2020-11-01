package pattern

import (
	"fmt"
	"testing"
)

func TestSingleton2(t *testing.T) {

	for i := 0; i < 10; i++ {
		go getInstance()
	}
	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
