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

// NewInstallOperatorParams creates a new InstallOperatorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInstallOperatorParams() *InstallOperatorParams {
	return &InstallOperatorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInstallOperatorParamsWithTimeout creates a new InstallOperatorParams object
// with the ability to set a timeout on a request.
func NewInstallOperatorParamsWithTimeout(timeout time.Duration) *InstallOperatorParams {
	return &InstallOperatorParams{
		timeout: timeout,
	}
}

// NewInstallOperatorParamsWithContext creates a new InstallOperatorParams object
// with the ability to set a context for a request.
func NewInstallOperatorParamsWithContext(ctx context.Context) *InstallOperatorParams {
	return &InstallOperatorParams{
		Context: ctx,
	}
}

// NewInstallOperatorParamsWithHTTPClient creates a new InstallOperatorParams object
// with the ability to set a custom HTTPClient for a request.
func NewInstallOperatorParamsWithHTTPClient(client *http.Client) *InstallOperatorParams {
	return &InstallOperatorParams{
		HTTPClient: client,
	}
}

/* InstallOperatorParams contains all the parameters to send to the API endpoint
   for the install operator operation.

   Typically these are written to a http.Request.
*/
type InstallOperatorParams struct {

	// Body.
	Body InstallOperatorBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the install operator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallOperatorParams) WithDefaults() *InstallOperatorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the install operator params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InstallOperatorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the install operator params
func (o *InstallOperatorParams) WithTimeout(timeout time.Duration) *InstallOperatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the install operator params
func (o *InstallOperatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the install operator params
func (o *InstallOperatorParams) WithContext(ctx context.Context) *InstallOperatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the install operator params
func (o *InstallOperatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the install operator params
func (o *InstallOperatorParams) WithHTTPClient(client *http.Client) *InstallOperatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the install operator params
func (o *InstallOperatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the install operator params
func (o *InstallOperatorParams) WithBody(body InstallOperatorBody) *InstallOperatorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the install operator params
func (o *InstallOperatorParams) SetBody(body InstallOperatorBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InstallOperatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
