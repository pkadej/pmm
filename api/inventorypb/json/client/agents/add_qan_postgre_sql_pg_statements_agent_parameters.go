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

// NewAddQANPostgreSQLPgStatementsAgentParams creates a new AddQANPostgreSQLPgStatementsAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddQANPostgreSQLPgStatementsAgentParams() *AddQANPostgreSQLPgStatementsAgentParams {
	return &AddQANPostgreSQLPgStatementsAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddQANPostgreSQLPgStatementsAgentParamsWithTimeout creates a new AddQANPostgreSQLPgStatementsAgentParams object
// with the ability to set a timeout on a request.
func NewAddQANPostgreSQLPgStatementsAgentParamsWithTimeout(timeout time.Duration) *AddQANPostgreSQLPgStatementsAgentParams {
	return &AddQANPostgreSQLPgStatementsAgentParams{
		timeout: timeout,
	}
}

// NewAddQANPostgreSQLPgStatementsAgentParamsWithContext creates a new AddQANPostgreSQLPgStatementsAgentParams object
// with the ability to set a context for a request.
func NewAddQANPostgreSQLPgStatementsAgentParamsWithContext(ctx context.Context) *AddQANPostgreSQLPgStatementsAgentParams {
	return &AddQANPostgreSQLPgStatementsAgentParams{
		Context: ctx,
	}
}

// NewAddQANPostgreSQLPgStatementsAgentParamsWithHTTPClient creates a new AddQANPostgreSQLPgStatementsAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddQANPostgreSQLPgStatementsAgentParamsWithHTTPClient(client *http.Client) *AddQANPostgreSQLPgStatementsAgentParams {
	return &AddQANPostgreSQLPgStatementsAgentParams{
		HTTPClient: client,
	}
}

/* AddQANPostgreSQLPgStatementsAgentParams contains all the parameters to send to the API endpoint
   for the add QAN postgre SQL pg statements agent operation.

   Typically these are written to a http.Request.
*/
type AddQANPostgreSQLPgStatementsAgentParams struct {

	// Body.
	Body AddQANPostgreSQLPgStatementsAgentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add QAN postgre SQL pg statements agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddQANPostgreSQLPgStatementsAgentParams) WithDefaults() *AddQANPostgreSQLPgStatementsAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add QAN postgre SQL pg statements agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddQANPostgreSQLPgStatementsAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add QAN postgre SQL pg statements agent params
func (o *AddQANPostgreSQLPgStatementsAgentParams) WithTimeout(timeout time.Duration) *AddQANPostgreSQLPgStatementsAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add QAN postgre SQL pg statements agent params
func (o *AddQANPostgreSQLPgStatementsAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add QAN postgre SQL pg statements agent params
func (o *AddQANPostgreSQLPgStatementsAgentParams) WithContext(ctx context.Context) *AddQANPostgreSQLPgStatementsAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add QAN postgre SQL pg statements agent params
func (o *AddQANPostgreSQLPgStatementsAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add QAN postgre SQL pg statements agent params
func (o *AddQANPostgreSQLPgStatementsAgentParams) WithHTTPClient(client *http.Client) *AddQANPostgreSQLPgStatementsAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add QAN postgre SQL pg statements agent params
func (o *AddQANPostgreSQLPgStatementsAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add QAN postgre SQL pg statements agent params
func (o *AddQANPostgreSQLPgStatementsAgentParams) WithBody(body AddQANPostgreSQLPgStatementsAgentBody) *AddQANPostgreSQLPgStatementsAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add QAN postgre SQL pg statements agent params
func (o *AddQANPostgreSQLPgStatementsAgentParams) SetBody(body AddQANPostgreSQLPgStatementsAgentBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddQANPostgreSQLPgStatementsAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
