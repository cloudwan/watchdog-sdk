// Code generated by protoc-gen-goten-client
// API: PcapService
// DO NOT EDIT!!!

package pcap_client

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
	descriptorsInitialized         bool
	pcapServiceDescriptor          *PcapServiceDescriptor
	reportPcapDescriptor           *ReportPcapDescriptor
	getPcapDescriptor              *GetPcapDescriptor
	getPcapFileFromAgentDescriptor *GetPcapFileFromAgentDescriptor
)

type ReportPcapDescriptor struct{}

type ReportPcapDescriptorClientMsgHandle struct{}

type ReportPcapDescriptorServerMsgHandle struct{}

func (d *ReportPcapDescriptor) NewEmptyClientMsg() proto.Message {
	return &ReportPcapRequest{}
}

func (d *ReportPcapDescriptor) NewEmptyServerMsg() proto.Message {
	return &empty.Empty{}
}

func (d *ReportPcapDescriptor) IsUnary() bool {
	return true
}

func (d *ReportPcapDescriptor) IsClientStream() bool {
	return false
}

func (d *ReportPcapDescriptor) IsServerStream() bool {
	return false
}

func (d *ReportPcapDescriptor) IsCollection() bool {
	return false
}

func (d *ReportPcapDescriptor) IsPlural() bool {
	return false
}

func (d *ReportPcapDescriptor) HasResource() bool {
	return true
}

func (d *ReportPcapDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ReportPcapDescriptor) GetVerb() string {
	return "reportPcap"
}

func (d *ReportPcapDescriptor) GetMethodName() string {
	return "ReportPcap"
}

func (d *ReportPcapDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.PcapService/ReportPcap"
}

func (d *ReportPcapDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *ReportPcapDescriptor) GetApiName() string {
	return "PcapService"
}

func (d *ReportPcapDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *ReportPcapDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ReportPcapDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return pcapServiceDescriptor
}

func (d *ReportPcapDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return probe.GetDescriptor()
}

func (d *ReportPcapDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ReportPcapDescriptorClientMsgHandle{}
}

func (d *ReportPcapDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ReportPcapDescriptorServerMsgHandle{}
}

func (h *ReportPcapDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ReportPcapRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ReportPcapRequest) *probe.Name
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

func (h *ReportPcapDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ReportPcapRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ReportPcapRequest) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ReportPcapDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ReportPcapRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*ReportPcapRequest) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *ReportPcapDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
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

