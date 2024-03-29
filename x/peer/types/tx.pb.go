// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: m0/peer/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// this line is used by starport scaffolding # proto/tx/message
type MsgCreatePeerID struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index         string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	CertIssuer    string `protobuf:"bytes,3,opt,name=certIssuer,proto3" json:"certIssuer,omitempty"`
	CertSerialNum string `protobuf:"bytes,4,opt,name=certSerialNum,proto3" json:"certSerialNum,omitempty"`
}

func (m *MsgCreatePeerID) Reset()         { *m = MsgCreatePeerID{} }
func (m *MsgCreatePeerID) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePeerID) ProtoMessage()    {}
func (*MsgCreatePeerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e75709a1b06ee061, []int{0}
}
func (m *MsgCreatePeerID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePeerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePeerID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePeerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePeerID.Merge(m, src)
}
func (m *MsgCreatePeerID) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePeerID) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePeerID.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePeerID proto.InternalMessageInfo

func (m *MsgCreatePeerID) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePeerID) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *MsgCreatePeerID) GetCertIssuer() string {
	if m != nil {
		return m.CertIssuer
	}
	return ""
}

func (m *MsgCreatePeerID) GetCertSerialNum() string {
	if m != nil {
		return m.CertSerialNum
	}
	return ""
}

type MsgCreatePeerIDResponse struct {
}

func (m *MsgCreatePeerIDResponse) Reset()         { *m = MsgCreatePeerIDResponse{} }
func (m *MsgCreatePeerIDResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePeerIDResponse) ProtoMessage()    {}
func (*MsgCreatePeerIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e75709a1b06ee061, []int{1}
}
func (m *MsgCreatePeerIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePeerIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePeerIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePeerIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePeerIDResponse.Merge(m, src)
}
func (m *MsgCreatePeerIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePeerIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePeerIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePeerIDResponse proto.InternalMessageInfo

type MsgUpdatePeerID struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index         string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	CertIssuer    string `protobuf:"bytes,3,opt,name=certIssuer,proto3" json:"certIssuer,omitempty"`
	CertSerialNum string `protobuf:"bytes,4,opt,name=certSerialNum,proto3" json:"certSerialNum,omitempty"`
}

func (m *MsgUpdatePeerID) Reset()         { *m = MsgUpdatePeerID{} }
func (m *MsgUpdatePeerID) String() string { return proto.CompactTextString(m) }
func (*MsgUpdatePeerID) ProtoMessage()    {}
func (*MsgUpdatePeerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e75709a1b06ee061, []int{2}
}
func (m *MsgUpdatePeerID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdatePeerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdatePeerID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdatePeerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdatePeerID.Merge(m, src)
}
func (m *MsgUpdatePeerID) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdatePeerID) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdatePeerID.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdatePeerID proto.InternalMessageInfo

func (m *MsgUpdatePeerID) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdatePeerID) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *MsgUpdatePeerID) GetCertIssuer() string {
	if m != nil {
		return m.CertIssuer
	}
	return ""
}

func (m *MsgUpdatePeerID) GetCertSerialNum() string {
	if m != nil {
		return m.CertSerialNum
	}
	return ""
}

type MsgUpdatePeerIDResponse struct {
}

