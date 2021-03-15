// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: flyteidl/service/admin.proto

#include "flyteidl/service/admin.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

namespace flyteidl {
namespace service {
}  // namespace service
}  // namespace flyteidl
void InitDefaults_flyteidl_2fservice_2fadmin_2eproto() {
}

constexpr ::google::protobuf::Metadata* file_level_metadata_flyteidl_2fservice_2fadmin_2eproto = nullptr;
constexpr ::google::protobuf::EnumDescriptor const** file_level_enum_descriptors_flyteidl_2fservice_2fadmin_2eproto = nullptr;
constexpr ::google::protobuf::ServiceDescriptor const** file_level_service_descriptors_flyteidl_2fservice_2fadmin_2eproto = nullptr;
const ::google::protobuf::uint32 TableStruct_flyteidl_2fservice_2fadmin_2eproto::offsets[1] = {};
static constexpr ::google::protobuf::internal::MigrationSchema* schemas = nullptr;
static constexpr ::google::protobuf::Message* const* file_default_instances = nullptr;

::google::protobuf::internal::AssignDescriptorsTable assign_descriptors_table_flyteidl_2fservice_2fadmin_2eproto = {
  {}, AddDescriptors_flyteidl_2fservice_2fadmin_2eproto, "flyteidl/service/admin.proto", schemas,
  file_default_instances, TableStruct_flyteidl_2fservice_2fadmin_2eproto::offsets,
  file_level_metadata_flyteidl_2fservice_2fadmin_2eproto, 0, file_level_enum_descriptors_flyteidl_2fservice_2fadmin_2eproto, file_level_service_descriptors_flyteidl_2fservice_2fadmin_2eproto,
};

