// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: privacy/tx_privacy_data.proto

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

type TxPrivacyData struct {
	Index     string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Creator   string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	LockTime  uint64 `protobuf:"varint,3,opt,name=lock_time,json=lockTime,proto3" json:"lock_time,omitempty"`
	Fee       uint64 `protobuf:"varint,4,opt,name=fee,proto3" json:"fee,omitempty"`
	Info      []byte `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	SigPubKey []byte `protobuf:"bytes,6,opt,name=sig_pub_key,json=sigPubKey,proto3" json:"sig_pub_key,omitempty"`
	Sig       []byte `protobuf:"bytes,7,opt,name=sig,proto3" json:"sig,omitempty"`
	Proof     []byte `protobuf:"bytes,8,opt,name=proof,proto3" json:"proof,omitempty"`
	Messages  []byte `protobuf:"bytes,9,opt,name=messages,proto3" json:"messages,omitempty"`
}

func (m *TxPrivacyData) Reset()         { *m = TxPrivacyData{} }
func (m *TxPrivacyData) String() string { return proto.CompactTextString(m) }
func (*TxPrivacyData) ProtoMessage()    {}
func (*TxPrivacyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_eca4abe5e130beda, []int{0}
}
func (m *TxPrivacyData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxPrivacyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxPrivacyData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxPrivacyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxPrivacyData.Merge(m, src)
}
func (m *TxPrivacyData) XXX_Size() int {
	return m.Size()
}
func (m *TxPrivacyData) XXX_DiscardUnknown() {
	xxx_messageInfo_TxPrivacyData.DiscardUnknown(m)
}

var xxx_messageInfo_TxPrivacyData proto.InternalMessageInfo

func (m *TxPrivacyData) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *TxPrivacyData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TxPrivacyData) GetLockTime() uint64 {
	if m != nil {
		return m.LockTime
	}
	return 0
}

func (m *TxPrivacyData) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *TxPrivacyData) GetInfo() []byte {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *TxPrivacyData) GetSigPubKey() []byte {
	if m != nil {
		return m.SigPubKey
	}
	return nil
}

func (m *TxPrivacyData) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *TxPrivacyData) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *TxPrivacyData) GetMessages() []byte {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*TxPrivacyData)(nil), "privacy.privacy.TxPrivacyData")
}

func init() { proto.RegisterFile("privacy/tx_privacy_data.proto", fileDescriptor_eca4abe5e130beda) }

var fileDescriptor_eca4abe5e130beda = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x1b, 0xd7, 0x6d, 0x6d, 0x54, 0x94, 0x20, 0xf8, 0xa1, 0x18, 0x8a, 0xa7, 0x9e, 0x36,
	0xc4, 0x7f, 0x20, 0xde, 0xbc, 0x8c, 0xb2, 0x93, 0x97, 0x92, 0x76, 0x5f, 0x4b, 0x98, 0x5d, 0x4a,
	0x93, 0x49, 0xfb, 0x2f, 0xfc, 0x59, 0x1e, 0x77, 0xf4, 0x28, 0x2d, 0xf8, 0x3b, 0xa4, 0x69, 0xb7,
	0x53, 0xde, 0xe7, 0x79, 0x21, 0x7c, 0xbc, 0xf4, 0xa1, 0xac, 0xe4, 0xa7, 0x48, 0x9b, 0xa5, 0xa9,
	0xe3, 0x31, 0xc6, 0x1b, 0x61, 0xc4, 0xa2, 0xac, 0x94, 0x51, 0xec, 0x6a, 0x74, 0x8b, 0xf1, 0x7d,
	0xfc, 0x23, 0xf4, 0x72, 0x5d, 0xaf, 0x06, 0x7a, 0x15, 0x46, 0xb0, 0x1b, 0x3a, 0x95, 0xbb, 0x0d,
	0xd6, 0x40, 0x02, 0x12, 0xfa, 0xd1, 0x00, 0x0c, 0xe8, 0x3c, 0xad, 0x50, 0x18, 0x55, 0xc1, 0x99,
	0xf5, 0x47, 0x64, 0xf7, 0xd4, 0xff, 0x50, 0xe9, 0x36, 0x36, 0xb2, 0x40, 0x98, 0x04, 0x24, 0x74,
	0x23, 0xaf, 0x17, 0x6b, 0x59, 0x20, 0xbb, 0xa6, 0x93, 0x0c, 0x11, 0x5c, 0xab, 0xfb, 0xc8, 0x18,
	0x75, 0xe5, 0x2e, 0x53, 0x30, 0x0d, 0x48, 0x78, 0x11, 0xd9, 0xcc, 0x38, 0x3d, 0xd7, 0x32, 0x8f,
	0xcb, 0x7d, 0x12, 0x6f, 0xb1, 0x81, 0x99, 0xad, 0x7c, 0x2d, 0xf3, 0xd5, 0x3e, 0x79, 0xc3, 0xa6,
	0xff, 0x45, 0xcb, 0x1c, 0xe6, 0xd6, 0xf7, 0xb1, 0x3f, 0xb2, 0xac, 0x94, 0xca, 0xc0, 0xb3, 0x6e,
	0x00, 0x76, 0x47, 0xbd, 0x02, 0xb5, 0x16, 0x39, 0x6a, 0xf0, 0x6d, 0x71, 0xe2, 0x97, 0xa7, 0xef,
	0x96, 0x93, 0x43, 0xcb, 0xc9, 0x6f, 0xcb, 0xc9, 0x57, 0xc7, 0x9d, 0x43, 0xc7, 0x9d, 0x9f, 0x8e,
	0x3b, 0xef, 0xb7, 0xc7, 0xc9, 0xea, 0xe5, 0x69, 0xbc, 0xa6, 0x44, 0x9d, 0xcc, 0xec, 0x66, 0xcf,
	0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x89, 0x37, 0x51, 0x14, 0x54, 0x01, 0x00, 0x00,
}

func (m *TxPrivacyData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxPrivacyData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxPrivacyData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		i -= len(m.Messages)
		copy(dAtA[i:], m.Messages)
		i = encodeVarintTxPrivacyData(dAtA, i, uint64(len(m.Messages)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintTxPrivacyData(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintTxPrivacyData(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.SigPubKey) > 0 {
		i -= len(m.SigPubKey)
		copy(dAtA[i:], m.SigPubKey)
		i = encodeVarintTxPrivacyData(dAtA, i, uint64(len(m.SigPubKey)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintTxPrivacyData(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Fee != 0 {
		i = encodeVarintTxPrivacyData(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x20
	}
	if m.LockTime != 0 {
		i = encodeVarintTxPrivacyData(dAtA, i, uint64(m.LockTime))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTxPrivacyData(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTxPrivacyData(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTxPrivacyData(dAtA []byte, offset int, v uint64) int {
	offset -= sovTxPrivacyData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TxPrivacyData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTxPrivacyData(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTxPrivacyData(uint64(l))
	}
	if m.LockTime != 0 {
		n += 1 + sovTxPrivacyData(uint64(m.LockTime))
	}
	if m.Fee != 0 {
		n += 1 + sovTxPrivacyData(uint64(m.Fee))
	}
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovTxPrivacyData(uint64(l))
	}
	l = len(m.SigPubKey)
	if l > 0 {
		n += 1 + l + sovTxPrivacyData(uint64(l))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovTxPrivacyData(uint64(l))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovTxPrivacyData(uint64(l))
	}
	l = len(m.Messages)
	if l > 0 {
		n += 1 + l + sovTxPrivacyData(uint64(l))
	}
	return n
}

func sovTxPrivacyData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTxPrivacyData(x uint64) (n int) {
	return sovTxPrivacyData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TxPrivacyData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTxPrivacyData
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
			return fmt.Errorf("proto: TxPrivacyData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxPrivacyData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxPrivacyData
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
				return ErrInvalidLengthTxPrivacyData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxPrivacyData
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
					return ErrIntOverflowTxPrivacyData
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
				return ErrInvalidLengthTxPrivacyData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTxPrivacyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockTime", wireType)
			}
			m.LockTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxPrivacyData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxPrivacyData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fee |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxPrivacyData
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
				return ErrInvalidLengthTxPrivacyData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxPrivacyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = append(m.Info[:0], dAtA[iNdEx:postIndex]...)
			if m.Info == nil {
				m.Info = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigPubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxPrivacyData
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
				return ErrInvalidLengthTxPrivacyData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxPrivacyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SigPubKey = append(m.SigPubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.SigPubKey == nil {
				m.SigPubKey = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxPrivacyData
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
				return ErrInvalidLengthTxPrivacyData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxPrivacyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxPrivacyData
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
				return ErrInvalidLengthTxPrivacyData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxPrivacyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = append(m.Proof[:0], dAtA[iNdEx:postIndex]...)
			if m.Proof == nil {
				m.Proof = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTxPrivacyData
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
				return ErrInvalidLengthTxPrivacyData
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTxPrivacyData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages[:0], dAtA[iNdEx:postIndex]...)
			if m.Messages == nil {
				m.Messages = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTxPrivacyData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTxPrivacyData
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
func skipTxPrivacyData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTxPrivacyData
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
					return 0, ErrIntOverflowTxPrivacyData
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
					return 0, ErrIntOverflowTxPrivacyData
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
				return 0, ErrInvalidLengthTxPrivacyData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTxPrivacyData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTxPrivacyData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTxPrivacyData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTxPrivacyData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTxPrivacyData = fmt.Errorf("proto: unexpected end of group")
)
