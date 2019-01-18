// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAgentReader is a Reader for the GetAgent structure.
type GetAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAgentOK creates a GetAgentOK with default headers values
func NewGetAgentOK() *GetAgentOK {
	return &GetAgentOK{}
}

/*GetAgentOK handles this case with default header values.

(empty)
*/
type GetAgentOK struct {
	Payload *GetAgentOKBody
}

func (o *GetAgentOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/Get][%d] getAgentOK  %+v", 200, o.Payload)
}

func (o *GetAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAgentBody get agent body
swagger:model GetAgentBody
*/
type GetAgentBody struct {

	// Unique Agent identifier.
	ID string `json:"id,omitempty"`
}

// Validate validates this get agent body
func (o *GetAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentBody) UnmarshalBinary(b []byte) error {
	var res GetAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBody get agent o k body
swagger:model GetAgentOKBody
*/
type GetAgentOKBody struct {

	// mysqld exporter
	MysqldExporter *GetAgentOKBodyMysqldExporter `json:"mysqld_exporter,omitempty"`

	// node exporter
	NodeExporter *GetAgentOKBodyNodeExporter `json:"node_exporter,omitempty"`

	// pmm agent
	PMMAgent *GetAgentOKBodyPMMAgent `json:"pmm_agent,omitempty"`
}

// Validate validates this get agent o k body
func (o *GetAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMysqldExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNodeExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePMMAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAgentOKBody) validateMysqldExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MysqldExporter) { // not required
		return nil
	}

	if o.MysqldExporter != nil {
		if err := o.MysqldExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "mysqld_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *GetAgentOKBody) validateNodeExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeExporter) { // not required
		return nil
	}

	if o.NodeExporter != nil {
		if err := o.NodeExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "node_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *GetAgentOKBody) validatePMMAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.PMMAgent) { // not required
		return nil
	}

	if o.PMMAgent != nil {
		if err := o.PMMAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "pmm_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBody) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyMysqldExporter MySQLdExporter represents mysqld_exporter Agent configuration.
swagger:model GetAgentOKBodyMysqldExporter
*/
type GetAgentOKBodyMysqldExporter struct {

	// Agent desired status: enabled or disabled.
	Disabled bool `json:"disabled,omitempty"`

	// host node info
	HostNodeInfo *GetAgentOKBodyMysqldExporterHostNodeInfo `json:"host_node_info,omitempty"`

	// Unique Agent identifier.
	ID string `json:"id,omitempty"`

	// HTTP listen port for exposing metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// Agent process status: running or not.
	Running bool `json:"running,omitempty"`

	// Service identifier for which insights are provided by that Agent.
	ServiceID string `json:"service_id,omitempty"`

	// MySQL username for extracting metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this get agent o k body mysqld exporter
func (o *GetAgentOKBodyMysqldExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHostNodeInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAgentOKBodyMysqldExporter) validateHostNodeInfo(formats strfmt.Registry) error {

	if swag.IsZero(o.HostNodeInfo) { // not required
		return nil
	}

	if o.HostNodeInfo != nil {
		if err := o.HostNodeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "mysqld_exporter" + "." + "host_node_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyMysqldExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyMysqldExporter) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyMysqldExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyMysqldExporterHostNodeInfo HostNodeInfo describes the way Service or Agent runs on Node.
swagger:model GetAgentOKBodyMysqldExporterHostNodeInfo
*/
type GetAgentOKBodyMysqldExporterHostNodeInfo struct {

	// Docker container ID.
	ContainerID string `json:"container_id,omitempty"`

	// Docker container name.
	ContainerName string `json:"container_name,omitempty"`

	// Kubernetes pod name.
	KubernetesPodName string `json:"kubernetes_pod_name,omitempty"`

	// Kubernetes pod UID.
	KubernetesPodUID string `json:"kubernetes_pod_uid,omitempty"`

	// Node identifier where Service or Agent runs.
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this get agent o k body mysqld exporter host node info
func (o *GetAgentOKBodyMysqldExporterHostNodeInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyMysqldExporterHostNodeInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyMysqldExporterHostNodeInfo) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyMysqldExporterHostNodeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyNodeExporter NodeExporter represents node_exporter Agent configuration.
swagger:model GetAgentOKBodyNodeExporter
*/
type GetAgentOKBodyNodeExporter struct {

	// Agent desired status: enabled or disabled.
	Disabled bool `json:"disabled,omitempty"`

	// host node info
	HostNodeInfo *GetAgentOKBodyNodeExporterHostNodeInfo `json:"host_node_info,omitempty"`

	// Unique Agent identifier.
	ID string `json:"id,omitempty"`

	// HTTP listen port for exposing metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// Agent process status: running or not.
	Running bool `json:"running,omitempty"`
}

// Validate validates this get agent o k body node exporter
func (o *GetAgentOKBodyNodeExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHostNodeInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAgentOKBodyNodeExporter) validateHostNodeInfo(formats strfmt.Registry) error {

	if swag.IsZero(o.HostNodeInfo) { // not required
		return nil
	}

	if o.HostNodeInfo != nil {
		if err := o.HostNodeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "node_exporter" + "." + "host_node_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyNodeExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyNodeExporter) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyNodeExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyNodeExporterHostNodeInfo HostNodeInfo describes the way Service or Agent runs on Node.
swagger:model GetAgentOKBodyNodeExporterHostNodeInfo
*/
type GetAgentOKBodyNodeExporterHostNodeInfo struct {

	// Docker container ID.
	ContainerID string `json:"container_id,omitempty"`

	// Docker container name.
	ContainerName string `json:"container_name,omitempty"`

	// Kubernetes pod name.
	KubernetesPodName string `json:"kubernetes_pod_name,omitempty"`

	// Kubernetes pod UID.
	KubernetesPodUID string `json:"kubernetes_pod_uid,omitempty"`

	// Node identifier where Service or Agent runs.
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this get agent o k body node exporter host node info
func (o *GetAgentOKBodyNodeExporterHostNodeInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyNodeExporterHostNodeInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyNodeExporterHostNodeInfo) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyNodeExporterHostNodeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyPMMAgent PMMAgent represent pmm-agent Agent configuration.
swagger:model GetAgentOKBodyPMMAgent
*/
type GetAgentOKBodyPMMAgent struct {

	// host node info
	HostNodeInfo *GetAgentOKBodyPMMAgentHostNodeInfo `json:"host_node_info,omitempty"`

	// Unique Agent identifier.
	ID string `json:"id,omitempty"`

	// Agent process status: running and connected to pmm-managed, or not.
	Running bool `json:"running,omitempty"`
}

// Validate validates this get agent o k body PMM agent
func (o *GetAgentOKBodyPMMAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHostNodeInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAgentOKBodyPMMAgent) validateHostNodeInfo(formats strfmt.Registry) error {

	if swag.IsZero(o.HostNodeInfo) { // not required
		return nil
	}

	if o.HostNodeInfo != nil {
		if err := o.HostNodeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getAgentOK" + "." + "pmm_agent" + "." + "host_node_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyPMMAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyPMMAgent) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyPMMAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAgentOKBodyPMMAgentHostNodeInfo HostNodeInfo describes the way Service or Agent runs on Node.
swagger:model GetAgentOKBodyPMMAgentHostNodeInfo
*/
type GetAgentOKBodyPMMAgentHostNodeInfo struct {

	// Docker container ID.
	ContainerID string `json:"container_id,omitempty"`

	// Docker container name.
	ContainerName string `json:"container_name,omitempty"`

	// Kubernetes pod name.
	KubernetesPodName string `json:"kubernetes_pod_name,omitempty"`

	// Kubernetes pod UID.
	KubernetesPodUID string `json:"kubernetes_pod_uid,omitempty"`

	// Node identifier where Service or Agent runs.
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this get agent o k body PMM agent host node info
func (o *GetAgentOKBodyPMMAgentHostNodeInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAgentOKBodyPMMAgentHostNodeInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAgentOKBodyPMMAgentHostNodeInfo) UnmarshalBinary(b []byte) error {
	var res GetAgentOKBodyPMMAgentHostNodeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}