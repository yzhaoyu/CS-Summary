// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: smartsheet.proto

package smartsheetapi

import (
	proto "github.com/golang/protobuf/proto"
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

type ResourceType int32

const (
	ResourceType_RESOURCE_TYPE_UNKNOWN    ResourceType = 0
	ResourceType_RESOURCE_TYPE_VIEW       ResourceType = 1
	ResourceType_RESOURCE_TYPE_RECORD     ResourceType = 2
	ResourceType_RESOURCE_TYPE_FIELD      ResourceType = 3
	ResourceType_RESOURCE_TYPE_ATTACHMENT ResourceType = 4
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "RESOURCE_TYPE_UNKNOWN",
		1: "RESOURCE_TYPE_VIEW",
		2: "RESOURCE_TYPE_RECORD",
		3: "RESOURCE_TYPE_FIELD",
		4: "RESOURCE_TYPE_ATTACHMENT",
	}
	ResourceType_value = map[string]int32{
		"RESOURCE_TYPE_UNKNOWN":    0,
		"RESOURCE_TYPE_VIEW":       1,
		"RESOURCE_TYPE_RECORD":     2,
		"RESOURCE_TYPE_FIELD":      3,
		"RESOURCE_TYPE_ATTACHMENT": 4,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_smartsheet_proto_enumTypes[0].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_smartsheet_proto_enumTypes[0]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_smartsheet_proto_rawDescGZIP(), []int{0}
}

type CommandType int32

const (
	CommandType_COMMAND_TYPE_UNKNOWN       CommandType = 0
	CommandType_COMMAND_TYPE_ADDVIEW       CommandType = 1
	CommandType_COMMAND_TYPE_DELETEVIEWS   CommandType = 2
	CommandType_COMMAND_TYPE_ADDRECORDS    CommandType = 3
	CommandType_COMMAND_TYPE_UPDATERECORDS CommandType = 4
	CommandType_COMMAND_TYPE_DELETERECORDS CommandType = 5
	CommandType_COMMAND_TYPE_ADDFIELDS     CommandType = 6
	CommandType_COMMAND_TYPE_UPDATEFIELDS  CommandType = 7
	CommandType_COMMAND_TYPE_DELETEFIELDS  CommandType = 8
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0: "COMMAND_TYPE_UNKNOWN",
		1: "COMMAND_TYPE_ADDVIEW",
		2: "COMMAND_TYPE_DELETEVIEWS",
		3: "COMMAND_TYPE_ADDRECORDS",
		4: "COMMAND_TYPE_UPDATERECORDS",
		5: "COMMAND_TYPE_DELETERECORDS",
		6: "COMMAND_TYPE_ADDFIELDS",
		7: "COMMAND_TYPE_UPDATEFIELDS",
		8: "COMMAND_TYPE_DELETEFIELDS",
	}
	CommandType_value = map[string]int32{
		"COMMAND_TYPE_UNKNOWN":       0,
		"COMMAND_TYPE_ADDVIEW":       1,
		"COMMAND_TYPE_DELETEVIEWS":   2,
		"COMMAND_TYPE_ADDRECORDS":    3,
		"COMMAND_TYPE_UPDATERECORDS": 4,
		"COMMAND_TYPE_DELETERECORDS": 5,
		"COMMAND_TYPE_ADDFIELDS":     6,
		"COMMAND_TYPE_UPDATEFIELDS":  7,
		"COMMAND_TYPE_DELETEFIELDS":  8,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_smartsheet_proto_enumTypes[1].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_smartsheet_proto_enumTypes[1]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_smartsheet_proto_rawDescGZIP(), []int{1}
}

// smartsheet api 查询请求包
type GetSmartsheetResourceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PadID  string       `protobuf:"bytes,1,opt,name=padID,proto3" json:"padID,omitempty"`                                // 文档ID
	SubID  string       `protobuf:"bytes,2,opt,name=subID,proto3" json:"subID,omitempty"`                                // 子表ID
	Type   ResourceType `protobuf:"varint,3,opt,name=type,proto3,enum=smartsheetapi.ResourceType" json:"type,omitempty"` // 要获取的资源类型，包括view、field、record、attachment
	Params []byte       `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`                              // 每种资源涉及到的参数
}

func (x *GetSmartsheetResourceReq) Reset() {
	*x = GetSmartsheetResourceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartsheet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSmartsheetResourceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSmartsheetResourceReq) ProtoMessage() {}

