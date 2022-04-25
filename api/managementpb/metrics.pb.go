// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: managementpb/metrics.proto

package managementpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MetricsMode defines desired metrics mode for agent,
// it can be pull, push or auto mode chosen by server.
type MetricsMode int32

const (
	MetricsMode_AUTO MetricsMode = 0
	MetricsMode_PULL MetricsMode = 1
	MetricsMode_PUSH MetricsMode = 2
)

// Enum value maps for MetricsMode.
var (
	MetricsMode_name = map[int32]string{
		0: "AUTO",
		1: "PULL",
		2: "PUSH",
	}
	MetricsMode_value = map[string]int32{
		"AUTO": 0,
		"PULL": 1,
		"PUSH": 2,
	}
)

func (x MetricsMode) Enum() *MetricsMode {
	p := new(MetricsMode)
	*p = x
	return p
}

func (x MetricsMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricsMode) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_metrics_proto_enumTypes[0].Descriptor()
}

func (MetricsMode) Type() protoreflect.EnumType {
	return &file_managementpb_metrics_proto_enumTypes[0]
}

func (x MetricsMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricsMode.Descriptor instead.
func (MetricsMode) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_metrics_proto_rawDescGZIP(), []int{0}
}

var File_managementpb_metrics_proto protoreflect.FileDescriptor

var file_managementpb_metrics_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2a, 0x2b, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x55, 0x54, 0x4f, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50,
	0x55, 0x53, 0x48, 0x10, 0x02, 0x42, 0x8f, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2f, 0x70, 0x6d, 0x6d,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x0a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0xe2, 0x02, 0x16, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_metrics_proto_rawDescOnce sync.Once
	file_managementpb_metrics_proto_rawDescData = file_managementpb_metrics_proto_rawDesc
)

func file_managementpb_metrics_proto_rawDescGZIP() []byte {
	file_managementpb_metrics_proto_rawDescOnce.Do(func() {
		file_managementpb_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_metrics_proto_rawDescData)
	})
	return file_managementpb_metrics_proto_rawDescData
}

var file_managementpb_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_managementpb_metrics_proto_goTypes = []interface{}{
	(MetricsMode)(0), // 0: management.MetricsMode
}
var file_managementpb_metrics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_managementpb_metrics_proto_init() }
func file_managementpb_metrics_proto_init() {
	if File_managementpb_metrics_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_metrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_managementpb_metrics_proto_goTypes,
		DependencyIndexes: file_managementpb_metrics_proto_depIdxs,
		EnumInfos:         file_managementpb_metrics_proto_enumTypes,
	}.Build()
	File_managementpb_metrics_proto = out.File
	file_managementpb_metrics_proto_rawDesc = nil
	file_managementpb_metrics_proto_goTypes = nil
	file_managementpb_metrics_proto_depIdxs = nil
}
