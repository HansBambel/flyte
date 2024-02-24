// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/admin/project_attributes.proto

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

// Defines a set of custom matching attributes at the project level.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique project id for which this set of attributes will be applied.
	Project            string              `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	MatchingAttributes *MatchingAttributes `protobuf:"bytes,2,opt,name=matching_attributes,json=matchingAttributes,proto3" json:"matching_attributes,omitempty"`
	// Optional, org key applied to the project.
	Org string `protobuf:"bytes,3,opt,name=org,proto3" json:"org,omitempty"`
}

func (x *ProjectAttributes) Reset() {
	*x = ProjectAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAttributes) ProtoMessage() {}

func (x *ProjectAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAttributes.ProtoReflect.Descriptor instead.
func (*ProjectAttributes) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_attributes_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectAttributes) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ProjectAttributes) GetMatchingAttributes() *MatchingAttributes {
	if x != nil {
		return x.MatchingAttributes
	}
	return nil
}

func (x *ProjectAttributes) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

// Sets custom attributes for a project
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectAttributesUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +required
	Attributes *ProjectAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *ProjectAttributesUpdateRequest) Reset() {
	*x = ProjectAttributesUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectAttributesUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAttributesUpdateRequest) ProtoMessage() {}

func (x *ProjectAttributesUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAttributesUpdateRequest.ProtoReflect.Descriptor instead.
func (*ProjectAttributesUpdateRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_attributes_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectAttributesUpdateRequest) GetAttributes() *ProjectAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// Purposefully empty, may be populated in the future.
type ProjectAttributesUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProjectAttributesUpdateResponse) Reset() {
	*x = ProjectAttributesUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectAttributesUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAttributesUpdateResponse) ProtoMessage() {}

func (x *ProjectAttributesUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAttributesUpdateResponse.ProtoReflect.Descriptor instead.
func (*ProjectAttributesUpdateResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_attributes_proto_rawDescGZIP(), []int{2}
}

// Request to get an individual project level attribute override.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectAttributesGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique project id which this set of attributes references.
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Which type of matchable attributes to return.
	// +required
	ResourceType MatchableResource `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	// Optional, org key applied to the project.
	Org string `protobuf:"bytes,3,opt,name=org,proto3" json:"org,omitempty"`
}

func (x *ProjectAttributesGetRequest) Reset() {
	*x = ProjectAttributesGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectAttributesGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAttributesGetRequest) ProtoMessage() {}

