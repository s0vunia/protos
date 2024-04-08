// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: orchestrator/orchestrator.proto

package orchestrator

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

// OrchestratorClient is the client API for Orchestrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrchestratorClient interface {
	// Expression create expression
	CreateExpression(ctx context.Context, in *CreateExpressionRequest, opts ...grpc.CallOption) (*CreateExpressionResponse, error)
	// Expressions return expression
	GetExpression(ctx context.Context, in *GetExpressionRequest, opts ...grpc.CallOption) (*GetExpressionResponse, error)
	// Expressions return all expressions
	GetExpressions(ctx context.Context, in *GetExpressionsRequest, opts ...grpc.CallOption) (*GetExpressionsResponse, error)
	// GetAgents return all agents
	GetAgents(ctx context.Context, in *GetAgentsRequest, opts ...grpc.CallOption) (*GetAgentsResponse, error)
	// GetOperators return all operators
	GetOperators(ctx context.Context, in *GetOperatorsRequest, opts ...grpc.CallOption) (*GetOperatorsResponse, error)
}

type orchestratorClient struct {
	cc grpc.ClientConnInterface
}

func NewOrchestratorClient(cc grpc.ClientConnInterface) OrchestratorClient {
	return &orchestratorClient{cc}
}

func (c *orchestratorClient) CreateExpression(ctx context.Context, in *CreateExpressionRequest, opts ...grpc.CallOption) (*CreateExpressionResponse, error) {
	out := new(CreateExpressionResponse)
	err := c.cc.Invoke(ctx, "/orchestrator.Orchestrator/CreateExpression", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) GetExpression(ctx context.Context, in *GetExpressionRequest, opts ...grpc.CallOption) (*GetExpressionResponse, error) {
	out := new(GetExpressionResponse)
	err := c.cc.Invoke(ctx, "/orchestrator.Orchestrator/GetExpression", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) GetExpressions(ctx context.Context, in *GetExpressionsRequest, opts ...grpc.CallOption) (*GetExpressionsResponse, error) {
	out := new(GetExpressionsResponse)
	err := c.cc.Invoke(ctx, "/orchestrator.Orchestrator/GetExpressions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) GetAgents(ctx context.Context, in *GetAgentsRequest, opts ...grpc.CallOption) (*GetAgentsResponse, error) {
	out := new(GetAgentsResponse)
	err := c.cc.Invoke(ctx, "/orchestrator.Orchestrator/GetAgents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) GetOperators(ctx context.Context, in *GetOperatorsRequest, opts ...grpc.CallOption) (*GetOperatorsResponse, error) {
	out := new(GetOperatorsResponse)
	err := c.cc.Invoke(ctx, "/orchestrator.Orchestrator/GetOperators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorServer is the server API for Orchestrator service.
// All implementations must embed UnimplementedOrchestratorServer
// for forward compatibility
type OrchestratorServer interface {
	// Expression create expression
	CreateExpression(context.Context, *CreateExpressionRequest) (*CreateExpressionResponse, error)
	// Expressions return expression
	GetExpression(context.Context, *GetExpressionRequest) (*GetExpressionResponse, error)
	// Expressions return all expressions
	GetExpressions(context.Context, *GetExpressionsRequest) (*GetExpressionsResponse, error)
	// GetAgents return all agents
	GetAgents(context.Context, *GetAgentsRequest) (*GetAgentsResponse, error)
	// GetOperators return all operators
	GetOperators(context.Context, *GetOperatorsRequest) (*GetOperatorsResponse, error)
	mustEmbedUnimplementedOrchestratorServer()
}

// UnimplementedOrchestratorServer must be embedded to have forward compatible implementations.
type UnimplementedOrchestratorServer struct {
}

func (UnimplementedOrchestratorServer) CreateExpression(context.Context, *CreateExpressionRequest) (*CreateExpressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpression not implemented")
}
func (UnimplementedOrchestratorServer) GetExpression(context.Context, *GetExpressionRequest) (*GetExpressionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpression not implemented")
}
func (UnimplementedOrchestratorServer) GetExpressions(context.Context, *GetExpressionsRequest) (*GetExpressionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpressions not implemented")
}
func (UnimplementedOrchestratorServer) GetAgents(context.Context, *GetAgentsRequest) (*GetAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgents not implemented")
}
func (UnimplementedOrchestratorServer) GetOperators(context.Context, *GetOperatorsRequest) (*GetOperatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperators not implemented")
}
func (UnimplementedOrchestratorServer) mustEmbedUnimplementedOrchestratorServer() {}

// UnsafeOrchestratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrchestratorServer will
// result in compilation errors.
type UnsafeOrchestratorServer interface {
	mustEmbedUnimplementedOrchestratorServer()
}

func RegisterOrchestratorServer(s grpc.ServiceRegistrar, srv OrchestratorServer) {
	s.RegisterService(&Orchestrator_ServiceDesc, srv)
}

func _Orchestrator_CreateExpression_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExpressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).CreateExpression(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orchestrator.Orchestrator/CreateExpression",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).CreateExpression(ctx, req.(*CreateExpressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_GetExpression_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExpressionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetExpression(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orchestrator.Orchestrator/GetExpression",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetExpression(ctx, req.(*GetExpressionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_GetExpressions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExpressionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetExpressions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orchestrator.Orchestrator/GetExpressions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetExpressions(ctx, req.(*GetExpressionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_GetAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orchestrator.Orchestrator/GetAgents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetAgents(ctx, req.(*GetAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_GetOperators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).GetOperators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orchestrator.Orchestrator/GetOperators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).GetOperators(ctx, req.(*GetOperatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Orchestrator_ServiceDesc is the grpc.ServiceDesc for Orchestrator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Orchestrator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orchestrator.Orchestrator",
	HandlerType: (*OrchestratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExpression",
			Handler:    _Orchestrator_CreateExpression_Handler,
		},
		{
			MethodName: "GetExpression",
			Handler:    _Orchestrator_GetExpression_Handler,
		},
		{
			MethodName: "GetExpressions",
			Handler:    _Orchestrator_GetExpressions_Handler,
		},
		{
			MethodName: "GetAgents",
			Handler:    _Orchestrator_GetAgents_Handler,
		},
		{
			MethodName: "GetOperators",
			Handler:    _Orchestrator_GetOperators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orchestrator/orchestrator.proto",
}
