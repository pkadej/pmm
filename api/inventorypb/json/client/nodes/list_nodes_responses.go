// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// ListNodesReader is a Reader for the ListNodes structure.
type ListNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListNodesOK creates a ListNodesOK with default headers values
func NewListNodesOK() *ListNodesOK {
	return &ListNodesOK{}
}

/*ListNodesOK handles this case with default header values.

A successful response.
*/
type ListNodesOK struct {
	Payload *ListNodesOKBody
}

func (o *ListNodesOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/List][%d] listNodesOk  %+v", 200, o.Payload)
}

func (o *ListNodesOK) GetPayload() *ListNodesOKBody {
	return o.Payload
}

func (o *ListNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListNodesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNodesDefault creates a ListNodesDefault with default headers values
func NewListNodesDefault(code int) *ListNodesDefault {
	return &ListNodesDefault{
		_statusCode: code,
	}
}

/*ListNodesDefault handles this case with default header values.

An unexpected error response.
*/
type ListNodesDefault struct {
	_statusCode int

	Payload *ListNodesDefaultBody
}

// Code gets the status code for the list nodes default response
func (o *ListNodesDefault) Code() int {
	return o._statusCode
}

func (o *ListNodesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/List][%d] ListNodes default  %+v", o._statusCode, o.Payload)
}

func (o *ListNodesDefault) GetPayload() *ListNodesDefaultBody {
	return o.Payload
}

