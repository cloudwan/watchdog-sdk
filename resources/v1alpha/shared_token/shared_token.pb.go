// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha/shared_token.proto
// DO NOT EDIT!!!

package shared_token

import (
	"fmt"
	"reflect"
	"sync"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha/common"
	probe_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha/probe_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha/project"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = reflect.Method{}
	_ = sync.Once{}

	_ = protojson.MarshalOptions{}
	_ = proto.MarshalOptions{}
	_ = preflect.Value{}
	_ = protoimpl.DescBuilder{}
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &common.SoftwareVersion{}
	_ = &probe_group.ProbeGroup{}
	_ = &project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SharedToken Resource
type SharedToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of SharedToken
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Human readable name of SharedToken.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Maximium number of probes provisioned through the SharedToken.
	// Any negative value is interpreted as no quota.
	Quota int32 `protobuf:"varint,3,opt,name=quota,proto3" json:"quota,omitempty" firestore:"quota"`
	// probe_template describes pre defined values for probes to be generated.
	// Note that some fieds will be overritten by the system.
	// Template values may include place holders that will be substituted by the
	// system with values from probes.
	ProbeTemplate *SharedToken_ProbeTemplate `protobuf:"bytes,4,opt,name=probe_template,json=probeTemplate,proto3" json:"probe_template,omitempty" firestore:"probeTemplate"`
	// deduplication_field_paths ensures that if there is an existing Probe
	// resource that matches all of the paths provided, the resource will be
	// reused, instaed of creating a new Probe resource. When the existing Probe
	// resource has a service account provisioned, it will be revoked.
	DeduplicationFieldPaths []string `protobuf:"bytes,5,rep,name=deduplication_field_paths,json=deduplicationFieldPaths,proto3" json:"deduplication_field_paths,omitempty" firestore:"deduplicationFieldPaths"`
	// secret hods the string of the sahred token. Do disclose the value.
	// This filed cannot be updated.
	Secret string `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty" firestore:"secret"`
	// Metadata
	Metadata                   *ntt_meta.Meta `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	AssociateExistingProbeOnly bool           `protobuf:"varint,8,opt,name=associate_existing_probe_only,json=associateExistingProbeOnly,proto3" json:"associate_existing_probe_only,omitempty" firestore:"associateExistingProbeOnly"`
}

func (m *SharedToken) Reset() {
	*m = SharedToken{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha_shared_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SharedToken) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SharedToken) ProtoMessage() {}

func (m *SharedToken) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha_shared_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SharedToken) GotenMessage() {}

// Deprecated, Use SharedToken.ProtoReflect.Descriptor instead.
func (*SharedToken) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha_shared_token_proto_rawDescGZIP(), []int{0}
}

func (m *SharedToken) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SharedToken) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SharedToken) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SharedToken) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SharedToken) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SharedToken) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *SharedToken) GetQuota() int32 {
	if m != nil {
		return m.Quota
	}
	return int32(0)
}

func (m *SharedToken) GetProbeTemplate() *SharedToken_ProbeTemplate {
	if m != nil {
		return m.ProbeTemplate
	}
	return nil
}

func (m *SharedToken) GetDeduplicationFieldPaths() []string {
	if m != nil {
		return m.DeduplicationFieldPaths
	}
	return nil
}

func (m *SharedToken) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *SharedToken) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SharedToken) GetAssociateExistingProbeOnly() bool {
	if m != nil {
		return m.AssociateExistingProbeOnly
	}
	return false
}

func (m *SharedToken) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "SharedToken"))
	}
	m.Name = fv
}

func (m *SharedToken) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "SharedToken"))
	}
	m.DisplayName = fv
}

func (m *SharedToken) SetQuota(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Quota", "SharedToken"))
	}
	m.Quota = fv
}

func (m *SharedToken) SetProbeTemplate(fv *SharedToken_ProbeTemplate) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProbeTemplate", "SharedToken"))
	}
	m.ProbeTemplate = fv
}

func (m *SharedToken) SetDeduplicationFieldPaths(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DeduplicationFieldPaths", "SharedToken"))
	}
	m.DeduplicationFieldPaths = fv
}

func (m *SharedToken) SetSecret(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Secret", "SharedToken"))
	}
	m.Secret = fv
}

func (m *SharedToken) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "SharedToken"))
	}
	m.Metadata = fv
}

func (m *SharedToken) SetAssociateExistingProbeOnly(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AssociateExistingProbeOnly", "SharedToken"))
	}
	m.AssociateExistingProbeOnly = fv
}

type SharedToken_ProbeTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Metadata      *SharedToken_ProbeTemplate_Meta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	DisplayName   string                          `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	Spec          *SharedToken_ProbeTemplate_Spec `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty" firestore:"spec"`
}

