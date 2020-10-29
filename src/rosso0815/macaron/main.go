package main

import (
	"log"
	"net/http"

	"gopkg.in/macaron.v1"
)

func main() {
	//	m := macaron.Classic()
	//	m.Get("/", func() string {
	//		return "Hello world!"
	//	})
	//	m.Run()
	port := 4000
	m := macaron.Classic()
	m.Get("/", myHandler)

	log.Println("Server is running on port ", port)
	log.Println("Hi Andreas")

	log.Println(http.ListenAndServe("0.0.0.0:4000", m))

}

func myHandler(ctx *macaron.Context) string {
	return "the request path is: " + ctx.Req.RequestURI
}
