// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: yubikey/v1/yubikey.proto

package yubikey

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major uint32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor uint32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch uint32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yubikey_v1_yubikey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_yubikey_v1_yubikey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_yubikey_v1_yubikey_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetMajor() uint32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() uint32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Version) GetPatch() uint32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

type GetSerialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSerialsRequest) Reset() {
	*x = GetSerialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yubikey_v1_yubikey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSerialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSerialsRequest) ProtoMessage() {}

func (x *GetSerialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yubikey_v1_yubikey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSerialsRequest.ProtoReflect.Descriptor instead.
func (*GetSerialsRequest) Descriptor() ([]byte, []int) {
	return file_yubikey_v1_yubikey_proto_rawDescGZIP(), []int{1}
}

type GetSerialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serials []uint32 `protobuf:"varint,1,rep,packed,name=serials,proto3" json:"serials,omitempty"`
}

func (x *GetSerialsResponse) Reset() {
	*x = GetSerialsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yubikey_v1_yubikey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSerialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSerialsResponse) ProtoMessage() {}

func (x *GetSerialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yubikey_v1_yubikey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSerialsResponse.ProtoReflect.Descriptor instead.
func (*GetSerialsResponse) Descriptor() ([]byte, []int) {
	return file_yubikey_v1_yubikey_proto_rawDescGZIP(), []int{2}
}

func (x *GetSerialsResponse) GetSerials() []uint32 {
	if x != nil {
		return x.Serials
	}
	return nil
}

type GetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serial uint32 `protobuf:"varint,1,opt,name=serial,proto3" json:"serial,omitempty"`
}

func (x *GetVersionRequest) Reset() {
	*x = GetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yubikey_v1_yubikey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionRequest) ProtoMessage() {}

func (x *GetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yubikey_v1_yubikey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionRequest.ProtoReflect.Descriptor instead.
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return file_yubikey_v1_yubikey_proto_rawDescGZIP(), []int{3}
}

func (x *GetVersionRequest) GetSerial() uint32 {
	if x != nil {
		return x.Serial
	}
	return 0
}

type GetVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetVersionResponse) Reset() {
	*x = GetVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yubikey_v1_yubikey_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionResponse) ProtoMessage() {}

func (x *GetVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yubikey_v1_yubikey_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionResponse.ProtoReflect.Descriptor instead.
func (*GetVersionResponse) Descriptor() ([]byte, []int) {
	return file_yubikey_v1_yubikey_proto_rawDescGZIP(), []int{4}
}

func (x *GetVersionResponse) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

var File_yubikey_v1_yubikey_proto protoreflect.FileDescriptor

var file_yubikey_v1_yubikey_proto_rawDesc = []byte{
	0x0a, 0x18, 0x79, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x79, 0x75, 0x62,
	0x69, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x79, 0x75, 0x62, 0x69,
	0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x22, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x79, 0x75, 0x62,
	0x69, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xed, 0x01, 0x0a, 0x0e, 0x59, 0x75, 0x62,
	0x69, 0x6b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1d, 0x2e, 0x79, 0x75, 0x62, 0x69,
	0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x79, 0x75, 0x62, 0x69, 0x6b,
	0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x79, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x2f, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x71, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x79, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x79, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f,
	0x79, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f,
	0x7b, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x7d, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6b, 0x75, 0x6d, 0x69, 0x6e, 0x2f, 0x79,
	0x6b, 0x6d, 0x67, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x79, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79,
	0x2f, 0x76, 0x31, 0x3b, 0x79, 0x75, 0x62, 0x69, 0x6b, 0x65, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_yubikey_v1_yubikey_proto_rawDescOnce sync.Once
	file_yubikey_v1_yubikey_proto_rawDescData = file_yubikey_v1_yubikey_proto_rawDesc
)

func file_yubikey_v1_yubikey_proto_rawDescGZIP() []byte {
	file_yubikey_v1_yubikey_proto_rawDescOnce.Do(func() {
		file_yubikey_v1_yubikey_proto_rawDescData = protoimpl.X.CompressGZIP(file_yubikey_v1_yubikey_proto_rawDescData)
	})
	return file_yubikey_v1_yubikey_proto_rawDescData
}

var file_yubikey_v1_yubikey_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yubikey_v1_yubikey_proto_goTypes = []interface{}{
	(*Version)(nil),            // 0: yubikey.v1.Version
	(*GetSerialsRequest)(nil),  // 1: yubikey.v1.GetSerialsRequest
	(*GetSerialsResponse)(nil), // 2: yubikey.v1.GetSerialsResponse
	(*GetVersionRequest)(nil),  // 3: yubikey.v1.GetVersionRequest
	(*GetVersionResponse)(nil), // 4: yubikey.v1.GetVersionResponse
}
var file_yubikey_v1_yubikey_proto_depIdxs = []int32{
	0, // 0: yubikey.v1.GetVersionResponse.version:type_name -> yubikey.v1.Version
	1, // 1: yubikey.v1.YubikeyService.GetSerials:input_type -> yubikey.v1.GetSerialsRequest
	3, // 2: yubikey.v1.YubikeyService.GetVersion:input_type -> yubikey.v1.GetVersionRequest
	2, // 3: yubikey.v1.YubikeyService.GetSerials:output_type -> yubikey.v1.GetSerialsResponse
	4, // 4: yubikey.v1.YubikeyService.GetVersion:output_type -> yubikey.v1.GetVersionResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yubikey_v1_yubikey_proto_init() }
func file_yubikey_v1_yubikey_proto_init() {
	if File_yubikey_v1_yubikey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yubikey_v1_yubikey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_yubikey_v1_yubikey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSerialsRequest); i {
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
		file_yubikey_v1_yubikey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSerialsResponse); i {
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
		file_yubikey_v1_yubikey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionRequest); i {
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
		file_yubikey_v1_yubikey_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionResponse); i {
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
			RawDescriptor: file_yubikey_v1_yubikey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yubikey_v1_yubikey_proto_goTypes,
		DependencyIndexes: file_yubikey_v1_yubikey_proto_depIdxs,
		MessageInfos:      file_yubikey_v1_yubikey_proto_msgTypes,
	}.Build()
	File_yubikey_v1_yubikey_proto = out.File
	file_yubikey_v1_yubikey_proto_rawDesc = nil
	file_yubikey_v1_yubikey_proto_goTypes = nil
	file_yubikey_v1_yubikey_proto_depIdxs = nil
}
