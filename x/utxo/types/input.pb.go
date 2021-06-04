// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: utxo/input.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"

	types "github.com/cosmos/cosmos-sdk/types"
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

type Input struct {
	RefTx        string     `protobuf:"bytes,1,opt,name=ref_tx,json=refTx,proto3" json:"ref_tx,omitempty"`
	RefMsg       int32      `protobuf:"varint,2,opt,name=ref_msg,json=refMsg,proto3" json:"ref_msg,omitempty"`
	RefOffset    int32      `protobuf:"varint,3,opt,name=ref_offset,json=refOffset,proto3" json:"ref_offset,omitempty"`
	FromAddr     string     `protobuf:"bytes,4,opt,name=from_addr,json=fromAddr,proto3" json:"from_addr,omitempty"`
	Amount       types.Coin `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount"`
	FrozenHeight int64      `protobuf:"varint,7,opt,name=frozen_height,json=frozenHeight,proto3" json:"frozen_height,omitempty"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f017d14d9e0952f, []int{0}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Input.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return m.Size()
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetRefTx() string {
	if m != nil {
		return m.RefTx
	}
	return ""
}

func (m *Input) GetRefMsg() int32 {
	if m != nil {
		return m.RefMsg
	}
	return 0
}

func (m *Input) GetRefOffset() int32 {
	if m != nil {
		return m.RefOffset
	}
	return 0
}

func (m *Input) GetFromAddr() string {
	if m != nil {
		return m.FromAddr
	}
	return ""
}

func (m *Input) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *Input) GetFrozenHeight() int64 {
	if m != nil {
		return m.FrozenHeight
	}
	return 0
}

type Output struct {
	ToAddr       string     `protobuf:"bytes,1,opt,name=to_addr,json=toAddr,proto3" json:"to_addr,omitempty"`
	Amount       types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
	FrozenHeight int64      `protobuf:"varint,4,opt,name=frozen_height,json=frozenHeight,proto3" json:"frozen_height,omitempty"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f017d14d9e0952f, []int{1}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Output.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(m, src)
}
func (m *Output) XXX_Size() int {
	return m.Size()
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetToAddr() string {
	if m != nil {
		return m.ToAddr
	}
	return ""
}

func (m *Output) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *Output) GetFrozenHeight() int64 {
	if m != nil {
		return m.FrozenHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Input)(nil), "liubaninc.m0.utxo.Input")
	proto.RegisterType((*Output)(nil), "liubaninc.m0.utxo.Output")
}

func init() { proto.RegisterFile("utxo/input.proto", fileDescriptor_2f017d14d9e0952f) }

var fileDescriptor_2f017d14d9e0952f = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x73, 0x6d, 0x93, 0xda, 0x53, 0x41, 0x83, 0xd2, 0x58, 0x31, 0x86, 0xea, 0x90, 0xe9,
	0xae, 0xd5, 0xc1, 0x4d, 0xb0, 0x2e, 0x3a, 0x48, 0x21, 0x38, 0xb9, 0x94, 0xfc, 0xb9, 0xa4, 0x01,
	0x93, 0xb7, 0x5c, 0x2e, 0x12, 0x5d, 0xfc, 0x0a, 0x7e, 0xac, 0x8e, 0x75, 0x73, 0x12, 0x69, 0xbf,
	0x88, 0xdc, 0xa5, 0x38, 0xe8, 0xe8, 0x76, 0x3c, 0xbf, 0x87, 0xe3, 0xf7, 0xf2, 0xe0, 0x9d, 0x52,
	0x54, 0x40, 0xd3, 0x7c, 0x56, 0x0a, 0x32, 0xe3, 0x20, 0xc0, 0xdc, 0x7d, 0x4c, 0xcb, 0xc0, 0xcf,
	0xd3, 0x3c, 0x24, 0xd9, 0x80, 0x48, 0xdc, 0xdb, 0x4b, 0x20, 0x01, 0x45, 0xa9, 0x7c, 0xd5, 0xc5,
	0x9e, 0x1d, 0x42, 0x91, 0x41, 0x41, 0x03, 0xbf, 0x60, 0xf4, 0x69, 0x18, 0x30, 0xe1, 0x0f, 0x69,
	0x08, 0x69, 0x5e, 0xf3, 0xfe, 0x3b, 0xc2, 0xfa, 0xad, 0xfc, 0xd8, 0xdc, 0xc7, 0x06, 0x67, 0xf1,
	0x44, 0x54, 0x16, 0x72, 0x90, 0xdb, 0xf1, 0x74, 0xce, 0xe2, 0xfb, 0xca, 0xec, 0xe2, 0xb6, 0x8c,
	0xb3, 0x22, 0xb1, 0x1a, 0x0e, 0x72, 0x75, 0x4f, 0xb6, 0xee, 0x8a, 0xc4, 0x3c, 0xc2, 0x58, 0x02,
	0x88, 0xe3, 0x82, 0x09, 0xab, 0xa9, 0x58, 0x87, 0xb3, 0x78, 0xac, 0x02, 0xf3, 0x10, 0x77, 0x62,
	0x0e, 0xd9, 0xc4, 0x8f, 0x22, 0x6e, 0xb5, 0xd4, 0x8f, 0x1b, 0x32, 0xb8, 0x8a, 0x22, 0x6e, 0x5e,
	0x60, 0xc3, 0xcf, 0xa0, 0xcc, 0x85, 0xa5, 0x3b, 0xc8, 0xdd, 0x3c, 0x3b, 0x20, 0xb5, 0x26, 0x91,
	0x9a, 0x64, 0xad, 0x49, 0xae, 0x21, 0xcd, 0x47, 0xad, 0xf9, 0xe7, 0xb1, 0xe6, 0xad, 0xeb, 0xe6,
	0x09, 0xde, 0x8e, 0x39, 0xbc, 0xb0, 0x7c, 0x32, 0x65, 0x69, 0x32, 0x15, 0x56, 0xdb, 0x41, 0x6e,
	0xd3, 0xdb, 0xaa, 0xc3, 0x1b, 0x95, 0xf5, 0x5f, 0xb1, 0x31, 0x2e, 0x85, 0xbc, 0xa9, 0x8b, 0xdb,
	0x02, 0x6a, 0x85, 0xfa, 0x28, 0x43, 0xc0, 0x2f, 0x81, 0xc6, 0x3f, 0x05, 0x5a, 0x7f, 0x05, 0x46,
	0x97, 0xf3, 0xa5, 0x8d, 0x16, 0x4b, 0x1b, 0x7d, 0x2d, 0x6d, 0xf4, 0xb6, 0xb2, 0xb5, 0xc5, 0xca,
	0xd6, 0x3e, 0x56, 0xb6, 0xf6, 0x70, 0x9a, 0xa4, 0x62, 0x5a, 0x06, 0x24, 0x84, 0x8c, 0xfe, 0x4c,
	0x48, 0xb3, 0x01, 0xad, 0xa8, 0xda, 0x58, 0x3c, 0xcf, 0x58, 0x11, 0x18, 0x6a, 0x9b, 0xf3, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x3e, 0x12, 0x95, 0xf8, 0x01, 0x00, 0x00,
}

