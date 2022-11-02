// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: actor.proto

package proto

import (
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

type DataFormat int32

const (
	DataFormat_UNKNOWN DataFormat = 0
	DataFormat_JSON    DataFormat = 1
)

// Enum value maps for DataFormat.
var (
	DataFormat_name = map[int32]string{
		0: "UNKNOWN",
		1: "JSON",
	}
	DataFormat_value = map[string]int32{
		"UNKNOWN": 0,
		"JSON":    1,
	}
)

func (x DataFormat) Enum() *DataFormat {
	p := new(DataFormat)
	*p = x
	return p
}

func (x DataFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_actor_proto_enumTypes[0].Descriptor()
}

func (DataFormat) Type() protoreflect.EnumType {
	return &file_actor_proto_enumTypes[0]
}

func (x DataFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataFormat.Descriptor instead.
func (DataFormat) EnumDescriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{0}
}

type PID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BinarySource string `protobuf:"bytes,3,opt,name=binary_source,json=binarySource,proto3" json:"binary_source,omitempty"`
	Host         string `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Endpoint     string `protobuf:"bytes,5,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *PID) Reset() {
	*x = PID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PID) ProtoMessage() {}

func (x *PID) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PID.ProtoReflect.Descriptor instead.
func (*PID) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{0}
}

func (x *PID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PID) GetBinarySource() string {
	if x != nil {
		return x.BinarySource
	}
	return ""
}

func (x *PID) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PID) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type RemoteMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body   []byte            `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Header map[string]string `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	From   *PID              `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To     *PID              `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Format DataFormat        `protobuf:"varint,5,opt,name=format,proto3,enum=proto.DataFormat" json:"format,omitempty"`
}

func (x *RemoteMessage) Reset() {
	*x = RemoteMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteMessage) ProtoMessage() {}

func (x *RemoteMessage) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteMessage.ProtoReflect.Descriptor instead.
func (*RemoteMessage) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteMessage) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *RemoteMessage) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RemoteMessage) GetFrom() *PID {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *RemoteMessage) GetTo() *PID {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *RemoteMessage) GetFormat() DataFormat {
	if x != nil {
		return x.Format
	}
	return DataFormat_UNKNOWN
}

var File_actor_proto protoreflect.FileDescriptor

var file_actor_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x22, 0xff, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x29, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x23, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01, 0x32, 0x3b, 0x0a, 0x05, 0x41,
	0x63, 0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x72, 0x63, 0x6f, 0x6e, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x6b, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6c, 0x65, 0x74, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_actor_proto_rawDescOnce sync.Once
	file_actor_proto_rawDescData = file_actor_proto_rawDesc
)

func file_actor_proto_rawDescGZIP() []byte {
	file_actor_proto_rawDescOnce.Do(func() {
		file_actor_proto_rawDescData = protoimpl.X.CompressGZIP(file_actor_proto_rawDescData)
	})
	return file_actor_proto_rawDescData
}

var file_actor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_actor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_actor_proto_goTypes = []interface{}{
	(DataFormat)(0),       // 0: proto.DataFormat
	(*PID)(nil),           // 1: proto.PID
	(*RemoteMessage)(nil), // 2: proto.RemoteMessage
	nil,                   // 3: proto.RemoteMessage.HeaderEntry
}
var file_actor_proto_depIdxs = []int32{
	3, // 0: proto.RemoteMessage.header:type_name -> proto.RemoteMessage.HeaderEntry
	1, // 1: proto.RemoteMessage.from:type_name -> proto.PID
	1, // 2: proto.RemoteMessage.to:type_name -> proto.PID
	0, // 3: proto.RemoteMessage.format:type_name -> proto.DataFormat
	2, // 4: proto.Actor.Call:input_type -> proto.RemoteMessage
	2, // 5: proto.Actor.Call:output_type -> proto.RemoteMessage
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_actor_proto_init() }
func file_actor_proto_init() {
	if File_actor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_actor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PID); i {
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
		file_actor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteMessage); i {
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
			RawDescriptor: file_actor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_actor_proto_goTypes,
		DependencyIndexes: file_actor_proto_depIdxs,
		EnumInfos:         file_actor_proto_enumTypes,
		MessageInfos:      file_actor_proto_msgTypes,
	}.Build()
	File_actor_proto = out.File
	file_actor_proto_rawDesc = nil
	file_actor_proto_goTypes = nil
	file_actor_proto_depIdxs = nil
}
