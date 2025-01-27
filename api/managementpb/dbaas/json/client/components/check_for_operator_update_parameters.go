// Code generated by go-swagger; DO NOT EDIT.

package components

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCheckForOperatorUpdateParams creates a new CheckForOperatorUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckForOperatorUpdateParams() *CheckForOperatorUpdateParams {
	return &CheckForOperatorUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckForOperatorUpdateParamsWithTimeout creates a new CheckForOperatorUpdateParams object
// with the ability to set a timeout on a request.
func NewCheckForOperatorUpdateParamsWithTimeout(timeout time.Duration) *CheckForOperatorUpdateParams {
	return &CheckForOperatorUpdateParams{
		timeout: timeout,
	}
}

// NewCheckForOperatorUpdateParamsWithContext creates a new CheckForOperatorUpdateParams object
// with the ability to set a context for a request.
func NewCheckForOperatorUpdateParamsWithContext(ctx context.Context) *CheckForOperatorUpdateParams {
	return &CheckForOperatorUpdateParams{
		Context: ctx,
	}
}

// NewCheckForOperatorUpdateParamsWithHTTPClient creates a new CheckForOperatorUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckForOperatorUpdateParamsWithHTTPClient(client *http.Client) *CheckForOperatorUpdateParams {
	return &CheckForOperatorUpdateParams{
		HTTPClient: client,
	}
}

/* CheckForOperatorUpdateParams contains all the parameters to send to the API endpoint
   for the check for operator update operation.

   Typically these are written to a http.Request.
*/
type CheckForOperatorUpdateParams struct {

	// Body.
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the check for operator update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckForOperatorUpdateParams) WithDefaults() *CheckForOperatorUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the check for operator update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckForOperatorUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the check for operator update params
func (o *CheckForOperatorUpdateParams) WithTimeout(timeout time.Duration) *CheckForOperatorUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check for operator update params
func (o *CheckForOperatorUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check for operator update params
func (o *CheckForOperatorUpdateParams) WithContext(ctx context.Context) *CheckForOperatorUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check for operator update params
func (o *CheckForOperatorUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check for operator update params
func (o *CheckForOperatorUpdateParams) WithHTTPClient(client *http.Client) *CheckForOperatorUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check for operator update params
func (o *CheckForOperatorUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the check for operator update params
func (o *CheckForOperatorUpdateParams) WithBody(body interface{}) *CheckForOperatorUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the check for operator update params
func (o *CheckForOperatorUpdateParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CheckForOperatorUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
