// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/probing_target.proto
// DO NOT EDIT!!!

package probing_target

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
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probing_target_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
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
	_ = &probe.Probe{}
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProbingTarget_TargetType int32

const (
	ProbingTarget_UNMANAGED_TARGET ProbingTarget_TargetType = 0
	ProbingTarget_MANAGED_TARGET   ProbingTarget_TargetType = 1
)

var (
	ProbingTarget_TargetType_name = map[int32]string{
		0: "UNMANAGED_TARGET",
		1: "MANAGED_TARGET",
	}

	ProbingTarget_TargetType_value = map[string]int32{
		"UNMANAGED_TARGET": 0,
		"MANAGED_TARGET":   1,
	}
)

func (x ProbingTarget_TargetType) Enum() *ProbingTarget_TargetType {
	p := new(ProbingTarget_TargetType)
	*p = x
	return p
}

func (x ProbingTarget_TargetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (ProbingTarget_TargetType) Descriptor() preflect.EnumDescriptor {
	return watchdog_proto_v1alpha2_probing_target_proto_enumTypes[0].Descriptor()
}

func (ProbingTarget_TargetType) Type() preflect.EnumType {
	return &watchdog_proto_v1alpha2_probing_target_proto_enumTypes[0]
}

func (x ProbingTarget_TargetType) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use ProbingTarget_TargetType.ProtoReflect.Descriptor instead.
func (ProbingTarget_TargetType) EnumDescriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_probing_target_proto_rawDescGZIP(), []int{0, 0}
}

// ProbingTarget Resource
type ProbingTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of ProbingTarget
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Display name is a human readable representation of the target
	DisplayName string `protobuf:"bytes,6,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,13,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// group reference
	Group *probing_target_group.Reference `protobuf:"bytes,17,opt,customtype=Reference,name=group,proto3" json:"group,omitempty" firestore:"group"`
	// mode defines the mode for this probe target - icmp/udp/http/tcp
	Mode common.ProbingMode `protobuf:"varint,2,opt,name=mode,proto3,enum=ntt.watchdog.v1alpha2.ProbingMode" json:"mode,omitempty" firestore:"mode"`
	// ip_version defines which IP version should be used.
	IpVersion common.IpVersion `protobuf:"varint,4,opt,name=ip_version,json=ipVersion,proto3,enum=ntt.watchdog.v1alpha2.IpVersion" json:"ip_version,omitempty" firestore:"ipVersion"`
	// address is the actual target for the probe - IP addreses/http endpoint
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty" firestore:"address"`
	// Additional Category Name to enrich data
	Category string `protobuf:"bytes,12,opt,name=category,proto3" json:"category,omitempty" firestore:"category"`
	// Target location type
	LocationType common.LocationType `protobuf:"varint,14,opt,name=location_type,json=locationType,proto3,enum=ntt.watchdog.v1alpha2.LocationType" json:"location_type,omitempty" firestore:"locationType"`
	// Specific location if applicable
	Location *common.Location `protobuf:"bytes,15,opt,name=location,proto3" json:"location,omitempty" firestore:"location"`
	// per-target constraints
	ProbingConstraint *common.ProbingConstraint `protobuf:"bytes,16,opt,name=probing_constraint,json=probingConstraint,proto3" json:"probing_constraint,omitempty" firestore:"probingConstraint"`
	// Default Probing Settings like interval, tos, window size, path probing
	DefaultProbingSettings *common.ProbingSettings `protobuf:"bytes,18,opt,name=default_probing_settings,json=defaultProbingSettings,proto3" json:"default_probing_settings,omitempty" firestore:"defaultProbingSettings"`
	// Additional configuration for http probing
	HttpProbingConfig *common.HTTPProbingConfig `protobuf:"bytes,19,opt,name=http_probing_config,json=httpProbingConfig,proto3" json:"http_probing_config,omitempty" firestore:"httpProbingConfig"`
	Agent             *probe.Reference          `protobuf:"bytes,20,opt,customtype=Reference,name=agent,proto3" json:"agent,omitempty" firestore:"agent"`
	// addresses list the list of IP addresses that this target is reachable on
	// used for managed targets that might have dual stack
	Addresses  []string                 `protobuf:"bytes,21,rep,name=addresses,proto3" json:"addresses,omitempty" firestore:"addresses"`
	TargetType ProbingTarget_TargetType `protobuf:"varint,22,opt,name=target_type,json=targetType,proto3,enum=ntt.watchdog.v1alpha2.ProbingTarget_TargetType" json:"target_type,omitempty" firestore:"targetType"`
}

