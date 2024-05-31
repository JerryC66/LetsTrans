// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: document_processor.proto

package document_processor

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

// DocumentProcessorClient is the client API for DocumentProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentProcessorClient interface {
	ProcessDocument(ctx context.Context, in *ProcessDocumentRequest, opts ...grpc.CallOption) (*ProcessDocumentResponse, error)
	ModifyDocument(ctx context.Context, in *ModifyDocumentRequest, opts ...grpc.CallOption) (*ModifyDocumentResponse, error)
}

type documentProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentProcessorClient(cc grpc.ClientConnInterface) DocumentProcessorClient {
	return &documentProcessorClient{cc}
}

func (c *documentProcessorClient) ProcessDocument(ctx context.Context, in *ProcessDocumentRequest, opts ...grpc.CallOption) (*ProcessDocumentResponse, error) {
	out := new(ProcessDocumentResponse)
	err := c.cc.Invoke(ctx, "/document_processor.DocumentProcessor/ProcessDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentProcessorClient) ModifyDocument(ctx context.Context, in *ModifyDocumentRequest, opts ...grpc.CallOption) (*ModifyDocumentResponse, error) {
	out := new(ModifyDocumentResponse)
	err := c.cc.Invoke(ctx, "/document_processor.DocumentProcessor/ModifyDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentProcessorServer is the server API for DocumentProcessor service.
// All implementations must embed UnimplementedDocumentProcessorServer
// for forward compatibility
type DocumentProcessorServer interface {
	ProcessDocument(context.Context, *ProcessDocumentRequest) (*ProcessDocumentResponse, error)
	ModifyDocument(context.Context, *ModifyDocumentRequest) (*ModifyDocumentResponse, error)
	mustEmbedUnimplementedDocumentProcessorServer()
}

// UnimplementedDocumentProcessorServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentProcessorServer struct {
}

func (UnimplementedDocumentProcessorServer) ProcessDocument(context.Context, *ProcessDocumentRequest) (*ProcessDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessDocument not implemented")
}
func (UnimplementedDocumentProcessorServer) ModifyDocument(context.Context, *ModifyDocumentRequest) (*ModifyDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyDocument not implemented")
}
func (UnimplementedDocumentProcessorServer) mustEmbedUnimplementedDocumentProcessorServer() {}

// UnsafeDocumentProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentProcessorServer will
// result in compilation errors.
type UnsafeDocumentProcessorServer interface {
	mustEmbedUnimplementedDocumentProcessorServer()
}

func RegisterDocumentProcessorServer(s grpc.ServiceRegistrar, srv DocumentProcessorServer) {
	s.RegisterService(&DocumentProcessor_ServiceDesc, srv)
}

func _DocumentProcessor_ProcessDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentProcessorServer).ProcessDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_processor.DocumentProcessor/ProcessDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentProcessorServer).ProcessDocument(ctx, req.(*ProcessDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentProcessor_ModifyDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentProcessorServer).ModifyDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document_processor.DocumentProcessor/ModifyDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentProcessorServer).ModifyDocument(ctx, req.(*ModifyDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentProcessor_ServiceDesc is the grpc.ServiceDesc for DocumentProcessor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentProcessor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "document_processor.DocumentProcessor",
	HandlerType: (*DocumentProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessDocument",
			Handler:    _DocumentProcessor_ProcessDocument_Handler,
		},
		{
			MethodName: "ModifyDocument",
			Handler:    _DocumentProcessor_ModifyDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "document_processor.proto",
}