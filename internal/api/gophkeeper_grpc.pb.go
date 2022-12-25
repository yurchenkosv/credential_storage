// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.2
// source: gophkeeper.proto

package api

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

// CredentialServiceClient is the client API for CredentialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CredentialServiceClient interface {
	SaveCredentialsData(ctx context.Context, in *CredentialsData, opts ...grpc.CallOption) (*ServerResponse, error)
	SaveBankingData(ctx context.Context, in *BankingCardData, opts ...grpc.CallOption) (*ServerResponse, error)
	SaveTextData(ctx context.Context, in *TextData, opts ...grpc.CallOption) (*ServerResponse, error)
	SaveBinaryData(ctx context.Context, in *BinaryData, opts ...grpc.CallOption) (*ServerResponse, error)
	GetCredentialsData(ctx context.Context, in *CredentialsDataRequest, opts ...grpc.CallOption) (*CredentialsData, error)
	GetBankingCardData(ctx context.Context, in *BankingCardDataRequest, opts ...grpc.CallOption) (*BankingCardData, error)
	GetTextData(ctx context.Context, in *TextDataRequest, opts ...grpc.CallOption) (*TextData, error)
	GetBinaryData(ctx context.Context, in *BinaryDataRequest, opts ...grpc.CallOption) (*BinaryData, error)
}

type credentialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCredentialServiceClient(cc grpc.ClientConnInterface) CredentialServiceClient {
	return &credentialServiceClient{cc}
}

