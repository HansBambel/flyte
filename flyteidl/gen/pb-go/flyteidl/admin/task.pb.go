// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/admin/task.proto

package admin

import (
	reflect "reflect"
	sync "sync"

	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a request structure to create a revision of a task.
// See :ref:`ref_flyteidl.admin.Task` for more details
type TaskCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id represents the unique identifier of the task.
	// +required
	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Represents the specification for task.
	// +required
	Spec *TaskSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *TaskCreateRequest) Reset() {
	*x = TaskCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateRequest) ProtoMessage() {}

func (x *TaskCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateRequest.ProtoReflect.Descriptor instead.
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskCreateRequest) GetId() *core.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *TaskCreateRequest) GetSpec() *TaskSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// Represents a response structure if task creation succeeds.
type TaskCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TaskCreateResponse) Reset() {
	*x = TaskCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskCreateResponse) ProtoMessage() {}

func (x *TaskCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskCreateResponse.ProtoReflect.Descriptor instead.
func (*TaskCreateResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_task_proto_rawDescGZIP(), []int{1}
}

// Flyte workflows are composed of many ordered tasks. That is small, reusable, self-contained logical blocks
// arranged to process workflow inputs and produce a deterministic set of outputs.
// Tasks can come in many varieties tuned for specialized behavior.
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id represents the unique identifier of the task.
	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// closure encapsulates all the fields that maps to a compiled version of the task.
	Closure *TaskClosure `protobuf:"bytes,2,opt,name=closure,proto3" json:"closure,omitempty"`
	// One-liner overview of the entity.
	ShortDescription string `protobuf:"bytes,3,opt,name=short_description,json=shortDescription,proto3" json:"short_description,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_task_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetId() *core.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Task) GetClosure() *TaskClosure {
	if x != nil {
		return x.Closure
	}
	return nil
}

func (x *Task) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

// Represents a list of tasks returned from the admin.
// See :ref:`ref_flyteidl.admin.Task` for more details
type TaskList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of tasks returned based on the request.
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TaskList) Reset() {
	*x = TaskList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskList) ProtoMessage() {}

func (x *TaskList) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskList.ProtoReflect.Descriptor instead.
func (*TaskList) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskList) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *TaskList) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Represents a structure that encapsulates the user-configured specification of the task.
type TaskSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Template of the task that encapsulates all the metadata of the task.
	Template *core.TaskTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	// Represents the specification for description entity.
	Description *DescriptionEntity `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *TaskSpec) Reset() {
	*x = TaskSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSpec) ProtoMessage() {}

func (x *TaskSpec) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSpec.ProtoReflect.Descriptor instead.
func (*TaskSpec) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_task_proto_rawDescGZIP(), []int{4}
}

func (x *TaskSpec) GetTemplate() *core.TaskTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *TaskSpec) GetDescription() *DescriptionEntity {
	if x != nil {
		return x.Description
	}
	return nil
}

// Compute task attributes which include values derived from the TaskSpec, as well as plugin-specific data
// and task metadata.
type TaskClosure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents the compiled representation of the task from the specification provided.
	CompiledTask *core.CompiledTask `protobuf:"bytes,1,opt,name=compiled_task,json=compiledTask,proto3" json:"compiled_task,omitempty"`
	// Time at which the task was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *TaskClosure) Reset() {
	*x = TaskClosure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskClosure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskClosure) ProtoMessage() {}

func (x *TaskClosure) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskClosure.ProtoReflect.Descriptor instead.
func (*TaskClosure) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_task_proto_rawDescGZIP(), []int{5}
}

func (x *TaskClosure) GetCompiledTask() *core.CompiledTask {
	if x != nil {
		return x.CompiledTask
	}
	return nil
}

func (x *TaskClosure) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_flyteidl_admin_task_proto protoreflect.FileDescriptor

