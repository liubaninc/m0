// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/certificates.proto

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

type Certificates struct {
	Creator    string                  `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Identifier CertificatesIdentifier  `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier"`
	Items      []CertificateIdentifier `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Disable    bool                    `protobuf:"varint,12,opt,name=disable,proto3" json:"disable,omitempty"`
}

func (m *Certificates) Reset()         { *m = Certificates{} }
func (m *Certificates) String() string { return proto.CompactTextString(m) }
func (*Certificates) ProtoMessage()    {}
func (*Certificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5962fe4a45c1649, []int{0}
}
func (m *Certificates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Certificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Certificates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Certificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificates.Merge(m, src)
}
func (m *Certificates) XXX_Size() int {
	return m.Size()
}
func (m *Certificates) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificates.DiscardUnknown(m)
}

var xxx_messageInfo_Certificates proto.InternalMessageInfo

func (m *Certificates) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Certificates) GetIdentifier() CertificatesIdentifier {
	if m != nil {
		return m.Identifier
	}
	return CertificatesIdentifier{}
}

func (m *Certificates) GetItems() []CertificateIdentifier {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Certificates) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

type ChildCertificates struct {
	Identifier CertificatesIdentifier   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier"`
	Items      []CertificatesIdentifier `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
}

func (m *ChildCertificates) Reset()         { *m = ChildCertificates{} }
func (m *ChildCertificates) String() string { return proto.CompactTextString(m) }
func (*ChildCertificates) ProtoMessage()    {}
func (*ChildCertificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5962fe4a45c1649, []int{1}
}
func (m *ChildCertificates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChildCertificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChildCertificates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChildCertificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChildCertificates.Merge(m, src)
}
func (m *ChildCertificates) XXX_Size() int {
	return m.Size()
}
func (m *ChildCertificates) XXX_DiscardUnknown() {
	xxx_messageInfo_ChildCertificates.DiscardUnknown(m)
}

var xxx_messageInfo_ChildCertificates proto.InternalMessageInfo

func (m *ChildCertificates) GetIdentifier() CertificatesIdentifier {
	if m != nil {
		return m.Identifier
	}
	return CertificatesIdentifier{}
}

func (m *ChildCertificates) GetItems() []CertificatesIdentifier {
	if m != nil {
		return m.Items
	}
	return nil
}

type CertificatesIdentifier struct {
	Subject      string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyID string `protobuf:"bytes,2,opt,name=subjectKeyID,proto3" json:"subjectKeyID,omitempty"`
}

func (m *CertificatesIdentifier) Reset()         { *m = CertificatesIdentifier{} }
func (m *CertificatesIdentifier) String() string { return proto.CompactTextString(m) }
func (*CertificatesIdentifier) ProtoMessage()    {}
func (*CertificatesIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5962fe4a45c1649, []int{2}
}
func (m *CertificatesIdentifier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificatesIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificatesIdentifier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CertificatesIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificatesIdentifier.Merge(m, src)
}
func (m *CertificatesIdentifier) XXX_Size() int {
	return m.Size()
}
func (m *CertificatesIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificatesIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_CertificatesIdentifier proto.InternalMessageInfo

func (m *CertificatesIdentifier) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *CertificatesIdentifier) GetSubjectKeyID() string {
	if m != nil {
		return m.SubjectKeyID
	}
	return ""
}

func init() {
	proto.RegisterType((*Certificates)(nil), "liubaninc.m0.pki.Certificates")
	proto.RegisterType((*ChildCertificates)(nil), "liubaninc.m0.pki.ChildCertificates")
	proto.RegisterType((*CertificatesIdentifier)(nil), "liubaninc.m0.pki.CertificatesIdentifier")
}

func init() { proto.RegisterFile("pki/certificates.proto", fileDescriptor_c5962fe4a45c1649) }

var fileDescriptor_c5962fe4a45c1649 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xc8, 0xce, 0xd4,
	0x4f, 0x4e, 0x2d, 0x2a, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0xc8, 0xc9, 0x2c, 0x4d, 0x4a, 0xcc, 0xcb, 0xcc, 0x4b, 0xd6, 0xcb, 0x35,
	0xd0, 0x2b, 0xc8, 0xce, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xea, 0x83, 0x58, 0x10,
	0x75, 0x52, 0xa2, 0x68, 0xfa, 0x21, 0xc2, 0x4a, 0xd7, 0x19, 0xb9, 0x78, 0x9c, 0x91, 0x4c, 0x15,
	0x92, 0xe0, 0x62, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x82, 0x71, 0x85, 0xfc, 0xb8, 0xb8, 0x32, 0x53, 0x52, 0xf3, 0x40, 0x4a, 0x53, 0x8b, 0x24,
	0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x34, 0xf4, 0xd0, 0xad, 0xd7, 0x43, 0x36, 0xcd, 0x13, 0xae,
	0xde, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x24, 0x13, 0x84, 0x9c, 0xb9, 0x58, 0x33, 0x4b,
	0x52, 0x73, 0x8b, 0x25, 0x98, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0xd4, 0xf1, 0x1a, 0x85, 0x61, 0x12,
	0x44, 0x2f, 0xc8, 0xb9, 0x29, 0x99, 0xc5, 0x89, 0x49, 0x39, 0xa9, 0x12, 0x3c, 0x0a, 0x8c, 0x1a,
	0x1c, 0x41, 0x30, 0xae, 0xd2, 0x4a, 0x46, 0x2e, 0x41, 0xe7, 0x8c, 0xcc, 0x9c, 0x14, 0x14, 0xef,
	0x51, 0xdb, 0x13, 0x2e, 0xa8, 0x9e, 0x20, 0xd5, 0x28, 0x88, 0x66, 0xa5, 0x30, 0x2e, 0x31, 0xec,
	0xca, 0x40, 0xfe, 0x2b, 0x2e, 0x4d, 0xca, 0x4a, 0x4d, 0x2e, 0x81, 0x45, 0x07, 0x94, 0x2b, 0xa4,
	0xc4, 0xc5, 0x03, 0x65, 0x7a, 0xa7, 0x56, 0x7a, 0xba, 0x80, 0xfd, 0xc2, 0x19, 0x84, 0x22, 0xe6,
	0x64, 0x7b, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78,
	0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xca, 0xe9, 0x99, 0x25,
	0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x70, 0x27, 0xeb, 0xe7, 0x1a, 0xe8, 0x57, 0xe8,
	0x83, 0x12, 0x4a, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x8d, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xeb, 0x0d, 0x76, 0x0a, 0x7c, 0x02, 0x00, 0x00,
}

func (m *Certificates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Certificates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Certificates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Disable {
		i--
		if m.Disable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x60
	}
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCertificates(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Identifier.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCertificates(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintCertificates(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChildCertificates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChildCertificates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChildCertificates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCertificates(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Identifier.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCertificates(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func (m *CertificatesIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificatesIdentifier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CertificatesIdentifier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubjectKeyID) > 0 {
		i -= len(m.SubjectKeyID)
		copy(dAtA[i:], m.SubjectKeyID)
		i = encodeVarintCertificates(dAtA, i, uint64(len(m.SubjectKeyID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintCertificates(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCertificates(dAtA []byte, offset int, v uint64) int {
	offset -= sovCertificates(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Certificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovCertificates(uint64(l))
	}
	l = m.Identifier.Size()
	n += 1 + l + sovCertificates(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovCertificates(uint64(l))
		}
	}
	if m.Disable {
		n += 2
	}
	return n
}

func (m *ChildCertificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Identifier.Size()
	n += 1 + l + sovCertificates(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovCertificates(uint64(l))
		}
	}
	return n
}

func (m *CertificatesIdentifier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovCertificates(uint64(l))
	}
	l = len(m.SubjectKeyID)
	if l > 0 {
		n += 1 + l + sovCertificates(uint64(l))
	}
	return n
}

func sovCertificates(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCertificates(x uint64) (n int) {
	return sovCertificates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Certificates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertificates
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
			return fmt.Errorf("proto: Certificates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Certificates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificates
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
				return ErrInvalidLengthCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificates
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
				return ErrInvalidLengthCertificates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Identifier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificates
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
				return ErrInvalidLengthCertificates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, CertificateIdentifier{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificates
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
			m.Disable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCertificates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertificates
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
func (m *ChildCertificates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertificates
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
			return fmt.Errorf("proto: ChildCertificates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChildCertificates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificates
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
				return ErrInvalidLengthCertificates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Identifier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificates
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
				return ErrInvalidLengthCertificates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, CertificatesIdentifier{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCertificates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertificates
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
func (m *CertificatesIdentifier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertificates
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
			return fmt.Errorf("proto: CertificatesIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificatesIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificates
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
				return ErrInvalidLengthCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificates
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
				return ErrInvalidLengthCertificates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCertificates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertificates
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
func skipCertificates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCertificates
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
					return 0, ErrIntOverflowCertificates
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
					return 0, ErrIntOverflowCertificates
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
				return 0, ErrInvalidLengthCertificates
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCertificates
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCertificates
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCertificates        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCertificates          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCertificates = fmt.Errorf("proto: unexpected end of group")
)
