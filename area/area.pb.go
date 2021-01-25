// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.1
// source: area/area.proto

package area

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
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

type ParentCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Local string `protobuf:"bytes,2,opt,name=local,proto3" json:"local,omitempty"`
}

func (x *ParentCode) Reset() {
	*x = ParentCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_area_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParentCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParentCode) ProtoMessage() {}

func (x *ParentCode) ProtoReflect() protoreflect.Message {
	mi := &file_area_area_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParentCode.ProtoReflect.Descriptor instead.
func (*ParentCode) Descriptor() ([]byte, []int) {
	return file_area_area_proto_rawDescGZIP(), []int{0}
}

func (x *ParentCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ParentCode) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

type AreaDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Code       string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ParentCode string `protobuf:"bytes,9,opt,name=parentCode,proto3" json:"parentCode,omitempty"`
	CreateTime int64  `protobuf:"varint,10,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime int64  `protobuf:"varint,11,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *AreaDto) Reset() {
	*x = AreaDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_area_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaDto) ProtoMessage() {}

func (x *AreaDto) ProtoReflect() protoreflect.Message {
	mi := &file_area_area_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaDto.ProtoReflect.Descriptor instead.
func (*AreaDto) Descriptor() ([]byte, []int) {
	return file_area_area_proto_rawDescGZIP(), []int{1}
}

func (x *AreaDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AreaDto) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AreaDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AreaDto) GetParentCode() string {
	if x != nil {
		return x.ParentCode
	}
	return ""
}

func (x *AreaDto) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *AreaDto) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type AreaList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaDTOs []*AreaDto `protobuf:"bytes,1,rep,name=areaDTOs,proto3" json:"areaDTOs,omitempty"`
}

func (x *AreaList) Reset() {
	*x = AreaList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_area_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaList) ProtoMessage() {}

func (x *AreaList) ProtoReflect() protoreflect.Message {
	mi := &file_area_area_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaList.ProtoReflect.Descriptor instead.
func (*AreaList) Descriptor() ([]byte, []int) {
	return file_area_area_proto_rawDescGZIP(), []int{2}
}

func (x *AreaList) GetAreaDTOs() []*AreaDto {
	if x != nil {
		return x.AreaDTOs
	}
	return nil
}

type StateAndAreaList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateMap map[string]*AreaDto `protobuf:"bytes,1,rep,name=stateMap,proto3" json:"stateMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AreaMap  map[string]*AreaDto `protobuf:"bytes,2,rep,name=areaMap,proto3" json:"areaMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StateAndAreaList) Reset() {
	*x = StateAndAreaList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_area_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateAndAreaList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateAndAreaList) ProtoMessage() {}

func (x *StateAndAreaList) ProtoReflect() protoreflect.Message {
	mi := &file_area_area_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateAndAreaList.ProtoReflect.Descriptor instead.
func (*StateAndAreaList) Descriptor() ([]byte, []int) {
	return file_area_area_proto_rawDescGZIP(), []int{3}
}

func (x *StateAndAreaList) GetStateMap() map[string]*AreaDto {
	if x != nil {
		return x.StateMap
	}
	return nil
}

func (x *StateAndAreaList) GetAreaMap() map[string]*AreaDto {
	if x != nil {
		return x.AreaMap
	}
	return nil
}

type StateAndAreaCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StateCodes []string `protobuf:"bytes,1,rep,name=stateCodes,proto3" json:"stateCodes,omitempty"`
	AreaCodes  []string `protobuf:"bytes,2,rep,name=areaCodes,proto3" json:"areaCodes,omitempty"`
}

func (x *StateAndAreaCodeReq) Reset() {
	*x = StateAndAreaCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_area_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StateAndAreaCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StateAndAreaCodeReq) ProtoMessage() {}

