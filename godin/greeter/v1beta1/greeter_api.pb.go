// Code generated by protoc-gen-go. DO NOT EDIT.
// source: godin/greeter/v1beta1/greeter_api.proto

package greeterv1beta1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64a2bf718c069116, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64a2bf718c069116, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "godin.greeter.v1beta1.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "godin.greeter.v1beta1.HelloResponse")
}

func init() {
	proto.RegisterFile("godin/greeter/v1beta1/greeter_api.proto", fileDescriptor_64a2bf718c069116)
}

var fileDescriptor_64a2bf718c069116 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcf, 0x4f, 0xc9,
	0xcc, 0xd3, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0x84, 0xf1, 0xe3, 0x13, 0x0b, 0x32, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0xc1,
	0x0a, 0xf5, 0xa0, 0x12, 0x7a, 0x50, 0x85, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e, 0xa9, 0x39, 0x39, 0xf9,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x36, 0x17, 0x2f, 0x54, 0x4d, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x14, 0x17, 0x07, 0xd8, 0x9c, 0xcc, 0xbc, 0x74, 0xa8, 0x42, 0x38,
	0xdf, 0x28, 0x81, 0x8b, 0xcb, 0x1d, 0x62, 0x87, 0x63, 0x80, 0xa7, 0x50, 0x10, 0x17, 0x2b, 0x58,
	0xab, 0x90, 0xb2, 0x1e, 0x56, 0xfb, 0xf5, 0x90, 0x2d, 0x97, 0x52, 0xc1, 0xaf, 0x08, 0x62, 0xbb,
	0x53, 0x09, 0x97, 0x64, 0x72, 0x7e, 0x2e, 0x76, 0xa5, 0x4e, 0xfc, 0x30, 0xcb, 0x0b, 0x32, 0x03,
	0x40, 0xfe, 0x0e, 0x60, 0x8c, 0xe2, 0x83, 0xaa, 0x81, 0x2a, 0x59, 0xc4, 0xc4, 0xec, 0xee, 0x1e,
	0xb1, 0x8a, 0x49, 0xd4, 0x1d, 0x6c, 0x00, 0x54, 0xbd, 0x5e, 0x98, 0xa1, 0x13, 0x48, 0xf6, 0x14,
	0x54, 0x3c, 0x06, 0x2a, 0x1e, 0x03, 0x15, 0x4f, 0x62, 0x03, 0x07, 0xa3, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xd9, 0x69, 0x4d, 0x2c, 0x71, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterAPIClient is the client API for GreeterAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterAPIClient interface {
	// Say hello.
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type greeterAPIClient struct {
	cc *grpc.ClientConn
}

func NewGreeterAPIClient(cc *grpc.ClientConn) GreeterAPIClient {
	return &greeterAPIClient{cc}
}

func (c *greeterAPIClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/godin.greeter.v1beta1.GreeterAPI/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterAPIServer is the server API for GreeterAPI service.
type GreeterAPIServer interface {
	// Say hello.
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterGreeterAPIServer(s *grpc.Server, srv GreeterAPIServer) {
	s.RegisterService(&_GreeterAPI_serviceDesc, srv)
}

func _GreeterAPI_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterAPIServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/godin.greeter.v1beta1.GreeterAPI/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterAPIServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreeterAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "godin.greeter.v1beta1.GreeterAPI",
	HandlerType: (*GreeterAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreeterAPI_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "godin/greeter/v1beta1/greeter_api.proto",
}