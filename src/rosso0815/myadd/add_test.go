package myadd

import (
	"log"
	"testing"
)

func Test_Add_1(t *testing.T) {
	log.Println("@@@ Test Add")
	i := Add(1, 2)
	if i != 3 {
		t.Errorf("expected '%d' but got '%d'", 3, 3)
	}
}
