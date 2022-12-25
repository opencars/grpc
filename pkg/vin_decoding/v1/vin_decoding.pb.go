// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: vin_decoding/v1/vin_decoding.proto

package vin_decodingv1

import (
	v1 "github.com/opencars/grpc/pkg/common/v1"
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

	Vins []string `protobuf:"bytes,1,rep,name=vins,proto3" json:"vins,omitempty"`
}

func (x *DecodeRequest) Reset() {
	*x = DecodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_decoding_v1_vin_decoding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeRequest) ProtoMessage() {}

func (x *DecodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vin_decoding_v1_vin_decoding_proto_msgTypes[0]
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
	return file_vin_decoding_v1_vin_decoding_proto_rawDescGZIP(), []int{0}
}

func (x *DecodeRequest) GetVins() []string {
	if x != nil {
		return x.Vins
	}
	return nil
}

type DecodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*DecodeResultItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *DecodeResponse) Reset() {
	*x = DecodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_decoding_v1_vin_decoding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeResponse) ProtoMessage() {}

func (x *DecodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vin_decoding_v1_vin_decoding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodeResponse.ProtoReflect.Descriptor instead.
func (*DecodeResponse) Descriptor() ([]byte, []int) {
	return file_vin_decoding_v1_vin_decoding_proto_rawDescGZIP(), []int{1}
}

func (x *DecodeResponse) GetItems() []*DecodeResultItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type DecodeResultItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vehicle    *Vehicle    `protobuf:"bytes,1,opt,name=vehicle,proto3" json:"vehicle,omitempty"`
	DecodedVin *DecodedVIN `protobuf:"bytes,2,opt,name=decoded_vin,json=decodedVin,proto3" json:"decoded_vin,omitempty"`
	Error      *v1.Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DecodeResultItem) Reset() {
	*x = DecodeResultItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_decoding_v1_vin_decoding_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodeResultItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodeResultItem) ProtoMessage() {}

func (x *DecodeResultItem) ProtoReflect() protoreflect.Message {
	mi := &file_vin_decoding_v1_vin_decoding_proto_msgTypes[2]
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
	return file_vin_decoding_v1_vin_decoding_proto_rawDescGZIP(), []int{2}
}

func (x *DecodeResultItem) GetVehicle() *Vehicle {
	if x != nil {
		return x.Vehicle
	}
	return nil
}

func (x *DecodeResultItem) GetDecodedVin() *DecodedVIN {
	if x != nil {
		return x.DecodedVin
	}
	return nil
}

func (x *DecodeResultItem) GetError() *v1.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type DecodedVIN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vds string `protobuf:"bytes,1,opt,name=vds,proto3" json:"vds,omitempty"`
	Vis string `protobuf:"bytes,2,opt,name=vis,proto3" json:"vis,omitempty"`
	Wmi string `protobuf:"bytes,3,opt,name=wmi,proto3" json:"wmi,omitempty"`
}

func (x *DecodedVIN) Reset() {
	*x = DecodedVIN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_decoding_v1_vin_decoding_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodedVIN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodedVIN) ProtoMessage() {}

func (x *DecodedVIN) ProtoReflect() protoreflect.Message {
	mi := &file_vin_decoding_v1_vin_decoding_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodedVIN.ProtoReflect.Descriptor instead.
func (*DecodedVIN) Descriptor() ([]byte, []int) {
	return file_vin_decoding_v1_vin_decoding_proto_rawDescGZIP(), []int{3}
}

func (x *DecodedVIN) GetVds() string {
	if x != nil {
		return x.Vds
	}
	return ""
}

func (x *DecodedVIN) GetVis() string {
	if x != nil {
		return x.Vis
	}
	return ""
}

func (x *DecodedVIN) GetWmi() string {
	if x != nil {
		return x.Wmi
	}
	return ""
}

type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Check        bool   `protobuf:"varint,1,opt,name=check,proto3" json:"check,omitempty"`
	Country      string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Manufacturer string `protobuf:"bytes,3,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Region       string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	Year         uint32 `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vin_decoding_v1_vin_decoding_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_vin_decoding_v1_vin_decoding_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_vin_decoding_v1_vin_decoding_proto_rawDescGZIP(), []int{4}
}

func (x *Vehicle) GetCheck() bool {
	if x != nil {
		return x.Check
	}
	return false
}

func (x *Vehicle) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Vehicle) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *Vehicle) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Vehicle) GetYear() uint32 {
	if x != nil {
		return x.Year
	}
	return 0
}

var File_vin_decoding_v1_vin_decoding_proto protoreflect.FileDescriptor

