// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: proto/registration/registration.proto

package registration

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

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         string       `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Number       string       `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Vin          string       `protobuf:"bytes,3,opt,name=vin,proto3" json:"vin,omitempty"`
	Brand        string       `protobuf:"bytes,4,opt,name=brand,proto3" json:"brand,omitempty"`
	Model        string       `protobuf:"bytes,5,opt,name=model,proto3" json:"model,omitempty"`
	Year         int32        `protobuf:"varint,6,opt,name=year,proto3" json:"year,omitempty"`
	Capacity     int32        `protobuf:"varint,7,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Color        string       `protobuf:"bytes,8,opt,name=color,proto3" json:"color,omitempty"`
	Fuel         string       `protobuf:"bytes,9,opt,name=fuel,proto3" json:"fuel,omitempty"`
	Kind         string       `protobuf:"bytes,10,opt,name=kind,proto3" json:"kind,omitempty"`
	Body         string       `protobuf:"bytes,17,opt,name=body,proto3" json:"body,omitempty"`
	NumSeating   int32        `protobuf:"varint,11,opt,name=num_seating,json=numSeating,proto3" json:"num_seating,omitempty"`
	OwnWeight    int32        `protobuf:"varint,12,opt,name=own_weight,json=ownWeight,proto3" json:"own_weight,omitempty"`
	TotalWeight  int32        `protobuf:"varint,13,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty"`
	Date         *common.Date `protobuf:"bytes,14,opt,name=date,proto3" json:"date,omitempty"`
	FirstRegDate *common.Date `protobuf:"bytes,15,opt,name=first_reg_date,json=firstRegDate,proto3" json:"first_reg_date,omitempty"`
	Category     string       `protobuf:"bytes,16,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_registration_registration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_proto_registration_registration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_proto_registration_registration_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Record) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Record) GetVin() string {
	if x != nil {
		return x.Vin
	}
	return ""
}

func (x *Record) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Record) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Record) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Record) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Record) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Record) GetFuel() string {
	if x != nil {
		return x.Fuel
	}
	return ""
}

func (x *Record) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Record) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Record) GetNumSeating() int32 {
	if x != nil {
		return x.NumSeating
	}
	return 0
}

func (x *Record) GetOwnWeight() int32 {
	if x != nil {
		return x.OwnWeight
	}
	return 0
}

func (x *Record) GetTotalWeight() int32 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

func (x *Record) GetDate() *common.Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Record) GetFirstRegDate() *common.Date {
	if x != nil {
		return x.FirstRegDate
	}
	return nil
}

func (x *Record) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_proto_registration_registration_proto protoreflect.FileDescriptor

var file_proto_registration_registration_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9,
	0x03, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x76, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x75, 0x65,
	0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x75, 0x65, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x53,
	0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x5f, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x57,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x0e, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x67, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x61, 0x72,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_registration_registration_proto_rawDescOnce sync.Once
	file_proto_registration_registration_proto_rawDescData = file_proto_registration_registration_proto_rawDesc
)

func file_proto_registration_registration_proto_rawDescGZIP() []byte {
	file_proto_registration_registration_proto_rawDescOnce.Do(func() {
		file_proto_registration_registration_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_registration_registration_proto_rawDescData)
	})
	return file_proto_registration_registration_proto_rawDescData
}

var file_proto_registration_registration_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_registration_registration_proto_goTypes = []interface{}{
	(*Record)(nil),      // 0: registration.Record
	(*common.Date)(nil), // 1: common.Date
}
var file_proto_registration_registration_proto_depIdxs = []int32{
	1, // 0: registration.Record.date:type_name -> common.Date
	1, // 1: registration.Record.first_reg_date:type_name -> common.Date
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_registration_registration_proto_init() }
func file_proto_registration_registration_proto_init() {
	if File_proto_registration_registration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_registration_registration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
			RawDescriptor: file_proto_registration_registration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_registration_registration_proto_goTypes,
		DependencyIndexes: file_proto_registration_registration_proto_depIdxs,
		MessageInfos:      file_proto_registration_registration_proto_msgTypes,
	}.Build()
	File_proto_registration_registration_proto = out.File
	file_proto_registration_registration_proto_rawDesc = nil
	file_proto_registration_registration_proto_goTypes = nil
	file_proto_registration_registration_proto_depIdxs = nil
}
