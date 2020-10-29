package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type message struct {
	Name string
	Body string
	Time int64
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	m := message{"Alice 1", "Hello User Andreas", 1294706395881547000}
	b, _ := json.Marshal(m)
	w.Write([]byte(b))
}
func main() {
	log.Println("start")
	http.HandleFunc("/", productsHandler)
	http.ListenAndServe(":8080", nil)
	log.Println("done")
}
