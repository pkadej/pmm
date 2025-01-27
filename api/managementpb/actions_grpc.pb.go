// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: managementpb/actions.proto

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

// ActionsClient is the client API for Actions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionsClient interface {
	// GetAction gets result of a given Action.
	GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error)
	// StartMySQLExplainAction starts MySQL EXPLAIN Action with traditional output.
	StartMySQLExplainAction(ctx context.Context, in *StartMySQLExplainActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainActionResponse, error)
	// StartMySQLExplainJSONAction starts MySQL EXPLAIN Action with JSON output.
	StartMySQLExplainJSONAction(ctx context.Context, in *StartMySQLExplainJSONActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainJSONActionResponse, error)
	// StartMySQLExplainTraditionalJSONAction starts MySQL EXPLAIN Action with traditional JSON output.
	StartMySQLExplainTraditionalJSONAction(ctx context.Context, in *StartMySQLExplainTraditionalJSONActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainTraditionalJSONActionResponse, error)
	// StartMySQLShowCreateTableAction starts MySQL SHOW CREATE TABLE Action.
	StartMySQLShowCreateTableAction(ctx context.Context, in *StartMySQLShowCreateTableActionRequest, opts ...grpc.CallOption) (*StartMySQLShowCreateTableActionResponse, error)
	// StartMySQLShowTableStatusAction starts MySQL SHOW TABLE STATUS Action.
	StartMySQLShowTableStatusAction(ctx context.Context, in *StartMySQLShowTableStatusActionRequest, opts ...grpc.CallOption) (*StartMySQLShowTableStatusActionResponse, error)
	// StartMySQLShowIndexAction starts MySQL SHOW INDEX Action.
	StartMySQLShowIndexAction(ctx context.Context, in *StartMySQLShowIndexActionRequest, opts ...grpc.CallOption) (*StartMySQLShowIndexActionResponse, error)
	// StartPostgreSQLShowCreateTableAction starts PostgreSQL SHOW CREATE TABLE Action.
	StartPostgreSQLShowCreateTableAction(ctx context.Context, in *StartPostgreSQLShowCreateTableActionRequest, opts ...grpc.CallOption) (*StartPostgreSQLShowCreateTableActionResponse, error)
	// StartPostgreSQLShowIndexAction starts PostgreSQL SHOW INDEX Action.
	StartPostgreSQLShowIndexAction(ctx context.Context, in *StartPostgreSQLShowIndexActionRequest, opts ...grpc.CallOption) (*StartPostgreSQLShowIndexActionResponse, error)
	// StartMongoDBExplainAction starts MongoDB EXPLAIN Action.
	StartMongoDBExplainAction(ctx context.Context, in *StartMongoDBExplainActionRequest, opts ...grpc.CallOption) (*StartMongoDBExplainActionResponse, error)
	// StartPTSummaryAction starts pt-summary Action.
	StartPTSummaryAction(ctx context.Context, in *StartPTSummaryActionRequest, opts ...grpc.CallOption) (*StartPTSummaryActionResponse, error)
	// StartPTPgSummaryAction starts pt-pg-summary Action.
	StartPTPgSummaryAction(ctx context.Context, in *StartPTPgSummaryActionRequest, opts ...grpc.CallOption) (*StartPTPgSummaryActionResponse, error)
	// StartPTMongoDBSummaryAction starts pt-mongodb-summary Action.
	StartPTMongoDBSummaryAction(ctx context.Context, in *StartPTMongoDBSummaryActionRequest, opts ...grpc.CallOption) (*StartPTMongoDBSummaryActionResponse, error)
	// StartPTMySQLSummaryAction starts pt-mysql-summary Action.
	StartPTMySQLSummaryAction(ctx context.Context, in *StartPTMySQLSummaryActionRequest, opts ...grpc.CallOption) (*StartPTMySQLSummaryActionResponse, error)
	// CancelAction stops an Action.
	CancelAction(ctx context.Context, in *CancelActionRequest, opts ...grpc.CallOption) (*CancelActionResponse, error)
}

type actionsClient struct {
	cc grpc.ClientConnInterface
}

