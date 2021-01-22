// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
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
	0x08, 0x61, 0x72, 0x65, 0x61, 0x44, 0x54, 0x4f, 0x73, 0x32, 0xf4, 0x01, 0x0a, 0x04, 0x41, 0x72,
	0x65, 0x61, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0f, 0x2e, 0x64, 0x65, 0x61, 0x6c,
	0x65, 0x72, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x49, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x41,
	0x72, 0x65, 0x61, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x54, 0x6f,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x64, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x2e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x72, 0x65, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_area_area_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_area_area_proto_goTypes = []interface{}{
	(*ParentCode)(nil),      // 0: dealer.ParentCode
	(*AreaDto)(nil),         // 1: dealer.AreaDto
	(*AreaList)(nil),        // 2: dealer.AreaList
	(*common.IdDto)(nil),    // 3: common.IdDto
	(*common.LocalDto)(nil), // 4: common.LocalDto
	(*common.Response)(nil), // 5: common.Response
}
var file_area_area_proto_depIdxs = []int32{
	1, // 0: dealer.AreaList.areaDTOs:type_name -> dealer.AreaDto
	1, // 1: dealer.Area.Add:input_type -> dealer.AreaDto
	3, // 2: dealer.Area.Delete:input_type -> common.IdDto
	1, // 3: dealer.Area.Update:input_type -> dealer.AreaDto
	4, // 4: dealer.Area.TopList:input_type -> common.LocalDto
	0, // 5: dealer.Area.ChildList:input_type -> dealer.ParentCode
	5, // 6: dealer.Area.Add:output_type -> common.Response
	5, // 7: dealer.Area.Delete:output_type -> common.Response
	5, // 8: dealer.Area.Update:output_type -> common.Response
	5, // 9: dealer.Area.TopList:output_type -> common.Response
	5, // 10: dealer.Area.ChildList:output_type -> common.Response
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_area_area_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
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