func (m *ProbingTarget) Reset() {
	*m = ProbingTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_probing_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProbingTarget) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProbingTarget) ProtoMessage() {}

func (m *ProbingTarget) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_probing_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProbingTarget) GotenMessage() {}

// Deprecated, Use ProbingTarget.ProtoReflect.Descriptor instead.
func (*ProbingTarget) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_probing_target_proto_rawDescGZIP(), []int{0}
}

func (m *ProbingTarget) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProbingTarget) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProbingTarget) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProbingTarget) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProbingTarget) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ProbingTarget) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ProbingTarget) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ProbingTarget) GetGroup() *probing_target_group.Reference {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *ProbingTarget) GetMode() common.ProbingMode {
	if m != nil {
		return m.Mode
	}
	return common.ProbingMode_PROBING_MODE_UNSPECIFIED
}

func (m *ProbingTarget) GetIpVersion() common.IpVersion {
	if m != nil {
		return m.IpVersion
	}
	return common.IpVersion_IP_VERSION_UNSPECIFIED
}

func (m *ProbingTarget) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ProbingTarget) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ProbingTarget) GetLocationType() common.LocationType {
	if m != nil {
		return m.LocationType
	}
	return common.LocationType_LOCATION_TYPE_UNSPECIFIED
}

func (m *ProbingTarget) GetLocation() *common.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *ProbingTarget) GetProbingConstraint() *common.ProbingConstraint {
	if m != nil {
		return m.ProbingConstraint
	}
	return nil
}

func (m *ProbingTarget) GetDefaultProbingSettings() *common.ProbingSettings {
	if m != nil {
		return m.DefaultProbingSettings
	}
	return nil
}

func (m *ProbingTarget) GetHttpProbingConfig() *common.HTTPProbingConfig {
	if m != nil {
		return m.HttpProbingConfig
	}
	return nil
}

func (m *ProbingTarget) GetAgent() *probe.Reference {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *ProbingTarget) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ProbingTarget) GetTargetType() ProbingTarget_TargetType {
	if m != nil {
		return m.TargetType
	}
	return ProbingTarget_UNMANAGED_TARGET
}

func (m *ProbingTarget) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ProbingTarget"))
	}
	m.Name = fv
}

func (m *ProbingTarget) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "ProbingTarget"))
	}
	m.DisplayName = fv
}

func (m *ProbingTarget) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "ProbingTarget"))
	}
	m.Metadata = fv
}

func (m *ProbingTarget) SetGroup(fv *probing_target_group.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Group", "ProbingTarget"))
	}
	m.Group = fv
}

func (m *ProbingTarget) SetMode(fv common.ProbingMode) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Mode", "ProbingTarget"))
	}
	m.Mode = fv
}

func (m *ProbingTarget) SetIpVersion(fv common.IpVersion) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "IpVersion", "ProbingTarget"))
	}
	m.IpVersion = fv
}

func (m *ProbingTarget) SetAddress(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Address", "ProbingTarget"))
	}
	m.Address = fv
}

func (m *ProbingTarget) SetCategory(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Category", "ProbingTarget"))
	}
	m.Category = fv
}

func (m *ProbingTarget) SetLocationType(fv common.LocationType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LocationType", "ProbingTarget"))
	}
	m.LocationType = fv
}

