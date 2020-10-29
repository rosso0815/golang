// Code generated by go-swagger; DO NOT EDIT.

package web_mvc_links_handler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new web mvc links handler API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for web mvc links handler API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
LinksUsingGET links
*/
func (a *Client) LinksUsingGET(params *LinksUsingGETParams) (*LinksUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLinksUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "linksUsingGET",
		Method:             "GET",
		PathPattern:        "/actuator",
		ProducesMediaTypes: []string{"application/json", "application/vnd.spring-boot.actuator.v2+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LinksUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LinksUsingGETOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
