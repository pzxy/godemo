// Code generated by protoc-gen-go. DO NOT EDIT.
// source: leetcode.proto

package proto

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

type CallRequest struct {
	CallString           string   `protobuf:"bytes,1,opt,name=CallString,proto3" json:"CallString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallRequest) Reset()         { *m = CallRequest{} }
func (m *CallRequest) String() string { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()    {}
func (*CallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{0}
}

func (m *CallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallRequest.Unmarshal(m, b)
}
func (m *CallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallRequest.Marshal(b, m, deterministic)
}
func (m *CallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallRequest.Merge(m, src)
}
func (m *CallRequest) XXX_Size() int {
	return xxx_messageInfo_CallRequest.Size(m)
}
func (m *CallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallRequest proto.InternalMessageInfo

func (m *CallRequest) GetCallString() string {
	if m != nil {
		return m.CallString
	}
	return ""
}

type CallResponse struct {
	CallString           string   `protobuf:"bytes,1,opt,name=CallString,proto3" json:"CallString,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallResponse) Reset()         { *m = CallResponse{} }
func (m *CallResponse) String() string { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()    {}
func (*CallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{1}
}

func (m *CallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallResponse.Unmarshal(m, b)
}
func (m *CallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallResponse.Marshal(b, m, deterministic)
}
func (m *CallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallResponse.Merge(m, src)
}
func (m *CallResponse) XXX_Size() int {
	return xxx_messageInfo_CallResponse.Size(m)
}
func (m *CallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallResponse proto.InternalMessageInfo

func (m *CallResponse) GetCallString() string {
	if m != nil {
		return m.CallString
	}
	return ""
}

type StreamPoint struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamPoint) Reset()         { *m = StreamPoint{} }
func (m *StreamPoint) String() string { return proto.CompactTextString(m) }
func (*StreamPoint) ProtoMessage()    {}
func (*StreamPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{2}
}

func (m *StreamPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamPoint.Unmarshal(m, b)
}
func (m *StreamPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamPoint.Marshal(b, m, deterministic)
}
func (m *StreamPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamPoint.Merge(m, src)
}
func (m *StreamPoint) XXX_Size() int {
	return xxx_messageInfo_StreamPoint.Size(m)
}
func (m *StreamPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamPoint.DiscardUnknown(m)
}

var xxx_messageInfo_StreamPoint proto.InternalMessageInfo

func (m *StreamPoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamPoint) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StreamRequest struct {
	Pt                   *StreamPoint `protobuf:"bytes,1,opt,name=pt,proto3" json:"pt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{3}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetPt() *StreamPoint {
	if m != nil {
		return m.Pt
	}
	return nil
}

type StreamResponse struct {
	Pt                   *StreamPoint `protobuf:"bytes,1,opt,name=pt,proto3" json:"pt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca53982754088a9d, []int{4}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetPt() *StreamPoint {
	if m != nil {
		return m.Pt
	}
	return nil
}

func init() {
	proto.RegisterType((*CallRequest)(nil), "proto.CallRequest")
	proto.RegisterType((*CallResponse)(nil), "proto.CallResponse")
	proto.RegisterType((*StreamPoint)(nil), "proto.StreamPoint")
	proto.RegisterType((*StreamRequest)(nil), "proto.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "proto.StreamResponse")
}

func init() { proto.RegisterFile("leetcode.proto", fileDescriptor_ca53982754088a9d) }

var fileDescriptor_ca53982754088a9d = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0xd0, 0x14, 0x9d, 0xa8, 0x87, 0xb1, 0x42, 0xe9, 0x41, 0xca, 0x9e, 0x72, 0x31,
	0x94, 0x56, 0x29, 0x7a, 0xd5, 0xa3, 0x07, 0xd9, 0xfc, 0x82, 0xd8, 0x0e, 0x12, 0x48, 0x32, 0x71,
	0x33, 0xe9, 0x4f, 0xf7, 0x2c, 0xcd, 0x66, 0x21, 0xf5, 0xa2, 0x39, 0xed, 0xce, 0xe3, 0x7d, 0x3b,
	0x6f, 0x76, 0x00, 0xf6, 0x54, 0x72, 0x52, 0x5b, 0x16, 0xc6, 0xb0, 0x3b, 0xf4, 0x3d, 0x44, 0x2f,
	0x59, 0x51, 0x18, 0xfa, 0x6a, 0xa9, 0x11, 0xbc, 0x03, 0x38, 0x96, 0xa9, 0xd8, 0xbc, 0xfa, 0x9c,
	0xab, 0xa5, 0x8a, 0x2f, 0xcc, 0x40, 0xd1, 0x09, 0x5c, 0x3a, 0x7b, 0x53, 0x73, 0xd5, 0xd0, 0x9f,
	0xfe, 0x2d, 0x44, 0xa9, 0x58, 0xca, 0xca, 0x77, 0xce, 0x2b, 0x41, 0x84, 0x49, 0x95, 0x95, 0xd4,
	0x1b, 0xbb, 0x3b, 0xce, 0x20, 0x3c, 0x64, 0x45, 0x4b, 0xf3, 0x60, 0xa9, 0xe2, 0xd0, 0xb8, 0x42,
	0x6f, 0xe0, 0xca, 0x81, 0x3e, 0x99, 0x86, 0xa0, 0x96, 0x0e, 0x8c, 0xd6, 0xe8, 0x66, 0x48, 0x06,
	0x4f, 0x9b, 0xa0, 0x16, 0xfd, 0x00, 0xd7, 0x1e, 0xea, 0xf3, 0xfd, 0x83, 0x5a, 0x7f, 0x2b, 0xdf,
	0x2b, 0x25, 0x7b, 0xc8, 0x77, 0x84, 0x5b, 0x98, 0xbc, 0xe5, 0x8d, 0xe0, 0xec, 0x84, 0xe8, 0x93,
	0x2c, 0x6e, 0x7f, 0xa9, 0xae, 0x95, 0x3e, 0x5b, 0x29, 0x7c, 0x82, 0xa9, 0xa1, 0x1d, 0xdb, 0xfd,
	0x48, 0x34, 0x56, 0xf8, 0x0c, 0xa1, 0xe1, 0x56, 0x68, 0x34, 0xb9, 0x52, 0xf8, 0x08, 0xe7, 0xc7,
	0x3f, 0x7f, 0xa5, 0x92, 0xd1, 0x4f, 0x39, 0xd8, 0xea, 0xe2, 0xe6, 0x44, 0xf3, 0xe8, 0xc7, 0xb4,
	0x53, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0x79, 0x8e, 0xc9, 0x17, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	//服务器流
	List(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamService_ListClient, error)
	//客户端流
	Record(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordClient, error)
	//双向流
	Route(ctx context.Context, opts ...grpc.CallOption) (StreamService_RouteClient, error)
	//函数调用
	CallDemo(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) List(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/proto.StreamService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceListClient struct {
	grpc.ClientStream
}

func (x *streamServiceListClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Record(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[1], "/proto.StreamService/Record", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRecordClient{stream}
	return x, nil
}

type StreamService_RecordClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceRecordClient struct {
	grpc.ClientStream
}

func (x *streamServiceRecordClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRecordClient) CloseAndRecv() (*StreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Route(ctx context.Context, opts ...grpc.CallOption) (StreamService_RouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[2], "/proto.StreamService/Route", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRouteClient{stream}
	return x, nil
}

type StreamService_RouteClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceRouteClient struct {
	grpc.ClientStream
}

func (x *streamServiceRouteClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRouteClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) CallDemo(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/proto.StreamService/CallDemo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	//服务器流
	List(*StreamRequest, StreamService_ListServer) error
	//客户端流
	Record(StreamService_RecordServer) error
	//双向流
	Route(StreamService_RouteServer) error
	//函数调用
	CallDemo(context.Context, *CallRequest) (*CallResponse, error)
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) List(req *StreamRequest, srv StreamService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedStreamServiceServer) Record(srv StreamService_RecordServer) error {
	return status.Errorf(codes.Unimplemented, "method Record not implemented")
}
func (*UnimplementedStreamServiceServer) Route(srv StreamService_RouteServer) error {
	return status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (*UnimplementedStreamServiceServer) CallDemo(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallDemo not implemented")
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).List(m, &streamServiceListServer{stream})
}

type StreamService_ListServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type streamServiceListServer struct {
	grpc.ServerStream
}

func (x *streamServiceListServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_Record_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Record(&streamServiceRecordServer{stream})
}

type StreamService_RecordServer interface {
	SendAndClose(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceRecordServer struct {
	grpc.ServerStream
}

func (x *streamServiceRecordServer) SendAndClose(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRecordServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_Route_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Route(&streamServiceRouteServer{stream})
}

type StreamService_RouteServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceRouteServer struct {
	grpc.ServerStream
}

func (x *streamServiceRouteServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRouteServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_CallDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).CallDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StreamService/CallDemo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).CallDemo(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallDemo",
			Handler:    _StreamService_CallDemo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _StreamService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Record",
			Handler:       _StreamService_Record_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Route",
			Handler:       _StreamService_Route_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "leetcode.proto",
}
