// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client/client.proto

package go_micro_client

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

type Request struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Endpoint             string   `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{0}
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

func (m *Request) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Request) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	Body                 []byte   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{1}
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

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Message struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	ContentType          string   `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d3551c163a1d198, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Message) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.client.Request")
	proto.RegisterType((*Response)(nil), "go.micro.client.Response")
	proto.RegisterType((*Message)(nil), "go.micro.client.Message")
}

func init() {
	proto.RegisterFile("client/client.proto", fileDescriptor_4d3551c163a1d198)
}

var fileDescriptor_4d3551c163a1d198 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x97, 0xad, 0xb6, 0xf3, 0x39, 0x10, 0x9e, 0x1e, 0x62, 0x0f, 0x32, 0x7b, 0xea, 0x29,
	0x8a, 0x9e, 0xc5, 0x43, 0xcf, 0x82, 0x54, 0xf1, 0x2a, 0x6d, 0xf6, 0x98, 0x81, 0x2e, 0x89, 0x4d,
	0x36, 0xe8, 0x1f, 0xe9, 0xff, 0x24, 0xa4, 0xd5, 0x89, 0x6e, 0x97, 0x9d, 0x92, 0xef, 0xfd, 0xe0,
	0x7b, 0x5f, 0xbe, 0xc0, 0x99, 0x6c, 0x14, 0x69, 0x7f, 0xdd, 0x1f, 0xc2, 0xb6, 0xc6, 0x1b, 0x3c,
	0x5d, 0x1a, 0xb1, 0x52, 0xb2, 0x35, 0xa2, 0x1f, 0x67, 0x1b, 0x48, 0x4a, 0xfa, 0x58, 0x93, 0xf3,
	0xc8, 0x21, 0x71, 0xd4, 0x6e, 0x94, 0x24, 0xce, 0xe6, 0x2c, 0x3f, 0x2e, 0xbf, 0x25, 0xa6, 0x30,
	0x25, 0xbd, 0xb0, 0x46, 0x69, 0xcf, 0xc7, 0x01, 0xfd, 0x68, 0xbc, 0x82, 0x99, 0x34, 0xda, 0x93,
	0xf6, 0x6f, 0xbe, 0xb3, 0xc4, 0x27, 0x81, 0x9f, 0x0c, 0xb3, 0x97, 0xce, 0x12, 0x22, 0x44, 0xb5,
	0x59, 0x74, 0x3c, 0x9a, 0xb3, 0x7c, 0x56, 0x86, 0x7b, 0x76, 0x09, 0xd3, 0x92, 0x9c, 0x35, 0xda,
	0x6d, 0x39, 0xfb, 0xc5, 0x5f, 0x21, 0x79, 0x24, 0xe7, 0xaa, 0x25, 0xe1, 0x39, 0x1c, 0x79, 0x63,
	0x95, 0x1c, 0x52, 0xf5, 0xe2, 0xdf, 0xde, 0xf1, 0xfe, 0xbd, 0x93, 0xad, 0xef, 0xed, 0x27, 0x83,
	0xb8, 0x08, 0x4f, 0xc7, 0x7b, 0x88, 0x8a, 0xaa, 0x69, 0x90, 0x8b, 0x3f, 0xa5, 0x88, 0xa1, 0x91,
	0xf4, 0x62, 0x07, 0xe9, 0x33, 0x67, 0x23, 0x2c, 0x20, 0x7e, 0xf6, 0x2d, 0x55, 0xab, 0x03, 0x0d,
	0x72, 0x76, 0xc3, 0xf0, 0x01, 0x92, 0xa7, 0x75, 0xdd, 0x28, 0xf7, 0xbe, 0xc3, 0x65, 0x28, 0x20,
	0xdd, 0x4b, 0xb2, 0x51, 0x1d, 0x87, 0x7f, 0xbd, 0xfb, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x12, 0xce,
	0x36, 0x9d, 0xee, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClientClient is the client API for Client service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientClient interface {
	// Call allows a single request to be made
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// Stream is a bidirectional stream
	Stream(ctx context.Context, opts ...grpc.CallOption) (Client_StreamClient, error)
	// Publish publishes a message and returns an empty Message
	Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type clientClient struct {
	cc grpc.ClientConnInterface
}

func NewClientClient(cc grpc.ClientConnInterface) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.client.Client/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Client_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Client_serviceDesc.Streams[0], "/go.micro.client.Client/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStreamClient{stream}
	return x, nil
}

type Client_StreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type clientStreamClient struct {
	grpc.ClientStream
}

func (x *clientStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientClient) Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/go.micro.client.Client/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServer is the server API for Client service.
type ClientServer interface {
	// Call allows a single request to be made
	Call(context.Context, *Request) (*Response, error)
	// Stream is a bidirectional stream
	Stream(Client_StreamServer) error
	// Publish publishes a message and returns an empty Message
	Publish(context.Context, *Message) (*Message, error)
}

// UnimplementedClientServer can be embedded to have forward compatible implementations.
type UnimplementedClientServer struct {
}

func (*UnimplementedClientServer) Call(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedClientServer) Stream(srv Client_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (*UnimplementedClientServer) Publish(ctx context.Context, req *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}

func RegisterClientServer(s *grpc.Server, srv ClientServer) {
	s.RegisterService(&_Client_serviceDesc, srv)
}

func _Client_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.client.Client/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientServer).Stream(&clientStreamServer{stream})
}

type Client_StreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type clientStreamServer struct {
	grpc.ServerStream
}

func (x *clientStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Client_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.client.Client/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Publish(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Client_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.client.Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Client_Call_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Client_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Client_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "client/client.proto",
}
