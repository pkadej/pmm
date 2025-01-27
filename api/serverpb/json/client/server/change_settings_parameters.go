// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewChangeSettingsParams creates a new ChangeSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeSettingsParams() *ChangeSettingsParams {
	return &ChangeSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeSettingsParamsWithTimeout creates a new ChangeSettingsParams object
// with the ability to set a timeout on a request.
func NewChangeSettingsParamsWithTimeout(timeout time.Duration) *ChangeSettingsParams {
	return &ChangeSettingsParams{
		timeout: timeout,
	}
}

// NewChangeSettingsParamsWithContext creates a new ChangeSettingsParams object
// with the ability to set a context for a request.
func NewChangeSettingsParamsWithContext(ctx context.Context) *ChangeSettingsParams {
	return &ChangeSettingsParams{
		Context: ctx,
	}
}

// NewChangeSettingsParamsWithHTTPClient creates a new ChangeSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeSettingsParamsWithHTTPClient(client *http.Client) *ChangeSettingsParams {
	return &ChangeSettingsParams{
		HTTPClient: client,
	}
}

/* ChangeSettingsParams contains all the parameters to send to the API endpoint
   for the change settings operation.

   Typically these are written to a http.Request.
*/
type ChangeSettingsParams struct {

	// Body.
	Body ChangeSettingsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeSettingsParams) WithDefaults() *ChangeSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change settings params
func (o *ChangeSettingsParams) WithTimeout(timeout time.Duration) *ChangeSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change settings params
func (o *ChangeSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change settings params
func (o *ChangeSettingsParams) WithContext(ctx context.Context) *ChangeSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change settings params
func (o *ChangeSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change settings params
func (o *ChangeSettingsParams) WithHTTPClient(client *http.Client) *ChangeSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change settings params
func (o *ChangeSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change settings params
func (o *ChangeSettingsParams) WithBody(body ChangeSettingsBody) *ChangeSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change settings params
func (o *ChangeSettingsParams) SetBody(body ChangeSettingsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
