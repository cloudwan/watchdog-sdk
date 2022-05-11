// Code generated by protoc-gen-goten-client
// API: GeoResolverService
// DO NOT EDIT!!!

package geo_resolver_client

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
	descriptorsInitialized       bool
	geoResolverServiceDescriptor *GeoResolverServiceDescriptor
	resolveGeoIPDescriptor       *ResolveGeoIPDescriptor
	bulkResolveGeoIPDescriptor   *BulkResolveGeoIPDescriptor
	resolveEnvironmentDescriptor *ResolveEnvironmentDescriptor
)

type ResolveGeoIPDescriptor struct{}

type ResolveGeoIPDescriptorClientMsgHandle struct{}

type ResolveGeoIPDescriptorServerMsgHandle struct{}

func (d *ResolveGeoIPDescriptor) NewEmptyClientMsg() proto.Message {
	return &ResolveGeoIPRequest{}
}

func (d *ResolveGeoIPDescriptor) NewEmptyServerMsg() proto.Message {
	return &ResolveGeoIPResponse{}
}

func (d *ResolveGeoIPDescriptor) IsUnary() bool {
	return true
}

func (d *ResolveGeoIPDescriptor) IsClientStream() bool {
	return false
}

func (d *ResolveGeoIPDescriptor) IsServerStream() bool {
	return false
}

func (d *ResolveGeoIPDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *ResolveGeoIPDescriptor) IsPluralSubject() bool {
	return false
}

func (d *ResolveGeoIPDescriptor) HasSubjectResource() bool {
	return true
}

func (d *ResolveGeoIPDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ResolveGeoIPDescriptor) GetVerb() string {
	return "resolveGeoIP"
}

func (d *ResolveGeoIPDescriptor) GetMethodName() string {
	return "ResolveGeoIP"
}

func (d *ResolveGeoIPDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.GeoResolverService/ResolveGeoIP"
}

func (d *ResolveGeoIPDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *ResolveGeoIPDescriptor) GetApiName() string {
	return "GeoResolverService"
}

func (d *ResolveGeoIPDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *ResolveGeoIPDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ResolveGeoIPDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return geoResolverServiceDescriptor
}

func (d *ResolveGeoIPDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return probe.GetDescriptor()
}

func (d *ResolveGeoIPDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ResolveGeoIPDescriptorClientMsgHandle{}
}

func (d *ResolveGeoIPDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ResolveGeoIPDescriptorServerMsgHandle{}
}

func (h *ResolveGeoIPDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ResolveGeoIPRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*ResolveGeoIPRequest) *probe.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if ref := typedMsg.GetName(); ref != nil {
		return &ref.Name
	}
	return (*probe.Name)(nil)
}

func (h *ResolveGeoIPDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ResolveGeoIPRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*ResolveGeoIPRequest) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *ResolveGeoIPDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ResolveGeoIPRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectCollectionName(*ResolveGeoIPRequest) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractSubjectCollectionName(typedMsg)
	}
	return nil
}

func (h *ResolveGeoIPDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ResolveGeoIPResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*ResolveGeoIPResponse) *probe.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *ResolveGeoIPDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ResolveGeoIPResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*ResolveGeoIPResponse) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *ResolveGeoIPDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ResolveGeoIPResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectCollectionName(*ResolveGeoIPResponse) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractSubjectCollectionName(typedMsg)
	}
	return nil
}

func GetResolveGeoIPDescriptor() *ResolveGeoIPDescriptor {
	return resolveGeoIPDescriptor
}

type BulkResolveGeoIPDescriptor struct{}

type BulkResolveGeoIPDescriptorClientMsgHandle struct{}

type BulkResolveGeoIPDescriptorServerMsgHandle struct{}

func (d *BulkResolveGeoIPDescriptor) NewEmptyClientMsg() proto.Message {
	return &BulkResolveGeoIPRequest{}
}

func (d *BulkResolveGeoIPDescriptor) NewEmptyServerMsg() proto.Message {
	return &BulkResolveGeoIPResponse{}
}

func (d *BulkResolveGeoIPDescriptor) IsUnary() bool {
	return true
}

func (d *BulkResolveGeoIPDescriptor) IsClientStream() bool {
	return false
}

func (d *BulkResolveGeoIPDescriptor) IsServerStream() bool {
	return false
}

