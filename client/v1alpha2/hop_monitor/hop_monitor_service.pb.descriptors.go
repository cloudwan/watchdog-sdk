// Code generated by protoc-gen-goten-client
// API: HopMonitorService
// DO NOT EDIT!!!

package hop_monitor_client

import (
	"google.golang.org/protobuf/proto"

	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(proto.Message)
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &probe.Probe{}
)

var (
	descriptorsInitialized      bool
	hopMonitorServiceDescriptor *HopMonitorServiceDescriptor
	runHopMonitorDescriptor     *RunHopMonitorDescriptor
)

type RunHopMonitorDescriptor struct{}

type RunHopMonitorDescriptorClientMsgHandle struct{}

type RunHopMonitorDescriptorServerMsgHandle struct{}

func (d *RunHopMonitorDescriptor) NewEmptyClientMsg() proto.Message {
	return &RunHopMonitorRequest{}
}

func (d *RunHopMonitorDescriptor) NewEmptyServerMsg() proto.Message {
	return &RunHopMonitorResponse{}
}

func (d *RunHopMonitorDescriptor) IsUnary() bool {
	return false
}

func (d *RunHopMonitorDescriptor) IsClientStream() bool {
	return false
}

func (d *RunHopMonitorDescriptor) IsServerStream() bool {
	return true
}

func (d *RunHopMonitorDescriptor) IsCollection() bool {
	return false
}

func (d *RunHopMonitorDescriptor) IsPlural() bool {
	return false
}

func (d *RunHopMonitorDescriptor) HasResource() bool {
	return true
}

func (d *RunHopMonitorDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *RunHopMonitorDescriptor) GetVerb() string {
	return "runHopMonitor"
}

func (d *RunHopMonitorDescriptor) GetMethodName() string {
	return "RunHopMonitor"
}

func (d *RunHopMonitorDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.HopMonitorService/RunHopMonitor"
}

func (d *RunHopMonitorDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *RunHopMonitorDescriptor) GetApiName() string {
	return "HopMonitorService"
}

func (d *RunHopMonitorDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *RunHopMonitorDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *RunHopMonitorDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return hopMonitorServiceDescriptor
}

func (d *RunHopMonitorDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return probe.GetDescriptor()
}

func (d *RunHopMonitorDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &RunHopMonitorDescriptorClientMsgHandle{}
}

func (d *RunHopMonitorDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &RunHopMonitorDescriptorServerMsgHandle{}
}

func (h *RunHopMonitorDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*RunHopMonitorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*RunHopMonitorRequest) *probe.Name
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

func (h *RunHopMonitorDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*RunHopMonitorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*RunHopMonitorRequest) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *RunHopMonitorDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*RunHopMonitorRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*RunHopMonitorRequest) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *RunHopMonitorDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*RunHopMonitorResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*RunHopMonitorResponse) *probe.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *RunHopMonitorDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*RunHopMonitorResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*RunHopMonitorResponse) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *RunHopMonitorDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*RunHopMonitorResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*RunHopMonitorResponse) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func GetRunHopMonitorDescriptor() *RunHopMonitorDescriptor {
	return runHopMonitorDescriptor
}

type HopMonitorServiceDescriptor struct{}

func (d *HopMonitorServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		runHopMonitorDescriptor,
	}
}

func (d *HopMonitorServiceDescriptor) GetFullAPIName() string {
	return "/ntt.watchdog.v1alpha2.HopMonitorService"
}

func (d *HopMonitorServiceDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *HopMonitorServiceDescriptor) GetApiName() string {
	return "HopMonitorService"
}

func (d *HopMonitorServiceDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *HopMonitorServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetHopMonitorServiceDescriptor() *HopMonitorServiceDescriptor {
	return hopMonitorServiceDescriptor
}

func initDescriptors() {
	hopMonitorServiceDescriptor = &HopMonitorServiceDescriptor{}
	runHopMonitorDescriptor = &RunHopMonitorDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(hopMonitorServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(runHopMonitorDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