func (m *ProbingTarget) SetLocation(fv *common.Location) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Location", "ProbingTarget"))
	}
	m.Location = fv
}

func (m *ProbingTarget) SetProbingConstraint(fv *common.ProbingConstraint) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProbingConstraint", "ProbingTarget"))
	}
	m.ProbingConstraint = fv
}

func (m *ProbingTarget) SetDefaultProbingSettings(fv *common.ProbingSettings) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DefaultProbingSettings", "ProbingTarget"))
	}
	m.DefaultProbingSettings = fv
}

func (m *ProbingTarget) SetHttpProbingConfig(fv *common.HTTPProbingConfig) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "HttpProbingConfig", "ProbingTarget"))
	}
	m.HttpProbingConfig = fv
}

func (m *ProbingTarget) SetAgent(fv *probe.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Agent", "ProbingTarget"))
	}
	m.Agent = fv
}

func (m *ProbingTarget) SetAddresses(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Addresses", "ProbingTarget"))
	}
	m.Addresses = fv
}

func (m *ProbingTarget) SetTargetType(fv ProbingTarget_TargetType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TargetType", "ProbingTarget"))
	}
	m.TargetType = fv
}

var watchdog_proto_v1alpha2_probing_target_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_probing_target_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e,
	0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc7, 0x0c, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x15, 0xb2, 0xda, 0x21, 0x11, 0x0a, 0x0f, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e,
	0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xb2, 0xda,
	0x21, 0x18, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x10, 0x01, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x43, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x4d,
	0x6f, 0x64, 0x65, 0x42, 0x0b, 0xca, 0xc6, 0x27, 0x07, 0x3a, 0x05, 0x10, 0x01, 0x22, 0x01, 0x00,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x49, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xca, 0xc6,
	0x27, 0x04, 0x3a, 0x02, 0x10, 0x01, 0x52, 0x09, 0x69, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xca, 0xc6, 0x27, 0x04, 0x2a, 0x02, 0x68, 0x01, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x52, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xca,
	0xc6, 0x27, 0x04, 0x3a, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x6a, 0x0a, 0x18, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x08, 0xca, 0xc6, 0x27, 0x04, 0x62, 0x02, 0x08, 0x01, 0x52,
	0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x58, 0x0a, 0x13, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x48, 0x54, 0x54,
	0x50, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11,
	0x68, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x25, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0f, 0xb2, 0xda, 0x21, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x10,
	0x02, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x4d, 0x41, 0x4e, 0x41,
	0x47, 0x45, 0x44, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x44, 0x5f, 0x54, 0x41, 0x52, 0x47, 0x45, 0x54, 0x10, 0x01,
	0x3a, 0xb1, 0x04, 0xea, 0x41, 0x57, 0x0a, 0x21, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x62,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x32, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x70, 0x72,
	0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x7d, 0x92, 0xd9, 0x21,
	0x2f, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x1a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0xaa, 0xd9, 0x21, 0xdb, 0x01, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x0a, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x0a, 0x0d, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x0a, 0x0f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x0a, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0d,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x12,
	0x14, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x1a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x14, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0xad,
	0x01, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x22, 0x0a, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x22, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x12, 0x70, 0x72, 0x6f, 0x62,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x22, 0x18,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x13, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x70,
	0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x05, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xe2, 0xde,
	0x21, 0x02, 0x08, 0x01, 0x4a, 0x04, 0x08, 0x0b, 0x10, 0x0c, 0x42, 0xdf, 0x03, 0xe8, 0xde, 0x21,
	0x01, 0xd2, 0xff, 0xd0, 0x02, 0x52, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e,
	0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x12, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72,
	0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x3b, 0x70, 0x72, 0x6f,
	0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x44,
	0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0xf2, 0x85, 0xd1, 0x02, 0x5a, 0x0a, 0x18, 0x70, 0x72, 0x6f, 0x62, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e,
	0x63, 0x65, 0x72, 0x12, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0xa2, 0x80, 0xd1, 0x02, 0x54, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e,
	0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72,
	0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_probing_target_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_probing_target_proto_rawDescData = watchdog_proto_v1alpha2_probing_target_proto_rawDesc
)

