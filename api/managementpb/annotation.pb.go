// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: managementpb/annotation.proto

package managementpb

import (
	context "context"
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

// AddAnnotationRequest is a params to add new annotation.
type AddAnnotationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An annotation description. Required.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Tags are used to filter annotations.
	Tags []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	// Used for annotate node.
	NodeName string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Used for annotate services.
	ServiceNames []string `protobuf:"bytes,4,rep,name=service_names,json=serviceNames,proto3" json:"service_names,omitempty"`
}

func (x *AddAnnotationRequest) Reset() {
	*x = AddAnnotationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_annotation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAnnotationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAnnotationRequest) ProtoMessage() {}

func (x *AddAnnotationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_annotation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAnnotationRequest.ProtoReflect.Descriptor instead.
func (*AddAnnotationRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_annotation_proto_rawDescGZIP(), []int{0}
}

func (x *AddAnnotationRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AddAnnotationRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *AddAnnotationRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AddAnnotationRequest) GetServiceNames() []string {
	if x != nil {
		return x.ServiceNames
	}
	return nil
}

type AddAnnotationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddAnnotationResponse) Reset() {
	*x = AddAnnotationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_annotation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddAnnotationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddAnnotationResponse) ProtoMessage() {}

func (x *AddAnnotationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_annotation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddAnnotationResponse.ProtoReflect.Descriptor instead.
func (*AddAnnotationResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_annotation_proto_rawDescGZIP(), []int{1}
}

var File_managementpb_annotation_proto protoreflect.FileDescriptor

var file_managementpb_annotation_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f,
	0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x88, 0x01, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x17, 0x0a, 0x15,
	0x41, 0x64, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8d, 0x01, 0x0a, 0x0a, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7f, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x23, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x41,
	0x64, 0x64, 0x3a, 0x01, 0x2a, 0x42, 0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_annotation_proto_rawDescOnce sync.Once
	file_managementpb_annotation_proto_rawDescData = file_managementpb_annotation_proto_rawDesc
)

func file_managementpb_annotation_proto_rawDescGZIP() []byte {
	file_managementpb_annotation_proto_rawDescOnce.Do(func() {
		file_managementpb_annotation_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_annotation_proto_rawDescData)
	})
	return file_managementpb_annotation_proto_rawDescData
}

var file_managementpb_annotation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_managementpb_annotation_proto_goTypes = []interface{}{
	(*AddAnnotationRequest)(nil),  // 0: management.AddAnnotationRequest
	(*AddAnnotationResponse)(nil), // 1: management.AddAnnotationResponse
}
var file_managementpb_annotation_proto_depIdxs = []int32{
	0, // 0: management.Annotation.AddAnnotation:input_type -> management.AddAnnotationRequest
	1, // 1: management.Annotation.AddAnnotation:output_type -> management.AddAnnotationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_managementpb_annotation_proto_init() }
func file_managementpb_annotation_proto_init() {
	if File_managementpb_annotation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managementpb_annotation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAnnotationRequest); i {
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
		file_managementpb_annotation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddAnnotationResponse); i {
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
			RawDescriptor: file_managementpb_annotation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_annotation_proto_goTypes,
		DependencyIndexes: file_managementpb_annotation_proto_depIdxs,
		MessageInfos:      file_managementpb_annotation_proto_msgTypes,
	}.Build()
	File_managementpb_annotation_proto = out.File
	file_managementpb_annotation_proto_rawDesc = nil
	file_managementpb_annotation_proto_goTypes = nil
	file_managementpb_annotation_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnnotationClient is the client API for Annotation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnnotationClient interface {
	// AddAnnotation adds annotation.
	AddAnnotation(ctx context.Context, in *AddAnnotationRequest, opts ...grpc.CallOption) (*AddAnnotationResponse, error)
}

type annotationClient struct {
	cc grpc.ClientConnInterface
}

func NewAnnotationClient(cc grpc.ClientConnInterface) AnnotationClient {
	return &annotationClient{cc}
}

func (c *annotationClient) AddAnnotation(ctx context.Context, in *AddAnnotationRequest, opts ...grpc.CallOption) (*AddAnnotationResponse, error) {
	out := new(AddAnnotationResponse)
	err := c.cc.Invoke(ctx, "/management.Annotation/AddAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnnotationServer is the server API for Annotation service.
type AnnotationServer interface {
	// AddAnnotation adds annotation.
	AddAnnotation(context.Context, *AddAnnotationRequest) (*AddAnnotationResponse, error)
}

// UnimplementedAnnotationServer can be embedded to have forward compatible implementations.
type UnimplementedAnnotationServer struct {
}

func (*UnimplementedAnnotationServer) AddAnnotation(context.Context, *AddAnnotationRequest) (*AddAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAnnotation not implemented")
}

func RegisterAnnotationServer(s *grpc.Server, srv AnnotationServer) {
	s.RegisterService(&_Annotation_serviceDesc, srv)
}

func _Annotation_AddAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnotationServer).AddAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Annotation/AddAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnotationServer).AddAnnotation(ctx, req.(*AddAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Annotation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.Annotation",
	HandlerType: (*AnnotationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddAnnotation",
			Handler:    _Annotation_AddAnnotation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/annotation.proto",
}
