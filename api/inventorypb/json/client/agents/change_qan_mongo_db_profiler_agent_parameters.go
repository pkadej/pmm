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

// NewChangeQANMongoDBProfilerAgentParams creates a new ChangeQANMongoDBProfilerAgentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeQANMongoDBProfilerAgentParams() *ChangeQANMongoDBProfilerAgentParams {
	return &ChangeQANMongoDBProfilerAgentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeQANMongoDBProfilerAgentParamsWithTimeout creates a new ChangeQANMongoDBProfilerAgentParams object
// with the ability to set a timeout on a request.
func NewChangeQANMongoDBProfilerAgentParamsWithTimeout(timeout time.Duration) *ChangeQANMongoDBProfilerAgentParams {
	return &ChangeQANMongoDBProfilerAgentParams{
		timeout: timeout,
	}
}

// NewChangeQANMongoDBProfilerAgentParamsWithContext creates a new ChangeQANMongoDBProfilerAgentParams object
// with the ability to set a context for a request.
func NewChangeQANMongoDBProfilerAgentParamsWithContext(ctx context.Context) *ChangeQANMongoDBProfilerAgentParams {
	return &ChangeQANMongoDBProfilerAgentParams{
		Context: ctx,
	}
}

// NewChangeQANMongoDBProfilerAgentParamsWithHTTPClient creates a new ChangeQANMongoDBProfilerAgentParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeQANMongoDBProfilerAgentParamsWithHTTPClient(client *http.Client) *ChangeQANMongoDBProfilerAgentParams {
	return &ChangeQANMongoDBProfilerAgentParams{
		HTTPClient: client,
	}
}

/* ChangeQANMongoDBProfilerAgentParams contains all the parameters to send to the API endpoint
   for the change QAN mongo DB profiler agent operation.

   Typically these are written to a http.Request.
*/
type ChangeQANMongoDBProfilerAgentParams struct {

	// Body.
	Body ChangeQANMongoDBProfilerAgentBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change QAN mongo DB profiler agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeQANMongoDBProfilerAgentParams) WithDefaults() *ChangeQANMongoDBProfilerAgentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change QAN mongo DB profiler agent params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeQANMongoDBProfilerAgentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change QAN mongo DB profiler agent params
func (o *ChangeQANMongoDBProfilerAgentParams) WithTimeout(timeout time.Duration) *ChangeQANMongoDBProfilerAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change QAN mongo DB profiler agent params
func (o *ChangeQANMongoDBProfilerAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change QAN mongo DB profiler agent params
func (o *ChangeQANMongoDBProfilerAgentParams) WithContext(ctx context.Context) *ChangeQANMongoDBProfilerAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change QAN mongo DB profiler agent params
func (o *ChangeQANMongoDBProfilerAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change QAN mongo DB profiler agent params
func (o *ChangeQANMongoDBProfilerAgentParams) WithHTTPClient(client *http.Client) *ChangeQANMongoDBProfilerAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change QAN mongo DB profiler agent params
func (o *ChangeQANMongoDBProfilerAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change QAN mongo DB profiler agent params
func (o *ChangeQANMongoDBProfilerAgentParams) WithBody(body ChangeQANMongoDBProfilerAgentBody) *ChangeQANMongoDBProfilerAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change QAN mongo DB profiler agent params
func (o *ChangeQANMongoDBProfilerAgentParams) SetBody(body ChangeQANMongoDBProfilerAgentBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeQANMongoDBProfilerAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
