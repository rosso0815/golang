package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"pfistera/vue_httpd/handlers"
)

func main() {

	log.Println("start")

	handlers.Hello()

	r := mux.NewRouter()

	r.HandleFunc("/api/products", handlers.ProductHandler)
	http.Handle("/api", r)

	http.Handle("/", http.FileServer(http.Dir("./src")))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	log.Println("finish")
}
