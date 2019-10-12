// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package proto

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
	return fileDescriptor_bb17ef3f514bfe54, []int{0}
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

type StreamReq struct {
	Pt                   *StreamPoint `protobuf:"bytes,1,opt,name=pt,proto3" json:"pt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StreamReq) Reset()         { *m = StreamReq{} }
func (m *StreamReq) String() string { return proto.CompactTextString(m) }
func (*StreamReq) ProtoMessage()    {}
func (*StreamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{1}
}

func (m *StreamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamReq.Unmarshal(m, b)
}
func (m *StreamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamReq.Marshal(b, m, deterministic)
}
func (m *StreamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamReq.Merge(m, src)
}
func (m *StreamReq) XXX_Size() int {
	return xxx_messageInfo_StreamReq.Size(m)
}
func (m *StreamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamReq.DiscardUnknown(m)
}

var xxx_messageInfo_StreamReq proto.InternalMessageInfo

func (m *StreamReq) GetPt() *StreamPoint {
	if m != nil {
		return m.Pt
	}
	return nil
}

type StreamRes struct {
	Pt                   *StreamPoint `protobuf:"bytes,1,opt,name=pt,proto3" json:"pt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StreamRes) Reset()         { *m = StreamRes{} }
func (m *StreamRes) String() string { return proto.CompactTextString(m) }
func (*StreamRes) ProtoMessage()    {}
func (*StreamRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{2}
}

func (m *StreamRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRes.Unmarshal(m, b)
}
func (m *StreamRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRes.Marshal(b, m, deterministic)
}
func (m *StreamRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRes.Merge(m, src)
}
func (m *StreamRes) XXX_Size() int {
	return xxx_messageInfo_StreamRes.Size(m)
}
func (m *StreamRes) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRes.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRes proto.InternalMessageInfo

func (m *StreamRes) GetPt() *StreamPoint {
	if m != nil {
		return m.Pt
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamPoint)(nil), "proto.StreamPoint")
	proto.RegisterType((*StreamReq)(nil), "proto.StreamReq")
	proto.RegisterType((*StreamRes)(nil), "proto.StreamRes")
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54) }

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xe6, 0x5c, 0xdc,
	0xc1, 0x60, 0xe1, 0x80, 0xfc, 0xcc, 0xbc, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x84, 0x8b, 0xb5, 0x2c, 0x31, 0xa7,
	0x34, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0x35, 0x08, 0xc2, 0x51, 0xd2, 0xe7, 0xe2, 0x84, 0x68,
	0x0c, 0x4a, 0x2d, 0x14, 0x52, 0xe2, 0x62, 0x2a, 0x28, 0x01, 0x6b, 0xe2, 0x36, 0x12, 0x82, 0x58,
	0xa0, 0x87, 0x64, 0x6c, 0x10, 0x53, 0x41, 0x09, 0xb2, 0x86, 0x62, 0x62, 0x34, 0x18, 0x1d, 0x65,
	0xe4, 0xe2, 0x85, 0x88, 0x05, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x59, 0x71, 0x09, 0x80,
	0x98, 0xa9, 0x45, 0xc1, 0x99, 0x29, 0xa9, 0x10, 0x29, 0x21, 0x01, 0x14, 0xdd, 0x41, 0xa9, 0x85,
	0x52, 0xe8, 0x22, 0xc5, 0x4a, 0x0c, 0x06, 0x8c, 0x20, 0xbd, 0xce, 0x39, 0x99, 0xa9, 0x79, 0x25,
	0xa4, 0xea, 0xd5, 0x60, 0x14, 0xb2, 0xe1, 0x12, 0x70, 0xc9, 0x2f, 0x4d, 0xca, 0x49, 0x25, 0x5d,
	0xaf, 0x01, 0x63, 0x12, 0x1b, 0x58, 0xd8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x8b, 0x34,
	0x64, 0x80, 0x01, 0x00, 0x00,
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
	ServerSideStream(ctx context.Context, in *StreamReq, opts ...grpc.CallOption) (StreamService_ServerSideStreamClient, error)
	ClientSideStream(ctx context.Context, opts ...grpc.CallOption) (StreamService_ClientSideStreamClient, error)
	DoubleSideStream(ctx context.Context, opts ...grpc.CallOption) (StreamService_DoubleSideStreamClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) ServerSideStream(ctx context.Context, in *StreamReq, opts ...grpc.CallOption) (StreamService_ServerSideStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/proto.StreamService/ServerSideStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceServerSideStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ServerSideStreamClient interface {
	Recv() (*StreamRes, error)
	grpc.ClientStream
}

type streamServiceServerSideStreamClient struct {
	grpc.ClientStream
}

func (x *streamServiceServerSideStreamClient) Recv() (*StreamRes, error) {
	m := new(StreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) ClientSideStream(ctx context.Context, opts ...grpc.CallOption) (StreamService_ClientSideStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[1], "/proto.StreamService/ClientSideStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceClientSideStreamClient{stream}
	return x, nil
}

type StreamService_ClientSideStreamClient interface {
	Send(*StreamReq) error
	CloseAndRecv() (*StreamRes, error)
	grpc.ClientStream
}

type streamServiceClientSideStreamClient struct {
	grpc.ClientStream
}

func (x *streamServiceClientSideStreamClient) Send(m *StreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceClientSideStreamClient) CloseAndRecv() (*StreamRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) DoubleSideStream(ctx context.Context, opts ...grpc.CallOption) (StreamService_DoubleSideStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[2], "/proto.StreamService/DoubleSideStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceDoubleSideStreamClient{stream}
	return x, nil
}

type StreamService_DoubleSideStreamClient interface {
	Send(*StreamReq) error
	Recv() (*StreamRes, error)
	grpc.ClientStream
}

type streamServiceDoubleSideStreamClient struct {
	grpc.ClientStream
}

func (x *streamServiceDoubleSideStreamClient) Send(m *StreamReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceDoubleSideStreamClient) Recv() (*StreamRes, error) {
	m := new(StreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	ServerSideStream(*StreamReq, StreamService_ServerSideStreamServer) error
	ClientSideStream(StreamService_ClientSideStreamServer) error
	DoubleSideStream(StreamService_DoubleSideStreamServer) error
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_ServerSideStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).ServerSideStream(m, &streamServiceServerSideStreamServer{stream})
}

type StreamService_ServerSideStreamServer interface {
	Send(*StreamRes) error
	grpc.ServerStream
}

type streamServiceServerSideStreamServer struct {
	grpc.ServerStream
}

func (x *streamServiceServerSideStreamServer) Send(m *StreamRes) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_ClientSideStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).ClientSideStream(&streamServiceClientSideStreamServer{stream})
}

type StreamService_ClientSideStreamServer interface {
	SendAndClose(*StreamRes) error
	Recv() (*StreamReq, error)
	grpc.ServerStream
}

type streamServiceClientSideStreamServer struct {
	grpc.ServerStream
}

func (x *streamServiceClientSideStreamServer) SendAndClose(m *StreamRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceClientSideStreamServer) Recv() (*StreamReq, error) {
	m := new(StreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_DoubleSideStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).DoubleSideStream(&streamServiceDoubleSideStreamServer{stream})
}

type StreamService_DoubleSideStreamServer interface {
	Send(*StreamRes) error
	Recv() (*StreamReq, error)
	grpc.ServerStream
}

type streamServiceDoubleSideStreamServer struct {
	grpc.ServerStream
}

func (x *streamServiceDoubleSideStreamServer) Send(m *StreamRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceDoubleSideStreamServer) Recv() (*StreamReq, error) {
	m := new(StreamReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSideStream",
			Handler:       _StreamService_ServerSideStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientSideStream",
			Handler:       _StreamService_ClientSideStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DoubleSideStream",
			Handler:       _StreamService_DoubleSideStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
