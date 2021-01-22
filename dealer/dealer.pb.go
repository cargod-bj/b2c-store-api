// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: dealer/dealer.proto

package dealer

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

type DealersDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DealersDto []*DealerDto `protobuf:"bytes,1,rep,name=dealersDto,proto3" json:"dealersDto,omitempty"`
}

func (x *DealersDto) Reset() {
	*x = DealersDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dealer_dealer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DealersDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DealersDto) ProtoMessage() {}

func (x *DealersDto) ProtoReflect() protoreflect.Message {
	mi := &file_dealer_dealer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DealersDto.ProtoReflect.Descriptor instead.
func (*DealersDto) Descriptor() ([]byte, []int) {
	return file_dealer_dealer_proto_rawDescGZIP(), []int{0}
}

func (x *DealersDto) GetDealersDto() []*DealerDto {
	if x != nil {
		return x.DealersDto
	}
	return nil
}

type DealerDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	DealerId       string `protobuf:"bytes,2,opt,name=DealerId,proto3" json:"DealerId,omitempty"`
	DealershipType uint32 `protobuf:"varint,3,opt,name=DealershipType,proto3" json:"DealershipType,omitempty"`
	Ssm            string `protobuf:"bytes,4,opt,name=Ssm,proto3" json:"Ssm,omitempty"`
	Name           string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Address        string `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	Logo           string `protobuf:"bytes,7,opt,name=Logo,proto3" json:"Logo,omitempty"`
	Email          string `protobuf:"bytes,8,opt,name=Email,proto3" json:"Email,omitempty"`
	Telephone      string `protobuf:"bytes,9,opt,name=Telephone,proto3" json:"Telephone,omitempty"`
	CreateTime     int64  `protobuf:"varint,10,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime     int64  `protobuf:"varint,11,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	State          string `protobuf:"bytes,12,opt,name=State,proto3" json:"State,omitempty"`
	Area           string `protobuf:"bytes,13,opt,name=Area,proto3" json:"Area,omitempty"`
	DealershipName string `protobuf:"bytes,14,opt,name=DealershipName,proto3" json:"DealershipName,omitempty"`
	Status         uint32 `protobuf:"varint,15,opt,name=Status,proto3" json:"Status,omitempty"`
	Remark         string `protobuf:"bytes,16,opt,name=Remark,proto3" json:"Remark,omitempty"`
}

func (x *DealerDto) Reset() {
	*x = DealerDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dealer_dealer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DealerDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DealerDto) ProtoMessage() {}

func (x *DealerDto) ProtoReflect() protoreflect.Message {
	mi := &file_dealer_dealer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DealerDto.ProtoReflect.Descriptor instead.
func (*DealerDto) Descriptor() ([]byte, []int) {
	return file_dealer_dealer_proto_rawDescGZIP(), []int{1}
}

func (x *DealerDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DealerDto) GetDealerId() string {
	if x != nil {
		return x.DealerId
	}
	return ""
}

func (x *DealerDto) GetDealershipType() uint32 {
	if x != nil {
		return x.DealershipType
	}
	return 0
}

func (x *DealerDto) GetSsm() string {
	if x != nil {
		return x.Ssm
	}
	return ""
}

func (x *DealerDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DealerDto) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DealerDto) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *DealerDto) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DealerDto) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *DealerDto) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *DealerDto) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *DealerDto) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *DealerDto) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *DealerDto) GetDealershipName() string {
	if x != nil {
		return x.DealershipName
	}
	return ""
}

