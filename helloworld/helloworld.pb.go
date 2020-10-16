// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.5.1
// source: helloworld/helloworld.proto

package helloworld

import (
	context "context"
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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto    string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor       string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda      string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino     string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	Prioritario string `protobuf:"bytes,6,opt,name=prioritario,proto3" json:"prioritario,omitempty"`
	Estado      string `protobuf:"bytes,7,opt,name=estado,proto3" json:"estado,omitempty"`
	Idpaquete   string `protobuf:"bytes,8,opt,name=idpaquete,proto3" json:"idpaquete,omitempty"`
	Tipo        string `protobuf:"bytes,9,opt,name=tipo,proto3" json:"tipo,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *Message) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *Message) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *Message) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *Message) GetPrioritario() string {
	if x != nil {
		return x.Prioritario
	}
	return ""
}

func (x *Message) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *Message) GetIdpaquete() string {
	if x != nil {
		return x.Idpaquete
	}
	return ""
}

func (x *Message) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

type CodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CodeRequest) Reset() {
	*x = CodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeRequest) ProtoMessage() {}

func (x *CodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeRequest.ProtoReflect.Descriptor instead.
func (*CodeRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{1}
}

func (x *CodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type PaqueteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idpaquete   string `protobuf:"bytes,1,opt,name=idpaquete,proto3" json:"idpaquete,omitempty"`
	Idcamion    string `protobuf:"bytes,2,opt,name=idcamion,proto3" json:"idcamion,omitempty"`
	Seguimiento string `protobuf:"bytes,3,opt,name=seguimiento,proto3" json:"seguimiento,omitempty"`
	Tipo        string `protobuf:"bytes,4,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor       string `protobuf:"bytes,5,opt,name=valor,proto3" json:"valor,omitempty"`
	Intentos    int32  `protobuf:"varint,6,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Estado      string `protobuf:"bytes,7,opt,name=estado,proto3" json:"estado,omitempty"`
	Producto    string `protobuf:"bytes,8,opt,name=producto,proto3" json:"producto,omitempty"`
	Origen      string `protobuf:"bytes,9,opt,name=origen,proto3" json:"origen,omitempty"`
	Destino     string `protobuf:"bytes,10,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *PaqueteRequest) Reset() {
	*x = PaqueteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld_helloworld_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaqueteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaqueteRequest) ProtoMessage() {}

func (x *PaqueteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld_helloworld_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaqueteRequest.ProtoReflect.Descriptor instead.
func (*PaqueteRequest) Descriptor() ([]byte, []int) {
	return file_helloworld_helloworld_proto_rawDescGZIP(), []int{2}
}

func (x *PaqueteRequest) GetIdpaquete() string {
	if x != nil {
		return x.Idpaquete
	}
	return ""
}

func (x *PaqueteRequest) GetIdcamion() string {
	if x != nil {
		return x.Idcamion
	}
	return ""
}

func (x *PaqueteRequest) GetSeguimiento() string {
	if x != nil {
		return x.Seguimiento
	}
	return ""
}

func (x *PaqueteRequest) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *PaqueteRequest) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *PaqueteRequest) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *PaqueteRequest) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *PaqueteRequest) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *PaqueteRequest) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *PaqueteRequest) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

var File_helloworld_helloworld_proto protoreflect.FileDescriptor

var file_helloworld_helloworld_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0xe9, 0x01, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x69, 0x70, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x98, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x71,
	0x75, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x64, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x64, 0x70, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x63,
	0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x63,
	0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69,
	0x65, 0x6e, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x6f, 0x32, 0xd4, 0x01, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x06, 0x42, 0x75, 0x73, 0x63, 0x61, 0x72, 0x12, 0x17, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x49, 0x0a, 0x0d, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65,
	0x12, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50, 0x61,
	0x71, 0x75, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_helloworld_helloworld_proto_rawDescOnce sync.Once
	file_helloworld_helloworld_proto_rawDescData = file_helloworld_helloworld_proto_rawDesc
)

func file_helloworld_helloworld_proto_rawDescGZIP() []byte {
	file_helloworld_helloworld_proto_rawDescOnce.Do(func() {
		file_helloworld_helloworld_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld_helloworld_proto_rawDescData)
	})
	return file_helloworld_helloworld_proto_rawDescData
}

