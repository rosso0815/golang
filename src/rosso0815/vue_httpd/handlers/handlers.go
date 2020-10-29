package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Hello for base testing
func Hello() {
	log.Println("Hello")
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//vars["title"] // the book title slug
	//vars["page"]  // the page

	log.Println("mux.vars=", vars)
}
