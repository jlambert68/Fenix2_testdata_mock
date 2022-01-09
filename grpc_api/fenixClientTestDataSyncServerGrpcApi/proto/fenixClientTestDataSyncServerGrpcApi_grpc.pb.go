// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// FenixClientTestDataGrpcServicesClient is the client API for FenixClientTestDataGrpcServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FenixClientTestDataGrpcServicesClient interface {
	//Fenix Server can check if Fenix Client Testdata sync server is alive with this service
	AreYouAlive(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix Server asks Fenix client to register itself with the Fenix Testdata sync server
	RegisterTestDataClient(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix Server asks Fenix client to send TestData MerkleHash to Fenix Testdata sync server with this service
	SendMerkleHash(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix Server asks Fenix client to send TestData MerkleTree to Fenix Testdata sync server with this service
	SendMerkleTree(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix Server asks Fenix client to send TestDataHeaders to Fenix Testdata sync server with this service
	SendTestDataHeaders(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error)
	// Fenix Server asks Fenix client to  send TestData rows, based on list of MerklePaths, to Fenix Testdata sync server with this service
	SendTestDataRows(ctx context.Context, in *MerklePathsMessage, opts ...grpc.CallOption) (*AckNackResponse, error)
}

type fenixClientTestDataGrpcServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewFenixClientTestDataGrpcServicesClient(cc grpc.ClientConnInterface) FenixClientTestDataGrpcServicesClient {
	return &fenixClientTestDataGrpcServicesClient{cc}
}

func (c *fenixClientTestDataGrpcServicesClient) AreYouAlive(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/AreYouAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixClientTestDataGrpcServicesClient) RegisterTestDataClient(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/RegisterTestDataClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixClientTestDataGrpcServicesClient) SendMerkleHash(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/SendMerkleHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixClientTestDataGrpcServicesClient) SendMerkleTree(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/SendMerkleTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixClientTestDataGrpcServicesClient) SendTestDataHeaders(ctx context.Context, in *EmptyParameter, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/SendTestDataHeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fenixClientTestDataGrpcServicesClient) SendTestDataRows(ctx context.Context, in *MerklePathsMessage, opts ...grpc.CallOption) (*AckNackResponse, error) {
	out := new(AckNackResponse)
	err := c.cc.Invoke(ctx, "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/SendTestDataRows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FenixClientTestDataGrpcServicesServer is the server API for FenixClientTestDataGrpcServices service.
// All implementations must embed UnimplementedFenixClientTestDataGrpcServicesServer
// for forward compatibility
type FenixClientTestDataGrpcServicesServer interface {
	//Fenix Server can check if Fenix Client Testdata sync server is alive with this service
	AreYouAlive(context.Context, *EmptyParameter) (*AckNackResponse, error)
	// Fenix Server asks Fenix client to register itself with the Fenix Testdata sync server
	RegisterTestDataClient(context.Context, *EmptyParameter) (*AckNackResponse, error)
	// Fenix Server asks Fenix client to send TestData MerkleHash to Fenix Testdata sync server with this service
	SendMerkleHash(context.Context, *EmptyParameter) (*AckNackResponse, error)
	// Fenix Server asks Fenix client to send TestData MerkleTree to Fenix Testdata sync server with this service
	SendMerkleTree(context.Context, *EmptyParameter) (*AckNackResponse, error)
	// Fenix Server asks Fenix client to send TestDataHeaders to Fenix Testdata sync server with this service
	SendTestDataHeaders(context.Context, *EmptyParameter) (*AckNackResponse, error)
	// Fenix Server asks Fenix client to  send TestData rows, based on list of MerklePaths, to Fenix Testdata sync server with this service
	SendTestDataRows(context.Context, *MerklePathsMessage) (*AckNackResponse, error)
	mustEmbedUnimplementedFenixClientTestDataGrpcServicesServer()
}

// UnimplementedFenixClientTestDataGrpcServicesServer must be embedded to have forward compatible implementations.
type UnimplementedFenixClientTestDataGrpcServicesServer struct {
}

func (UnimplementedFenixClientTestDataGrpcServicesServer) AreYouAlive(context.Context, *EmptyParameter) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AreYouAlive not implemented")
}
func (UnimplementedFenixClientTestDataGrpcServicesServer) RegisterTestDataClient(context.Context, *EmptyParameter) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTestDataClient not implemented")
}
func (UnimplementedFenixClientTestDataGrpcServicesServer) SendMerkleHash(context.Context, *EmptyParameter) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMerkleHash not implemented")
}
func (UnimplementedFenixClientTestDataGrpcServicesServer) SendMerkleTree(context.Context, *EmptyParameter) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMerkleTree not implemented")
}
func (UnimplementedFenixClientTestDataGrpcServicesServer) SendTestDataHeaders(context.Context, *EmptyParameter) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTestDataHeaders not implemented")
}
func (UnimplementedFenixClientTestDataGrpcServicesServer) SendTestDataRows(context.Context, *MerklePathsMessage) (*AckNackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTestDataRows not implemented")
}
func (UnimplementedFenixClientTestDataGrpcServicesServer) mustEmbedUnimplementedFenixClientTestDataGrpcServicesServer() {
}

// UnsafeFenixClientTestDataGrpcServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FenixClientTestDataGrpcServicesServer will
// result in compilation errors.
type UnsafeFenixClientTestDataGrpcServicesServer interface {
	mustEmbedUnimplementedFenixClientTestDataGrpcServicesServer()
}

func RegisterFenixClientTestDataGrpcServicesServer(s grpc.ServiceRegistrar, srv FenixClientTestDataGrpcServicesServer) {
	s.RegisterService(&FenixClientTestDataGrpcServices_ServiceDesc, srv)
}

func _FenixClientTestDataGrpcServices_AreYouAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixClientTestDataGrpcServicesServer).AreYouAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/AreYouAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixClientTestDataGrpcServicesServer).AreYouAlive(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixClientTestDataGrpcServices_RegisterTestDataClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixClientTestDataGrpcServicesServer).RegisterTestDataClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/RegisterTestDataClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixClientTestDataGrpcServicesServer).RegisterTestDataClient(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixClientTestDataGrpcServices_SendMerkleHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixClientTestDataGrpcServicesServer).SendMerkleHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/SendMerkleHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixClientTestDataGrpcServicesServer).SendMerkleHash(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixClientTestDataGrpcServices_SendMerkleTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixClientTestDataGrpcServicesServer).SendMerkleTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/SendMerkleTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixClientTestDataGrpcServicesServer).SendMerkleTree(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixClientTestDataGrpcServices_SendTestDataHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixClientTestDataGrpcServicesServer).SendTestDataHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/SendTestDataHeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixClientTestDataGrpcServicesServer).SendTestDataHeaders(ctx, req.(*EmptyParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _FenixClientTestDataGrpcServices_SendTestDataRows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MerklePathsMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FenixClientTestDataGrpcServicesServer).SendTestDataRows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices/SendTestDataRows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FenixClientTestDataGrpcServicesServer).SendTestDataRows(ctx, req.(*MerklePathsMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// FenixClientTestDataGrpcServices_ServiceDesc is the grpc.ServiceDesc for FenixClientTestDataGrpcServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FenixClientTestDataGrpcServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fenixClientTestDataSyncServerGrpcApi.FenixClientTestDataGrpcServices",
	HandlerType: (*FenixClientTestDataGrpcServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AreYouAlive",
			Handler:    _FenixClientTestDataGrpcServices_AreYouAlive_Handler,
		},
		{
			MethodName: "RegisterTestDataClient",
			Handler:    _FenixClientTestDataGrpcServices_RegisterTestDataClient_Handler,
		},
		{
			MethodName: "SendMerkleHash",
			Handler:    _FenixClientTestDataGrpcServices_SendMerkleHash_Handler,
		},
		{
			MethodName: "SendMerkleTree",
			Handler:    _FenixClientTestDataGrpcServices_SendMerkleTree_Handler,
		},
		{
			MethodName: "SendTestDataHeaders",
			Handler:    _FenixClientTestDataGrpcServices_SendTestDataHeaders_Handler,
		},
		{
			MethodName: "SendTestDataRows",
			Handler:    _FenixClientTestDataGrpcServices_SendTestDataRows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fenixClientTestDataSyncServerGrpcApi.proto",
}