func (m *SharedToken_ProbeTemplate) Reset() {
	*m = SharedToken_ProbeTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha_shared_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SharedToken_ProbeTemplate) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SharedToken_ProbeTemplate) ProtoMessage() {}

func (m *SharedToken_ProbeTemplate) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha_shared_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SharedToken_ProbeTemplate) GotenMessage() {}

// Deprecated, Use SharedToken_ProbeTemplate.ProtoReflect.Descriptor instead.
func (*SharedToken_ProbeTemplate) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha_shared_token_proto_rawDescGZIP(), []int{0, 0}
}

func (m *SharedToken_ProbeTemplate) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SharedToken_ProbeTemplate) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SharedToken_ProbeTemplate) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SharedToken_ProbeTemplate) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SharedToken_ProbeTemplate) GetMetadata() *SharedToken_ProbeTemplate_Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SharedToken_ProbeTemplate) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *SharedToken_ProbeTemplate) GetSpec() *SharedToken_ProbeTemplate_Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *SharedToken_ProbeTemplate) SetMetadata(fv *SharedToken_ProbeTemplate_Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "SharedToken_ProbeTemplate"))
	}
	m.Metadata = fv
}

func (m *SharedToken_ProbeTemplate) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "SharedToken_ProbeTemplate"))
	}
	m.DisplayName = fv
}

func (m *SharedToken_ProbeTemplate) SetSpec(fv *SharedToken_ProbeTemplate_Spec) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Spec", "SharedToken_ProbeTemplate"))
	}
	m.Spec = fv
}

type SharedToken_ProbeTemplate_Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Tags          []string          `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" firestore:"tags"`
	Labels        map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"labels"`
	Annotations   map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"annotations"`
}

func (m *SharedToken_ProbeTemplate_Meta) Reset() {
	*m = SharedToken_ProbeTemplate_Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha_shared_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SharedToken_ProbeTemplate_Meta) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SharedToken_ProbeTemplate_Meta) ProtoMessage() {}

func (m *SharedToken_ProbeTemplate_Meta) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha_shared_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SharedToken_ProbeTemplate_Meta) GotenMessage() {}

// Deprecated, Use SharedToken_ProbeTemplate_Meta.ProtoReflect.Descriptor instead.
func (*SharedToken_ProbeTemplate_Meta) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha_shared_token_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (m *SharedToken_ProbeTemplate_Meta) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SharedToken_ProbeTemplate_Meta) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SharedToken_ProbeTemplate_Meta) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SharedToken_ProbeTemplate_Meta) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SharedToken_ProbeTemplate_Meta) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SharedToken_ProbeTemplate_Meta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *SharedToken_ProbeTemplate_Meta) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *SharedToken_ProbeTemplate_Meta) SetTags(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Tags", "SharedToken_ProbeTemplate_Meta"))
	}
	m.Tags = fv
}

func (m *SharedToken_ProbeTemplate_Meta) SetLabels(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Labels", "SharedToken_ProbeTemplate_Meta"))
	}
	m.Labels = fv
}

func (m *SharedToken_ProbeTemplate_Meta) SetAnnotations(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Annotations", "SharedToken_ProbeTemplate_Meta"))
	}
	m.Annotations = fv
}

type SharedToken_ProbeTemplate_Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Probe Group allows optional grouping of probes.
	ProbeGroup *probe_group.Reference `protobuf:"bytes,11,opt,customtype=Reference,name=probe_group,json=probeGroup,proto3" json:"probe_group,omitempty" firestore:"probeGroup"`
	// Address defines the expected address of the probe.
	PrimaryLocation   *common.Location              `protobuf:"bytes,6,opt,name=primary_location,json=primaryLocation,proto3" json:"primary_location,omitempty" firestore:"primaryLocation"`
	LocationDiscovery *common.LocationDiscoverySpec `protobuf:"bytes,8,opt,name=location_discovery,json=locationDiscovery,proto3" json:"location_discovery,omitempty" firestore:"locationDiscovery"`
	// Contact Information
	ContactInfo      *common.ContactInformation `protobuf:"bytes,9,opt,name=contact_info,json=contactInfo,proto3" json:"contact_info,omitempty" firestore:"contactInfo"`
	DisableSpeedtest bool                       `protobuf:"varint,13,opt,name=disable_speedtest,json=disableSpeedtest,proto3" json:"disable_speedtest,omitempty" firestore:"disableSpeedtest"`
}

