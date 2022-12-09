// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: proto/core/action.proto

package core

import (
	common "github.com/opencars/grpc/pkg/common"
	operation "github.com/opencars/grpc/pkg/operation"
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

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number      string       `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Vin         string       `protobuf:"bytes,2,opt,name=vin,proto3" json:"vin,omitempty"`
	Brand       string       `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	Model       string       `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Year        int32        `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
	Capacity    int32        `protobuf:"varint,6,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Color       string       `protobuf:"bytes,7,opt,name=color,proto3" json:"color,omitempty"`
	Fuel        string       `protobuf:"bytes,8,opt,name=fuel,proto3" json:"fuel,omitempty"`
	Kind        string       `protobuf:"bytes,9,opt,name=kind,proto3" json:"kind,omitempty"`
	Date        *common.Date `protobuf:"bytes,10,opt,name=date,proto3" json:"date,omitempty"`
	OwnWeight   int32        `protobuf:"varint,11,opt,name=own_weight,json=ownWeight,proto3" json:"own_weight,omitempty"`
	TotalWeight int32        `protobuf:"varint,12,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty"`
	// Operation:
	Body       string                  `protobuf:"bytes,13,opt,name=body,proto3" json:"body,omitempty"`
	Purpose    string                  `protobuf:"bytes,14,opt,name=purpose,proto3" json:"purpose,omitempty"`
	Action     *operation.RecordAction `protobuf:"bytes,15,opt,name=action,proto3" json:"action,omitempty"`
	Department *operation.Department   `protobuf:"bytes,16,opt,name=department,proto3" json:"department,omitempty"`
	Owner      *operation.Owner        `protobuf:"bytes,17,opt,name=owner,proto3" json:"owner,omitempty"`
	// Registration:
	Code         string       `protobuf:"bytes,18,opt,name=code,proto3" json:"code,omitempty"`
	NumSeating   int32        `protobuf:"varint,19,opt,name=num_seating,json=numSeating,proto3" json:"num_seating,omitempty"`
	FirstRegDate *common.Date `protobuf:"bytes,20,opt,name=first_reg_date,json=firstRegDate,proto3" json:"first_reg_date,omitempty"`
	Category     string       `protobuf:"bytes,21,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_core_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_proto_core_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_proto_core_action_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Action) GetVin() string {
	if x != nil {
		return x.Vin
	}
	return ""
}

func (x *Action) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Action) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Action) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Action) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Action) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Action) GetFuel() string {
	if x != nil {
		return x.Fuel
	}
	return ""
}

func (x *Action) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Action) GetDate() *common.Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Action) GetOwnWeight() int32 {
	if x != nil {
		return x.OwnWeight
	}
	return 0
}

func (x *Action) GetTotalWeight() int32 {
	if x != nil {
		return x.TotalWeight
	}
	return 0
}

func (x *Action) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Action) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
}

func (x *Action) GetAction() *operation.RecordAction {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *Action) GetDepartment() *operation.Department {
	if x != nil {
		return x.Department
	}
	return nil
}

func (x *Action) GetOwner() *operation.Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Action) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Action) GetNumSeating() int32 {
	if x != nil {
		return x.NumSeating
	}
	return 0
}

func (x *Action) GetFirstRegDate() *common.Date {
	if x != nil {
		return x.FirstRegDate
	}
	return nil
}

func (x *Action) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_proto_core_action_proto protoreflect.FileDescriptor

var file_proto_core_action_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x1a,
	0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x04, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x69, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x75, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x75, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77,
	0x6e, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6f, 0x77, 0x6e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0a, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x32, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x52, 0x65, 0x67, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42,
	0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x63, 0x61, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_core_action_proto_rawDescOnce sync.Once
	file_proto_core_action_proto_rawDescData = file_proto_core_action_proto_rawDesc
)

func file_proto_core_action_proto_rawDescGZIP() []byte {
	file_proto_core_action_proto_rawDescOnce.Do(func() {
		file_proto_core_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_core_action_proto_rawDescData)
	})
	return file_proto_core_action_proto_rawDescData
}

var file_proto_core_action_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_core_action_proto_goTypes = []interface{}{
	(*Action)(nil),                 // 0: core.Action
	(*common.Date)(nil),            // 1: common.Date
	(*operation.RecordAction)(nil), // 2: operation.RecordAction
	(*operation.Department)(nil),   // 3: operation.Department
	(*operation.Owner)(nil),        // 4: operation.Owner
}
var file_proto_core_action_proto_depIdxs = []int32{
	1, // 0: core.Action.date:type_name -> common.Date
	2, // 1: core.Action.action:type_name -> operation.RecordAction
	3, // 2: core.Action.department:type_name -> operation.Department
	4, // 3: core.Action.owner:type_name -> operation.Owner
	1, // 4: core.Action.first_reg_date:type_name -> common.Date
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_core_action_proto_init() }
func file_proto_core_action_proto_init() {
	if File_proto_core_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_core_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
			RawDescriptor: file_proto_core_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_core_action_proto_goTypes,
		DependencyIndexes: file_proto_core_action_proto_depIdxs,
		MessageInfos:      file_proto_core_action_proto_msgTypes,
	}.Build()
	File_proto_core_action_proto = out.File
	file_proto_core_action_proto_rawDesc = nil
	file_proto_core_action_proto_goTypes = nil
	file_proto_core_action_proto_depIdxs = nil
}