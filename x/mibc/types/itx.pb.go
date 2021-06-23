// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mibc/itx.proto

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

type Itx struct {
	Creator         string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id              uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	SourceHash      string `protobuf:"bytes,3,opt,name=sourceHash,proto3" json:"sourceHash,omitempty"`
	DestinationHash string `protobuf:"bytes,4,opt,name=destinationHash,proto3" json:"destinationHash,omitempty"`
	Source          bool   `protobuf:"varint,5,opt,name=source,proto3" json:"source,omitempty"`
	Log             string `protobuf:"bytes,6,opt,name=log,proto3" json:"log,omitempty"`
}

func (m *Itx) Reset()         { *m = Itx{} }
func (m *Itx) String() string { return proto.CompactTextString(m) }
func (*Itx) ProtoMessage()    {}
func (*Itx) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9cfe4c56c792caf, []int{0}
}
func (m *Itx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Itx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Itx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Itx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Itx.Merge(m, src)
}
func (m *Itx) XXX_Size() int {
	return m.Size()
}
func (m *Itx) XXX_DiscardUnknown() {
	xxx_messageInfo_Itx.DiscardUnknown(m)
}

var xxx_messageInfo_Itx proto.InternalMessageInfo

func (m *Itx) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Itx) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Itx) GetSourceHash() string {
	if m != nil {
		return m.SourceHash
	}
	return ""
}

func (m *Itx) GetDestinationHash() string {
	if m != nil {
		return m.DestinationHash
	}
	return ""
}

func (m *Itx) GetSource() bool {
	if m != nil {
		return m.Source
	}
	return false
}

func (m *Itx) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func init() {
	proto.RegisterType((*Itx)(nil), "liubaninc.m0.mibc.Itx")
}

func init() { proto.RegisterFile("mibc/itx.proto", fileDescriptor_c9cfe4c56c792caf) }

var fileDescriptor_c9cfe4c56c792caf = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcd, 0x4c, 0x4a,
	0xd6, 0xcf, 0x2c, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcc, 0xc9, 0x2c, 0x4d,
	0x4a, 0xcc, 0xcb, 0xcc, 0x4b, 0xd6, 0xcb, 0x35, 0xd0, 0x03, 0x49, 0x2a, 0x2d, 0x66, 0xe4, 0x62,
	0xf6, 0x2c, 0xa9, 0x10, 0x92, 0xe0, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x98, 0x14,
	0x18, 0x35, 0x58, 0x82, 0x98, 0x32, 0x53, 0x84, 0xe4, 0xb8, 0xb8, 0x8a, 0xf3, 0x4b, 0x8b, 0x92,
	0x53, 0x3d, 0x12, 0x8b, 0x33, 0x24, 0x98, 0xc1, 0x8a, 0x91, 0x44, 0x84, 0x34, 0xb8, 0xf8, 0x53,
	0x52, 0x8b, 0x4b, 0x32, 0xf3, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0xc0, 0x8a, 0x58, 0xc0, 0x8a, 0xd0,
	0x85, 0x85, 0xc4, 0xb8, 0xd8, 0x20, 0xfa, 0x24, 0x58, 0x15, 0x18, 0x35, 0x38, 0x82, 0xa0, 0x3c,
	0x21, 0x01, 0x2e, 0xe6, 0x9c, 0xfc, 0x74, 0x09, 0x36, 0xb0, 0x2e, 0x10, 0xd3, 0xc9, 0xee, 0xc4,
	0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1,
	0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x54, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93,
	0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xe1, 0xbe, 0xd3, 0xcf, 0x35, 0xd0, 0xaf, 0xd0, 0x07, 0x7b, 0xbe,
	0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x7f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbc, 0x71, 0x10, 0x22, 0x11, 0x01, 0x00, 0x00,
}

func (m *Itx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Itx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Itx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintItx(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0x32
	}
	if m.Source {
		i--
		if m.Source {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.DestinationHash) > 0 {
		i -= len(m.DestinationHash)
		copy(dAtA[i:], m.DestinationHash)
		i = encodeVarintItx(dAtA, i, uint64(len(m.DestinationHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourceHash) > 0 {
		i -= len(m.SourceHash)
		copy(dAtA[i:], m.SourceHash)
		i = encodeVarintItx(dAtA, i, uint64(len(m.SourceHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintItx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintItx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintItx(dAtA []byte, offset int, v uint64) int {
	offset -= sovItx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Itx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovItx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovItx(uint64(m.Id))
	}
	l = len(m.SourceHash)
	if l > 0 {
		n += 1 + l + sovItx(uint64(l))
	}
	l = len(m.DestinationHash)
	if l > 0 {
		n += 1 + l + sovItx(uint64(l))
	}
	if m.Source {
		n += 2
	}
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovItx(uint64(l))
	}
	return n
}

func sovItx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozItx(x uint64) (n int) {
	return sovItx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Itx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowItx
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
			return fmt.Errorf("proto: Itx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Itx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItx
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
				return ErrInvalidLengthItx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItx
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
				return ErrInvalidLengthItx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItx
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
				return ErrInvalidLengthItx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItx
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
			m.Source = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowItx
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
				return ErrInvalidLengthItx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthItx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipItx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthItx
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
func skipItx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowItx
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
					return 0, ErrIntOverflowItx
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
					return 0, ErrIntOverflowItx
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
				return 0, ErrInvalidLengthItx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupItx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthItx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthItx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowItx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupItx = fmt.Errorf("proto: unexpected end of group")
)
