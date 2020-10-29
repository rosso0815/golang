package myinterface

import (
	"log"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	log.Println("M")
}
