package main

import "github.com/yuin/gopher-lua"

func main() {

	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile("hello.lua"); err != nil {
		panic(err)
	}

	if err := L.DoString(`print("hello andreas")`); err != nil {
		panic(err)
	}
}