func (x *ProjectAttributesGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAttributesGetRequest.ProtoReflect.Descriptor instead.
func (*ProjectAttributesGetRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_attributes_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectAttributesGetRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ProjectAttributesGetRequest) GetResourceType() MatchableResource {
	if x != nil {
		return x.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

func (x *ProjectAttributesGetRequest) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

// Response to get an individual project level attribute override.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectAttributesGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes *ProjectAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *ProjectAttributesGetResponse) Reset() {
	*x = ProjectAttributesGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectAttributesGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAttributesGetResponse) ProtoMessage() {}

func (x *ProjectAttributesGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAttributesGetResponse.ProtoReflect.Descriptor instead.
func (*ProjectAttributesGetResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_attributes_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectAttributesGetResponse) GetAttributes() *ProjectAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// Request to delete a set matchable project level attribute override.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectAttributesDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique project id which this set of attributes references.
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Which type of matchable attributes to delete.
	// +required
	ResourceType MatchableResource `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	// Optional, org key applied to the project.
	Org string `protobuf:"bytes,3,opt,name=org,proto3" json:"org,omitempty"`
}

func (x *ProjectAttributesDeleteRequest) Reset() {
	*x = ProjectAttributesDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectAttributesDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAttributesDeleteRequest) ProtoMessage() {}

func (x *ProjectAttributesDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAttributesDeleteRequest.ProtoReflect.Descriptor instead.
func (*ProjectAttributesDeleteRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_attributes_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectAttributesDeleteRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ProjectAttributesDeleteRequest) GetResourceType() MatchableResource {
	if x != nil {
		return x.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

func (x *ProjectAttributesDeleteRequest) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

// Purposefully empty, may be populated in the future.
type ProjectAttributesDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProjectAttributesDeleteResponse) Reset() {
	*x = ProjectAttributesDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectAttributesDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectAttributesDeleteResponse) ProtoMessage() {}

func (x *ProjectAttributesDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_attributes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectAttributesDeleteResponse.ProtoReflect.Descriptor instead.
func (*ProjectAttributesDeleteResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_attributes_proto_rawDescGZIP(), []int{6}
}

var File_flyteidl_admin_project_attributes_proto protoreflect.FileDescriptor

var file_flyteidl_admin_project_attributes_proto_rawDesc = []byte{
	0x0a, 0x27, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x27, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x53, 0x0a, 0x13, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x5f, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x52, 0x12, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x22, 0x63, 0x0a, 0x1e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x21,
	0x0a, 0x1f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x91, 0x01, 0x0a, 0x1b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x46, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6f, 0x72, 0x67, 0x22, 0x61, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x1e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x46, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6f, 0x72, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x22,
	0x21, 0x0a, 0x1f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0xc2, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x16, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67,
	0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0xa2, 0x02, 0x03, 0x46, 0x41, 0x58, 0xaa, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x1a, 0x46, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_admin_project_attributes_proto_rawDescOnce sync.Once
	file_flyteidl_admin_project_attributes_proto_rawDescData = file_flyteidl_admin_project_attributes_proto_rawDesc
)

func file_flyteidl_admin_project_attributes_proto_rawDescGZIP() []byte {
	file_flyteidl_admin_project_attributes_proto_rawDescOnce.Do(func() {
		file_flyteidl_admin_project_attributes_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_admin_project_attributes_proto_rawDescData)
	})
	return file_flyteidl_admin_project_attributes_proto_rawDescData
}

var file_flyteidl_admin_project_attributes_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_flyteidl_admin_project_attributes_proto_goTypes = []interface{}{
	(*ProjectAttributes)(nil),               // 0: flyteidl.admin.ProjectAttributes
	(*ProjectAttributesUpdateRequest)(nil),  // 1: flyteidl.admin.ProjectAttributesUpdateRequest
	(*ProjectAttributesUpdateResponse)(nil), // 2: flyteidl.admin.ProjectAttributesUpdateResponse
	(*ProjectAttributesGetRequest)(nil),     // 3: flyteidl.admin.ProjectAttributesGetRequest
	(*ProjectAttributesGetResponse)(nil),    // 4: flyteidl.admin.ProjectAttributesGetResponse
	(*ProjectAttributesDeleteRequest)(nil),  // 5: flyteidl.admin.ProjectAttributesDeleteRequest
	(*ProjectAttributesDeleteResponse)(nil), // 6: flyteidl.admin.ProjectAttributesDeleteResponse
	(*MatchingAttributes)(nil),              // 7: flyteidl.admin.MatchingAttributes
	(MatchableResource)(0),                  // 8: flyteidl.admin.MatchableResource
}
var file_flyteidl_admin_project_attributes_proto_depIdxs = []int32{
	7, // 0: flyteidl.admin.ProjectAttributes.matching_attributes:type_name -> flyteidl.admin.MatchingAttributes
	0, // 1: flyteidl.admin.ProjectAttributesUpdateRequest.attributes:type_name -> flyteidl.admin.ProjectAttributes
	8, // 2: flyteidl.admin.ProjectAttributesGetRequest.resource_type:type_name -> flyteidl.admin.MatchableResource
	0, // 3: flyteidl.admin.ProjectAttributesGetResponse.attributes:type_name -> flyteidl.admin.ProjectAttributes
	8, // 4: flyteidl.admin.ProjectAttributesDeleteRequest.resource_type:type_name -> flyteidl.admin.MatchableResource
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_flyteidl_admin_project_attributes_proto_init() }
func file_flyteidl_admin_project_attributes_proto_init() {
	if File_flyteidl_admin_project_attributes_proto != nil {
		return
	}
	file_flyteidl_admin_matchable_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_admin_project_attributes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectAttributes); i {
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
		file_flyteidl_admin_project_attributes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectAttributesUpdateRequest); i {
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
		file_flyteidl_admin_project_attributes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectAttributesUpdateResponse); i {
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
		file_flyteidl_admin_project_attributes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectAttributesGetRequest); i {
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
		file_flyteidl_admin_project_attributes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectAttributesGetResponse); i {
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
		file_flyteidl_admin_project_attributes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectAttributesDeleteRequest); i {
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
		file_flyteidl_admin_project_attributes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectAttributesDeleteResponse); i {
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
			RawDescriptor: file_flyteidl_admin_project_attributes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_admin_project_attributes_proto_goTypes,
		DependencyIndexes: file_flyteidl_admin_project_attributes_proto_depIdxs,
		MessageInfos:      file_flyteidl_admin_project_attributes_proto_msgTypes,
	}.Build()
	File_flyteidl_admin_project_attributes_proto = out.File
	file_flyteidl_admin_project_attributes_proto_rawDesc = nil
	file_flyteidl_admin_project_attributes_proto_goTypes = nil
	file_flyteidl_admin_project_attributes_proto_depIdxs = nil
}
