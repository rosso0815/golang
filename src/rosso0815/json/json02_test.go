package main

import (
	"encoding/json"
	"fmt"

	"testing"
)

type Cmds struct {
	searchText string                 `json:"search"`
	B          int                    `json:"b"`
	X          map[string]interface{} `json:"-"` // Rest of the fields should go here.
}

func TestJson02(t *testing.T) {
	s := `{"search":1, "b":2, "x":1, "y":1}`

	j := Cmds{}
	if err := json.Unmarshal([]byte(s), &j.X); err != nil {
		panic(err)
	}

	//	if n, ok := j.X["a"].(float64); ok {
	//		f.A = int(n)
	//	}

	if n, ok := j.X["b"].(float64); ok {
		j.B = int(n)
	}

	//delete(f.X, "a")
	delete(j.X, "b")

	fmt.Printf("%+v", j)
}
