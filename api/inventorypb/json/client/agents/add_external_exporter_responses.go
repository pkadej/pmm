// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddExternalExporterReader is a Reader for the AddExternalExporter structure.
type AddExternalExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddExternalExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddExternalExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddExternalExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddExternalExporterOK creates a AddExternalExporterOK with default headers values
func NewAddExternalExporterOK() *AddExternalExporterOK {
	return &AddExternalExporterOK{}
}

/* AddExternalExporterOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddExternalExporterOK struct {
	Payload *AddExternalExporterOKBody
}

func (o *AddExternalExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddExternalExporter][%d] addExternalExporterOk  %+v", 200, o.Payload)
}
func (o *AddExternalExporterOK) GetPayload() *AddExternalExporterOKBody {
	return o.Payload
}

func (o *AddExternalExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddExternalExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddExternalExporterDefault creates a AddExternalExporterDefault with default headers values
func NewAddExternalExporterDefault(code int) *AddExternalExporterDefault {
	return &AddExternalExporterDefault{
		_statusCode: code,
	}
}

/* AddExternalExporterDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddExternalExporterDefault struct {
	_statusCode int

	Payload *AddExternalExporterDefaultBody
}

// Code gets the status code for the add external exporter default response
func (o *AddExternalExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddExternalExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddExternalExporter][%d] AddExternalExporter default  %+v", o._statusCode, o.Payload)
}
func (o *AddExternalExporterDefault) GetPayload() *AddExternalExporterDefaultBody {
	return o.Payload
}

func (o *AddExternalExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddExternalExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddExternalExporterBody add external exporter body
swagger:model AddExternalExporterBody
*/
type AddExternalExporterBody struct {

	// The node identifier where this instance is run.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// HTTP basic auth username for collecting metrics.
	Username string `json:"username,omitempty"`

	// HTTP basic auth password for collecting metrics.
	Password string `json:"password,omitempty"`

	// Scheme to generate URI to exporter metrics endpoints(default: http).
	Scheme string `json:"scheme,omitempty"`

	// Path under which metrics are exposed, used to generate URI(default: /metrics).
	MetricsPath string `json:"metrics_path,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Enables push metrics mode for exporter.
	PushMetrics bool `json:"push_metrics,omitempty"`
}

// Validate validates this add external exporter body
func (o *AddExternalExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add external exporter body based on context it is used
func (o *AddExternalExporterBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddExternalExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddExternalExporterBody) UnmarshalBinary(b []byte) error {
	var res AddExternalExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddExternalExporterDefaultBody add external exporter default body
swagger:model AddExternalExporterDefaultBody
*/
type AddExternalExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddExternalExporterDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add external exporter default body
func (o *AddExternalExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddExternalExporterDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddExternalExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddExternalExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add external exporter default body based on the context it is used
func (o *AddExternalExporterDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddExternalExporterDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddExternalExporter default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddExternalExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddExternalExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddExternalExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddExternalExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddExternalExporterDefaultBodyDetailsItems0 add external exporter default body details items0
swagger:model AddExternalExporterDefaultBodyDetailsItems0
*/
type AddExternalExporterDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this add external exporter default body details items0
func (o *AddExternalExporterDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add external exporter default body details items0 based on context it is used
func (o *AddExternalExporterDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddExternalExporterDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddExternalExporterDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddExternalExporterDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddExternalExporterOKBody add external exporter OK body
swagger:model AddExternalExporterOKBody
*/
type AddExternalExporterOKBody struct {

	// external exporter
	ExternalExporter *AddExternalExporterOKBodyExternalExporter `json:"external_exporter,omitempty"`
}

// Validate validates this add external exporter OK body
func (o *AddExternalExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExternalExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddExternalExporterOKBody) validateExternalExporter(formats strfmt.Registry) error {
	if swag.IsZero(o.ExternalExporter) { // not required
		return nil
	}

	if o.ExternalExporter != nil {
		if err := o.ExternalExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addExternalExporterOk" + "." + "external_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addExternalExporterOk" + "." + "external_exporter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add external exporter OK body based on the context it is used
func (o *AddExternalExporterOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateExternalExporter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddExternalExporterOKBody) contextValidateExternalExporter(ctx context.Context, formats strfmt.Registry) error {

	if o.ExternalExporter != nil {
		if err := o.ExternalExporter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addExternalExporterOk" + "." + "external_exporter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addExternalExporterOk" + "." + "external_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddExternalExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddExternalExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddExternalExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddExternalExporterOKBodyExternalExporter ExternalExporter runs on any Node type, including Remote Node.
swagger:model AddExternalExporterOKBodyExternalExporter
*/
type AddExternalExporterOKBodyExternalExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// If disabled, metrics from this exporter will not be collected.
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// HTTP basic auth username for collecting metrics.
	Username string `json:"username,omitempty"`

	// Scheme to generate URI to exporter metrics endpoints.
	Scheme string `json:"scheme,omitempty"`

	// Path under which metrics are exposed, used to generate URI.
	MetricsPath string `json:"metrics_path,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// Path to exec process
	ProcessExecPath string `json:"process_exec_path,omitempty"`
}

// Validate validates this add external exporter OK body external exporter
func (o *AddExternalExporterOKBodyExternalExporter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add external exporter OK body external exporter based on context it is used
func (o *AddExternalExporterOKBodyExternalExporter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddExternalExporterOKBodyExternalExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddExternalExporterOKBodyExternalExporter) UnmarshalBinary(b []byte) error {
	var res AddExternalExporterOKBodyExternalExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
