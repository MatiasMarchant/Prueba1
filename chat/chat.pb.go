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

type Ordenseguimiento struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nordenseguimiento string `protobuf:"bytes,1,opt,name=Nordenseguimiento,proto3" json:"Nordenseguimiento,omitempty"`
}

func (x *Ordenseguimiento) Reset() {
	*x = Ordenseguimiento{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ordenseguimiento) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ordenseguimiento) ProtoMessage() {}

func (x *Ordenseguimiento) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ordenseguimiento.ProtoReflect.Descriptor instead.
func (*Ordenseguimiento) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Ordenseguimiento) GetNordenseguimiento() string {
	if x != nil {
		return x.Nordenseguimiento
	}
	return ""
}

type Estado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Estado string `protobuf:"bytes,1,opt,name=Estado,proto3" json:"Estado,omitempty"`
}

func (x *Estado) Reset() {
	*x = Estado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Estado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Estado) ProtoMessage() {}

func (x *Estado) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Estado.ProtoReflect.Descriptor instead.
func (*Estado) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *Estado) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type ColaPaquete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idpaquete   string `protobuf:"bytes,1,opt,name=idpaquete,proto3" json:"idpaquete,omitempty"`
	Seguimiento string `protobuf:"bytes,2,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	Tipo        string `protobuf:"bytes,3,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor       string `protobuf:"bytes,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Intentos    string `protobuf:"bytes,5,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Estado      string `protobuf:"bytes,6,opt,name=estado,proto3" json:"estado,omitempty"`
	Origen      string `protobuf:"bytes,7,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino     string `protobuf:"bytes,8,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *ColaPaquete) Reset() {
	*x = ColaPaquete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColaPaquete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColaPaquete) ProtoMessage() {}

func (x *ColaPaquete) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColaPaquete.ProtoReflect.Descriptor instead.
func (*ColaPaquete) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ColaPaquete) GetIdpaquete() string {
	if x != nil {
		return x.Idpaquete
	}
	return ""
}

func (x *ColaPaquete) GetSeguimiento() string {
	if x != nil {
		return x.Seguimiento
	}
	return ""
}

func (x *ColaPaquete) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *ColaPaquete) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *ColaPaquete) GetIntentos() string {
	if x != nil {
		return x.Intentos
	}
	return ""
}

func (x *ColaPaquete) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *ColaPaquete) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *ColaPaquete) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type PaqueteEnviado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idpaquete   string `protobuf:"bytes,1,opt,name=idpaquete,proto3" json:"idpaquete,omitempty"`
	Seguimiento string `protobuf:"bytes,2,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	Tipo        string `protobuf:"bytes,3,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor       string `protobuf:"bytes,4,opt,name=valor,proto3" json:"valor,omitempty"`
	Intentos    string `protobuf:"bytes,5,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Estado      string `protobuf:"bytes,6,opt,name=estado,proto3" json:"estado,omitempty"`
	Origen      string `protobuf:"bytes,7,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino     string `protobuf:"bytes,8,opt,name=destino,proto3" json:"destino,omitempty"`
	Idcamion    string `protobuf:"bytes,9,opt,name=idcamion,proto3" json:"idcamion,omitempty"`
}

func (x *PaqueteEnviado) Reset() {
	*x = PaqueteEnviado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaqueteEnviado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaqueteEnviado) ProtoMessage() {}

func (x *PaqueteEnviado) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaqueteEnviado.ProtoReflect.Descriptor instead.
func (*PaqueteEnviado) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{5}
}

func (x *PaqueteEnviado) GetIdpaquete() string {
	if x != nil {
		return x.Idpaquete
	}
	return ""
}

func (x *PaqueteEnviado) GetSeguimiento() string {
	if x != nil {
		return x.Seguimiento
	}
	return ""
}

func (x *PaqueteEnviado) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *PaqueteEnviado) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *PaqueteEnviado) GetIntentos() string {
	if x != nil {
		return x.Intentos
	}
	return ""
}

