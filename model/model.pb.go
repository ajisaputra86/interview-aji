// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: model.proto

package model

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

type CheckInOutType int32

const (
	CheckInOutType_IN  CheckInOutType = 0
	CheckInOutType_OUT CheckInOutType = 1
)

// Enum value maps for CheckInOutType.
var (
	CheckInOutType_name = map[int32]string{
		0: "IN",
		1: "OUT",
	}
	CheckInOutType_value = map[string]int32{
		"IN":  0,
		"OUT": 1,
	}
)

func (x CheckInOutType) Enum() *CheckInOutType {
	p := new(CheckInOutType)
	*p = x
	return p
}

func (x CheckInOutType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckInOutType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_proto_enumTypes[0].Descriptor()
}

func (CheckInOutType) Type() protoreflect.EnumType {
	return &file_model_proto_enumTypes[0]
}

func (x CheckInOutType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckInOutType.Descriptor instead.
func (CheckInOutType) EnumDescriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

type CheckInOutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string         `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Type       CheckInOutType `protobuf:"varint,2,opt,name=type,proto3,enum=v1.CheckInOutType" json:"type,omitempty"`
	LocationId int64          `protobuf:"varint,3,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
}

func (x *CheckInOutRequest) Reset() {
	*x = CheckInOutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInOutRequest) ProtoMessage() {}

func (x *CheckInOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInOutRequest.ProtoReflect.Descriptor instead.
func (*CheckInOutRequest) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *CheckInOutRequest) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *CheckInOutRequest) GetType() CheckInOutType {
	if x != nil {
		return x.Type
	}
	return CheckInOutType_IN
}

func (x *CheckInOutRequest) GetLocationId() int64 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *Location) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *ListLocation) Reset() {
	*x = ListLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocation) ProtoMessage() {}

func (x *ListLocation) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocation.ProtoReflect.Descriptor instead.
func (*ListLocation) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{2}
}

func (x *ListLocation) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

type CheckInOutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CheckinoutTime string         `protobuf:"bytes,2,opt,name=checkinout_time,json=checkinoutTime,proto3" json:"checkinout_time,omitempty"`
	CheckType      CheckInOutType `protobuf:"varint,3,opt,name=check_type,json=checkType,proto3,enum=v1.CheckInOutType" json:"check_type,omitempty"`
	Location       *Location      `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CheckInOutResponse) Reset() {
	*x = CheckInOutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInOutResponse) ProtoMessage() {}

func (x *CheckInOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInOutResponse.ProtoReflect.Descriptor instead.
func (*CheckInOutResponse) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3}
}

func (x *CheckInOutResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CheckInOutResponse) GetCheckinoutTime() string {
	if x != nil {
		return x.CheckinoutTime
	}
	return ""
}

func (x *CheckInOutResponse) GetCheckType() CheckInOutType {
	if x != nil {
		return x.CheckType
	}
	return CheckInOutType_IN
}

func (x *CheckInOutResponse) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type CheckInOutRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CheckinoutType CheckInOutType `protobuf:"varint,2,opt,name=checkinout_type,json=checkinoutType,proto3,enum=v1.CheckInOutType" json:"checkinout_type,omitempty"`
	Location       *Location      `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CheckInOutRecord) Reset() {
	*x = CheckInOutRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckInOutRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckInOutRecord) ProtoMessage() {}

func (x *CheckInOutRecord) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckInOutRecord.ProtoReflect.Descriptor instead.
func (*CheckInOutRecord) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{4}
}

func (x *CheckInOutRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CheckInOutRecord) GetCheckinoutType() CheckInOutType {
	if x != nil {
		return x.CheckinoutType
	}
	return CheckInOutType_IN
}

func (x *CheckInOutRecord) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x31, 0x22, 0x7c, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x4f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x4f, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x2e, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x3a, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xaa, 0x01, 0x0a, 0x12,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x6f, 0x75, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x4f, 0x75, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x6e, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a,
	0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x6e, 0x4f, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x69, 0x6e, 0x6f, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x21, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x4f,
	0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x6a, 0x72, 0x69, 0x72, 0x61, 0x68, 0x6d, 0x61,
	0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x61, 0x6a, 0x69, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData = file_model_proto_rawDesc
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_rawDescData)
	})
	return file_model_proto_rawDescData
}

var file_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_model_proto_goTypes = []interface{}{
	(CheckInOutType)(0),        // 0: v1.CheckInOutType
	(*CheckInOutRequest)(nil),  // 1: v1.CheckInOutRequest
	(*Location)(nil),           // 2: v1.Location
	(*ListLocation)(nil),       // 3: v1.ListLocation
	(*CheckInOutResponse)(nil), // 4: v1.CheckInOutResponse
	(*CheckInOutRecord)(nil),   // 5: v1.CheckInOutRecord
}
var file_model_proto_depIdxs = []int32{
	0, // 0: v1.CheckInOutRequest.type:type_name -> v1.CheckInOutType
	2, // 1: v1.ListLocation.locations:type_name -> v1.Location
	0, // 2: v1.CheckInOutResponse.check_type:type_name -> v1.CheckInOutType
	2, // 3: v1.CheckInOutResponse.location:type_name -> v1.Location
	0, // 4: v1.CheckInOutRecord.checkinout_type:type_name -> v1.CheckInOutType
	2, // 5: v1.CheckInOutRecord.location:type_name -> v1.Location
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInOutRequest); i {
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
		file_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocation); i {
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
		file_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInOutResponse); i {
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
		file_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckInOutRecord); i {
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
			RawDescriptor: file_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		EnumInfos:         file_model_proto_enumTypes,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_rawDesc = nil
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}
