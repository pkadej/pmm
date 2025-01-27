// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewChangeProxySQLExporterParams creates a new ChangeProxySQLExporterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeProxySQLExporterParams() *ChangeProxySQLExporterParams {
	return &ChangeProxySQLExporterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeProxySQLExporterParamsWithTimeout creates a new ChangeProxySQLExporterParams object
// with the ability to set a timeout on a request.
func NewChangeProxySQLExporterParamsWithTimeout(timeout time.Duration) *ChangeProxySQLExporterParams {
	return &ChangeProxySQLExporterParams{
		timeout: timeout,
	}
}

// NewChangeProxySQLExporterParamsWithContext creates a new ChangeProxySQLExporterParams object
// with the ability to set a context for a request.
func NewChangeProxySQLExporterParamsWithContext(ctx context.Context) *ChangeProxySQLExporterParams {
	return &ChangeProxySQLExporterParams{
		Context: ctx,
	}
}

// NewChangeProxySQLExporterParamsWithHTTPClient creates a new ChangeProxySQLExporterParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeProxySQLExporterParamsWithHTTPClient(client *http.Client) *ChangeProxySQLExporterParams {
	return &ChangeProxySQLExporterParams{
		HTTPClient: client,
	}
}

/* ChangeProxySQLExporterParams contains all the parameters to send to the API endpoint
   for the change proxy SQL exporter operation.

   Typically these are written to a http.Request.
*/
type ChangeProxySQLExporterParams struct {

	// Body.
	Body ChangeProxySQLExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change proxy SQL exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeProxySQLExporterParams) WithDefaults() *ChangeProxySQLExporterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change proxy SQL exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeProxySQLExporterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change proxy SQL exporter params
func (o *ChangeProxySQLExporterParams) WithTimeout(timeout time.Duration) *ChangeProxySQLExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change proxy SQL exporter params
func (o *ChangeProxySQLExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change proxy SQL exporter params
func (o *ChangeProxySQLExporterParams) WithContext(ctx context.Context) *ChangeProxySQLExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change proxy SQL exporter params
func (o *ChangeProxySQLExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change proxy SQL exporter params
func (o *ChangeProxySQLExporterParams) WithHTTPClient(client *http.Client) *ChangeProxySQLExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change proxy SQL exporter params
func (o *ChangeProxySQLExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change proxy SQL exporter params
func (o *ChangeProxySQLExporterParams) WithBody(body ChangeProxySQLExporterBody) *ChangeProxySQLExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change proxy SQL exporter params
func (o *ChangeProxySQLExporterParams) SetBody(body ChangeProxySQLExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeProxySQLExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
