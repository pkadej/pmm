// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// AddQANPostgreSQLPgStatMonitorAgentReader is a Reader for the AddQANPostgreSQLPgStatMonitorAgent structure.
type AddQANPostgreSQLPgStatMonitorAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddQANPostgreSQLPgStatMonitorAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddQANPostgreSQLPgStatMonitorAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddQANPostgreSQLPgStatMonitorAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddQANPostgreSQLPgStatMonitorAgentOK creates a AddQANPostgreSQLPgStatMonitorAgentOK with default headers values
func NewAddQANPostgreSQLPgStatMonitorAgentOK() *AddQANPostgreSQLPgStatMonitorAgentOK {
	return &AddQANPostgreSQLPgStatMonitorAgentOK{}
}

/*AddQANPostgreSQLPgStatMonitorAgentOK handles this case with default header values.

A successful response.
*/
type AddQANPostgreSQLPgStatMonitorAgentOK struct {
	Payload *AddQANPostgreSQLPgStatMonitorAgentOKBody
}

func (o *AddQANPostgreSQLPgStatMonitorAgentOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddQANPostgreSQLPgStatMonitorAgent][%d] addQanPostgreSqlPgStatMonitorAgentOk  %+v", 200, o.Payload)
}

func (o *AddQANPostgreSQLPgStatMonitorAgentOK) GetPayload() *AddQANPostgreSQLPgStatMonitorAgentOKBody {
	return o.Payload
}

func (o *AddQANPostgreSQLPgStatMonitorAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddQANPostgreSQLPgStatMonitorAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddQANPostgreSQLPgStatMonitorAgentDefault creates a AddQANPostgreSQLPgStatMonitorAgentDefault with default headers values
func NewAddQANPostgreSQLPgStatMonitorAgentDefault(code int) *AddQANPostgreSQLPgStatMonitorAgentDefault {
	return &AddQANPostgreSQLPgStatMonitorAgentDefault{
		_statusCode: code,
	}
}

/*AddQANPostgreSQLPgStatMonitorAgentDefault handles this case with default header values.

An unexpected error response.
*/
type AddQANPostgreSQLPgStatMonitorAgentDefault struct {
	_statusCode int

	Payload *AddQANPostgreSQLPgStatMonitorAgentDefaultBody
}

// Code gets the status code for the add QAN postgre SQL pg stat monitor agent default response
func (o *AddQANPostgreSQLPgStatMonitorAgentDefault) Code() int {
	return o._statusCode
}

func (o *AddQANPostgreSQLPgStatMonitorAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddQANPostgreSQLPgStatMonitorAgent][%d] AddQANPostgreSQLPgStatMonitorAgent default  %+v", o._statusCode, o.Payload)
}

func (o *AddQANPostgreSQLPgStatMonitorAgentDefault) GetPayload() *AddQANPostgreSQLPgStatMonitorAgentDefaultBody {
	return o.Payload
}

func (o *AddQANPostgreSQLPgStatMonitorAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddQANPostgreSQLPgStatMonitorAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddQANPostgreSQLPgStatMonitorAgentBody add QAN postgre SQL pg stat monitor agent body
swagger:model AddQANPostgreSQLPgStatMonitorAgentBody
*/
type AddQANPostgreSQLPgStatMonitorAgentBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmmAgentId,omitempty"`

	// Service identifier.
	ServiceID string `json:"serviceId,omitempty"`

	// PostgreSQL username for getting pg stat monitor data.
	Username string `json:"username,omitempty"`

	// PostgreSQL password for getting pg stat monitor data.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tlsSkipVerify,omitempty"`

	// Disable query examples.
	DisableQueryExamples bool `json:"disableQueryExamples,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"customLabels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skipConnectionCheck,omitempty"`

	// TLS CA certificate.
	TLSCa string `json:"tlsCa,omitempty"`

	// TLS Certifcate.
	TLSCert string `json:"tlsCert,omitempty"`

	// TLS Certificate Key.
	TLSKey string `json:"tlsKey,omitempty"`
}

