// Code generated by go-swagger; DO NOT EDIT.

package basic_error_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewErrorHTMLUsingGETParams creates a new ErrorHTMLUsingGETParams object
// with the default values initialized.
func NewErrorHTMLUsingGETParams() *ErrorHTMLUsingGETParams {

	return &ErrorHTMLUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewErrorHTMLUsingGETParamsWithTimeout creates a new ErrorHTMLUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewErrorHTMLUsingGETParamsWithTimeout(timeout time.Duration) *ErrorHTMLUsingGETParams {

	return &ErrorHTMLUsingGETParams{

		timeout: timeout,
	}
}

// NewErrorHTMLUsingGETParamsWithContext creates a new ErrorHTMLUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewErrorHTMLUsingGETParamsWithContext(ctx context.Context) *ErrorHTMLUsingGETParams {

	return &ErrorHTMLUsingGETParams{

		Context: ctx,
	}
}

// NewErrorHTMLUsingGETParamsWithHTTPClient creates a new ErrorHTMLUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewErrorHTMLUsingGETParamsWithHTTPClient(client *http.Client) *ErrorHTMLUsingGETParams {

	return &ErrorHTMLUsingGETParams{
		HTTPClient: client,
	}
}

/*ErrorHTMLUsingGETParams contains all the parameters to send to the API endpoint
for the error Html using g e t operation typically these are written to a http.Request
*/
type ErrorHTMLUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the error Html using g e t params
func (o *ErrorHTMLUsingGETParams) WithTimeout(timeout time.Duration) *ErrorHTMLUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the error Html using g e t params
func (o *ErrorHTMLUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the error Html using g e t params
func (o *ErrorHTMLUsingGETParams) WithContext(ctx context.Context) *ErrorHTMLUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the error Html using g e t params
func (o *ErrorHTMLUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the error Html using g e t params
func (o *ErrorHTMLUsingGETParams) WithHTTPClient(client *http.Client) *ErrorHTMLUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the error Html using g e t params
func (o *ErrorHTMLUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ErrorHTMLUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
