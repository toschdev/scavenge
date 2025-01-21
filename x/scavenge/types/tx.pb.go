// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scavenge/scavenge/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b36a47d3ea84630, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b36a47d3ea84630, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgCreateQuestion struct {
	Creator         string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Question        string `protobuf:"bytes,2,opt,name=question,proto3" json:"question,omitempty"`
	EncryptedAnswer string `protobuf:"bytes,3,opt,name=encryptedAnswer,proto3" json:"encryptedAnswer,omitempty"`
	Bounty          uint64 `protobuf:"varint,4,opt,name=bounty,proto3" json:"bounty,omitempty"`
}

func (m *MsgCreateQuestion) Reset()         { *m = MsgCreateQuestion{} }
func (m *MsgCreateQuestion) String() string { return proto.CompactTextString(m) }
func (*MsgCreateQuestion) ProtoMessage()    {}
func (*MsgCreateQuestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b36a47d3ea84630, []int{2}
}
func (m *MsgCreateQuestion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateQuestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateQuestion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateQuestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateQuestion.Merge(m, src)
}
func (m *MsgCreateQuestion) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateQuestion) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateQuestion.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateQuestion proto.InternalMessageInfo

func (m *MsgCreateQuestion) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateQuestion) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *MsgCreateQuestion) GetEncryptedAnswer() string {
	if m != nil {
		return m.EncryptedAnswer
	}
	return ""
}

func (m *MsgCreateQuestion) GetBounty() uint64 {
	if m != nil {
		return m.Bounty
	}
	return 0
}

type MsgCreateQuestionResponse struct {
}

func (m *MsgCreateQuestionResponse) Reset()         { *m = MsgCreateQuestionResponse{} }
func (m *MsgCreateQuestionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateQuestionResponse) ProtoMessage()    {}
func (*MsgCreateQuestionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b36a47d3ea84630, []int{3}
}
func (m *MsgCreateQuestionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateQuestionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateQuestionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateQuestionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateQuestionResponse.Merge(m, src)
}
func (m *MsgCreateQuestionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateQuestionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateQuestionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateQuestionResponse proto.InternalMessageInfo

type MsgCommitAnswer struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	QuestionId uint64 `protobuf:"varint,2,opt,name=questionId,proto3" json:"questionId,omitempty"`
	HashAnswer string `protobuf:"bytes,3,opt,name=hashAnswer,proto3" json:"hashAnswer,omitempty"`
}

func (m *MsgCommitAnswer) Reset()         { *m = MsgCommitAnswer{} }
func (m *MsgCommitAnswer) String() string { return proto.CompactTextString(m) }
func (*MsgCommitAnswer) ProtoMessage()    {}
func (*MsgCommitAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b36a47d3ea84630, []int{4}
}
func (m *MsgCommitAnswer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCommitAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCommitAnswer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCommitAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCommitAnswer.Merge(m, src)
}
func (m *MsgCommitAnswer) XXX_Size() int {
	return m.Size()
}
func (m *MsgCommitAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCommitAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCommitAnswer proto.InternalMessageInfo

func (m *MsgCommitAnswer) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCommitAnswer) GetQuestionId() uint64 {
	if m != nil {
		return m.QuestionId
	}
	return 0
}

func (m *MsgCommitAnswer) GetHashAnswer() string {
	if m != nil {
		return m.HashAnswer
	}
	return ""
}

type MsgCommitAnswerResponse struct {
}

func (m *MsgCommitAnswerResponse) Reset()         { *m = MsgCommitAnswerResponse{} }
func (m *MsgCommitAnswerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCommitAnswerResponse) ProtoMessage()    {}
func (*MsgCommitAnswerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b36a47d3ea84630, []int{5}
}
func (m *MsgCommitAnswerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCommitAnswerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCommitAnswerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCommitAnswerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCommitAnswerResponse.Merge(m, src)
}
func (m *MsgCommitAnswerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCommitAnswerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCommitAnswerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCommitAnswerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "scavenge.scavenge.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "scavenge.scavenge.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgCreateQuestion)(nil), "scavenge.scavenge.MsgCreateQuestion")
	proto.RegisterType((*MsgCreateQuestionResponse)(nil), "scavenge.scavenge.MsgCreateQuestionResponse")
	proto.RegisterType((*MsgCommitAnswer)(nil), "scavenge.scavenge.MsgCommitAnswer")
	proto.RegisterType((*MsgCommitAnswerResponse)(nil), "scavenge.scavenge.MsgCommitAnswerResponse")
}

func init() { proto.RegisterFile("scavenge/scavenge/tx.proto", fileDescriptor_6b36a47d3ea84630) }

