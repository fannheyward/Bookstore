// Code generated by protoc-gen-go.
// source: pb/book.proto
// DO NOT EDIT!

/*
Package im_fann_book is a generated protocol buffer package.

It is generated from these files:
	pb/book.proto

It has these top-level messages:
	Book
	BookResp
*/
package im_fann_book

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Book struct {
	Id    int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
}

func (m *Book) Reset()                    { *m = Book{} }
func (m *Book) String() string            { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()               {}
func (*Book) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Book) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type BookResp struct {
	Books []*Book `protobuf:"bytes,1,rep,name=books" json:"books,omitempty"`
}

func (m *BookResp) Reset()                    { *m = BookResp{} }
func (m *BookResp) String() string            { return proto.CompactTextString(m) }
func (*BookResp) ProtoMessage()               {}
func (*BookResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BookResp) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
}

func init() {
	proto.RegisterType((*Book)(nil), "im.fann.book.Book")
	proto.RegisterType((*BookResp)(nil), "im.fann.book.BookResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BookService service

type BookServiceClient interface {
	ListBooks(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*BookResp, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) ListBooks(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*BookResp, error) {
	out := new(BookResp)
	err := grpc.Invoke(ctx, "/im.fann.book.BookService/ListBooks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BookService service

type BookServiceServer interface {
	ListBooks(context.Context, *google_protobuf1.Empty) (*BookResp, error)
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/im.fann.book.BookService/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ListBooks(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "im.fann.book.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _BookService_ListBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/book.proto",
}

func init() { proto.RegisterFile("pb/book.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xd2, 0x4f,
	0xca, 0xcf, 0xcf, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc9, 0xcc, 0xd5, 0x4b, 0x4b,
	0xcc, 0xcb, 0xd3, 0x03, 0x89, 0x49, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16,
	0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0xd4, 0x4a,
	0x49, 0x43, 0x65, 0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0xd4, 0xdc, 0x82, 0x92, 0x4a, 0x88, 0xa4,
	0x92, 0x0e, 0x17, 0x8b, 0x53, 0x7e, 0x7e, 0xb6, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x73, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x66, 0x49, 0x4e,
	0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0xa3, 0x64, 0xc2, 0xc5, 0x01, 0x52, 0x1d,
	0x94, 0x5a, 0x5c, 0x20, 0xa4, 0xc1, 0xc5, 0x0a, 0xb2, 0xbc, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83,
	0xdb, 0x48, 0x48, 0x0f, 0xd9, 0x49, 0x7a, 0x60, 0x65, 0x10, 0x05, 0x46, 0xb1, 0x5c, 0xdc, 0x20,
	0x6e, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x1f, 0x17, 0xa7, 0x4f, 0x66, 0x71, 0x09,
	0x48, 0xa8, 0x58, 0x48, 0x4c, 0x0f, 0xe2, 0x3a, 0x3d, 0x98, 0xeb, 0xf4, 0x5c, 0x41, 0xae, 0x93,
	0x12, 0xc3, 0x62, 0x5c, 0x6a, 0x71, 0x81, 0x92, 0x60, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xb8, 0x85,
	0x38, 0xf5, 0xcb, 0x0c, 0xc1, 0x01, 0x52, 0xec, 0x24, 0xc6, 0x85, 0x12, 0x1a, 0x4e, 0x60, 0x0f,
	0x05, 0x30, 0x26, 0xb1, 0x81, 0x8d, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x96, 0xbc,
	0x92, 0x3b, 0x01, 0x00, 0x00,
}