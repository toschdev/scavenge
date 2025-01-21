// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scavenge/scavenge/query.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6197257e32ffca5f, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6197257e32ffca5f, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryListQuestionsRequest struct {
}

func (m *QueryListQuestionsRequest) Reset()         { *m = QueryListQuestionsRequest{} }
func (m *QueryListQuestionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryListQuestionsRequest) ProtoMessage()    {}
func (*QueryListQuestionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6197257e32ffca5f, []int{2}
}
func (m *QueryListQuestionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListQuestionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListQuestionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListQuestionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListQuestionsRequest.Merge(m, src)
}
func (m *QueryListQuestionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryListQuestionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListQuestionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListQuestionsRequest proto.InternalMessageInfo

type QueryListQuestionsResponse struct {
}

func (m *QueryListQuestionsResponse) Reset()         { *m = QueryListQuestionsResponse{} }
func (m *QueryListQuestionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryListQuestionsResponse) ProtoMessage()    {}
func (*QueryListQuestionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6197257e32ffca5f, []int{3}
}
func (m *QueryListQuestionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryListQuestionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryListQuestionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryListQuestionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryListQuestionsResponse.Merge(m, src)
}
func (m *QueryListQuestionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryListQuestionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryListQuestionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryListQuestionsResponse proto.InternalMessageInfo

type QueryShowQuestionRequest struct {
	QuestionId uint64 `protobuf:"varint,1,opt,name=questionId,proto3" json:"questionId,omitempty"`
}

func (m *QueryShowQuestionRequest) Reset()         { *m = QueryShowQuestionRequest{} }
func (m *QueryShowQuestionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryShowQuestionRequest) ProtoMessage()    {}
func (*QueryShowQuestionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6197257e32ffca5f, []int{4}
}
func (m *QueryShowQuestionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryShowQuestionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryShowQuestionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryShowQuestionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryShowQuestionRequest.Merge(m, src)
}
func (m *QueryShowQuestionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryShowQuestionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryShowQuestionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryShowQuestionRequest proto.InternalMessageInfo

func (m *QueryShowQuestionRequest) GetQuestionId() uint64 {
	if m != nil {
		return m.QuestionId
	}
	return 0
}

type QueryShowQuestionResponse struct {
}

func (m *QueryShowQuestionResponse) Reset()         { *m = QueryShowQuestionResponse{} }
func (m *QueryShowQuestionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryShowQuestionResponse) ProtoMessage()    {}
func (*QueryShowQuestionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6197257e32ffca5f, []int{5}
}
func (m *QueryShowQuestionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryShowQuestionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryShowQuestionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryShowQuestionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryShowQuestionResponse.Merge(m, src)
}
func (m *QueryShowQuestionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryShowQuestionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryShowQuestionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryShowQuestionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "scavenge.scavenge.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "scavenge.scavenge.QueryParamsResponse")
	proto.RegisterType((*QueryListQuestionsRequest)(nil), "scavenge.scavenge.QueryListQuestionsRequest")
	proto.RegisterType((*QueryListQuestionsResponse)(nil), "scavenge.scavenge.QueryListQuestionsResponse")
	proto.RegisterType((*QueryShowQuestionRequest)(nil), "scavenge.scavenge.QueryShowQuestionRequest")
	proto.RegisterType((*QueryShowQuestionResponse)(nil), "scavenge.scavenge.QueryShowQuestionResponse")
}

func init() { proto.RegisterFile("scavenge/scavenge/query.proto", fileDescriptor_6197257e32ffca5f) }

var fileDescriptor_6197257e32ffca5f = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x4e, 0x4e, 0x2c,
	0x4b, 0xcd, 0x4b, 0x4f, 0xd5, 0x87, 0x33, 0x0a, 0x4b, 0x53, 0x8b, 0x2a, 0xf5, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x04, 0x61, 0xa2, 0x7a, 0x30, 0x86, 0x94, 0x60, 0x62, 0x6e, 0x66, 0x5e, 0xbe,
	0x3e, 0x98, 0x84, 0xa8, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0xa8,
	0xa8, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e,
	0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x54, 0x56, 0x2b, 0x39, 0xbf, 0x38, 0x37, 0xbf,
	0x58, 0x3f, 0x29, 0xb1, 0x18, 0x6a, 0xa5, 0x7e, 0x99, 0x61, 0x52, 0x6a, 0x49, 0xa2, 0xa1, 0x7e,
	0x41, 0x62, 0x7a, 0x66, 0x1e, 0x58, 0x31, 0x54, 0xad, 0x1c, 0xa6, 0x23, 0x0b, 0x12, 0x8b, 0x12,
	0x73, 0xa1, 0x66, 0x29, 0x89, 0x70, 0x09, 0x05, 0x82, 0x4c, 0x08, 0x00, 0x0b, 0x06, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x28, 0x05, 0x73, 0x09, 0xa3, 0x88, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x0a, 0xd9, 0x70, 0xb1, 0x41, 0x34, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x49, 0xea, 0x61,
	0xf8, 0x51, 0x0f, 0xa2, 0xc5, 0x89, 0xf3, 0xc4, 0x3d, 0x79, 0x86, 0x15, 0xcf, 0x37, 0x68, 0x31,
	0x06, 0x41, 0xf5, 0x28, 0x49, 0x73, 0x49, 0x82, 0x0d, 0xf5, 0xc9, 0x2c, 0x2e, 0x09, 0x04, 0x59,
	0x03, 0xf2, 0x12, 0xcc, 0x46, 0x19, 0x2e, 0x29, 0x6c, 0x92, 0x10, 0x8b, 0x95, 0xac, 0xb8, 0x24,
	0xc0, 0xb2, 0xc1, 0x19, 0xf9, 0xe5, 0x30, 0x59, 0xa8, 0x4e, 0x21, 0x39, 0x2e, 0xae, 0x42, 0xa8,
	0x90, 0x67, 0x0a, 0xd8, 0x61, 0x2c, 0x41, 0x48, 0x22, 0x70, 0x6b, 0x51, 0xf5, 0x42, 0x0c, 0x36,
	0x3a, 0xc0, 0xcc, 0xc5, 0x0a, 0x96, 0x15, 0xaa, 0xe2, 0x62, 0x83, 0x38, 0x5d, 0x48, 0x15, 0x8b,
	0xaf, 0x30, 0xc3, 0x48, 0x4a, 0x8d, 0x90, 0x32, 0xa8, 0xdb, 0x15, 0x9b, 0x2e, 0x3f, 0x99, 0xcc,
	0x24, 0x2d, 0x24, 0xa9, 0x8f, 0x2b, 0x2a, 0x84, 0xa6, 0x33, 0x72, 0xf1, 0xa2, 0x78, 0x5c, 0x48,
	0x07, 0x97, 0xe1, 0xd8, 0x02, 0x4f, 0x4a, 0x97, 0x48, 0xd5, 0x50, 0x17, 0x69, 0x82, 0x5d, 0xa4,
	0x2c, 0xa4, 0x88, 0xc5, 0x45, 0x39, 0x99, 0xc5, 0x25, 0xf1, 0x85, 0x70, 0x77, 0x2c, 0x60, 0xe4,
	0xe2, 0x41, 0x0e, 0x38, 0x21, 0x6d, 0x5c, 0x56, 0x61, 0x89, 0x1a, 0x29, 0x1d, 0xe2, 0x14, 0x43,
	0x9d, 0x65, 0x0a, 0x76, 0x96, 0xbe, 0x90, 0x2e, 0x16, 0x67, 0x15, 0x67, 0xe4, 0x97, 0xc3, 0x9d,
	0xa5, 0x5f, 0x8d, 0x88, 0xde, 0x5a, 0x27, 0xe3, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96,
	0x63, 0x88, 0x92, 0x84, 0x6b, 0xaf, 0x40, 0x98, 0x54, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06,
	0x4e, 0xfd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0xec, 0x71, 0xe8, 0xc4, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of ListQuestions items.
	ListQuestions(ctx context.Context, in *QueryListQuestionsRequest, opts ...grpc.CallOption) (*QueryListQuestionsResponse, error)
	// Queries a list of ShowQuestion items.
	ShowQuestion(ctx context.Context, in *QueryShowQuestionRequest, opts ...grpc.CallOption) (*QueryShowQuestionResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/scavenge.scavenge.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListQuestions(ctx context.Context, in *QueryListQuestionsRequest, opts ...grpc.CallOption) (*QueryListQuestionsResponse, error) {
	out := new(QueryListQuestionsResponse)
	err := c.cc.Invoke(ctx, "/scavenge.scavenge.Query/ListQuestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ShowQuestion(ctx context.Context, in *QueryShowQuestionRequest, opts ...grpc.CallOption) (*QueryShowQuestionResponse, error) {
	out := new(QueryShowQuestionResponse)
	err := c.cc.Invoke(ctx, "/scavenge.scavenge.Query/ShowQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of ListQuestions items.
	ListQuestions(context.Context, *QueryListQuestionsRequest) (*QueryListQuestionsResponse, error)
	// Queries a list of ShowQuestion items.
	ShowQuestion(context.Context, *QueryShowQuestionRequest) (*QueryShowQuestionResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ListQuestions(ctx context.Context, req *QueryListQuestionsRequest) (*QueryListQuestionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuestions not implemented")
}
func (*UnimplementedQueryServer) ShowQuestion(ctx context.Context, req *QueryShowQuestionRequest) (*QueryShowQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowQuestion not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
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
		FullMethod: "/scavenge.scavenge.Query/Params",
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
		FullMethod: "/scavenge.scavenge.Query/ListQuestions",
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
		FullMethod: "/scavenge.scavenge.Query/ShowQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ShowQuestion(ctx, req.(*QueryShowQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var Query_serviceDesc = _Query_serviceDesc
var _Query_serviceDesc = grpc.ServiceDesc{
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scavenge/scavenge/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryListQuestionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListQuestionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListQuestionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryListQuestionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryListQuestionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryListQuestionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryShowQuestionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryShowQuestionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryShowQuestionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.QuestionId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.QuestionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryShowQuestionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryShowQuestionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryShowQuestionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryListQuestionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryListQuestionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryShowQuestionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QuestionId != 0 {
		n += 1 + sovQuery(uint64(m.QuestionId))
	}
	return n
}

func (m *QueryShowQuestionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListQuestionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListQuestionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListQuestionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryListQuestionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryListQuestionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryListQuestionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryShowQuestionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryShowQuestionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryShowQuestionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuestionId", wireType)
			}
			m.QuestionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuestionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryShowQuestionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryShowQuestionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryShowQuestionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
