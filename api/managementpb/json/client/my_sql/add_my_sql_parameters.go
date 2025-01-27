// Code generated by go-swagger; DO NOT EDIT.

package my_sql

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

// NewAddMySQLParams creates a new AddMySQLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddMySQLParams() *AddMySQLParams {
	return &AddMySQLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddMySQLParamsWithTimeout creates a new AddMySQLParams object
// with the ability to set a timeout on a request.
func NewAddMySQLParamsWithTimeout(timeout time.Duration) *AddMySQLParams {
	return &AddMySQLParams{
		timeout: timeout,
	}
}

// NewAddMySQLParamsWithContext creates a new AddMySQLParams object
// with the ability to set a context for a request.
func NewAddMySQLParamsWithContext(ctx context.Context) *AddMySQLParams {
	return &AddMySQLParams{
		Context: ctx,
	}
}

// NewAddMySQLParamsWithHTTPClient creates a new AddMySQLParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddMySQLParamsWithHTTPClient(client *http.Client) *AddMySQLParams {
	return &AddMySQLParams{
		HTTPClient: client,
	}
}

/* AddMySQLParams contains all the parameters to send to the API endpoint
   for the add my SQL operation.

   Typically these are written to a http.Request.
*/
type AddMySQLParams struct {

	// Body.
	Body AddMySQLBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add my SQL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddMySQLParams) WithDefaults() *AddMySQLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add my SQL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddMySQLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add my SQL params
func (o *AddMySQLParams) WithTimeout(timeout time.Duration) *AddMySQLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add my SQL params
func (o *AddMySQLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add my SQL params
func (o *AddMySQLParams) WithContext(ctx context.Context) *AddMySQLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add my SQL params
func (o *AddMySQLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add my SQL params
func (o *AddMySQLParams) WithHTTPClient(client *http.Client) *AddMySQLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add my SQL params
func (o *AddMySQLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add my SQL params
func (o *AddMySQLParams) WithBody(body AddMySQLBody) *AddMySQLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add my SQL params
func (o *AddMySQLParams) SetBody(body AddMySQLBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddMySQLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
