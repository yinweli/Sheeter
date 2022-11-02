// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: file2Data.proto

package sheeterProto

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

type File2Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reward *Reward `protobuf:"bytes,1,opt,name=Reward,proto3,oneof" json:"Reward,omitempty"`  //
	Enable *bool   `protobuf:"varint,2,opt,name=Enable,proto3,oneof" json:"Enable,omitempty"` // 是否啟用
	Ignore *int64  `protobuf:"varint,3,opt,name=Ignore,proto3,oneof" json:"Ignore,omitempty"` // 忽略
	Key    *int64  `protobuf:"varint,4,opt,name=Key,proto3,oneof" json:"Key,omitempty"`       // 索引
	Name   *string `protobuf:"bytes,5,opt,name=Name,proto3,oneof" json:"Name,omitempty"`      // 名稱
}

func (x *File2Data) Reset() {
	*x = File2Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file2Data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File2Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File2Data) ProtoMessage() {}

func (x *File2Data) ProtoReflect() protoreflect.Message {
	mi := &file_file2Data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File2Data.ProtoReflect.Descriptor instead.
func (*File2Data) Descriptor() ([]byte, []int) {
	return file_file2Data_proto_rawDescGZIP(), []int{0}
}

func (x *File2Data) GetReward() *Reward {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *File2Data) GetEnable() bool {
	if x != nil && x.Enable != nil {
		return *x.Enable
	}
	return false
}

func (x *File2Data) GetIgnore() int64 {
	if x != nil && x.Ignore != nil {
		return *x.Ignore
	}
	return 0
}

func (x *File2Data) GetKey() int64 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *File2Data) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type File2DataStorer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas map[int64]*File2Data `protobuf:"bytes,1,rep,name=Datas,proto3" json:"Datas,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *File2DataStorer) Reset() {
	*x = File2DataStorer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file2Data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File2DataStorer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File2DataStorer) ProtoMessage() {}

func (x *File2DataStorer) ProtoReflect() protoreflect.Message {
	mi := &file_file2Data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File2DataStorer.ProtoReflect.Descriptor instead.
func (*File2DataStorer) Descriptor() ([]byte, []int) {
	return file_file2Data_proto_rawDescGZIP(), []int{1}
}

func (x *File2DataStorer) GetDatas() map[int64]*File2Data {
	if x != nil {
		return x.Datas
	}
	return nil
}

var File_file2Data_proto protoreflect.FileDescriptor

var file_file2Data_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x32, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x01,
	0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x32, 0x44, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x06, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x48, 0x00, 0x52, 0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01,
	0x52, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x49,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x06, 0x49,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x03, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x4b, 0x65,
	0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x0f, 0x46,
	0x69, 0x6c, 0x65, 0x32, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x72, 0x12, 0x3e,
	0x0a, 0x05, 0x44, 0x61, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x32, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x44, 0x61, 0x74, 0x61, 0x73, 0x1a, 0x51,
	0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x32, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file2Data_proto_rawDescOnce sync.Once
	file_file2Data_proto_rawDescData = file_file2Data_proto_rawDesc
)

func file_file2Data_proto_rawDescGZIP() []byte {
	file_file2Data_proto_rawDescOnce.Do(func() {
		file_file2Data_proto_rawDescData = protoimpl.X.CompressGZIP(file_file2Data_proto_rawDescData)
	})
	return file_file2Data_proto_rawDescData
}

var file_file2Data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_file2Data_proto_goTypes = []interface{}{
	(*File2Data)(nil),       // 0: sheeterProto.File2Data
	(*File2DataStorer)(nil), // 1: sheeterProto.File2DataStorer
	nil,                     // 2: sheeterProto.File2DataStorer.DatasEntry
	(*Reward)(nil),          // 3: sheeterProto.Reward
}
var file_file2Data_proto_depIdxs = []int32{
	3, // 0: sheeterProto.File2Data.Reward:type_name -> sheeterProto.Reward
	2, // 1: sheeterProto.File2DataStorer.Datas:type_name -> sheeterProto.File2DataStorer.DatasEntry
	0, // 2: sheeterProto.File2DataStorer.DatasEntry.value:type_name -> sheeterProto.File2Data
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_file2Data_proto_init() }
func file_file2Data_proto_init() {
	if File_file2Data_proto != nil {
		return
	}
	file_reward_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_file2Data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File2Data); i {
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
		file_file2Data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File2DataStorer); i {
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
	file_file2Data_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file2Data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file2Data_proto_goTypes,
		DependencyIndexes: file_file2Data_proto_depIdxs,
		MessageInfos:      file_file2Data_proto_msgTypes,
	}.Build()
	File_file2Data_proto = out.File
	file_file2Data_proto_rawDesc = nil
	file_file2Data_proto_goTypes = nil
	file_file2Data_proto_depIdxs = nil
}