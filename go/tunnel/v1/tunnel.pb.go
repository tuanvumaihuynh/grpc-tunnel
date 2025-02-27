// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: tunnel/v1/tunnel.proto

package tunnelv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TunnelRequest represents a client request to send data through the tunnel.
type TunnelRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TunnelRequest) Reset() {
	*x = TunnelRequest{}
	mi := &file_tunnel_v1_tunnel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TunnelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelRequest) ProtoMessage() {}

func (x *TunnelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_v1_tunnel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelRequest.ProtoReflect.Descriptor instead.
func (*TunnelRequest) Descriptor() ([]byte, []int) {
	return file_tunnel_v1_tunnel_proto_rawDescGZIP(), []int{0}
}

func (x *TunnelRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// TunnelResponse represents the server's response containing data sent back through the tunnel.
type TunnelResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TunnelResponse) Reset() {
	*x = TunnelResponse{}
	mi := &file_tunnel_v1_tunnel_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TunnelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelResponse) ProtoMessage() {}

func (x *TunnelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_v1_tunnel_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelResponse.ProtoReflect.Descriptor instead.
func (*TunnelResponse) Descriptor() ([]byte, []int) {
	return file_tunnel_v1_tunnel_proto_rawDescGZIP(), []int{1}
}

func (x *TunnelResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_tunnel_v1_tunnel_proto protoreflect.FileDescriptor

var file_tunnel_v1_tunnel_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x22, 0x23, 0x0a, 0x0d, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x0e, 0x54, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x54,
	0x0a, 0x0d, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x43, 0x0a, 0x06, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x2e, 0x74, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x9e, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x75, 0x61, 0x6e, 0x76, 0x75, 0x6d, 0x61, 0x69, 0x68, 0x75, 0x79, 0x6e, 0x68,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x67, 0x6f, 0x2f,
	0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x54, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x15, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_tunnel_v1_tunnel_proto_rawDescOnce sync.Once
	file_tunnel_v1_tunnel_proto_rawDescData []byte
)

func file_tunnel_v1_tunnel_proto_rawDescGZIP() []byte {
	file_tunnel_v1_tunnel_proto_rawDescOnce.Do(func() {
		file_tunnel_v1_tunnel_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_tunnel_v1_tunnel_proto_rawDesc), len(file_tunnel_v1_tunnel_proto_rawDesc)))
	})
	return file_tunnel_v1_tunnel_proto_rawDescData
}

var file_tunnel_v1_tunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tunnel_v1_tunnel_proto_goTypes = []any{
	(*TunnelRequest)(nil),  // 0: tunnel.v1.TunnelRequest
	(*TunnelResponse)(nil), // 1: tunnel.v1.TunnelResponse
}
var file_tunnel_v1_tunnel_proto_depIdxs = []int32{
	0, // 0: tunnel.v1.TunnelService.Tunnel:input_type -> tunnel.v1.TunnelRequest
	1, // 1: tunnel.v1.TunnelService.Tunnel:output_type -> tunnel.v1.TunnelResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tunnel_v1_tunnel_proto_init() }
func file_tunnel_v1_tunnel_proto_init() {
	if File_tunnel_v1_tunnel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_tunnel_v1_tunnel_proto_rawDesc), len(file_tunnel_v1_tunnel_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tunnel_v1_tunnel_proto_goTypes,
		DependencyIndexes: file_tunnel_v1_tunnel_proto_depIdxs,
		MessageInfos:      file_tunnel_v1_tunnel_proto_msgTypes,
	}.Build()
	File_tunnel_v1_tunnel_proto = out.File
	file_tunnel_v1_tunnel_proto_goTypes = nil
	file_tunnel_v1_tunnel_proto_depIdxs = nil
}
