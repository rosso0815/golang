// Code generated by go-swagger; DO NOT EDIT.

package timesheet_rest_controller

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

// NewRetrieveAllTimeSheetUsingGETParams creates a new RetrieveAllTimeSheetUsingGETParams object
// with the default values initialized.
func NewRetrieveAllTimeSheetUsingGETParams() *RetrieveAllTimeSheetUsingGETParams {

	return &RetrieveAllTimeSheetUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRetrieveAllTimeSheetUsingGETParamsWithTimeout creates a new RetrieveAllTimeSheetUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRetrieveAllTimeSheetUsingGETParamsWithTimeout(timeout time.Duration) *RetrieveAllTimeSheetUsingGETParams {

	return &RetrieveAllTimeSheetUsingGETParams{

		timeout: timeout,
	}
}

// NewRetrieveAllTimeSheetUsingGETParamsWithContext creates a new RetrieveAllTimeSheetUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewRetrieveAllTimeSheetUsingGETParamsWithContext(ctx context.Context) *RetrieveAllTimeSheetUsingGETParams {

	return &RetrieveAllTimeSheetUsingGETParams{

		Context: ctx,
	}
}

// NewRetrieveAllTimeSheetUsingGETParamsWithHTTPClient creates a new RetrieveAllTimeSheetUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRetrieveAllTimeSheetUsingGETParamsWithHTTPClient(client *http.Client) *RetrieveAllTimeSheetUsingGETParams {

	return &RetrieveAllTimeSheetUsingGETParams{
		HTTPClient: client,
	}
}

/*RetrieveAllTimeSheetUsingGETParams contains all the parameters to send to the API endpoint
for the retrieve all time sheet using g e t operation typically these are written to a http.Request
*/
type RetrieveAllTimeSheetUsingGETParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the retrieve all time sheet using g e t params
func (o *RetrieveAllTimeSheetUsingGETParams) WithTimeout(timeout time.Duration) *RetrieveAllTimeSheetUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the retrieve all time sheet using g e t params
func (o *RetrieveAllTimeSheetUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the retrieve all time sheet using g e t params
func (o *RetrieveAllTimeSheetUsingGETParams) WithContext(ctx context.Context) *RetrieveAllTimeSheetUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the retrieve all time sheet using g e t params
func (o *RetrieveAllTimeSheetUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the retrieve all time sheet using g e t params
func (o *RetrieveAllTimeSheetUsingGETParams) WithHTTPClient(client *http.Client) *RetrieveAllTimeSheetUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the retrieve all time sheet using g e t params
func (o *RetrieveAllTimeSheetUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RetrieveAllTimeSheetUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
