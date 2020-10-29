package main

import l "log"
import c "pfistera/config01"

type InfoAll interface {
	ShowInfo()
	ShowSalary()
}

type Person2 struct {
	c.Person
	//Firstname string
	//Lastname  string

	Gender string
}

// used for interface
func (p *Person2) ShowInfo() {
	l.Printf("interface->ShowInfo => Person firstname = %s ; lastname = %s\n", p.Firstname, p.Lastname)
}

func (p *Person2) ShowSalary() {
	l.Printf("ShowPersonSalary 100Baht\n")
}

// used for person internally
func (p *Person2) ChangeFirstname(t string) {
	l.Printf("ChangeFirstname => firstname = %s, before %s\n", t, p.Firstname)
	p.Firstname = t
}

func (p *Person2) GetInfo() {
	l.Println("Person2 - GetInfo Firstname = ", p.Firstname)
}

type Boss struct {
	Person2
	Title string
}

func (p Boss) ShowInfo() {
	l.Printf("ShowInfoBoss => firstname = %s ; lastname = %s\n", p.Firstname, p.Lastname)
}

func (p Boss) ShowSalary() {
	l.Printf("ShowBossSalary 1000Baht\n")
}

func main() {

	c.Hello()

	l.Println("start")

	var p1 Person2
	var info InfoAll
	info = &p1

	p1.Firstname = "Andreas"
	p1.Lastname = "Pfister"

	info.ShowInfo()
	info.ShowSalary()

	//p1 := &Person{Firstname: "Andreas", Lastname: "Pfister"}
	//p1.ShowPersonInfo()
	p1.ChangeFirstname("Peter")
	p1.ShowInfo()
	p1.Firstname = "Heinz"
	p1.ShowInfo()

	var b1 Boss
	b1.Person.Firstname = "Peter"
	b1.Person.Lastname = "Meier"
	b1.Title = "Master"
	info = &b1
	info.ShowInfo()
	info.ShowSalary()

	c.MakeInfo("info", &p1)

	l.Println("ende")
}
