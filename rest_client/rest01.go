package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Info struct {
	Name    string
	Vorname string
}

func main() {
	log.Print("start")

	// url
	url := "http://localhost/pfistera/cgi/rest/sayHi.cgi"

	//	get data
	var netClient = &http.Client{
		Timeout: time.Second * 60,
	}

	response, err := netClient.Get(url)
	if err != nil {
		log.Print("response failed error = ", err)
	}
	log.Print("response=", response)
	log.Print("responseBody=", response.Body)

	// convert to JSON
	var data Info
	decoder := json.NewDecoder(response.Body)

	err = decoder.Decode(&data)
	if err != nil {
		log.Print("json-error")
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
		//	switch v := err.(type) {
		//	case *json.SyntaxError:
		//		fmt.Println(string(body[v.Offset-40 : v.Offset]))
		//	}
	}

	//json.Unmarshal(response, &data)
	log.Print("Results:", data.Name, " ", data.Vorname)

	log.Print("ende")
	os.Exit(0)
}