func (d *BulkResolveGeoIPDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *BulkResolveGeoIPDescriptor) IsPluralSubject() bool {
	return false
}

func (d *BulkResolveGeoIPDescriptor) HasSubjectResource() bool {
	return true
}

func (d *BulkResolveGeoIPDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *BulkResolveGeoIPDescriptor) GetVerb() string {
	return "bulkResolveGeoIP"
}

func (d *BulkResolveGeoIPDescriptor) GetMethodName() string {
	return "BulkResolveGeoIP"
}

func (d *BulkResolveGeoIPDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.GeoResolverService/BulkResolveGeoIP"
}

func (d *BulkResolveGeoIPDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *BulkResolveGeoIPDescriptor) GetApiName() string {
	return "GeoResolverService"
}

func (d *BulkResolveGeoIPDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *BulkResolveGeoIPDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *BulkResolveGeoIPDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return geoResolverServiceDescriptor
}

func (d *BulkResolveGeoIPDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return probe.GetDescriptor()
}

func (d *BulkResolveGeoIPDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &BulkResolveGeoIPDescriptorClientMsgHandle{}
}

func (d *BulkResolveGeoIPDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &BulkResolveGeoIPDescriptorServerMsgHandle{}
}

func (h *BulkResolveGeoIPDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BulkResolveGeoIPRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*BulkResolveGeoIPRequest) *probe.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if ref := typedMsg.GetName(); ref != nil {
		return &ref.Name
	}
	return (*probe.Name)(nil)
}

func (h *BulkResolveGeoIPDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*BulkResolveGeoIPRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*BulkResolveGeoIPRequest) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *BulkResolveGeoIPDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BulkResolveGeoIPRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectCollectionName(*BulkResolveGeoIPRequest) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractSubjectCollectionName(typedMsg)
	}
	return nil
}

func (h *BulkResolveGeoIPDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BulkResolveGeoIPResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*BulkResolveGeoIPResponse) *probe.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *BulkResolveGeoIPDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*BulkResolveGeoIPResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*BulkResolveGeoIPResponse) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *BulkResolveGeoIPDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*BulkResolveGeoIPResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectCollectionName(*BulkResolveGeoIPResponse) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractSubjectCollectionName(typedMsg)
	}
	return nil
}

func GetBulkResolveGeoIPDescriptor() *BulkResolveGeoIPDescriptor {
	return bulkResolveGeoIPDescriptor
}

type ResolveEnvironmentDescriptor struct{}

type ResolveEnvironmentDescriptorClientMsgHandle struct{}

type ResolveEnvironmentDescriptorServerMsgHandle struct{}

func (d *ResolveEnvironmentDescriptor) NewEmptyClientMsg() proto.Message {
	return &ResolveEnvironmentRequest{}
}

func (d *ResolveEnvironmentDescriptor) NewEmptyServerMsg() proto.Message {
	return &ResolveEnvironmentResponse{}
}

func (d *ResolveEnvironmentDescriptor) IsUnary() bool {
	return true
}

func (d *ResolveEnvironmentDescriptor) IsClientStream() bool {
	return false
}

func (d *ResolveEnvironmentDescriptor) IsServerStream() bool {
	return false
}

func (d *ResolveEnvironmentDescriptor) IsCollectionSubject() bool {
	return false
}

func (d *ResolveEnvironmentDescriptor) IsPluralSubject() bool {
	return false
}

func (d *ResolveEnvironmentDescriptor) HasSubjectResource() bool {
	return true
}

func (d *ResolveEnvironmentDescriptor) RequestHasResourceBody() bool {
	return false
}

func (d *ResolveEnvironmentDescriptor) GetVerb() string {
	return "resolveEnvironment"
}

func (d *ResolveEnvironmentDescriptor) GetMethodName() string {
	return "ResolveEnvironment"
}

func (d *ResolveEnvironmentDescriptor) GetFullMethodName() string {
	return "/ntt.watchdog.v1alpha2.GeoResolverService/ResolveEnvironment"
}

func (d *ResolveEnvironmentDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *ResolveEnvironmentDescriptor) GetApiName() string {
	return "GeoResolverService"
}

