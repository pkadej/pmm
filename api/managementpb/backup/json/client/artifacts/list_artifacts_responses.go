// Code generated by go-swagger; DO NOT EDIT.

package artifacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListArtifactsReader is a Reader for the ListArtifacts structure.
type ListArtifactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArtifactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArtifactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListArtifactsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListArtifactsOK creates a ListArtifactsOK with default headers values
func NewListArtifactsOK() *ListArtifactsOK {
	return &ListArtifactsOK{}
}

/*ListArtifactsOK handles this case with default header values.

A successful response.
*/
type ListArtifactsOK struct {
	Payload *ListArtifactsOKBody
}

func (o *ListArtifactsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Artifacts/List][%d] listArtifactsOk  %+v", 200, o.Payload)
}

func (o *ListArtifactsOK) GetPayload() *ListArtifactsOKBody {
	return o.Payload
}

func (o *ListArtifactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListArtifactsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArtifactsDefault creates a ListArtifactsDefault with default headers values
func NewListArtifactsDefault(code int) *ListArtifactsDefault {
	return &ListArtifactsDefault{
		_statusCode: code,
	}
}

/*ListArtifactsDefault handles this case with default header values.

An unexpected error response.
*/
type ListArtifactsDefault struct {
	_statusCode int

	Payload *ListArtifactsDefaultBody
}

// Code gets the status code for the list artifacts default response
func (o *ListArtifactsDefault) Code() int {
	return o._statusCode
}

func (o *ListArtifactsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Artifacts/List][%d] ListArtifacts default  %+v", o._statusCode, o.Payload)
}

func (o *ListArtifactsDefault) GetPayload() *ListArtifactsDefaultBody {
	return o.Payload
}

func (o *ListArtifactsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListArtifactsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ArtifactsItems0 Artifact represents single backup artifact.
swagger:model ArtifactsItems0
*/
type ArtifactsItems0 struct {

	// Machine-readable artifact ID.
	ArtifactID string `json:"artifactId,omitempty"`

	// Artifact name
	Name string `json:"name,omitempty"`

	// Database vendor e.g. PostgreSQL, MongoDB, MySQL.
	Vendor string `json:"vendor,omitempty"`

	// Machine-readable location ID.
	LocationID string `json:"locationId,omitempty"`

	// Location name.
	LocationName string `json:"locationName,omitempty"`

	// Machine-readable service ID.
	ServiceID string `json:"serviceId,omitempty"`

	// Service name.
	ServiceName string `json:"serviceName,omitempty"`

	// DataModel is a model used for performing a backup.
	// Enum: [DATA_MODEL_INVALID PHYSICAL LOGICAL]
	DataModel *string `json:"dataModel,omitempty"`

	// BackupStatus shows the current status of execution of backup.
	// Enum: [BACKUP_STATUS_INVALID BACKUP_STATUS_PENDING BACKUP_STATUS_IN_PROGRESS BACKUP_STATUS_PAUSED BACKUP_STATUS_SUCCESS BACKUP_STATUS_ERROR BACKUP_STATUS_DELETING BACKUP_STATUS_FAILED_TO_DELETE]
	Status *string `json:"status,omitempty"`

	// Artifact creation time.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// BackupMode specifies backup mode.
	// Enum: [BACKUP_MODE_INVALID SNAPSHOT INCREMENTAL PITR]
	Mode *string `json:"mode,omitempty"`
}

// Validate validates this artifacts items0
func (o *ArtifactsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDataModel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var artifactsItems0TypeDataModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DATA_MODEL_INVALID","PHYSICAL","LOGICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		artifactsItems0TypeDataModelPropEnum = append(artifactsItems0TypeDataModelPropEnum, v)
	}
}

const (

	// ArtifactsItems0DataModelDATAMODELINVALID captures enum value "DATA_MODEL_INVALID"
	ArtifactsItems0DataModelDATAMODELINVALID string = "DATA_MODEL_INVALID"

	// ArtifactsItems0DataModelPHYSICAL captures enum value "PHYSICAL"
	ArtifactsItems0DataModelPHYSICAL string = "PHYSICAL"

	// ArtifactsItems0DataModelLOGICAL captures enum value "LOGICAL"
	ArtifactsItems0DataModelLOGICAL string = "LOGICAL"
)

// prop value enum
func (o *ArtifactsItems0) validateDataModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, artifactsItems0TypeDataModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ArtifactsItems0) validateDataModel(formats strfmt.Registry) error {

	if swag.IsZero(o.DataModel) { // not required
		return nil
	}

	// value enum
	if err := o.validateDataModelEnum("dataModel", "body", *o.DataModel); err != nil {
		return err
	}

	return nil
}

var artifactsItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BACKUP_STATUS_INVALID","BACKUP_STATUS_PENDING","BACKUP_STATUS_IN_PROGRESS","BACKUP_STATUS_PAUSED","BACKUP_STATUS_SUCCESS","BACKUP_STATUS_ERROR","BACKUP_STATUS_DELETING","BACKUP_STATUS_FAILED_TO_DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		artifactsItems0TypeStatusPropEnum = append(artifactsItems0TypeStatusPropEnum, v)
	}
}

const (

	// ArtifactsItems0StatusBACKUPSTATUSINVALID captures enum value "BACKUP_STATUS_INVALID"
	ArtifactsItems0StatusBACKUPSTATUSINVALID string = "BACKUP_STATUS_INVALID"

	// ArtifactsItems0StatusBACKUPSTATUSPENDING captures enum value "BACKUP_STATUS_PENDING"
	ArtifactsItems0StatusBACKUPSTATUSPENDING string = "BACKUP_STATUS_PENDING"

	// ArtifactsItems0StatusBACKUPSTATUSINPROGRESS captures enum value "BACKUP_STATUS_IN_PROGRESS"
	ArtifactsItems0StatusBACKUPSTATUSINPROGRESS string = "BACKUP_STATUS_IN_PROGRESS"

	// ArtifactsItems0StatusBACKUPSTATUSPAUSED captures enum value "BACKUP_STATUS_PAUSED"
	ArtifactsItems0StatusBACKUPSTATUSPAUSED string = "BACKUP_STATUS_PAUSED"

	// ArtifactsItems0StatusBACKUPSTATUSSUCCESS captures enum value "BACKUP_STATUS_SUCCESS"
	ArtifactsItems0StatusBACKUPSTATUSSUCCESS string = "BACKUP_STATUS_SUCCESS"

	// ArtifactsItems0StatusBACKUPSTATUSERROR captures enum value "BACKUP_STATUS_ERROR"
	ArtifactsItems0StatusBACKUPSTATUSERROR string = "BACKUP_STATUS_ERROR"

	// ArtifactsItems0StatusBACKUPSTATUSDELETING captures enum value "BACKUP_STATUS_DELETING"
	ArtifactsItems0StatusBACKUPSTATUSDELETING string = "BACKUP_STATUS_DELETING"

	// ArtifactsItems0StatusBACKUPSTATUSFAILEDTODELETE captures enum value "BACKUP_STATUS_FAILED_TO_DELETE"
	ArtifactsItems0StatusBACKUPSTATUSFAILEDTODELETE string = "BACKUP_STATUS_FAILED_TO_DELETE"
)

// prop value enum
func (o *ArtifactsItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, artifactsItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ArtifactsItems0) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

func (o *ArtifactsItems0) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var artifactsItems0TypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BACKUP_MODE_INVALID","SNAPSHOT","INCREMENTAL","PITR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		artifactsItems0TypeModePropEnum = append(artifactsItems0TypeModePropEnum, v)
	}
}

const (

	// ArtifactsItems0ModeBACKUPMODEINVALID captures enum value "BACKUP_MODE_INVALID"
	ArtifactsItems0ModeBACKUPMODEINVALID string = "BACKUP_MODE_INVALID"

	// ArtifactsItems0ModeSNAPSHOT captures enum value "SNAPSHOT"
	ArtifactsItems0ModeSNAPSHOT string = "SNAPSHOT"

	// ArtifactsItems0ModeINCREMENTAL captures enum value "INCREMENTAL"
	ArtifactsItems0ModeINCREMENTAL string = "INCREMENTAL"

	// ArtifactsItems0ModePITR captures enum value "PITR"
	ArtifactsItems0ModePITR string = "PITR"
)

// prop value enum
func (o *ArtifactsItems0) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, artifactsItems0TypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ArtifactsItems0) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(o.Mode) { // not required
		return nil
	}

	// value enum
	if err := o.validateModeEnum("mode", "body", *o.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ArtifactsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ArtifactsItems0) UnmarshalBinary(b []byte) error {
	var res ArtifactsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListArtifactsDefaultBody list artifacts default body
swagger:model ListArtifactsDefaultBody
*/
type ListArtifactsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list artifacts default body
func (o *ListArtifactsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactsDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ListArtifacts default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListArtifactsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListArtifactsOKBody list artifacts OK body
swagger:model ListArtifactsOKBody
*/
type ListArtifactsOKBody struct {

	// artifacts
	Artifacts []*ArtifactsItems0 `json:"artifacts"`
}

// Validate validates this list artifacts OK body
func (o *ListArtifactsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateArtifacts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListArtifactsOKBody) validateArtifacts(formats strfmt.Registry) error {

	if swag.IsZero(o.Artifacts) { // not required
		return nil
	}

	for i := 0; i < len(o.Artifacts); i++ {
		if swag.IsZero(o.Artifacts[i]) { // not required
			continue
		}

		if o.Artifacts[i] != nil {
			if err := o.Artifacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listArtifactsOk" + "." + "artifacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListArtifactsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListArtifactsOKBody) UnmarshalBinary(b []byte) error {
	var res ListArtifactsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
