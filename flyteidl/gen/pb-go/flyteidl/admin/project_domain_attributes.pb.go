// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/admin/project_domain_attributes.proto

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

// Defines a set of custom matching attributes which defines resource defaults for a project and domain.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectDomainAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique project id for which this set of attributes will be applied.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id for which this set of attributes will be applied.
	Domain             string              `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	MatchingAttributes *MatchingAttributes `protobuf:"bytes,3,opt,name=matching_attributes,json=matchingAttributes,proto3" json:"matching_attributes,omitempty"`
	// Optional, org key applied to the attributes.
	Org string `protobuf:"bytes,4,opt,name=org,proto3" json:"org,omitempty"`
}

func (x *ProjectDomainAttributes) Reset() {
	*x = ProjectDomainAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectDomainAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDomainAttributes) ProtoMessage() {}

func (x *ProjectDomainAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectDomainAttributes.ProtoReflect.Descriptor instead.
func (*ProjectDomainAttributes) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_domain_attributes_proto_rawDescGZIP(), []int{0}
}

func (x *ProjectDomainAttributes) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ProjectDomainAttributes) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ProjectDomainAttributes) GetMatchingAttributes() *MatchingAttributes {
	if x != nil {
		return x.MatchingAttributes
	}
	return nil
}

func (x *ProjectDomainAttributes) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

// Sets custom attributes for a project-domain combination.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectDomainAttributesUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +required
	Attributes *ProjectDomainAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *ProjectDomainAttributesUpdateRequest) Reset() {
	*x = ProjectDomainAttributesUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectDomainAttributesUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDomainAttributesUpdateRequest) ProtoMessage() {}

func (x *ProjectDomainAttributesUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectDomainAttributesUpdateRequest.ProtoReflect.Descriptor instead.
func (*ProjectDomainAttributesUpdateRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_domain_attributes_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectDomainAttributesUpdateRequest) GetAttributes() *ProjectDomainAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// Purposefully empty, may be populated in the future.
type ProjectDomainAttributesUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProjectDomainAttributesUpdateResponse) Reset() {
	*x = ProjectDomainAttributesUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectDomainAttributesUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDomainAttributesUpdateResponse) ProtoMessage() {}

func (x *ProjectDomainAttributesUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectDomainAttributesUpdateResponse.ProtoReflect.Descriptor instead.
func (*ProjectDomainAttributesUpdateResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_domain_attributes_proto_rawDescGZIP(), []int{2}
}

// Request to get an individual project domain attribute override.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectDomainAttributesGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique project id which this set of attributes references.
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id which this set of attributes references.
	// +required
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Which type of matchable attributes to return.
	// +required
	ResourceType MatchableResource `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	// Optional, org key applied to the attributes.
	Org string `protobuf:"bytes,4,opt,name=org,proto3" json:"org,omitempty"`
}

func (x *ProjectDomainAttributesGetRequest) Reset() {
	*x = ProjectDomainAttributesGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectDomainAttributesGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDomainAttributesGetRequest) ProtoMessage() {}