var fileDescriptor_6b36a47d3ea84630 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x6d, 0x08, 0xe4, 0x11, 0x51, 0xe5, 0x54, 0x51, 0xdb, 0x48, 0x26, 0x32, 0x0c,
	0x51, 0x04, 0xb1, 0x68, 0xa5, 0x0e, 0x15, 0x4b, 0xd3, 0x89, 0x21, 0x12, 0x18, 0xb1, 0x30, 0x14,
	0xb9, 0xf1, 0xc9, 0xf1, 0x60, 0x9f, 0xb9, 0xbb, 0x94, 0x7a, 0x43, 0x8c, 0x4c, 0xac, 0x7c, 0x03,
	0xc6, 0x08, 0xb1, 0xb3, 0x76, 0x8c, 0x98, 0x98, 0x10, 0x4a, 0x86, 0x7c, 0x0d, 0x64, 0xdf, 0xd9,
	0x49, 0x9c, 0x44, 0x74, 0x49, 0xee, 0xfd, 0xdf, 0xff, 0xee, 0xfd, 0xde, 0xbd, 0x33, 0x18, 0x7c,
	0xe0, 0x5e, 0x92, 0xc8, 0x27, 0x76, 0xb1, 0x10, 0x57, 0xdd, 0x98, 0x51, 0x41, 0x71, 0x33, 0x97,
	0xba, 0xf9, 0xc2, 0x68, 0xba, 0x61, 0x10, 0x51, 0x3b, 0xfb, 0x95, 0x2e, 0xe3, 0x60, 0x40, 0x79,
	0x48, 0xb9, 0x1d, 0x72, 0xdf, 0xbe, 0x7c, 0x96, 0xfe, 0xa9, 0x84, 0x2e, 0x13, 0xef, 0xb2, 0xc8,
	0x96, 0x81, 0x4a, 0xed, 0xfb, 0xd4, 0xa7, 0x52, 0x4f, 0x57, 0x4a, 0x35, 0xd7, 0x59, 0x62, 0x97,
	0xb9, 0xa1, 0xda, 0x65, 0xfd, 0x44, 0xb0, 0xd7, 0xe7, 0xfe, 0x9b, 0xd8, 0x73, 0x05, 0x79, 0x99,
	0x65, 0xf0, 0x31, 0xd4, 0xdd, 0x91, 0x18, 0x52, 0x16, 0x88, 0x44, 0x43, 0x2d, 0xd4, 0xae, 0xf7,
	0xb4, 0x5f, 0x3f, 0x9e, 0xee, 0xab, 0x72, 0xa7, 0x9e, 0xc7, 0x08, 0xe7, 0xaf, 0x05, 0x0b, 0x22,
	0xdf, 0x59, 0x58, 0xf1, 0x73, 0xa8, 0xc9, 0xb3, 0xb5, 0x9d, 0x16, 0x6a, 0xdf, 0x3d, 0xd4, 0xbb,
	0x6b, 0xcd, 0x76, 0x65, 0x89, 0x5e, 0xfd, 0xfa, 0xcf, 0xc3, 0xca, 0xb7, 0xf9, 0xb8, 0x83, 0x1c,
	0xb5, 0xe7, 0xe4, 0xf8, 0xd3, 0x7c, 0xdc, 0x59, 0x9c, 0xf6, 0x79, 0x3e, 0xee, 0x3c, 0x2a, 0x98,
	0xaf, 0x16, 0xf8, 0x25, 0x5a, 0x4b, 0x87, 0x83, 0x92, 0xe4, 0x10, 0x1e, 0xd3, 0x88, 0x13, 0xeb,
	0x2b, 0x82, 0x66, 0x9f, 0xfb, 0x67, 0x8c, 0xb8, 0x82, 0xbc, 0x1a, 0x11, 0x2e, 0x02, 0x1a, 0x61,
	0x0d, 0x6e, 0x0f, 0x52, 0x85, 0x32, 0xd9, 0x9c, 0x93, 0x87, 0xd8, 0x80, 0x3b, 0xef, 0x95, 0x2b,
	0x6b, 0xa1, 0xee, 0x14, 0x31, 0x6e, 0xc3, 0x1e, 0x89, 0x06, 0x2c, 0x89, 0x05, 0xf1, 0x4e, 0x23,
	0xfe, 0x81, 0x30, 0x6d, 0x37, 0xb3, 0x94, 0x65, 0x7c, 0x1f, 0x6a, 0x17, 0x74, 0x14, 0x89, 0x44,
	0xab, 0xb6, 0x50, 0xbb, 0xea, 0xa8, 0xe8, 0xa4, 0x91, 0x36, 0x98, 0xd7, 0xb2, 0x1e, 0x80, 0xbe,
	0x86, 0x56, 0x80, 0x27, 0xd9, 0x50, 0xce, 0x68, 0x18, 0x06, 0x42, 0x9d, 0xba, 0x9d, 0xda, 0x04,
	0xc8, 0x29, 0x5f, 0x78, 0x19, 0x77, 0xd5, 0x59, 0x52, 0xd2, 0xfc, 0xd0, 0xe5, 0xc3, 0x15, 0xe8,
	0x25, 0xa5, 0xc4, 0x25, 0xaf, 0x73, 0xb9, 0x74, 0x4e, 0x75, 0xf8, 0x7d, 0x07, 0x76, 0xfb, 0xdc,
	0xc7, 0xe7, 0xd0, 0x58, 0x79, 0x2f, 0xd6, 0x86, 0x39, 0x97, 0x46, 0x62, 0x74, 0xfe, 0xef, 0xc9,
	0xeb, 0x60, 0x0f, 0xee, 0x95, 0x46, 0xf6, 0x78, 0xf3, 0xee, 0x55, 0x97, 0xf1, 0xe4, 0x26, 0xae,
	0xa2, 0xca, 0x39, 0x34, 0x56, 0x2e, 0x78, 0x4b, 0x17, 0xcb, 0x9e, 0x6d, 0x5d, 0x6c, 0xba, 0x2d,
	0xe3, 0xd6, 0xc7, 0xf4, 0x79, 0xf7, 0x8e, 0xae, 0xa7, 0x26, 0x9a, 0x4c, 0x4d, 0xf4, 0x77, 0x6a,
	0xa2, 0x2f, 0x33, 0xb3, 0x32, 0x99, 0x99, 0x95, 0xdf, 0x33, 0xb3, 0xf2, 0x56, 0xdf, 0xf4, 0xba,
	0x45, 0x12, 0x13, 0x7e, 0x51, 0xcb, 0x3e, 0xce, 0xa3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b,
	0xf3, 0xd7, 0x26, 0x4a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateQuestion(ctx context.Context, in *MsgCreateQuestion, opts ...grpc.CallOption) (*MsgCreateQuestionResponse, error)
	CommitAnswer(ctx context.Context, in *MsgCommitAnswer, opts ...grpc.CallOption) (*MsgCommitAnswerResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/scavenge.scavenge.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateQuestion(ctx context.Context, in *MsgCreateQuestion, opts ...grpc.CallOption) (*MsgCreateQuestionResponse, error) {
	out := new(MsgCreateQuestionResponse)
	err := c.cc.Invoke(ctx, "/scavenge.scavenge.Msg/CreateQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CommitAnswer(ctx context.Context, in *MsgCommitAnswer, opts ...grpc.CallOption) (*MsgCommitAnswerResponse, error) {
	out := new(MsgCommitAnswerResponse)
	err := c.cc.Invoke(ctx, "/scavenge.scavenge.Msg/CommitAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreateQuestion(context.Context, *MsgCreateQuestion) (*MsgCreateQuestionResponse, error)
	CommitAnswer(context.Context, *MsgCommitAnswer) (*MsgCommitAnswerResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) CreateQuestion(ctx context.Context, req *MsgCreateQuestion) (*MsgCreateQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestion not implemented")
}
func (*UnimplementedMsgServer) CommitAnswer(ctx context.Context, req *MsgCommitAnswer) (*MsgCommitAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitAnswer not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scavenge.scavenge.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateQuestion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scavenge.scavenge.Msg/CreateQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateQuestion(ctx, req.(*MsgCreateQuestion))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CommitAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCommitAnswer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CommitAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scavenge.scavenge.Msg/CommitAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CommitAnswer(ctx, req.(*MsgCommitAnswer))
	}
	return interceptor(ctx, in, info, handler)
}

var Msg_serviceDesc = _Msg_serviceDesc
var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scavenge.scavenge.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateQuestion",
			Handler:    _Msg_CreateQuestion_Handler,
		},
		{
			MethodName: "CommitAnswer",
			Handler:    _Msg_CommitAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scavenge/scavenge/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCreateQuestion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateQuestion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateQuestion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Bounty != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Bounty))
		i--
		dAtA[i] = 0x20
	}
	if len(m.EncryptedAnswer) > 0 {
		i -= len(m.EncryptedAnswer)
		copy(dAtA[i:], m.EncryptedAnswer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EncryptedAnswer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Question) > 0 {
		i -= len(m.Question)
		copy(dAtA[i:], m.Question)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Question)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateQuestionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateQuestionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateQuestionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCommitAnswer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCommitAnswer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCommitAnswer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HashAnswer) > 0 {
		i -= len(m.HashAnswer)
		copy(dAtA[i:], m.HashAnswer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HashAnswer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.QuestionId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.QuestionId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCommitAnswerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCommitAnswerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCommitAnswerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCreateQuestion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Question)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.EncryptedAnswer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Bounty != 0 {
		n += 1 + sovTx(uint64(m.Bounty))
	}
	return n
}

func (m *MsgCreateQuestionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCommitAnswer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.QuestionId != 0 {
		n += 1 + sovTx(uint64(m.QuestionId))
	}
	l = len(m.HashAnswer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCommitAnswerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateQuestion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateQuestion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateQuestion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Question", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Question = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedAnswer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedAnswer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bounty", wireType)
			}
			m.Bounty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bounty |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateQuestionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateQuestionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateQuestionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCommitAnswer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCommitAnswer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCommitAnswer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuestionId", wireType)
			}
			m.QuestionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashAnswer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashAnswer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCommitAnswerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCommitAnswerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCommitAnswerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