func (o *ListNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListNodesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerItems0 ContainerNode represents a Docker container.
swagger:model ContainerItems0
*/
type ContainerItems0 struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs.
	MachineID string `json:"machine_id,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this container items0
func (o *ContainerItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerItems0) UnmarshalBinary(b []byte) error {
	var res ContainerItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GenericItems0 GenericNode represents a bare metal server or virtual machine.
swagger:model GenericItems0
*/
type GenericItems0 struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this generic items0
func (o *GenericItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GenericItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GenericItems0) UnmarshalBinary(b []byte) error {
	var res GenericItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListNodesBody list nodes body
swagger:model ListNodesBody
*/
type ListNodesBody struct {

	// NodeType describes supported Node types.
	// Enum: [NODE_TYPE_INVALID GENERIC_NODE CONTAINER_NODE REMOTE_NODE REMOTE_RDS_NODE REMOTE_AZURE_DATABASE_NODE]
	NodeType *string `json:"node_type,omitempty"`
}

// Validate validates this list nodes body
func (o *ListNodesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listNodesBodyTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_TYPE_INVALID","GENERIC_NODE","CONTAINER_NODE","REMOTE_NODE","REMOTE_RDS_NODE","REMOTE_AZURE_DATABASE_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listNodesBodyTypeNodeTypePropEnum = append(listNodesBodyTypeNodeTypePropEnum, v)
	}
}

const (

	// ListNodesBodyNodeTypeNODETYPEINVALID captures enum value "NODE_TYPE_INVALID"
	ListNodesBodyNodeTypeNODETYPEINVALID string = "NODE_TYPE_INVALID"

	// ListNodesBodyNodeTypeGENERICNODE captures enum value "GENERIC_NODE"
	ListNodesBodyNodeTypeGENERICNODE string = "GENERIC_NODE"

	// ListNodesBodyNodeTypeCONTAINERNODE captures enum value "CONTAINER_NODE"
	ListNodesBodyNodeTypeCONTAINERNODE string = "CONTAINER_NODE"

	// ListNodesBodyNodeTypeREMOTENODE captures enum value "REMOTE_NODE"
	ListNodesBodyNodeTypeREMOTENODE string = "REMOTE_NODE"

	// ListNodesBodyNodeTypeREMOTERDSNODE captures enum value "REMOTE_RDS_NODE"
	ListNodesBodyNodeTypeREMOTERDSNODE string = "REMOTE_RDS_NODE"

	// ListNodesBodyNodeTypeREMOTEAZUREDATABASENODE captures enum value "REMOTE_AZURE_DATABASE_NODE"
	ListNodesBodyNodeTypeREMOTEAZUREDATABASENODE string = "REMOTE_AZURE_DATABASE_NODE"
)

// prop value enum
func (o *ListNodesBody) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listNodesBodyTypeNodeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListNodesBody) validateNodeType(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeType) { // not required
		return nil
	}

	// value enum
	if err := o.validateNodeTypeEnum("body"+"."+"node_type", "body", *o.NodeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListNodesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListNodesBody) UnmarshalBinary(b []byte) error {
	var res ListNodesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListNodesDefaultBody list nodes default body
swagger:model ListNodesDefaultBody
*/
type ListNodesDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list nodes default body
func (o *ListNodesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListNodesDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ListNodes default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListNodesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListNodesDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListNodesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListNodesOKBody list nodes OK body
swagger:model ListNodesOKBody
*/
type ListNodesOKBody struct {

	// generic
	Generic []*GenericItems0 `json:"generic"`

	// container
	Container []*ContainerItems0 `json:"container"`

	// remote
	Remote []*RemoteItems0 `json:"remote"`

	// remote rds
	RemoteRDS []*RemoteRDSItems0 `json:"remote_rds"`

	// remote azure database
	RemoteAzureDatabase []*RemoteAzureDatabaseItems0 `json:"remote_azure_database"`
}

// Validate validates this list nodes OK body
func (o *ListNodesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteRDS(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteAzureDatabase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListNodesOKBody) validateGeneric(formats strfmt.Registry) error {

	if swag.IsZero(o.Generic) { // not required
		return nil
	}

	for i := 0; i < len(o.Generic); i++ {
		if swag.IsZero(o.Generic[i]) { // not required
			continue
		}

		if o.Generic[i] != nil {
			if err := o.Generic[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listNodesOk" + "." + "generic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListNodesOKBody) validateContainer(formats strfmt.Registry) error {

	if swag.IsZero(o.Container) { // not required
		return nil
	}

	for i := 0; i < len(o.Container); i++ {
		if swag.IsZero(o.Container[i]) { // not required
			continue
		}

		if o.Container[i] != nil {
			if err := o.Container[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listNodesOk" + "." + "container" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListNodesOKBody) validateRemote(formats strfmt.Registry) error {

	if swag.IsZero(o.Remote) { // not required
		return nil
	}

	for i := 0; i < len(o.Remote); i++ {
		if swag.IsZero(o.Remote[i]) { // not required
			continue
		}

		if o.Remote[i] != nil {
			if err := o.Remote[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listNodesOk" + "." + "remote" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListNodesOKBody) validateRemoteRDS(formats strfmt.Registry) error {

	if swag.IsZero(o.RemoteRDS) { // not required
		return nil
	}

	for i := 0; i < len(o.RemoteRDS); i++ {
		if swag.IsZero(o.RemoteRDS[i]) { // not required
			continue
		}

		if o.RemoteRDS[i] != nil {
			if err := o.RemoteRDS[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listNodesOk" + "." + "remote_rds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListNodesOKBody) validateRemoteAzureDatabase(formats strfmt.Registry) error {

	if swag.IsZero(o.RemoteAzureDatabase) { // not required
		return nil
	}

	for i := 0; i < len(o.RemoteAzureDatabase); i++ {
		if swag.IsZero(o.RemoteAzureDatabase[i]) { // not required
			continue
		}

		if o.RemoteAzureDatabase[i] != nil {
			if err := o.RemoteAzureDatabase[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listNodesOk" + "." + "remote_azure_database" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListNodesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListNodesOKBody) UnmarshalBinary(b []byte) error {
	var res ListNodesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoteAzureDatabaseItems0 RemoteAzureDatabaseNode represents remote AzureDatabase Node. Agents can't run on Remote AzureDatabase Nodes.
swagger:model RemoteAzureDatabaseItems0
*/
type RemoteAzureDatabaseItems0 struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this remote azure database items0
func (o *RemoteAzureDatabaseItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoteAzureDatabaseItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoteAzureDatabaseItems0) UnmarshalBinary(b []byte) error {
	var res RemoteAzureDatabaseItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoteItems0 RemoteNode represents generic remote Node. It's a node where we don't run pmm-agents. Only external exporters can run on Remote Nodes.
swagger:model RemoteItems0
*/
type RemoteItems0 struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// Node address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this remote items0
func (o *RemoteItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoteItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoteItems0) UnmarshalBinary(b []byte) error {
	var res RemoteItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoteRDSItems0 RemoteRDSNode represents remote RDS Node. Agents can't run on Remote RDS Nodes.
swagger:model RemoteRDSItems0
*/
type RemoteRDSItems0 struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// DB instance identifier.
	Address string `json:"address,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this remote RDS items0
func (o *RemoteRDSItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoteRDSItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoteRDSItems0) UnmarshalBinary(b []byte) error {
	var res RemoteRDSItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
