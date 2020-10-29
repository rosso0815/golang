package config01

import l "log"

type Person struct {
	Firstname string
	Lastname  string
}

type PersonInfo interface {
	GetInfo()
}

func init() {
	l.Println("@@@ config01 init")
}

func Hello() {
	l.Println("@@@ config01 hello")
}

func MakeInfo(s string, p PersonInfo) {
	l.Println("@@@ config01 MakeInfo")
	p.GetInfo()
}
