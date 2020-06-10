// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: location/location.proto

package location

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type LocationDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Code        string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	MapUrl      string `protobuf:"bytes,5,opt,name=MapUrl,proto3" json:"MapUrl,omitempty"`
	Latitude    string `protobuf:"bytes,6,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude   string `protobuf:"bytes,7,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	CreateTime  int64  `protobuf:"varint,8,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime  int64  `protobuf:"varint,9,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *LocationDto) Reset() {
	*x = LocationDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationDto) ProtoMessage() {}

func (x *LocationDto) ProtoReflect() protoreflect.Message {
	mi := &file_location_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationDto.ProtoReflect.Descriptor instead.
func (*LocationDto) Descriptor() ([]byte, []int) {
	return file_location_location_proto_rawDescGZIP(), []int{0}
}

func (x *LocationDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LocationDto) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *LocationDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LocationDto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *LocationDto) GetMapUrl() string {
	if x != nil {
		return x.MapUrl
	}
	return ""
}

func (x *LocationDto) GetLatitude() string {
	if x != nil {
		return x.Latitude
	}
	return ""
}

func (x *LocationDto) GetLongitude() string {
	if x != nil {
		return x.Longitude
	}
	return ""
}

func (x *LocationDto) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *LocationDto) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

type PagedList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNo   uint32     `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize uint32     `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	Total    uint32     `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	List     []*any.Any `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PagedList) Reset() {
	*x = PagedList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagedList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagedList) ProtoMessage() {}

func (x *PagedList) ProtoReflect() protoreflect.Message {
	mi := &file_location_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagedList.ProtoReflect.Descriptor instead.
func (*PagedList) Descriptor() ([]byte, []int) {
	return file_location_location_proto_rawDescGZIP(), []int{1}
}

func (x *PagedList) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *PagedList) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PagedList) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PagedList) GetList() []*any.Any {
	if x != nil {
		return x.List
	}
	return nil
}

var File_location_location_proto protoreflect.FileDescriptor

var file_location_location_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x0b, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x61, 0x70, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x4c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x7f, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xd0, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x15, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x15, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x28, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62,
	0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_location_location_proto_rawDescOnce sync.Once
	file_location_location_proto_rawDescData = file_location_location_proto_rawDesc
)

func file_location_location_proto_rawDescGZIP() []byte {
	file_location_location_proto_rawDescOnce.Do(func() {
		file_location_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_location_location_proto_rawDescData)
	})
	return file_location_location_proto_rawDescData
}

var file_location_location_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_location_location_proto_goTypes = []interface{}{
	(*LocationDto)(nil),     // 0: location.LocationDto
	(*PagedList)(nil),       // 1: location.PagedList
	(*any.Any)(nil),         // 2: google.protobuf.Any
	(*common.Page)(nil),     // 3: common.Page
	(*common.Response)(nil), // 4: common.Response
}
var file_location_location_proto_depIdxs = []int32{
	2, // 0: location.PagedList.list:type_name -> google.protobuf.Any
	0, // 1: location.Location.Add:input_type -> location.LocationDto
	0, // 2: location.Location.Delete:input_type -> location.LocationDto
	0, // 3: location.Location.Update:input_type -> location.LocationDto
	3, // 4: location.Location.List:input_type -> common.Page
	4, // 5: location.Location.Add:output_type -> common.Response
	4, // 6: location.Location.Delete:output_type -> common.Response
	4, // 7: location.Location.Update:output_type -> common.Response
	4, // 8: location.Location.List:output_type -> common.Response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_location_location_proto_init() }
func file_location_location_proto_init() {
	if File_location_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_location_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationDto); i {
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
		file_location_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagedList); i {
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
			RawDescriptor: file_location_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_location_location_proto_goTypes,
		DependencyIndexes: file_location_location_proto_depIdxs,
		MessageInfos:      file_location_location_proto_msgTypes,
	}.Build()
	File_location_location_proto = out.File
	file_location_location_proto_rawDesc = nil
	file_location_location_proto_goTypes = nil
	file_location_location_proto_depIdxs = nil
}