func (m *SharedToken_ProbeTemplate_Spec) Reset() {
	*m = SharedToken_ProbeTemplate_Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha_shared_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *SharedToken_ProbeTemplate_Spec) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*SharedToken_ProbeTemplate_Spec) ProtoMessage() {}

func (m *SharedToken_ProbeTemplate_Spec) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha_shared_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*SharedToken_ProbeTemplate_Spec) GotenMessage() {}

// Deprecated, Use SharedToken_ProbeTemplate_Spec.ProtoReflect.Descriptor instead.
func (*SharedToken_ProbeTemplate_Spec) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha_shared_token_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (m *SharedToken_ProbeTemplate_Spec) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *SharedToken_ProbeTemplate_Spec) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *SharedToken_ProbeTemplate_Spec) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *SharedToken_ProbeTemplate_Spec) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *SharedToken_ProbeTemplate_Spec) GetProbeGroup() *probe_group.Reference {
	if m != nil {
		return m.ProbeGroup
	}
	return nil
}

func (m *SharedToken_ProbeTemplate_Spec) GetPrimaryLocation() *common.Location {
	if m != nil {
		return m.PrimaryLocation
	}
	return nil
}

func (m *SharedToken_ProbeTemplate_Spec) GetLocationDiscovery() *common.LocationDiscoverySpec {
	if m != nil {
		return m.LocationDiscovery
	}
	return nil
}

func (m *SharedToken_ProbeTemplate_Spec) GetContactInfo() *common.ContactInformation {
	if m != nil {
		return m.ContactInfo
	}
	return nil
}

func (m *SharedToken_ProbeTemplate_Spec) GetDisableSpeedtest() bool {
	if m != nil {
		return m.DisableSpeedtest
	}
	return false
}

func (m *SharedToken_ProbeTemplate_Spec) SetProbeGroup(fv *probe_group.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProbeGroup", "SharedToken_ProbeTemplate_Spec"))
	}
	m.ProbeGroup = fv
}

func (m *SharedToken_ProbeTemplate_Spec) SetPrimaryLocation(fv *common.Location) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PrimaryLocation", "SharedToken_ProbeTemplate_Spec"))
	}
	m.PrimaryLocation = fv
}

func (m *SharedToken_ProbeTemplate_Spec) SetLocationDiscovery(fv *common.LocationDiscoverySpec) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LocationDiscovery", "SharedToken_ProbeTemplate_Spec"))
	}
	m.LocationDiscovery = fv
}

func (m *SharedToken_ProbeTemplate_Spec) SetContactInfo(fv *common.ContactInformation) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ContactInfo", "SharedToken_ProbeTemplate_Spec"))
	}
	m.ContactInfo = fv
}

func (m *SharedToken_ProbeTemplate_Spec) SetDisableSpeedtest(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisableSpeedtest", "SharedToken_ProbeTemplate_Spec"))
	}
	m.DisableSpeedtest = fv
}

var watchdog_proto_v1alpha_shared_token_proto preflect.FileDescriptor

var watchdog_proto_v1alpha_shared_token_proto_rawDesc = []byte{
	0x0a, 0x29, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64,
	0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x0c,
	0x0a, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xb2, 0xda, 0x21,
	0x0f, 0x0a, 0x0d, 0x0a, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x12,
	0x56, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x19, 0x64, 0x65, 0x64, 0x75, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xca, 0xc6, 0x27, 0x04,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x17, 0x64, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x41, 0x0a, 0x1d, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x5f,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x6f,
	0x6e, 0x6c, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x61, 0x73, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x1a, 0x9e, 0x07, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x62, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x1a, 0xe6, 0x02, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0c, 0xca,
	0xc6, 0x27, 0x08, 0x42, 0x06, 0x12, 0x02, 0x08, 0x20, 0x18, 0x01, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x58, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x40, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x67, 0x0a, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x45, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0xe4, 0x02, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x3b, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xb2,
	0xda, 0x21, 0x10, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x10, 0x01, 0xe2, 0xde, 0x21, 0x02, 0x08, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x49, 0x0a, 0x10, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x5a, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x4b, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x74, 0x65, 0x73, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65,
	0x65, 0x64, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x8a, 0x02, 0xea, 0x41, 0x51, 0x0a, 0x1f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x7d, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x7b,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x7d, 0x92, 0xd9, 0x21,
	0x2b, 0x0a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a, 0x07, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xfa, 0xde, 0x21, 0x0d,
	0x0a, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0xb2, 0xdf, 0x21,
	0x0a, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x65, 0x77, 0xda, 0x94, 0x23, 0x08, 0x12,
	0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x58, 0x1a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x22, 0x0e,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x19,
	0x64, 0x65, 0x64, 0x75, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x42, 0xc9, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x4d, 0x0a,
	0x12, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x92, 0x8c, 0xd1, 0x02,
	0x57, 0x0a, 0x17, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x10, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0xd2, 0x84, 0xd1, 0x02, 0x43, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2, 0x80, 0xd1, 0x02, 0x4f,
	0x0a, 0x13, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha_shared_token_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha_shared_token_proto_rawDescData = watchdog_proto_v1alpha_shared_token_proto_rawDesc
)

