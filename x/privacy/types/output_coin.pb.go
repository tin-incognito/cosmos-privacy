// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: privacy/output_coin.proto

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

type OutputCoin struct {
	Index               string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Creator             string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	IsConfidentialAsset bool   `protobuf:"varint,3,opt,name=is_confidential_asset,json=isConfidentialAsset,proto3" json:"is_confidential_asset,omitempty"`
	Value               []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *OutputCoin) Reset()         { *m = OutputCoin{} }
func (m *OutputCoin) String() string { return proto.CompactTextString(m) }
func (*OutputCoin) ProtoMessage()    {}
func (*OutputCoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_4793c75a9286f400, []int{0}
}
func (m *OutputCoin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutputCoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutputCoin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutputCoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputCoin.Merge(m, src)
}
func (m *OutputCoin) XXX_Size() int {
	return m.Size()
}
func (m *OutputCoin) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputCoin.DiscardUnknown(m)
}

var xxx_messageInfo_OutputCoin proto.InternalMessageInfo

func (m *OutputCoin) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *OutputCoin) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *OutputCoin) GetIsConfidentialAsset() bool {
	if m != nil {
		return m.IsConfidentialAsset
	}
	return false
}

func (m *OutputCoin) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*OutputCoin)(nil), "privacy.privacy.OutputCoin")
}

func init() { proto.RegisterFile("privacy/output_coin.proto", fileDescriptor_4793c75a9286f400) }

var fileDescriptor_4793c75a9286f400 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2c,
	0x4b, 0x4c, 0xae, 0xd4, 0xcf, 0x2f, 0x2d, 0x29, 0x28, 0x2d, 0x89, 0x4f, 0xce, 0xcf, 0xcc, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0x4a, 0xe9, 0x41, 0x69, 0xa5, 0x36, 0x46, 0x2e,
	0x2e, 0x7f, 0xb0, 0x32, 0xe7, 0xfc, 0xcc, 0x3c, 0x21, 0x11, 0x2e, 0xd6, 0xcc, 0xbc, 0x94, 0xd4,
	0x0a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0x3d, 0xb9, 0x28,
	0x35, 0xb1, 0x24, 0xbf, 0x48, 0x82, 0x09, 0x2c, 0x0e, 0xe3, 0x0a, 0x19, 0x71, 0x89, 0x66, 0x16,
	0xc7, 0x27, 0xe7, 0xe7, 0xa5, 0x65, 0xa6, 0xa4, 0xe6, 0x95, 0x64, 0x26, 0xe6, 0xc4, 0x27, 0x16,
	0x17, 0xa7, 0x96, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x04, 0x09, 0x67, 0x16, 0x3b, 0x23, 0xc9,
	0x39, 0x82, 0xa4, 0x40, 0x76, 0x94, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0xb0, 0x28, 0x30, 0x6a, 0xf0,
	0x04, 0x41, 0x38, 0x4e, 0x86, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25,
	0x0e, 0xf3, 0x4e, 0x85, 0x3e, 0x8c, 0x55, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0x93,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x57, 0xe9, 0x7b, 0xc5, 0xf0, 0x00, 0x00, 0x00,
}

func (m *OutputCoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutputCoin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutputCoin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintOutputCoin(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if m.IsConfidentialAsset {
		i--
		if m.IsConfidentialAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintOutputCoin(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintOutputCoin(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOutputCoin(dAtA []byte, offset int, v uint64) int {
	offset -= sovOutputCoin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutputCoin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovOutputCoin(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovOutputCoin(uint64(l))
	}
	if m.IsConfidentialAsset {
		n += 2
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovOutputCoin(uint64(l))
	}
	return n
}

func sovOutputCoin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOutputCoin(x uint64) (n int) {
	return sovOutputCoin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutputCoin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOutputCoin
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
			return fmt.Errorf("proto: OutputCoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutputCoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutputCoin
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
				return ErrInvalidLengthOutputCoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutputCoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutputCoin
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
				return ErrInvalidLengthOutputCoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOutputCoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsConfidentialAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutputCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsConfidentialAsset = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOutputCoin
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
				return ErrInvalidLengthOutputCoin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOutputCoin
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
			skippy, err := skipOutputCoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOutputCoin
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
func skipOutputCoin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOutputCoin
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
					return 0, ErrIntOverflowOutputCoin
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
					return 0, ErrIntOverflowOutputCoin
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
				return 0, ErrInvalidLengthOutputCoin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOutputCoin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOutputCoin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOutputCoin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOutputCoin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOutputCoin = fmt.Errorf("proto: unexpected end of group")
)
