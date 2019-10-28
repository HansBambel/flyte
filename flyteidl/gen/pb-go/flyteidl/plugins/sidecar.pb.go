// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/sidecar.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "k8s.io/api/core/v1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A sidecar job brings up the desired pod_spec.
// The plugin executor is responsible for keeping the pod alive until the primary container terminates
// or the task itself times out.
type SidecarJob struct {
	PodSpec              *v1.PodSpec `protobuf:"bytes,1,opt,name=pod_spec,json=podSpec,proto3" json:"pod_spec,omitempty"`
	PrimaryContainerName string      `protobuf:"bytes,2,opt,name=primary_container_name,json=primaryContainerName,proto3" json:"primary_container_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SidecarJob) Reset()         { *m = SidecarJob{} }
func (m *SidecarJob) String() string { return proto.CompactTextString(m) }
func (*SidecarJob) ProtoMessage()    {}
func (*SidecarJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccf0cecf5c059, []int{0}
}

func (m *SidecarJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SidecarJob.Unmarshal(m, b)
}
func (m *SidecarJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SidecarJob.Marshal(b, m, deterministic)
}
func (m *SidecarJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SidecarJob.Merge(m, src)
}
func (m *SidecarJob) XXX_Size() int {
	return xxx_messageInfo_SidecarJob.Size(m)
}
func (m *SidecarJob) XXX_DiscardUnknown() {
	xxx_messageInfo_SidecarJob.DiscardUnknown(m)
}

var xxx_messageInfo_SidecarJob proto.InternalMessageInfo

func (m *SidecarJob) GetPodSpec() *v1.PodSpec {
	if m != nil {
		return m.PodSpec
	}
	return nil
}

func (m *SidecarJob) GetPrimaryContainerName() string {
	if m != nil {
		return m.PrimaryContainerName
	}
	return ""
}

func init() {
	proto.RegisterType((*SidecarJob)(nil), "flyteidl.plugins.SidecarJob")
}

func init() { proto.RegisterFile("flyteidl/plugins/sidecar.proto", fileDescriptor_b34ccf0cecf5c059) }

var fileDescriptor_b34ccf0cecf5c059 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x3f, 0x4b, 0x03, 0x41,
	0x10, 0x47, 0x39, 0x0b, 0xff, 0xac, 0x8d, 0x1c, 0x22, 0x41, 0x41, 0x42, 0xaa, 0x34, 0xce, 0x10,
	0xa3, 0x62, 0xad, 0x9d, 0x85, 0x48, 0xd2, 0xd9, 0x1c, 0x7b, 0xbb, 0x93, 0x75, 0xf0, 0x6e, 0x67,
	0xd8, 0xdb, 0x04, 0xce, 0x4f, 0x2f, 0x7a, 0x39, 0x0b, 0xbb, 0x81, 0xc7, 0x3c, 0x7e, 0xcf, 0x5c,
	0x6f, 0x9a, 0x3e, 0x13, 0xfb, 0x06, 0xb5, 0xd9, 0x06, 0x8e, 0x1d, 0x76, 0xec, 0xc9, 0xd9, 0x04,
	0x9a, 0x24, 0x4b, 0x79, 0x36, 0x72, 0xd8, 0xf3, 0xcb, 0xd9, 0xe7, 0x63, 0x07, 0x2c, 0x68, 0x95,
	0xd1, 0x49, 0x22, 0xdc, 0x2d, 0x30, 0x50, 0xa4, 0x64, 0x33, 0xf9, 0xe1, 0x6b, 0xf6, 0x65, 0xcc,
	0x7a, 0xd0, 0xbc, 0x48, 0x5d, 0x3e, 0x98, 0x63, 0x15, 0x5f, 0x75, 0x4a, 0x6e, 0x52, 0x4c, 0x8b,
	0xf9, 0xe9, 0xed, 0x15, 0x0c, 0x12, 0xb0, 0xca, 0xf0, 0x23, 0x81, 0xdd, 0x02, 0xde, 0xc4, 0xaf,
	0x95, 0xdc, 0xea, 0x48, 0x87, 0xa3, 0xbc, 0x33, 0x17, 0x9a, 0xb8, 0xb5, 0xa9, 0xaf, 0x9c, 0xc4,
	0x6c, 0x39, 0x52, 0xaa, 0xa2, 0x6d, 0x69, 0x72, 0x30, 0x2d, 0xe6, 0x27, 0xab, 0xf3, 0x3d, 0x7d,
	0x1e, 0xe1, 0xab, 0x6d, 0xe9, 0xe9, 0xfe, 0x7d, 0x19, 0x38, 0x7f, 0x6c, 0x6b, 0x70, 0xd2, 0x62,
	0xd3, 0x6f, 0x32, 0xfe, 0x35, 0x06, 0x8a, 0xa8, 0xf5, 0x4d, 0x10, 0xfc, 0x9f, 0x5d, 0x1f, 0xfe,
	0x2e, 0x5f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x71, 0x89, 0x82, 0x45, 0x11, 0x01, 0x00, 0x00,
}