const char descriptor_table_protodef_flyteidl_2fservice_2fadmin_2eproto[] =
  "\n\034flyteidl/service/admin.proto\022\020flyteidl"
  ".service\032\034google/api/annotations.proto\032\034"
  "flyteidl/admin/project.proto\032.flyteidl/a"
  "dmin/project_domain_attributes.proto\032\031fl"
  "yteidl/admin/task.proto\032\035flyteidl/admin/"
  "workflow.proto\032(flyteidl/admin/workflow_"
  "attributes.proto\032 flyteidl/admin/launch_"
  "plan.proto\032\032flyteidl/admin/event.proto\032\036"
  "flyteidl/admin/execution.proto\032\'flyteidl"
  "/admin/matchable_resource.proto\032#flyteid"
  "l/admin/node_execution.proto\032#flyteidl/a"
  "dmin/task_execution.proto\032\034flyteidl/admi"
  "n/version.proto\032\033flyteidl/admin/common.p"
  "roto\032,protoc-gen-swagger/options/annotat"
  "ions.proto2\245a\n\014AdminService\022\305\002\n\nCreateTa"
  "sk\022!.flyteidl.admin.TaskCreateRequest\032\"."
  "flyteidl.admin.TaskCreateResponse\"\357\001\202\323\344\223"
  "\002\022\"\r/api/v1/tasks:\001*\222A\323\001\032&Create and reg"
  "ister a task definition.JB\n\003400\022;\n9Retur"
  "ned for bad request that may have failed"
  " validation.Je\n\003409\022^\n\\Returned for a re"
  "quest that references an identical entit"
  "y that has already been registered.\022\262\001\n\007"
  "GetTask\022 .flyteidl.admin.ObjectGetReques"
  "t\032\024.flyteidl.admin.Task\"o\202\323\344\223\002\?\022=/api/v1"
  "/tasks/{id.project}/{id.domain}/{id.name"
  "}/{id.version}\222A\'\032%Retrieve an existing "
  "task definition.\022\336\001\n\013ListTaskIds\0220.flyte"
  "idl.admin.NamedEntityIdentifierListReque"
  "st\032).flyteidl.admin.NamedEntityIdentifie"
  "rList\"r\202\323\344\223\002%\022#/api/v1/task_ids/{project"
  "}/{domain}\222AD\032BFetch existing task defin"
  "ition identifiers matching input filters"
  ".\022\353\001\n\tListTasks\022#.flyteidl.admin.Resourc"
  "eListRequest\032\030.flyteidl.admin.TaskList\"\236"
  "\001\202\323\344\223\002\\\0220/api/v1/tasks/{id.project}/{id."
  "domain}/{id.name}Z(\022&/api/v1/tasks/{id.p"
  "roject}/{id.domain}\222A9\0327Fetch existing t"
  "ask definitions matching input filters.\022"
  "\331\002\n\016CreateWorkflow\022%.flyteidl.admin.Work"
  "flowCreateRequest\032&.flyteidl.admin.Workf"
  "lowCreateResponse\"\367\001\202\323\344\223\002\026\"\021/api/v1/work"
  "flows:\001*\222A\327\001\032*Create and register a work"
  "flow definition.JB\n\003400\022;\n9Returned for "
  "bad request that may have failed validat"
  "ion.Je\n\003409\022^\n\\Returned for a request th"
  "at references an identical entity that h"
  "as already been registered.\022\302\001\n\013GetWorkf"
  "low\022 .flyteidl.admin.ObjectGetRequest\032\030."
  "flyteidl.admin.Workflow\"w\202\323\344\223\002C\022A/api/v1"
  "/workflows/{id.project}/{id.domain}/{id."
  "name}/{id.version}\222A+\032)Retrieve an exist"
  "ing workflow definition.\022\355\001\n\017ListWorkflo"
  "wIds\0220.flyteidl.admin.NamedEntityIdentif"
  "ierListRequest\032).flyteidl.admin.NamedEnt"
  "ityIdentifierList\"}\202\323\344\223\002)\022\'/api/v1/workf"
  "low_ids/{project}/{domain}\222AK\032IFetch an "
  "existing workflow definition identifiers"
  " matching input filters.\022\377\001\n\rListWorkflo"
  "ws\022#.flyteidl.admin.ResourceListRequest\032"
  "\034.flyteidl.admin.WorkflowList\"\252\001\202\323\344\223\002d\0224"
  "/api/v1/workflows/{id.project}/{id.domai"
  "n}/{id.name}Z,\022*/api/v1/workflows/{id.pr"
  "oject}/{id.domain}\222A=\032;Fetch existing wo"
  "rkflow definitions matching input filter"
  "s.\022\345\002\n\020CreateLaunchPlan\022\'.flyteidl.admin"
  ".LaunchPlanCreateRequest\032(.flyteidl.admi"
  "n.LaunchPlanCreateResponse\"\375\001\202\323\344\223\002\031\"\024/ap"
  "i/v1/launch_plans:\001*\222A\332\001\032-Create and reg"
  "ister a launch plan definition.JB\n\003400\022;"
  "\n9Returned for bad request that may have"
  " failed validation.Je\n\003409\022^\n\\Returned f"
  "or a request that references an identica"
  "l entity that has already been registere"
  "d.\022\314\001\n\rGetLaunchPlan\022 .flyteidl.admin.Ob"
  "jectGetRequest\032\032.flyteidl.admin.LaunchPl"
  "an\"}\202\323\344\223\002F\022D/api/v1/launch_plans/{id.pro"
  "ject}/{id.domain}/{id.name}/{id.version}"
  "\222A.\032,Retrieve an existing launch plan de"
  "finition.\022\363\001\n\023GetActiveLaunchPlan\022\'.flyt"
  "eidl.admin.ActiveLaunchPlanRequest\032\032.fly"
  "teidl.admin.LaunchPlan\"\226\001\202\323\344\223\002@\022>/api/v1"
  "/active_launch_plans/{id.project}/{id.do"
  "main}/{id.name}\222AM\032KRetrieve the active "
  "launch plan version specified by input r"
  "equest filters.\022\353\001\n\025ListActiveLaunchPlan"
  "s\022+.flyteidl.admin.ActiveLaunchPlanListR"
  "equest\032\036.flyteidl.admin.LaunchPlanList\"\204"
  "\001\202\323\344\223\0020\022./api/v1/active_launch_plans/{pr"
  "oject}/{domain}\222AK\032IFetch the active lau"
  "nch plan versions specified by input req"
  "uest filters.\022\363\001\n\021ListLaunchPlanIds\0220.fl"
  "yteidl.admin.NamedEntityIdentifierListRe"
  "quest\032).flyteidl.admin.NamedEntityIdenti"
  "fierList\"\200\001\202\323\344\223\002,\022*/api/v1/launch_plan_i"
  "ds/{project}/{domain}\222AK\032IFetch existing"
  " launch plan definition identifiers matc"
  "hing input filters.\022\214\002\n\017ListLaunchPlans\022"
  "#.flyteidl.admin.ResourceListRequest\032\036.f"
  "lyteidl.admin.LaunchPlanList\"\263\001\202\323\344\223\002j\0227/"
  "api/v1/launch_plans/{id.project}/{id.dom"
  "ain}/{id.name}Z/\022-/api/v1/launch_plans/{"
  "id.project}/{id.domain}\222A@\032>Fetch existi"
  "ng launch plan definitions matching inpu"
  "t filters.\022\300\006\n\020UpdateLaunchPlan\022\'.flytei"
  "dl.admin.LaunchPlanUpdateRequest\032(.flyte"
  "idl.admin.LaunchPlanUpdateResponse\"\330\005\202\323\344"
  "\223\002I\032D/api/v1/launch_plans/{id.project}/{"
  "id.domain}/{id.name}/{id.version}:\001*\222A\205\005"
  "\032\202\005Update the status of an existing laun"
  "ch plan definition. At most one launch p"
  "lan version for a given {project, domain"
  ", name} can be active at a time. If this"
  " call sets a launch plan to active and e"
  "xisting version is already active, the r"
  "esult of this call will be that the form"
  "erly active launch plan will be made ina"
  "ctive and specified launch plan in this "
  "request will be made active. In the even"
  "t that the formerly active launch plan h"
  "ad a schedule associated it with it, thi"
  "s schedule will be disabled. If the refe"
  "rence launch plan in this request is bei"
  "ng set to active and has a schedule asso"
  "ciated with it, the schedule will be ena"
  "bled.\022\242\001\n\017CreateExecution\022&.flyteidl.adm"
  "in.ExecutionCreateRequest\032\'.flyteidl.adm"
  "in.ExecutionCreateResponse\">\202\323\344\223\002\027\"\022/api"
  "/v1/executions:\001*\222A\036\032\034Create a workflow "
  "execution.\022\261\001\n\021RelaunchExecution\022(.flyte"
  "idl.admin.ExecutionRelaunchRequest\032\'.fly"
  "teidl.admin.ExecutionCreateResponse\"I\202\323\344"
  "\223\002 \"\033/api/v1/executions/relaunch:\001*\222A \032\036"
  "Relaunch a workflow execution.\022\302\001\n\014GetEx"
  "ecution\022+.flyteidl.admin.WorkflowExecuti"
  "onGetRequest\032\031.flyteidl.admin.Execution\""
  "j\202\323\344\223\0027\0225/api/v1/executions/{id.project}"
  "/{id.domain}/{id.name}\222A*\032(Retrieve an e"
  "xisting workflow execution.\022\202\002\n\020GetExecu"
  "tionData\022/.flyteidl.admin.WorkflowExecut"
  "ionGetDataRequest\0320.flyteidl.admin.Workf"
  "lowExecutionGetDataResponse\"\212\001\202\323\344\223\002<\022:/a"
  "pi/v1/data/executions/{id.project}/{id.d"
  "omain}/{id.name}\222AE\032CRetrieve input and "
  "output data from an existing workflow ex"
  "ecution.\022\310\001\n\016ListExecutions\022#.flyteidl.a"
  "dmin.ResourceListRequest\032\035.flyteidl.admi"
  "n.ExecutionList\"r\202\323\344\223\002-\022+/api/v1/executi"
  "ons/{id.project}/{id.domain}\222A<\032:Fetch e"
  "xisting workflow executions matching inp"
  "ut filters.\022\364\001\n\022TerminateExecution\022).fly"
  "teidl.admin.ExecutionTerminateRequest\032*."
  "flyteidl.admin.ExecutionTerminateRespons"
  "e\"\206\001\202\323\344\223\002:*5/api/v1/executions/{id.proje"
  "ct}/{id.domain}/{id.name}:\001*\222AC\032ATermina"
  "te the active workflow execution specifi"
  "ed in the request.\022\374\001\n\020GetNodeExecution\022"
  "\'.flyteidl.admin.NodeExecutionGetRequest"
  "\032\035.flyteidl.admin.NodeExecution\"\237\001\202\323\344\223\002p"
  "\022n/api/v1/node_executions/{id.execution_"
  "id.project}/{id.execution_id.domain}/{id"
  ".execution_id.name}/{id.node_id}\222A&\032$Ret"
  "rieve an existing node execution.\022\232\002\n\022Li"
  "stNodeExecutions\022(.flyteidl.admin.NodeEx"
  "ecutionListRequest\032!.flyteidl.admin.Node"
  "ExecutionList\"\266\001\202\323\344\223\002u\022s/api/v1/node_exe"
  "cutions/{workflow_execution_id.project}/"
  "{workflow_execution_id.domain}/{workflow"
  "_execution_id.name}\222A8\0326Fetch existing n"
  "ode executions matching input filters.\022\357"
  "\004\n\031ListNodeExecutionsForTask\022/.flyteidl."
  "admin.NodeExecutionForTaskListRequest\032!."
  "flyteidl.admin.NodeExecutionList\"\375\003\202\323\344\223\002"
  "\254\003\022\251\003/api/v1/children/task_executions/{t"
  "ask_execution_id.node_execution_id.execu"
  "tion_id.project}/{task_execution_id.node"
  "_execution_id.execution_id.domain}/{task"
  "_execution_id.node_execution_id.executio"
  "n_id.name}/{task_execution_id.node_execu"
  "tion_id.node_id}/{task_execution_id.task"
  "_id.project}/{task_execution_id.task_id."
  "domain}/{task_execution_id.task_id.name}"
  "/{task_execution_id.task_id.version}/{ta"
  "sk_execution_id.retry_attempt}\222AG\032EFetch"
  " child node executions launched by the s"
  "pecified task execution.\022\263\002\n\024GetNodeExec"
  "utionData\022+.flyteidl.admin.NodeExecution"
  "GetDataRequest\032,.flyteidl.admin.NodeExec"
  "utionGetDataResponse\"\277\001\202\323\344\223\002u\022s/api/v1/d"
  "ata/node_executions/{id.execution_id.pro"
  "ject}/{id.execution_id.domain}/{id.execu"
  "tion_id.name}/{id.node_id}\222AA\032\?Retrieve "
  "input and output data from an existing n"
  "ode execution.\022\227\001\n\017RegisterProject\022&.fly"
  "teidl.admin.ProjectRegisterRequest\032\'.fly"
  "teidl.admin.ProjectRegisterResponse\"3\202\323\344"
  "\223\002\025\"\020/api/v1/projects:\001*\222A\025\032\023Register a "
  "project.\022\207\001\n\rUpdateProject\022\027.flyteidl.ad"
  "min.Project\032%.flyteidl.admin.ProjectUpda"
  "teResponse\"6\202\323\344\223\002\032\032\025/api/v1/projects/{id"
  "}:\001*\222A\023\032\021Update a project.\022\205\001\n\014ListProje"
  "cts\022\".flyteidl.admin.ProjectListRequest\032"
  "\030.flyteidl.admin.Projects\"7\202\323\344\223\002\022\022\020/api/"
  "v1/projects\222A\034\032\032Fetch registered project"
  "s.\022\335\001\n\023CreateWorkflowEvent\022-.flyteidl.ad"
  "min.WorkflowExecutionEventRequest\032..flyt"
  "eidl.admin.WorkflowExecutionEventRespons"
  "e\"g\202\323\344\223\002\035\"\030/api/v1/events/workflows:\001*\222A"
  "A\032\?Create a workflow execution event rec"
  "ording a phase transition.\022\311\001\n\017CreateNod"
  "eEvent\022).flyteidl.admin.NodeExecutionEve"
  "ntRequest\032*.flyteidl.admin.NodeExecution"
  "EventResponse\"_\202\323\344\223\002\031\"\024/api/v1/events/no"
  "des:\001*\222A=\032;Create a node execution event"
  " recording a phase transition.\022\311\001\n\017Creat"
  "eTaskEvent\022).flyteidl.admin.TaskExecutio"
  "nEventRequest\032*.flyteidl.admin.TaskExecu"
  "tionEventResponse\"_\202\323\344\223\002\031\"\024/api/v1/event"
  "s/tasks:\001*\222A=\032;Create a task execution e"
  "vent recording a phase transition.\022\251\003\n\020G"
  "etTaskExecution\022\'.flyteidl.admin.TaskExe"
  "cutionGetRequest\032\035.flyteidl.admin.TaskEx"
  "ecution\"\314\002\202\323\344\223\002\234\002\022\231\002/api/v1/task_executi"
  "ons/{id.node_execution_id.execution_id.p"
  "roject}/{id.node_execution_id.execution_"
  "id.domain}/{id.node_execution_id.executi"
  "on_id.name}/{id.node_execution_id.node_i"
  "d}/{id.task_id.project}/{id.task_id.doma"
  "in}/{id.task_id.name}/{id.task_id.versio"
  "n}/{id.retry_attempt}\222A&\032$Retrieve an ex"
  "isting task execution.\022\323\002\n\022ListTaskExecu"
  "tions\022(.flyteidl.admin.TaskExecutionList"
  "Request\032!.flyteidl.admin.TaskExecutionLi"
  "st\"\357\001\202\323\344\223\002\255\001\022\252\001/api/v1/task_executions/{"
  "node_execution_id.execution_id.project}/"
  "{node_execution_id.execution_id.domain}/"
  "{node_execution_id.execution_id.name}/{n"
  "ode_execution_id.node_id}\222A8\0326Fetch exis"
  "ting task executions matching input filt"
  "ers.\022\340\003\n\024GetTaskExecutionData\022+.flyteidl"
  ".admin.TaskExecutionGetDataRequest\032,.fly"
  "teidl.admin.TaskExecutionGetDataResponse"
  "\"\354\002\202\323\344\223\002\241\002\022\236\002/api/v1/data/task_execution"
  "s/{id.node_execution_id.execution_id.pro"
  "ject}/{id.node_execution_id.execution_id"
  ".domain}/{id.node_execution_id.execution"
  "_id.name}/{id.node_execution_id.node_id}"
  "/{id.task_id.project}/{id.task_id.domain"
  "}/{id.task_id.name}/{id.task_id.version}"
  "/{id.retry_attempt}\222AA\032\?Retrieve input a"
  "nd output data from an existing task exe"
  "cution.\022\277\002\n\035UpdateProjectDomainAttribute"
  "s\0224.flyteidl.admin.ProjectDomainAttribut"
  "esUpdateRequest\0325.flyteidl.admin.Project"
  "DomainAttributesUpdateResponse\"\260\001\202\323\344\223\002O\032"
  "J/api/v1/project_domain_attributes/{attr"
  "ibutes.project}/{attributes.domain}:\001*\222A"
  "X\032VUpdate the customized resource attrib"
  "utes associated with a project-domain co"
  "mbination\022\237\002\n\032GetProjectDomainAttributes"
  "\0221.flyteidl.admin.ProjectDomainAttribute"
  "sGetRequest\0322.flyteidl.admin.ProjectDoma"
  "inAttributesGetResponse\"\231\001\202\323\344\223\0026\0224/api/v"
  "1/project_domain_attributes/{project}/{d"
  "omain}\222AZ\032XRetrieve the customized resou"
  "rce attributes associated with a project"
  "-domain combination\022\251\002\n\035DeleteProjectDom"
  "ainAttributes\0224.flyteidl.admin.ProjectDo"
  "mainAttributesDeleteRequest\0325.flyteidl.a"
  "dmin.ProjectDomainAttributesDeleteRespon"
  "se\"\232\001\202\323\344\223\0029*4/api/v1/project_domain_attr"
  "ibutes/{project}/{domain}:\001*\222AX\032VDelete "
  "the customized resource attributes assoc"
  "iated with a project-domain combination\022"
  "\316\002\n\030UpdateWorkflowAttributes\022/.flyteidl."
  "admin.WorkflowAttributesUpdateRequest\0320."
  "flyteidl.admin.WorkflowAttributesUpdateR"
  "esponse\"\316\001\202\323\344\223\002_\032Z/api/v1/workflow_attri"
  "butes/{attributes.project}/{attributes.d"
  "omain}/{attributes.workflow}:\001*\222Af\032dUpda"
  "te the customized resource attributes as"
  "sociated with a project, domain and work"
  "flow combination\022\243\002\n\025GetWorkflowAttribut"
  "es\022,.flyteidl.admin.WorkflowAttributesGe"
  "tRequest\032-.flyteidl.admin.WorkflowAttrib"
  "utesGetResponse\"\254\001\202\323\344\223\002;\0229/api/v1/workfl"
  "ow_attributes/{project}/{domain}/{workfl"
  "ow}\222Ah\032fRetrieve the customized resource"
  " attributes associated with a project, d"
  "omain and workflow combination\022\255\002\n\030Delet"
  "eWorkflowAttributes\022/.flyteidl.admin.Wor"
  "kflowAttributesDeleteRequest\0320.flyteidl."
  "admin.WorkflowAttributesDeleteResponse\"\255"
  "\001\202\323\344\223\002>*9/api/v1/workflow_attributes/{pr"
  "oject}/{domain}/{workflow}:\001*\222Af\032dDelete"
  " the customized resource attributes asso"
  "ciated with a project, domain and workfl"
  "ow combination\022\341\001\n\027ListMatchableAttribut"
  "es\022..flyteidl.admin.ListMatchableAttribu"
  "tesRequest\032/.flyteidl.admin.ListMatchabl"
  "eAttributesResponse\"e\202\323\344\223\002\036\022\034/api/v1/mat"
  "chable_attributes\222A>\032<Retrieve a list of"
  " MatchableAttributesConfiguration object"
  "s.\022\200\002\n\021ListNamedEntities\022&.flyteidl.admi"
  "n.NamedEntityListRequest\032\037.flyteidl.admi"
  "n.NamedEntityList\"\241\001\202\323\344\223\002;\0229/api/v1/name"
  "d_entities/{resource_type}/{project}/{do"
  "main}\222A]\032[Retrieve a list of NamedEntity"
  " objects sharing a common resource type,"
  " project, and domain.\022\312\001\n\016GetNamedEntity"
  "\022%.flyteidl.admin.NamedEntityGetRequest\032"
  "\033.flyteidl.admin.NamedEntity\"t\202\323\344\223\002K\022I/a"
  "pi/v1/named_entities/{resource_type}/{id"
  ".project}/{id.domain}/{id.name}\222A \032\036Retr"
  "ieve a NamedEntity object.\022\363\001\n\021UpdateNam"
  "edEntity\022(.flyteidl.admin.NamedEntityUpd"
  "ateRequest\032).flyteidl.admin.NamedEntityU"
  "pdateResponse\"\210\001\202\323\344\223\002N\032I/api/v1/named_en"
  "tities/{resource_type}/{id.project}/{id."
  "domain}/{id.name}:\001*\222A1\032/Update the fiel"
  "ds associated with a NamedEntity\022\277\001\n\nGet"
  "Version\022!.flyteidl.admin.GetVersionReque"
  "st\032\".flyteidl.admin.GetVersionResponse\"j"
  "\202\323\344\223\002\021\022\017/api/v1/version\222AP\032NRetrieve the"
  " Version (including the Build  informati"
  "on) for FlyteAdmin serviceB9Z7github.com"
  "/flyteorg/flyteidl/gen/pb-go/flyteidl/se"
  "rviceb\006proto3"
  ;
