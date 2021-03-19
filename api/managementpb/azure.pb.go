// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: managementpb/azure.proto

package managementpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AddAzureDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Azure region.
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// AWS availability zone.
	Az string `protobuf:"bytes,2,opt,name=az,proto3" json:"az,omitempty"`
	// AWS instance ID.
	InstanceId string `protobuf:"bytes,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// Azure instance class.
	NodeModel string `protobuf:"bytes,4,opt,name=node_model,json=nodeModel,proto3" json:"node_model,omitempty"`
	// Address used to connect to it.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Access port.
	Port uint32 `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	// Unique across all Nodes user-defined name. Defaults to Azure Database instance ID.
	NodeName string `protobuf:"bytes,8,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Unique across all Services user-defined name. Defaults to Azure Database instance ID.
	ServiceName string `protobuf:"bytes,9,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,10,opt,name=environment,proto3" json:"environment,omitempty"`
	// Username for scraping metrics.
	Username string `protobuf:"bytes,13,opt,name=username,proto3" json:"username,omitempty"`
	// Password for scraping metrics.
	Password string `protobuf:"bytes,14,opt,name=password,proto3" json:"password,omitempty"`
	// Azure database client ID
	AzureDatabaseClientId string `protobuf:"bytes,29,opt,name=azure_database_client_id,json=azureDatabaseClientId,proto3" json:"azure_database_client_id,omitempty"`
	// Azure database client secret
	AzureDatabaseClientSecret string `protobuf:"bytes,30,opt,name=azure_database_client_secret,json=azureDatabaseClientSecret,proto3" json:"azure_database_client_secret,omitempty"`
	// Azure database tanant ID
	AzureDatabaseTenantId string `protobuf:"bytes,31,opt,name=azure_database_tenant_id,json=azureDatabaseTenantId,proto3" json:"azure_database_tenant_id,omitempty"`
	// Azure database subscription ID
	AzureDatabaseSubscriptionId string `protobuf:"bytes,32,opt,name=azure_database_subscription_id,json=azureDatabaseSubscriptionId,proto3" json:"azure_database_subscription_id,omitempty"`
	// Azure database resource type (mysql, maria, postgres)
	AzureDatabaseResourceType string `protobuf:"bytes,33,opt,name=azure_database_resource_type,json=azureDatabaseResourceType,proto3" json:"azure_database_resource_type,omitempty"`
	// If true, adds azure_database_exporter.
	AzureDatabaseExporter bool `protobuf:"varint,17,opt,name=azure_database_exporter,json=azureDatabaseExporter,proto3" json:"azure_database_exporter,omitempty"`
	// If true, adds qan-mysql-perfschema-agent.
	QanMysqlPerfschema bool `protobuf:"varint,18,opt,name=qan_mysql_perfschema,json=qanMysqlPerfschema,proto3" json:"qan_mysql_perfschema,omitempty"`
	// Custom user-assigned labels for Node and Service.
	CustomLabels map[string]string `protobuf:"bytes,19,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Skip connection check.
	SkipConnectionCheck bool `protobuf:"varint,20,opt,name=skip_connection_check,json=skipConnectionCheck,proto3" json:"skip_connection_check,omitempty"`
	// Use TLS for database connections.
	Tls bool `protobuf:"varint,21,opt,name=tls,proto3" json:"tls,omitempty"`
	// Skip TLS certificate and hostname validation.
	TlsSkipVerify bool `protobuf:"varint,22,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty"`
	// Disable query examples.
	DisableQueryExamples bool `protobuf:"varint,23,opt,name=disable_query_examples,json=disableQueryExamples,proto3" json:"disable_query_examples,omitempty"`
	// Tablestats group collectors will be disabled if there are more than that number of tables.
	// If zero, server's default value is used.
	// Use negative value to disable them.
	TablestatsGroupTableLimit int32 `protobuf:"varint,24,opt,name=tablestats_group_table_limit,json=tablestatsGroupTableLimit,proto3" json:"tablestats_group_table_limit,omitempty"`
	// Disable basic metrics.
	DisableBasicMetrics bool `protobuf:"varint,25,opt,name=disable_basic_metrics,json=disableBasicMetrics,proto3" json:"disable_basic_metrics,omitempty"`
	// Disable enhanced metrics.
	DisableEnhancedMetrics bool `protobuf:"varint,26,opt,name=disable_enhanced_metrics,json=disableEnhancedMetrics,proto3" json:"disable_enhanced_metrics,omitempty"`
	// Defines metrics flow model for this exporter.
	// Push metrics mode is not allowed.
	MetricsMode MetricsMode `protobuf:"varint,27,opt,name=metrics_mode,json=metricsMode,proto3,enum=management.MetricsMode" json:"metrics_mode,omitempty"`
}

