// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: proto/operation/operation.proto

package operation

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Owner_Entity int32

const (
	Owner_UNKNOWN    Owner_Entity = 0
	Owner_INDIVIDUAL Owner_Entity = 1
	Owner_LEGAL      Owner_Entity = 2
)

// Enum value maps for Owner_Entity.
var (
	Owner_Entity_name = map[int32]string{
		0: "UNKNOWN",
		1: "INDIVIDUAL",
		2: "LEGAL",
	}
	Owner_Entity_value = map[string]int32{
		"UNKNOWN":    0,
		"INDIVIDUAL": 1,
		"LEGAL":      2,
	}
)

func (x Owner_Entity) Enum() *Owner_Entity {
	p := new(Owner_Entity)
	*p = x
	return p
}

func (x Owner_Entity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Owner_Entity) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_operation_operation_proto_enumTypes[0].Descriptor()
}

func (Owner_Entity) Type() protoreflect.EnumType {
	return &file_proto_operation_operation_proto_enumTypes[0]
}

func (x Owner_Entity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Owner_Entity.Descriptor instead.
func (Owner_Entity) EnumDescriptor() ([]byte, []int) {
	return file_proto_operation_operation_proto_rawDescGZIP(), []int{1, 0}
}

type Department struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Department) Reset() {
	*x = Department{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_operation_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Department) ProtoMessage() {}

func (x *Department) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operation_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Department.ProtoReflect.Descriptor instead.
func (*Department) Descriptor() ([]byte, []int) {
	return file_proto_operation_operation_proto_rawDescGZIP(), []int{0}
}

func (x *Department) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Department) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity       Owner_Entity     `protobuf:"varint,1,opt,name=entity,proto3,enum=operation.Owner_Entity" json:"entity,omitempty"`
	Registration *Owner_Territory `protobuf:"bytes,2,opt,name=registration,proto3" json:"registration,omitempty"`
}

func (x *Owner) Reset() {
	*x = Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_operation_operation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operation_operation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_proto_operation_operation_proto_rawDescGZIP(), []int{1}
}

func (x *Owner) GetEntity() Owner_Entity {
	if x != nil {
		return x.Entity
	}
	return Owner_UNKNOWN
}

func (x *Owner) GetRegistration() *Owner_Territory {
	if x != nil {
		return x.Registration
	}
	return nil
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number      string       `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Brand       string       `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Model       string       `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Year        int32        `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
	Capacity    int32        `protobuf:"varint,5,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Color       string       `protobuf:"bytes,6,opt,name=color,proto3" json:"color,omitempty"`
	Fuel        string       `protobuf:"bytes,7,opt,name=fuel,proto3" json:"fuel,omitempty"`
	Kind        string       `protobuf:"bytes,8,opt,name=kind,proto3" json:"kind,omitempty"`
	Body        string       `protobuf:"bytes,9,opt,name=body,proto3" json:"body,omitempty"`
	Purpose     string       `protobuf:"bytes,10,opt,name=purpose,proto3" json:"purpose,omitempty"`
	OwnWeight   int32        `protobuf:"varint,11,opt,name=own_weight,json=ownWeight,proto3" json:"own_weight,omitempty"`
	TotalWeight int32        `protobuf:"varint,12,opt,name=total_weight,json=totalWeight,proto3" json:"total_weight,omitempty"`
	Date        *common.Date `protobuf:"bytes,13,opt,name=date,proto3" json:"date,omitempty"`
	Department  *Department  `protobuf:"bytes,14,opt,name=department,proto3" json:"department,omitempty"`
	Owner       *Owner       `protobuf:"bytes,15,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_operation_operation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operation_operation_proto_msgTypes[2]
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
	return file_proto_operation_operation_proto_rawDescGZIP(), []int{2}
}

func (x *Record) GetNumber() string {
	if x != nil {
		return x.Number
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

func (x *Record) GetPurpose() string {
	if x != nil {
		return x.Purpose
	}
	return ""
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

func (x *Record) GetDepartment() *Department {
	if x != nil {
		return x.Department
	}
	return nil
}

func (x *Record) GetOwner() *Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

type Owner_Territory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Owner_Territory) Reset() {
	*x = Owner_Territory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_operation_operation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner_Territory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner_Territory) ProtoMessage() {}

func (x *Owner_Territory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_operation_operation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner_Territory.ProtoReflect.Descriptor instead.
func (*Owner_Territory) Descriptor() ([]byte, []int) {
	return file_proto_operation_operation_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Owner_Territory) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_proto_operation_operation_proto protoreflect.FileDescriptor

var file_proto_operation_operation_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x05,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x54,
	0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x0a, 0x09, 0x54, 0x65, 0x72, 0x72, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x30, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x49, 0x4e, 0x44, 0x49, 0x56, 0x49, 0x44, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x10, 0x02, 0x22, 0xab, 0x03, 0x0a, 0x06, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x75, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x75,
	0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75,
	0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75, 0x72,
	0x70, 0x6f, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x57, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x26, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x61, 0x72, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_operation_operation_proto_rawDescOnce sync.Once
	file_proto_operation_operation_proto_rawDescData = file_proto_operation_operation_proto_rawDesc
)

func file_proto_operation_operation_proto_rawDescGZIP() []byte {
	file_proto_operation_operation_proto_rawDescOnce.Do(func() {
		file_proto_operation_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_operation_operation_proto_rawDescData)
	})
	return file_proto_operation_operation_proto_rawDescData
}

var file_proto_operation_operation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_operation_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_operation_operation_proto_goTypes = []interface{}{
	(Owner_Entity)(0),       // 0: operation.Owner.Entity
	(*Department)(nil),      // 1: operation.Department
	(*Owner)(nil),           // 2: operation.Owner
	(*Record)(nil),          // 3: operation.Record
	(*Owner_Territory)(nil), // 4: operation.Owner.Territory
	(*common.Date)(nil),     // 5: common.Date
}
var file_proto_operation_operation_proto_depIdxs = []int32{
	0, // 0: operation.Owner.entity:type_name -> operation.Owner.Entity
	4, // 1: operation.Owner.registration:type_name -> operation.Owner.Territory
	5, // 2: operation.Record.date:type_name -> common.Date
	1, // 3: operation.Record.department:type_name -> operation.Department
	2, // 4: operation.Record.owner:type_name -> operation.Owner
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_operation_operation_proto_init() }
func file_proto_operation_operation_proto_init() {
	if File_proto_operation_operation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_operation_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Department); i {
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
		file_proto_operation_operation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner); i {
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
		file_proto_operation_operation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_operation_operation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner_Territory); i {
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
			RawDescriptor: file_proto_operation_operation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_operation_operation_proto_goTypes,
		DependencyIndexes: file_proto_operation_operation_proto_depIdxs,
		EnumInfos:         file_proto_operation_operation_proto_enumTypes,
		MessageInfos:      file_proto_operation_operation_proto_msgTypes,
	}.Build()
	File_proto_operation_operation_proto = out.File
	file_proto_operation_operation_proto_rawDesc = nil
	file_proto_operation_operation_proto_goTypes = nil
	file_proto_operation_operation_proto_depIdxs = nil
}
