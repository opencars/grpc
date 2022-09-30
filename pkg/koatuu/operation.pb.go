// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: proto/koatuu/operation.proto

package koatuu

import (
	common "github.com/opencars/grpc/pkg/common"
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

type DecodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []string `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *DecodeRequest) Reset() {
	*x = DecodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_koatuu_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeRequest) ProtoMessage() {}

func (x *DecodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_koatuu_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeRequest.ProtoReflect.Descriptor instead.
func (*DecodeRequest) Descriptor() ([]byte, []int) {
	return file_proto_koatuu_operation_proto_rawDescGZIP(), []int{0}
}

func (x *DecodeRequest) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

type DecodeResultList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*DecodeResultItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *DecodeResultList) Reset() {
	*x = DecodeResultList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_koatuu_operation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeResultList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeResultList) ProtoMessage() {}

func (x *DecodeResultList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_koatuu_operation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeResultList.ProtoReflect.Descriptor instead.
func (*DecodeResultList) Descriptor() ([]byte, []int) {
	return file_proto_koatuu_operation_proto_rawDescGZIP(), []int{1}
}

func (x *DecodeResultList) GetItems() []*DecodeResultItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type DecodeResultItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Summary string        `protobuf:"bytes,1,opt,name=summary,proto3" json:"summary,omitempty"`
	Error   *common.Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DecodeResultItem) Reset() {
	*x = DecodeResultItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_koatuu_operation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeResultItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeResultItem) ProtoMessage() {}

func (x *DecodeResultItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_koatuu_operation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeResultItem.ProtoReflect.Descriptor instead.
func (*DecodeResultItem) Descriptor() ([]byte, []int) {
	return file_proto_koatuu_operation_proto_rawDescGZIP(), []int{2}
}

func (x *DecodeResultItem) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *DecodeResultItem) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_proto_koatuu_operation_proto protoreflect.FileDescriptor

var file_proto_koatuu_operation_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x6f, 0x61, 0x74, 0x75, 0x75, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6b, 0x6f, 0x61, 0x74, 0x75, 0x75, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x25, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x10, 0x44, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x6f, 0x61,
	0x74, 0x75, 0x75, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x51, 0x0a, 0x10, 0x44,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x44,
	0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x15, 0x2e, 0x6b, 0x6f, 0x61, 0x74, 0x75, 0x75, 0x2e, 0x44, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6b, 0x6f, 0x61,
	0x74, 0x75, 0x75, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x61, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6b, 0x6f, 0x61, 0x74, 0x75, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_koatuu_operation_proto_rawDescOnce sync.Once
	file_proto_koatuu_operation_proto_rawDescData = file_proto_koatuu_operation_proto_rawDesc
)

func file_proto_koatuu_operation_proto_rawDescGZIP() []byte {
	file_proto_koatuu_operation_proto_rawDescOnce.Do(func() {
		file_proto_koatuu_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_koatuu_operation_proto_rawDescData)
	})
	return file_proto_koatuu_operation_proto_rawDescData
}

var file_proto_koatuu_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_koatuu_operation_proto_goTypes = []interface{}{
	(*DecodeRequest)(nil),    // 0: koatuu.DecodeRequest
	(*DecodeResultList)(nil), // 1: koatuu.DecodeResultList
	(*DecodeResultItem)(nil), // 2: koatuu.DecodeResultItem
	(*common.Error)(nil),     // 3: common.Error
}
var file_proto_koatuu_operation_proto_depIdxs = []int32{
	2, // 0: koatuu.DecodeResultList.items:type_name -> koatuu.DecodeResultItem
	3, // 1: koatuu.DecodeResultItem.error:type_name -> common.Error
	0, // 2: koatuu.Service.Decode:input_type -> koatuu.DecodeRequest
	1, // 3: koatuu.Service.Decode:output_type -> koatuu.DecodeResultList
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_koatuu_operation_proto_init() }
func file_proto_koatuu_operation_proto_init() {
	if File_proto_koatuu_operation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_koatuu_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeRequest); i {
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
		file_proto_koatuu_operation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeResultList); i {
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
		file_proto_koatuu_operation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeResultItem); i {
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
			RawDescriptor: file_proto_koatuu_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_koatuu_operation_proto_goTypes,
		DependencyIndexes: file_proto_koatuu_operation_proto_depIdxs,
		MessageInfos:      file_proto_koatuu_operation_proto_msgTypes,
	}.Build()
	File_proto_koatuu_operation_proto = out.File
	file_proto_koatuu_operation_proto_rawDesc = nil
	file_proto_koatuu_operation_proto_goTypes = nil
	file_proto_koatuu_operation_proto_depIdxs = nil
}