var file_helloworld_helloworld_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_helloworld_helloworld_proto_goTypes = []interface{}{
	(*Message)(nil),        // 0: helloworld.Message
	(*CodeRequest)(nil),    // 1: helloworld.CodeRequest
	(*PaqueteRequest)(nil), // 2: helloworld.PaqueteRequest
}
var file_helloworld_helloworld_proto_depIdxs = []int32{
	0, // 0: helloworld.HelloworldService.SayHello:input_type -> helloworld.Message
	1, // 1: helloworld.HelloworldService.Buscar:input_type -> helloworld.CodeRequest
	2, // 2: helloworld.HelloworldService.EnviarPaquete:input_type -> helloworld.PaqueteRequest
	0, // 3: helloworld.HelloworldService.SayHello:output_type -> helloworld.Message
	1, // 4: helloworld.HelloworldService.Buscar:output_type -> helloworld.CodeRequest
	2, // 5: helloworld.HelloworldService.EnviarPaquete:output_type -> helloworld.PaqueteRequest
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_helloworld_helloworld_proto_init() }
func file_helloworld_helloworld_proto_init() {
	if File_helloworld_helloworld_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld_helloworld_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_helloworld_helloworld_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeRequest); i {
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
		file_helloworld_helloworld_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaqueteRequest); i {
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
			RawDescriptor: file_helloworld_helloworld_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld_helloworld_proto_goTypes,
		DependencyIndexes: file_helloworld_helloworld_proto_depIdxs,
		MessageInfos:      file_helloworld_helloworld_proto_msgTypes,
	}.Build()
	File_helloworld_helloworld_proto = out.File
	file_helloworld_helloworld_proto_rawDesc = nil
	file_helloworld_helloworld_proto_goTypes = nil
	file_helloworld_helloworld_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloworldServiceClient is the client API for HelloworldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloworldServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	Buscar(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeRequest, error)
	EnviarPaquete(ctx context.Context, in *PaqueteRequest, opts ...grpc.CallOption) (*PaqueteRequest, error)
}

type helloworldServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloworldServiceClient(cc grpc.ClientConnInterface) HelloworldServiceClient {
	return &helloworldServiceClient{cc}
}

func (c *helloworldServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/helloworld.HelloworldService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldServiceClient) Buscar(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeRequest, error) {
	out := new(CodeRequest)
	err := c.cc.Invoke(ctx, "/helloworld.HelloworldService/Buscar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldServiceClient) EnviarPaquete(ctx context.Context, in *PaqueteRequest, opts ...grpc.CallOption) (*PaqueteRequest, error) {
	out := new(PaqueteRequest)
	err := c.cc.Invoke(ctx, "/helloworld.HelloworldService/EnviarPaquete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloworldServiceServer is the server API for HelloworldService service.
type HelloworldServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	Buscar(context.Context, *CodeRequest) (*CodeRequest, error)
	EnviarPaquete(context.Context, *PaqueteRequest) (*PaqueteRequest, error)
}

// UnimplementedHelloworldServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloworldServiceServer struct {
}

func (*UnimplementedHelloworldServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloworldServiceServer) Buscar(context.Context, *CodeRequest) (*CodeRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Buscar not implemented")
}
func (*UnimplementedHelloworldServiceServer) EnviarPaquete(context.Context, *PaqueteRequest) (*PaqueteRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarPaquete not implemented")
}

func RegisterHelloworldServiceServer(s *grpc.Server, srv HelloworldServiceServer) {
	s.RegisterService(&_HelloworldService_serviceDesc, srv)
}

func _HelloworldService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloworldService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloworldService_Buscar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServiceServer).Buscar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloworldService/Buscar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServiceServer).Buscar(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloworldService_EnviarPaquete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaqueteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServiceServer).EnviarPaquete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.HelloworldService/EnviarPaquete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServiceServer).EnviarPaquete(ctx, req.(*PaqueteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloworldService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.HelloworldService",
	HandlerType: (*HelloworldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloworldService_SayHello_Handler,
		},
		{
			MethodName: "Buscar",
			Handler:    _HelloworldService_Buscar_Handler,
		},
		{
			MethodName: "EnviarPaquete",
			Handler:    _HelloworldService_EnviarPaquete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}
