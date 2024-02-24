// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/core/workflow_closure.proto

package core

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

// Defines an enclosed package of workflow and tasks it references.
type WorkflowClosure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required. Workflow template.
	Workflow *WorkflowTemplate `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
	// optional. A collection of tasks referenced by the workflow. Only needed if the workflow
	// references tasks.
	Tasks []*TaskTemplate `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *WorkflowClosure) Reset() {
	*x = WorkflowClosure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_workflow_closure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowClosure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowClosure) ProtoMessage() {}

func (x *WorkflowClosure) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_workflow_closure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowClosure.ProtoReflect.Descriptor instead.
func (*WorkflowClosure) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_workflow_closure_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowClosure) GetWorkflow() *WorkflowTemplate {
	if x != nil {
		return x.Workflow
	}
	return nil
}

func (x *WorkflowClosure) GetTasks() []*TaskTemplate {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_flyteidl_core_workflow_closure_proto protoreflect.FileDescriptor

var file_flyteidl_core_workflow_closure_proto_rawDesc = []byte{
	0x0a, 0x24, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x1c, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81,
	0x01, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6c, 0x6f, 0x73, 0x75,
	0x72, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12,
	0x31, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x42, 0xba, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x14, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x43, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x46,
	0x43, 0x58, 0xaa, 0x02, 0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x43, 0x6f,
	0x72, 0x65, 0xca, 0x02, 0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x43, 0x6f,
	0x72, 0x65, 0xe2, 0x02, 0x19, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x43, 0x6f,
	0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x43, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_core_workflow_closure_proto_rawDescOnce sync.Once
	file_flyteidl_core_workflow_closure_proto_rawDescData = file_flyteidl_core_workflow_closure_proto_rawDesc
)

func file_flyteidl_core_workflow_closure_proto_rawDescGZIP() []byte {
	file_flyteidl_core_workflow_closure_proto_rawDescOnce.Do(func() {
		file_flyteidl_core_workflow_closure_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_core_workflow_closure_proto_rawDescData)
	})
	return file_flyteidl_core_workflow_closure_proto_rawDescData
}

var file_flyteidl_core_workflow_closure_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flyteidl_core_workflow_closure_proto_goTypes = []interface{}{
	(*WorkflowClosure)(nil),  // 0: flyteidl.core.WorkflowClosure
	(*WorkflowTemplate)(nil), // 1: flyteidl.core.WorkflowTemplate
	(*TaskTemplate)(nil),     // 2: flyteidl.core.TaskTemplate
}
var file_flyteidl_core_workflow_closure_proto_depIdxs = []int32{
	1, // 0: flyteidl.core.WorkflowClosure.workflow:type_name -> flyteidl.core.WorkflowTemplate
	2, // 1: flyteidl.core.WorkflowClosure.tasks:type_name -> flyteidl.core.TaskTemplate
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_flyteidl_core_workflow_closure_proto_init() }
func file_flyteidl_core_workflow_closure_proto_init() {
	if File_flyteidl_core_workflow_closure_proto != nil {
		return
	}
	file_flyteidl_core_workflow_proto_init()
	file_flyteidl_core_tasks_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_core_workflow_closure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowClosure); i {
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
			RawDescriptor: file_flyteidl_core_workflow_closure_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_core_workflow_closure_proto_goTypes,
		DependencyIndexes: file_flyteidl_core_workflow_closure_proto_depIdxs,
		MessageInfos:      file_flyteidl_core_workflow_closure_proto_msgTypes,
	}.Build()
	File_flyteidl_core_workflow_closure_proto = out.File
	file_flyteidl_core_workflow_closure_proto_rawDesc = nil
	file_flyteidl_core_workflow_closure_proto_goTypes = nil
	file_flyteidl_core_workflow_closure_proto_depIdxs = nil
}