func (x *DealerDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DealerDto) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type DealerListDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	DealerId          string `protobuf:"bytes,2,opt,name=DealerId,proto3" json:"DealerId,omitempty"`
	DealershipType    uint32 `protobuf:"varint,3,opt,name=DealershipType,proto3" json:"DealershipType,omitempty"`
	Ssm               string `protobuf:"bytes,4,opt,name=Ssm,proto3" json:"Ssm,omitempty"`
	Name              string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Address           string `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	Logo              string `protobuf:"bytes,7,opt,name=Logo,proto3" json:"Logo,omitempty"`
	Email             string `protobuf:"bytes,8,opt,name=Email,proto3" json:"Email,omitempty"`
	Telephone         string `protobuf:"bytes,9,opt,name=Telephone,proto3" json:"Telephone,omitempty"`
	CreateTime        int64  `protobuf:"varint,10,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime        int64  `protobuf:"varint,11,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	State             string `protobuf:"bytes,12,opt,name=State,proto3" json:"State,omitempty"`
	Area              string `protobuf:"bytes,13,opt,name=Area,proto3" json:"Area,omitempty"`
	DealershipName    string `protobuf:"bytes,14,opt,name=DealershipName,proto3" json:"DealershipName,omitempty"`
	Status            uint32 `protobuf:"varint,15,opt,name=Status,proto3" json:"Status,omitempty"`
	Remark            string `protobuf:"bytes,16,opt,name=Remark,proto3" json:"Remark,omitempty"`
	StateName         string `protobuf:"bytes,17,opt,name=StateName,proto3" json:"StateName,omitempty"`
	AreaName          string `protobuf:"bytes,18,opt,name=AreaName,proto3" json:"AreaName,omitempty"`
	StatusDescription string `protobuf:"bytes,19,opt,name=StatusDescription,proto3" json:"StatusDescription,omitempty"`
}

func (x *DealerListDTO) Reset() {
	*x = DealerListDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dealer_dealer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DealerListDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DealerListDTO) ProtoMessage() {}

func (x *DealerListDTO) ProtoReflect() protoreflect.Message {
	mi := &file_dealer_dealer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DealerListDTO.ProtoReflect.Descriptor instead.
func (*DealerListDTO) Descriptor() ([]byte, []int) {
	return file_dealer_dealer_proto_rawDescGZIP(), []int{2}
}

func (x *DealerListDTO) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DealerListDTO) GetDealerId() string {
	if x != nil {
		return x.DealerId
	}
	return ""
}

func (x *DealerListDTO) GetDealershipType() uint32 {
	if x != nil {
		return x.DealershipType
	}
	return 0
}

func (x *DealerListDTO) GetSsm() string {
	if x != nil {
		return x.Ssm
	}
	return ""
}

func (x *DealerListDTO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DealerListDTO) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DealerListDTO) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *DealerListDTO) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DealerListDTO) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *DealerListDTO) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *DealerListDTO) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *DealerListDTO) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *DealerListDTO) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *DealerListDTO) GetDealershipName() string {
	if x != nil {
		return x.DealershipName
	}
	return ""
}

func (x *DealerListDTO) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DealerListDTO) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *DealerListDTO) GetStateName() string {
	if x != nil {
		return x.StateName
	}
	return ""
}

func (x *DealerListDTO) GetAreaName() string {
	if x != nil {
		return x.AreaName
	}
	return ""
}

func (x *DealerListDTO) GetStatusDescription() string {
	if x != nil {
		return x.StatusDescription
	}
	return ""
}

type SyncDealerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DealerId       string `protobuf:"bytes,2,opt,name=DealerId,proto3" json:"DealerId,omitempty"`
	DealershipType uint32 `protobuf:"varint,3,opt,name=DealershipType,proto3" json:"DealershipType,omitempty"`
	Name           string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *SyncDealerData) Reset() {
	*x = SyncDealerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dealer_dealer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncDealerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncDealerData) ProtoMessage() {}

