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

// AddMySQLdExporterReader is a Reader for the AddMySQLdExporter structure.
type AddMySQLdExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddMySQLdExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddMySQLdExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddMySQLdExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddMySQLdExporterOK creates a AddMySQLdExporterOK with default headers values
func NewAddMySQLdExporterOK() *AddMySQLdExporterOK {
	return &AddMySQLdExporterOK{}
}

/*AddMySQLdExporterOK handles this case with default header values.

A successful response.
*/
type AddMySQLdExporterOK struct {
	Payload *AddMySQLdExporterOKBody
}

func (o *AddMySQLdExporterOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddMySQLdExporter][%d] addMySQLdExporterOk  %+v", 200, o.Payload)
}

func (o *AddMySQLdExporterOK) GetPayload() *AddMySQLdExporterOKBody {
	return o.Payload
}

func (o *AddMySQLdExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMySQLdExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddMySQLdExporterDefault creates a AddMySQLdExporterDefault with default headers values
func NewAddMySQLdExporterDefault(code int) *AddMySQLdExporterDefault {
	return &AddMySQLdExporterDefault{
		_statusCode: code,
	}
}

/*AddMySQLdExporterDefault handles this case with default header values.

An unexpected error response.
*/
type AddMySQLdExporterDefault struct {
	_statusCode int

	Payload *AddMySQLdExporterDefaultBody
}

// Code gets the status code for the add my s q ld exporter default response
func (o *AddMySQLdExporterDefault) Code() int {
	return o._statusCode
}

func (o *AddMySQLdExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Agents/AddMySQLdExporter][%d] AddMySQLdExporter default  %+v", o._statusCode, o.Payload)
}

func (o *AddMySQLdExporterDefault) GetPayload() *AddMySQLdExporterDefaultBody {
	return o.Payload
}

func (o *AddMySQLdExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddMySQLdExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddMySQLdExporterBody add my s q ld exporter body
swagger:model AddMySQLdExporterBody
*/
type AddMySQLdExporterBody struct {

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// MySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Certificate Authority certificate chain.
	TLSCa string `json:"tls_ca,omitempty"`

	// Client certificate.
	TLSCert string `json:"tls_cert,omitempty"`

	// Password for decrypting tls_cert.
	TLSKey string `json:"tls_key,omitempty"`

	// Tablestats group collectors will be disabled if there are more than that number of tables.
	// 0 means tablestats group collectors are always enabled (no limit).
	// Negative value means tablestats group collectors are always disabled.
	TablestatsGroupTableLimit int32 `json:"tablestats_group_table_limit,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Enables push metrics mode for exporter.
	PushMetrics bool `json:"push_metrics,omitempty"`

	// List of collector names to disable in this exporter.
	DisableCollectors []string `json:"disable_collectors"`

	// Custom password for exporter endpoint /metrics.
	AgentPassword string `json:"agent_password,omitempty"`
}

// Validate validates this add my s q ld exporter body
func (o *AddMySQLdExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLdExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLdExporterBody) UnmarshalBinary(b []byte) error {
	var res AddMySQLdExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLdExporterDefaultBody add my s q ld exporter default body
swagger:model AddMySQLdExporterDefaultBody
*/
type AddMySQLdExporterDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this add my s q ld exporter default body
func (o *AddMySQLdExporterDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMySQLdExporterDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("AddMySQLdExporter default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLdExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLdExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddMySQLdExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLdExporterOKBody add my s q ld exporter OK body
swagger:model AddMySQLdExporterOKBody
*/
type AddMySQLdExporterOKBody struct {

	// Actual table count at the moment of adding.
	TableCount int32 `json:"table_count,omitempty"`

	// mysqld exporter
	MysqldExporter *AddMySQLdExporterOKBodyMysqldExporter `json:"mysqld_exporter,omitempty"`
}

// Validate validates this add my s q ld exporter OK body
func (o *AddMySQLdExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddMySQLdExporterOKBody) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MysqldExporter) { // not required
		return nil
	}

	if o.MysqldExporter != nil {
		if err := o.MysqldExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addMySQLdExporterOk" + "." + "mysqld_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLdExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLdExporterOKBody) UnmarshalBinary(b []byte) error {
	var res AddMySQLdExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddMySQLdExporterOKBodyMysqldExporter MySQLdExporter runs on Generic or Container Node and exposes MySQL Service metrics.
swagger:model AddMySQLdExporterOKBodyMysqldExporter
*/
type AddMySQLdExporterOKBodyMysqldExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for scraping metrics.
	Username string `json:"username,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// Certificate Authority certificate chain.
	TLSCa string `json:"tls_ca,omitempty"`

	// Client certificate.
	TLSCert string `json:"tls_cert,omitempty"`

	// Password for decrypting tls_cert.
	TLSKey string `json:"tls_key,omitempty"`

	// Tablestats group collectors are disabled if there are more than that number of tables.
	// 0 means tablestats group collectors are always enabled (no limit).
	// Negative value means tablestats group collectors are always disabled.
	TablestatsGroupTableLimit int32 `json:"tablestats_group_table_limit,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// True if exporter uses push metrics mode.
	PushMetricsEnabled bool `json:"push_metrics_enabled,omitempty"`

	// List of disabled collector names.
	DisabledCollectors []string `json:"disabled_collectors"`

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

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// True if tablestats group collectors are currently disabled.
	TablestatsGroupDisabled bool `json:"tablestats_group_disabled,omitempty"`

	// Path to exec process
	ProcessExecPath string `json:"process_exec_path,omitempty"`
}

// Validate validates this add my s q ld exporter OK body mysqld exporter
func (o *AddMySQLdExporterOKBodyMysqldExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum = append(addMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddMySQLdExporterOKBodyMysqldExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddMySQLdExporterOKBodyMysqldExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddMySQLdExporterOKBodyMysqldExporterStatusSTARTING captures enum value "STARTING"
	AddMySQLdExporterOKBodyMysqldExporterStatusSTARTING string = "STARTING"

	// AddMySQLdExporterOKBodyMysqldExporterStatusRUNNING captures enum value "RUNNING"
	AddMySQLdExporterOKBodyMysqldExporterStatusRUNNING string = "RUNNING"

	// AddMySQLdExporterOKBodyMysqldExporterStatusWAITING captures enum value "WAITING"
	AddMySQLdExporterOKBodyMysqldExporterStatusWAITING string = "WAITING"

	// AddMySQLdExporterOKBodyMysqldExporterStatusSTOPPING captures enum value "STOPPING"
	AddMySQLdExporterOKBodyMysqldExporterStatusSTOPPING string = "STOPPING"

	// AddMySQLdExporterOKBodyMysqldExporterStatusDONE captures enum value "DONE"
	AddMySQLdExporterOKBodyMysqldExporterStatusDONE string = "DONE"

	// AddMySQLdExporterOKBodyMysqldExporterStatusUNKNOWN captures enum value "UNKNOWN"
	AddMySQLdExporterOKBodyMysqldExporterStatusUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (o *AddMySQLdExporterOKBodyMysqldExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addMySQLdExporterOkBodyMysqldExporterTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *AddMySQLdExporterOKBodyMysqldExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addMySQLdExporterOk"+"."+"mysqld_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddMySQLdExporterOKBodyMysqldExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMySQLdExporterOKBodyMysqldExporter) UnmarshalBinary(b []byte) error {
	var res AddMySQLdExporterOKBodyMysqldExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
