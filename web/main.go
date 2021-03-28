package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	Port = ":8080"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "The time is now " + time.Now().String()
	fmt.Fprintln(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func main() {
	log.Println("Hello world!")
	osPort := os.Getenv("PORT")
	if len(osPort) < 1 {
		log.Println("Port empty")

	} else {
		Port = osPort
	}
	log.Println("Port = ", Port)
	http.HandleFunc("/static", serveStatic)
	http.HandleFunc("/", serveDynamic)
	http.ListenAndServe(Port, nil)
}
