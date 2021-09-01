// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator/last_validator_power.proto

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

type LastValidatorPower struct {
	Index string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Power int64  `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *LastValidatorPower) Reset()         { *m = LastValidatorPower{} }
func (m *LastValidatorPower) String() string { return proto.CompactTextString(m) }
func (*LastValidatorPower) ProtoMessage()    {}
func (*LastValidatorPower) Descriptor() ([]byte, []int) {
	return fileDescriptor_830cb29c5288b722, []int{0}
}
func (m *LastValidatorPower) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastValidatorPower) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastValidatorPower.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastValidatorPower) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastValidatorPower.Merge(m, src)
}
func (m *LastValidatorPower) XXX_Size() int {
	return m.Size()
}
func (m *LastValidatorPower) XXX_DiscardUnknown() {
	xxx_messageInfo_LastValidatorPower.DiscardUnknown(m)
}

var xxx_messageInfo_LastValidatorPower proto.InternalMessageInfo

func (m *LastValidatorPower) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *LastValidatorPower) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func init() {
	proto.RegisterType((*LastValidatorPower)(nil), "liubaninc.m0.validator.LastValidatorPower")
}

func init() {
	proto.RegisterFile("validator/last_validator_power.proto", fileDescriptor_830cb29c5288b722)
}

var fileDescriptor_830cb29c5288b722 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0xcf, 0x49, 0x2c, 0x2e, 0x89, 0x87, 0x73, 0xe3, 0x0b, 0xf2,
	0xcb, 0x53, 0x8b, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0x72, 0x32, 0x4b, 0x93, 0x12,
	0xf3, 0x32, 0xf3, 0x92, 0xf5, 0x72, 0x0d, 0xf4, 0xe0, 0x6a, 0x94, 0x1c, 0xb8, 0x84, 0x7c, 0x12,
	0x8b, 0x4b, 0xc2, 0x60, 0x02, 0x01, 0x20, 0x3d, 0x42, 0x22, 0x5c, 0xac, 0x99, 0x79, 0x29, 0xa9,
	0x15, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x0e, 0x48, 0x14, 0x6c, 0xa4, 0x04, 0xb3,
	0x02, 0xa3, 0x06, 0x73, 0x10, 0x84, 0xe3, 0xe4, 0x7c, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72,
	0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7,
	0x72, 0x0c, 0x51, 0x9a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x70,
	0xeb, 0xf5, 0x73, 0x0d, 0xf4, 0x2b, 0xf4, 0x11, 0x6e, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62,
	0x03, 0xbb, 0xd2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x19, 0xfa, 0xe6, 0x86, 0xcd, 0x00, 0x00,
	0x00,
}

func (m *LastValidatorPower) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastValidatorPower) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastValidatorPower) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintLastValidatorPower(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintLastValidatorPower(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintLastValidatorPower(dAtA []byte, offset int, v uint64) int {
	offset -= sovLastValidatorPower(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LastValidatorPower) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovLastValidatorPower(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovLastValidatorPower(uint64(m.Power))
	}
	return n
}

func sovLastValidatorPower(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLastValidatorPower(x uint64) (n int) {
	return sovLastValidatorPower(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LastValidatorPower) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLastValidatorPower
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
			return fmt.Errorf("proto: LastValidatorPower: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastValidatorPower: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLastValidatorPower
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
				return ErrInvalidLengthLastValidatorPower
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLastValidatorPower
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLastValidatorPower
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLastValidatorPower(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLastValidatorPower
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
func skipLastValidatorPower(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLastValidatorPower
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
					return 0, ErrIntOverflowLastValidatorPower
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
					return 0, ErrIntOverflowLastValidatorPower
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
				return 0, ErrInvalidLengthLastValidatorPower
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLastValidatorPower
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLastValidatorPower
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLastValidatorPower        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLastValidatorPower          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLastValidatorPower = fmt.Errorf("proto: unexpected end of group")
)