func (h *ReportPcapDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
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

func (h *ReportPcapDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
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

func GetReportPcapDescriptor() *ReportPcapDescriptor {
	return reportPcapDescriptor
}

type GetPcapDescriptor struct{}

type GetPcapDescriptorClientMsgHandle struct{}

type GetPcapDescriptorServerMsgHandle struct{}

func (d *GetPcapDescriptor) NewEmptyClientMsg() proto.Message {
	return &GetPcapRequest{}
}

func (d *GetPcapDescriptor) NewEmptyServerMsg() proto.Message {
	return &GetPcapResponse{}
}

func (d *GetPcapDescriptor) IsUnary() bool {
	return true
}

func (d *GetPcapDescriptor) IsClientStream() bool {
	return false
}

func (d *GetPcapDescriptor) IsServerStream() bool {
	return false
}

func (d *GetPcapDescriptor) IsCollection() bool {
	return false
}

func (d *GetPcapDescriptor) IsPlural() bool {
	return false
}

func (d *GetPcapDescriptor) HasResource() bool {
	return false
}

func (d *GetPcapDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *GetPcapDescriptor) GetVerb() string {
	return "getPcap"
}

func (d *GetPcapDescriptor) GetMethodName() string {
	return "GetPcap"
}

func (d *GetPcapDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.PcapService/GetPcap"
}

func (d *GetPcapDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *GetPcapDescriptor) GetApiName() string {
	return "PcapService"
}

func (d *GetPcapDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *GetPcapDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *GetPcapDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return pcapServiceDescriptor
}

func (d *GetPcapDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return nil
}

func (d *GetPcapDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetPcapDescriptorClientMsgHandle{}
}

func (d *GetPcapDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetPcapDescriptorServerMsgHandle{}
}

func (h *GetPcapDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetPcapDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	return nil
}

func (h *GetPcapDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetPcapDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetPcapDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	return nil
}

func (h *GetPcapDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetGetPcapDescriptor() *GetPcapDescriptor {
	return getPcapDescriptor
}

type GetPcapFileFromAgentDescriptor struct{}

type GetPcapFileFromAgentDescriptorClientMsgHandle struct{}

type GetPcapFileFromAgentDescriptorServerMsgHandle struct{}

func (d *GetPcapFileFromAgentDescriptor) NewEmptyClientMsg() proto.Message {
	return &GetPcapFileFromAgentRequest{}
}

func (d *GetPcapFileFromAgentDescriptor) NewEmptyServerMsg() proto.Message {
	return &GetPcapResponse{}
}

func (d *GetPcapFileFromAgentDescriptor) IsUnary() bool {
	return false
}

func (d *GetPcapFileFromAgentDescriptor) IsClientStream() bool {
	return false
}

func (d *GetPcapFileFromAgentDescriptor) IsServerStream() bool {
	return true
}

func (d *GetPcapFileFromAgentDescriptor) IsCollection() bool {
	return false
}

func (d *GetPcapFileFromAgentDescriptor) IsPlural() bool {
	return false
}

func (d *GetPcapFileFromAgentDescriptor) HasResource() bool {
	return true
}

func (d *GetPcapFileFromAgentDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *GetPcapFileFromAgentDescriptor) GetVerb() string {
	return "getPcapFileFromAgent"
}

func (d *GetPcapFileFromAgentDescriptor) GetMethodName() string {
	return "GetPcapFileFromAgent"
}

func (d *GetPcapFileFromAgentDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.PcapService/GetPcapFileFromAgent"
}

func (d *GetPcapFileFromAgentDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *GetPcapFileFromAgentDescriptor) GetApiName() string {
	return "PcapService"
}

func (d *GetPcapFileFromAgentDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *GetPcapFileFromAgentDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *GetPcapFileFromAgentDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return pcapServiceDescriptor
}

func (d *GetPcapFileFromAgentDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return probe.GetDescriptor()
}

func (d *GetPcapFileFromAgentDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetPcapFileFromAgentDescriptorClientMsgHandle{}
}

func (d *GetPcapFileFromAgentDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetPcapFileFromAgentDescriptorServerMsgHandle{}
}

func (h *GetPcapFileFromAgentDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*GetPcapFileFromAgentRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*GetPcapFileFromAgentRequest) *probe.Name
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

func (h *GetPcapFileFromAgentDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*GetPcapFileFromAgentRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*GetPcapFileFromAgentRequest) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *GetPcapFileFromAgentDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*GetPcapFileFromAgentRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*GetPcapFileFromAgentRequest) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *GetPcapFileFromAgentDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*GetPcapResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*GetPcapResponse) *probe.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *GetPcapFileFromAgentDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*GetPcapResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*GetPcapResponse) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *GetPcapFileFromAgentDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*GetPcapResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*GetPcapResponse) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func GetGetPcapFileFromAgentDescriptor() *GetPcapFileFromAgentDescriptor {
	return getPcapFileFromAgentDescriptor
}

type PcapServiceDescriptor struct{}

func (d *PcapServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		reportPcapDescriptor,
		getPcapDescriptor,
		getPcapFileFromAgentDescriptor,
	}
}

func (d *PcapServiceDescriptor) GetFullAPIName() string {
	return "/ntt.watchdog.v1alpha2.PcapService"
}

func (d *PcapServiceDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *PcapServiceDescriptor) GetApiName() string {
	return "PcapService"
}

func (d *PcapServiceDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *PcapServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetPcapServiceDescriptor() *PcapServiceDescriptor {
	return pcapServiceDescriptor
}

func initDescriptors() {
	pcapServiceDescriptor = &PcapServiceDescriptor{}
	reportPcapDescriptor = &ReportPcapDescriptor{}
	getPcapDescriptor = &GetPcapDescriptor{}
	getPcapFileFromAgentDescriptor = &GetPcapFileFromAgentDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(pcapServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(reportPcapDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(getPcapDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(getPcapFileFromAgentDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
