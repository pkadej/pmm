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

// NewAddMySQLdExporterParams creates a new AddMySQLdExporterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddMySQLdExporterParams() *AddMySQLdExporterParams {
	return &AddMySQLdExporterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddMySQLdExporterParamsWithTimeout creates a new AddMySQLdExporterParams object
// with the ability to set a timeout on a request.
func NewAddMySQLdExporterParamsWithTimeout(timeout time.Duration) *AddMySQLdExporterParams {
	return &AddMySQLdExporterParams{
		timeout: timeout,
	}
}

// NewAddMySQLdExporterParamsWithContext creates a new AddMySQLdExporterParams object
// with the ability to set a context for a request.
func NewAddMySQLdExporterParamsWithContext(ctx context.Context) *AddMySQLdExporterParams {
	return &AddMySQLdExporterParams{
		Context: ctx,
	}
}

// NewAddMySQLdExporterParamsWithHTTPClient creates a new AddMySQLdExporterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddMySQLdExporterParamsWithHTTPClient(client *http.Client) *AddMySQLdExporterParams {
	return &AddMySQLdExporterParams{
		HTTPClient: client,
	}
}

/* AddMySQLdExporterParams contains all the parameters to send to the API endpoint
   for the add my s q ld exporter operation.

   Typically these are written to a http.Request.
*/
type AddMySQLdExporterParams struct {

	// Body.
	Body AddMySQLdExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add my s q ld exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddMySQLdExporterParams) WithDefaults() *AddMySQLdExporterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add my s q ld exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddMySQLdExporterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add my s q ld exporter params
func (o *AddMySQLdExporterParams) WithTimeout(timeout time.Duration) *AddMySQLdExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add my s q ld exporter params
func (o *AddMySQLdExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add my s q ld exporter params
func (o *AddMySQLdExporterParams) WithContext(ctx context.Context) *AddMySQLdExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add my s q ld exporter params
func (o *AddMySQLdExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add my s q ld exporter params
func (o *AddMySQLdExporterParams) WithHTTPClient(client *http.Client) *AddMySQLdExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add my s q ld exporter params
func (o *AddMySQLdExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add my s q ld exporter params
func (o *AddMySQLdExporterParams) WithBody(body AddMySQLdExporterBody) *AddMySQLdExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add my s q ld exporter params
func (o *AddMySQLdExporterParams) SetBody(body AddMySQLdExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddMySQLdExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