func (x *GetSmartsheetResourceReq) ProtoReflect() protoreflect.Message {
	mi := &file_smartsheet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSmartsheetResourceReq.ProtoReflect.Descriptor instead.
func (*GetSmartsheetResourceReq) Descriptor() ([]byte, []int) {
	return file_smartsheet_proto_rawDescGZIP(), []int{0}
}

func (x *GetSmartsheetResourceReq) GetPadID() string {
	if x != nil {
		return x.PadID
	}
	return ""
}

func (x *GetSmartsheetResourceReq) GetSubID() string {
	if x != nil {
		return x.SubID
	}
	return ""
}

func (x *GetSmartsheetResourceReq) GetType() ResourceType {
	if x != nil {
		return x.Type
	}
	return ResourceType_RESOURCE_TYPE_UNKNOWN
}

func (x *GetSmartsheetResourceReq) GetParams() []byte {
	if x != nil {
		return x.Params
	}
	return nil
}

// smartsheet api 查询回包
type GetSmartsheetResourceRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // 后台返回的数据
}

func (x *GetSmartsheetResourceRsp) Reset() {
	*x = GetSmartsheetResourceRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartsheet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSmartsheetResourceRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSmartsheetResourceRsp) ProtoMessage() {}

func (x *GetSmartsheetResourceRsp) ProtoReflect() protoreflect.Message {
	mi := &file_smartsheet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSmartsheetResourceRsp.ProtoReflect.Descriptor instead.
func (*GetSmartsheetResourceRsp) Descriptor() ([]byte, []int) {
	return file_smartsheet_proto_rawDescGZIP(), []int{1}
}

func (x *GetSmartsheetResourceRsp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// smartsheet api 编辑请求包
type EditSmartsheetResourceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PadID   string      `protobuf:"bytes,1,opt,name=padID,proto3" json:"padID,omitempty"`
	SubID   string      `protobuf:"bytes,2,opt,name=subID,proto3" json:"subID,omitempty"`
	Command CommandType `protobuf:"varint,3,opt,name=command,proto3,enum=smartsheetapi.CommandType" json:"command,omitempty"`
	Params  []byte      `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *EditSmartsheetResourceReq) Reset() {
	*x = EditSmartsheetResourceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartsheet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditSmartsheetResourceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditSmartsheetResourceReq) ProtoMessage() {}

func (x *EditSmartsheetResourceReq) ProtoReflect() protoreflect.Message {
	mi := &file_smartsheet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditSmartsheetResourceReq.ProtoReflect.Descriptor instead.
func (*EditSmartsheetResourceReq) Descriptor() ([]byte, []int) {
	return file_smartsheet_proto_rawDescGZIP(), []int{2}
}

func (x *EditSmartsheetResourceReq) GetPadID() string {
	if x != nil {
		return x.PadID
	}
	return ""
}

func (x *EditSmartsheetResourceReq) GetSubID() string {
	if x != nil {
		return x.SubID
	}
	return ""
}

func (x *EditSmartsheetResourceReq) GetCommand() CommandType {
	if x != nil {
		return x.Command
	}
	return CommandType_COMMAND_TYPE_UNKNOWN
}

func (x *EditSmartsheetResourceReq) GetParams() []byte {
	if x != nil {
		return x.Params
	}
	return nil
}

// smartsheet api 编辑请求回包
type EditSmartsheetResourceRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"` // 后台返回的数据
}

func (x *EditSmartsheetResourceRsp) Reset() {
	*x = EditSmartsheetResourceRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smartsheet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditSmartsheetResourceRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditSmartsheetResourceRsp) ProtoMessage() {}

func (x *EditSmartsheetResourceRsp) ProtoReflect() protoreflect.Message {
	mi := &file_smartsheet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditSmartsheetResourceRsp.ProtoReflect.Descriptor instead.
func (*EditSmartsheetResourceRsp) Descriptor() ([]byte, []int) {
	return file_smartsheet_proto_rawDescGZIP(), []int{3}
}

func (x *EditSmartsheetResourceRsp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_smartsheet_proto protoreflect.FileDescriptor

var file_smartsheet_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x61, 0x70,
	0x69, 0x22, 0x8f, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x61, 0x64, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x75, 0x62, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x75, 0x62, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x73, 0x68, 0x65, 0x65, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x22, 0x2e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x73,
	0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x95, 0x01, 0x0a, 0x19, 0x45, 0x64, 0x69, 0x74, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x64, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x61, 0x64, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x75, 0x62, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x75, 0x62, 0x49, 0x44, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x2f, 0x0a, 0x19, 0x45,
	0x64, 0x69, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x92, 0x01, 0x0a,
	0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x10, 0x01,
	0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x45, 0x4c,
	0x44, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x41, 0x43, 0x48, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x04, 0x2a, 0x96, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43,
	0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x56,
	0x49, 0x45, 0x57, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x56, 0x49, 0x45, 0x57,
	0x53, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x03,
	0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x04,
	0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x53, 0x10, 0x05,
	0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x44, 0x44, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x10, 0x06, 0x12, 0x1d, 0x0a, 0x19,
	0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19, 0x43,
	0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x53, 0x10, 0x08, 0x32, 0xf3, 0x01, 0x0a, 0x18, 0x53,
	0x6d, 0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x50, 0x49,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x61,
	0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x73, 0x70, 0x12, 0x6c, 0x0a, 0x16, 0x45, 0x64, 0x69, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x73,
	0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x2e, 0x73,
	0x6d, 0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x64, 0x69,
	0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x68,
	0x65, 0x65, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74,
	0x73, 0x68, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x73, 0x70,
	0x42, 0x1b, 0x5a, 0x19, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x68, 0x65, 0x65, 0x74, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smartsheet_proto_rawDescOnce sync.Once
	file_smartsheet_proto_rawDescData = file_smartsheet_proto_rawDesc
)

func file_smartsheet_proto_rawDescGZIP() []byte {
	file_smartsheet_proto_rawDescOnce.Do(func() {
		file_smartsheet_proto_rawDescData = protoimpl.X.CompressGZIP(file_smartsheet_proto_rawDescData)
	})
	return file_smartsheet_proto_rawDescData
}

var file_smartsheet_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_smartsheet_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_smartsheet_proto_goTypes = []interface{}{
	(ResourceType)(0),                 // 0: smartsheetapi.ResourceType
	(CommandType)(0),                  // 1: smartsheetapi.CommandType
	(*GetSmartsheetResourceReq)(nil),  // 2: smartsheetapi.GetSmartsheetResourceReq
	(*GetSmartsheetResourceRsp)(nil),  // 3: smartsheetapi.GetSmartsheetResourceRsp
	(*EditSmartsheetResourceReq)(nil), // 4: smartsheetapi.EditSmartsheetResourceReq
	(*EditSmartsheetResourceRsp)(nil), // 5: smartsheetapi.EditSmartsheetResourceRsp
}
var file_smartsheet_proto_depIdxs = []int32{
	0, // 0: smartsheetapi.GetSmartsheetResourceReq.type:type_name -> smartsheetapi.ResourceType
	1, // 1: smartsheetapi.EditSmartsheetResourceReq.command:type_name -> smartsheetapi.CommandType
	2, // 2: smartsheetapi.SmartsheetOpenAPIService.GetSmartsheetResource:input_type -> smartsheetapi.GetSmartsheetResourceReq
	4, // 3: smartsheetapi.SmartsheetOpenAPIService.EditSmartsheetResource:input_type -> smartsheetapi.EditSmartsheetResourceReq
	3, // 4: smartsheetapi.SmartsheetOpenAPIService.GetSmartsheetResource:output_type -> smartsheetapi.GetSmartsheetResourceRsp
	5, // 5: smartsheetapi.SmartsheetOpenAPIService.EditSmartsheetResource:output_type -> smartsheetapi.EditSmartsheetResourceRsp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_smartsheet_proto_init() }
func file_smartsheet_proto_init() {
	if File_smartsheet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smartsheet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSmartsheetResourceReq); i {
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
		file_smartsheet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSmartsheetResourceRsp); i {
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
		file_smartsheet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditSmartsheetResourceReq); i {
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
		file_smartsheet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditSmartsheetResourceRsp); i {
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
			RawDescriptor: file_smartsheet_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smartsheet_proto_goTypes,
		DependencyIndexes: file_smartsheet_proto_depIdxs,
		EnumInfos:         file_smartsheet_proto_enumTypes,
		MessageInfos:      file_smartsheet_proto_msgTypes,
	}.Build()
	File_smartsheet_proto = out.File
	file_smartsheet_proto_rawDesc = nil
	file_smartsheet_proto_goTypes = nil
	file_smartsheet_proto_depIdxs = nil
}