func (x *AddAzureDatabaseRequest) Reset() {
	*x = AddAzureDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_azure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAzureDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAzureDatabaseRequest) ProtoMessage() {}

func (x *AddAzureDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_azure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAzureDatabaseRequest.ProtoReflect.Descriptor instead.
func (*AddAzureDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_azure_proto_rawDescGZIP(), []int{0}
}

func (x *AddAzureDatabaseRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAz() string {
	if x != nil {
		return x.Az
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetNodeModel() string {
	if x != nil {
		return x.NodeModel
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *AddAzureDatabaseRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureDatabaseClientId() string {
	if x != nil {
		return x.AzureDatabaseClientId
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureDatabaseClientSecret() string {
	if x != nil {
		return x.AzureDatabaseClientSecret
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureDatabaseTenantId() string {
	if x != nil {
		return x.AzureDatabaseTenantId
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureDatabaseSubscriptionId() string {
	if x != nil {
		return x.AzureDatabaseSubscriptionId
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureDatabaseResourceType() string {
	if x != nil {
		return x.AzureDatabaseResourceType
	}
	return ""
}

func (x *AddAzureDatabaseRequest) GetAzureDatabaseExporter() bool {
	if x != nil {
		return x.AzureDatabaseExporter
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetQanMysqlPerfschema() bool {
	if x != nil {
		return x.QanMysqlPerfschema
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

func (x *AddAzureDatabaseRequest) GetSkipConnectionCheck() bool {
	if x != nil {
		return x.SkipConnectionCheck
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetTls() bool {
	if x != nil {
		return x.Tls
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetTlsSkipVerify() bool {
	if x != nil {
		return x.TlsSkipVerify
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetDisableQueryExamples() bool {
	if x != nil {
		return x.DisableQueryExamples
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetTablestatsGroupTableLimit() int32 {
	if x != nil {
		return x.TablestatsGroupTableLimit
	}
	return 0
}

func (x *AddAzureDatabaseRequest) GetDisableBasicMetrics() bool {
	if x != nil {
		return x.DisableBasicMetrics
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetDisableEnhancedMetrics() bool {
	if x != nil {
		return x.DisableEnhancedMetrics
	}
	return false
}

func (x *AddAzureDatabaseRequest) GetMetricsMode() MetricsMode {
	if x != nil {
		return x.MetricsMode
	}
	return MetricsMode_AUTO
}

type AddAzureDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddAzureDatabaseResponse) Reset() {
	*x = AddAzureDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_azure_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAzureDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAzureDatabaseResponse) ProtoMessage() {}

func (x *AddAzureDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_azure_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAzureDatabaseResponse.ProtoReflect.Descriptor instead.
func (*AddAzureDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_azure_proto_rawDescGZIP(), []int{1}
}

var File_managementpb_azure_proto protoreflect.FileDescriptor

var file_managementpb_azure_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x61,
	0x7a, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x0a, 0x0a, 0x17, 0x41, 0x64, 0x64,
	0x41, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x7a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x61, 0x7a, 0x12, 0x27, 0x0a, 0x0b, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x10, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x37, 0x0a, 0x18,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x1c, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x43, 0x0a, 0x1e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x1c, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x61, 0x7a, 0x75, 0x72,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x30, 0x0a,
	0x14, 0x71, 0x61, 0x6e, 0x5f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x71, 0x61, 0x6e,
	0x4d, 0x79, 0x73, 0x71, 0x6c, 0x50, 0x65, 0x72, 0x66, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x5a, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x73,
	0x6b, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73, 0x6b, 0x69, 0x70,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x6c,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x53,
	0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12,
	0x3f, 0x0a, 0x1c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x32, 0x0a, 0x15, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x69, 0x63, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x45,
	0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x3a,
	0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1a, 0x0a, 0x18, 0x41,
	0x64, 0x64, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9c, 0x01, 0x0a, 0x0d, 0x41, 0x7a, 0x75, 0x72,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x8a, 0x01, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x23,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41,
	0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x64, 0x64, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x25, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x41, 0x64, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_azure_proto_rawDescOnce sync.Once
	file_managementpb_azure_proto_rawDescData = file_managementpb_azure_proto_rawDesc
)

func file_managementpb_azure_proto_rawDescGZIP() []byte {
	file_managementpb_azure_proto_rawDescOnce.Do(func() {
		file_managementpb_azure_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_azure_proto_rawDescData)
	})
	return file_managementpb_azure_proto_rawDescData
}

var file_managementpb_azure_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_managementpb_azure_proto_goTypes = []interface{}{
	(*AddAzureDatabaseRequest)(nil),  // 0: management.AddAzureDatabaseRequest
	(*AddAzureDatabaseResponse)(nil), // 1: management.AddAzureDatabaseResponse
	nil,                              // 2: management.AddAzureDatabaseRequest.CustomLabelsEntry
	(MetricsMode)(0),                 // 3: management.MetricsMode
}
var file_managementpb_azure_proto_depIdxs = []int32{
	2, // 0: management.AddAzureDatabaseRequest.custom_labels:type_name -> management.AddAzureDatabaseRequest.CustomLabelsEntry
	3, // 1: management.AddAzureDatabaseRequest.metrics_mode:type_name -> management.MetricsMode
	0, // 2: management.AzureDatabase.AddAzureDatabase:input_type -> management.AddAzureDatabaseRequest
	1, // 3: management.AzureDatabase.AddAzureDatabase:output_type -> management.AddAzureDatabaseResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_managementpb_azure_proto_init() }
func file_managementpb_azure_proto_init() {
	if File_managementpb_azure_proto != nil {
		return
	}
	file_managementpb_metrics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_managementpb_azure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAzureDatabaseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_azure_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAzureDatabaseResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_azure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_azure_proto_goTypes,
		DependencyIndexes: file_managementpb_azure_proto_depIdxs,
		MessageInfos:      file_managementpb_azure_proto_msgTypes,
	}.Build()
	File_managementpb_azure_proto = out.File
	file_managementpb_azure_proto_rawDesc = nil
	file_managementpb_azure_proto_goTypes = nil
	file_managementpb_azure_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AzureDatabaseClient is the client API for AzureDatabase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AzureDatabaseClient interface {
	// AddAzureDatabase adds Azure Database instance.
	AddAzureDatabase(ctx context.Context, in *AddAzureDatabaseRequest, opts ...grpc.CallOption) (*AddAzureDatabaseResponse, error)
}

type azureDatabaseClient struct {
	cc grpc.ClientConnInterface
}

func NewAzureDatabaseClient(cc grpc.ClientConnInterface) AzureDatabaseClient {
	return &azureDatabaseClient{cc}
}

func (c *azureDatabaseClient) AddAzureDatabase(ctx context.Context, in *AddAzureDatabaseRequest, opts ...grpc.CallOption) (*AddAzureDatabaseResponse, error) {
	out := new(AddAzureDatabaseResponse)
	err := c.cc.Invoke(ctx, "/management.AzureDatabase/AddAzureDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AzureDatabaseServer is the server API for AzureDatabase service.
type AzureDatabaseServer interface {
	// AddAzureDatabase adds Azure Database instance.
	AddAzureDatabase(context.Context, *AddAzureDatabaseRequest) (*AddAzureDatabaseResponse, error)
}

// UnimplementedAzureDatabaseServer can be embedded to have forward compatible implementations.
type UnimplementedAzureDatabaseServer struct {
}

func (*UnimplementedAzureDatabaseServer) AddAzureDatabase(context.Context, *AddAzureDatabaseRequest) (*AddAzureDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAzureDatabase not implemented")
}

func RegisterAzureDatabaseServer(s *grpc.Server, srv AzureDatabaseServer) {
	s.RegisterService(&_AzureDatabase_serviceDesc, srv)
}

func _AzureDatabase_AddAzureDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAzureDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AzureDatabaseServer).AddAzureDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.AzureDatabase/AddAzureDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AzureDatabaseServer).AddAzureDatabase(ctx, req.(*AddAzureDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AzureDatabase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.AzureDatabase",
	HandlerType: (*AzureDatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAzureDatabase",
			Handler:    _AzureDatabase_AddAzureDatabase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/azure.proto",
}