var file_flyteidl_admin_task_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1e, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c,
	0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x14, 0x0a, 0x12,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x29, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x07, 0x63, 0x6c, 0x6f, 0x73, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6c, 0x6f,
	0x73, 0x75, 0x72, 0x65, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x08, 0x54, 0x61,
	0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x08, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x70, 0x65, 0x63, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x43,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6c, 0x6f, 0x73,
	0x75, 0x72, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0xb5, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d,
	0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0xa2, 0x02, 0x03, 0x46, 0x41, 0x58, 0xaa, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x1a, 0x46, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x3a, 0x3a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_admin_task_proto_rawDescOnce sync.Once
	file_flyteidl_admin_task_proto_rawDescData = file_flyteidl_admin_task_proto_rawDesc
)

func file_flyteidl_admin_task_proto_rawDescGZIP() []byte {
	file_flyteidl_admin_task_proto_rawDescOnce.Do(func() {
		file_flyteidl_admin_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_admin_task_proto_rawDescData)
	})
	return file_flyteidl_admin_task_proto_rawDescData
}

var file_flyteidl_admin_task_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_flyteidl_admin_task_proto_goTypes = []interface{}{
	(*TaskCreateRequest)(nil),     // 0: flyteidl.admin.TaskCreateRequest
	(*TaskCreateResponse)(nil),    // 1: flyteidl.admin.TaskCreateResponse
	(*Task)(nil),                  // 2: flyteidl.admin.Task
	(*TaskList)(nil),              // 3: flyteidl.admin.TaskList
	(*TaskSpec)(nil),              // 4: flyteidl.admin.TaskSpec
	(*TaskClosure)(nil),           // 5: flyteidl.admin.TaskClosure
	(*core.Identifier)(nil),       // 6: flyteidl.core.Identifier
	(*core.TaskTemplate)(nil),     // 7: flyteidl.core.TaskTemplate
	(*DescriptionEntity)(nil),     // 8: flyteidl.admin.DescriptionEntity
	(*core.CompiledTask)(nil),     // 9: flyteidl.core.CompiledTask
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_flyteidl_admin_task_proto_depIdxs = []int32{
	6,  // 0: flyteidl.admin.TaskCreateRequest.id:type_name -> flyteidl.core.Identifier
	4,  // 1: flyteidl.admin.TaskCreateRequest.spec:type_name -> flyteidl.admin.TaskSpec
	6,  // 2: flyteidl.admin.Task.id:type_name -> flyteidl.core.Identifier
	5,  // 3: flyteidl.admin.Task.closure:type_name -> flyteidl.admin.TaskClosure
	2,  // 4: flyteidl.admin.TaskList.tasks:type_name -> flyteidl.admin.Task
	7,  // 5: flyteidl.admin.TaskSpec.template:type_name -> flyteidl.core.TaskTemplate
	8,  // 6: flyteidl.admin.TaskSpec.description:type_name -> flyteidl.admin.DescriptionEntity
	9,  // 7: flyteidl.admin.TaskClosure.compiled_task:type_name -> flyteidl.core.CompiledTask
	10, // 8: flyteidl.admin.TaskClosure.created_at:type_name -> google.protobuf.Timestamp
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_flyteidl_admin_task_proto_init() }
func file_flyteidl_admin_task_proto_init() {
	if File_flyteidl_admin_task_proto != nil {
		return
	}
	file_flyteidl_admin_description_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_admin_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskCreateRequest); i {
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
		file_flyteidl_admin_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskCreateResponse); i {
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
		file_flyteidl_admin_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_flyteidl_admin_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskList); i {
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
		file_flyteidl_admin_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskSpec); i {
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
		file_flyteidl_admin_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskClosure); i {
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
			RawDescriptor: file_flyteidl_admin_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_admin_task_proto_goTypes,
		DependencyIndexes: file_flyteidl_admin_task_proto_depIdxs,
		MessageInfos:      file_flyteidl_admin_task_proto_msgTypes,
	}.Build()
	File_flyteidl_admin_task_proto = out.File
	file_flyteidl_admin_task_proto_rawDesc = nil
	file_flyteidl_admin_task_proto_goTypes = nil
	file_flyteidl_admin_task_proto_depIdxs = nil
}
