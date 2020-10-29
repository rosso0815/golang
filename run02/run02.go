package main

import (
	"fmt"
)

//------------------------------------------------------------------------------
type User struct {
	FirstName, LastName string
}

func (u *User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
func (u *User) AllText() string {
	return fmt.Sprintf("AllUser => %s, %s", u.FirstName, u.LastName)
}

//------------------------------------------------------------------------------
type Worker struct {
	FirstName, LastName, Job string
}

func (u *Worker) Name() string {
	return fmt.Sprintf("%s %s %s", u.FirstName, u.LastName, u.Job)
}

func (u *Worker) AllText() string {
	return fmt.Sprintf("AllWorker => %s, %s, %s", u.FirstName, u.LastName, u.Job)
}

//------------------------------------------------------------------------------
type Namer interface {
	Name() string
	AllText() string
}

//------------------------------------------------------------------------------
func Greet(n Namer) string {
	return fmt.Sprintf("Dear %s %s", n.Name(), n.AllText())
}

func main() {
	u := &User{"Matt", "Aimonetti"}
	fmt.Println(Greet(u))

	w := &Worker{"Andreas", "Pfister", "Programmer"}
	fmt.Println(Greet(w))
}
