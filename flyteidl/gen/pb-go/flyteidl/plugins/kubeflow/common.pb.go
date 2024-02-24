// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/plugins/kubeflow/common.proto

package plugins

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

type RestartPolicy int32

const (
	RestartPolicy_RESTART_POLICY_NEVER      RestartPolicy = 0
	RestartPolicy_RESTART_POLICY_ON_FAILURE RestartPolicy = 1
	RestartPolicy_RESTART_POLICY_ALWAYS     RestartPolicy = 2
)

// Enum value maps for RestartPolicy.
var (
	RestartPolicy_name = map[int32]string{
		0: "RESTART_POLICY_NEVER",
		1: "RESTART_POLICY_ON_FAILURE",
		2: "RESTART_POLICY_ALWAYS",
	}
	RestartPolicy_value = map[string]int32{
		"RESTART_POLICY_NEVER":      0,
		"RESTART_POLICY_ON_FAILURE": 1,
		"RESTART_POLICY_ALWAYS":     2,
	}
)

func (x RestartPolicy) Enum() *RestartPolicy {
	p := new(RestartPolicy)
	*p = x
	return p
}

func (x RestartPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RestartPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_flyteidl_plugins_kubeflow_common_proto_enumTypes[0].Descriptor()
}

func (RestartPolicy) Type() protoreflect.EnumType {
	return &file_flyteidl_plugins_kubeflow_common_proto_enumTypes[0]
}

