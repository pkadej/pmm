// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: managementpb/rds.proto

package managementpb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RDSClient is the client API for RDS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RDSClient interface {
	// DiscoverRDS discovers RDS instances.
	DiscoverRDS(ctx context.Context, in *DiscoverRDSRequest, opts ...grpc.CallOption) (*DiscoverRDSResponse, error)
	// AddRDS adds RDS instance.
	AddRDS(ctx context.Context, in *AddRDSRequest, opts ...grpc.CallOption) (*AddRDSResponse, error)
}

type rDSClient struct {
	cc grpc.ClientConnInterface
}

func NewRDSClient(cc grpc.ClientConnInterface) RDSClient {
	return &rDSClient{cc}
}

func (c *rDSClient) DiscoverRDS(ctx context.Context, in *DiscoverRDSRequest, opts ...grpc.CallOption) (*DiscoverRDSResponse, error) {
	out := new(DiscoverRDSResponse)
	err := c.cc.Invoke(ctx, "/management.RDS/DiscoverRDS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rDSClient) AddRDS(ctx context.Context, in *AddRDSRequest, opts ...grpc.CallOption) (*AddRDSResponse, error) {
	out := new(AddRDSResponse)
	err := c.cc.Invoke(ctx, "/management.RDS/AddRDS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RDSServer is the server API for RDS service.
// All implementations must embed UnimplementedRDSServer
// for forward compatibility
type RDSServer interface {
	// DiscoverRDS discovers RDS instances.
	DiscoverRDS(context.Context, *DiscoverRDSRequest) (*DiscoverRDSResponse, error)
	// AddRDS adds RDS instance.
	AddRDS(context.Context, *AddRDSRequest) (*AddRDSResponse, error)
	mustEmbedUnimplementedRDSServer()
}

// UnimplementedRDSServer must be embedded to have forward compatible implementations.
type UnimplementedRDSServer struct {
}

func (UnimplementedRDSServer) DiscoverRDS(context.Context, *DiscoverRDSRequest) (*DiscoverRDSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverRDS not implemented")
}
func (UnimplementedRDSServer) AddRDS(context.Context, *AddRDSRequest) (*AddRDSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRDS not implemented")
}
func (UnimplementedRDSServer) mustEmbedUnimplementedRDSServer() {}

// UnsafeRDSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RDSServer will
// result in compilation errors.
type UnsafeRDSServer interface {
	mustEmbedUnimplementedRDSServer()
}

func RegisterRDSServer(s grpc.ServiceRegistrar, srv RDSServer) {
	s.RegisterService(&RDS_ServiceDesc, srv)
}

func _RDS_DiscoverRDS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverRDSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDSServer).DiscoverRDS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.RDS/DiscoverRDS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDSServer).DiscoverRDS(ctx, req.(*DiscoverRDSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RDS_AddRDS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRDSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDSServer).AddRDS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.RDS/AddRDS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDSServer).AddRDS(ctx, req.(*AddRDSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RDS_ServiceDesc is the grpc.ServiceDesc for RDS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RDS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.RDS",
	HandlerType: (*RDSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiscoverRDS",
			Handler:    _RDS_DiscoverRDS_Handler,
		},
		{
			MethodName: "AddRDS",
			Handler:    _RDS_AddRDS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/rds.proto",
}
