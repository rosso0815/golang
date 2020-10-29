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

// ErrorHTMLUsingHEADReader is a Reader for the ErrorHTMLUsingHEAD structure.
type ErrorHTMLUsingHEADReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ErrorHTMLUsingHEADReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewErrorHTMLUsingHEADOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewErrorHTMLUsingHEADNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewErrorHTMLUsingHEADUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewErrorHTMLUsingHEADForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewErrorHTMLUsingHEADOK creates a ErrorHTMLUsingHEADOK with default headers values
func NewErrorHTMLUsingHEADOK() *ErrorHTMLUsingHEADOK {
	return &ErrorHTMLUsingHEADOK{}
}

/*ErrorHTMLUsingHEADOK handles this case with default header values.

OK
*/
type ErrorHTMLUsingHEADOK struct {
	Payload *models.ModelAndView
}

func (o *ErrorHTMLUsingHEADOK) Error() string {
	return fmt.Sprintf("[HEAD /error][%d] errorHtmlUsingHEADOK  %+v", 200, o.Payload)
}

func (o *ErrorHTMLUsingHEADOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelAndView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewErrorHTMLUsingHEADNoContent creates a ErrorHTMLUsingHEADNoContent with default headers values
func NewErrorHTMLUsingHEADNoContent() *ErrorHTMLUsingHEADNoContent {
	return &ErrorHTMLUsingHEADNoContent{}
}

/*ErrorHTMLUsingHEADNoContent handles this case with default header values.

No Content
*/
type ErrorHTMLUsingHEADNoContent struct {
}

func (o *ErrorHTMLUsingHEADNoContent) Error() string {
	return fmt.Sprintf("[HEAD /error][%d] errorHtmlUsingHEADNoContent ", 204)
}

func (o *ErrorHTMLUsingHEADNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorHTMLUsingHEADUnauthorized creates a ErrorHTMLUsingHEADUnauthorized with default headers values
func NewErrorHTMLUsingHEADUnauthorized() *ErrorHTMLUsingHEADUnauthorized {
	return &ErrorHTMLUsingHEADUnauthorized{}
}

/*ErrorHTMLUsingHEADUnauthorized handles this case with default header values.

Unauthorized
*/
type ErrorHTMLUsingHEADUnauthorized struct {
}

func (o *ErrorHTMLUsingHEADUnauthorized) Error() string {
	return fmt.Sprintf("[HEAD /error][%d] errorHtmlUsingHEADUnauthorized ", 401)
}

func (o *ErrorHTMLUsingHEADUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorHTMLUsingHEADForbidden creates a ErrorHTMLUsingHEADForbidden with default headers values
func NewErrorHTMLUsingHEADForbidden() *ErrorHTMLUsingHEADForbidden {
	return &ErrorHTMLUsingHEADForbidden{}
}

/*ErrorHTMLUsingHEADForbidden handles this case with default header values.

Forbidden
*/
type ErrorHTMLUsingHEADForbidden struct {
}

func (o *ErrorHTMLUsingHEADForbidden) Error() string {
	return fmt.Sprintf("[HEAD /error][%d] errorHtmlUsingHEADForbidden ", 403)
}

func (o *ErrorHTMLUsingHEADForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}