func (x *SyncDealerData) ProtoReflect() protoreflect.Message {
	mi := &file_dealer_dealer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncDealerData.ProtoReflect.Descriptor instead.
func (*SyncDealerData) Descriptor() ([]byte, []int) {
	return file_dealer_dealer_proto_rawDescGZIP(), []int{3}
}

func (x *SyncDealerData) GetDealerId() string {
	if x != nil {
		return x.DealerId
	}
	return ""
}

func (x *SyncDealerData) GetDealershipType() uint32 {
	if x != nil {
		return x.DealershipType
	}
	return 0
}

func (x *SyncDealerData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type IdsDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []uint64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *IdsDto) Reset() {
	*x = IdsDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dealer_dealer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsDto) ProtoMessage() {}

func (x *IdsDto) ProtoReflect() protoreflect.Message {
	mi := &file_dealer_dealer_proto_msgTypes[4]
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
	return file_dealer_dealer_proto_rawDescGZIP(), []int{4}
}

func (x *IdsDto) GetId() []uint64 {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_dealer_dealer_proto protoreflect.FileDescriptor

var file_dealer_dealer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x1a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64,
	0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x6c,
	0x65, 0x72, 0x73, 0x44, 0x74, 0x6f, 0x12, 0x31, 0x0a, 0x0a, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72,
	0x73, 0x44, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x65, 0x61,
	0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x0a, 0x64,
	0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x44, 0x74, 0x6f, 0x22, 0xa9, 0x03, 0x0a, 0x09, 0x44, 0x65,
	0x61, 0x6c, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x44, 0x65, 0x61,
	0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53,
	0x73, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x73, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4c,
	0x6f, 0x67, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x65,
	0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x26, 0x0a,
	0x0e, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x95, 0x04, 0x0a, 0x0d, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x54, 0x4f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x44, 0x65, 0x61,
	0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53,
	0x73, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x73, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4c,
	0x6f, 0x67, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x65,
	0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x26, 0x0a,
	0x0e, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x72, 0x65, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x68, 0x0a,
	0x0e, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x44,
	0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x49, 0x64, 0x73, 0x44, 0x74,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x32, 0xd5, 0x02, 0x0a, 0x06, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x03,
	0x41, 0x64, 0x64, 0x12, 0x11, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x61,
	0x6c, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x44, 0x65,
	0x61, 0x6c, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x61, 0x6c, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x44, 0x65,
	0x61, 0x6c, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x54,
	0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x49,
	0x64, 0x73, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62,
	0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dealer_dealer_proto_rawDescOnce sync.Once
	file_dealer_dealer_proto_rawDescData = file_dealer_dealer_proto_rawDesc
)

func file_dealer_dealer_proto_rawDescGZIP() []byte {
	file_dealer_dealer_proto_rawDescOnce.Do(func() {
		file_dealer_dealer_proto_rawDescData = protoimpl.X.CompressGZIP(file_dealer_dealer_proto_rawDescData)
	})
	return file_dealer_dealer_proto_rawDescData
}

var file_dealer_dealer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_dealer_dealer_proto_goTypes = []interface{}{
	(*DealersDto)(nil),        // 0: dealer.DealersDto
	(*DealerDto)(nil),         // 1: dealer.DealerDto
	(*DealerListDTO)(nil),     // 2: dealer.DealerListDTO
	(*SyncDealerData)(nil),    // 3: dealer.SyncDealerData
	(*IdsDto)(nil),            // 4: dealer.IdsDto
	(*common.Page)(nil),       // 5: common.Page
	(*common.IdLocalDTO)(nil), // 6: common.IdLocalDTO
	(*common.Response)(nil),   // 7: common.Response
}
var file_dealer_dealer_proto_depIdxs = []int32{
	1, // 0: dealer.DealersDto.dealersDto:type_name -> dealer.DealerDto
	1, // 1: dealer.Dealer.Add:input_type -> dealer.DealerDto
	1, // 2: dealer.Dealer.Delete:input_type -> dealer.DealerDto
	1, // 3: dealer.Dealer.Update:input_type -> dealer.DealerDto
	5, // 4: dealer.Dealer.List:input_type -> common.Page
	5, // 5: dealer.Dealer.SyncDealer:input_type -> common.Page
	6, // 6: dealer.Dealer.Get:input_type -> common.IdLocalDTO
	4, // 7: dealer.Dealer.GetListByIds:input_type -> dealer.IdsDto
	7, // 8: dealer.Dealer.Add:output_type -> common.Response
	7, // 9: dealer.Dealer.Delete:output_type -> common.Response
	7, // 10: dealer.Dealer.Update:output_type -> common.Response
	7, // 11: dealer.Dealer.List:output_type -> common.Response
	7, // 12: dealer.Dealer.SyncDealer:output_type -> common.Response
	7, // 13: dealer.Dealer.Get:output_type -> common.Response
	7, // 14: dealer.Dealer.GetListByIds:output_type -> common.Response
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dealer_dealer_proto_init() }
func file_dealer_dealer_proto_init() {
	if File_dealer_dealer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dealer_dealer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DealersDto); i {
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
		file_dealer_dealer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DealerDto); i {
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
		file_dealer_dealer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DealerListDTO); i {
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
		file_dealer_dealer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncDealerData); i {
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
		file_dealer_dealer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dealer_dealer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dealer_dealer_proto_goTypes,
		DependencyIndexes: file_dealer_dealer_proto_depIdxs,
		MessageInfos:      file_dealer_dealer_proto_msgTypes,
	}.Build()
	File_dealer_dealer_proto = out.File
	file_dealer_dealer_proto_rawDesc = nil
	file_dealer_dealer_proto_goTypes = nil
	file_dealer_dealer_proto_depIdxs = nil
}
