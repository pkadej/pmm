// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// NewConnectParams creates a new ConnectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConnectParams() *ConnectParams {
	return &ConnectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConnectParamsWithTimeout creates a new ConnectParams object
// with the ability to set a timeout on a request.
func NewConnectParamsWithTimeout(timeout time.Duration) *ConnectParams {
	return &ConnectParams{
		timeout: timeout,
	}
}

// NewConnectParamsWithContext creates a new ConnectParams object
// with the ability to set a context for a request.
func NewConnectParamsWithContext(ctx context.Context) *ConnectParams {
	return &ConnectParams{
		Context: ctx,
	}
}

// NewConnectParamsWithHTTPClient creates a new ConnectParams object
// with the ability to set a custom HTTPClient for a request.
func NewConnectParamsWithHTTPClient(client *http.Client) *ConnectParams {
	return &ConnectParams{
		HTTPClient: client,
	}
}

/* ConnectParams contains all the parameters to send to the API endpoint
   for the connect operation.

   Typically these are written to a http.Request.
*/
type ConnectParams struct {

	// Body.
	Body ConnectBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the connect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectParams) WithDefaults() *ConnectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the connect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the connect params
func (o *ConnectParams) WithTimeout(timeout time.Duration) *ConnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connect params
func (o *ConnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connect params
func (o *ConnectParams) WithContext(ctx context.Context) *ConnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connect params
func (o *ConnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connect params
func (o *ConnectParams) WithHTTPClient(client *http.Client) *ConnectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connect params
func (o *ConnectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the connect params
func (o *ConnectParams) WithBody(body ConnectBody) *ConnectParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the connect params
func (o *ConnectParams) SetBody(body ConnectBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
