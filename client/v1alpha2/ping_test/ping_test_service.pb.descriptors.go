// Code generated by protoc-gen-goten-client
// API: PingTestService
// DO NOT EDIT!!!

package ping_test_client

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
	_ = proto.Message(nil)
	_ = gotenclient.MethodDescriptor(nil)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &probe.Probe{}
)

var (
	descriptorsInitialized    bool
	pingTestServiceDescriptor *PingTestServiceDescriptor
	runPingTestDescriptor     *RunPingTestDescriptor
)

type RunPingTestDescriptor struct{}

type RunPingTestDescriptorClientMsgHandle struct{}

type RunPingTestDescriptorServerMsgHandle struct{}

func (d *RunPingTestDescriptor) NewEmptyClientMsg() proto.Message {
	return &RunPingTestRequest{}
}

func (d *RunPingTestDescriptor) NewEmptyServerMsg() proto.Message {
	return &RunPingTestResponse{}
}

func (d *RunPingTestDescriptor) IsUnary() bool {
	return false
}

func (d *RunPingTestDescriptor) IsClientStream() bool {
	return false
}

func (d *RunPingTestDescriptor) IsServerStream() bool {
	return true
}

func (d *RunPingTestDescriptor) IsCollection() bool {
	return false
}

func (d *RunPingTestDescriptor) IsPlural() bool {
	return false
}

func (d *RunPingTestDescriptor) HasResource() bool {
	return true
}

func (d *RunPingTestDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *RunPingTestDescriptor) GetVerb() string {
	return "runPingTest"
}

func (d *RunPingTestDescriptor) GetMethodName() string {
	return "RunPingTest"
}

func (d *RunPingTestDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.PingTestService/RunPingTest"
}

func (d *RunPingTestDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *RunPingTestDescriptor) GetApiName() string {
	return "PingTestService"
}

func (d *RunPingTestDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *RunPingTestDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *RunPingTestDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return pingTestServiceDescriptor
}

func (d *RunPingTestDescriptor) GetResourceDescriptor() gotenresource.Descriptor {
	return probe.GetDescriptor()
}

func (d *RunPingTestDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &RunPingTestDescriptorClientMsgHandle{}
}

func (d *RunPingTestDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &RunPingTestDescriptorServerMsgHandle{}
}

func (h *RunPingTestDescriptorClientMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*RunPingTestRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*RunPingTestRequest) *probe.Name
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

func (h *RunPingTestDescriptorClientMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*RunPingTestRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*RunPingTestRequest) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *RunPingTestDescriptorClientMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*RunPingTestRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*RunPingTestRequest) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func (h *RunPingTestDescriptorServerMsgHandle) ExtractResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*RunPingTestResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceName(*RunPingTestResponse) *probe.Name
	})
	if ok {
		return override.OverrideExtractResourceName(typedMsg)
	}
	return nil
}

func (h *RunPingTestDescriptorServerMsgHandle) ExtractResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*RunPingTestResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractResourceNames(*RunPingTestResponse) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractResourceNames(typedMsg))
	}
	return nil
}

func (h *RunPingTestDescriptorServerMsgHandle) ExtractCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*RunPingTestResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractCollectionName(*RunPingTestResponse) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractCollectionName(typedMsg)
	}
	return nil
}

func GetRunPingTestDescriptor() *RunPingTestDescriptor {
	return runPingTestDescriptor
}

type PingTestServiceDescriptor struct{}

func (d *PingTestServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		runPingTestDescriptor,
	}
}

func (d *PingTestServiceDescriptor) GetFullAPIName() string {
	return "/ntt.watchdog.v1alpha2.PingTestService"
}

func (d *PingTestServiceDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *PingTestServiceDescriptor) GetApiName() string {
	return "PingTestService"
}

func (d *PingTestServiceDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *PingTestServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetPingTestServiceDescriptor() *PingTestServiceDescriptor {
	return pingTestServiceDescriptor
}

func initDescriptors() {
	pingTestServiceDescriptor = &PingTestServiceDescriptor{}
	runPingTestDescriptor = &RunPingTestDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(pingTestServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(runPingTestDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
