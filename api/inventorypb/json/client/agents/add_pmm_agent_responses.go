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

// AddPMMAgentReader is a Reader for the AddPMMAgent structure.
type AddPMMAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPMMAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddPMMAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddPMMAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPMMAgentOK creates a AddPMMAgentOK with default headers values
func NewAddPMMAgentOK() *AddPMMAgentOK {
	return &AddPMMAgentOK{}
}

/* AddPMMAgentOK describes a response with status code 200, with default header values.

A successful response.
*/
type AddPMMAgentOK struct {
	Payload *AddPMMAgentOKBody
}

func (o *AddPMMAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddPMMAgent][%d] addPmmAgentOk  %+v", 200, o.Payload)
}
func (o *AddPMMAgentOK) GetPayload() *AddPMMAgentOKBody {
	return o.Payload
}

func (o *AddPMMAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPMMAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPMMAgentDefault creates a AddPMMAgentDefault with default headers values
func NewAddPMMAgentDefault(code int) *AddPMMAgentDefault {
	return &AddPMMAgentDefault{
		_statusCode: code,
	}
}

/* AddPMMAgentDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AddPMMAgentDefault struct {
	_statusCode int

	Payload *AddPMMAgentDefaultBody
}

// Code gets the status code for the add PMM agent default response
func (o *AddPMMAgentDefault) Code() int {
	return o._statusCode
}

func (o *AddPMMAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddPMMAgent][%d] AddPMMAgent default  %+v", o._statusCode, o.Payload)
}
func (o *AddPMMAgentDefault) GetPayload() *AddPMMAgentDefaultBody {
	return o.Payload
}

func (o *AddPMMAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPMMAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddPMMAgentBody add PMM agent body
swagger:model AddPMMAgentBody
*/
type AddPMMAgentBody struct {

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add PMM agent body
func (o *AddPMMAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add PMM agent body based on context it is used
func (o *AddPMMAgentBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPMMAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPMMAgentBody) UnmarshalBinary(b []byte) error {
	var res AddPMMAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPMMAgentDefaultBody add PMM agent default body
swagger:model AddPMMAgentDefaultBody
*/
type AddPMMAgentDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*AddPMMAgentDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this add PMM agent default body
func (o *AddPMMAgentDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPMMAgentDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AddPMMAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddPMMAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this add PMM agent default body based on the context it is used
func (o *AddPMMAgentDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPMMAgentDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AddPMMAgent default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AddPMMAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPMMAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPMMAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddPMMAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPMMAgentDefaultBodyDetailsItems0 add PMM agent default body details items0
swagger:model AddPMMAgentDefaultBodyDetailsItems0
*/
type AddPMMAgentDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this add PMM agent default body details items0
func (o *AddPMMAgentDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add PMM agent default body details items0 based on context it is used
func (o *AddPMMAgentDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPMMAgentDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPMMAgentDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AddPMMAgentDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPMMAgentOKBody add PMM agent OK body
swagger:model AddPMMAgentOKBody
*/
type AddPMMAgentOKBody struct {

	// pmm agent
	PMMAgent *AddPMMAgentOKBodyPMMAgent `json:"pmm_agent,omitempty"`
}

// Validate validates this add PMM agent OK body
func (o *AddPMMAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePMMAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPMMAgentOKBody) validatePMMAgent(formats strfmt.Registry) error {
	if swag.IsZero(o.PMMAgent) { // not required
		return nil
	}

	if o.PMMAgent != nil {
		if err := o.PMMAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPmmAgentOk" + "." + "pmm_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addPmmAgentOk" + "." + "pmm_agent")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this add PMM agent OK body based on the context it is used
func (o *AddPMMAgentOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePMMAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPMMAgentOKBody) contextValidatePMMAgent(ctx context.Context, formats strfmt.Registry) error {

	if o.PMMAgent != nil {
		if err := o.PMMAgent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPmmAgentOk" + "." + "pmm_agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addPmmAgentOk" + "." + "pmm_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPMMAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPMMAgentOKBody) UnmarshalBinary(b []byte) error {
	var res AddPMMAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPMMAgentOKBodyPMMAgent PMMAgent runs on Generic or Container Node.
swagger:model AddPMMAgentOKBodyPMMAgent
*/
type AddPMMAgentOKBodyPMMAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if Agent is running and connected to pmm-managed.
	Connected bool `json:"connected,omitempty"`

	// Path to exec process
	ProcessExecPath string `json:"process_exec_path,omitempty"`
}

// Validate validates this add PMM agent OK body PMM agent
func (o *AddPMMAgentOKBodyPMMAgent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add PMM agent OK body PMM agent based on context it is used
func (o *AddPMMAgentOKBodyPMMAgent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPMMAgentOKBodyPMMAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPMMAgentOKBodyPMMAgent) UnmarshalBinary(b []byte) error {
	var res AddPMMAgentOKBodyPMMAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
