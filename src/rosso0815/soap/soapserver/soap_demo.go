package soapserver

import (
	"encoding/xml"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://codenotfound.com/services/helloworld"

// NewHelloWorld_PortType creates an initializes a HelloWorld_PortType.
func NewHelloWorld_PortType(cli *soap.Client) HelloWorld_PortType {
	return &helloWorld_PortType{cli}
}

// HelloWorld_PortType was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type HelloWorld_PortType interface {
	// Echo was auto-generated from WSDL.
	Echo(α *Echo) (β *Echo, err error)

	// SayHello was auto-generated from WSDL.
	SayHello(α *Person) (β *Greeting, err error)
}

// NewOperation was auto-generated from WSDL.
type NewOperation struct {
	In string `xml:"in,omitempty" json:"in,omitempty" yaml:"in,omitempty"`
}

// NewOperationResponse was auto-generated from WSDL.
type NewOperationResponse struct {
	Out string `xml:"out,omitempty" json:"out,omitempty" yaml:"out,omitempty"`
}

// NewOperationResponse1 was auto-generated from WSDL.
type NewOperationResponse1 struct {
	Out string `xml:"out,omitempty" json:"out,omitempty" yaml:"out,omitempty"`
}

// Echo was auto-generated from WSDL.
type Echo struct {
	XMLName xml.Name `xml:"http://codenotfound.com/services/helloworld echo" json:"-" yaml:"-"`
	In      string   `xml:"in,omitempty" json:"in,omitempty" yaml:"in,omitempty"`
}

// GetEcho was auto-generated from WSDL.
type GetEcho struct {
	In string `xml:"in,omitempty" json:"in,omitempty" yaml:"in,omitempty"`
}

// GetEchoResponse was auto-generated from WSDL.
type GetEchoResponse struct {
	Out string `xml:"out,omitempty" json:"out,omitempty" yaml:"out,omitempty"`
}

// Greeting was auto-generated from WSDL.
type Greeting struct {
	Greeting string `xml:"greeting,omitempty" json:"greeting,omitempty" yaml:"greeting,omitempty"`
}

// Person was auto-generated from WSDL.
type Person struct {
	XMLName   xml.Name `xml:"http://codenotfound.com/services/helloworld person" json:"-" yaml:"-"`
	FirstName string   `xml:"firstName,omitempty" json:"firstName,omitempty" yaml:"firstName,omitempty"`
	LastName  string   `xml:"lastName,omitempty" json:"lastName,omitempty" yaml:"lastName,omitempty"`
}

// helloWorld_PortType implements the HelloWorld_PortType interface.
type helloWorld_PortType struct {
	cli *soap.Client
}

// Echo was auto-generated from WSDL.
func (p *helloWorld_PortType) Echo(α *Echo) (β *Echo, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M Echo `xml:"echo"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SayHello was auto-generated from WSDL.
func (p *helloWorld_PortType) SayHello(α *Person) (β *Greeting, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M Greeting `xml:"greeting"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}
