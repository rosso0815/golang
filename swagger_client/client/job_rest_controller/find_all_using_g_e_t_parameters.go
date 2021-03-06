// Code generated by go-swagger; DO NOT EDIT.

package job_rest_controller

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

// NewFindAllUsingGETParams creates a new FindAllUsingGETParams object
// with the default values initialized.
func NewFindAllUsingGETParams() *FindAllUsingGETParams {

	return &FindAllUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindAllUsingGETParamsWithTimeout creates a new FindAllUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindAllUsingGETParamsWithTimeout(timeout time.Duration) *FindAllUsingGETParams {

	return &FindAllUsingGETParams{

		timeout: timeout,
	}
}

// NewFindAllUsingGETParamsWithContext creates a new FindAllUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindAllUsingGETParamsWithContext(ctx context.Context) *FindAllUsingGETParams {

	return &FindAllUsingGETParams{

		Context: ctx,
	}
}

// NewFindAllUsingGETParamsWithHTTPClient creates a new FindAllUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindAllUsingGETParamsWithHTTPClient(client *http.Client) *FindAllUsingGETParams {

	return &FindAllUsingGETParams{
		HTTPClient: client,
	}
}

/*FindAllUsingGETParams contains all the parameters to send to the API endpoint
for the find all using g e t operation typically these are written to a http.Request
*/
type FindAllUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find all using g e t params
func (o *FindAllUsingGETParams) WithTimeout(timeout time.Duration) *FindAllUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find all using g e t params
func (o *FindAllUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find all using g e t params
func (o *FindAllUsingGETParams) WithContext(ctx context.Context) *FindAllUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find all using g e t params
func (o *FindAllUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find all using g e t params
func (o *FindAllUsingGETParams) WithHTTPClient(client *http.Client) *FindAllUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find all using g e t params
func (o *FindAllUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindAllUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
