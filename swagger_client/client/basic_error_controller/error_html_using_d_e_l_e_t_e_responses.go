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

// ErrorHTMLUsingDELETEReader is a Reader for the ErrorHTMLUsingDELETE structure.
type ErrorHTMLUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ErrorHTMLUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewErrorHTMLUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 204:
		result := NewErrorHTMLUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewErrorHTMLUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewErrorHTMLUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewErrorHTMLUsingDELETEOK creates a ErrorHTMLUsingDELETEOK with default headers values
func NewErrorHTMLUsingDELETEOK() *ErrorHTMLUsingDELETEOK {
	return &ErrorHTMLUsingDELETEOK{}
}

/*ErrorHTMLUsingDELETEOK handles this case with default header values.

OK
*/
type ErrorHTMLUsingDELETEOK struct {
	Payload *models.ModelAndView
}

func (o *ErrorHTMLUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /error][%d] errorHtmlUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *ErrorHTMLUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelAndView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewErrorHTMLUsingDELETENoContent creates a ErrorHTMLUsingDELETENoContent with default headers values
func NewErrorHTMLUsingDELETENoContent() *ErrorHTMLUsingDELETENoContent {
	return &ErrorHTMLUsingDELETENoContent{}
}

/*ErrorHTMLUsingDELETENoContent handles this case with default header values.

No Content
*/
type ErrorHTMLUsingDELETENoContent struct {
}

func (o *ErrorHTMLUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /error][%d] errorHtmlUsingDELETENoContent ", 204)
}

func (o *ErrorHTMLUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorHTMLUsingDELETEUnauthorized creates a ErrorHTMLUsingDELETEUnauthorized with default headers values
func NewErrorHTMLUsingDELETEUnauthorized() *ErrorHTMLUsingDELETEUnauthorized {
	return &ErrorHTMLUsingDELETEUnauthorized{}
}

/*ErrorHTMLUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type ErrorHTMLUsingDELETEUnauthorized struct {
}

func (o *ErrorHTMLUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /error][%d] errorHtmlUsingDELETEUnauthorized ", 401)
}

func (o *ErrorHTMLUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorHTMLUsingDELETEForbidden creates a ErrorHTMLUsingDELETEForbidden with default headers values
func NewErrorHTMLUsingDELETEForbidden() *ErrorHTMLUsingDELETEForbidden {
	return &ErrorHTMLUsingDELETEForbidden{}
}

/*ErrorHTMLUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type ErrorHTMLUsingDELETEForbidden struct {
}

func (o *ErrorHTMLUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /error][%d] errorHtmlUsingDELETEForbidden ", 403)
}

func (o *ErrorHTMLUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