func (x *PaqueteEnviado) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *PaqueteEnviado) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *PaqueteEnviado) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *PaqueteEnviado) GetIdcamion() string {
	if x != nil {
		return x.Idcamion
	}
	return ""
}

type IdCamion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idcamion string `protobuf:"bytes,1,opt,name=idcamion,proto3" json:"idcamion,omitempty"`
}

func (x *IdCamion) Reset() {
	*x = IdCamion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdCamion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdCamion) ProtoMessage() {}

func (x *IdCamion) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdCamion.ProtoReflect.Descriptor instead.
func (*IdCamion) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{6}
}

func (x *IdCamion) GetIdcamion() string {
	if x != nil {
		return x.Idcamion
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
	0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22, 0x40, 0x0a, 0x10, 0x4f, 0x72, 0x64,
	0x65, 0x6e, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x2c, 0x0a,
	0x11, 0x4e, 0x6f, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e,
	0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x4e, 0x6f, 0x72, 0x64, 0x65, 0x6e,
	0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x45,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0xdd, 0x01,
	0x0a, 0x0b, 0x43, 0x6f, 0x6c, 0x61, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x64, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x64, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22, 0xfc, 0x01,
	0x0a, 0x0e, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x64, 0x6f,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x69, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x08,
	0x49, 0x64, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x63, 0x61,
	0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x63, 0x61,
	0x6d, 0x69, 0x6f, 0x6e, 0x32, 0x8c, 0x04, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x25, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x14, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x76, 0x69,
	0x61, 0x64, 0x6f, 0x1a, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65,
	0x74, 0x65, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x64, 0x6f, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x1b, 0x45,
	0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x72, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x43, 0x61,
	0x6d, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x0e, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x49, 0x64, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x6f, 0x6c, 0x61, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x69, 0x62, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x50,
	0x79, 0x6d, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x70, 0x79, 0x6d, 0x65, 0x73, 0x1a, 0x16, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d,
	0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x11, 0x52, 0x65, 0x64, 0x65, 0x63,
	0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x50, 0x79, 0x6d, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65,
	0x70, 0x79, 0x6d, 0x65, 0x73, 0x1a, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x70, 0x79, 0x6d, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x69, 0x62, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e,
	0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x1a, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x67,
	0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x12, 0x52, 0x65,
	0x64, 0x65, 0x63, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x18, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x11, 0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f,
	0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x16, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65,
	0x6e, 0x74, 0x6f, 0x1a, 0x0c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x45, 0x73, 0x74, 0x61, 0x64,
	0x6f, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_chat_proto_goTypes = []interface{}{
	(*Ordenclientepymes)(nil),  // 0: chat.Ordenclientepymes
	(*Ordenclienteretail)(nil), // 1: chat.Ordenclienteretail
	(*Ordenseguimiento)(nil),   // 2: chat.Ordenseguimiento
	(*Estado)(nil),             // 3: chat.Estado
	(*ColaPaquete)(nil),        // 4: chat.ColaPaquete
	(*PaqueteEnviado)(nil),     // 5: chat.PaqueteEnviado
	(*IdCamion)(nil),           // 6: chat.IdCamion
}
var file_chat_proto_depIdxs = []int32{
	5, // 0: chat.ChatService.ActualizarRegistroPaqueteCamionNormal:input_type -> chat.PaqueteEnviado
	6, // 1: chat.ChatService.EntregarPaqueteCamionNormal:input_type -> chat.IdCamion
	0, // 2: chat.ChatService.RecibirOrdenPymes:input_type -> chat.Ordenclientepymes
	0, // 3: chat.ChatService.RedecirOrdenPymes:input_type -> chat.Ordenclientepymes
	1, // 4: chat.ChatService.RecibirOrdenRetail:input_type -> chat.Ordenclienteretail
	1, // 5: chat.ChatService.RedecirOrdenRetail:input_type -> chat.Ordenclienteretail
	2, // 6: chat.ChatService.CodigoSeguimiento:input_type -> chat.Ordenseguimiento
	5, // 7: chat.ChatService.ActualizarRegistroPaqueteCamionNormal:output_type -> chat.PaqueteEnviado
	4, // 8: chat.ChatService.EntregarPaqueteCamionNormal:output_type -> chat.ColaPaquete
	2, // 9: chat.ChatService.RecibirOrdenPymes:output_type -> chat.Ordenseguimiento
	0, // 10: chat.ChatService.RedecirOrdenPymes:output_type -> chat.Ordenclientepymes
	2, // 11: chat.ChatService.RecibirOrdenRetail:output_type -> chat.Ordenseguimiento
	1, // 12: chat.ChatService.RedecirOrdenRetail:output_type -> chat.Ordenclienteretail
	3, // 13: chat.ChatService.CodigoSeguimiento:output_type -> chat.Estado
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ordenseguimiento); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Estado); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColaPaquete); i {
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
		file_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaqueteEnviado); i {
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
		file_chat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdCamion); i {
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
			NumMessages:   7,
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
	ActualizarRegistroPaqueteCamionNormal(ctx context.Context, in *PaqueteEnviado, opts ...grpc.CallOption) (*PaqueteEnviado, error)
	EntregarPaqueteCamionNormal(ctx context.Context, in *IdCamion, opts ...grpc.CallOption) (*ColaPaquete, error)
	RecibirOrdenPymes(ctx context.Context, in *Ordenclientepymes, opts ...grpc.CallOption) (*Ordenseguimiento, error)
	RedecirOrdenPymes(ctx context.Context, in *Ordenclientepymes, opts ...grpc.CallOption) (*Ordenclientepymes, error)
	RecibirOrdenRetail(ctx context.Context, in *Ordenclienteretail, opts ...grpc.CallOption) (*Ordenseguimiento, error)
	RedecirOrdenRetail(ctx context.Context, in *Ordenclienteretail, opts ...grpc.CallOption) (*Ordenclienteretail, error)
	CodigoSeguimiento(ctx context.Context, in *Ordenseguimiento, opts ...grpc.CallOption) (*Estado, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) ActualizarRegistroPaqueteCamionNormal(ctx context.Context, in *PaqueteEnviado, opts ...grpc.CallOption) (*PaqueteEnviado, error) {
	out := new(PaqueteEnviado)
	err := c.cc.Invoke(ctx, "/chat.ChatService/ActualizarRegistroPaqueteCamionNormal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) EntregarPaqueteCamionNormal(ctx context.Context, in *IdCamion, opts ...grpc.CallOption) (*ColaPaquete, error) {
	out := new(ColaPaquete)
	err := c.cc.Invoke(ctx, "/chat.ChatService/EntregarPaqueteCamionNormal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RecibirOrdenPymes(ctx context.Context, in *Ordenclientepymes, opts ...grpc.CallOption) (*Ordenseguimiento, error) {
	out := new(Ordenseguimiento)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RecibirOrdenPymes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RedecirOrdenPymes(ctx context.Context, in *Ordenclientepymes, opts ...grpc.CallOption) (*Ordenclientepymes, error) {
	out := new(Ordenclientepymes)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RedecirOrdenPymes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) RecibirOrdenRetail(ctx context.Context, in *Ordenclienteretail, opts ...grpc.CallOption) (*Ordenseguimiento, error) {
	out := new(Ordenseguimiento)
	err := c.cc.Invoke(ctx, "/chat.ChatService/RecibirOrdenRetail", in, out, opts...)
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

func (c *chatServiceClient) CodigoSeguimiento(ctx context.Context, in *Ordenseguimiento, opts ...grpc.CallOption) (*Estado, error) {
	out := new(Estado)
	err := c.cc.Invoke(ctx, "/chat.ChatService/CodigoSeguimiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	ActualizarRegistroPaqueteCamionNormal(context.Context, *PaqueteEnviado) (*PaqueteEnviado, error)
	EntregarPaqueteCamionNormal(context.Context, *IdCamion) (*ColaPaquete, error)
	RecibirOrdenPymes(context.Context, *Ordenclientepymes) (*Ordenseguimiento, error)
	RedecirOrdenPymes(context.Context, *Ordenclientepymes) (*Ordenclientepymes, error)
	RecibirOrdenRetail(context.Context, *Ordenclienteretail) (*Ordenseguimiento, error)
	RedecirOrdenRetail(context.Context, *Ordenclienteretail) (*Ordenclienteretail, error)
	CodigoSeguimiento(context.Context, *Ordenseguimiento) (*Estado, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) ActualizarRegistroPaqueteCamionNormal(context.Context, *PaqueteEnviado) (*PaqueteEnviado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActualizarRegistroPaqueteCamionNormal not implemented")
}
func (*UnimplementedChatServiceServer) EntregarPaqueteCamionNormal(context.Context, *IdCamion) (*ColaPaquete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntregarPaqueteCamionNormal not implemented")
}
func (*UnimplementedChatServiceServer) RecibirOrdenPymes(context.Context, *Ordenclientepymes) (*Ordenseguimiento, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecibirOrdenPymes not implemented")
}
func (*UnimplementedChatServiceServer) RedecirOrdenPymes(context.Context, *Ordenclientepymes) (*Ordenclientepymes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedecirOrdenPymes not implemented")
}
func (*UnimplementedChatServiceServer) RecibirOrdenRetail(context.Context, *Ordenclienteretail) (*Ordenseguimiento, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecibirOrdenRetail not implemented")
}
func (*UnimplementedChatServiceServer) RedecirOrdenRetail(context.Context, *Ordenclienteretail) (*Ordenclienteretail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedecirOrdenRetail not implemented")
}
func (*UnimplementedChatServiceServer) CodigoSeguimiento(context.Context, *Ordenseguimiento) (*Estado, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CodigoSeguimiento not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_ActualizarRegistroPaqueteCamionNormal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaqueteEnviado)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).ActualizarRegistroPaqueteCamionNormal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/ActualizarRegistroPaqueteCamionNormal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).ActualizarRegistroPaqueteCamionNormal(ctx, req.(*PaqueteEnviado))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_EntregarPaqueteCamionNormal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdCamion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).EntregarPaqueteCamionNormal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/EntregarPaqueteCamionNormal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).EntregarPaqueteCamionNormal(ctx, req.(*IdCamion))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_RecibirOrdenPymes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ordenclientepymes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RecibirOrdenPymes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RecibirOrdenPymes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RecibirOrdenPymes(ctx, req.(*Ordenclientepymes))
	}
	return interceptor(ctx, in, info, handler)
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

func _ChatService_RecibirOrdenRetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ordenclienteretail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).RecibirOrdenRetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/RecibirOrdenRetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).RecibirOrdenRetail(ctx, req.(*Ordenclienteretail))
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

func _ChatService_CodigoSeguimiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ordenseguimiento)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CodigoSeguimiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/CodigoSeguimiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CodigoSeguimiento(ctx, req.(*Ordenseguimiento))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActualizarRegistroPaqueteCamionNormal",
			Handler:    _ChatService_ActualizarRegistroPaqueteCamionNormal_Handler,
		},
		{
			MethodName: "EntregarPaqueteCamionNormal",
			Handler:    _ChatService_EntregarPaqueteCamionNormal_Handler,
		},
		{
			MethodName: "RecibirOrdenPymes",
			Handler:    _ChatService_RecibirOrdenPymes_Handler,
		},
		{
			MethodName: "RedecirOrdenPymes",
			Handler:    _ChatService_RedecirOrdenPymes_Handler,
		},
		{
			MethodName: "RecibirOrdenRetail",
			Handler:    _ChatService_RecibirOrdenRetail_Handler,
		},
		{
			MethodName: "RedecirOrdenRetail",
			Handler:    _ChatService_RedecirOrdenRetail_Handler,
		},
		{
			MethodName: "CodigoSeguimiento",
			Handler:    _ChatService_CodigoSeguimiento_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
