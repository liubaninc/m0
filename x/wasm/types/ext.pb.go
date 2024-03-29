// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: m0/wasm/ext.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	xmodel "github.com/liubaninc/m0/x/wasm/xmodel"
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

type InputExt struct {
	Bucket    string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	RefTx     string `protobuf:"bytes,3,opt,name=ref_tx,json=refTx,proto3" json:"ref_tx,omitempty"`
	RefMsg    int32  `protobuf:"varint,4,opt,name=ref_msg,json=refMsg,proto3" json:"ref_msg,omitempty"`
	RefOffset int32  `protobuf:"varint,5,opt,name=ref_offset,json=refOffset,proto3" json:"ref_offset,omitempty"`
}

func (m *InputExt) Reset()         { *m = InputExt{} }
func (m *InputExt) String() string { return proto.CompactTextString(m) }
func (*InputExt) ProtoMessage()    {}
func (*InputExt) Descriptor() ([]byte, []int) {
	return fileDescriptor_5339723941bf1e72, []int{0}
}
func (m *InputExt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InputExt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InputExt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InputExt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputExt.Merge(m, src)
}
func (m *InputExt) XXX_Size() int {
	return m.Size()
}
func (m *InputExt) XXX_DiscardUnknown() {
	xxx_messageInfo_InputExt.DiscardUnknown(m)
}

var xxx_messageInfo_InputExt proto.InternalMessageInfo

func (m *InputExt) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *InputExt) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *InputExt) GetRefTx() string {
	if m != nil {
		return m.RefTx
	}
	return ""
}

func (m *InputExt) GetRefMsg() int32 {
	if m != nil {
		return m.RefMsg
	}
	return 0
}

func (m *InputExt) GetRefOffset() int32 {
	if m != nil {
		return m.RefOffset
	}
	return 0
}

type OutputExt struct {
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *OutputExt) Reset()         { *m = OutputExt{} }
func (m *OutputExt) String() string { return proto.CompactTextString(m) }
func (*OutputExt) ProtoMessage()    {}
func (*OutputExt) Descriptor() ([]byte, []int) {
	return fileDescriptor_5339723941bf1e72, []int{1}
}
func (m *OutputExt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutputExt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutputExt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutputExt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputExt.Merge(m, src)
}
func (m *OutputExt) XXX_Size() int {
	return m.Size()
}
func (m *OutputExt) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputExt.DiscardUnknown(m)
}

var xxx_messageInfo_OutputExt proto.InternalMessageInfo

func (m *OutputExt) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *OutputExt) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *OutputExt) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type InvokeRequest struct {
	ModuleName     string                  `protobuf:"bytes,1,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	ContractName   string                  `protobuf:"bytes,2,opt,name=contract_name,json=contractName,proto3" json:"contract_name,omitempty"`
	MethodName     string                  `protobuf:"bytes,3,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	Args           string                  `protobuf:"bytes,4,opt,name=args,proto3" json:"args,omitempty"`
	ResourceLimits []*xmodel.ResourceLimit `protobuf:"bytes,5,rep,name=resource_limits,json=resourceLimits,proto3" json:"resource_limits,omitempty"`
	// amount is the amount transfer to the contract
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *InvokeRequest) Reset()         { *m = InvokeRequest{} }
func (m *InvokeRequest) String() string { return proto.CompactTextString(m) }
func (*InvokeRequest) ProtoMessage()    {}
func (*InvokeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5339723941bf1e72, []int{2}
}
func (m *InvokeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InvokeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InvokeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InvokeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvokeRequest.Merge(m, src)
}
func (m *InvokeRequest) XXX_Size() int {
	return m.Size()
}
func (m *InvokeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InvokeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InvokeRequest proto.InternalMessageInfo

func (m *InvokeRequest) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *InvokeRequest) GetContractName() string {
	if m != nil {
		return m.ContractName
	}
	return ""
}

func (m *InvokeRequest) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *InvokeRequest) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *InvokeRequest) GetResourceLimits() []*xmodel.ResourceLimit {
	if m != nil {
		return m.ResourceLimits
	}
	return nil
}

func (m *InvokeRequest) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterType((*InputExt)(nil), "liubaninc.m0.wasm.InputExt")
	proto.RegisterType((*OutputExt)(nil), "liubaninc.m0.wasm.OutputExt")
	proto.RegisterType((*InvokeRequest)(nil), "liubaninc.m0.wasm.InvokeRequest")
}

func init() { proto.RegisterFile("m0/wasm/ext.proto", fileDescriptor_5339723941bf1e72) }