func watchdog_proto_v1alpha_shared_token_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha_shared_token_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha_shared_token_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha_shared_token_proto_rawDescData)
	})
	return watchdog_proto_v1alpha_shared_token_proto_rawDescData
}

var watchdog_proto_v1alpha_shared_token_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var watchdog_proto_v1alpha_shared_token_proto_goTypes = []interface{}{
	(*SharedToken)(nil),                    // 0: ntt.watchdog.v1alpha.SharedToken
	(*SharedToken_ProbeTemplate)(nil),      // 1: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate
	(*SharedToken_ProbeTemplate_Meta)(nil), // 2: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Meta
	(*SharedToken_ProbeTemplate_Spec)(nil), // 3: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Spec
	nil,                                    // 4: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Meta.LabelsEntry
	nil,                                    // 5: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Meta.AnnotationsEntry
	(*ntt_meta.Meta)(nil),                  // 6: ntt.types.Meta
	(*common.Location)(nil),                // 7: ntt.watchdog.v1alpha.Location
	(*common.LocationDiscoverySpec)(nil),   // 8: ntt.watchdog.v1alpha.LocationDiscoverySpec
	(*common.ContactInformation)(nil),      // 9: ntt.watchdog.v1alpha.ContactInformation
}
var watchdog_proto_v1alpha_shared_token_proto_depIdxs = []int32{
	1, // 0: ntt.watchdog.v1alpha.SharedToken.probe_template:type_name -> ntt.watchdog.v1alpha.SharedToken.ProbeTemplate
	6, // 1: ntt.watchdog.v1alpha.SharedToken.metadata:type_name -> ntt.types.Meta
	2, // 2: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.metadata:type_name -> ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Meta
	3, // 3: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.spec:type_name -> ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Spec
	4, // 4: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Meta.labels:type_name -> ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Meta.LabelsEntry
	5, // 5: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Meta.annotations:type_name -> ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Meta.AnnotationsEntry
	7, // 6: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Spec.primary_location:type_name -> ntt.watchdog.v1alpha.Location
	8, // 7: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Spec.location_discovery:type_name -> ntt.watchdog.v1alpha.LocationDiscoverySpec
	9, // 8: ntt.watchdog.v1alpha.SharedToken.ProbeTemplate.Spec.contact_info:type_name -> ntt.watchdog.v1alpha.ContactInformation
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha_shared_token_proto_init() }
func watchdog_proto_v1alpha_shared_token_proto_init() {
	if watchdog_proto_v1alpha_shared_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha_shared_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedToken); i {
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
		watchdog_proto_v1alpha_shared_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedToken_ProbeTemplate); i {
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
		watchdog_proto_v1alpha_shared_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedToken_ProbeTemplate_Meta); i {
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
		watchdog_proto_v1alpha_shared_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedToken_ProbeTemplate_Spec); i {
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
			RawDescriptor: watchdog_proto_v1alpha_shared_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha_shared_token_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha_shared_token_proto_depIdxs,
		MessageInfos:      watchdog_proto_v1alpha_shared_token_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha_shared_token_proto = out.File
	watchdog_proto_v1alpha_shared_token_proto_rawDesc = nil
	watchdog_proto_v1alpha_shared_token_proto_goTypes = nil
	watchdog_proto_v1alpha_shared_token_proto_depIdxs = nil
}
