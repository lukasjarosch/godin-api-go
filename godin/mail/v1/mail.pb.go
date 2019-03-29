// Code generated by protoc-gen-go. DO NOT EDIT.
// source: godin/mail/v1/mail.proto

package mailv1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Mail struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Mail) Reset()         { *m = Mail{} }
func (m *Mail) String() string { return proto.CompactTextString(m) }
func (*Mail) ProtoMessage()    {}
func (*Mail) Descriptor() ([]byte, []int) {
	return fileDescriptor_715c08cf773e2f00, []int{0}
}

func (m *Mail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mail.Unmarshal(m, b)
}
func (m *Mail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mail.Marshal(b, m, deterministic)
}
func (m *Mail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mail.Merge(m, src)
}
func (m *Mail) XXX_Size() int {
	return xxx_messageInfo_Mail.Size(m)
}
func (m *Mail) XXX_DiscardUnknown() {
	xxx_messageInfo_Mail.DiscardUnknown(m)
}

var xxx_messageInfo_Mail proto.InternalMessageInfo

func (m *Mail) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Mail) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Mail)(nil), "godin.mail.v1.Mail")
}

func init() { proto.RegisterFile("godin/mail/v1/mail.proto", fileDescriptor_715c08cf773e2f00) }

var fileDescriptor_715c08cf773e2f00 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xcf, 0x4f, 0xc9,
	0xcc, 0xd3, 0xcf, 0x4d, 0xcc, 0xcc, 0xd1, 0x2f, 0x33, 0x04, 0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xbc, 0x60, 0x19, 0x3d, 0xb0, 0x48, 0x99, 0xa1, 0x92, 0x01, 0x17, 0x8b, 0x6f, 0x62,
	0x66, 0x8e, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x84, 0x23, 0x24, 0xc4, 0xc5, 0x92, 0x94, 0x9f, 0x52, 0x29, 0xc1, 0x04, 0x16, 0x04,
	0xb3, 0x9d, 0xfc, 0xb8, 0x04, 0x93, 0xf3, 0x73, 0xf5, 0x50, 0x8c, 0x71, 0xe2, 0x04, 0x19, 0x12,
	0x00, 0xb2, 0x20, 0x80, 0x31, 0x8a, 0x0d, 0x24, 0x5a, 0x66, 0xb8, 0x88, 0x89, 0xd9, 0xdd, 0x37,
	0x62, 0x15, 0x13, 0xaf, 0x3b, 0x58, 0x29, 0x48, 0x85, 0x5e, 0x98, 0xe1, 0x29, 0x28, 0x3f, 0x06,
	0xc4, 0x8f, 0x09, 0x33, 0x4c, 0x62, 0x03, 0xbb, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x14,
	0x9d, 0xde, 0xdb, 0xb3, 0x00, 0x00, 0x00,
}
