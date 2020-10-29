package main

import (
	"encoding/json"
	"log"
	"testing"
)

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}

func TestJson01(t *testing.T) {

	log.Print("100 start")

	// Create an instance of the Box struct.
	box := Box{
		Width:  10,
		Height: 20,
		Color:  "blue",
		Open:   false,
		//Info:   "text",
	}

	// Create JSON from the instance data.
	// ... Ignore errors.
	b, _ := json.Marshal(box)

	// Convert bytes to string.
	s := string(b)
	log.Print(s)

	// and versa

	//b == []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)

	type Message struct {
		Name string
		Body string
		Time int64
	}

	log.Print("1 unmarshall")

	c := []byte(`[{"Name":"Bob","Food":"Pickle"},{"Name":"Beppo"}]`)
	//c := []byte(`{"Name":"Bob","Body":"Pickle"}`)
	var m []Message
	err := json.Unmarshal(c, &m)
	if err != nil {
		log.Fatal("json.unmarshall ", err)
	} else {
		log.Print("json.unmarshall", m)
	}

	log.Print("done")
}