func (x *StateAndAreaCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_area_area_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StateAndAreaCodeReq.ProtoReflect.Descriptor instead.
func (*StateAndAreaCodeReq) Descriptor() ([]byte, []int) {
	return file_area_area_proto_rawDescGZIP(), []int{4}
}

func (x *StateAndAreaCodeReq) GetStateCodes() []string {
	if x != nil {
		return x.StateCodes
	}
	return nil
}

func (x *StateAndAreaCodeReq) GetAreaCodes() []string {
	if x != nil {
		return x.AreaCodes
	}
	return nil
}

var File_area_area_proto protoreflect.FileDescriptor

var file_area_area_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x72, 0x65, 0x61, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f,
	0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0xa1, 0x01, 0x0a,
	0x07, 0x41, 0x72, 0x65, 0x61, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x37, 0x0a, 0x08, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08,
	0x61, 0x72, 0x65, 0x61, 0x44, 0x54, 0x4f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x44, 0x74, 0x6f, 0x52,
	0x08, 0x61, 0x72, 0x65, 0x61, 0x44, 0x54, 0x4f, 0x73, 0x22, 0xb2, 0x02, 0x0a, 0x10, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x42,
	0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x41,
	0x6e, 0x64, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4d,
	0x61, 0x70, 0x12, 0x3f, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x4d, 0x61, 0x70, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x64, 0x41, 0x72, 0x65, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x41, 0x72,
	0x65, 0x61, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x72, 0x65, 0x61,
	0x4d, 0x61, 0x70, 0x1a, 0x4c, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x41,
	0x72, 0x65, 0x61, 0x44, 0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x4b, 0x0a, 0x0c, 0x41, 0x72, 0x65, 0x61, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x72, 0x65, 0x61,
	0x44, 0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x53,
	0x0a, 0x13, 0x53, 0x74, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x32, 0xbf, 0x02, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x2a, 0x0a, 0x03,
	0x41, 0x64, 0x64, 0x12, 0x0f, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x72, 0x65,
	0x61, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x0f, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x44, 0x74, 0x6f,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x54, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x12, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x41, 0x72, 0x65, 0x61, 0x42, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x64, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x65, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_area_area_proto_rawDescOnce sync.Once
	file_area_area_proto_rawDescData = file_area_area_proto_rawDesc
)

func file_area_area_proto_rawDescGZIP() []byte {
	file_area_area_proto_rawDescOnce.Do(func() {
		file_area_area_proto_rawDescData = protoimpl.X.CompressGZIP(file_area_area_proto_rawDescData)
	})
	return file_area_area_proto_rawDescData
}

var file_area_area_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_area_area_proto_goTypes = []interface{}{
	(*ParentCode)(nil),          // 0: dealer.ParentCode
	(*AreaDto)(nil),             // 1: dealer.AreaDto
	(*AreaList)(nil),            // 2: dealer.AreaList
	(*StateAndAreaList)(nil),    // 3: dealer.StateAndAreaList
	(*StateAndAreaCodeReq)(nil), // 4: dealer.StateAndAreaCodeReq
	nil,                         // 5: dealer.StateAndAreaList.StateMapEntry
	nil,                         // 6: dealer.StateAndAreaList.AreaMapEntry
	(*common.IdDto)(nil),        // 7: common.IdDto
	(*common.LocalDto)(nil),     // 8: common.LocalDto
	(*common.Response)(nil),     // 9: common.Response
}
var file_area_area_proto_depIdxs = []int32{
	1,  // 0: dealer.AreaList.areaDTOs:type_name -> dealer.AreaDto
	5,  // 1: dealer.StateAndAreaList.stateMap:type_name -> dealer.StateAndAreaList.StateMapEntry
	6,  // 2: dealer.StateAndAreaList.areaMap:type_name -> dealer.StateAndAreaList.AreaMapEntry
	1,  // 3: dealer.StateAndAreaList.StateMapEntry.value:type_name -> dealer.AreaDto
	1,  // 4: dealer.StateAndAreaList.AreaMapEntry.value:type_name -> dealer.AreaDto
	1,  // 5: dealer.Area.Add:input_type -> dealer.AreaDto
	7,  // 6: dealer.Area.Delete:input_type -> common.IdDto
	1,  // 7: dealer.Area.Update:input_type -> dealer.AreaDto
	8,  // 8: dealer.Area.TopList:input_type -> common.LocalDto
	0,  // 9: dealer.Area.ChildList:input_type -> dealer.ParentCode
	4,  // 10: dealer.Area.GetStateAndAreaByCodes:input_type -> dealer.StateAndAreaCodeReq
	9,  // 11: dealer.Area.Add:output_type -> common.Response
	9,  // 12: dealer.Area.Delete:output_type -> common.Response
	9,  // 13: dealer.Area.Update:output_type -> common.Response
	9,  // 14: dealer.Area.TopList:output_type -> common.Response
	9,  // 15: dealer.Area.ChildList:output_type -> common.Response
	9,  // 16: dealer.Area.GetStateAndAreaByCodes:output_type -> common.Response
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_area_area_proto_init() }
func file_area_area_proto_init() {
	if File_area_area_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_area_area_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParentCode); i {
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
		file_area_area_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreaDto); i {
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
		file_area_area_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreaList); i {
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
		file_area_area_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateAndAreaList); i {
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
		file_area_area_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StateAndAreaCodeReq); i {
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
			RawDescriptor: file_area_area_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_area_area_proto_goTypes,
		DependencyIndexes: file_area_area_proto_depIdxs,
		MessageInfos:      file_area_area_proto_msgTypes,
	}.Build()
	File_area_area_proto = out.File
	file_area_area_proto_rawDesc = nil
	file_area_area_proto_goTypes = nil
	file_area_area_proto_depIdxs = nil
}
