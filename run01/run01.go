package main

//import "os"
import "fmt"

//define a Rectangle struct that has a length and a width
type Rectangle struct {
	length, width int
}

//write a function Area that can apply to a Rectangle type
func (r Rectangle) Area() int {
	return r.length * r.width
}

type Car interface {
	GetBrand() string
}

// based on car we have Volkswagen
type Volkswagen struct {
	brand string
}

func (c Volkswagen) GetBrand() string {
	return "VW Polo"
}

func main() {
	fmt.Println("hello world")
	fmt.Println("")
	r := Rectangle{length: 5, width: 3} //define a new Rectangle instance with values for its properties
	fmt.Println("Rectangle details are: ", r)
	fmt.Println("Rectangle's area is: ", r.Area())

	fmt.Println("start car")
	polo := Volkswagen{brand: "VW Polo"}
	var i Car
	i = polo
	fmt.Println("car brand=", i.GetBrand())

	//argsWithProg := os.Args
	//args := len(os.Args)
	//fmt.Println("len=", args)
	//argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.
	//arg := os.Args[3]

	//fmt.Println(argsWithProg)
	//fmt.Println(argsWithoutProg)
	//fmt.Println(arg)

	//os.Exit(3)
}
