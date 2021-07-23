// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// this line is used by starport scaffolding # 3
type QueryGetValidatorRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryGetValidatorRequest) Reset()         { *m = QueryGetValidatorRequest{} }
func (m *QueryGetValidatorRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetValidatorRequest) ProtoMessage()    {}
func (*QueryGetValidatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b4d73ed8fedd8d, []int{0}
}
func (m *QueryGetValidatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetValidatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetValidatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetValidatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetValidatorRequest.Merge(m, src)
}
func (m *QueryGetValidatorRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetValidatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetValidatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetValidatorRequest proto.InternalMessageInfo

func (m *QueryGetValidatorRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryGetValidatorResponse struct {
	Validator *Validator `protobuf:"bytes,1,opt,name=Validator,proto3" json:"Validator,omitempty"`
}

func (m *QueryGetValidatorResponse) Reset()         { *m = QueryGetValidatorResponse{} }
func (m *QueryGetValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetValidatorResponse) ProtoMessage()    {}
func (*QueryGetValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b4d73ed8fedd8d, []int{1}
}
func (m *QueryGetValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetValidatorResponse.Merge(m, src)
}
func (m *QueryGetValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetValidatorResponse proto.InternalMessageInfo

func (m *QueryGetValidatorResponse) GetValidator() *Validator {
	if m != nil {
		return m.Validator
	}
	return nil
}

type QueryAllValidatorRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllValidatorRequest) Reset()         { *m = QueryAllValidatorRequest{} }
func (m *QueryAllValidatorRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllValidatorRequest) ProtoMessage()    {}
func (*QueryAllValidatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b4d73ed8fedd8d, []int{2}
}
func (m *QueryAllValidatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllValidatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllValidatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllValidatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllValidatorRequest.Merge(m, src)
}
func (m *QueryAllValidatorRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllValidatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllValidatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllValidatorRequest proto.InternalMessageInfo

func (m *QueryAllValidatorRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllValidatorResponse struct {
	Validator  []*Validator        `protobuf:"bytes,1,rep,name=Validator,proto3" json:"Validator,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllValidatorResponse) Reset()         { *m = QueryAllValidatorResponse{} }
func (m *QueryAllValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllValidatorResponse) ProtoMessage()    {}
func (*QueryAllValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b4d73ed8fedd8d, []int{3}
}
func (m *QueryAllValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllValidatorResponse.Merge(m, src)
}
func (m *QueryAllValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllValidatorResponse proto.InternalMessageInfo

func (m *QueryAllValidatorResponse) GetValidator() []*Validator {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *QueryAllValidatorResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetValidatorRequest)(nil), "liubaninc.m0.validator.QueryGetValidatorRequest")
	proto.RegisterType((*QueryGetValidatorResponse)(nil), "liubaninc.m0.validator.QueryGetValidatorResponse")
	proto.RegisterType((*QueryAllValidatorRequest)(nil), "liubaninc.m0.validator.QueryAllValidatorRequest")
	proto.RegisterType((*QueryAllValidatorResponse)(nil), "liubaninc.m0.validator.QueryAllValidatorResponse")
}

func init() { proto.RegisterFile("validator/query.proto", fileDescriptor_31b4d73ed8fedd8d) }

var fileDescriptor_31b4d73ed8fedd8d = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6b, 0xdb, 0x30,
	0x18, 0xc6, 0xa3, 0x8c, 0x6d, 0x44, 0xdb, 0x49, 0xb0, 0x91, 0x84, 0x61, 0x16, 0x6f, 0x6c, 0xcb,
	0xc6, 0xa4, 0xfc, 0xd9, 0x7d, 0x64, 0x83, 0xe5, 0xba, 0xe6, 0xd0, 0x43, 0x29, 0x04, 0x39, 0x11,
	0xae, 0x41, 0xb6, 0x1c, 0x4b, 0x4e, 0x1b, 0x4a, 0x2f, 0xfd, 0x04, 0x85, 0xde, 0x7b, 0xcb, 0x77,
	0xe9, 0x31, 0xd0, 0x4b, 0x8f, 0x25, 0xe9, 0x07, 0x29, 0xb1, 0x15, 0xdb, 0xa1, 0x0e, 0x26, 0xb7,
	0x28, 0x7a, 0x9f, 0xe7, 0xfd, 0xbd, 0xcf, 0x2b, 0xc3, 0x77, 0x53, 0xca, 0x9d, 0x31, 0x55, 0x22,
	0x20, 0x93, 0x90, 0x05, 0x33, 0xec, 0x07, 0x42, 0x09, 0xf4, 0x9e, 0x3b, 0xa1, 0x45, 0x3d, 0xc7,
	0x1b, 0x61, 0xb7, 0x85, 0x93, 0x9a, 0xfa, 0x07, 0x5b, 0x08, 0x9b, 0x33, 0x42, 0x7d, 0x87, 0x50,
	0xcf, 0x13, 0x8a, 0x2a, 0x47, 0x78, 0x32, 0x56, 0xd5, 0xbf, 0x8f, 0x84, 0x74, 0x85, 0x24, 0x16,
	0x95, 0x2c, 0xb6, 0x23, 0xd3, 0xb6, 0xc5, 0x14, 0x6d, 0x13, 0x9f, 0xda, 0x8e, 0x17, 0x15, 0xeb,
	0xda, 0xcf, 0x69, 0x63, 0x4e, 0xa5, 0x1a, 0x26, 0xc7, 0xa1, 0x2f, 0x4e, 0x59, 0xa0, 0xab, 0x6a,
	0x69, 0x55, 0xf2, 0x2b, 0xbe, 0x32, 0x7f, 0xc1, 0xea, 0xc1, 0xba, 0x45, 0x9f, 0xa9, 0xc3, 0xcd,
	0xd5, 0x80, 0x4d, 0x42, 0x26, 0x15, 0xaa, 0xc2, 0xd7, 0x74, 0x3c, 0x0e, 0x98, 0x94, 0x55, 0xf0,
	0x11, 0x7c, 0xab, 0x0c, 0x36, 0x47, 0xf3, 0x18, 0xd6, 0x72, 0x54, 0xd2, 0x17, 0x9e, 0x64, 0xe8,
	0x37, 0xac, 0x24, 0x7f, 0x46, 0xc2, 0x37, 0x9d, 0x06, 0xce, 0x4f, 0x02, 0xa7, 0xea, 0x54, 0x63,
	0x5a, 0x9a, 0xa9, 0xc7, 0xf9, 0x33, 0xa6, 0x7f, 0x10, 0xa6, 0x21, 0x68, 0xf7, 0x2f, 0x38, 0x4e,
	0x0c, 0xaf, 0x13, 0xc3, 0xf1, 0x02, 0x74, 0x62, 0xf8, 0x3f, 0xb5, 0x99, 0xd6, 0x0e, 0x32, 0x4a,
	0x73, 0x0e, 0xf4, 0x08, 0xdb, 0x4d, 0xf2, 0x47, 0x78, 0xb1, 0xef, 0x08, 0xa8, 0xbf, 0x85, 0x59,
	0x8e, 0x30, 0xbf, 0x16, 0x62, 0xc6, 0xdd, 0xb3, 0x9c, 0x9d, 0x45, 0x19, 0xbe, 0x8c, 0x38, 0xd1,
	0x1c, 0x64, 0xa0, 0x50, 0x6b, 0x17, 0xce, 0xae, 0x6d, 0xd6, 0xdb, 0x7b, 0x28, 0x62, 0x10, 0xb3,
	0x7b, 0x79, 0xf7, 0x78, 0x5d, 0xfe, 0x89, 0x7e, 0x90, 0x44, 0x4a, 0xdc, 0x16, 0xc9, 0x79, 0x4d,
	0xe4, 0x5c, 0x3f, 0x8d, 0x0b, 0x74, 0x03, 0xe0, 0xdb, 0xc4, 0xaa, 0xc7, 0x79, 0x01, 0x6a, 0xce,
	0x92, 0x0b, 0x50, 0xf3, 0x36, 0x66, 0x36, 0x23, 0xd4, 0x4f, 0xa8, 0x51, 0x88, 0xfa, 0xe7, 0xef,
	0xed, 0xd2, 0x00, 0x8b, 0xa5, 0x01, 0x1e, 0x96, 0x06, 0xb8, 0x5a, 0x19, 0xa5, 0xc5, 0xca, 0x28,
	0xdd, 0xaf, 0x8c, 0xd2, 0x51, 0xd3, 0x76, 0xd4, 0x49, 0x68, 0xe1, 0x91, 0x70, 0xb7, 0x6d, 0xce,
	0x32, 0x46, 0x6a, 0xe6, 0x33, 0x69, 0xbd, 0x8a, 0x3e, 0x9f, 0xee, 0x53, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0xa7, 0xc9, 0x31, 0xfa, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a validator by id.
	Validator(ctx context.Context, in *QueryGetValidatorRequest, opts ...grpc.CallOption) (*QueryGetValidatorResponse, error)
	// Queries a list of validator items.
	ValidatorAll(ctx context.Context, in *QueryAllValidatorRequest, opts ...grpc.CallOption) (*QueryAllValidatorResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Validator(ctx context.Context, in *QueryGetValidatorRequest, opts ...grpc.CallOption) (*QueryGetValidatorResponse, error) {
	out := new(QueryGetValidatorResponse)
	err := c.cc.Invoke(ctx, "/liubaninc.m0.validator.Query/Validator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ValidatorAll(ctx context.Context, in *QueryAllValidatorRequest, opts ...grpc.CallOption) (*QueryAllValidatorResponse, error) {
	out := new(QueryAllValidatorResponse)
	err := c.cc.Invoke(ctx, "/liubaninc.m0.validator.Query/ValidatorAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a validator by id.
	Validator(context.Context, *QueryGetValidatorRequest) (*QueryGetValidatorResponse, error)
	// Queries a list of validator items.
	ValidatorAll(context.Context, *QueryAllValidatorRequest) (*QueryAllValidatorResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Validator(ctx context.Context, req *QueryGetValidatorRequest) (*QueryGetValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validator not implemented")
}
func (*UnimplementedQueryServer) ValidatorAll(ctx context.Context, req *QueryAllValidatorRequest) (*QueryAllValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Validator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Validator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liubaninc.m0.validator.Query/Validator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Validator(ctx, req.(*QueryGetValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ValidatorAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ValidatorAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liubaninc.m0.validator.Query/ValidatorAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ValidatorAll(ctx, req.(*QueryAllValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "liubaninc.m0.validator.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validator",
			Handler:    _Query_Validator_Handler,
		},
		{
			MethodName: "ValidatorAll",
			Handler:    _Query_ValidatorAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "validator/query.proto",
}

func (m *QueryGetValidatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetValidatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetValidatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Validator != nil {
		{
			size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllValidatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllValidatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllValidatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validator) > 0 {
		for iNdEx := len(m.Validator) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validator[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetValidatorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Validator != nil {
		l = m.Validator.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllValidatorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Validator) > 0 {
		for _, e := range m.Validator {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetValidatorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetValidatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetValidatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validator == nil {
				m.Validator = &Validator{}
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllValidatorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllValidatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllValidatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = append(m.Validator, &Validator{})
			if err := m.Validator[len(m.Validator)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)