// Validate validates this add QAN postgre SQL pg stat monitor agent body
func (o *AddQANPostgreSQLPgStatMonitorAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatMonitorAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatMonitorAgentBody) UnmarshalBinary(b []byte) error {
	var res AddQANPostgreSQLPgStatMonitorAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANPostgreSQLPgStatMonitorAgentDefaultBody add QAN postgre SQL pg stat monitor agent default body
swagger:model AddQANPostgreSQLPgStatMonitorAgentDefaultBody
*/
type AddQANPostgreSQLPgStatMonitorAgentDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add QAN postgre SQL pg stat monitor agent default body
func (o *AddQANPostgreSQLPgStatMonitorAgentDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANPostgreSQLPgStatMonitorAgentDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddQANPostgreSQLPgStatMonitorAgent default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatMonitorAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatMonitorAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddQANPostgreSQLPgStatMonitorAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANPostgreSQLPgStatMonitorAgentOKBody add QAN postgre SQL pg stat monitor agent OK body
swagger:model AddQANPostgreSQLPgStatMonitorAgentOKBody
*/
type AddQANPostgreSQLPgStatMonitorAgentOKBody struct {

	// qan postgresql pgstatmonitor agent
	QANPostgresqlPgstatmonitorAgent *AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent `json:"qanPostgresqlPgstatmonitorAgent,omitempty"`
}

// Validate validates this add QAN postgre SQL pg stat monitor agent OK body
func (o *AddQANPostgreSQLPgStatMonitorAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateQANPostgresqlPgstatmonitorAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddQANPostgreSQLPgStatMonitorAgentOKBody) validateQANPostgresqlPgstatmonitorAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANPostgresqlPgstatmonitorAgent) { // not required
		return nil
	}

	if o.QANPostgresqlPgstatmonitorAgent != nil {
		if err := o.QANPostgresqlPgstatmonitorAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addQanPostgreSqlPgStatMonitorAgentOk" + "." + "qanPostgresqlPgstatmonitorAgent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatMonitorAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatMonitorAgentOKBody) UnmarshalBinary(b []byte) error {
	var res AddQANPostgreSQLPgStatMonitorAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent QANPostgreSQLPgStatMonitorAgent runs within pmm-agent and sends PostgreSQL Query Analytics data to the PMM Server.
swagger:model AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent
*/
type AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agentId,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmmAgentId,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"serviceId,omitempty"`

	// PostgreSQL username for getting pg stat monitor data.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tlsSkipVerify,omitempty"`

	// True if query examples are disabled.
	QueryExamplesDisabled bool `json:"queryExamplesDisabled,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"customLabels,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	//  - UNKNOWN: Agent is not connected, we don't know anything about it's state.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE UNKNOWN]
	Status *string `json:"status,omitempty"`
}

// Validate validates this add QAN postgre SQL pg stat monitor agent OK body QAN postgresql pgstatmonitor agent
func (o *AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum = append(addQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum, v)
	}
}

const (

	// AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusSTARTING captures enum value "STARTING"
	AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusSTARTING string = "STARTING"

	// AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusRUNNING captures enum value "RUNNING"
	AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusRUNNING string = "RUNNING"

	// AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusWAITING captures enum value "WAITING"
	AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusWAITING string = "WAITING"

	// AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusSTOPPING captures enum value "STOPPING"
	AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusSTOPPING string = "STOPPING"

	// AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusDONE captures enum value "DONE"
	AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusDONE string = "DONE"

	// AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusUNKNOWN captures enum value "UNKNOWN"
	AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgentStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addQanPostgreSqlPgStatMonitorAgentOkBodyQanPostgresqlPgstatmonitorAgentTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addQanPostgreSqlPgStatMonitorAgentOk"+"."+"qanPostgresqlPgstatmonitorAgent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent) UnmarshalBinary(b []byte) error {
	var res AddQANPostgreSQLPgStatMonitorAgentOKBodyQANPostgresqlPgstatmonitorAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