var fileDescriptor_5339723941bf1e72 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0x9b, 0xda, 0x90, 0x69, 0x0b, 0x74, 0x54, 0x8a, 0xa9, 0x84, 0x13, 0x05, 0x16, 0xd9,
	0xe0, 0x49, 0x61, 0xcf, 0xa2, 0x88, 0x45, 0x55, 0xa0, 0x92, 0xc5, 0x8a, 0x4d, 0x34, 0x76, 0xae,
	0x5d, 0x2b, 0x1e, 0x4f, 0x98, 0x47, 0x70, 0xd7, 0xfc, 0x00, 0xdf, 0xc1, 0x37, 0xf0, 0x01, 0x5d,
	0x76, 0xc9, 0x0a, 0x50, 0xf2, 0x23, 0x68, 0x1e, 0xa0, 0xb2, 0xec, 0xea, 0xde, 0x7b, 0xce, 0xf1,
	0xf1, 0x1d, 0x9d, 0x8b, 0xf6, 0xd9, 0x94, 0x7c, 0xa6, 0x92, 0x11, 0xe8, 0x54, 0xba, 0x14, 0x5c,
	0x71, 0xbc, 0xdf, 0xd4, 0x3a, 0xa7, 0x6d, 0xdd, 0x16, 0x29, 0x9b, 0xa6, 0x86, 0x3c, 0x3a, 0xa8,
	0x78, 0xc5, 0x2d, 0x4b, 0x4c, 0xe7, 0x84, 0x47, 0x49, 0xc1, 0x25, 0xe3, 0x92, 0xe4, 0x54, 0x02,
	0x59, 0x1d, 0xe7, 0xa0, 0xe8, 0x31, 0x29, 0x78, 0xdd, 0x7a, 0xfe, 0x90, 0x4d, 0x49, 0xc7, 0xf8,
	0x1c, 0x1a, 0x5f, 0x1c, 0x3e, 0xfe, 0x12, 0xa0, 0xbb, 0xa7, 0xed, 0x52, 0xab, 0x37, 0x9d, 0xc2,
	0x87, 0x28, 0xca, 0x75, 0xb1, 0x00, 0x15, 0x07, 0xa3, 0x60, 0x32, 0xc8, 0xfc, 0x84, 0x1f, 0xa0,
	0xfe, 0x02, 0x2e, 0xe3, 0x2d, 0x0b, 0x9a, 0x16, 0x3f, 0x44, 0x91, 0x80, 0x72, 0xa6, 0xba, 0xb8,
	0x6f, 0xc1, 0x50, 0x40, 0xf9, 0xa1, 0xc3, 0x8f, 0xd0, 0x1d, 0x03, 0x33, 0x59, 0xc5, 0xdb, 0xa3,
	0x60, 0x12, 0x66, 0x46, 0xf5, 0x4e, 0x56, 0xf8, 0x09, 0x42, 0x86, 0xe0, 0x65, 0x29, 0x41, 0xc5,
	0xa1, 0xe5, 0x06, 0x02, 0xca, 0x73, 0x0b, 0x8c, 0xcf, 0xd0, 0xe0, 0x5c, 0xab, 0x5b, 0x6f, 0x71,
	0x80, 0xc2, 0x15, 0x6d, 0x34, 0xd8, 0x25, 0x76, 0x33, 0x37, 0x8c, 0xbf, 0x6f, 0xa1, 0xbd, 0xd3,
	0x76, 0xc5, 0x17, 0x90, 0xc1, 0x27, 0x0d, 0x52, 0xe1, 0x21, 0xda, 0x61, 0x7c, 0xae, 0x1b, 0x98,
	0xb5, 0x94, 0x81, 0xb7, 0x45, 0x0e, 0x7a, 0x4f, 0x19, 0xe0, 0xa7, 0x68, 0xaf, 0xe0, 0xad, 0x12,
	0xb4, 0x50, 0x4e, 0xe2, 0x7e, 0xb2, 0xfb, 0x17, 0xb4, 0x22, 0xe3, 0x02, 0xea, 0x82, 0xcf, 0x9d,
	0xa4, 0xef, 0x5d, 0x2c, 0x64, 0x05, 0x18, 0x6d, 0x53, 0x51, 0x49, 0xfb, 0xf4, 0x41, 0x66, 0x7b,
	0x7c, 0x86, 0xee, 0x0b, 0x90, 0x5c, 0x8b, 0x02, 0x66, 0x4d, 0xcd, 0x6a, 0x25, 0xe3, 0x70, 0xd4,
	0x9f, 0xec, 0xbc, 0x18, 0xa7, 0xff, 0x45, 0xeb, 0x43, 0xc9, 0xbc, 0xf6, 0xad, 0x91, 0x66, 0xf7,
	0xc4, 0xcd, 0x51, 0xe2, 0x02, 0x45, 0x94, 0x71, 0xdd, 0xaa, 0x38, 0xb2, 0x1e, 0x8f, 0x53, 0x97,
	0x7a, 0x6a, 0x52, 0x4f, 0x7d, 0xea, 0xe9, 0x6b, 0x5e, 0xb7, 0x27, 0xd3, 0xab, 0x9f, 0xc3, 0xde,
	0xb7, 0x5f, 0xc3, 0x49, 0x55, 0xab, 0x0b, 0x9d, 0xa7, 0x05, 0x67, 0xc4, 0x9f, 0x88, 0x2b, 0xcf,
	0xe5, 0x7c, 0x41, 0xd4, 0xe5, 0x12, 0xa4, 0xfd, 0x40, 0x66, 0xde, 0xfa, 0xe4, 0xd5, 0xd5, 0x3a,
	0x09, 0xae, 0xd7, 0x49, 0xf0, 0x7b, 0x9d, 0x04, 0x5f, 0x37, 0x49, 0xef, 0x7a, 0x93, 0xf4, 0x7e,
	0x6c, 0x92, 0xde, 0xc7, 0x67, 0x37, 0xbc, 0xfe, 0x2d, 0x4f, 0xcc, 0x61, 0xb9, 0xb3, 0xb5, 0x6e,
	0x79, 0x64, 0x0f, 0xeb, 0xe5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0xe6, 0xcb, 0xc0, 0xce,
	0x02, 0x00, 0x00,
}