func (x RestartPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RestartPolicy.Descriptor instead.
func (RestartPolicy) EnumDescriptor() ([]byte, []int) {
	return file_flyteidl_plugins_kubeflow_common_proto_rawDescGZIP(), []int{0}
}

type CleanPodPolicy int32

const (
	CleanPodPolicy_CLEANPOD_POLICY_NONE    CleanPodPolicy = 0
	CleanPodPolicy_CLEANPOD_POLICY_RUNNING CleanPodPolicy = 1
	CleanPodPolicy_CLEANPOD_POLICY_ALL     CleanPodPolicy = 2
)

// Enum value maps for CleanPodPolicy.
var (
	CleanPodPolicy_name = map[int32]string{
		0: "CLEANPOD_POLICY_NONE",
		1: "CLEANPOD_POLICY_RUNNING",
		2: "CLEANPOD_POLICY_ALL",
	}
	CleanPodPolicy_value = map[string]int32{
		"CLEANPOD_POLICY_NONE":    0,
		"CLEANPOD_POLICY_RUNNING": 1,
		"CLEANPOD_POLICY_ALL":     2,
	}
)

func (x CleanPodPolicy) Enum() *CleanPodPolicy {
	p := new(CleanPodPolicy)
	*p = x
	return p
}

func (x CleanPodPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CleanPodPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_flyteidl_plugins_kubeflow_common_proto_enumTypes[1].Descriptor()
}

func (CleanPodPolicy) Type() protoreflect.EnumType {
	return &file_flyteidl_plugins_kubeflow_common_proto_enumTypes[1]
}

func (x CleanPodPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CleanPodPolicy.Descriptor instead.
func (CleanPodPolicy) EnumDescriptor() ([]byte, []int) {
	return file_flyteidl_plugins_kubeflow_common_proto_rawDescGZIP(), []int{1}
}

type RunPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the policy to kill pods after the job completes. Default to None.
	CleanPodPolicy CleanPodPolicy `protobuf:"varint,1,opt,name=clean_pod_policy,json=cleanPodPolicy,proto3,enum=flyteidl.plugins.kubeflow.CleanPodPolicy" json:"clean_pod_policy,omitempty"`
	// TTL to clean up jobs. Default to infinite.
	TtlSecondsAfterFinished int32 `protobuf:"varint,2,opt,name=ttl_seconds_after_finished,json=ttlSecondsAfterFinished,proto3" json:"ttl_seconds_after_finished,omitempty"`
	// Specifies the duration in seconds relative to the startTime that the job may be active
	// before the system tries to terminate it; value must be positive integer.
	ActiveDeadlineSeconds int32 `protobuf:"varint,3,opt,name=active_deadline_seconds,json=activeDeadlineSeconds,proto3" json:"active_deadline_seconds,omitempty"`
	// Number of retries before marking this job failed.
	BackoffLimit int32 `protobuf:"varint,4,opt,name=backoff_limit,json=backoffLimit,proto3" json:"backoff_limit,omitempty"`
}

func (x *RunPolicy) Reset() {
	*x = RunPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_plugins_kubeflow_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunPolicy) ProtoMessage() {}

func (x *RunPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_plugins_kubeflow_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunPolicy.ProtoReflect.Descriptor instead.
func (*RunPolicy) Descriptor() ([]byte, []int) {
	return file_flyteidl_plugins_kubeflow_common_proto_rawDescGZIP(), []int{0}
}

func (x *RunPolicy) GetCleanPodPolicy() CleanPodPolicy {
	if x != nil {
		return x.CleanPodPolicy
	}
	return CleanPodPolicy_CLEANPOD_POLICY_NONE
}

func (x *RunPolicy) GetTtlSecondsAfterFinished() int32 {
	if x != nil {
		return x.TtlSecondsAfterFinished
	}
	return 0
}

func (x *RunPolicy) GetActiveDeadlineSeconds() int32 {
	if x != nil {
		return x.ActiveDeadlineSeconds
	}
	return 0
}

func (x *RunPolicy) GetBackoffLimit() int32 {
	if x != nil {
		return x.BackoffLimit
	}
	return 0
}

var File_flyteidl_plugins_kubeflow_common_proto protoreflect.FileDescriptor

var file_flyteidl_plugins_kubeflow_common_proto_rawDesc = []byte{
	0x0a, 0x26, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x66,
	0x6c, 0x6f, 0x77, 0x22, 0xfa, 0x01, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x53, 0x0a, 0x10, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x64, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x50, 0x6f, 0x64,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0e, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x50, 0x6f, 0x64,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3b, 0x0a, 0x1a, 0x74, 0x74, 0x6c, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x74, 0x74, 0x6c, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x41, 0x66, 0x74, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x62,
	0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x2a, 0x63, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x50, 0x4f, 0x4c,
	0x49, 0x43, 0x59, 0x5f, 0x4e, 0x45, 0x56, 0x45, 0x52, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x52,
	0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x4f, 0x4e,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x41, 0x4c, 0x57,
	0x41, 0x59, 0x53, 0x10, 0x02, 0x2a, 0x60, 0x0a, 0x0e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x50, 0x6f,
	0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4c, 0x45, 0x41, 0x4e,
	0x50, 0x4f, 0x44, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4c, 0x45, 0x41, 0x4e, 0x50, 0x4f, 0x44, 0x5f, 0x50, 0x4f,
	0x4c, 0x49, 0x43, 0x59, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x17,
	0x0a, 0x13, 0x43, 0x4c, 0x45, 0x41, 0x4e, 0x50, 0x4f, 0x44, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43,
	0x59, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x42, 0xf1, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x46, 0x50, 0x4b, 0xaa, 0x02, 0x19,
	0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x19, 0x46, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x5c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x5c, 0x4b, 0x75, 0x62,
	0x65, 0x66, 0x6c, 0x6f, 0x77, 0xe2, 0x02, 0x25, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x5c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f,
	0x77, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b,
	0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_plugins_kubeflow_common_proto_rawDescOnce sync.Once
	file_flyteidl_plugins_kubeflow_common_proto_rawDescData = file_flyteidl_plugins_kubeflow_common_proto_rawDesc
)

func file_flyteidl_plugins_kubeflow_common_proto_rawDescGZIP() []byte {
	file_flyteidl_plugins_kubeflow_common_proto_rawDescOnce.Do(func() {
		file_flyteidl_plugins_kubeflow_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_plugins_kubeflow_common_proto_rawDescData)
	})
	return file_flyteidl_plugins_kubeflow_common_proto_rawDescData
}

var file_flyteidl_plugins_kubeflow_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_flyteidl_plugins_kubeflow_common_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flyteidl_plugins_kubeflow_common_proto_goTypes = []interface{}{
	(RestartPolicy)(0),  // 0: flyteidl.plugins.kubeflow.RestartPolicy
	(CleanPodPolicy)(0), // 1: flyteidl.plugins.kubeflow.CleanPodPolicy
	(*RunPolicy)(nil),   // 2: flyteidl.plugins.kubeflow.RunPolicy
}
var file_flyteidl_plugins_kubeflow_common_proto_depIdxs = []int32{
	1, // 0: flyteidl.plugins.kubeflow.RunPolicy.clean_pod_policy:type_name -> flyteidl.plugins.kubeflow.CleanPodPolicy
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flyteidl_plugins_kubeflow_common_proto_init() }
func file_flyteidl_plugins_kubeflow_common_proto_init() {
	if File_flyteidl_plugins_kubeflow_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_plugins_kubeflow_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunPolicy); i {
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
			RawDescriptor: file_flyteidl_plugins_kubeflow_common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_plugins_kubeflow_common_proto_goTypes,
		DependencyIndexes: file_flyteidl_plugins_kubeflow_common_proto_depIdxs,
		EnumInfos:         file_flyteidl_plugins_kubeflow_common_proto_enumTypes,
		MessageInfos:      file_flyteidl_plugins_kubeflow_common_proto_msgTypes,
	}.Build()
	File_flyteidl_plugins_kubeflow_common_proto = out.File
	file_flyteidl_plugins_kubeflow_common_proto_rawDesc = nil
	file_flyteidl_plugins_kubeflow_common_proto_goTypes = nil
	file_flyteidl_plugins_kubeflow_common_proto_depIdxs = nil
}
