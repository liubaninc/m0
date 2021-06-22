// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mibc/itx.proto

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

type Itx struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Hash    string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Source  bool   `protobuf:"varint,4,opt,name=source,proto3" json:"source,omitempty"`
	Chain   string `protobuf:"bytes,5,opt,name=chain,proto3" json:"chain,omitempty"`
	Log     string `protobuf:"bytes,6,opt,name=log,proto3" json:"log,omitempty"`
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

func (m *Itx) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Itx) GetSource() bool {
	if m != nil {
		return m.Source
	}
	return false
}

func (m *Itx) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
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
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xb1, 0x4e, 0x84, 0x40,
	0x10, 0x86, 0x59, 0xe0, 0x50, 0xa7, 0xb8, 0xe8, 0xe4, 0x62, 0x36, 0x16, 0x1b, 0x62, 0x2c, 0xa8,
	0xd8, 0x4b, 0xec, 0x2d, 0xec, 0x6c, 0x29, 0xed, 0x60, 0x8f, 0xc0, 0x24, 0x07, 0x73, 0x81, 0x25,
	0xc1, 0xda, 0x17, 0xf0, 0xb1, 0x2c, 0xaf, 0xb4, 0x34, 0xf0, 0x22, 0x86, 0x45, 0xed, 0xfe, 0x6f,
	0xe6, 0x6b, 0x3e, 0xd8, 0x36, 0x54, 0x18, 0x4d, 0x76, 0x4c, 0x4f, 0x1d, 0x5b, 0xc6, 0x9b, 0x23,
	0x0d, 0x45, 0xde, 0x52, 0x6b, 0xd2, 0x66, 0x9f, 0x2e, 0xcf, 0xbb, 0x5d, 0xc5, 0x15, 0xbb, 0xaf,
	0x5e, 0xd6, 0x2a, 0xde, 0xbf, 0x0b, 0x08, 0x5e, 0xec, 0x88, 0x12, 0x2e, 0x4c, 0x57, 0xe6, 0x96,
	0x3b, 0x29, 0x62, 0x91, 0x5c, 0x65, 0x7f, 0x88, 0x5b, 0xf0, 0xe9, 0x20, 0xfd, 0x58, 0x24, 0x61,
	0xe6, 0xd3, 0x01, 0x11, 0xc2, 0x3a, 0xef, 0x6b, 0x19, 0x38, 0xcd, 0x6d, 0xbc, 0x85, 0xa8, 0xe7,
	0xa1, 0x33, 0xa5, 0x0c, 0x63, 0x91, 0x5c, 0x66, 0xbf, 0x84, 0x3b, 0xd8, 0x98, 0x3a, 0xa7, 0x56,
	0x6e, 0x9c, 0xbc, 0x02, 0x5e, 0x43, 0x70, 0xe4, 0x4a, 0x46, 0xee, 0xb6, 0xcc, 0xe7, 0xa7, 0xcf,
	0x49, 0x89, 0xf3, 0xa4, 0xc4, 0xf7, 0xa4, 0xc4, 0xc7, 0xac, 0xbc, 0xf3, 0xac, 0xbc, 0xaf, 0x59,
	0x79, 0xaf, 0x0f, 0x15, 0xd9, 0x7a, 0x28, 0x52, 0xc3, 0x8d, 0xfe, 0x6f, 0xd2, 0xcd, 0x5e, 0x8f,
	0xda, 0x25, 0xdb, 0xb7, 0x53, 0xd9, 0x17, 0x91, 0x8b, 0x79, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xc8, 0xc9, 0x16, 0x8a, 0x07, 0x01, 0x00, 0x00,
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
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintItx(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Source {
		i--
		if m.Source {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintItx(dAtA, i, uint64(len(m.Hash)))
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
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovItx(uint64(l))
	}
	if m.Source {
		n += 2
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovItx(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
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
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
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
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
