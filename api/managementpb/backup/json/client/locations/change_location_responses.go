// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// ChangeLocationReader is a Reader for the ChangeLocation structure.
type ChangeLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeLocationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeLocationOK creates a ChangeLocationOK with default headers values
func NewChangeLocationOK() *ChangeLocationOK {
	return &ChangeLocationOK{}
}

/*ChangeLocationOK handles this case with default header values.

A successful response.
*/
type ChangeLocationOK struct {
	Payload interface{}
}

func (o *ChangeLocationOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Locations/Change][%d] changeLocationOk  %+v", 200, o.Payload)
}

func (o *ChangeLocationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeLocationDefault creates a ChangeLocationDefault with default headers values
func NewChangeLocationDefault(code int) *ChangeLocationDefault {
	return &ChangeLocationDefault{
		_statusCode: code,
	}
}

/*ChangeLocationDefault handles this case with default header values.

An unexpected error response.
*/
type ChangeLocationDefault struct {
	_statusCode int

	Payload *ChangeLocationDefaultBody
}

// Code gets the status code for the change location default response
func (o *ChangeLocationDefault) Code() int {
	return o._statusCode
}

func (o *ChangeLocationDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Locations/Change][%d] ChangeLocation default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeLocationDefault) GetPayload() *ChangeLocationDefaultBody {
	return o.Payload
}

func (o *ChangeLocationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeLocationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeLocationBody change location body
swagger:model ChangeLocationBody
*/
type ChangeLocationBody struct {

	// Machine-readable ID.
	LocationID string `json:"location_id,omitempty"`

	// Location name
	Name string `json:"name,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// pmm client config
	PMMClientConfig *ChangeLocationParamsBodyPMMClientConfig `json:"pmm_client_config,omitempty"`

	// pmm server config
	PMMServerConfig *ChangeLocationParamsBodyPMMServerConfig `json:"pmm_server_config,omitempty"`

	// s3 config
	S3Config *ChangeLocationParamsBodyS3Config `json:"s3_config,omitempty"`
}

// Validate validates this change location body
func (o *ChangeLocationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePMMClientConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMMServerConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateS3Config(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeLocationBody) validatePMMClientConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.PMMClientConfig) { // not required
		return nil
	}

	if o.PMMClientConfig != nil {
		if err := o.PMMClientConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pmm_client_config")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeLocationBody) validatePMMServerConfig(formats strfmt.Registry) error {

	if swag.IsZero(o.PMMServerConfig) { // not required
		return nil
	}

	if o.PMMServerConfig != nil {
		if err := o.PMMServerConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "pmm_server_config")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeLocationBody) validateS3Config(formats strfmt.Registry) error {

	if swag.IsZero(o.S3Config) { // not required
		return nil
	}

	if o.S3Config != nil {
		if err := o.S3Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "s3_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeLocationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeLocationBody) UnmarshalBinary(b []byte) error {
	var res ChangeLocationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeLocationDefaultBody change location default body
swagger:model ChangeLocationDefaultBody
*/
type ChangeLocationDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change location default body
func (o *ChangeLocationDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeLocationDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ChangeLocation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeLocationDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeLocationDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeLocationDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeLocationParamsBodyPMMClientConfig PMMClientLocationConfig represents file system config inside pmm-client.
swagger:model ChangeLocationParamsBodyPMMClientConfig
*/
type ChangeLocationParamsBodyPMMClientConfig struct {

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this change location params body PMM client config
func (o *ChangeLocationParamsBodyPMMClientConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeLocationParamsBodyPMMClientConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeLocationParamsBodyPMMClientConfig) UnmarshalBinary(b []byte) error {
	var res ChangeLocationParamsBodyPMMClientConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeLocationParamsBodyPMMServerConfig PMMServerLocationConfig represents file system config inside pmm-server.
swagger:model ChangeLocationParamsBodyPMMServerConfig
*/
type ChangeLocationParamsBodyPMMServerConfig struct {

	// path
	Path string `json:"path,omitempty"`
}

// Validate validates this change location params body PMM server config
func (o *ChangeLocationParamsBodyPMMServerConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeLocationParamsBodyPMMServerConfig) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeLocationParamsBodyPMMServerConfig) UnmarshalBinary(b []byte) error {
	var res ChangeLocationParamsBodyPMMServerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeLocationParamsBodyS3Config S3LocationConfig represents S3 bucket configuration.
swagger:model ChangeLocationParamsBodyS3Config
*/
type ChangeLocationParamsBodyS3Config struct {

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// access key
	AccessKey string `json:"access_key,omitempty"`

	// secret key
	SecretKey string `json:"secret_key,omitempty"`

	// bucket name
	BucketName string `json:"bucket_name,omitempty"`
}

// Validate validates this change location params body s3 config
func (o *ChangeLocationParamsBodyS3Config) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeLocationParamsBodyS3Config) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeLocationParamsBodyS3Config) UnmarshalBinary(b []byte) error {
	var res ChangeLocationParamsBodyS3Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