func (m *Input) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Input) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Input) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FrozenHeight != 0 {
		i = encodeVarintInput(dAtA, i, uint64(m.FrozenHeight))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintInput(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.FromAddr) > 0 {
		i -= len(m.FromAddr)
		copy(dAtA[i:], m.FromAddr)
		i = encodeVarintInput(dAtA, i, uint64(len(m.FromAddr)))
		i--
		dAtA[i] = 0x22
	}
	if m.RefOffset != 0 {
		i = encodeVarintInput(dAtA, i, uint64(m.RefOffset))
		i--
		dAtA[i] = 0x18
	}
	if m.RefMsg != 0 {
		i = encodeVarintInput(dAtA, i, uint64(m.RefMsg))
		i--
		dAtA[i] = 0x10
	}
	if len(m.RefTx) > 0 {
		i -= len(m.RefTx)
		copy(dAtA[i:], m.RefTx)
		i = encodeVarintInput(dAtA, i, uint64(len(m.RefTx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Output) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Output) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Output) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FrozenHeight != 0 {
		i = encodeVarintInput(dAtA, i, uint64(m.FrozenHeight))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintInput(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ToAddr) > 0 {
		i -= len(m.ToAddr)
		copy(dAtA[i:], m.ToAddr)
		i = encodeVarintInput(dAtA, i, uint64(len(m.ToAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInput(dAtA []byte, offset int, v uint64) int {
	offset -= sovInput(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Input) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RefTx)
	if l > 0 {
		n += 1 + l + sovInput(uint64(l))
	}
	if m.RefMsg != 0 {
		n += 1 + sovInput(uint64(m.RefMsg))
	}
	if m.RefOffset != 0 {
		n += 1 + sovInput(uint64(m.RefOffset))
	}
	l = len(m.FromAddr)
	if l > 0 {
		n += 1 + l + sovInput(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovInput(uint64(l))
	if m.FrozenHeight != 0 {
		n += 1 + sovInput(uint64(m.FrozenHeight))
	}
	return n
}

func (m *Output) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ToAddr)
	if l > 0 {
		n += 1 + l + sovInput(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovInput(uint64(l))
	if m.FrozenHeight != 0 {
		n += 1 + sovInput(uint64(m.FrozenHeight))
	}
	return n
}

func sovInput(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInput(x uint64) (n int) {
	return sovInput(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Input) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInput
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
			return fmt.Errorf("proto: Input: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Input: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefTx", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
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
				return ErrInvalidLengthInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefTx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefMsg", wireType)
			}
			m.RefMsg = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RefMsg |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefOffset", wireType)
			}
			m.RefOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RefOffset |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
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
				return ErrInvalidLengthInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
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
				return ErrInvalidLengthInput
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrozenHeight", wireType)
			}
			m.FrozenHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FrozenHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInput(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInput
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
func (m *Output) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInput
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
			return fmt.Errorf("proto: Output: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Output: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
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
				return ErrInvalidLengthInput
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
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
				return ErrInvalidLengthInput
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInput
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrozenHeight", wireType)
			}
			m.FrozenHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInput
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FrozenHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInput(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInput
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
func skipInput(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInput
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
					return 0, ErrIntOverflowInput
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
					return 0, ErrIntOverflowInput
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
				return 0, ErrInvalidLengthInput
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInput
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInput
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInput        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInput          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInput = fmt.Errorf("proto: unexpected end of group")
)
