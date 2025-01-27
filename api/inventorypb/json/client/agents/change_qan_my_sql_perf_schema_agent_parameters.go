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

// NewChangeQANMySQLPerfSchemaAgentParams creates a new ChangeQANMySQLPerfSchemaAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeQANMySQLPerfSchemaAgentParams() *ChangeQANMySQLPerfSchemaAgentParams {
	return &ChangeQANMySQLPerfSchemaAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeQANMySQLPerfSchemaAgentParamsWithTimeout creates a new ChangeQANMySQLPerfSchemaAgentParams object
// with the ability to set a timeout on a request.
func NewChangeQANMySQLPerfSchemaAgentParamsWithTimeout(timeout time.Duration) *ChangeQANMySQLPerfSchemaAgentParams {
	return &ChangeQANMySQLPerfSchemaAgentParams{
		timeout: timeout,
	}
}

// NewChangeQANMySQLPerfSchemaAgentParamsWithContext creates a new ChangeQANMySQLPerfSchemaAgentParams object
// with the ability to set a context for a request.
func NewChangeQANMySQLPerfSchemaAgentParamsWithContext(ctx context.Context) *ChangeQANMySQLPerfSchemaAgentParams {
	return &ChangeQANMySQLPerfSchemaAgentParams{
		Context: ctx,
	}
}

// NewChangeQANMySQLPerfSchemaAgentParamsWithHTTPClient creates a new ChangeQANMySQLPerfSchemaAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeQANMySQLPerfSchemaAgentParamsWithHTTPClient(client *http.Client) *ChangeQANMySQLPerfSchemaAgentParams {
	return &ChangeQANMySQLPerfSchemaAgentParams{
		HTTPClient: client,
	}
}

/* ChangeQANMySQLPerfSchemaAgentParams contains all the parameters to send to the API endpoint
   for the change QAN my SQL perf schema agent operation.

   Typically these are written to a http.Request.
*/
type ChangeQANMySQLPerfSchemaAgentParams struct {

	// Body.
	Body ChangeQANMySQLPerfSchemaAgentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change QAN my SQL perf schema agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeQANMySQLPerfSchemaAgentParams) WithDefaults() *ChangeQANMySQLPerfSchemaAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change QAN my SQL perf schema agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeQANMySQLPerfSchemaAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change QAN my SQL perf schema agent params
func (o *ChangeQANMySQLPerfSchemaAgentParams) WithTimeout(timeout time.Duration) *ChangeQANMySQLPerfSchemaAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change QAN my SQL perf schema agent params
func (o *ChangeQANMySQLPerfSchemaAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change QAN my SQL perf schema agent params
func (o *ChangeQANMySQLPerfSchemaAgentParams) WithContext(ctx context.Context) *ChangeQANMySQLPerfSchemaAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change QAN my SQL perf schema agent params
func (o *ChangeQANMySQLPerfSchemaAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change QAN my SQL perf schema agent params
func (o *ChangeQANMySQLPerfSchemaAgentParams) WithHTTPClient(client *http.Client) *ChangeQANMySQLPerfSchemaAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change QAN my SQL perf schema agent params
func (o *ChangeQANMySQLPerfSchemaAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change QAN my SQL perf schema agent params
func (o *ChangeQANMySQLPerfSchemaAgentParams) WithBody(body ChangeQANMySQLPerfSchemaAgentBody) *ChangeQANMySQLPerfSchemaAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change QAN my SQL perf schema agent params
func (o *ChangeQANMySQLPerfSchemaAgentParams) SetBody(body ChangeQANMySQLPerfSchemaAgentBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeQANMySQLPerfSchemaAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
