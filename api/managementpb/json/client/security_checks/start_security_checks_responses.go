// Code generated by go-swagger; DO NOT EDIT.

package security_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StartSecurityChecksReader is a Reader for the StartSecurityChecks structure.
type StartSecurityChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartSecurityChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartSecurityChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartSecurityChecksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartSecurityChecksOK creates a StartSecurityChecksOK with default headers values
func NewStartSecurityChecksOK() *StartSecurityChecksOK {
	return &StartSecurityChecksOK{}
}

/*StartSecurityChecksOK handles this case with default header values.

A successful response.
*/
type StartSecurityChecksOK struct {
	Payload interface{}
}

func (o *StartSecurityChecksOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/Start][%d] startSecurityChecksOk  %+v", 200, o.Payload)
}

func (o *StartSecurityChecksOK) GetPayload() interface{} {
	return o.Payload
}

func (o *StartSecurityChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartSecurityChecksDefault creates a StartSecurityChecksDefault with default headers values
func NewStartSecurityChecksDefault(code int) *StartSecurityChecksDefault {
	return &StartSecurityChecksDefault{
		_statusCode: code,
	}
}

/*StartSecurityChecksDefault handles this case with default header values.

An unexpected error response.
*/
type StartSecurityChecksDefault struct {
	_statusCode int

	Payload *StartSecurityChecksDefaultBody
}

// Code gets the status code for the start security checks default response
func (o *StartSecurityChecksDefault) Code() int {
	return o._statusCode
}

func (o *StartSecurityChecksDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/Start][%d] StartSecurityChecks default  %+v", o._statusCode, o.Payload)
}

func (o *StartSecurityChecksDefault) GetPayload() *StartSecurityChecksDefaultBody {
	return o.Payload
}

func (o *StartSecurityChecksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartSecurityChecksDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartSecurityChecksBody start security checks body
swagger:model StartSecurityChecksBody
*/
type StartSecurityChecksBody struct {

	// Names of the checks that should be started.
	Names []string `json:"names"`
}

// Validate validates this start security checks body
func (o *StartSecurityChecksBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartSecurityChecksBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartSecurityChecksBody) UnmarshalBinary(b []byte) error {
	var res StartSecurityChecksBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartSecurityChecksDefaultBody start security checks default body
swagger:model StartSecurityChecksDefaultBody
*/
type StartSecurityChecksDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this start security checks default body
func (o *StartSecurityChecksDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartSecurityChecksDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("StartSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartSecurityChecksDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartSecurityChecksDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartSecurityChecksDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
