# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/service/external_plugin_service.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import literals_pb2 as flyteidl_dot_core_dot_literals__pb2
from flyteidl.core import tasks_pb2 as flyteidl_dot_core_dot_tasks__pb2
from flyteidl.core import interface_pb2 as flyteidl_dot_core_dot_interface__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n.flyteidl/service/external_plugin_service.proto\x12\x10\x66lyteidl.service\x1a\x1c\x66lyteidl/core/literals.proto\x1a\x19\x66lyteidl/core/tasks.proto\x1a\x1d\x66lyteidl/core/interface.proto\"\xa4\x01\n\x11TaskCreateRequest\x12\x31\n\x06inputs\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.LiteralMapR\x06inputs\x12\x37\n\x08template\x18\x02 \x01(\x0b\x32\x1b.flyteidl.core.TaskTemplateR\x08template\x12#\n\routput_prefix\x18\x03 \x01(\tR\x0coutputPrefix\"+\n\x12TaskCreateResponse\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\"D\n\x0eTaskGetRequest\x12\x1b\n\ttask_type\x18\x01 \x01(\tR\x08taskType\x12\x15\n\x06job_id\x18\x02 \x01(\tR\x05jobId\"u\n\x0fTaskGetResponse\x12-\n\x05state\x18\x01 \x01(\x0e\x32\x17.flyteidl.service.StateR\x05state\x12\x33\n\x07outputs\x18\x02 \x01(\x0b\x32\x19.flyteidl.core.LiteralMapR\x07outputs\"G\n\x11TaskDeleteRequest\x12\x1b\n\ttask_type\x18\x01 \x01(\tR\x08taskType\x12\x15\n\x06job_id\x18\x02 \x01(\tR\x05jobId\"\x14\n\x12TaskDeleteResponse*^\n\x05State\x12\x15\n\x11RETRYABLE_FAILURE\x10\x00\x12\x15\n\x11PERMANENT_FAILURE\x10\x01\x12\x0b\n\x07PENDING\x10\x02\x12\x0b\n\x07RUNNING\x10\x03\x12\r\n\tSUCCEEDED\x10\x04\x32\x9f\x02\n\x15\x45xternalPluginService\x12Y\n\nCreateTask\x12#.flyteidl.service.TaskCreateRequest\x1a$.flyteidl.service.TaskCreateResponse\"\x00\x12P\n\x07GetTask\x12 .flyteidl.service.TaskGetRequest\x1a!.flyteidl.service.TaskGetResponse\"\x00\x12Y\n\nDeleteTask\x12#.flyteidl.service.TaskDeleteRequest\x1a$.flyteidl.service.TaskDeleteResponse\"\x00\x42\xcc\x01\n\x14\x63om.flyteidl.serviceB\x1a\x45xternalPluginServiceProtoP\x01Z7github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service\xa2\x02\x03\x46SX\xaa\x02\x10\x46lyteidl.Service\xca\x02\x10\x46lyteidl\\Service\xe2\x02\x1c\x46lyteidl\\Service\\GPBMetadata\xea\x02\x11\x46lyteidl::Serviceb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'flyteidl.service.external_plugin_service_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\024com.flyteidl.serviceB\032ExternalPluginServiceProtoP\001Z7github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service\242\002\003FSX\252\002\020Flyteidl.Service\312\002\020Flyteidl\\Service\342\002\034Flyteidl\\Service\\GPBMetadata\352\002\021Flyteidl::Service'
  _STATE._serialized_start=652
  _STATE._serialized_end=746
  _TASKCREATEREQUEST._serialized_start=157
  _TASKCREATEREQUEST._serialized_end=321
  _TASKCREATERESPONSE._serialized_start=323
  _TASKCREATERESPONSE._serialized_end=366
  _TASKGETREQUEST._serialized_start=368
  _TASKGETREQUEST._serialized_end=436
  _TASKGETRESPONSE._serialized_start=438
  _TASKGETRESPONSE._serialized_end=555
  _TASKDELETEREQUEST._serialized_start=557
  _TASKDELETEREQUEST._serialized_end=628
  _TASKDELETERESPONSE._serialized_start=630
  _TASKDELETERESPONSE._serialized_end=650
  _EXTERNALPLUGINSERVICE._serialized_start=749
  _EXTERNALPLUGINSERVICE._serialized_end=1036
# @@protoc_insertion_point(module_scope)
