// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: m0/storage/encrypt_storage.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type EncryptStorage struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index     string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Md5       string `protobuf:"bytes,3,opt,name=md5,proto3" json:"md5,omitempty"`
	Encrypted string `protobuf:"bytes,4,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	Envelope  string `protobuf:"bytes,5,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (m *EncryptStorage) Reset()         { *m = EncryptStorage{} }
func (m *EncryptStorage) String() string { return proto.CompactTextString(m) }
func (*EncryptStorage) ProtoMessage()    {}
func (*EncryptStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_51624de81e293fc6, []int{0}
}
func (m *EncryptStorage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EncryptStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EncryptStorage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EncryptStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptStorage.Merge(m, src)
}
func (m *EncryptStorage) XXX_Size() int {
	return m.Size()
}
func (m *EncryptStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptStorage.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptStorage proto.InternalMessageInfo

func (m *EncryptStorage) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *EncryptStorage) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *EncryptStorage) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *EncryptStorage) GetEncrypted() string {
	if m != nil {
		return m.Encrypted
	}
	return ""
}

func (m *EncryptStorage) GetEnvelope() string {
	if m != nil {
		return m.Envelope
	}
	return ""
}

func init() {
	proto.RegisterType((*EncryptStorage)(nil), "liubaninc.m0.storage.EncryptStorage")
}

func init() { proto.RegisterFile("m0/storage/encrypt_storage.proto", fileDescriptor_51624de81e293fc6) }

var fileDescriptor_51624de81e293fc6 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0x35, 0xd0, 0x2f,
	0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x4f, 0xcd, 0x4b, 0x2e, 0xaa, 0x2c, 0x28, 0x89, 0x87,
	0xf2, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0x72, 0x32, 0x4b, 0x93, 0x12, 0xf3, 0x32,
	0xf3, 0x92, 0xf5, 0x72, 0x0d, 0xf4, 0xa0, 0x72, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x05,
	0xfa, 0x20, 0x16, 0x44, 0xad, 0x52, 0x0f, 0x23, 0x17, 0x9f, 0x2b, 0xc4, 0x94, 0x60, 0x88, 0x42,
	0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x18, 0x57, 0x48, 0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x09, 0x2c,
	0x0e, 0xe1, 0x08, 0x09, 0x70, 0x31, 0xe7, 0xa6, 0x98, 0x4a, 0x30, 0x83, 0xc5, 0x40, 0x4c, 0x21,
	0x19, 0x2e, 0x4e, 0xa8, 0xcb, 0x52, 0x53, 0x24, 0x58, 0xc0, 0xe2, 0x08, 0x01, 0x21, 0x29, 0x2e,
	0x8e, 0xd4, 0xbc, 0xb2, 0xd4, 0x9c, 0xfc, 0x82, 0x54, 0x09, 0x56, 0xb0, 0x24, 0x9c, 0xef, 0xe4,
	0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c,
	0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xea, 0xe9, 0x99, 0x25, 0x19,
	0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x70, 0xff, 0xe9, 0xe7, 0x1a, 0xe8, 0x57, 0xc0, 0x43,
	0xa3, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x31, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x8a, 0x20, 0xad, 0x88, 0x28, 0x01, 0x00, 0x00,
}

func (m *EncryptStorage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EncryptStorage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EncryptStorage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Envelope) > 0 {
		i -= len(m.Envelope)
		copy(dAtA[i:], m.Envelope)
		i = encodeVarintEncryptStorage(dAtA, i, uint64(len(m.Envelope)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Encrypted) > 0 {
		i -= len(m.Encrypted)
		copy(dAtA[i:], m.Encrypted)
		i = encodeVarintEncryptStorage(dAtA, i, uint64(len(m.Encrypted)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Md5) > 0 {
		i -= len(m.Md5)
		copy(dAtA[i:], m.Md5)
		i = encodeVarintEncryptStorage(dAtA, i, uint64(len(m.Md5)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintEncryptStorage(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEncryptStorage(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEncryptStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovEncryptStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EncryptStorage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEncryptStorage(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovEncryptStorage(uint64(l))
	}
	l = len(m.Md5)
	if l > 0 {
		n += 1 + l + sovEncryptStorage(uint64(l))
	}
	l = len(m.Encrypted)
	if l > 0 {
		n += 1 + l + sovEncryptStorage(uint64(l))
	}
	l = len(m.Envelope)
	if l > 0 {
		n += 1 + l + sovEncryptStorage(uint64(l))
	}
	return n
}

func sovEncryptStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEncryptStorage(x uint64) (n int) {
	return sovEncryptStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EncryptStorage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEncryptStorage
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
			return fmt.Errorf("proto: EncryptStorage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EncryptStorage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncryptStorage
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
				return ErrInvalidLengthEncryptStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEncryptStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncryptStorage
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
				return ErrInvalidLengthEncryptStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEncryptStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Md5", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncryptStorage
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
				return ErrInvalidLengthEncryptStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEncryptStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Md5 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encrypted", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncryptStorage
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
				return ErrInvalidLengthEncryptStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEncryptStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Encrypted = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Envelope", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEncryptStorage
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
				return ErrInvalidLengthEncryptStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEncryptStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Envelope = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEncryptStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEncryptStorage
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
func skipEncryptStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEncryptStorage
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
					return 0, ErrIntOverflowEncryptStorage
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
					return 0, ErrIntOverflowEncryptStorage
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
				return 0, ErrInvalidLengthEncryptStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEncryptStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEncryptStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEncryptStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEncryptStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEncryptStorage = fmt.Errorf("proto: unexpected end of group")
)
