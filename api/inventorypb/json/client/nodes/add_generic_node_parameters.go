// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// NewAddGenericNodeParams creates a new AddGenericNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddGenericNodeParams() *AddGenericNodeParams {
	return &AddGenericNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddGenericNodeParamsWithTimeout creates a new AddGenericNodeParams object
// with the ability to set a timeout on a request.
func NewAddGenericNodeParamsWithTimeout(timeout time.Duration) *AddGenericNodeParams {
	return &AddGenericNodeParams{
		timeout: timeout,
	}
}

// NewAddGenericNodeParamsWithContext creates a new AddGenericNodeParams object
// with the ability to set a context for a request.
func NewAddGenericNodeParamsWithContext(ctx context.Context) *AddGenericNodeParams {
	return &AddGenericNodeParams{
		Context: ctx,
	}
}

// NewAddGenericNodeParamsWithHTTPClient creates a new AddGenericNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddGenericNodeParamsWithHTTPClient(client *http.Client) *AddGenericNodeParams {
	return &AddGenericNodeParams{
		HTTPClient: client,
	}
}

/* AddGenericNodeParams contains all the parameters to send to the API endpoint
   for the add generic node operation.

   Typically these are written to a http.Request.
*/
type AddGenericNodeParams struct {

	// Body.
	Body AddGenericNodeBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add generic node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddGenericNodeParams) WithDefaults() *AddGenericNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add generic node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddGenericNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add generic node params
func (o *AddGenericNodeParams) WithTimeout(timeout time.Duration) *AddGenericNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add generic node params
func (o *AddGenericNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add generic node params
func (o *AddGenericNodeParams) WithContext(ctx context.Context) *AddGenericNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add generic node params
func (o *AddGenericNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add generic node params
func (o *AddGenericNodeParams) WithHTTPClient(client *http.Client) *AddGenericNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add generic node params
func (o *AddGenericNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add generic node params
func (o *AddGenericNodeParams) WithBody(body AddGenericNodeBody) *AddGenericNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add generic node params
func (o *AddGenericNodeParams) SetBody(body AddGenericNodeBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddGenericNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
