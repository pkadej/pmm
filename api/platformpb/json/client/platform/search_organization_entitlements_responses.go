// Code generated by go-swagger; DO NOT EDIT.

package platform

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
	"github.com/go-openapi/validate"
)

// SearchOrganizationEntitlementsReader is a Reader for the SearchOrganizationEntitlements structure.
type SearchOrganizationEntitlementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchOrganizationEntitlementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOrganizationEntitlementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchOrganizationEntitlementsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchOrganizationEntitlementsOK creates a SearchOrganizationEntitlementsOK with default headers values
func NewSearchOrganizationEntitlementsOK() *SearchOrganizationEntitlementsOK {
	return &SearchOrganizationEntitlementsOK{}
}

/*SearchOrganizationEntitlementsOK handles this case with default header values.

A successful response.
*/
type SearchOrganizationEntitlementsOK struct {
	Payload *SearchOrganizationEntitlementsOKBody
}

func (o *SearchOrganizationEntitlementsOK) Error() string {
	return fmt.Sprintf("[POST /v1/Platform/SearchOrganizationEntitlements][%d] searchOrganizationEntitlementsOk  %+v", 200, o.Payload)
}

func (o *SearchOrganizationEntitlementsOK) GetPayload() *SearchOrganizationEntitlementsOKBody {
	return o.Payload
}

func (o *SearchOrganizationEntitlementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchOrganizationEntitlementsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchOrganizationEntitlementsDefault creates a SearchOrganizationEntitlementsDefault with default headers values
func NewSearchOrganizationEntitlementsDefault(code int) *SearchOrganizationEntitlementsDefault {
	return &SearchOrganizationEntitlementsDefault{
		_statusCode: code,
	}
}

/*SearchOrganizationEntitlementsDefault handles this case with default header values.

An unexpected error response.
*/
type SearchOrganizationEntitlementsDefault struct {
	_statusCode int

	Payload *SearchOrganizationEntitlementsDefaultBody
}

// Code gets the status code for the search organization entitlements default response
func (o *SearchOrganizationEntitlementsDefault) Code() int {
	return o._statusCode
}

func (o *SearchOrganizationEntitlementsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Platform/SearchOrganizationEntitlements][%d] SearchOrganizationEntitlements default  %+v", o._statusCode, o.Payload)
}

func (o *SearchOrganizationEntitlementsDefault) GetPayload() *SearchOrganizationEntitlementsDefaultBody {
	return o.Payload
}

func (o *SearchOrganizationEntitlementsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SearchOrganizationEntitlementsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*EntitlementsItems0 OrganizationEntitlement contains information about Organization entitlement.
swagger:model EntitlementsItems0
*/
type EntitlementsItems0 struct {

	// Entitlement number.
	Number string `json:"number,omitempty"`

	// Entitlement name.
	Name string `json:"name,omitempty"`

	// Entitlement short summary.
	Summary string `json:"summary,omitempty"`

	// Entitlement tier.
	Tier string `json:"tier,omitempty"`

	// Total units covered by this entitlement.
	TotalUnits string `json:"total_units,omitempty"`

	// Flag indicates that unlimited units are covered.
	UnlimitedUnits bool `json:"unlimited_units,omitempty"`

	// Support level covered by this entitlement.
	SupportLevel string `json:"support_level,omitempty"`

	// Percona product families covered by this entitlement.
	SoftwareFamilies []string `json:"software_families"`

	// Entitlement start data.
	// Note: only date is used here but not time.
	// Format: date-time
	StartDate strfmt.DateTime `json:"start_date,omitempty"`

	// Entitlement end date.
	// Note: only date is used here but not time.
	// Format: date-time
	EndDate strfmt.DateTime `json:"end_date,omitempty"`

	// platform
	Platform *EntitlementsItems0Platform `json:"platform,omitempty"`
}

// Validate validates this entitlements items0
func (o *EntitlementsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *EntitlementsItems0) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(o.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date", "body", "date-time", o.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *EntitlementsItems0) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(o.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("end_date", "body", "date-time", o.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *EntitlementsItems0) validatePlatform(formats strfmt.Registry) error {

	if swag.IsZero(o.Platform) { // not required
		return nil
	}

	if o.Platform != nil {
		if err := o.Platform.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("platform")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *EntitlementsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EntitlementsItems0) UnmarshalBinary(b []byte) error {
	var res EntitlementsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*EntitlementsItems0Platform Platform indicates platform specific entitlements.
swagger:model EntitlementsItems0Platform
*/
type EntitlementsItems0Platform struct {

	// Flag indicates that security advisors are covered by this entitlement.
	SecurityAdvisor string `json:"security_advisor,omitempty"`

	// Flag indicates that config advisors are covered by this entitlement.
	ConfigAdvisor string `json:"config_advisor,omitempty"`
}

// Validate validates this entitlements items0 platform
func (o *EntitlementsItems0Platform) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *EntitlementsItems0Platform) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EntitlementsItems0Platform) UnmarshalBinary(b []byte) error {
	var res EntitlementsItems0Platform
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SearchOrganizationEntitlementsDefaultBody search organization entitlements default body
swagger:model SearchOrganizationEntitlementsDefaultBody
*/
type SearchOrganizationEntitlementsDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this search organization entitlements default body
func (o *SearchOrganizationEntitlementsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationEntitlementsDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("SearchOrganizationEntitlements default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsDefaultBody) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationEntitlementsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SearchOrganizationEntitlementsOKBody search organization entitlements OK body
swagger:model SearchOrganizationEntitlementsOKBody
*/
type SearchOrganizationEntitlementsOKBody struct {

	// entitlements
	Entitlements []*EntitlementsItems0 `json:"entitlements"`
}

// Validate validates this search organization entitlements OK body
func (o *SearchOrganizationEntitlementsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEntitlements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationEntitlementsOKBody) validateEntitlements(formats strfmt.Registry) error {

	if swag.IsZero(o.Entitlements) { // not required
		return nil
	}

	for i := 0; i < len(o.Entitlements); i++ {
		if swag.IsZero(o.Entitlements[i]) { // not required
			continue
		}

		if o.Entitlements[i] != nil {
			if err := o.Entitlements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchOrganizationEntitlementsOk" + "." + "entitlements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationEntitlementsOKBody) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationEntitlementsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
