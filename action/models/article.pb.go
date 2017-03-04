// Code generated by protoc-gen-go.
// source: article.proto
// DO NOT EDIT!

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Article struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UserId      uint64 `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Data        string `protobuf:"bytes,6,opt,name=data" json:"data,omitempty"`
	CreatedUtc  int64  `protobuf:"varint,14,opt,name=created_utc,json=createdUtc" json:"created_utc,omitempty"`
	ModifiedUtc int64  `protobuf:"varint,15,opt,name=modified_utc,json=modifiedUtc" json:"modified_utc,omitempty"`
	Deleted     bool   `protobuf:"varint,16,opt,name=deleted" json:"deleted,omitempty"`
	DeletedUtc  int64  `protobuf:"varint,17,opt,name=deleted_utc,json=deletedUtc" json:"deleted_utc,omitempty"`
}

func (m *Article) Reset()                    { *m = Article{} }
func (m *Article) String() string            { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()               {}
func (*Article) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Article) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Article) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Article) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Article) GetCreatedUtc() int64 {
	if m != nil {
		return m.CreatedUtc
	}
	return 0
}

func (m *Article) GetModifiedUtc() int64 {
	if m != nil {
		return m.ModifiedUtc
	}
	return 0
}

func (m *Article) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *Article) GetDeletedUtc() int64 {
	if m != nil {
		return m.DeletedUtc
	}
	return 0
}

func init() {
	proto.RegisterType((*Article)(nil), "models.Article")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ArticleRPC service

type ArticleRPCClient interface {
	GetArticleById(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error)
}

type articleRPCClient struct {
	cc *grpc.ClientConn
}

func NewArticleRPCClient(cc *grpc.ClientConn) ArticleRPCClient {
	return &articleRPCClient{cc}
}

func (c *articleRPCClient) GetArticleById(ctx context.Context, in *Article, opts ...grpc.CallOption) (*Article, error) {
	out := new(Article)
	err := grpc.Invoke(ctx, "/models.ArticleRPC/GetArticleById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ArticleRPC service

type ArticleRPCServer interface {
	GetArticleById(context.Context, *Article) (*Article, error)
}

func RegisterArticleRPCServer(s *grpc.Server, srv ArticleRPCServer) {
	s.RegisterService(&_ArticleRPC_serviceDesc, srv)
}

func _ArticleRPC_GetArticleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticleRPCServer).GetArticleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.ArticleRPC/GetArticleById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticleRPCServer).GetArticleById(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArticleRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.ArticleRPC",
	HandlerType: (*ArticleRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArticleById",
			Handler:    _ArticleRPC_GetArticleById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "article.proto",
}

func init() { proto.RegisterFile("article.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x71, 0xa8, 0x12, 0xb8, 0x42, 0x0a, 0xb7, 0x60, 0xb1, 0x10, 0x3a, 0x79, 0xca, 0x00,
	0xfc, 0x01, 0xca, 0x80, 0xba, 0x21, 0x4b, 0x9d, 0x2b, 0xe3, 0x3b, 0x24, 0x4b, 0xa9, 0x8c, 0xdc,
	0xeb, 0xc0, 0x0f, 0xe4, 0x7f, 0xa1, 0x38, 0xce, 0xc2, 0xf6, 0xde, 0xe7, 0xe7, 0x27, 0xbd, 0x83,
	0x6b, 0x97, 0x24, 0xf8, 0x81, 0xfb, 0xef, 0x14, 0x25, 0x62, 0x7d, 0x88, 0xc4, 0xc3, 0x71, 0xfd,
	0xab, 0xa0, 0x79, 0x9d, 0x5e, 0xb0, 0x85, 0x2a, 0x90, 0x56, 0x9d, 0x32, 0x0b, 0x5b, 0x05, 0xc2,
	0x3b, 0x68, 0x4e, 0x47, 0x4e, 0xfb, 0x40, 0xba, 0xca, 0xb0, 0x1e, 0xed, 0x96, 0x10, 0x61, 0x41,
	0x4e, 0x9c, 0xae, 0x3b, 0x65, 0x2e, 0x6d, 0xd6, 0xf8, 0x00, 0x4b, 0x9f, 0xd8, 0x09, 0xd3, 0xfe,
	0x24, 0x5e, 0xb7, 0x9d, 0x32, 0xe7, 0x16, 0x0a, 0xda, 0x89, 0xc7, 0x47, 0xb8, 0x3a, 0x44, 0x0a,
	0x5f, 0xa1, 0x24, 0x56, 0x39, 0xb1, 0x9c, 0xd9, 0x18, 0xd1, 0xd0, 0x10, 0x0f, 0x2c, 0x4c, 0xfa,
	0xa6, 0x53, 0xe6, 0xc2, 0xce, 0x76, 0x6c, 0x2f, 0x32, 0xff, 0xbd, 0x9d, 0xda, 0x0b, 0xda, 0x89,
	0x7f, 0xda, 0x00, 0x94, 0x19, 0xf6, 0xe3, 0x0d, 0x5f, 0xa0, 0x7d, 0x67, 0x29, 0x60, 0xf3, 0xb3,
	0x25, 0x5c, 0xf5, 0xd3, 0xe0, 0xbe, 0xc0, 0xfb, 0xff, 0x60, 0x7d, 0xf6, 0x59, 0xe7, 0xd3, 0x3c,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x26, 0xae, 0xce, 0x03, 0x2b, 0x01, 0x00, 0x00,
}