func (m *MsgUpdatePeerIDResponse) Reset()         { *m = MsgUpdatePeerIDResponse{} }
func (m *MsgUpdatePeerIDResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdatePeerIDResponse) ProtoMessage()    {}
func (*MsgUpdatePeerIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e75709a1b06ee061, []int{3}
}
func (m *MsgUpdatePeerIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdatePeerIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdatePeerIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdatePeerIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdatePeerIDResponse.Merge(m, src)
}
func (m *MsgUpdatePeerIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdatePeerIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdatePeerIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdatePeerIDResponse proto.InternalMessageInfo

type MsgDeletePeerID struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index   string `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *MsgDeletePeerID) Reset()         { *m = MsgDeletePeerID{} }
func (m *MsgDeletePeerID) String() string { return proto.CompactTextString(m) }
func (*MsgDeletePeerID) ProtoMessage()    {}
func (*MsgDeletePeerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e75709a1b06ee061, []int{4}
}
func (m *MsgDeletePeerID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeletePeerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeletePeerID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeletePeerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeletePeerID.Merge(m, src)
}
func (m *MsgDeletePeerID) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeletePeerID) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeletePeerID.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeletePeerID proto.InternalMessageInfo

func (m *MsgDeletePeerID) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeletePeerID) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type MsgDeletePeerIDResponse struct {
}

func (m *MsgDeletePeerIDResponse) Reset()         { *m = MsgDeletePeerIDResponse{} }
func (m *MsgDeletePeerIDResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeletePeerIDResponse) ProtoMessage()    {}
func (*MsgDeletePeerIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e75709a1b06ee061, []int{5}
}
func (m *MsgDeletePeerIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeletePeerIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeletePeerIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeletePeerIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeletePeerIDResponse.Merge(m, src)
}
func (m *MsgDeletePeerIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeletePeerIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeletePeerIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeletePeerIDResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreatePeerID)(nil), "liubaninc.m0.peer.MsgCreatePeerID")
	proto.RegisterType((*MsgCreatePeerIDResponse)(nil), "liubaninc.m0.peer.MsgCreatePeerIDResponse")
	proto.RegisterType((*MsgUpdatePeerID)(nil), "liubaninc.m0.peer.MsgUpdatePeerID")
	proto.RegisterType((*MsgUpdatePeerIDResponse)(nil), "liubaninc.m0.peer.MsgUpdatePeerIDResponse")
	proto.RegisterType((*MsgDeletePeerID)(nil), "liubaninc.m0.peer.MsgDeletePeerID")
	proto.RegisterType((*MsgDeletePeerIDResponse)(nil), "liubaninc.m0.peer.MsgDeletePeerIDResponse")
}

func init() { proto.RegisterFile("m0/peer/tx.proto", fileDescriptor_e75709a1b06ee061) }

var fileDescriptor_e75709a1b06ee061 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x35, 0xd0, 0x2f,
	0x48, 0x4d, 0x2d, 0xd2, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xcc, 0xc9,
	0x2c, 0x4d, 0x4a, 0xcc, 0xcb, 0xcc, 0x4b, 0xd6, 0xcb, 0x35, 0xd0, 0x03, 0xc9, 0x29, 0xb5, 0x33,
	0x72, 0xf1, 0xfb, 0x16, 0xa7, 0x3b, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x06, 0xa4, 0xa6, 0x16, 0x79,
	0xba, 0x08, 0x49, 0x70, 0xb1, 0x27, 0x83, 0xf8, 0xf9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x5e, 0x4a, 0x6a, 0x85, 0x04, 0x13, 0x58, 0x1c,
	0xc2, 0x11, 0x92, 0xe3, 0xe2, 0x4a, 0x4e, 0x2d, 0x2a, 0xf1, 0x2c, 0x2e, 0x2e, 0x4d, 0x2d, 0x92,
	0x60, 0x06, 0x4b, 0x21, 0x89, 0x08, 0xa9, 0x70, 0xf1, 0x82, 0x78, 0xc1, 0xa9, 0x45, 0x99, 0x89,
	0x39, 0x7e, 0xa5, 0xb9, 0x12, 0x2c, 0x60, 0x25, 0xa8, 0x82, 0x4a, 0x92, 0x5c, 0xe2, 0x68, 0x0e,
	0x09, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x85, 0x39, 0x32, 0xb4, 0x20, 0x65, 0x70, 0x38,
	0x12, 0xd9, 0x21, 0x70, 0x47, 0x3a, 0x82, 0xdd, 0xe8, 0x92, 0x9a, 0x93, 0x4a, 0xae, 0x1b, 0xa1,
	0xa6, 0x23, 0x1b, 0x01, 0x33, 0xdd, 0x68, 0x29, 0x13, 0x17, 0xb3, 0x6f, 0x71, 0xba, 0x50, 0x1c,
	0x17, 0x0f, 0x4a, 0x5c, 0x29, 0xe9, 0x61, 0xc4, 0xa9, 0x1e, 0x5a, 0x30, 0x4a, 0x69, 0x11, 0x56,
	0x03, 0xb3, 0x07, 0x64, 0x3e, 0x4a, 0x30, 0xe3, 0x30, 0x1f, 0x59, 0x0d, 0x2e, 0xf3, 0xb1, 0x85,
	0x12, 0xc8, 0x7c, 0x94, 0x20, 0xc2, 0x61, 0x3e, 0xb2, 0x1a, 0x5c, 0xe6, 0x63, 0x0b, 0x27, 0x27,
	0xbb, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x52, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x9b, 0xa7, 0x9f, 0x6b, 0xa0, 0x5f, 0x01, 0xcd,
	0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0x9c, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xca, 0x7a, 0x43, 0xde, 0x3d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreatePeerID(ctx context.Context, in *MsgCreatePeerID, opts ...grpc.CallOption) (*MsgCreatePeerIDResponse, error)
	UpdatePeerID(ctx context.Context, in *MsgUpdatePeerID, opts ...grpc.CallOption) (*MsgUpdatePeerIDResponse, error)
	DeletePeerID(ctx context.Context, in *MsgDeletePeerID, opts ...grpc.CallOption) (*MsgDeletePeerIDResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePeerID(ctx context.Context, in *MsgCreatePeerID, opts ...grpc.CallOption) (*MsgCreatePeerIDResponse, error) {
	out := new(MsgCreatePeerIDResponse)
	err := c.cc.Invoke(ctx, "/liubaninc.m0.peer.Msg/CreatePeerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePeerID(ctx context.Context, in *MsgUpdatePeerID, opts ...grpc.CallOption) (*MsgUpdatePeerIDResponse, error) {
	out := new(MsgUpdatePeerIDResponse)
	err := c.cc.Invoke(ctx, "/liubaninc.m0.peer.Msg/UpdatePeerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeletePeerID(ctx context.Context, in *MsgDeletePeerID, opts ...grpc.CallOption) (*MsgDeletePeerIDResponse, error) {
	out := new(MsgDeletePeerIDResponse)
	err := c.cc.Invoke(ctx, "/liubaninc.m0.peer.Msg/DeletePeerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreatePeerID(context.Context, *MsgCreatePeerID) (*MsgCreatePeerIDResponse, error)
	UpdatePeerID(context.Context, *MsgUpdatePeerID) (*MsgUpdatePeerIDResponse, error)
	DeletePeerID(context.Context, *MsgDeletePeerID) (*MsgDeletePeerIDResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePeerID(ctx context.Context, req *MsgCreatePeerID) (*MsgCreatePeerIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePeerID not implemented")
}
func (*UnimplementedMsgServer) UpdatePeerID(ctx context.Context, req *MsgUpdatePeerID) (*MsgUpdatePeerIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePeerID not implemented")
}
func (*UnimplementedMsgServer) DeletePeerID(ctx context.Context, req *MsgDeletePeerID) (*MsgDeletePeerIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePeerID not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePeerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePeerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePeerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liubaninc.m0.peer.Msg/CreatePeerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePeerID(ctx, req.(*MsgCreatePeerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePeerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePeerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePeerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liubaninc.m0.peer.Msg/UpdatePeerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePeerID(ctx, req.(*MsgUpdatePeerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeletePeerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeletePeerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeletePeerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/liubaninc.m0.peer.Msg/DeletePeerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeletePeerID(ctx, req.(*MsgDeletePeerID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "liubaninc.m0.peer.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePeerID",
			Handler:    _Msg_CreatePeerID_Handler,
		},
		{
			MethodName: "UpdatePeerID",
			Handler:    _Msg_UpdatePeerID_Handler,
		},
		{
			MethodName: "DeletePeerID",
			Handler:    _Msg_DeletePeerID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "m0/peer/tx.proto",
}

func (m *MsgCreatePeerID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePeerID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePeerID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CertSerialNum) > 0 {
		i -= len(m.CertSerialNum)
		copy(dAtA[i:], m.CertSerialNum)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CertSerialNum)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CertIssuer) > 0 {
		i -= len(m.CertIssuer)
		copy(dAtA[i:], m.CertIssuer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CertIssuer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePeerIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePeerIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePeerIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdatePeerID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdatePeerID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdatePeerID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CertSerialNum) > 0 {
		i -= len(m.CertSerialNum)
		copy(dAtA[i:], m.CertSerialNum)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CertSerialNum)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CertIssuer) > 0 {
		i -= len(m.CertIssuer)
		copy(dAtA[i:], m.CertIssuer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CertIssuer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdatePeerIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdatePeerIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdatePeerIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeletePeerID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeletePeerID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeletePeerID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeletePeerIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeletePeerIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeletePeerIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreatePeerID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CertIssuer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CertSerialNum)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreatePeerIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdatePeerID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CertIssuer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CertSerialNum)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdatePeerIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeletePeerID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeletePeerIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreatePeerID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePeerID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePeerID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertIssuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertIssuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertSerialNum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertSerialNum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreatePeerIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePeerIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePeerIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdatePeerID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdatePeerID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdatePeerID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertIssuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertIssuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertSerialNum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertSerialNum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdatePeerIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdatePeerIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdatePeerIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeletePeerID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeletePeerID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeletePeerID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeletePeerIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeletePeerIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeletePeerIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
