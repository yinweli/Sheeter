// Code generated by sheeter. DO NOT EDIT.
// Sheeter: https://github.com/yinweli/Sheeter

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: benchmark114Data.proto

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

type Benchmark114Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reward *Reward `protobuf:"bytes,1,opt,name=Reward,proto3,oneof" json:"Reward,omitempty"`  //
	Enable *bool   `protobuf:"varint,2,opt,name=Enable,proto3,oneof" json:"Enable,omitempty"` // 是否啟用
	Key    *int64  `protobuf:"varint,3,opt,name=Key,proto3,oneof" json:"Key,omitempty"`       // 索引
	Name   *string `protobuf:"bytes,4,opt,name=Name,proto3,oneof" json:"Name,omitempty"`      // 名稱
}

func (x *Benchmark114Data) Reset() {
	*x = Benchmark114Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark114Data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Benchmark114Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Benchmark114Data) ProtoMessage() {}

func (x *Benchmark114Data) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark114Data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Benchmark114Data.ProtoReflect.Descriptor instead.
func (*Benchmark114Data) Descriptor() ([]byte, []int) {
	return file_benchmark114Data_proto_rawDescGZIP(), []int{0}
}

func (x *Benchmark114Data) GetReward() *Reward {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *Benchmark114Data) GetEnable() bool {
	if x != nil && x.Enable != nil {
		return *x.Enable
	}
	return false
}

func (x *Benchmark114Data) GetKey() int64 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *Benchmark114Data) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type Benchmark114DataStorer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas map[int64]*Benchmark114Data `protobuf:"bytes,1,rep,name=Datas,proto3" json:"Datas,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Benchmark114DataStorer) Reset() {
	*x = Benchmark114DataStorer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark114Data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Benchmark114DataStorer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Benchmark114DataStorer) ProtoMessage() {}

func (x *Benchmark114DataStorer) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark114Data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Benchmark114DataStorer.ProtoReflect.Descriptor instead.
func (*Benchmark114DataStorer) Descriptor() ([]byte, []int) {
	return file_benchmark114Data_proto_rawDescGZIP(), []int{1}
}

func (x *Benchmark114DataStorer) GetDatas() map[int64]*Benchmark114Data {
	if x != nil {
		return x.Datas
	}
	return nil
}

var File_benchmark114Data_proto protoreflect.FileDescriptor

var file_benchmark114Data_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x31, 0x31, 0x34, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x10, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61,
	0x72, 0x6b, 0x31, 0x31, 0x34, 0x44, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x06, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x48,
	0x00, 0x52, 0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x06,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x4b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x4b, 0x65, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xb9, 0x01, 0x0a, 0x16, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x31, 0x31,
	0x34, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x05, 0x44,
	0x61, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x68, 0x65,
	0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d,
	0x61, 0x72, 0x6b, 0x31, 0x31, 0x34, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x72,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x44, 0x61, 0x74,
	0x61, 0x73, 0x1a, 0x58, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x31, 0x31, 0x34, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10, 0x5a, 0x0e,
	0x2e, 0x3b, 0x73, 0x68, 0x65, 0x65, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_benchmark114Data_proto_rawDescOnce sync.Once
	file_benchmark114Data_proto_rawDescData = file_benchmark114Data_proto_rawDesc
)

func file_benchmark114Data_proto_rawDescGZIP() []byte {
	file_benchmark114Data_proto_rawDescOnce.Do(func() {
		file_benchmark114Data_proto_rawDescData = protoimpl.X.CompressGZIP(file_benchmark114Data_proto_rawDescData)
	})
	return file_benchmark114Data_proto_rawDescData
}

var file_benchmark114Data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_benchmark114Data_proto_goTypes = []interface{}{
	(*Benchmark114Data)(nil),       // 0: sheeterProto.Benchmark114Data
	(*Benchmark114DataStorer)(nil), // 1: sheeterProto.Benchmark114DataStorer
	nil,                            // 2: sheeterProto.Benchmark114DataStorer.DatasEntry
	(*Reward)(nil),                 // 3: sheeterProto.Reward
}
var file_benchmark114Data_proto_depIdxs = []int32{
	3, // 0: sheeterProto.Benchmark114Data.Reward:type_name -> sheeterProto.Reward
	2, // 1: sheeterProto.Benchmark114DataStorer.Datas:type_name -> sheeterProto.Benchmark114DataStorer.DatasEntry
	0, // 2: sheeterProto.Benchmark114DataStorer.DatasEntry.value:type_name -> sheeterProto.Benchmark114Data
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_benchmark114Data_proto_init() }
func file_benchmark114Data_proto_init() {
	if File_benchmark114Data_proto != nil {
		return
	}
	file_reward_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_benchmark114Data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Benchmark114Data); i {
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
		file_benchmark114Data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Benchmark114DataStorer); i {
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
	file_benchmark114Data_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_benchmark114Data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_benchmark114Data_proto_goTypes,
		DependencyIndexes: file_benchmark114Data_proto_depIdxs,
		MessageInfos:      file_benchmark114Data_proto_msgTypes,
	}.Build()
	File_benchmark114Data_proto = out.File
	file_benchmark114Data_proto_rawDesc = nil
	file_benchmark114Data_proto_goTypes = nil
	file_benchmark114Data_proto_depIdxs = nil
}
