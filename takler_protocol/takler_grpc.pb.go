// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: takler_protocol/takler.proto

package takler_protocol

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

// TaklerServerClient is the client API for TaklerServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaklerServerClient interface {
	// child command
	RunInitCommand(ctx context.Context, in *InitCommand, opts ...grpc.CallOption) (*ServiceResponse, error)
	RunCompleteCommand(ctx context.Context, in *CompleteCommand, opts ...grpc.CallOption) (*ServiceResponse, error)
	RunAbortCommand(ctx context.Context, in *AbortCommand, opts ...grpc.CallOption) (*ServiceResponse, error)
	RunEventCommand(ctx context.Context, in *EventCommand, opts ...grpc.CallOption) (*ServiceResponse, error)
	RunMeterCommand(ctx context.Context, in *MeterCommand, opts ...grpc.CallOption) (*ServiceResponse, error)
	RunRequeueCommand(ctx context.Context, in *RequeueCommand, opts ...grpc.CallOption) (*ServiceResponse, error)
	RunSuspendCommand(ctx context.Context, in *SuspendCommand, opts ...grpc.CallOption) (*ServiceResponse, error)
	RunResumeCommand(ctx context.Context, in *SuspendCommand, opts ...grpc.CallOption) (*ServiceResponse, error)
	RunRunCommand(ctx context.Context, in *RunCommand, opts ...grpc.CallOption) (*ServiceResponse, error)
	RunForceCommand(ctx context.Context, in *ForceCommand, opts ...grpc.CallOption) (*ServiceResponse, error)
	RunShowRequest(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error)
	RunPingRequest(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type taklerServerClient struct {
	cc grpc.ClientConnInterface
}

func NewTaklerServerClient(cc grpc.ClientConnInterface) TaklerServerClient {
	return &taklerServerClient{cc}
}

func (c *taklerServerClient) RunInitCommand(ctx context.Context, in *InitCommand, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunInitCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunCompleteCommand(ctx context.Context, in *CompleteCommand, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunCompleteCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunAbortCommand(ctx context.Context, in *AbortCommand, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunAbortCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunEventCommand(ctx context.Context, in *EventCommand, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunEventCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunMeterCommand(ctx context.Context, in *MeterCommand, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunMeterCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunRequeueCommand(ctx context.Context, in *RequeueCommand, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunRequeueCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunSuspendCommand(ctx context.Context, in *SuspendCommand, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunSuspendCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunResumeCommand(ctx context.Context, in *SuspendCommand, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunResumeCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunRunCommand(ctx context.Context, in *RunCommand, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunRunCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunForceCommand(ctx context.Context, in *ForceCommand, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunForceCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunShowRequest(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (*ShowResponse, error) {
	out := new(ShowResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunShowRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taklerServerClient) RunPingRequest(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/takler_protocol.TaklerServer/RunPingRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaklerServerServer is the server API for TaklerServer service.
// All implementations must embed UnimplementedTaklerServerServer
// for forward compatibility
type TaklerServerServer interface {
	// child command
	RunInitCommand(context.Context, *InitCommand) (*ServiceResponse, error)
	RunCompleteCommand(context.Context, *CompleteCommand) (*ServiceResponse, error)
	RunAbortCommand(context.Context, *AbortCommand) (*ServiceResponse, error)
	RunEventCommand(context.Context, *EventCommand) (*ServiceResponse, error)
	RunMeterCommand(context.Context, *MeterCommand) (*ServiceResponse, error)
	RunRequeueCommand(context.Context, *RequeueCommand) (*ServiceResponse, error)
	RunSuspendCommand(context.Context, *SuspendCommand) (*ServiceResponse, error)
	RunResumeCommand(context.Context, *SuspendCommand) (*ServiceResponse, error)
	RunRunCommand(context.Context, *RunCommand) (*ServiceResponse, error)
	RunForceCommand(context.Context, *ForceCommand) (*ServiceResponse, error)
	RunShowRequest(context.Context, *ShowRequest) (*ShowResponse, error)
	RunPingRequest(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedTaklerServerServer()
}

// UnimplementedTaklerServerServer must be embedded to have forward compatible implementations.
type UnimplementedTaklerServerServer struct {
}

func (UnimplementedTaklerServerServer) RunInitCommand(context.Context, *InitCommand) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunInitCommand not implemented")
}
func (UnimplementedTaklerServerServer) RunCompleteCommand(context.Context, *CompleteCommand) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCompleteCommand not implemented")
}
func (UnimplementedTaklerServerServer) RunAbortCommand(context.Context, *AbortCommand) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunAbortCommand not implemented")
}
func (UnimplementedTaklerServerServer) RunEventCommand(context.Context, *EventCommand) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunEventCommand not implemented")
}
func (UnimplementedTaklerServerServer) RunMeterCommand(context.Context, *MeterCommand) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunMeterCommand not implemented")
}
func (UnimplementedTaklerServerServer) RunRequeueCommand(context.Context, *RequeueCommand) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunRequeueCommand not implemented")
}
func (UnimplementedTaklerServerServer) RunSuspendCommand(context.Context, *SuspendCommand) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunSuspendCommand not implemented")
}
func (UnimplementedTaklerServerServer) RunResumeCommand(context.Context, *SuspendCommand) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunResumeCommand not implemented")
}
func (UnimplementedTaklerServerServer) RunRunCommand(context.Context, *RunCommand) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunRunCommand not implemented")
}
func (UnimplementedTaklerServerServer) RunForceCommand(context.Context, *ForceCommand) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunForceCommand not implemented")
}
func (UnimplementedTaklerServerServer) RunShowRequest(context.Context, *ShowRequest) (*ShowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunShowRequest not implemented")
}
func (UnimplementedTaklerServerServer) RunPingRequest(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunPingRequest not implemented")
}
func (UnimplementedTaklerServerServer) mustEmbedUnimplementedTaklerServerServer() {}

// UnsafeTaklerServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaklerServerServer will
// result in compilation errors.
type UnsafeTaklerServerServer interface {
	mustEmbedUnimplementedTaklerServerServer()
}

func RegisterTaklerServerServer(s grpc.ServiceRegistrar, srv TaklerServerServer) {
	s.RegisterService(&TaklerServer_ServiceDesc, srv)
}

func _TaklerServer_RunInitCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunInitCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunInitCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunInitCommand(ctx, req.(*InitCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunCompleteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunCompleteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunCompleteCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunCompleteCommand(ctx, req.(*CompleteCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunAbortCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbortCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunAbortCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunAbortCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunAbortCommand(ctx, req.(*AbortCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunEventCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunEventCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunEventCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunEventCommand(ctx, req.(*EventCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunMeterCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeterCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunMeterCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunMeterCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunMeterCommand(ctx, req.(*MeterCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunRequeueCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequeueCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunRequeueCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunRequeueCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunRequeueCommand(ctx, req.(*RequeueCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunSuspendCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuspendCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunSuspendCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunSuspendCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunSuspendCommand(ctx, req.(*SuspendCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunResumeCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuspendCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunResumeCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunResumeCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunResumeCommand(ctx, req.(*SuspendCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunRunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunRunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunRunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunRunCommand(ctx, req.(*RunCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunForceCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForceCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunForceCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunForceCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunForceCommand(ctx, req.(*ForceCommand))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunShowRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunShowRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunShowRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunShowRequest(ctx, req.(*ShowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaklerServer_RunPingRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaklerServerServer).RunPingRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/takler_protocol.TaklerServer/RunPingRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaklerServerServer).RunPingRequest(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaklerServer_ServiceDesc is the grpc.ServiceDesc for TaklerServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaklerServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "takler_protocol.TaklerServer",
	HandlerType: (*TaklerServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunInitCommand",
			Handler:    _TaklerServer_RunInitCommand_Handler,
		},
		{
			MethodName: "RunCompleteCommand",
			Handler:    _TaklerServer_RunCompleteCommand_Handler,
		},
		{
			MethodName: "RunAbortCommand",
			Handler:    _TaklerServer_RunAbortCommand_Handler,
		},
		{
			MethodName: "RunEventCommand",
			Handler:    _TaklerServer_RunEventCommand_Handler,
		},
		{
			MethodName: "RunMeterCommand",
			Handler:    _TaklerServer_RunMeterCommand_Handler,
		},
		{
			MethodName: "RunRequeueCommand",
			Handler:    _TaklerServer_RunRequeueCommand_Handler,
		},
		{
			MethodName: "RunSuspendCommand",
			Handler:    _TaklerServer_RunSuspendCommand_Handler,
		},
		{
			MethodName: "RunResumeCommand",
			Handler:    _TaklerServer_RunResumeCommand_Handler,
		},
		{
			MethodName: "RunRunCommand",
			Handler:    _TaklerServer_RunRunCommand_Handler,
		},
		{
			MethodName: "RunForceCommand",
			Handler:    _TaklerServer_RunForceCommand_Handler,
		},
		{
			MethodName: "RunShowRequest",
			Handler:    _TaklerServer_RunShowRequest_Handler,
		},
		{
			MethodName: "RunPingRequest",
			Handler:    _TaklerServer_RunPingRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "takler_protocol/takler.proto",
}
