// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: privacy/output_coin_length.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type OutputCoinLength struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *OutputCoinLength) Reset()         { *m = OutputCoinLength{} }
func (m *OutputCoinLength) String() string { return proto.CompactTextString(m) }
func (*OutputCoinLength) ProtoMessage()    {}
func (*OutputCoinLength) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f9dba7c0f2a8ea1, []int{0}
}
func (m *OutputCoinLength) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutputCoinLength) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutputCoinLength.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutputCoinLength) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputCoinLength.Merge(m, src)
}
func (m *OutputCoinLength) XXX_Size() int {
	return m.Size()
}
func (m *OutputCoinLength) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputCoinLength.DiscardUnknown(m)
}

var xxx_messageInfo_OutputCoinLength proto.InternalMessageInfo

func (m *OutputCoinLength) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*OutputCoinLength)(nil), "privacy.privacy.OutputCoinLength")
}

func init() { proto.RegisterFile("privacy/output_coin_length.proto", fileDescriptor_2f9dba7c0f2a8ea1) }

var fileDescriptor_2f9dba7c0f2a8ea1 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2c,
	0x4b, 0x4c, 0xae, 0xd4, 0xcf, 0x2f, 0x2d, 0x29, 0x28, 0x2d, 0x89, 0x4f, 0xce, 0xcf, 0xcc, 0x8b,
	0xcf, 0x49, 0xcd, 0x4b, 0x2f, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0xaa,
	0xd0, 0x83, 0xd2, 0x4a, 0x1a, 0x5c, 0x02, 0xfe, 0x60, 0xc5, 0xce, 0xf9, 0x99, 0x79, 0x3e, 0x60,
	0xa5, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x41, 0x10, 0x8e, 0x93, 0xe1, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x89,
	0xc3, 0xac, 0xad, 0xd0, 0x87, 0xb1, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x96, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x30, 0xdb, 0x5d, 0xb1, 0x98, 0x00, 0x00, 0x00,
}

func (m *OutputCoinLength) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutputCoinLength) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutputCoinLength) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintOutputCoinLength(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOutputCoinLength(dAtA []byte, offset int, v uint64) int {
	offset -= sovOutputCoinLength(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutputCoinLength) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovOutputCoinLength(uint64(l))
	}
	return n
}

func sovOutputCoinLength(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOutputCoinLength(x uint64) (n int) {
	return sovOutputCoinLength(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutputCoinLength) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOutputCoinLength
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
			return fmt.Errorf("proto: OutputCoinLength: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutputCoinLength: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutputCoinLength
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthOutputCoinLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOutputCoinLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOutputCoinLength(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOutputCoinLength
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
func skipOutputCoinLength(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOutputCoinLength
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
					return 0, ErrIntOverflowOutputCoinLength
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
					return 0, ErrIntOverflowOutputCoinLength
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
				return 0, ErrInvalidLengthOutputCoinLength
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOutputCoinLength
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOutputCoinLength
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOutputCoinLength        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOutputCoinLength          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOutputCoinLength = fmt.Errorf("proto: unexpected end of group")
)