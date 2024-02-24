// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/admin/version.proto

package admin

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Response for the GetVersion API
type GetVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The control plane version information. FlyteAdmin and related components
	// form the control plane of Flyte
	ControlPlaneVersion *Version `protobuf:"bytes,1,opt,name=control_plane_version,json=controlPlaneVersion,proto3" json:"control_plane_version,omitempty"`
}

func (x *GetVersionResponse) Reset() {
	*x = GetVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionResponse) ProtoMessage() {}

func (x *GetVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_version_proto_msgTypes[0]
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
	return file_flyteidl_admin_version_proto_rawDescGZIP(), []int{0}
}

func (x *GetVersionResponse) GetControlPlaneVersion() *Version {
	if x != nil {
		return x.ControlPlaneVersion
	}
	return nil
}

// Provides Version information for a component
type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the GIT sha of the build
	Build string `protobuf:"bytes,1,opt,name=Build,proto3" json:"Build,omitempty"`
	// Version for the build, should follow a semver
	Version string `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
	// Build timestamp
	BuildTime string `protobuf:"bytes,3,opt,name=BuildTime,proto3" json:"BuildTime,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_version_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_version_proto_msgTypes[1]
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
	return file_flyteidl_admin_version_proto_rawDescGZIP(), []int{1}
}

func (x *Version) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

func (x *Version) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Version) GetBuildTime() string {
	if x != nil {
		return x.BuildTime
	}
	return ""
}

// Empty request for GetVersion
type GetVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetVersionRequest) Reset() {
	*x = GetVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_version_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionRequest) ProtoMessage() {}

func (x *GetVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_version_proto_msgTypes[2]
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
	return file_flyteidl_admin_version_proto_rawDescGZIP(), []int{2}
}

var File_flyteidl_admin_version_proto protoreflect.FileDescriptor

var file_flyteidl_admin_version_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x61,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x57, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0xb8, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0xa2, 0x02, 0x03, 0x46, 0x41, 0x58, 0xaa, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x0e, 0x46, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x1a, 0x46, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x46, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_flyteidl_admin_version_proto_rawDescOnce sync.Once
	file_flyteidl_admin_version_proto_rawDescData = file_flyteidl_admin_version_proto_rawDesc
)

func file_flyteidl_admin_version_proto_rawDescGZIP() []byte {
	file_flyteidl_admin_version_proto_rawDescOnce.Do(func() {
		file_flyteidl_admin_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_admin_version_proto_rawDescData)
	})
	return file_flyteidl_admin_version_proto_rawDescData
}

var file_flyteidl_admin_version_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_flyteidl_admin_version_proto_goTypes = []interface{}{
	(*GetVersionResponse)(nil), // 0: flyteidl.admin.GetVersionResponse
	(*Version)(nil),            // 1: flyteidl.admin.Version
	(*GetVersionRequest)(nil),  // 2: flyteidl.admin.GetVersionRequest
}
var file_flyteidl_admin_version_proto_depIdxs = []int32{
	1, // 0: flyteidl.admin.GetVersionResponse.control_plane_version:type_name -> flyteidl.admin.Version
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flyteidl_admin_version_proto_init() }
func file_flyteidl_admin_version_proto_init() {
	if File_flyteidl_admin_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_admin_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_flyteidl_admin_version_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_flyteidl_admin_version_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_admin_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_admin_version_proto_goTypes,
		DependencyIndexes: file_flyteidl_admin_version_proto_depIdxs,
		MessageInfos:      file_flyteidl_admin_version_proto_msgTypes,
	}.Build()
	File_flyteidl_admin_version_proto = out.File
	file_flyteidl_admin_version_proto_rawDesc = nil
	file_flyteidl_admin_version_proto_goTypes = nil
	file_flyteidl_admin_version_proto_depIdxs = nil
}
