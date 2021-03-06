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

type BurpRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BurpRequest) Reset()         { *m = BurpRequest{} }
func (m *BurpRequest) String() string { return proto.CompactTextString(m) }
func (*BurpRequest) ProtoMessage()    {}
func (*BurpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64a2bf718c069116, []int{2}
}

func (m *BurpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BurpRequest.Unmarshal(m, b)
}
func (m *BurpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BurpRequest.Marshal(b, m, deterministic)
}
func (m *BurpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurpRequest.Merge(m, src)
}
func (m *BurpRequest) XXX_Size() int {
	return xxx_messageInfo_BurpRequest.Size(m)
}
func (m *BurpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BurpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BurpRequest proto.InternalMessageInfo

type BurpResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BurpResponse) Reset()         { *m = BurpResponse{} }
func (m *BurpResponse) String() string { return proto.CompactTextString(m) }
func (*BurpResponse) ProtoMessage()    {}
func (*BurpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64a2bf718c069116, []int{3}
}

func (m *BurpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BurpResponse.Unmarshal(m, b)
}
func (m *BurpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BurpResponse.Marshal(b, m, deterministic)
}
func (m *BurpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurpResponse.Merge(m, src)
}
func (m *BurpResponse) XXX_Size() int {
	return xxx_messageInfo_BurpResponse.Size(m)
}
func (m *BurpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BurpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BurpResponse proto.InternalMessageInfo

type GoodbyeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodbyeRequest) Reset()         { *m = GoodbyeRequest{} }
func (m *GoodbyeRequest) String() string { return proto.CompactTextString(m) }
func (*GoodbyeRequest) ProtoMessage()    {}
func (*GoodbyeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64a2bf718c069116, []int{4}
}

func (m *GoodbyeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodbyeRequest.Unmarshal(m, b)
}
func (m *GoodbyeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodbyeRequest.Marshal(b, m, deterministic)
}
func (m *GoodbyeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodbyeRequest.Merge(m, src)
}
func (m *GoodbyeRequest) XXX_Size() int {
	return xxx_messageInfo_GoodbyeRequest.Size(m)
}
func (m *GoodbyeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodbyeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GoodbyeRequest proto.InternalMessageInfo

type GoodbyeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoodbyeResponse) Reset()         { *m = GoodbyeResponse{} }
func (m *GoodbyeResponse) String() string { return proto.CompactTextString(m) }
func (*GoodbyeResponse) ProtoMessage()    {}
func (*GoodbyeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64a2bf718c069116, []int{5}
}

func (m *GoodbyeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoodbyeResponse.Unmarshal(m, b)
}
func (m *GoodbyeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoodbyeResponse.Marshal(b, m, deterministic)
}
func (m *GoodbyeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoodbyeResponse.Merge(m, src)
}
func (m *GoodbyeResponse) XXX_Size() int {
	return xxx_messageInfo_GoodbyeResponse.Size(m)
}
func (m *GoodbyeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GoodbyeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GoodbyeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HelloRequest)(nil), "godin.greeter.v1beta1.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "godin.greeter.v1beta1.HelloResponse")
	proto.RegisterType((*BurpRequest)(nil), "godin.greeter.v1beta1.BurpRequest")
	proto.RegisterType((*BurpResponse)(nil), "godin.greeter.v1beta1.BurpResponse")
	proto.RegisterType((*GoodbyeRequest)(nil), "godin.greeter.v1beta1.GoodbyeRequest")
	proto.RegisterType((*GoodbyeResponse)(nil), "godin.greeter.v1beta1.GoodbyeResponse")
}

func init() {
	proto.RegisterFile("godin/greeter/v1beta1/greeter_api.proto", fileDescriptor_64a2bf718c069116)
}

