// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: privacy/serial_number.proto

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

type SerialNumber struct {
	Index        string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Creator      string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	TokenId      []byte `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	SerialNumber []byte `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
}

func (m *SerialNumber) Reset()         { *m = SerialNumber{} }
func (m *SerialNumber) String() string { return proto.CompactTextString(m) }
func (*SerialNumber) ProtoMessage()    {}
func (*SerialNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9be2931f060f752, []int{0}
}
func (m *SerialNumber) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SerialNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SerialNumber.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SerialNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerialNumber.Merge(m, src)
}
func (m *SerialNumber) XXX_Size() int {
	return m.Size()
}
func (m *SerialNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_SerialNumber.DiscardUnknown(m)
}

var xxx_messageInfo_SerialNumber proto.InternalMessageInfo

func (m *SerialNumber) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SerialNumber) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SerialNumber) GetTokenId() []byte {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *SerialNumber) GetSerialNumber() []byte {
	if m != nil {
		return m.SerialNumber
	}
	return nil
}

func init() {
	proto.RegisterType((*SerialNumber)(nil), "privacy.privacy.SerialNumber")
}

func init() { proto.RegisterFile("privacy/serial_number.proto", fileDescriptor_e9be2931f060f752) }

var fileDescriptor_e9be2931f060f752 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2c,
	0x4b, 0x4c, 0xae, 0xd4, 0x2f, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0x89, 0xcf, 0x2b, 0xcd, 0x4d, 0x4a,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0x4a, 0xea, 0x41, 0x69, 0xa5, 0x3a,
	0x2e, 0x9e, 0x60, 0xb0, 0x3a, 0x3f, 0xb0, 0x32, 0x21, 0x11, 0x2e, 0xd6, 0xcc, 0xbc, 0x94, 0xd4,
	0x0a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b, 0x3d, 0xb9, 0x28,
	0x35, 0xb1, 0x24, 0xbf, 0x48, 0x82, 0x09, 0x2c, 0x0e, 0xe3, 0x0a, 0x49, 0x72, 0x71, 0x94, 0xe4,
	0x67, 0xa7, 0xe6, 0xc5, 0x67, 0xa6, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0xf0, 0x04, 0xb1, 0x83, 0xf9,
	0x9e, 0x29, 0x42, 0xca, 0x5c, 0xbc, 0x28, 0x4e, 0x90, 0x60, 0x01, 0xcb, 0xf3, 0x14, 0x23, 0xd9,
	0xe7, 0x64, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xe2, 0x30, 0x7f,
	0x54, 0xe8, 0xc3, 0x58, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xaf, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x68, 0xd9, 0x0f, 0xbc, 0xe9, 0x00, 0x00, 0x00,
}

func (m *SerialNumber) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SerialNumber) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SerialNumber) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SerialNumber) > 0 {
		i -= len(m.SerialNumber)
		copy(dAtA[i:], m.SerialNumber)
		i = encodeVarintSerialNumber(dAtA, i, uint64(len(m.SerialNumber)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintSerialNumber(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSerialNumber(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintSerialNumber(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSerialNumber(dAtA []byte, offset int, v uint64) int {
	offset -= sovSerialNumber(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SerialNumber) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovSerialNumber(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSerialNumber(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovSerialNumber(uint64(l))
	}
	l = len(m.SerialNumber)
	if l > 0 {
		n += 1 + l + sovSerialNumber(uint64(l))
	}
	return n
}

func sovSerialNumber(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSerialNumber(x uint64) (n int) {
	return sovSerialNumber(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SerialNumber) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSerialNumber
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
			return fmt.Errorf("proto: SerialNumber: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SerialNumber: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerialNumber
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
				return ErrInvalidLengthSerialNumber
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSerialNumber
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
					return ErrIntOverflowSerialNumber
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
				return ErrInvalidLengthSerialNumber
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSerialNumber
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerialNumber
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
				return ErrInvalidLengthSerialNumber
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSerialNumber
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = append(m.TokenId[:0], dAtA[iNdEx:postIndex]...)
			if m.TokenId == nil {
				m.TokenId = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialNumber", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSerialNumber
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
				return ErrInvalidLengthSerialNumber
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSerialNumber
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerialNumber = append(m.SerialNumber[:0], dAtA[iNdEx:postIndex]...)
			if m.SerialNumber == nil {
				m.SerialNumber = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSerialNumber(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSerialNumber
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
func skipSerialNumber(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSerialNumber
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
					return 0, ErrIntOverflowSerialNumber
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
					return 0, ErrIntOverflowSerialNumber
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
				return 0, ErrInvalidLengthSerialNumber
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSerialNumber
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSerialNumber
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSerialNumber        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSerialNumber          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSerialNumber = fmt.Errorf("proto: unexpected end of group")
)