func (x *ProjectDomainAttributesGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectDomainAttributesGetRequest.ProtoReflect.Descriptor instead.
func (*ProjectDomainAttributesGetRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_domain_attributes_proto_rawDescGZIP(), []int{3}
}

func (x *ProjectDomainAttributesGetRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ProjectDomainAttributesGetRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ProjectDomainAttributesGetRequest) GetResourceType() MatchableResource {
	if x != nil {
		return x.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

func (x *ProjectDomainAttributesGetRequest) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

// Response to get an individual project domain attribute override.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectDomainAttributesGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes *ProjectDomainAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *ProjectDomainAttributesGetResponse) Reset() {
	*x = ProjectDomainAttributesGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectDomainAttributesGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDomainAttributesGetResponse) ProtoMessage() {}

func (x *ProjectDomainAttributesGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectDomainAttributesGetResponse.ProtoReflect.Descriptor instead.
func (*ProjectDomainAttributesGetResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_domain_attributes_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectDomainAttributesGetResponse) GetAttributes() *ProjectDomainAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// Request to delete a set matchable project domain attribute override.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectDomainAttributesDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique project id which this set of attributes references.
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id which this set of attributes references.
	// +required
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Which type of matchable attributes to delete.
	// +required
	ResourceType MatchableResource `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	// Optional, org key applied to the attributes.
	Org string `protobuf:"bytes,4,opt,name=org,proto3" json:"org,omitempty"`
}

func (x *ProjectDomainAttributesDeleteRequest) Reset() {
	*x = ProjectDomainAttributesDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectDomainAttributesDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDomainAttributesDeleteRequest) ProtoMessage() {}

func (x *ProjectDomainAttributesDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectDomainAttributesDeleteRequest.ProtoReflect.Descriptor instead.
func (*ProjectDomainAttributesDeleteRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_domain_attributes_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectDomainAttributesDeleteRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ProjectDomainAttributesDeleteRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ProjectDomainAttributesDeleteRequest) GetResourceType() MatchableResource {
	if x != nil {
		return x.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

func (x *ProjectDomainAttributesDeleteRequest) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

// Purposefully empty, may be populated in the future.
type ProjectDomainAttributesDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProjectDomainAttributesDeleteResponse) Reset() {
	*x = ProjectDomainAttributesDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectDomainAttributesDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectDomainAttributesDeleteResponse) ProtoMessage() {}

func (x *ProjectDomainAttributesDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_project_domain_attributes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectDomainAttributesDeleteResponse.ProtoReflect.Descriptor instead.
func (*ProjectDomainAttributesDeleteResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_project_domain_attributes_proto_rawDescGZIP(), []int{6}
}

var File_flyteidl_admin_project_domain_attributes_proto protoreflect.FileDescriptor

var file_flyteidl_admin_project_domain_attributes_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x1a, 0x27, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x17, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x53, 0x0a, 0x13, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x12, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x72, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x22, 0x6f,
	0x0a, 0x24, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22,
	0x27, 0x0a, 0x25, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x21, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x46, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x22, 0x6d, 0x0a, 0x22, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x47, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x24, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x46, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6f, 0x72, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x22, 0x27,
	0x0a, 0x25, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xc8, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x1c,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xa2, 0x02, 0x03, 0x46, 0x41,
	0x58, 0xaa, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0xca, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x1a, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0f, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_admin_project_domain_attributes_proto_rawDescOnce sync.Once
	file_flyteidl_admin_project_domain_attributes_proto_rawDescData = file_flyteidl_admin_project_domain_attributes_proto_rawDesc
)

func file_flyteidl_admin_project_domain_attributes_proto_rawDescGZIP() []byte {
	file_flyteidl_admin_project_domain_attributes_proto_rawDescOnce.Do(func() {
		file_flyteidl_admin_project_domain_attributes_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_admin_project_domain_attributes_proto_rawDescData)
	})
	return file_flyteidl_admin_project_domain_attributes_proto_rawDescData
}

var file_flyteidl_admin_project_domain_attributes_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_flyteidl_admin_project_domain_attributes_proto_goTypes = []interface{}{
	(*ProjectDomainAttributes)(nil),               // 0: flyteidl.admin.ProjectDomainAttributes
	(*ProjectDomainAttributesUpdateRequest)(nil),  // 1: flyteidl.admin.ProjectDomainAttributesUpdateRequest
	(*ProjectDomainAttributesUpdateResponse)(nil), // 2: flyteidl.admin.ProjectDomainAttributesUpdateResponse
	(*ProjectDomainAttributesGetRequest)(nil),     // 3: flyteidl.admin.ProjectDomainAttributesGetRequest
	(*ProjectDomainAttributesGetResponse)(nil),    // 4: flyteidl.admin.ProjectDomainAttributesGetResponse
	(*ProjectDomainAttributesDeleteRequest)(nil),  // 5: flyteidl.admin.ProjectDomainAttributesDeleteRequest
	(*ProjectDomainAttributesDeleteResponse)(nil), // 6: flyteidl.admin.ProjectDomainAttributesDeleteResponse
	(*MatchingAttributes)(nil),                    // 7: flyteidl.admin.MatchingAttributes
	(MatchableResource)(0),                        // 8: flyteidl.admin.MatchableResource
}
var file_flyteidl_admin_project_domain_attributes_proto_depIdxs = []int32{
	7, // 0: flyteidl.admin.ProjectDomainAttributes.matching_attributes:type_name -> flyteidl.admin.MatchingAttributes
	0, // 1: flyteidl.admin.ProjectDomainAttributesUpdateRequest.attributes:type_name -> flyteidl.admin.ProjectDomainAttributes
	8, // 2: flyteidl.admin.ProjectDomainAttributesGetRequest.resource_type:type_name -> flyteidl.admin.MatchableResource
	0, // 3: flyteidl.admin.ProjectDomainAttributesGetResponse.attributes:type_name -> flyteidl.admin.ProjectDomainAttributes
	8, // 4: flyteidl.admin.ProjectDomainAttributesDeleteRequest.resource_type:type_name -> flyteidl.admin.MatchableResource
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_flyteidl_admin_project_domain_attributes_proto_init() }
func file_flyteidl_admin_project_domain_attributes_proto_init() {
	if File_flyteidl_admin_project_domain_attributes_proto != nil {
		return
	}
	file_flyteidl_admin_matchable_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_admin_project_domain_attributes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectDomainAttributes); i {
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
		file_flyteidl_admin_project_domain_attributes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectDomainAttributesUpdateRequest); i {
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
		file_flyteidl_admin_project_domain_attributes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectDomainAttributesUpdateResponse); i {
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
		file_flyteidl_admin_project_domain_attributes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectDomainAttributesGetRequest); i {
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
		file_flyteidl_admin_project_domain_attributes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectDomainAttributesGetResponse); i {
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
		file_flyteidl_admin_project_domain_attributes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectDomainAttributesDeleteRequest); i {
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
		file_flyteidl_admin_project_domain_attributes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectDomainAttributesDeleteResponse); i {
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
			RawDescriptor: file_flyteidl_admin_project_domain_attributes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_admin_project_domain_attributes_proto_goTypes,
		DependencyIndexes: file_flyteidl_admin_project_domain_attributes_proto_depIdxs,
		MessageInfos:      file_flyteidl_admin_project_domain_attributes_proto_msgTypes,
	}.Build()
	File_flyteidl_admin_project_domain_attributes_proto = out.File
	file_flyteidl_admin_project_domain_attributes_proto_rawDesc = nil
	file_flyteidl_admin_project_domain_attributes_proto_goTypes = nil
	file_flyteidl_admin_project_domain_attributes_proto_depIdxs = nil
}
