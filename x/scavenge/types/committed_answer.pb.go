// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: scavenge/scavenge/committed_answer.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type CommittedAnswer struct {
	QuestionId uint64 `protobuf:"varint,1,opt,name=questionId,proto3" json:"questionId,omitempty"`
	HashAnswer string `protobuf:"bytes,2,opt,name=hashAnswer,proto3" json:"hashAnswer,omitempty"`
	Creator    string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *CommittedAnswer) Reset()         { *m = CommittedAnswer{} }
func (m *CommittedAnswer) String() string { return proto.CompactTextString(m) }
func (*CommittedAnswer) ProtoMessage()    {}
func (*CommittedAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1d0eea702bc0a10, []int{0}
}
func (m *CommittedAnswer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommittedAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommittedAnswer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommittedAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommittedAnswer.Merge(m, src)
}
func (m *CommittedAnswer) XXX_Size() int {
	return m.Size()
}
func (m *CommittedAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_CommittedAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_CommittedAnswer proto.InternalMessageInfo

func (m *CommittedAnswer) GetQuestionId() uint64 {
	if m != nil {
		return m.QuestionId
	}
	return 0
}

func (m *CommittedAnswer) GetHashAnswer() string {
	if m != nil {
		return m.HashAnswer
	}
	return ""
}

func (m *CommittedAnswer) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*CommittedAnswer)(nil), "scavenge.scavenge.CommittedAnswer")
}

func init() {
	proto.RegisterFile("scavenge/scavenge/committed_answer.proto", fileDescriptor_f1d0eea702bc0a10)
}

var fileDescriptor_f1d0eea702bc0a10 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x4e, 0x4e, 0x2c,
	0x4b, 0xcd, 0x4b, 0x4f, 0xd5, 0x87, 0x33, 0x92, 0xf3, 0x73, 0x73, 0x33, 0x4b, 0x4a, 0x52, 0x53,
	0xe2, 0x13, 0xf3, 0x8a, 0xcb, 0x53, 0x8b, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x61,
	0x0a, 0xf4, 0x60, 0x0c, 0xa5, 0x6c, 0x2e, 0x7e, 0x67, 0x98, 0x62, 0x47, 0xb0, 0x5a, 0x21, 0x39,
	0x2e, 0xae, 0xc2, 0xd2, 0xd4, 0xe2, 0x92, 0xcc, 0xfc, 0x3c, 0xcf, 0x14, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0x96, 0x20, 0x24, 0x11, 0x90, 0x7c, 0x46, 0x62, 0x71, 0x06, 0x44, 0xb5, 0x04, 0x93, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x92, 0x88, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x7e,
	0x91, 0x04, 0x33, 0x58, 0x12, 0xc6, 0x75, 0x32, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0x49, 0xb8, 0xd3, 0x2b, 0x10, 0xbe, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62,
	0x03, 0xbb, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x3d, 0x3e, 0x57, 0xe7, 0x00, 0x00,
	0x00,
}

func (m *CommittedAnswer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommittedAnswer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommittedAnswer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintCommittedAnswer(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.HashAnswer) > 0 {
		i -= len(m.HashAnswer)
		copy(dAtA[i:], m.HashAnswer)
		i = encodeVarintCommittedAnswer(dAtA, i, uint64(len(m.HashAnswer)))
		i--
		dAtA[i] = 0x12
	}
	if m.QuestionId != 0 {
		i = encodeVarintCommittedAnswer(dAtA, i, uint64(m.QuestionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommittedAnswer(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommittedAnswer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CommittedAnswer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QuestionId != 0 {
		n += 1 + sovCommittedAnswer(uint64(m.QuestionId))
	}
	l = len(m.HashAnswer)
	if l > 0 {
		n += 1 + l + sovCommittedAnswer(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovCommittedAnswer(uint64(l))
	}
	return n
}

func sovCommittedAnswer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommittedAnswer(x uint64) (n int) {
	return sovCommittedAnswer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CommittedAnswer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommittedAnswer
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
			return fmt.Errorf("proto: CommittedAnswer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommittedAnswer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuestionId", wireType)
			}
			m.QuestionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittedAnswer
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashAnswer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittedAnswer
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
				return ErrInvalidLengthCommittedAnswer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittedAnswer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HashAnswer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommittedAnswer
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
				return ErrInvalidLengthCommittedAnswer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommittedAnswer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommittedAnswer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommittedAnswer
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
func skipCommittedAnswer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommittedAnswer
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
					return 0, ErrIntOverflowCommittedAnswer
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
					return 0, ErrIntOverflowCommittedAnswer
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
				return 0, ErrInvalidLengthCommittedAnswer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommittedAnswer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommittedAnswer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommittedAnswer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommittedAnswer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommittedAnswer = fmt.Errorf("proto: unexpected end of group")
)
