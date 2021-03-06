// Code generated by go-swagger; DO NOT EDIT.

package basic_error_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "pfistera/swagger_client/models"
)

// ErrorHTMLUsingOPTIONSReader is a Reader for the ErrorHTMLUsingOPTIONS structure.
type ErrorHTMLUsingOPTIONSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ErrorHTMLUsingOPTIONSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewErrorHTMLUsingOPTIONSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewErrorHTMLUsingOPTIONSNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewErrorHTMLUsingOPTIONSUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewErrorHTMLUsingOPTIONSForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewErrorHTMLUsingOPTIONSOK creates a ErrorHTMLUsingOPTIONSOK with default headers values
func NewErrorHTMLUsingOPTIONSOK() *ErrorHTMLUsingOPTIONSOK {
	return &ErrorHTMLUsingOPTIONSOK{}
}

/*ErrorHTMLUsingOPTIONSOK handles this case with default header values.

OK
*/
type ErrorHTMLUsingOPTIONSOK struct {
	Payload *models.ModelAndView
}

func (o *ErrorHTMLUsingOPTIONSOK) Error() string {
	return fmt.Sprintf("[OPTIONS /error][%d] errorHtmlUsingOPTIONSOK  %+v", 200, o.Payload)
}

func (o *ErrorHTMLUsingOPTIONSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelAndView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewErrorHTMLUsingOPTIONSNoContent creates a ErrorHTMLUsingOPTIONSNoContent with default headers values
func NewErrorHTMLUsingOPTIONSNoContent() *ErrorHTMLUsingOPTIONSNoContent {
	return &ErrorHTMLUsingOPTIONSNoContent{}
}

/*ErrorHTMLUsingOPTIONSNoContent handles this case with default header values.

No Content
*/
type ErrorHTMLUsingOPTIONSNoContent struct {
}

func (o *ErrorHTMLUsingOPTIONSNoContent) Error() string {
	return fmt.Sprintf("[OPTIONS /error][%d] errorHtmlUsingOPTIONSNoContent ", 204)
}

func (o *ErrorHTMLUsingOPTIONSNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorHTMLUsingOPTIONSUnauthorized creates a ErrorHTMLUsingOPTIONSUnauthorized with default headers values
func NewErrorHTMLUsingOPTIONSUnauthorized() *ErrorHTMLUsingOPTIONSUnauthorized {
	return &ErrorHTMLUsingOPTIONSUnauthorized{}
}

/*ErrorHTMLUsingOPTIONSUnauthorized handles this case with default header values.

Unauthorized
*/
type ErrorHTMLUsingOPTIONSUnauthorized struct {
}

func (o *ErrorHTMLUsingOPTIONSUnauthorized) Error() string {
	return fmt.Sprintf("[OPTIONS /error][%d] errorHtmlUsingOPTIONSUnauthorized ", 401)
}

func (o *ErrorHTMLUsingOPTIONSUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorHTMLUsingOPTIONSForbidden creates a ErrorHTMLUsingOPTIONSForbidden with default headers values
func NewErrorHTMLUsingOPTIONSForbidden() *ErrorHTMLUsingOPTIONSForbidden {
	return &ErrorHTMLUsingOPTIONSForbidden{}
}

/*ErrorHTMLUsingOPTIONSForbidden handles this case with default header values.

Forbidden
*/
type ErrorHTMLUsingOPTIONSForbidden struct {
}

func (o *ErrorHTMLUsingOPTIONSForbidden) Error() string {
	return fmt.Sprintf("[OPTIONS /error][%d] errorHtmlUsingOPTIONSForbidden ", 403)
}

func (o *ErrorHTMLUsingOPTIONSForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