func NewActionsClient(cc grpc.ClientConnInterface) ActionsClient {
	return &actionsClient{cc}
}

func (c *actionsClient) GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionResponse, error) {
	out := new(GetActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/GetAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartMySQLExplainAction(ctx context.Context, in *StartMySQLExplainActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainActionResponse, error) {
	out := new(StartMySQLExplainActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartMySQLExplainAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartMySQLExplainJSONAction(ctx context.Context, in *StartMySQLExplainJSONActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainJSONActionResponse, error) {
	out := new(StartMySQLExplainJSONActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartMySQLExplainJSONAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartMySQLExplainTraditionalJSONAction(ctx context.Context, in *StartMySQLExplainTraditionalJSONActionRequest, opts ...grpc.CallOption) (*StartMySQLExplainTraditionalJSONActionResponse, error) {
	out := new(StartMySQLExplainTraditionalJSONActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartMySQLExplainTraditionalJSONAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartMySQLShowCreateTableAction(ctx context.Context, in *StartMySQLShowCreateTableActionRequest, opts ...grpc.CallOption) (*StartMySQLShowCreateTableActionResponse, error) {
	out := new(StartMySQLShowCreateTableActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartMySQLShowCreateTableAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartMySQLShowTableStatusAction(ctx context.Context, in *StartMySQLShowTableStatusActionRequest, opts ...grpc.CallOption) (*StartMySQLShowTableStatusActionResponse, error) {
	out := new(StartMySQLShowTableStatusActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartMySQLShowTableStatusAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartMySQLShowIndexAction(ctx context.Context, in *StartMySQLShowIndexActionRequest, opts ...grpc.CallOption) (*StartMySQLShowIndexActionResponse, error) {
	out := new(StartMySQLShowIndexActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartMySQLShowIndexAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartPostgreSQLShowCreateTableAction(ctx context.Context, in *StartPostgreSQLShowCreateTableActionRequest, opts ...grpc.CallOption) (*StartPostgreSQLShowCreateTableActionResponse, error) {
	out := new(StartPostgreSQLShowCreateTableActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartPostgreSQLShowCreateTableAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartPostgreSQLShowIndexAction(ctx context.Context, in *StartPostgreSQLShowIndexActionRequest, opts ...grpc.CallOption) (*StartPostgreSQLShowIndexActionResponse, error) {
	out := new(StartPostgreSQLShowIndexActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartPostgreSQLShowIndexAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartMongoDBExplainAction(ctx context.Context, in *StartMongoDBExplainActionRequest, opts ...grpc.CallOption) (*StartMongoDBExplainActionResponse, error) {
	out := new(StartMongoDBExplainActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartMongoDBExplainAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartPTSummaryAction(ctx context.Context, in *StartPTSummaryActionRequest, opts ...grpc.CallOption) (*StartPTSummaryActionResponse, error) {
	out := new(StartPTSummaryActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartPTSummaryAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartPTPgSummaryAction(ctx context.Context, in *StartPTPgSummaryActionRequest, opts ...grpc.CallOption) (*StartPTPgSummaryActionResponse, error) {
	out := new(StartPTPgSummaryActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartPTPgSummaryAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartPTMongoDBSummaryAction(ctx context.Context, in *StartPTMongoDBSummaryActionRequest, opts ...grpc.CallOption) (*StartPTMongoDBSummaryActionResponse, error) {
	out := new(StartPTMongoDBSummaryActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartPTMongoDBSummaryAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) StartPTMySQLSummaryAction(ctx context.Context, in *StartPTMySQLSummaryActionRequest, opts ...grpc.CallOption) (*StartPTMySQLSummaryActionResponse, error) {
	out := new(StartPTMySQLSummaryActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/StartPTMySQLSummaryAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsClient) CancelAction(ctx context.Context, in *CancelActionRequest, opts ...grpc.CallOption) (*CancelActionResponse, error) {
	out := new(CancelActionResponse)
	err := c.cc.Invoke(ctx, "/management.Actions/CancelAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionsServer is the server API for Actions service.
// All implementations must embed UnimplementedActionsServer
// for forward compatibility
type ActionsServer interface {
	// GetAction gets result of a given Action.
	GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error)
	// StartMySQLExplainAction starts MySQL EXPLAIN Action with traditional output.
	StartMySQLExplainAction(context.Context, *StartMySQLExplainActionRequest) (*StartMySQLExplainActionResponse, error)
	// StartMySQLExplainJSONAction starts MySQL EXPLAIN Action with JSON output.
	StartMySQLExplainJSONAction(context.Context, *StartMySQLExplainJSONActionRequest) (*StartMySQLExplainJSONActionResponse, error)
	// StartMySQLExplainTraditionalJSONAction starts MySQL EXPLAIN Action with traditional JSON output.
	StartMySQLExplainTraditionalJSONAction(context.Context, *StartMySQLExplainTraditionalJSONActionRequest) (*StartMySQLExplainTraditionalJSONActionResponse, error)
	// StartMySQLShowCreateTableAction starts MySQL SHOW CREATE TABLE Action.
	StartMySQLShowCreateTableAction(context.Context, *StartMySQLShowCreateTableActionRequest) (*StartMySQLShowCreateTableActionResponse, error)
	// StartMySQLShowTableStatusAction starts MySQL SHOW TABLE STATUS Action.
	StartMySQLShowTableStatusAction(context.Context, *StartMySQLShowTableStatusActionRequest) (*StartMySQLShowTableStatusActionResponse, error)
	// StartMySQLShowIndexAction starts MySQL SHOW INDEX Action.
	StartMySQLShowIndexAction(context.Context, *StartMySQLShowIndexActionRequest) (*StartMySQLShowIndexActionResponse, error)
	// StartPostgreSQLShowCreateTableAction starts PostgreSQL SHOW CREATE TABLE Action.
	StartPostgreSQLShowCreateTableAction(context.Context, *StartPostgreSQLShowCreateTableActionRequest) (*StartPostgreSQLShowCreateTableActionResponse, error)
	// StartPostgreSQLShowIndexAction starts PostgreSQL SHOW INDEX Action.
	StartPostgreSQLShowIndexAction(context.Context, *StartPostgreSQLShowIndexActionRequest) (*StartPostgreSQLShowIndexActionResponse, error)
	// StartMongoDBExplainAction starts MongoDB EXPLAIN Action.
	StartMongoDBExplainAction(context.Context, *StartMongoDBExplainActionRequest) (*StartMongoDBExplainActionResponse, error)
	// StartPTSummaryAction starts pt-summary Action.
	StartPTSummaryAction(context.Context, *StartPTSummaryActionRequest) (*StartPTSummaryActionResponse, error)
	// StartPTPgSummaryAction starts pt-pg-summary Action.
	StartPTPgSummaryAction(context.Context, *StartPTPgSummaryActionRequest) (*StartPTPgSummaryActionResponse, error)
	// StartPTMongoDBSummaryAction starts pt-mongodb-summary Action.
	StartPTMongoDBSummaryAction(context.Context, *StartPTMongoDBSummaryActionRequest) (*StartPTMongoDBSummaryActionResponse, error)
	// StartPTMySQLSummaryAction starts pt-mysql-summary Action.
	StartPTMySQLSummaryAction(context.Context, *StartPTMySQLSummaryActionRequest) (*StartPTMySQLSummaryActionResponse, error)
	// CancelAction stops an Action.
	CancelAction(context.Context, *CancelActionRequest) (*CancelActionResponse, error)
	mustEmbedUnimplementedActionsServer()
}

// UnimplementedActionsServer must be embedded to have forward compatible implementations.
type UnimplementedActionsServer struct {
}

func (UnimplementedActionsServer) GetAction(context.Context, *GetActionRequest) (*GetActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAction not implemented")
}
func (UnimplementedActionsServer) StartMySQLExplainAction(context.Context, *StartMySQLExplainActionRequest) (*StartMySQLExplainActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLExplainAction not implemented")
}
func (UnimplementedActionsServer) StartMySQLExplainJSONAction(context.Context, *StartMySQLExplainJSONActionRequest) (*StartMySQLExplainJSONActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLExplainJSONAction not implemented")
}
func (UnimplementedActionsServer) StartMySQLExplainTraditionalJSONAction(context.Context, *StartMySQLExplainTraditionalJSONActionRequest) (*StartMySQLExplainTraditionalJSONActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLExplainTraditionalJSONAction not implemented")
}
func (UnimplementedActionsServer) StartMySQLShowCreateTableAction(context.Context, *StartMySQLShowCreateTableActionRequest) (*StartMySQLShowCreateTableActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLShowCreateTableAction not implemented")
}
func (UnimplementedActionsServer) StartMySQLShowTableStatusAction(context.Context, *StartMySQLShowTableStatusActionRequest) (*StartMySQLShowTableStatusActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLShowTableStatusAction not implemented")
}
func (UnimplementedActionsServer) StartMySQLShowIndexAction(context.Context, *StartMySQLShowIndexActionRequest) (*StartMySQLShowIndexActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMySQLShowIndexAction not implemented")
}
func (UnimplementedActionsServer) StartPostgreSQLShowCreateTableAction(context.Context, *StartPostgreSQLShowCreateTableActionRequest) (*StartPostgreSQLShowCreateTableActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPostgreSQLShowCreateTableAction not implemented")
}
func (UnimplementedActionsServer) StartPostgreSQLShowIndexAction(context.Context, *StartPostgreSQLShowIndexActionRequest) (*StartPostgreSQLShowIndexActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPostgreSQLShowIndexAction not implemented")
}
func (UnimplementedActionsServer) StartMongoDBExplainAction(context.Context, *StartMongoDBExplainActionRequest) (*StartMongoDBExplainActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMongoDBExplainAction not implemented")
}
func (UnimplementedActionsServer) StartPTSummaryAction(context.Context, *StartPTSummaryActionRequest) (*StartPTSummaryActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPTSummaryAction not implemented")
}
func (UnimplementedActionsServer) StartPTPgSummaryAction(context.Context, *StartPTPgSummaryActionRequest) (*StartPTPgSummaryActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPTPgSummaryAction not implemented")
}
func (UnimplementedActionsServer) StartPTMongoDBSummaryAction(context.Context, *StartPTMongoDBSummaryActionRequest) (*StartPTMongoDBSummaryActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPTMongoDBSummaryAction not implemented")
}
func (UnimplementedActionsServer) StartPTMySQLSummaryAction(context.Context, *StartPTMySQLSummaryActionRequest) (*StartPTMySQLSummaryActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPTMySQLSummaryAction not implemented")
}
func (UnimplementedActionsServer) CancelAction(context.Context, *CancelActionRequest) (*CancelActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAction not implemented")
}
func (UnimplementedActionsServer) mustEmbedUnimplementedActionsServer() {}

// UnsafeActionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionsServer will
// result in compilation errors.
type UnsafeActionsServer interface {
	mustEmbedUnimplementedActionsServer()
}

func RegisterActionsServer(s grpc.ServiceRegistrar, srv ActionsServer) {
	s.RegisterService(&Actions_ServiceDesc, srv)
}

func _Actions_GetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).GetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/GetAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).GetAction(ctx, req.(*GetActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartMySQLExplainAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLExplainActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartMySQLExplainAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartMySQLExplainAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartMySQLExplainAction(ctx, req.(*StartMySQLExplainActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartMySQLExplainJSONAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLExplainJSONActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartMySQLExplainJSONAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartMySQLExplainJSONAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartMySQLExplainJSONAction(ctx, req.(*StartMySQLExplainJSONActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartMySQLExplainTraditionalJSONAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLExplainTraditionalJSONActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartMySQLExplainTraditionalJSONAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartMySQLExplainTraditionalJSONAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartMySQLExplainTraditionalJSONAction(ctx, req.(*StartMySQLExplainTraditionalJSONActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartMySQLShowCreateTableAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLShowCreateTableActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartMySQLShowCreateTableAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartMySQLShowCreateTableAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartMySQLShowCreateTableAction(ctx, req.(*StartMySQLShowCreateTableActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartMySQLShowTableStatusAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLShowTableStatusActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartMySQLShowTableStatusAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartMySQLShowTableStatusAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartMySQLShowTableStatusAction(ctx, req.(*StartMySQLShowTableStatusActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartMySQLShowIndexAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMySQLShowIndexActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartMySQLShowIndexAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartMySQLShowIndexAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartMySQLShowIndexAction(ctx, req.(*StartMySQLShowIndexActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartPostgreSQLShowCreateTableAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPostgreSQLShowCreateTableActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartPostgreSQLShowCreateTableAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartPostgreSQLShowCreateTableAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartPostgreSQLShowCreateTableAction(ctx, req.(*StartPostgreSQLShowCreateTableActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartPostgreSQLShowIndexAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPostgreSQLShowIndexActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartPostgreSQLShowIndexAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartPostgreSQLShowIndexAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartPostgreSQLShowIndexAction(ctx, req.(*StartPostgreSQLShowIndexActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartMongoDBExplainAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMongoDBExplainActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartMongoDBExplainAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartMongoDBExplainAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartMongoDBExplainAction(ctx, req.(*StartMongoDBExplainActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartPTSummaryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPTSummaryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartPTSummaryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartPTSummaryAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartPTSummaryAction(ctx, req.(*StartPTSummaryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartPTPgSummaryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPTPgSummaryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartPTPgSummaryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartPTPgSummaryAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartPTPgSummaryAction(ctx, req.(*StartPTPgSummaryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartPTMongoDBSummaryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPTMongoDBSummaryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartPTMongoDBSummaryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartPTMongoDBSummaryAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartPTMongoDBSummaryAction(ctx, req.(*StartPTMongoDBSummaryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_StartPTMySQLSummaryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPTMySQLSummaryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).StartPTMySQLSummaryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/StartPTMySQLSummaryAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).StartPTMySQLSummaryAction(ctx, req.(*StartPTMySQLSummaryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Actions_CancelAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).CancelAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Actions/CancelAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).CancelAction(ctx, req.(*CancelActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Actions_ServiceDesc is the grpc.ServiceDesc for Actions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Actions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.Actions",
	HandlerType: (*ActionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAction",
			Handler:    _Actions_GetAction_Handler,
		},
		{
			MethodName: "StartMySQLExplainAction",
			Handler:    _Actions_StartMySQLExplainAction_Handler,
		},
		{
			MethodName: "StartMySQLExplainJSONAction",
			Handler:    _Actions_StartMySQLExplainJSONAction_Handler,
		},
		{
			MethodName: "StartMySQLExplainTraditionalJSONAction",
			Handler:    _Actions_StartMySQLExplainTraditionalJSONAction_Handler,
		},
		{
			MethodName: "StartMySQLShowCreateTableAction",
			Handler:    _Actions_StartMySQLShowCreateTableAction_Handler,
		},
		{
			MethodName: "StartMySQLShowTableStatusAction",
			Handler:    _Actions_StartMySQLShowTableStatusAction_Handler,
		},
		{
			MethodName: "StartMySQLShowIndexAction",
			Handler:    _Actions_StartMySQLShowIndexAction_Handler,
		},
		{
			MethodName: "StartPostgreSQLShowCreateTableAction",
			Handler:    _Actions_StartPostgreSQLShowCreateTableAction_Handler,
		},
		{
			MethodName: "StartPostgreSQLShowIndexAction",
			Handler:    _Actions_StartPostgreSQLShowIndexAction_Handler,
		},
		{
			MethodName: "StartMongoDBExplainAction",
			Handler:    _Actions_StartMongoDBExplainAction_Handler,
		},
		{
			MethodName: "StartPTSummaryAction",
			Handler:    _Actions_StartPTSummaryAction_Handler,
		},
		{
			MethodName: "StartPTPgSummaryAction",
			Handler:    _Actions_StartPTPgSummaryAction_Handler,
		},
		{
			MethodName: "StartPTMongoDBSummaryAction",
			Handler:    _Actions_StartPTMongoDBSummaryAction_Handler,
		},
		{
			MethodName: "StartPTMySQLSummaryAction",
			Handler:    _Actions_StartPTMySQLSummaryAction_Handler,
		},
		{
			MethodName: "CancelAction",
			Handler:    _Actions_CancelAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/actions.proto",
}
