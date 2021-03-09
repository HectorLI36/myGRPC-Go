// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: slice.proto

package slice

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SliceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content      string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	SliceNumber  int32  `protobuf:"varint,3,opt,name=slice_number,json=sliceNumber,proto3" json:"slice_number,omitempty"`
	SlicePerPage int32  `protobuf:"varint,4,opt,name=slice_per_page,json=slicePerPage,proto3" json:"slice_per_page,omitempty"`
	Offset       int32  `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	BookId       int32  `protobuf:"varint,6,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (x *SliceResponse) Reset() {
	*x = SliceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SliceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceResponse) ProtoMessage() {}

func (x *SliceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_slice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceResponse.ProtoReflect.Descriptor instead.
func (*SliceResponse) Descriptor() ([]byte, []int) {
	return file_slice_proto_rawDescGZIP(), []int{0}
}

func (x *SliceResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SliceResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SliceResponse) GetSliceNumber() int32 {
	if x != nil {
		return x.SliceNumber
	}
	return 0
}

func (x *SliceResponse) GetSlicePerPage() int32 {
	if x != nil {
		return x.SlicePerPage
	}
	return 0
}

func (x *SliceResponse) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SliceResponse) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

type SliceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId       int32 `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Offset       int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	SlicePerPage int32 `protobuf:"varint,3,opt,name=slice_per_page,json=slicePerPage,proto3" json:"slice_per_page,omitempty"`
}

func (x *SliceRequest) Reset() {
	*x = SliceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_slice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SliceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SliceRequest) ProtoMessage() {}

func (x *SliceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_slice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SliceRequest.ProtoReflect.Descriptor instead.
func (*SliceRequest) Descriptor() ([]byte, []int) {
	return file_slice_proto_rawDescGZIP(), []int{1}
}

func (x *SliceRequest) GetBookId() int32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *SliceRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SliceRequest) GetSlicePerPage() int32 {
	if x != nil {
		return x.SlicePerPage
	}
	return 0
}

var File_slice_proto protoreflect.FileDescriptor

var file_slice_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01,
	0x0a, 0x0d, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x6c, 0x69, 0x63,
	0x65, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x0c, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6c,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65,
	0x32, 0x38, 0x0a, 0x0b, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x0d, 0x2e, 0x53, 0x6c,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x53, 0x6c, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_slice_proto_rawDescOnce sync.Once
	file_slice_proto_rawDescData = file_slice_proto_rawDesc
)

func file_slice_proto_rawDescGZIP() []byte {
	file_slice_proto_rawDescOnce.Do(func() {
		file_slice_proto_rawDescData = protoimpl.X.CompressGZIP(file_slice_proto_rawDescData)
	})
	return file_slice_proto_rawDescData
}

var file_slice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_slice_proto_goTypes = []interface{}{
	(*SliceResponse)(nil), // 0: SliceResponse
	(*SliceRequest)(nil),  // 1: SliceRequest
}
var file_slice_proto_depIdxs = []int32{
	1, // 0: SliceServer.GetSlice:input_type -> SliceRequest
	0, // 1: SliceServer.GetSlice:output_type -> SliceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_slice_proto_init() }
func file_slice_proto_init() {
	if File_slice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_slice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SliceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_slice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SliceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_slice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_slice_proto_goTypes,
		DependencyIndexes: file_slice_proto_depIdxs,
		MessageInfos:      file_slice_proto_msgTypes,
	}.Build()
	File_slice_proto = out.File
	file_slice_proto_rawDesc = nil
	file_slice_proto_goTypes = nil
	file_slice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SliceServerClient is the client API for SliceServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SliceServerClient interface {
	GetSlice(ctx context.Context, in *SliceRequest, opts ...grpc.CallOption) (*SliceResponse, error)
}

type sliceServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSliceServerClient(cc grpc.ClientConnInterface) SliceServerClient {
	return &sliceServerClient{cc}
}

func (c *sliceServerClient) GetSlice(ctx context.Context, in *SliceRequest, opts ...grpc.CallOption) (*SliceResponse, error) {
	out := new(SliceResponse)
	err := c.cc.Invoke(ctx, "/SliceServer/GetSlice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SliceServerServer is the server API for SliceServer service.
type SliceServerServer interface {
	GetSlice(context.Context, *SliceRequest) (*SliceResponse, error)
}

// UnimplementedSliceServerServer can be embedded to have forward compatible implementations.
type UnimplementedSliceServerServer struct {
}

func (*UnimplementedSliceServerServer) GetSlice(context.Context, *SliceRequest) (*SliceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlice not implemented")
}

func RegisterSliceServerServer(s *grpc.Server, srv SliceServerServer) {
	s.RegisterService(&_SliceServer_serviceDesc, srv)
}

func _SliceServer_GetSlice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SliceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SliceServerServer).GetSlice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SliceServer/GetSlice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SliceServerServer).GetSlice(ctx, req.(*SliceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SliceServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SliceServer",
	HandlerType: (*SliceServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSlice",
			Handler:    _SliceServer_GetSlice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "slice.proto",
}