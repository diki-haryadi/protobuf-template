// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: go-micro-template/article/v1/article.proto

package articlev1

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

// ArticleServiceClient is the client API for ArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticleServiceClient interface {
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error)
	GetArticleById(ctx context.Context, in *GetArticleByIdRequest, opts ...grpc.CallOption) (*GetArticleByIdResponse, error)
}

type articleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArticleServiceClient(cc grpc.ClientConnInterface) ArticleServiceClient {
	return &articleServiceClient{cc}
}

func (c *articleServiceClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	out := new(CreateArticleResponse)
	err := c.cc.Invoke(ctx, "/go-micro-template.article.v1.ArticleService/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleServiceClient) GetArticleById(ctx context.Context, in *GetArticleByIdRequest, opts ...grpc.CallOption) (*GetArticleByIdResponse, error) {
	out := new(GetArticleByIdResponse)
	err := c.cc.Invoke(ctx, "/go-micro-template.article.v1.ArticleService/GetArticleById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticleServiceServer is the server API for ArticleService service.
// All implementations should embed UnimplementedArticleServiceServer
// for forward compatibility
type ArticleServiceServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error)
	GetArticleById(context.Context, *GetArticleByIdRequest) (*GetArticleByIdResponse, error)
}

// UnimplementedArticleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedArticleServiceServer struct {
}

func (UnimplementedArticleServiceServer) CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedArticleServiceServer) GetArticleById(context.Context, *GetArticleByIdRequest) (*GetArticleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleById not implemented")
}

// UnsafeArticleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticleServiceServer will
// result in compilation errors.
type UnsafeArticleServiceServer interface {
	mustEmbedUnimplementedArticleServiceServer()
}

func RegisterArticleServiceServer(s grpc.ServiceRegistrar, srv ArticleServiceServer) {
	s.RegisterService(&ArticleService_ServiceDesc, srv)
}

func _ArticleService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go-micro-template.article.v1.ArticleService/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArticleService_GetArticleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleServiceServer).GetArticleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go-micro-template.article.v1.ArticleService/GetArticleById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleServiceServer).GetArticleById(ctx, req.(*GetArticleByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticleService_ServiceDesc is the grpc.ServiceDesc for ArticleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go-micro-template.article.v1.ArticleService",
	HandlerType: (*ArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _ArticleService_CreateArticle_Handler,
		},
		{
			MethodName: "GetArticleById",
			Handler:    _ArticleService_GetArticleById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go-micro-template/article/v1/article.proto",
}