func (m *InputExt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InputExt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InputExt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RefOffset != 0 {
		i = encodeVarintExt(dAtA, i, uint64(m.RefOffset))
		i--
		dAtA[i] = 0x28
	}
	if m.RefMsg != 0 {
		i = encodeVarintExt(dAtA, i, uint64(m.RefMsg))
		i--
		dAtA[i] = 0x20
	}
	if len(m.RefTx) > 0 {
		i -= len(m.RefTx)
		copy(dAtA[i:], m.RefTx)
		i = encodeVarintExt(dAtA, i, uint64(len(m.RefTx)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintExt(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Bucket) > 0 {
		i -= len(m.Bucket)
		copy(dAtA[i:], m.Bucket)
		i = encodeVarintExt(dAtA, i, uint64(len(m.Bucket)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OutputExt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutputExt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutputExt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintExt(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintExt(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Bucket) > 0 {
		i -= len(m.Bucket)
		copy(dAtA[i:], m.Bucket)
		i = encodeVarintExt(dAtA, i, uint64(len(m.Bucket)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InvokeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InvokeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InvokeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ResourceLimits) > 0 {
		for iNdEx := len(m.ResourceLimits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ResourceLimits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintExt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Args) > 0 {
		i -= len(m.Args)
		copy(dAtA[i:], m.Args)
		i = encodeVarintExt(dAtA, i, uint64(len(m.Args)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MethodName) > 0 {
		i -= len(m.MethodName)
		copy(dAtA[i:], m.MethodName)
		i = encodeVarintExt(dAtA, i, uint64(len(m.MethodName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ContractName) > 0 {
		i -= len(m.ContractName)
		copy(dAtA[i:], m.ContractName)
		i = encodeVarintExt(dAtA, i, uint64(len(m.ContractName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintExt(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExt(dAtA []byte, offset int, v uint64) int {
	offset -= sovExt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InputExt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bucket)
	if l > 0 {
		n += 1 + l + sovExt(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovExt(uint64(l))
	}
	l = len(m.RefTx)
	if l > 0 {
		n += 1 + l + sovExt(uint64(l))
	}
	if m.RefMsg != 0 {
		n += 1 + sovExt(uint64(m.RefMsg))
	}
	if m.RefOffset != 0 {
		n += 1 + sovExt(uint64(m.RefOffset))
	}
	return n
}

func (m *OutputExt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bucket)
	if l > 0 {
		n += 1 + l + sovExt(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovExt(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovExt(uint64(l))
	}
	return n
}

func (m *InvokeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovExt(uint64(l))
	}
	l = len(m.ContractName)
	if l > 0 {
		n += 1 + l + sovExt(uint64(l))
	}
	l = len(m.MethodName)
	if l > 0 {
		n += 1 + l + sovExt(uint64(l))
	}
	l = len(m.Args)
	if l > 0 {
		n += 1 + l + sovExt(uint64(l))
	}
	if len(m.ResourceLimits) > 0 {
		for _, e := range m.ResourceLimits {
			l = e.Size()
			n += 1 + l + sovExt(uint64(l))
		}
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovExt(uint64(l))
		}
	}
	return n
}

func sovExt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExt(x uint64) (n int) {
	return sovExt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InputExt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExt
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
			return fmt.Errorf("proto: InputExt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InputExt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bucket", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bucket = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefTx", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefTx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefMsg", wireType)
			}
			m.RefMsg = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefOffset", wireType)
			}
			m.RefOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
		default:
			iNdEx = preIndex
			skippy, err := skipExt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExt
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
func (m *OutputExt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExt
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
			return fmt.Errorf("proto: OutputExt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutputExt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bucket", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bucket = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthExt
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
			skippy, err := skipExt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExt
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
func (m *InvokeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExt
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
			return fmt.Errorf("proto: InvokeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InvokeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MethodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceLimits = append(m.ResourceLimits, &xmodel.ResourceLimit{})
			if err := m.ResourceLimits[len(m.ResourceLimits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExt
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
				return ErrInvalidLengthExt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExt
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
func skipExt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExt
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
					return 0, ErrIntOverflowExt
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
					return 0, ErrIntOverflowExt
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
				return 0, ErrInvalidLengthExt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExt = fmt.Errorf("proto: unexpected end of group")
)
