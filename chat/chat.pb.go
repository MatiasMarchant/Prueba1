// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: chat.proto

package chat

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

type Ordenclientepymes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto    string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor       int32  `protobuf:"varint,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda      string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino     string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	Prioritario bool   `protobuf:"varint,6,opt,name=prioritario,proto3" json:"prioritario,omitempty"`
}

func (x *Ordenclientepymes) Reset() {
	*x = Ordenclientepymes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ordenclientepymes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ordenclientepymes) ProtoMessage() {}

func (x *Ordenclientepymes) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ordenclientepymes.ProtoReflect.Descriptor instead.
func (*Ordenclientepymes) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Ordenclientepymes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ordenclientepymes) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *Ordenclientepymes) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Ordenclientepymes) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *Ordenclientepymes) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *Ordenclientepymes) GetPrioritario() bool {
	if x != nil {
		return x.Prioritario
	}
	return false
}

type Ordenclienteretail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor    int32  `protobuf:"varint,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda   string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino  string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *Ordenclienteretail) Reset() {
	*x = Ordenclienteretail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ordenclienteretail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ordenclienteretail) ProtoMessage() {}

func (x *Ordenclienteretail) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ordenclienteretail.ProtoReflect.Descriptor instead.
func (*Ordenclienteretail) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Ordenclienteretail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ordenclienteretail) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *Ordenclienteretail) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Ordenclienteretail) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *Ordenclienteretail) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x65, 0x70, 0x79, 0x6d, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69,
	0x65, 0x6e, 0x64, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x6e,
	0x64, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x22, 0x88,
	0x01, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x32, 0xa2, 0x01, 0x0a, 0x0b, 0x43, 0x68,
	0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x11, 0x52, 0x65, 0x64,
	0x65, 0x63, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x50, 0x79, 0x6d, 0x65, 0x73, 0x12, 0x17,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x65, 0x70, 0x79, 0x6d, 0x65, 0x73, 0x1a, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x70, 0x79, 0x6d, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x4a, 0x0a, 0x12, 0x52, 0x65, 0x64, 0x65, 0x63, 0x69, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x6e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chat_proto_goTypes = []interface{}{
	(*Ordenclientepymes)(nil),  // 0: chat.Ordenclientepymes
	(*Ordenclienteretail)(nil), // 1: chat.Ordenclienteretail
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.RedecirOrdenPymes:input_type -> chat.Ordenclientepymes
	1, // 1: chat.ChatService.RedecirOrdenRetail:input_type -> chat.Ordenclienteretail
	0, // 2: chat.ChatService.RedecirOrdenPymes:output_type -> chat.Ordenclientepymes
	1, // 3: chat.ChatService.RedecirOrdenRetail:output_type -> chat.Ordenclienteretail
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ordenclientepymes); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ordenclienteretail); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	RedecirOrdenPymes(ctx context.Context, in *Ordenclientepymes, opts ...grpc.CallOption) (*Ordenclientepymes, error)
	RedecirOrdenRetail(ctx context.Context, in *Ordenclienteretail, opts ...grpc.CallOption) (*Ordenclienteretail, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) RedecirOrdenPymes(ctx context.Context, in *Ordenclientepymes, opts ...grpc.CallOption) (*Ordenclientepymes, error) {
	out := new(Ordenclientepymes)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RedecirOrdenPymes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RedecirOrdenRetail(ctx context.Context, in *Ordenclienteretail, opts ...grpc.CallOption) (*Ordenclienteretail, error) {
	out := new(Ordenclienteretail)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RedecirOrdenRetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	RedecirOrdenPymes(context.Context, *Ordenclientepymes) (*Ordenclientepymes, error)
	RedecirOrdenRetail(context.Context, *Ordenclienteretail) (*Ordenclienteretail, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) RedecirOrdenPymes(context.Context, *Ordenclientepymes) (*Ordenclientepymes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedecirOrdenPymes not implemented")
}
func (*UnimplementedChatServiceServer) RedecirOrdenRetail(context.Context, *Ordenclienteretail) (*Ordenclienteretail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedecirOrdenRetail not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_RedecirOrdenPymes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ordenclientepymes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RedecirOrdenPymes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RedecirOrdenPymes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RedecirOrdenPymes(ctx, req.(*Ordenclientepymes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RedecirOrdenRetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ordenclienteretail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RedecirOrdenRetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RedecirOrdenRetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RedecirOrdenRetail(ctx, req.(*Ordenclienteretail))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RedecirOrdenPymes",
			Handler:    _ChatService_RedecirOrdenPymes_Handler,
		},
		{
			MethodName: "RedecirOrdenRetail",
			Handler:    _ChatService_RedecirOrdenRetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