func (d *ResolveEnvironmentDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *ResolveEnvironmentDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func (d *ResolveEnvironmentDescriptor) GetApiDescriptor() gotenclient.ApiDescriptor {
	return geoResolverServiceDescriptor
}

func (d *ResolveEnvironmentDescriptor) GetSubjectResourceDescriptor() gotenresource.Descriptor {
	return probe.GetDescriptor()
}

func (d *ResolveEnvironmentDescriptor) GetClientMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ResolveEnvironmentDescriptorClientMsgHandle{}
}

func (d *ResolveEnvironmentDescriptor) GetServerMsgReflectHandle() gotenclient.MethodMsgHandle {
	return &ResolveEnvironmentDescriptorServerMsgHandle{}
}

func (h *ResolveEnvironmentDescriptorClientMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ResolveEnvironmentRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*ResolveEnvironmentRequest) *probe.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	if ref := typedMsg.GetName(); ref != nil {
		return &ref.Name
	}
	return (*probe.Name)(nil)
}

func (h *ResolveEnvironmentDescriptorClientMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ResolveEnvironmentRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*ResolveEnvironmentRequest) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *ResolveEnvironmentDescriptorClientMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ResolveEnvironmentRequest)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectCollectionName(*ResolveEnvironmentRequest) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractSubjectCollectionName(typedMsg)
	}
	return nil
}

func (h *ResolveEnvironmentDescriptorServerMsgHandle) ExtractSubjectResourceName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ResolveEnvironmentResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceName(*ResolveEnvironmentResponse) *probe.Name
	})
	if ok {
		return override.OverrideExtractSubjectResourceName(typedMsg)
	}
	return nil
}

func (h *ResolveEnvironmentDescriptorServerMsgHandle) ExtractSubjectResourceNames(msg proto.Message) gotenresource.NameList {
	typedMsg := msg.(*ResolveEnvironmentResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectResourceNames(*ResolveEnvironmentResponse) []*probe.Name
	})
	if ok {
		return probe.ProbeNameList(override.OverrideExtractSubjectResourceNames(typedMsg))
	}
	return nil
}

func (h *ResolveEnvironmentDescriptorServerMsgHandle) ExtractSubjectCollectionName(msg proto.Message) gotenresource.Name {
	typedMsg := msg.(*ResolveEnvironmentResponse)
	var asInterface interface{} = h
	override, ok := asInterface.(interface {
		OverrideExtractSubjectCollectionName(*ResolveEnvironmentResponse) *probe.ParentName
	})
	if ok {
		return override.OverrideExtractSubjectCollectionName(typedMsg)
	}
	return nil
}

func GetResolveEnvironmentDescriptor() *ResolveEnvironmentDescriptor {
	return resolveEnvironmentDescriptor
}

type GeoResolverServiceDescriptor struct{}

func (d *GeoResolverServiceDescriptor) AllMethodDescriptors() []gotenclient.MethodDescriptor {
	return []gotenclient.MethodDescriptor{
		resolveGeoIPDescriptor,
		bulkResolveGeoIPDescriptor,
		resolveEnvironmentDescriptor,
	}
}

func (d *GeoResolverServiceDescriptor) GetFullAPIName() string {
	return "/ntt.watchdog.v1alpha2.GeoResolverService"
}

func (d *GeoResolverServiceDescriptor) GetProtoPkgName() string {
	return "ntt.watchdog.v1alpha2"
}

func (d *GeoResolverServiceDescriptor) GetApiName() string {
	return "GeoResolverService"
}

func (d *GeoResolverServiceDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *GeoResolverServiceDescriptor) GetServiceVersion() string {
	return "v1alpha2"
}

func GetGeoResolverServiceDescriptor() *GeoResolverServiceDescriptor {
	return geoResolverServiceDescriptor
}

func initDescriptors() {
	geoResolverServiceDescriptor = &GeoResolverServiceDescriptor{}
	resolveGeoIPDescriptor = &ResolveGeoIPDescriptor{}
	bulkResolveGeoIPDescriptor = &BulkResolveGeoIPDescriptor{}
	resolveEnvironmentDescriptor = &ResolveEnvironmentDescriptor{}
	gotenclient.GetRegistry().RegisterApiDescriptor(geoResolverServiceDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(resolveGeoIPDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(bulkResolveGeoIPDescriptor)
	gotenclient.GetRegistry().RegisterMethodDescriptor(resolveEnvironmentDescriptor)
}

func init() {
	if !descriptorsInitialized {
		initDescriptors()
		descriptorsInitialized = true
	}
}
