// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the user's name.
type Request struct {
	Params               []byte   `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

// The response message containing the greetings
type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Result               []byte   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Err                  string   `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Response) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "rpc.Request")
	proto.RegisterType((*Response)(nil), "rpc.Response")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x4a, 0x03, 0x41,
	0x10, 0x84, 0x8d, 0xf9, 0x33, 0x4d, 0x14, 0x99, 0x83, 0x04, 0x4f, 0x31, 0x07, 0x09, 0x8a, 0x11,
	0xf4, 0x09, 0x34, 0x87, 0x44, 0x88, 0x10, 0x26, 0xbe, 0xc0, 0x38, 0xdb, 0x49, 0x96, 0x6c, 0xa6,
	0xc7, 0xee, 0x59, 0xc4, 0x37, 0xf5, 0x71, 0x64, 0x7f, 0xf0, 0xba, 0x73, 0xab, 0x82, 0xaf, 0x66,
	0xaa, 0xa0, 0x61, 0xc0, 0xde, 0xce, 0x3c, 0x53, 0x20, 0xd5, 0x66, 0x6f, 0x27, 0x37, 0xd0, 0xd7,
	0xf8, 0x95, 0xa3, 0x04, 0x75, 0x05, 0x3d, 0x6f, 0xd8, 0x1c, 0x65, 0xd4, 0x1a, 0xb7, 0xa6, 0x43,
	0x5d, 0xbb, 0xc9, 0x12, 0xce, 0x34, 0x8a, 0x27, 0x27, 0xa8, 0x14, 0x74, 0x2c, 0x25, 0x58, 0x12,
	0x5d, 0x5d, 0xea, 0x22, 0xc7, 0x28, 0x79, 0x16, 0x46, 0xa7, 0x55, 0xae, 0x72, 0xea, 0x12, 0xda,
	0xc8, 0x3c, 0x6a, 0x8f, 0x5b, 0xd3, 0x81, 0x2e, 0xe4, 0xd3, 0x6f, 0x07, 0xfa, 0x0b, 0x46, 0x0c,
	0xc8, 0xea, 0x1e, 0x60, 0x81, 0xe1, 0xc5, 0x5a, 0xca, 0x5d, 0x50, 0xc3, 0x59, 0xd1, 0xab, 0x6e,
	0x72, 0x7d, 0x5e, 0xbb, 0xea, 0xd3, 0xc9, 0x89, 0x7a, 0x84, 0x8b, 0x0d, 0xba, 0xe4, 0x1d, 0x45,
	0xcc, 0x0e, 0xb5, 0xf9, 0x6e, 0x0a, 0x54, 0xaf, 0xd7, 0x7c, 0x13, 0xfc, 0x00, 0xc3, 0x05, 0x86,
	0xd7, 0x8c, 0xec, 0x61, 0x69, 0x64, 0x1f, 0x51, 0xe6, 0x1f, 0xc7, 0x74, 0xb7, 0x0f, 0x11, 0x65,
	0x56, 0x46, 0x42, 0x1c, 0x7c, 0x07, 0x83, 0x39, 0xb9, 0x6d, 0xca, 0x47, 0x4c, 0x22, 0x57, 0xca,
	0x6e, 0x4d, 0x94, 0x45, 0xc0, 0x73, 0xe3, 0x92, 0x34, 0x31, 0x01, 0x25, 0x6e, 0xe3, 0xfc, 0xc7,
	0x66, 0xb8, 0xc9, 0x3d, 0x72, 0x63, 0xe0, 0x16, 0xba, 0x1f, 0x74, 0x40, 0x17, 0x31, 0x6f, 0x8d,
	0xc8, 0xf2, 0xe6, 0xb6, 0x14, 0xc1, 0xae, 0xc8, 0x9a, 0x2c, 0x82, 0xfd, 0xec, 0x95, 0x37, 0xfd,
	0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x70, 0x12, 0xad, 0x08, 0xe0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	GetAccount(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	SendMessageRaw(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetMessage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetBlockHash(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetBlockHeight(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	LastHeight(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Confirmed(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetMsgPool(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Candidates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetCycleSupers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Token(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	PeersInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	LocalInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetAccount(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SendMessageRaw(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/SendMessageRaw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetMessage(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetBlockHash(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/GetBlockHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetBlockHeight(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/GetBlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) LastHeight(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/LastHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Confirmed(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/Confirmed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetMsgPool(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/GetMsgPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Candidates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/Candidates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetCycleSupers(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/GetCycleSupers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Token(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) PeersInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/PeersInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) LocalInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Greeter/LocalInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	GetAccount(context.Context, *Request) (*Response, error)
	SendMessageRaw(context.Context, *Request) (*Response, error)
	GetMessage(context.Context, *Request) (*Response, error)
	GetBlockHash(context.Context, *Request) (*Response, error)
	GetBlockHeight(context.Context, *Request) (*Response, error)
	LastHeight(context.Context, *Request) (*Response, error)
	Confirmed(context.Context, *Request) (*Response, error)
	GetMsgPool(context.Context, *Request) (*Response, error)
	Candidates(context.Context, *Request) (*Response, error)
	GetCycleSupers(context.Context, *Request) (*Response, error)
	Token(context.Context, *Request) (*Response, error)
	PeersInfo(context.Context, *Request) (*Response, error)
	LocalInfo(context.Context, *Request) (*Response, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) GetAccount(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (*UnimplementedGreeterServer) SendMessageRaw(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageRaw not implemented")
}
func (*UnimplementedGreeterServer) GetMessage(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (*UnimplementedGreeterServer) GetBlockHash(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHash not implemented")
}
func (*UnimplementedGreeterServer) GetBlockHeight(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeight not implemented")
}
func (*UnimplementedGreeterServer) LastHeight(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastHeight not implemented")
}
func (*UnimplementedGreeterServer) Confirmed(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Confirmed not implemented")
}
func (*UnimplementedGreeterServer) GetMsgPool(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsgPool not implemented")
}
func (*UnimplementedGreeterServer) Candidates(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Candidates not implemented")
}
func (*UnimplementedGreeterServer) GetCycleSupers(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCycleSupers not implemented")
}
func (*UnimplementedGreeterServer) Token(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (*UnimplementedGreeterServer) PeersInfo(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeersInfo not implemented")
}
func (*UnimplementedGreeterServer) LocalInfo(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocalInfo not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetAccount(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SendMessageRaw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SendMessageRaw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/SendMessageRaw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SendMessageRaw(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetMessage(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetBlockHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetBlockHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/GetBlockHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetBlockHash(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/GetBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetBlockHeight(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_LastHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).LastHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/LastHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).LastHeight(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Confirmed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Confirmed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/Confirmed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Confirmed(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetMsgPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetMsgPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/GetMsgPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetMsgPool(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Candidates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Candidates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/Candidates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Candidates(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetCycleSupers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetCycleSupers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/GetCycleSupers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetCycleSupers(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Token(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_PeersInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).PeersInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/PeersInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).PeersInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_LocalInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).LocalInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Greeter/LocalInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).LocalInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccount",
			Handler:    _Greeter_GetAccount_Handler,
		},
		{
			MethodName: "SendMessageRaw",
			Handler:    _Greeter_SendMessageRaw_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _Greeter_GetMessage_Handler,
		},
		{
			MethodName: "GetBlockHash",
			Handler:    _Greeter_GetBlockHash_Handler,
		},
		{
			MethodName: "GetBlockHeight",
			Handler:    _Greeter_GetBlockHeight_Handler,
		},
		{
			MethodName: "LastHeight",
			Handler:    _Greeter_LastHeight_Handler,
		},
		{
			MethodName: "Confirmed",
			Handler:    _Greeter_Confirmed_Handler,
		},
		{
			MethodName: "GetMsgPool",
			Handler:    _Greeter_GetMsgPool_Handler,
		},
		{
			MethodName: "Candidates",
			Handler:    _Greeter_Candidates_Handler,
		},
		{
			MethodName: "GetCycleSupers",
			Handler:    _Greeter_GetCycleSupers_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _Greeter_Token_Handler,
		},
		{
			MethodName: "PeersInfo",
			Handler:    _Greeter_PeersInfo_Handler,
		},
		{
			MethodName: "LocalInfo",
			Handler:    _Greeter_LocalInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}