var fileDescriptor_64a2bf718c069116 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcf, 0x4f, 0xc9,
	0xcc, 0xd3, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0xd2, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0x84, 0xf1, 0xe3, 0x13, 0x0b, 0x32, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0xc1,
	0x0a, 0xf5, 0xa0, 0x12, 0x7a, 0x50, 0x85, 0x4a, 0x4a, 0x5c, 0x3c, 0x1e, 0xa9, 0x39, 0x39, 0xf9,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x36, 0x17, 0x2f, 0x54, 0x4d, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x14, 0x17, 0x07, 0xd8, 0x9c, 0xcc, 0xbc, 0x74, 0xa8, 0x42, 0x38,
	0x5f, 0x89, 0x97, 0x8b, 0xdb, 0xa9, 0xb4, 0xa8, 0x00, 0x6a, 0x9e, 0x12, 0x1f, 0x17, 0x0f, 0x84,
	0x0b, 0xd1, 0xaa, 0x24, 0xc0, 0xc5, 0xe7, 0x9e, 0x9f, 0x9f, 0x92, 0x54, 0x99, 0x0a, 0x53, 0x21,
	0xc8, 0xc5, 0x0f, 0x17, 0x81, 0x28, 0x32, 0xea, 0x66, 0xe2, 0xe2, 0x72, 0x87, 0x38, 0xd4, 0x31,
	0xc0, 0x53, 0x28, 0x88, 0x8b, 0x15, 0x6c, 0xbf, 0x90, 0xb2, 0x1e, 0x56, 0x4f, 0xe8, 0x21, 0xfb,
	0x40, 0x4a, 0x05, 0xbf, 0x22, 0xa8, 0x17, 0xfc, 0xb9, 0x58, 0x40, 0xee, 0x12, 0x52, 0xc2, 0xa1,
	0x1a, 0xc9, 0x0f, 0x52, 0xca, 0x78, 0xd5, 0x40, 0x0d, 0x8c, 0xe0, 0x62, 0x87, 0x7a, 0x43, 0x48,
	0x15, 0x87, 0x7a, 0x54, 0x8f, 0x4b, 0xa9, 0x11, 0x52, 0x06, 0x31, 0xd9, 0xa9, 0x84, 0x4b, 0x32,
	0x39, 0x3f, 0x17, 0xbb, 0x62, 0x27, 0x7e, 0x58, 0x38, 0x15, 0x64, 0x06, 0x80, 0xe2, 0x39, 0x80,
	0x31, 0x8a, 0x0f, 0xaa, 0x06, 0xaa, 0x64, 0x11, 0x13, 0xb3, 0xbb, 0x7b, 0xc4, 0x2a, 0x26, 0x51,
	0x77, 0xb0, 0x01, 0x50, 0xf5, 0x7a, 0x61, 0x86, 0x4e, 0x20, 0xd9, 0x53, 0x50, 0xf1, 0x18, 0xa8,
	0x78, 0x0c, 0x54, 0x3c, 0x89, 0x0d, 0x9c, 0x6c, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x12,
	0xff, 0xf0, 0x61, 0x61, 0x02, 0x00, 0x00,
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
	Burp(ctx context.Context, in *BurpRequest, opts ...grpc.CallOption) (*BurpResponse, error)
	Goodbye(ctx context.Context, in *GoodbyeRequest, opts ...grpc.CallOption) (*GoodbyeResponse, error)
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

func (c *greeterAPIClient) Burp(ctx context.Context, in *BurpRequest, opts ...grpc.CallOption) (*BurpResponse, error) {
	out := new(BurpResponse)
	err := c.cc.Invoke(ctx, "/godin.greeter.v1beta1.GreeterAPI/Burp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterAPIClient) Goodbye(ctx context.Context, in *GoodbyeRequest, opts ...grpc.CallOption) (*GoodbyeResponse, error) {
	out := new(GoodbyeResponse)
	err := c.cc.Invoke(ctx, "/godin.greeter.v1beta1.GreeterAPI/Goodbye", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterAPIServer is the server API for GreeterAPI service.
type GreeterAPIServer interface {
	// Say hello.
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	Burp(context.Context, *BurpRequest) (*BurpResponse, error)
	Goodbye(context.Context, *GoodbyeRequest) (*GoodbyeResponse, error)
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

func _GreeterAPI_Burp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BurpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterAPIServer).Burp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/godin.greeter.v1beta1.GreeterAPI/Burp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterAPIServer).Burp(ctx, req.(*BurpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterAPI_Goodbye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodbyeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterAPIServer).Goodbye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/godin.greeter.v1beta1.GreeterAPI/Goodbye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterAPIServer).Goodbye(ctx, req.(*GoodbyeRequest))
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
		{
			MethodName: "Burp",
			Handler:    _GreeterAPI_Burp_Handler,
		},
		{
			MethodName: "Goodbye",
			Handler:    _GreeterAPI_Goodbye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "godin/greeter/v1beta1/greeter_api.proto",
}
