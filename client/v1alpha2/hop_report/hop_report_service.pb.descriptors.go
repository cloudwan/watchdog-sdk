// Code generated by protoc-gen-goten-client
// API: HopReportService
// DO NOT EDIT!!!

package hop_report_client

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
	descriptorsInitialized     bool
	hopReportServiceDescriptor *HopReportServiceDescriptor
	reportHopsDescriptor       *ReportHopsDescriptor
	getHopReportsDescriptor    *GetHopReportsDescriptor
)

type ReportHopsDescriptor struct{}

type ReportHopsDescriptorClientMsgHandle struct{}

type ReportHopsDescriptorServerMsgHandle struct{}

func (d *ReportHopsDescriptor) NewEmptyClientMsg() proto.Message {
	return &ReportHopsRequest{}
}

func (d *ReportHopsDescriptor) NewEmptyServerMsg() proto.Message {
	return &empty.Empty{}
}

func (d *ReportHopsDescriptor) IsUnary() bool {
	return true
}

func (d *ReportHopsDescriptor) IsClientStream() bool {
	return false
}

func (d *ReportHopsDescriptor) IsServerStream() bool {
	return false
}

func (d *ReportHopsDescriptor) IsCollection() bool {
	return false
}

func (d *ReportHopsDescriptor) IsPlural() bool {
	return false
}

func (d *ReportHopsDescriptor) HasResource() bool {
	return true
}

func (d *ReportHopsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ReportHopsDescriptor) GetVerb() string {
	return "reportHops"
}

func (d *ReportHopsDescriptor) GetMethodName() string {
	return "ReportHops"
}

func (d *ReportHopsDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.HopReportService/ReportHops"
}

func (d *ReportHopsDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *ReportHopsDescriptor) GetApiName() string {
	return "HopReportService"
}

func (d *ReportHopsDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *ReportHopsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ReportHopsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return hopReportServiceDescriptor
}

func (d *ReportHopsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return probe.GetDescriptor()
}

func (d *ReportHopsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ReportHopsDescriptorClientMsgHandle{}
}

func (d *ReportHopsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ReportHopsDescriptorServerMsgHandle{}
}

func (h *ReportHopsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ReportHopsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ReportHopsRequest) *probe.Name
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

func (h *ReportHopsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ReportHopsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ReportHopsRequest) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ReportHopsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ReportHopsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*ReportHopsRequest) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *ReportHopsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
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

func (h *ReportHopsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
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

func (h *ReportHopsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
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

func GetReportHopsDescriptor() *ReportHopsDescriptor {
	return reportHopsDescriptor
}

type GetHopReportsDescriptor struct{}

type GetHopReportsDescriptorClientMsgHandle struct{}

type GetHopReportsDescriptorServerMsgHandle struct{}

func (d *GetHopReportsDescriptor) NewEmptyClientMsg() proto.Message {
	return &GetHopReportsRequest{}
}

func (d *GetHopReportsDescriptor) NewEmptyServerMsg() proto.Message {
	return &GetHopReportsResponse{}
}

func (d *GetHopReportsDescriptor) IsUnary() bool {
	return true
}

func (d *GetHopReportsDescriptor) IsClientStream() bool {
	return false
}

func (d *GetHopReportsDescriptor) IsServerStream() bool {
	return false
}

func (d *GetHopReportsDescriptor) IsCollection() bool {
	return false
}

func (d *GetHopReportsDescriptor) IsPlural() bool {
	return false
}

func (d *GetHopReportsDescriptor) HasResource() bool {
	return false
}

func (d *GetHopReportsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *GetHopReportsDescriptor) GetVerb() string {
	return "getHopReports"
}

func (d *GetHopReportsDescriptor) GetMethodName() string {
	return "GetHopReports"
}

func (d *GetHopReportsDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.HopReportService/GetHopReports"
}

func (d *GetHopReportsDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *GetHopReportsDescriptor) GetApiName() string {
	return "HopReportService"
}

func (d *GetHopReportsDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *GetHopReportsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *GetHopReportsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return hopReportServiceDescriptor
}

func (d *GetHopReportsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return nil
}

func (d *GetHopReportsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetHopReportsDescriptorClientMsgHandle{}
}

func (d *GetHopReportsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetHopReportsDescriptorServerMsgHandle{}
}

func (h *GetHopReportsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetHopReportsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	return nil
}

func (h *GetHopReportsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetHopReportsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetHopReportsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	return nil
}

func (h *GetHopReportsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetGetHopReportsDescriptor() *GetHopReportsDescriptor {
	return getHopReportsDescriptor
}

type HopReportServiceDescriptor struct{}

func (d *HopReportServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		reportHopsDescriptor,
		getHopReportsDescriptor,
	}
}

func (d *HopReportServiceDescriptor) GetFullAPIName() string {
	return "/ntt.watchdog.v1alpha2.HopReportService"
}

func (d *HopReportServiceDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *HopReportServiceDescriptor) GetApiName() string {
	return "HopReportService"
}

func (d *HopReportServiceDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *HopReportServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetHopReportServiceDescriptor() *HopReportServiceDescriptor {
	return hopReportServiceDescriptor
}

func initDescriptors() {
	hopReportServiceDescriptor = &HopReportServiceDescriptor{}
	reportHopsDescriptor = &ReportHopsDescriptor{}
	getHopReportsDescriptor = &GetHopReportsDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(hopReportServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(reportHopsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(getHopReportsDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
