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

// NewAddAzureExporterParams creates a new AddAzureExporterParams object
// with the default values initialized.
func NewAddAzureExporterParams() *AddAzureExporterParams {
	var ()
	return &AddAzureExporterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAzureExporterParamsWithTimeout creates a new AddAzureExporterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAzureExporterParamsWithTimeout(timeout time.Duration) *AddAzureExporterParams {
	var ()
	return &AddAzureExporterParams{

		timeout: timeout,
	}
}

// NewAddAzureExporterParamsWithContext creates a new AddAzureExporterParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAzureExporterParamsWithContext(ctx context.Context) *AddAzureExporterParams {
	var ()
	return &AddAzureExporterParams{

		Context: ctx,
	}
}

// NewAddAzureExporterParamsWithHTTPClient creates a new AddAzureExporterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAzureExporterParamsWithHTTPClient(client *http.Client) *AddAzureExporterParams {
	var ()
	return &AddAzureExporterParams{
		HTTPClient: client,
	}
}

/*AddAzureExporterParams contains all the parameters to send to the API endpoint
for the add azure exporter operation typically these are written to a http.Request
*/
type AddAzureExporterParams struct {

	/*Body*/
	Body AddAzureExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add azure exporter params
func (o *AddAzureExporterParams) WithTimeout(timeout time.Duration) *AddAzureExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add azure exporter params
func (o *AddAzureExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add azure exporter params
func (o *AddAzureExporterParams) WithContext(ctx context.Context) *AddAzureExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add azure exporter params
func (o *AddAzureExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add azure exporter params
func (o *AddAzureExporterParams) WithHTTPClient(client *http.Client) *AddAzureExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add azure exporter params
func (o *AddAzureExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add azure exporter params
func (o *AddAzureExporterParams) WithBody(body AddAzureExporterBody) *AddAzureExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add azure exporter params
func (o *AddAzureExporterParams) SetBody(body AddAzureExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAzureExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