func (c *credentialServiceClient) SaveCredentialsData(ctx context.Context, in *CredentialsData, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/api.CredentialService/SaveCredentialsData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialServiceClient) SaveBankingData(ctx context.Context, in *BankingCardData, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/api.CredentialService/SaveBankingData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialServiceClient) SaveTextData(ctx context.Context, in *TextData, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/api.CredentialService/SaveTextData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialServiceClient) SaveBinaryData(ctx context.Context, in *BinaryData, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/api.CredentialService/SaveBinaryData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialServiceClient) GetCredentialsData(ctx context.Context, in *CredentialsDataRequest, opts ...grpc.CallOption) (*CredentialsData, error) {
	out := new(CredentialsData)
	err := c.cc.Invoke(ctx, "/api.CredentialService/GetCredentialsData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialServiceClient) GetBankingCardData(ctx context.Context, in *BankingCardDataRequest, opts ...grpc.CallOption) (*BankingCardData, error) {
	out := new(BankingCardData)
	err := c.cc.Invoke(ctx, "/api.CredentialService/GetBankingCardData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialServiceClient) GetTextData(ctx context.Context, in *TextDataRequest, opts ...grpc.CallOption) (*TextData, error) {
	out := new(TextData)
	err := c.cc.Invoke(ctx, "/api.CredentialService/GetTextData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialServiceClient) GetBinaryData(ctx context.Context, in *BinaryDataRequest, opts ...grpc.CallOption) (*BinaryData, error) {
	out := new(BinaryData)
	err := c.cc.Invoke(ctx, "/api.CredentialService/GetBinaryData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CredentialServiceServer is the server API for CredentialService service.
// All implementations should embed UnimplementedCredentialServiceServer
// for forward compatibility
type CredentialServiceServer interface {
	SaveCredentialsData(context.Context, *CredentialsData) (*ServerResponse, error)
	SaveBankingData(context.Context, *BankingCardData) (*ServerResponse, error)
	SaveTextData(context.Context, *TextData) (*ServerResponse, error)
	SaveBinaryData(context.Context, *BinaryData) (*ServerResponse, error)
	GetCredentialsData(context.Context, *CredentialsDataRequest) (*CredentialsData, error)
	GetBankingCardData(context.Context, *BankingCardDataRequest) (*BankingCardData, error)
	GetTextData(context.Context, *TextDataRequest) (*TextData, error)
	GetBinaryData(context.Context, *BinaryDataRequest) (*BinaryData, error)
}

// UnimplementedCredentialServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCredentialServiceServer struct {
}

func (UnimplementedCredentialServiceServer) SaveCredentialsData(context.Context, *CredentialsData) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCredentialsData not implemented")
}
func (UnimplementedCredentialServiceServer) SaveBankingData(context.Context, *BankingCardData) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBankingData not implemented")
}
func (UnimplementedCredentialServiceServer) SaveTextData(context.Context, *TextData) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTextData not implemented")
}
func (UnimplementedCredentialServiceServer) SaveBinaryData(context.Context, *BinaryData) (*ServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBinaryData not implemented")
}
func (UnimplementedCredentialServiceServer) GetCredentialsData(context.Context, *CredentialsDataRequest) (*CredentialsData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCredentialsData not implemented")
}
func (UnimplementedCredentialServiceServer) GetBankingCardData(context.Context, *BankingCardDataRequest) (*BankingCardData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBankingCardData not implemented")
}
func (UnimplementedCredentialServiceServer) GetTextData(context.Context, *TextDataRequest) (*TextData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTextData not implemented")
}
func (UnimplementedCredentialServiceServer) GetBinaryData(context.Context, *BinaryDataRequest) (*BinaryData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBinaryData not implemented")
}

// UnsafeCredentialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CredentialServiceServer will
// result in compilation errors.
type UnsafeCredentialServiceServer interface {
	mustEmbedUnimplementedCredentialServiceServer()
}

func RegisterCredentialServiceServer(s grpc.ServiceRegistrar, srv CredentialServiceServer) {
	s.RegisterService(&CredentialService_ServiceDesc, srv)
}

func _CredentialService_SaveCredentialsData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialsData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).SaveCredentialsData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CredentialService/SaveCredentialsData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).SaveCredentialsData(ctx, req.(*CredentialsData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialService_SaveBankingData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankingCardData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).SaveBankingData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CredentialService/SaveBankingData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).SaveBankingData(ctx, req.(*BankingCardData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialService_SaveTextData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).SaveTextData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CredentialService/SaveTextData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).SaveTextData(ctx, req.(*TextData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialService_SaveBinaryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).SaveBinaryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CredentialService/SaveBinaryData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).SaveBinaryData(ctx, req.(*BinaryData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialService_GetCredentialsData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialsDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).GetCredentialsData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CredentialService/GetCredentialsData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).GetCredentialsData(ctx, req.(*CredentialsDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialService_GetBankingCardData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankingCardDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).GetBankingCardData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CredentialService/GetBankingCardData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).GetBankingCardData(ctx, req.(*BankingCardDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialService_GetTextData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).GetTextData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CredentialService/GetTextData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).GetTextData(ctx, req.(*TextDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CredentialService_GetBinaryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinaryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CredentialServiceServer).GetBinaryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CredentialService/GetBinaryData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CredentialServiceServer).GetBinaryData(ctx, req.(*BinaryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CredentialService_ServiceDesc is the grpc.ServiceDesc for CredentialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CredentialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.CredentialService",
	HandlerType: (*CredentialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveCredentialsData",
			Handler:    _CredentialService_SaveCredentialsData_Handler,
		},
		{
			MethodName: "SaveBankingData",
			Handler:    _CredentialService_SaveBankingData_Handler,
		},
		{
			MethodName: "SaveTextData",
			Handler:    _CredentialService_SaveTextData_Handler,
		},
		{
			MethodName: "SaveBinaryData",
			Handler:    _CredentialService_SaveBinaryData_Handler,
		},
		{
			MethodName: "GetCredentialsData",
			Handler:    _CredentialService_GetCredentialsData_Handler,
		},
		{
			MethodName: "GetBankingCardData",
			Handler:    _CredentialService_GetBankingCardData_Handler,
		},
		{
			MethodName: "GetTextData",
			Handler:    _CredentialService_GetTextData_Handler,
		},
		{
			MethodName: "GetBinaryData",
			Handler:    _CredentialService_GetBinaryData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gophkeeper.proto",
}