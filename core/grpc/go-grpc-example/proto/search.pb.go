// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

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

type SearchReq struct {
	Req                  string   `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchReq) Reset()         { *m = SearchReq{} }
func (m *SearchReq) String() string { return proto.CompactTextString(m) }
func (*SearchReq) ProtoMessage()    {}
func (*SearchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{0}
}

func (m *SearchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReq.Unmarshal(m, b)
}
func (m *SearchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReq.Marshal(b, m, deterministic)
}
func (m *SearchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReq.Merge(m, src)
}
func (m *SearchReq) XXX_Size() int {
	return xxx_messageInfo_SearchReq.Size(m)
}
func (m *SearchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReq.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReq proto.InternalMessageInfo

func (m *SearchReq) GetReq() string {
	if m != nil {
		return m.Req
	}
	return ""
}

type SearchRes struct {
	Res                  string   `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRes) Reset()         { *m = SearchRes{} }
func (m *SearchRes) String() string { return proto.CompactTextString(m) }
func (*SearchRes) ProtoMessage()    {}
func (*SearchRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{1}
}

func (m *SearchRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRes.Unmarshal(m, b)
}
func (m *SearchRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRes.Marshal(b, m, deterministic)
}
func (m *SearchRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRes.Merge(m, src)
}
func (m *SearchRes) XXX_Size() int {
	return xxx_messageInfo_SearchRes.Size(m)
}
func (m *SearchRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRes.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRes proto.InternalMessageInfo

func (m *SearchRes) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func init() {
	proto.RegisterType((*SearchReq)(nil), "proto.SearchReq")
	proto.RegisterType((*SearchRes)(nil), "proto.SearchRes")
}

func init() { proto.RegisterFile("search.proto", fileDescriptor_453745cff914010e) }

var fileDescriptor_453745cff914010e = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xb2, 0x5c, 0x9c,
	0xc1, 0x60, 0xe1, 0xa0, 0xd4, 0x42, 0x21, 0x01, 0x2e, 0xe6, 0xa2, 0xd4, 0x42, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0xce, 0x20, 0x10, 0x13, 0x59, 0xba, 0x18, 0x22, 0x5d, 0x8c, 0x90, 0x2e, 0x36, 0xb2,
	0xe7, 0xe2, 0x85, 0x48, 0x07, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0xe9, 0x71, 0xb1, 0x41,
	0x04, 0x84, 0x04, 0x20, 0xf6, 0xe8, 0xc1, 0x4d, 0x97, 0x42, 0x17, 0x29, 0x56, 0x62, 0x48, 0x62,
	0x03, 0x0b, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x32, 0x4d, 0x22, 0x6b, 0x9c, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchRes, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchRes, error) {
	out := new(SearchRes)
	err := c.cc.Invoke(ctx, "/proto.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	Search(context.Context, *SearchReq) (*SearchRes, error)
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*SearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}