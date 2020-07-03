// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: store/store.proto

package store

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type IdDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdDTO) Reset() {
	*x = IdDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdDTO) ProtoMessage() {}

func (x *IdDTO) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdDTO.ProtoReflect.Descriptor instead.
func (*IdDTO) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{0}
}

func (x *IdDTO) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 指定的一组ids
type IdsDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdsDto) Reset() {
	*x = IdsDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsDto) ProtoMessage() {}

func (x *IdsDto) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdsDto.ProtoReflect.Descriptor instead.
func (*IdsDto) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{1}
}

func (x *IdsDto) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type StoreDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Telephone   string               `protobuf:"bytes,3,opt,name=telephone,proto3" json:"telephone,omitempty"`
	ManagerId   string               `protobuf:"bytes,4,opt,name=manager_id,json=managerId,proto3" json:"manager_id,omitempty"`
	Logo        string               `protobuf:"bytes,5,opt,name=logo,proto3" json:"logo,omitempty"`
	Photos      string               `protobuf:"bytes,6,opt,name=photos,proto3" json:"photos,omitempty"`
	Address     string               `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Description string               `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Country     string               `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	CreateTime  *timestamp.Timestamp `protobuf:"bytes,10,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime  *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Status      uint32               `protobuf:"varint,12,opt,name=status,proto3" json:"status,omitempty"`
	DealerId    string               `protobuf:"bytes,13,opt,name=DealerId,proto3" json:"DealerId,omitempty"`
}

func (x *StoreDTO) Reset() {
	*x = StoreDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreDTO) ProtoMessage() {}

func (x *StoreDTO) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreDTO.ProtoReflect.Descriptor instead.
func (*StoreDTO) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{2}
}

func (x *StoreDTO) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StoreDTO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoreDTO) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *StoreDTO) GetManagerId() string {
	if x != nil {
		return x.ManagerId
	}
	return ""
}

func (x *StoreDTO) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *StoreDTO) GetPhotos() string {
	if x != nil {
		return x.Photos
	}
	return ""
}

func (x *StoreDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *StoreDTO) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *StoreDTO) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *StoreDTO) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *StoreDTO) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *StoreDTO) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *StoreDTO) GetDealerId() string {
	if x != nil {
		return x.DealerId
	}
	return ""
}

// store的简单信息dto
type StoreSimpleDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StoreSimpleDto) Reset() {
	*x = StoreSimpleDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreSimpleDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreSimpleDto) ProtoMessage() {}

func (x *StoreSimpleDto) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreSimpleDto.ProtoReflect.Descriptor instead.
func (*StoreSimpleDto) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{3}
}

func (x *StoreSimpleDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StoreSimpleDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StoreCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *Page  `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StoreCondition) Reset() {
	*x = StoreCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreCondition) ProtoMessage() {}

func (x *StoreCondition) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreCondition.ProtoReflect.Descriptor instead.
func (*StoreCondition) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{4}
}

func (x *StoreCondition) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *StoreCondition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo   uint32 `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize uint32 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{5}
}

func (x *Page) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *Page) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *any.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Code string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{6}
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Response) GetData() *any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type AddStoreData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddStoreData) Reset() {
	*x = AddStoreData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStoreData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStoreData) ProtoMessage() {}

func (x *AddStoreData) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStoreData.ProtoReflect.Descriptor instead.
func (*AddStoreData) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{7}
}

func (x *AddStoreData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_store_store_proto protoreflect.FileDescriptor

var file_store_store_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a,
	0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x64, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x06, 0x49, 0x64,
	0x73, 0x44, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x99, 0x03, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x34, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x3a, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5a, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x1e, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0xd7, 0x02, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x49, 0x64, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x0c, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x44, 0x54, 0x4f,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x0d, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x73, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x73, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_store_proto_rawDescOnce sync.Once
	file_store_store_proto_rawDescData = file_store_store_proto_rawDesc
)

func file_store_store_proto_rawDescGZIP() []byte {
	file_store_store_proto_rawDescOnce.Do(func() {
		file_store_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_store_proto_rawDescData)
	})
	return file_store_store_proto_rawDescData
}

var file_store_store_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_store_store_proto_goTypes = []interface{}{
	(*IdDTO)(nil),               // 0: store.IdDTO
	(*IdsDto)(nil),              // 1: store.IdsDto
	(*StoreDTO)(nil),            // 2: store.StoreDTO
	(*StoreSimpleDto)(nil),      // 3: store.StoreSimpleDto
	(*StoreCondition)(nil),      // 4: store.StoreCondition
	(*Page)(nil),                // 5: store.Page
	(*Response)(nil),            // 6: store.Response
	(*AddStoreData)(nil),        // 7: store.AddStoreData
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*any.Any)(nil),             // 9: google.protobuf.Any
	(*common.Response)(nil),     // 10: common.Response
}
var file_store_store_proto_depIdxs = []int32{
	8,  // 0: store.StoreDTO.createTime:type_name -> google.protobuf.Timestamp
	8,  // 1: store.StoreDTO.updateTime:type_name -> google.protobuf.Timestamp
	5,  // 2: store.StoreCondition.page:type_name -> store.Page
	9,  // 3: store.Response.data:type_name -> google.protobuf.Any
	2,  // 4: store.Store.Add:input_type -> store.StoreDTO
	4,  // 5: store.Store.List:input_type -> store.StoreCondition
	2,  // 6: store.Store.Update:input_type -> store.StoreDTO
	0,  // 7: store.Store.Delete:input_type -> store.IdDTO
	0,  // 8: store.Store.Get:input_type -> store.IdDTO
	1,  // 9: store.Store.ListSimpleInfoByIds:input_type -> store.IdsDto
	1,  // 10: store.Store.ListLocation:input_type -> store.IdsDto
	10, // 11: store.Store.Add:output_type -> common.Response
	10, // 12: store.Store.List:output_type -> common.Response
	10, // 13: store.Store.Update:output_type -> common.Response
	10, // 14: store.Store.Delete:output_type -> common.Response
	10, // 15: store.Store.Get:output_type -> common.Response
	10, // 16: store.Store.ListSimpleInfoByIds:output_type -> common.Response
	10, // 17: store.Store.ListLocation:output_type -> common.Response
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_store_store_proto_init() }
func file_store_store_proto_init() {
	if File_store_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdDTO); i {
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
		file_store_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdsDto); i {
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
		file_store_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreDTO); i {
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
		file_store_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreSimpleDto); i {
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
		file_store_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreCondition); i {
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
		file_store_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_store_store_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_store_store_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStoreData); i {
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
			RawDescriptor: file_store_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_store_store_proto_goTypes,
		DependencyIndexes: file_store_store_proto_depIdxs,
		MessageInfos:      file_store_store_proto_msgTypes,
	}.Build()
	File_store_store_proto = out.File
	file_store_store_proto_rawDesc = nil
	file_store_store_proto_goTypes = nil
	file_store_store_proto_depIdxs = nil
}
