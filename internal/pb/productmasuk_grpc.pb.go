// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: productmasuk.proto

package pb

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

// ProductMasukServiceClient is the client API for ProductMasukService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductMasukServiceClient interface {
	CreateProductMasuk(ctx context.Context, in *CreateProductMasukInput, opts ...grpc.CallOption) (*ProductMasukResponse, error)
	GetProductMasuks(ctx context.Context, in *ProductMasuksRequest, opts ...grpc.CallOption) (*ProductMasuksResponse, error)
	GetProductMasuk(ctx context.Context, in *ProductMasukRequest, opts ...grpc.CallOption) (*ProductMasukResponse, error)
	UpdateProductMasuk(ctx context.Context, in *UpdateProductMasukInput, opts ...grpc.CallOption) (*ProductMasukResponse, error)
	DeleteProductMasuk(ctx context.Context, in *ProductMasukRequest, opts ...grpc.CallOption) (*DeleteProductMasukResponse, error)
}

type productMasukServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductMasukServiceClient(cc grpc.ClientConnInterface) ProductMasukServiceClient {
	return &productMasukServiceClient{cc}
}

func (c *productMasukServiceClient) CreateProductMasuk(ctx context.Context, in *CreateProductMasukInput, opts ...grpc.CallOption) (*ProductMasukResponse, error) {
	out := new(ProductMasukResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductMasukService/CreateProductMasuk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productMasukServiceClient) GetProductMasuks(ctx context.Context, in *ProductMasuksRequest, opts ...grpc.CallOption) (*ProductMasuksResponse, error) {
	out := new(ProductMasuksResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductMasukService/GetProductMasuks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productMasukServiceClient) GetProductMasuk(ctx context.Context, in *ProductMasukRequest, opts ...grpc.CallOption) (*ProductMasukResponse, error) {
	out := new(ProductMasukResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductMasukService/GetProductMasuk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productMasukServiceClient) UpdateProductMasuk(ctx context.Context, in *UpdateProductMasukInput, opts ...grpc.CallOption) (*ProductMasukResponse, error) {
	out := new(ProductMasukResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductMasukService/UpdateProductMasuk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productMasukServiceClient) DeleteProductMasuk(ctx context.Context, in *ProductMasukRequest, opts ...grpc.CallOption) (*DeleteProductMasukResponse, error) {
	out := new(DeleteProductMasukResponse)
	err := c.cc.Invoke(ctx, "/pb.ProductMasukService/DeleteProductMasuk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductMasukServiceServer is the server API for ProductMasukService service.
// All implementations must embed UnimplementedProductMasukServiceServer
// for forward compatibility
type ProductMasukServiceServer interface {
	CreateProductMasuk(context.Context, *CreateProductMasukInput) (*ProductMasukResponse, error)
	GetProductMasuks(context.Context, *ProductMasuksRequest) (*ProductMasuksResponse, error)
	GetProductMasuk(context.Context, *ProductMasukRequest) (*ProductMasukResponse, error)
	UpdateProductMasuk(context.Context, *UpdateProductMasukInput) (*ProductMasukResponse, error)
	DeleteProductMasuk(context.Context, *ProductMasukRequest) (*DeleteProductMasukResponse, error)
	mustEmbedUnimplementedProductMasukServiceServer()
}

// UnimplementedProductMasukServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductMasukServiceServer struct {
}

func (UnimplementedProductMasukServiceServer) CreateProductMasuk(context.Context, *CreateProductMasukInput) (*ProductMasukResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProductMasuk not implemented")
}
func (UnimplementedProductMasukServiceServer) GetProductMasuks(context.Context, *ProductMasuksRequest) (*ProductMasuksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductMasuks not implemented")
}
func (UnimplementedProductMasukServiceServer) GetProductMasuk(context.Context, *ProductMasukRequest) (*ProductMasukResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductMasuk not implemented")
}
func (UnimplementedProductMasukServiceServer) UpdateProductMasuk(context.Context, *UpdateProductMasukInput) (*ProductMasukResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductMasuk not implemented")
}
func (UnimplementedProductMasukServiceServer) DeleteProductMasuk(context.Context, *ProductMasukRequest) (*DeleteProductMasukResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProductMasuk not implemented")
}
func (UnimplementedProductMasukServiceServer) mustEmbedUnimplementedProductMasukServiceServer() {}

// UnsafeProductMasukServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductMasukServiceServer will
// result in compilation errors.
type UnsafeProductMasukServiceServer interface {
	mustEmbedUnimplementedProductMasukServiceServer()
}

func RegisterProductMasukServiceServer(s grpc.ServiceRegistrar, srv ProductMasukServiceServer) {
	s.RegisterService(&ProductMasukService_ServiceDesc, srv)
}

func _ProductMasukService_CreateProductMasuk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductMasukInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductMasukServiceServer).CreateProductMasuk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductMasukService/CreateProductMasuk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductMasukServiceServer).CreateProductMasuk(ctx, req.(*CreateProductMasukInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductMasukService_GetProductMasuks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductMasuksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductMasukServiceServer).GetProductMasuks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductMasukService/GetProductMasuks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductMasukServiceServer).GetProductMasuks(ctx, req.(*ProductMasuksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductMasukService_GetProductMasuk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductMasukRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductMasukServiceServer).GetProductMasuk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductMasukService/GetProductMasuk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductMasukServiceServer).GetProductMasuk(ctx, req.(*ProductMasukRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductMasukService_UpdateProductMasuk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductMasukInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductMasukServiceServer).UpdateProductMasuk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductMasukService/UpdateProductMasuk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductMasukServiceServer).UpdateProductMasuk(ctx, req.(*UpdateProductMasukInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductMasukService_DeleteProductMasuk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductMasukRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductMasukServiceServer).DeleteProductMasuk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductMasukService/DeleteProductMasuk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductMasukServiceServer).DeleteProductMasuk(ctx, req.(*ProductMasukRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductMasukService_ServiceDesc is the grpc.ServiceDesc for ProductMasukService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductMasukService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProductMasukService",
	HandlerType: (*ProductMasukServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProductMasuk",
			Handler:    _ProductMasukService_CreateProductMasuk_Handler,
		},
		{
			MethodName: "GetProductMasuks",
			Handler:    _ProductMasukService_GetProductMasuks_Handler,
		},
		{
			MethodName: "GetProductMasuk",
			Handler:    _ProductMasukService_GetProductMasuk_Handler,
		},
		{
			MethodName: "UpdateProductMasuk",
			Handler:    _ProductMasukService_UpdateProductMasuk_Handler,
		},
		{
			MethodName: "DeleteProductMasuk",
			Handler:    _ProductMasukService_DeleteProductMasuk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "productmasuk.proto",
}
