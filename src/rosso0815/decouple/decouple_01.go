package config

type Config interface {
    Foo() string
}

package foo

type Foo struct{}

func (f *Foo) Foo() string {
    return "foo"
}

package boo

type Boo struct{}

func (b *Boo) Foo() string {
    return "boo" 
}

package main

func foo(cfg config.Config) string{
    return cfg.Foo()
}

func main() {
    // here you inject an instance of Foo into foo(Config)

    log.Print(foo(&foo.Foo{}))

    // here you inject an instance of Boo into foo(Config)
    log.Print(foo(&boo.Boo{})
}
