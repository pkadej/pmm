// Code generated by go-swagger; DO NOT EDIT.

package actions

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

// CancelActionReader is a Reader for the CancelAction structure.
type CancelActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCancelActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelActionOK creates a CancelActionOK with default headers values
func NewCancelActionOK() *CancelActionOK {
	return &CancelActionOK{}
}

/* CancelActionOK describes a response with status code 200, with default header values.

A successful response.
*/
type CancelActionOK struct {
	Payload interface{}
}

func (o *CancelActionOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/Cancel][%d] cancelActionOk  %+v", 200, o.Payload)
}
func (o *CancelActionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CancelActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelActionDefault creates a CancelActionDefault with default headers values
func NewCancelActionDefault(code int) *CancelActionDefault {
	return &CancelActionDefault{
		_statusCode: code,
	}
}

/* CancelActionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CancelActionDefault struct {
	_statusCode int

	Payload *CancelActionDefaultBody
}

// Code gets the status code for the cancel action default response
func (o *CancelActionDefault) Code() int {
	return o._statusCode
}

func (o *CancelActionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Actions/Cancel][%d] CancelAction default  %+v", o._statusCode, o.Payload)
}
func (o *CancelActionDefault) GetPayload() *CancelActionDefaultBody {
	return o.Payload
}

func (o *CancelActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CancelActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CancelActionBody cancel action body
swagger:model CancelActionBody
*/
type CancelActionBody struct {

	// Unique Action ID. Required.
	ActionID string `json:"action_id,omitempty"`
}

// Validate validates this cancel action body
func (o *CancelActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cancel action body based on context it is used
func (o *CancelActionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CancelActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelActionBody) UnmarshalBinary(b []byte) error {
	var res CancelActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CancelActionDefaultBody cancel action default body
swagger:model CancelActionDefaultBody
*/
type CancelActionDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*CancelActionDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this cancel action default body
func (o *CancelActionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CancelActionDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("CancelAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CancelAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cancel action default body based on the context it is used
func (o *CancelActionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CancelActionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CancelAction default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CancelAction default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CancelActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res CancelActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CancelActionDefaultBodyDetailsItems0 cancel action default body details items0
swagger:model CancelActionDefaultBodyDetailsItems0
*/
type CancelActionDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this cancel action default body details items0
func (o *CancelActionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cancel action default body details items0 based on context it is used
func (o *CancelActionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CancelActionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CancelActionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res CancelActionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
