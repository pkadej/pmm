// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNodeParams creates a new GetNodeParams object
// with the default values initialized.
func NewGetNodeParams() *GetNodeParams {
	var ()
	return &GetNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodeParamsWithTimeout creates a new GetNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodeParamsWithTimeout(timeout time.Duration) *GetNodeParams {
	var ()
	return &GetNodeParams{

		timeout: timeout,
	}
}

// NewGetNodeParamsWithContext creates a new GetNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNodeParamsWithContext(ctx context.Context) *GetNodeParams {
	var ()
	return &GetNodeParams{

		Context: ctx,
	}
}

// NewGetNodeParamsWithHTTPClient creates a new GetNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNodeParamsWithHTTPClient(client *http.Client) *GetNodeParams {
	var ()
	return &GetNodeParams{
		HTTPClient: client,
	}
}

/*GetNodeParams contains all the parameters to send to the API endpoint
for the get node operation typically these are written to a http.Request
*/
type GetNodeParams struct {

	/*Body*/
	Body GetNodeBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get node params
func (o *GetNodeParams) WithTimeout(timeout time.Duration) *GetNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node params
func (o *GetNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node params
func (o *GetNodeParams) WithContext(ctx context.Context) *GetNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node params
func (o *GetNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node params
func (o *GetNodeParams) WithHTTPClient(client *http.Client) *GetNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node params
func (o *GetNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get node params
func (o *GetNodeParams) WithBody(body GetNodeBody) *GetNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get node params
func (o *GetNodeParams) SetBody(body GetNodeBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