func watchdog_proto_v1alpha2_probing_target_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_probing_target_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_probing_target_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_probing_target_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_probing_target_proto_rawDescData
}

var watchdog_proto_v1alpha2_probing_target_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var watchdog_proto_v1alpha2_probing_target_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var watchdog_proto_v1alpha2_probing_target_proto_goTypes = []interface{}{
	(ProbingTarget_TargetType)(0),    // 0: ntt.watchdog.v1alpha2.ProbingTarget_TargetType
	(*ProbingTarget)(nil),            // 1: ntt.watchdog.v1alpha2.ProbingTarget
	(*ntt_meta.Meta)(nil),            // 2: ntt.types.Meta
	(common.ProbingMode)(0),          // 3: ntt.watchdog.v1alpha2.ProbingMode
	(common.IpVersion)(0),            // 4: ntt.watchdog.v1alpha2.IpVersion
	(common.LocationType)(0),         // 5: ntt.watchdog.v1alpha2.LocationType
	(*common.Location)(nil),          // 6: ntt.watchdog.v1alpha2.Location
	(*common.ProbingConstraint)(nil), // 7: ntt.watchdog.v1alpha2.ProbingConstraint
	(*common.ProbingSettings)(nil),   // 8: ntt.watchdog.v1alpha2.ProbingSettings
	(*common.HTTPProbingConfig)(nil), // 9: ntt.watchdog.v1alpha2.HTTPProbingConfig
}
var watchdog_proto_v1alpha2_probing_target_proto_depIdxs = []int32{
	2, // 0: ntt.watchdog.v1alpha2.ProbingTarget.metadata:type_name -> ntt.types.Meta
	3, // 1: ntt.watchdog.v1alpha2.ProbingTarget.mode:type_name -> ntt.watchdog.v1alpha2.ProbingMode
	4, // 2: ntt.watchdog.v1alpha2.ProbingTarget.ip_version:type_name -> ntt.watchdog.v1alpha2.IpVersion
	5, // 3: ntt.watchdog.v1alpha2.ProbingTarget.location_type:type_name -> ntt.watchdog.v1alpha2.LocationType
	6, // 4: ntt.watchdog.v1alpha2.ProbingTarget.location:type_name -> ntt.watchdog.v1alpha2.Location
	7, // 5: ntt.watchdog.v1alpha2.ProbingTarget.probing_constraint:type_name -> ntt.watchdog.v1alpha2.ProbingConstraint
	8, // 6: ntt.watchdog.v1alpha2.ProbingTarget.default_probing_settings:type_name -> ntt.watchdog.v1alpha2.ProbingSettings
	9, // 7: ntt.watchdog.v1alpha2.ProbingTarget.http_probing_config:type_name -> ntt.watchdog.v1alpha2.HTTPProbingConfig
	0, // 8: ntt.watchdog.v1alpha2.ProbingTarget.target_type:type_name -> ntt.watchdog.v1alpha2.ProbingTarget_TargetType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_probing_target_proto_init() }
func watchdog_proto_v1alpha2_probing_target_proto_init() {
	if watchdog_proto_v1alpha2_probing_target_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_probing_target_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbingTarget); i {
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
			RawDescriptor: watchdog_proto_v1alpha2_probing_target_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_probing_target_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_probing_target_proto_depIdxs,
		EnumInfos:         watchdog_proto_v1alpha2_probing_target_proto_enumTypes,
		MessageInfos:      watchdog_proto_v1alpha2_probing_target_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_probing_target_proto = out.File
	watchdog_proto_v1alpha2_probing_target_proto_rawDesc = nil
	watchdog_proto_v1alpha2_probing_target_proto_goTypes = nil
	watchdog_proto_v1alpha2_probing_target_proto_depIdxs = nil
}
