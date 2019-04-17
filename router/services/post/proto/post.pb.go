// Code generated by protoc-gen-go. DO NOT EDIT.
// source: post.proto

package proto

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CategoryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryRequest) Reset()         { *m = CategoryRequest{} }
func (m *CategoryRequest) String() string { return proto.CompactTextString(m) }
func (*CategoryRequest) ProtoMessage()    {}
func (*CategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{0}
}
func (m *CategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryRequest.Unmarshal(m, b)
}
func (m *CategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryRequest.Marshal(b, m, deterministic)
}
func (m *CategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryRequest.Merge(m, src)
}
func (m *CategoryRequest) XXX_Size() int {
	return xxx_messageInfo_CategoryRequest.Size(m)
}
func (m *CategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryRequest proto.InternalMessageInfo

type Category struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{1}
}
func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Category) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type CategoryResponse struct {
	Categories           []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	Result               bool        `protobuf:"varint,2,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CategoryResponse) Reset()         { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()    {}
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{2}
}
func (m *CategoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryResponse.Unmarshal(m, b)
}
func (m *CategoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryResponse.Marshal(b, m, deterministic)
}
func (m *CategoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryResponse.Merge(m, src)
}
func (m *CategoryResponse) XXX_Size() int {
	return xxx_messageInfo_CategoryResponse.Size(m)
}
func (m *CategoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryResponse proto.InternalMessageInfo

func (m *CategoryResponse) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *CategoryResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type PostRequest struct {
	Category             int32    `protobuf:"varint,1,opt,name=Category,proto3" json:"Category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostRequest) Reset()         { *m = PostRequest{} }
func (m *PostRequest) String() string { return proto.CompactTextString(m) }
func (*PostRequest) ProtoMessage()    {}
func (*PostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{3}
}
func (m *PostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostRequest.Unmarshal(m, b)
}
func (m *PostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostRequest.Marshal(b, m, deterministic)
}
func (m *PostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostRequest.Merge(m, src)
}
func (m *PostRequest) XXX_Size() int {
	return xxx_messageInfo_PostRequest.Size(m)
}
func (m *PostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostRequest proto.InternalMessageInfo

func (m *PostRequest) GetCategory() int32 {
	if m != nil {
		return m.Category
	}
	return 0
}

type Post struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Category             int32    `protobuf:"varint,4,opt,name=category,proto3" json:"category,omitempty"`
	Created              string   `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{4}
}
func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Post) GetCategory() int32 {
	if m != nil {
		return m.Category
	}
	return 0
}

func (m *Post) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

type PostResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	Result               bool     `protobuf:"varint,2,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostResponse) Reset()         { *m = PostResponse{} }
func (m *PostResponse) String() string { return proto.CompactTextString(m) }
func (*PostResponse) ProtoMessage()    {}
func (*PostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e114ad14deab1dd1, []int{5}
}
func (m *PostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResponse.Unmarshal(m, b)
}
func (m *PostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResponse.Marshal(b, m, deterministic)
}
func (m *PostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResponse.Merge(m, src)
}
func (m *PostResponse) XXX_Size() int {
	return xxx_messageInfo_PostResponse.Size(m)
}
func (m *PostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostResponse proto.InternalMessageInfo

func (m *PostResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *PostResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*CategoryRequest)(nil), "post.CategoryRequest")
	proto.RegisterType((*Category)(nil), "post.Category")
	proto.RegisterType((*CategoryResponse)(nil), "post.CategoryResponse")
	proto.RegisterType((*PostRequest)(nil), "post.PostRequest")
	proto.RegisterType((*Post)(nil), "post.Post")
	proto.RegisterType((*PostResponse)(nil), "post.PostResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostsClient is the client API for Posts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostsClient interface {
	GetCategories(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	GetPosts(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
}

type postsClient struct {
	cc *grpc.ClientConn
}

func NewPostsClient(cc *grpc.ClientConn) PostsClient {
	return &postsClient{cc}
}

func (c *postsClient) GetCategories(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, "/post.Posts/GetCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsClient) GetPosts(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/post.Posts/GetPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostsServer is the server API for Posts service.
type PostsServer interface {
	GetCategories(context.Context, *CategoryRequest) (*CategoryResponse, error)
	GetPosts(context.Context, *PostRequest) (*PostResponse, error)
}

func RegisterPostsServer(s *grpc.Server, srv PostsServer) {
	s.RegisterService(&_Posts_serviceDesc, srv)
}

func _Posts_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.Posts/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetCategories(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Posts_GetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostsServer).GetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.Posts/GetPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostsServer).GetPosts(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Posts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "post.Posts",
	HandlerType: (*PostsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategories",
			Handler:    _Posts_GetCategories_Handler,
		},
		{
			MethodName: "GetPosts",
			Handler:    _Posts_GetPosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post.proto",
}

func init() { proto.RegisterFile("post.proto", fileDescriptor_e114ad14deab1dd1) }

var fileDescriptor_e114ad14deab1dd1 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0x69, 0x52, 0xe3, 0x54, 0xab, 0x1d, 0xb4, 0x2c, 0x3d, 0x85, 0x3d, 0xc5, 0x4b, 0x90,
	0xf6, 0x0f, 0x08, 0x3d, 0xd4, 0x63, 0xc9, 0xd1, 0x5b, 0x4d, 0x06, 0x09, 0x14, 0x37, 0x66, 0xa7,
	0xa2, 0x17, 0x7f, 0xbb, 0x64, 0x3f, 0x9a, 0x50, 0x10, 0x7a, 0x9b, 0xf7, 0xf2, 0xe6, 0xe5, 0xbd,
	0x61, 0x01, 0x1a, 0xa5, 0x39, 0x6f, 0x5a, 0xc5, 0x0a, 0xa3, 0x6e, 0x96, 0x33, 0xb8, 0x5d, 0xef,
	0x98, 0xde, 0x55, 0xfb, 0x53, 0xd0, 0xe7, 0x81, 0x34, 0xcb, 0x27, 0x48, 0x3c, 0x85, 0x53, 0x08,
	0xeb, 0x4a, 0x04, 0x69, 0x90, 0xc5, 0x45, 0x58, 0x57, 0x78, 0x0f, 0x31, 0xd7, 0xbc, 0x27, 0x11,
	0xa6, 0x41, 0x76, 0x55, 0x58, 0x20, 0x5f, 0xe1, 0xae, 0x37, 0xd1, 0x8d, 0xfa, 0xd0, 0x84, 0x39,
	0x40, 0x69, 0xb9, 0x9a, 0xb4, 0x08, 0xd2, 0x51, 0x36, 0x59, 0x4e, 0x73, 0xf3, 0xff, 0xa3, 0x76,
	0xa0, 0xc0, 0x39, 0x8c, 0x0b, 0xd2, 0x87, 0x3d, 0x1b, 0xeb, 0xa4, 0x70, 0x48, 0x3e, 0xc2, 0x64,
	0xab, 0x34, 0xbb, 0x70, 0xb8, 0xe8, 0xc3, 0xb9, 0x58, 0x47, 0x2c, 0xbf, 0x20, 0xea, 0xa4, 0xe7,
	0x85, 0x46, 0x84, 0x88, 0xe9, 0x9b, 0xc5, 0xc8, 0x90, 0x66, 0xee, 0xdc, 0x4b, 0xef, 0x1e, 0x59,
	0x77, 0x8f, 0x51, 0xc0, 0x65, 0xd9, 0xd2, 0x8e, 0xa9, 0x12, 0xb1, 0x59, 0xf1, 0x50, 0xbe, 0xc0,
	0xb5, 0x8d, 0xe8, 0xaa, 0xa7, 0x10, 0x77, 0x3d, 0x7d, 0x6b, 0xb0, 0xad, 0x8d, 0xc4, 0x7e, 0xf8,
	0xaf, 0xec, 0xf2, 0x17, 0xe2, 0xad, 0x11, 0x3c, 0xc3, 0xcd, 0x86, 0x78, 0xdd, 0x9f, 0xe7, 0xe1,
	0xe4, 0x74, 0xf6, 0x1c, 0x8b, 0xf9, 0x29, 0x6d, 0x23, 0xc8, 0x0b, 0x5c, 0x41, 0xb2, 0x21, 0xb6,
	0x6e, 0xb3, 0x41, 0x02, 0xb7, 0x88, 0x43, 0xca, 0x2f, 0xbd, 0x8d, 0xcd, 0xd3, 0x58, 0xfd, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xab, 0x39, 0xa7, 0x68, 0x28, 0x02, 0x00, 0x00,
}
