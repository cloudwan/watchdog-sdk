// Code generated by protoc-gen-goten-client
// API: HTTPMetricsService
// DO NOT EDIT!!!

package http_metrics_client

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
	descriptorsInitialized       bool
	httpMetricsServiceDescriptor *HTTPMetricsServiceDescriptor
	reportHTTPMetricsDescriptor  *ReportHTTPMetricsDescriptor
	getHTTPMetricsDescriptor     *GetHTTPMetricsDescriptor
)

type ReportHTTPMetricsDescriptor struct{}

type ReportHTTPMetricsDescriptorClientMsgHandle struct{}

type ReportHTTPMetricsDescriptorServerMsgHandle struct{}

func (d *ReportHTTPMetricsDescriptor) NewEmptyClientMsg() proto.Message {
	return &ReportHTTPMetricsRequest{}
}

func (d *ReportHTTPMetricsDescriptor) NewEmptyServerMsg() proto.Message {
	return &empty.Empty{}
}

func (d *ReportHTTPMetricsDescriptor) IsUnary() bool {
	return true
}

func (d *ReportHTTPMetricsDescriptor) IsClientStream() bool {
	return false
}

func (d *ReportHTTPMetricsDescriptor) IsServerStream() bool {
	return false
}

func (d *ReportHTTPMetricsDescriptor) IsCollection() bool {
	return false
}

func (d *ReportHTTPMetricsDescriptor) IsPlural() bool {
	return false
}

func (d *ReportHTTPMetricsDescriptor) HasResource() bool {
	return true
}

func (d *ReportHTTPMetricsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ReportHTTPMetricsDescriptor) GetVerb() string {
	return "reportHTTPMetrics"
}

func (d *ReportHTTPMetricsDescriptor) GetMethodName() string {
	return "ReportHTTPMetrics"
}

func (d *ReportHTTPMetricsDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.HTTPMetricsService/ReportHTTPMetrics"
}

func (d *ReportHTTPMetricsDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *ReportHTTPMetricsDescriptor) GetApiName() string {
	return "HTTPMetricsService"
}

func (d *ReportHTTPMetricsDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *ReportHTTPMetricsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ReportHTTPMetricsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return httpMetricsServiceDescriptor
}

func (d *ReportHTTPMetricsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return probe.GetDescriptor()
}

func (d *ReportHTTPMetricsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ReportHTTPMetricsDescriptorClientMsgHandle{}
}

func (d *ReportHTTPMetricsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ReportHTTPMetricsDescriptorServerMsgHandle{}
}

func (h *ReportHTTPMetricsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ReportHTTPMetricsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*ReportHTTPMetricsRequest) *probe.Name
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

func (h *ReportHTTPMetricsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ReportHTTPMetricsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*ReportHTTPMetricsRequest) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *ReportHTTPMetricsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ReportHTTPMetricsRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*ReportHTTPMetricsRequest) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *ReportHTTPMetricsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
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

func (h *ReportHTTPMetricsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
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

func (h *ReportHTTPMetricsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
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

func GetReportHTTPMetricsDescriptor() *ReportHTTPMetricsDescriptor {
	return reportHTTPMetricsDescriptor
}

type GetHTTPMetricsDescriptor struct{}

type GetHTTPMetricsDescriptorClientMsgHandle struct{}

type GetHTTPMetricsDescriptorServerMsgHandle struct{}

func (d *GetHTTPMetricsDescriptor) NewEmptyClientMsg() proto.Message {
	return &GetHTTPMetricsRequest{}
}

func (d *GetHTTPMetricsDescriptor) NewEmptyServerMsg() proto.Message {
	return &GetHTTPMetricsResponse{}
}

func (d *GetHTTPMetricsDescriptor) IsUnary() bool {
	return true
}

func (d *GetHTTPMetricsDescriptor) IsClientStream() bool {
	return false
}

func (d *GetHTTPMetricsDescriptor) IsServerStream() bool {
	return false
}

func (d *GetHTTPMetricsDescriptor) IsCollection() bool {
	return false
}

func (d *GetHTTPMetricsDescriptor) IsPlural() bool {
	return false
}

func (d *GetHTTPMetricsDescriptor) HasResource() bool {
	return false
}

func (d *GetHTTPMetricsDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *GetHTTPMetricsDescriptor) GetVerb() string {
	return "getHTTPMetrics"
}

func (d *GetHTTPMetricsDescriptor) GetMethodName() string {
	return "GetHTTPMetrics"
}

func (d *GetHTTPMetricsDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.HTTPMetricsService/GetHTTPMetrics"
}

func (d *GetHTTPMetricsDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *GetHTTPMetricsDescriptor) GetApiName() string {
	return "HTTPMetricsService"
}

func (d *GetHTTPMetricsDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *GetHTTPMetricsDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *GetHTTPMetricsDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return httpMetricsServiceDescriptor
}

func (d *GetHTTPMetricsDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return nil
}

func (d *GetHTTPMetricsDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetHTTPMetricsDescriptorClientMsgHandle{}
}

func (d *GetHTTPMetricsDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &GetHTTPMetricsDescriptorServerMsgHandle{}
}

func (h *GetHTTPMetricsDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetHTTPMetricsDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	return nil
}

func (h *GetHTTPMetricsDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetHTTPMetricsDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	return nil
}

func (h *GetHTTPMetricsDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	return nil
}

func (h *GetHTTPMetricsDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	return nil
}

func GetGetHTTPMetricsDescriptor() *GetHTTPMetricsDescriptor {
	return getHTTPMetricsDescriptor
}

type HTTPMetricsServiceDescriptor struct{}

func (d *HTTPMetricsServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		reportHTTPMetricsDescriptor,
		getHTTPMetricsDescriptor,
	}
}

func (d *HTTPMetricsServiceDescriptor) GetFullAPIName() string {
	return "/ntt.watchdog.v1alpha2.HTTPMetricsService"
}

func (d *HTTPMetricsServiceDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *HTTPMetricsServiceDescriptor) GetApiName() string {
	return "HTTPMetricsService"
}

func (d *HTTPMetricsServiceDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *HTTPMetricsServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetHTTPMetricsServiceDescriptor() *HTTPMetricsServiceDescriptor {
	return httpMetricsServiceDescriptor
}

func initDescriptors() {
	httpMetricsServiceDescriptor = &HTTPMetricsServiceDescriptor{}
	reportHTTPMetricsDescriptor = &ReportHTTPMetricsDescriptor{}
	getHTTPMetricsDescriptor = &GetHTTPMetricsDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(httpMetricsServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(reportHTTPMetricsDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(getHTTPMetricsDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
