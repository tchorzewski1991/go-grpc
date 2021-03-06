// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blogpb/blog.proto

package blogpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRequest) Reset()         { *m = CreateBlogRequest{} }
func (m *CreateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRequest) ProtoMessage()    {}
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{1}
}

func (m *CreateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRequest.Unmarshal(m, b)
}
func (m *CreateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRequest.Marshal(b, m, deterministic)
}
func (m *CreateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRequest.Merge(m, src)
}
func (m *CreateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRequest.Size(m)
}
func (m *CreateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRequest proto.InternalMessageInfo

func (m *CreateBlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogResponse) Reset()         { *m = CreateBlogResponse{} }
func (m *CreateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBlogResponse) ProtoMessage()    {}
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{2}
}

func (m *CreateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogResponse.Unmarshal(m, b)
}
func (m *CreateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogResponse.Marshal(b, m, deterministic)
}
func (m *CreateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogResponse.Merge(m, src)
}
func (m *CreateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBlogResponse.Size(m)
}
func (m *CreateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogResponse proto.InternalMessageInfo

func (m *CreateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ReadBlogRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogRequest) Reset()         { *m = ReadBlogRequest{} }
func (m *ReadBlogRequest) String() string { return proto.CompactTextString(m) }
func (*ReadBlogRequest) ProtoMessage()    {}
func (*ReadBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{3}
}

func (m *ReadBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogRequest.Unmarshal(m, b)
}
func (m *ReadBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogRequest.Marshal(b, m, deterministic)
}
func (m *ReadBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogRequest.Merge(m, src)
}
func (m *ReadBlogRequest) XXX_Size() int {
	return xxx_messageInfo_ReadBlogRequest.Size(m)
}
func (m *ReadBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogRequest proto.InternalMessageInfo

func (m *ReadBlogRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

type ReadBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogResponse) Reset()         { *m = ReadBlogResponse{} }
func (m *ReadBlogResponse) String() string { return proto.CompactTextString(m) }
func (*ReadBlogResponse) ProtoMessage()    {}
func (*ReadBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{4}
}

func (m *ReadBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogResponse.Unmarshal(m, b)
}
func (m *ReadBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogResponse.Marshal(b, m, deterministic)
}
func (m *ReadBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogResponse.Merge(m, src)
}
func (m *ReadBlogResponse) XXX_Size() int {
	return xxx_messageInfo_ReadBlogResponse.Size(m)
}
func (m *ReadBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogResponse proto.InternalMessageInfo

func (m *ReadBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ListBlogRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogRequest) Reset()         { *m = ListBlogRequest{} }
func (m *ListBlogRequest) String() string { return proto.CompactTextString(m) }
func (*ListBlogRequest) ProtoMessage()    {}
func (*ListBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{5}
}

func (m *ListBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogRequest.Unmarshal(m, b)
}
func (m *ListBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogRequest.Marshal(b, m, deterministic)
}
func (m *ListBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogRequest.Merge(m, src)
}
func (m *ListBlogRequest) XXX_Size() int {
	return xxx_messageInfo_ListBlogRequest.Size(m)
}
func (m *ListBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogRequest proto.InternalMessageInfo

type ListBlogResponse struct {
	Blogs                []*Blog  `protobuf:"bytes,1,rep,name=blogs,proto3" json:"blogs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogResponse) Reset()         { *m = ListBlogResponse{} }
func (m *ListBlogResponse) String() string { return proto.CompactTextString(m) }
func (*ListBlogResponse) ProtoMessage()    {}
func (*ListBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cd072c3eda6f7ba, []int{6}
}

func (m *ListBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogResponse.Unmarshal(m, b)
}
func (m *ListBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogResponse.Marshal(b, m, deterministic)
}
func (m *ListBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogResponse.Merge(m, src)
}
func (m *ListBlogResponse) XXX_Size() int {
	return xxx_messageInfo_ListBlogResponse.Size(m)
}
func (m *ListBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogResponse proto.InternalMessageInfo

func (m *ListBlogResponse) GetBlogs() []*Blog {
	if m != nil {
		return m.Blogs
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogRequest)(nil), "blog.CreateBlogRequest")
	proto.RegisterType((*CreateBlogResponse)(nil), "blog.CreateBlogResponse")
	proto.RegisterType((*ReadBlogRequest)(nil), "blog.ReadBlogRequest")
	proto.RegisterType((*ReadBlogResponse)(nil), "blog.ReadBlogResponse")
	proto.RegisterType((*ListBlogRequest)(nil), "blog.ListBlogRequest")
	proto.RegisterType((*ListBlogResponse)(nil), "blog.ListBlogResponse")
}

func init() { proto.RegisterFile("blogpb/blog.proto", fileDescriptor_1cd072c3eda6f7ba) }

var fileDescriptor_1cd072c3eda6f7ba = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0xd3, 0x52, 0x0a, 0x0c, 0x09, 0x7f, 0x26, 0x28, 0x1b, 0x34, 0x86, 0xec, 0xc9, 0x78,
	0xa0, 0x06, 0x38, 0x79, 0xc4, 0x83, 0x21, 0x31, 0x31, 0xe2, 0xcd, 0x0b, 0x29, 0x74, 0x53, 0x37,
	0x69, 0xba, 0xb5, 0xbb, 0x78, 0x31, 0x5e, 0x7c, 0x05, 0x1f, 0xcd, 0x57, 0xf0, 0x29, 0x3c, 0x99,
	0xdd, 0x2d, 0x42, 0xca, 0x41, 0x4f, 0x30, 0xdf, 0xcc, 0xfc, 0x66, 0xe6, 0xeb, 0x42, 0x77, 0x95,
	0x88, 0x38, 0x5b, 0x05, 0xfa, 0x67, 0x94, 0xe5, 0x42, 0x09, 0xf4, 0xf4, 0xff, 0xc1, 0x69, 0x2c,
	0x44, 0x9c, 0xb0, 0x20, 0xcc, 0x78, 0x10, 0xa6, 0xa9, 0x50, 0xa1, 0xe2, 0x22, 0x95, 0xb6, 0x86,
	0xae, 0xc1, 0x9b, 0x25, 0x22, 0xc6, 0x16, 0xb8, 0x3c, 0x22, 0xce, 0xd0, 0x39, 0x6f, 0x2c, 0x5c,
	0x1e, 0xe1, 0x09, 0x34, 0xc2, 0x8d, 0x7a, 0x12, 0xf9, 0x92, 0x47, 0xc4, 0x35, 0x72, 0xdd, 0x0a,
	0xf3, 0x08, 0x7b, 0x50, 0x55, 0x5c, 0x25, 0x8c, 0x54, 0x4c, 0xc2, 0x06, 0x48, 0xa0, 0xb6, 0x16,
	0xa9, 0x62, 0xa9, 0x22, 0x9e, 0xd1, 0xb7, 0x21, 0x9d, 0x40, 0xf7, 0x3a, 0x67, 0xa1, 0x62, 0x7a,
	0xd4, 0x82, 0x3d, 0x6f, 0x98, 0x54, 0x78, 0x06, 0x66, 0x3f, 0x33, 0xb3, 0x39, 0x86, 0x91, 0x59,
	0xdc, 0x14, 0x18, 0x9d, 0x4e, 0x01, 0xf7, 0x9b, 0x64, 0x26, 0x52, 0xc9, 0xfe, 0xec, 0xba, 0x80,
	0xf6, 0x82, 0x85, 0xd1, 0xfe, 0xa0, 0x3e, 0xd4, 0x74, 0x6a, 0xf9, 0x7b, 0x9f, 0xaf, 0xc3, 0x79,
	0x44, 0xc7, 0xd0, 0xd9, 0xd5, 0xfe, 0x93, 0xdf, 0x85, 0xf6, 0x2d, 0x97, 0x6a, 0x8f, 0x4f, 0xa7,
	0xd0, 0xd9, 0x49, 0x05, 0x66, 0x08, 0x55, 0x5d, 0x2e, 0x89, 0x33, 0xac, 0x94, 0x38, 0x36, 0x31,
	0xfe, 0x76, 0xa0, 0xa9, 0xe3, 0x07, 0x96, 0xbf, 0xf0, 0x35, 0xc3, 0x7b, 0xf0, 0xed, 0xb9, 0xd8,
	0xb7, 0xc5, 0x07, 0x8e, 0x0d, 0xc8, 0x61, 0xc2, 0x8e, 0xa3, 0xbd, 0xf7, 0xcf, 0xaf, 0x0f, 0xb7,
	0x45, 0x7d, 0xf3, 0xf9, 0xe5, 0x95, 0xd9, 0x15, 0xef, 0xc0, 0xd3, 0xf7, 0xe1, 0x91, 0xed, 0x2b,
	0xf9, 0x32, 0x38, 0x2e, 0xcb, 0x05, 0x8c, 0x18, 0x18, 0x62, 0xc7, 0xc2, 0x82, 0xd7, 0xc2, 0xbd,
	0x37, 0xbc, 0x01, 0x4f, 0x5f, 0xba, 0x05, 0x96, 0x8c, 0xd8, 0x02, 0xcb, 0x66, 0xd0, 0x96, 0x01,
	0xd6, 0xb1, 0xd8, 0xee, 0xd2, 0x99, 0xd5, 0x1f, 0x7d, 0xfb, 0x5c, 0x57, 0xbe, 0x79, 0x86, 0x93,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x3e, 0xb9, 0xad, 0xbf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	Create(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error)
	Read(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error)
	List(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (BlogService_ListClient, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) Create(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error) {
	out := new(CreateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Read(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error) {
	out := new(ReadBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) List(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (BlogService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlogService_serviceDesc.Streams[0], "/blog.BlogService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogService_ListClient interface {
	Recv() (*ListBlogResponse, error)
	grpc.ClientStream
}

type blogServiceListClient struct {
	grpc.ClientStream
}

func (x *blogServiceListClient) Recv() (*ListBlogResponse, error) {
	m := new(ListBlogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	Create(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error)
	Read(context.Context, *ReadBlogRequest) (*ReadBlogResponse, error)
	List(*ListBlogRequest, BlogService_ListServer) error
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Create(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Read(ctx, req.(*ReadBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServiceServer).List(m, &blogServiceListServer{stream})
}

type BlogService_ListServer interface {
	Send(*ListBlogResponse) error
	grpc.ServerStream
}

type blogServiceListServer struct {
	grpc.ServerStream
}

func (x *blogServiceListServer) Send(m *ListBlogResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BlogService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _BlogService_Read_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _BlogService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blogpb/blog.proto",
}