var file_vin_decoding_v1_vin_decoding_proto_rawDesc = []byte{
	0x0a, 0x22, 0x76, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x76, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0d,
	0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x76, 0x69, 0x6e,
	0x73, 0x22, 0x49, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xac, 0x01, 0x0a,
	0x10, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x32, 0x0a, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64,
	0x5f, 0x76, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x69, 0x6e,
	0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63,
	0x6f, 0x64, 0x65, 0x64, 0x56, 0x49, 0x4e, 0x52, 0x0a, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64,
	0x56, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x42, 0x0a, 0x0a, 0x44,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x56, 0x49, 0x4e, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x69, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x77, 0x6d, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x6d, 0x69, 0x22,
	0x89, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6d,
	0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x32, 0x54, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1e, 0x2e, 0x76, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x76, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0xbd, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x6e, 0x5f, 0x64, 0x65,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x56, 0x69, 0x6e, 0x44, 0x65,
	0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x61,
	0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x69, 0x6e, 0x5f,
	0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69, 0x6e, 0x5f,
	0x64, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58,
	0xaa, 0x02, 0x0e, 0x56, 0x69, 0x6e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0e, 0x56, 0x69, 0x6e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1a, 0x56, 0x69, 0x6e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0f, 0x56, 0x69, 0x6e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vin_decoding_v1_vin_decoding_proto_rawDescOnce sync.Once
	file_vin_decoding_v1_vin_decoding_proto_rawDescData = file_vin_decoding_v1_vin_decoding_proto_rawDesc
)

func file_vin_decoding_v1_vin_decoding_proto_rawDescGZIP() []byte {
	file_vin_decoding_v1_vin_decoding_proto_rawDescOnce.Do(func() {
		file_vin_decoding_v1_vin_decoding_proto_rawDescData = protoimpl.X.CompressGZIP(file_vin_decoding_v1_vin_decoding_proto_rawDescData)
	})
	return file_vin_decoding_v1_vin_decoding_proto_rawDescData
}

var file_vin_decoding_v1_vin_decoding_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vin_decoding_v1_vin_decoding_proto_goTypes = []interface{}{
	(*DecodeRequest)(nil),    // 0: vin_decoding.v1.DecodeRequest
	(*DecodeResponse)(nil),   // 1: vin_decoding.v1.DecodeResponse
	(*DecodeResultItem)(nil), // 2: vin_decoding.v1.DecodeResultItem
	(*DecodedVIN)(nil),       // 3: vin_decoding.v1.DecodedVIN
	(*Vehicle)(nil),          // 4: vin_decoding.v1.Vehicle
	(*v1.Error)(nil),         // 5: common.v1.Error
}
var file_vin_decoding_v1_vin_decoding_proto_depIdxs = []int32{
	2, // 0: vin_decoding.v1.DecodeResponse.items:type_name -> vin_decoding.v1.DecodeResultItem
	4, // 1: vin_decoding.v1.DecodeResultItem.vehicle:type_name -> vin_decoding.v1.Vehicle
	3, // 2: vin_decoding.v1.DecodeResultItem.decoded_vin:type_name -> vin_decoding.v1.DecodedVIN
	5, // 3: vin_decoding.v1.DecodeResultItem.error:type_name -> common.v1.Error
	0, // 4: vin_decoding.v1.Service.Decode:input_type -> vin_decoding.v1.DecodeRequest
	1, // 5: vin_decoding.v1.Service.Decode:output_type -> vin_decoding.v1.DecodeResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_vin_decoding_v1_vin_decoding_proto_init() }
func file_vin_decoding_v1_vin_decoding_proto_init() {
	if File_vin_decoding_v1_vin_decoding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vin_decoding_v1_vin_decoding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_vin_decoding_v1_vin_decoding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodeResponse); i {
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
		file_vin_decoding_v1_vin_decoding_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_vin_decoding_v1_vin_decoding_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodedVIN); i {
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
		file_vin_decoding_v1_vin_decoding_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vehicle); i {
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
			RawDescriptor: file_vin_decoding_v1_vin_decoding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vin_decoding_v1_vin_decoding_proto_goTypes,
		DependencyIndexes: file_vin_decoding_v1_vin_decoding_proto_depIdxs,
		MessageInfos:      file_vin_decoding_v1_vin_decoding_proto_msgTypes,
	}.Build()
	File_vin_decoding_v1_vin_decoding_proto = out.File
	file_vin_decoding_v1_vin_decoding_proto_rawDesc = nil
	file_vin_decoding_v1_vin_decoding_proto_goTypes = nil
	file_vin_decoding_v1_vin_decoding_proto_depIdxs = nil
}
