package playinterface

import (
	"log"
	"testing"
)

// from https://www.integralist.co.uk/posts/go-interfaces/

type Base interface {
	ShowInfo()
	//Bar(s string) (string, error)
}

type base1 struct{ base string }

func (b *base1) ShowInfo() {
	log.Println("@@@ showInfo of base1 , base =", b.base)
}

type base2 struct{ base string }

func (b *base2) ShowInfo() {
	log.Println("@@@ showInfo of base2 , base =", b.base)
}

func doStuffWithBase(base Base) {
	log.Println("doStuffWith")
	base.ShowInfo()
}

func Test_Base(t *testing.T) {
	log.Println("@@@ Test_Car")

	b1 := base1{base: "base_1"}
	doStuffWithBase(&b1)

	b2 := base2{base: "base_2"}
	doStuffWithBase(&b2)
}
