// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: scavenge/scavenge/query.proto

package scavenge

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
	Query_Params_FullMethodName        = "/scavenge.scavenge.Query/Params"
	Query_ListQuestions_FullMethodName = "/scavenge.scavenge.Query/ListQuestions"
	Query_ShowQuestion_FullMethodName  = "/scavenge.scavenge.Query/ShowQuestion"
	Query_ListCommits_FullMethodName   = "/scavenge.scavenge.Query/ListCommits"
	Query_ShowCommit_FullMethodName    = "/scavenge.scavenge.Query/ShowCommit"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of ListQuestions items.
	ListQuestions(ctx context.Context, in *QueryListQuestionsRequest, opts ...grpc.CallOption) (*QueryListQuestionsResponse, error)
	// Queries a list of ShowQuestion items.
	ShowQuestion(ctx context.Context, in *QueryShowQuestionRequest, opts ...grpc.CallOption) (*QueryShowQuestionResponse, error)
	// Queries a list of ListCommits items.
	ListCommits(ctx context.Context, in *QueryListCommitsRequest, opts ...grpc.CallOption) (*QueryListCommitsResponse, error)
	// Queries a list of ShowCommit items.
	ShowCommit(ctx context.Context, in *QueryShowCommitRequest, opts ...grpc.CallOption) (*QueryShowCommitResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListQuestions(ctx context.Context, in *QueryListQuestionsRequest, opts ...grpc.CallOption) (*QueryListQuestionsResponse, error) {
	out := new(QueryListQuestionsResponse)
	err := c.cc.Invoke(ctx, Query_ListQuestions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ShowQuestion(ctx context.Context, in *QueryShowQuestionRequest, opts ...grpc.CallOption) (*QueryShowQuestionResponse, error) {
	out := new(QueryShowQuestionResponse)
	err := c.cc.Invoke(ctx, Query_ShowQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListCommits(ctx context.Context, in *QueryListCommitsRequest, opts ...grpc.CallOption) (*QueryListCommitsResponse, error) {
	out := new(QueryListCommitsResponse)
	err := c.cc.Invoke(ctx, Query_ListCommits_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ShowCommit(ctx context.Context, in *QueryShowCommitRequest, opts ...grpc.CallOption) (*QueryShowCommitResponse, error) {
	out := new(QueryShowCommitResponse)
	err := c.cc.Invoke(ctx, Query_ShowCommit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ListQuestions items.
	ListQuestions(context.Context, *QueryListQuestionsRequest) (*QueryListQuestionsResponse, error)
	// Queries a list of ShowQuestion items.
	ShowQuestion(context.Context, *QueryShowQuestionRequest) (*QueryShowQuestionResponse, error)
	// Queries a list of ListCommits items.
	ListCommits(context.Context, *QueryListCommitsRequest) (*QueryListCommitsResponse, error)
	// Queries a list of ShowCommit items.
	ShowCommit(context.Context, *QueryShowCommitRequest) (*QueryShowCommitResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) ListQuestions(context.Context, *QueryListQuestionsRequest) (*QueryListQuestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuestions not implemented")
}
func (UnimplementedQueryServer) ShowQuestion(context.Context, *QueryShowQuestionRequest) (*QueryShowQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowQuestion not implemented")
}
func (UnimplementedQueryServer) ListCommits(context.Context, *QueryListCommitsRequest) (*QueryListCommitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommits not implemented")
}
func (UnimplementedQueryServer) ShowCommit(context.Context, *QueryShowCommitRequest) (*QueryShowCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowCommit not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListQuestionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListQuestions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListQuestions(ctx, req.(*QueryListQuestionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ShowQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryShowQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ShowQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ShowQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ShowQuestion(ctx, req.(*QueryShowQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListCommits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryListCommitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListCommits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListCommits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListCommits(ctx, req.(*QueryListCommitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ShowCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryShowCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ShowCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ShowCommit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ShowCommit(ctx, req.(*QueryShowCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scavenge.scavenge.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ListQuestions",
			Handler:    _Query_ListQuestions_Handler,
		},
		{
			MethodName: "ShowQuestion",
			Handler:    _Query_ShowQuestion_Handler,
		},
		{
			MethodName: "ListCommits",
			Handler:    _Query_ListCommits_Handler,
		},
		{
			MethodName: "ShowCommit",
			Handler:    _Query_ShowCommit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scavenge/scavenge/query.proto",
}
