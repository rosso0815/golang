package main

import (
	"fmt"

	//"/path/to/generated/example"

	"github.com/fiorix/wsdl2go/soap"

	"pfistera/soap/soapserver"
)

func main() {

	fmt.Println("start soap")

	cli := soap.Client{
		URL:       "http://localhost:9090/codenotfound/ws/helloworld?wsdl",
		Namespace: soapserver.Namespace,
	}

	soapService := soapserver.NewHelloWorld_PortType(&cli)

	//echoReply, err := soapService.Echo(&soapserver.Echo{In: "Hi"})
	echoReply, err := soapService.Echo(&soapserver.Echo{In: "2"})

	if err != nil {
		fmt.Println("ERROR", err)
	} else {
		fmt.Println("echoReplay", echoReply)
	}

}
