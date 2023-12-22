// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v1/api.proto

package apiv1

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

const (
	ScannerService_Start_FullMethodName = "/com.facehunter.scanner.api.v1.ScannerService/Start"
)

// ScannerServiceClient is the client API for ScannerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScannerServiceClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
}

type scannerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScannerServiceClient(cc grpc.ClientConnInterface) ScannerServiceClient {
	return &scannerServiceClient{cc}
}

func (c *scannerServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, ScannerService_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScannerServiceServer is the server API for ScannerService service.
// All implementations should embed UnimplementedScannerServiceServer
// for forward compatibility
type ScannerServiceServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
}

// UnimplementedScannerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedScannerServiceServer struct {
}

func (UnimplementedScannerServiceServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}

// UnsafeScannerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScannerServiceServer will
// result in compilation errors.
type UnsafeScannerServiceServer interface {
	mustEmbedUnimplementedScannerServiceServer()
}

func RegisterScannerServiceServer(s grpc.ServiceRegistrar, srv ScannerServiceServer) {
	s.RegisterService(&ScannerService_ServiceDesc, srv)
}

func _ScannerService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScannerServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ScannerService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScannerServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ScannerService_ServiceDesc is the grpc.ServiceDesc for ScannerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScannerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.facehunter.scanner.api.v1.ScannerService",
	HandlerType: (*ScannerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _ScannerService_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api.proto",
}