::google::protobuf::internal::DescriptorTable descriptor_table_flyteidl_2fservice_2fadmin_2eproto = {
  false, InitDefaults_flyteidl_2fservice_2fadmin_2eproto, 
  descriptor_table_protodef_flyteidl_2fservice_2fadmin_2eproto,
  "flyteidl/service/admin.proto", &assign_descriptors_table_flyteidl_2fservice_2fadmin_2eproto, 13093,
};

void AddDescriptors_flyteidl_2fservice_2fadmin_2eproto() {
  static constexpr ::google::protobuf::internal::InitFunc deps[15] =
  {
    ::AddDescriptors_google_2fapi_2fannotations_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fproject_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fproject_5fdomain_5fattributes_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2ftask_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fworkflow_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fworkflow_5fattributes_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2flaunch_5fplan_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fevent_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fexecution_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fmatchable_5fresource_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fnode_5fexecution_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2ftask_5fexecution_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fversion_2eproto,
    ::AddDescriptors_flyteidl_2fadmin_2fcommon_2eproto,
    ::AddDescriptors_protoc_2dgen_2dswagger_2foptions_2fannotations_2eproto,
  };
 ::google::protobuf::internal::AddDescriptors(&descriptor_table_flyteidl_2fservice_2fadmin_2eproto, deps, 15);
}

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_flyteidl_2fservice_2fadmin_2eproto = []() { AddDescriptors_flyteidl_2fservice_2fadmin_2eproto(); return true; }();
namespace flyteidl {
namespace service {

// @@protoc_insertion_point(namespace_scope)
}  // namespace service
}  // namespace flyteidl
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
