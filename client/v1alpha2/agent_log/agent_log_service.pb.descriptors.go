// Code generated by protoc-gen-goten-client
// API: AgentLogService
// DO NOT EDIT!!!

package agent_log_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(proto.Message)
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &empty.Empty{}
	_ = &probe.Probe{}
)

var (
	descriptorsInitialized    bool
	agentLogServiceDescriptor *AgentLogServiceDescriptor
	reportAgentLogDescriptor  *ReportAgentLogDescriptor
	getAgentLogsDescriptor    *GetAgentLogsDescriptor
)

type ReportAgentLogDescriptor struct{}

type ReportAgentLogDescriptorClientMsgHandle struct{}

type ReportAgentLogDescriptorServerMsgHandle struct{}

func (d *ReportAgentLogDescriptor) NewEmptyClientMsg() proto.Message {
	return &ReportAgentLogRequest{}
}

func (d *ReportAgentLogDescriptor) NewEmptyServerMsg() proto.Message {
	return &empty.Empty{}
}

func (d *ReportAgentLogDescriptor) IsUnary() bool {
	return true
}

func (d *ReportAgentLogDescriptor) IsClientStream() bool {
	return false
}

func (d *ReportAgentLogDescriptor) IsServerStream() bool {
	return false
}

func (d *ReportAgentLogDescriptor) IsCollection() bool {
	return false
}

func (d *ReportAgentLogDescriptor) IsPlural() bool {
	return false
}

func (d *ReportAgentLogDescriptor) HasResource() bool {
	return true
}

func (d *ReportAgentLogDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ReportAgentLogDescriptor) GetVerb() string {
	return "reportAgentLog"
}

func (d *ReportAgentLogDescriptor) GetMethodName() string {
	return "ReportAgentLog"
}

func (d *ReportAgentLogDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.AgentLogService/ReportAgentLog"
}

func (d *ReportAgentLogDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *ReportAgentLogDescriptor) GetApiName() string {
	return "AgentLogService"
}

func (d *ReportAgentLogDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *ReportAgentLogDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ReportAgentLogDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return agentLogServiceDescriptor
}

func (d *ReportAgentLogDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return probe.GetDescriptor()
}

func (d *ReportAgentLogDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ReportAgentLogDescriptorClientMsgHandle{}
}

func (d *ReportAgentLogDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ReportAgentLogDescriptorServerMsgHandle{}
}

func (h *ReportAgentLogDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ReportAgentLogRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ReportAgentLogRequest) *probe.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	{
		if ref := typedMsg.GetName(); ref != nil {
			return &ref.Name
		}
	}
	return (*probe.Name)(nil)
}

func (h *ReportAgentLogDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ReportAgentLogRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ReportAgentLogRequest) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ReportAgentLogDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ReportAgentLogRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*ReportAgentLogRequest) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *ReportAgentLogDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*empty.Empty)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*empty.Empty) *probe.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *ReportAgentLogDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*empty.Empty)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*empty.Empty) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ReportAgentLogDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*empty.Empty)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*empty.Empty) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func GetReportAgentLogDescriptor() *ReportAgentLogDescriptor {
	return reportAgentLogDescriptor
}

type GetAgentLogsDescriptor struct{}

type GetAgentLogsDescriptorClientMsgHandle struct{}

type GetAgentLogsDescriptorServerMsgHandle struct{}

func (d *GetAgentLogsDescriptor) NewEmptyClientMsg() proto.Message {
	return &GetAgentLogsRequest{}
}

func (d *GetAgentLogsDescriptor) NewEmptyServerMsg() proto.Message {
	return &GetAgentLogsResponse{}
}

func (d *GetAgentLogsDescriptor) IsUnary() bool {
	return false
}

func (d *GetAgentLogsDescriptor) IsClientStream() bool {
	return false
}

func (d *GetAgentLogsDescriptor) IsServerStream() bool {
	return true
}

func (d *GetAgentLogsDescriptor) IsCollection() bool {
	return false
}

func (d *GetAgentLogsDescriptor) IsPlural() bool {
	return false
}

func (d *GetAgentLogsDescriptor) HasResource() bool {
	return false
}

func (d *GetAgentLogsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *GetAgentLogsDescriptor) GetVerb() string {
	return "getAgentLogs"
}

func (d *GetAgentLogsDescriptor) GetMethodName() string {
	return "GetAgentLogs"
}

func (d *GetAgentLogsDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.AgentLogService/GetAgentLogs"
}

func (d *GetAgentLogsDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *GetAgentLogsDescriptor) GetApiName() string {
	return "AgentLogService"
}

func (d *GetAgentLogsDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *GetAgentLogsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *GetAgentLogsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return agentLogServiceDescriptor
}

func (d *GetAgentLogsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return nil
}

func (d *GetAgentLogsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetAgentLogsDescriptorClientMsgHandle{}
}

func (d *GetAgentLogsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetAgentLogsDescriptorServerMsgHandle{}
}

func (h *GetAgentLogsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetAgentLogsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	return nil
}

func (h *GetAgentLogsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetAgentLogsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetAgentLogsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	return nil
}

func (h *GetAgentLogsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetGetAgentLogsDescriptor() *GetAgentLogsDescriptor {
	return getAgentLogsDescriptor
}

type AgentLogServiceDescriptor struct{}

func (d *AgentLogServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		reportAgentLogDescriptor,
		getAgentLogsDescriptor,
	}
}

func (d *AgentLogServiceDescriptor) GetFullAPIName() string {
	return "/ntt.watchdog.v1alpha2.AgentLogService"
}

func (d *AgentLogServiceDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *AgentLogServiceDescriptor) GetApiName() string {
	return "AgentLogService"
}

func (d *AgentLogServiceDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *AgentLogServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetAgentLogServiceDescriptor() *AgentLogServiceDescriptor {
	return agentLogServiceDescriptor
}

func initDescriptors() {
	agentLogServiceDescriptor = &AgentLogServiceDescriptor{}
	reportAgentLogDescriptor = &ReportAgentLogDescriptor{}
	getAgentLogsDescriptor = &GetAgentLogsDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(agentLogServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(reportAgentLogDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(getAgentLogsDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
