// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/project.proto
// DO NOT EDIT!!!

package project

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
	policy "github.com/cloudwan/edgelq-sdk/meta/multi_region/proto/policy"
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
	_ = &policy.Policy{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Project_ProbeStatusExposureMode int32

const (
	// Same as Private
	Project_ProbeStatusExposureModeUnspecified Project_ProbeStatusExposureMode = 0
	// Probe status is not available for the public. Authentication and
	// authorization is required.
	Project_Private Project_ProbeStatusExposureMode = 1
	// Any internet users can access prove status including provisioning status,
	// connection status and ISP names, from serial numbers without any
	// authentication.
	Project_InsecurePublic Project_ProbeStatusExposureMode = 2
)

var (
	Project_ProbeStatusExposureMode_name = map[int32]string{
		0: "ProbeStatusExposureModeUnspecified",
		1: "Private",
		2: "InsecurePublic",
	}

	Project_ProbeStatusExposureMode_value = map[string]int32{
		"ProbeStatusExposureModeUnspecified": 0,
		"Private":                            1,
		"InsecurePublic":                     2,
	}
)

func (x Project_ProbeStatusExposureMode) Enum() *Project_ProbeStatusExposureMode {
	p := new(Project_ProbeStatusExposureMode)
	*p = x
	return p
}

func (x Project_ProbeStatusExposureMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (Project_ProbeStatusExposureMode) Descriptor() preflect.EnumDescriptor {
	return watchdog_proto_v1alpha2_project_proto_enumTypes[0].Descriptor()
}

func (Project_ProbeStatusExposureMode) Type() preflect.EnumType {
	return &watchdog_proto_v1alpha2_project_proto_enumTypes[0]
}

func (x Project_ProbeStatusExposureMode) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use Project_ProbeStatusExposureMode.ProtoReflect.Descriptor instead.
func (Project_ProbeStatusExposureMode) EnumDescriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_project_proto_rawDescGZIP(), []int{0, 0}
}

// Project Resource
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of Project
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
	// Multi region policy
	MultiRegionPolicy *policy.Policy `protobuf:"bytes,6,opt,name=multi_region_policy,json=multiRegionPolicy,proto3" json:"multi_region_policy,omitempty" firestore:"multiRegionPolicy"`
	// Behaviour mode for probeStatus.
	// Note that this configuration affects all Proe resources in the project,
	// To make this option work, enable PublicDeviceInfo
	ProbeStatusExposureMode Project_ProbeStatusExposureMode `protobuf:"varint,2,opt,name=probe_status_exposure_mode,json=probeStatusExposureMode,proto3,enum=ntt.watchdog.v1alpha2.Project_ProbeStatusExposureMode" json:"probe_status_exposure_mode,omitempty" firestore:"probeStatusExposureMode"`
	// InternetQualityRatings will be posted to this Slack webook URL when
	// configured.
	SlackWebhookUrl string `protobuf:"bytes,3,opt,name=slack_webhook_url,json=slackWebhookUrl,proto3" json:"slack_webhook_url,omitempty" firestore:"slackWebhookUrl"`
	// InternetQualityRatings will be posted to this Microsoft Teams webook URL
	// when configured.
	TeamsWebhookUrl string `protobuf:"bytes,4,opt,name=teams_webhook_url,json=teamsWebhookUrl,proto3" json:"teams_webhook_url,omitempty" firestore:"teamsWebhookUrl"`
	// The name of target group that is used to calculate
	// "Interenet Destination" section.
	InternetSummaryTargetGroup string `protobuf:"bytes,7,opt,name=internet_summary_target_group,json=internetSummaryTargetGroup,proto3" json:"internet_summary_target_group,omitempty" firestore:"internetSummaryTargetGroup"`
	// Preferred locale for the project
	// https://iso639-3.sil.org/sites/iso639-3/files/downloads/iso-639-3.tab
	PreferredLocale *Project_Locale `protobuf:"bytes,8,opt,name=preferred_locale,json=preferredLocale,proto3" json:"preferred_locale,omitempty" firestore:"preferredLocale"`
	// URI could be stun or http
	ExternalIpCheckUri []string `protobuf:"bytes,9,rep,name=external_ip_check_uri,json=externalIpCheckUri,proto3" json:"external_ip_check_uri,omitempty" firestore:"externalIpCheckUri"`
}

func (m *Project) Reset() {
	*m = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Project) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Project) ProtoMessage() {}

func (m *Project) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Project) GotenMessage() {}

// Deprecated, Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_project_proto_rawDescGZIP(), []int{0}
}

func (m *Project) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Project) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Project) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Project) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Project) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Project) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Project) GetMultiRegionPolicy() *policy.Policy {
	if m != nil {
		return m.MultiRegionPolicy
	}
	return nil
}

func (m *Project) GetProbeStatusExposureMode() Project_ProbeStatusExposureMode {
	if m != nil {
		return m.ProbeStatusExposureMode
	}
	return Project_ProbeStatusExposureModeUnspecified
}

func (m *Project) GetSlackWebhookUrl() string {
	if m != nil {
		return m.SlackWebhookUrl
	}
	return ""
}

func (m *Project) GetTeamsWebhookUrl() string {
	if m != nil {
		return m.TeamsWebhookUrl
	}
	return ""
}

func (m *Project) GetInternetSummaryTargetGroup() string {
	if m != nil {
		return m.InternetSummaryTargetGroup
	}
	return ""
}

func (m *Project) GetPreferredLocale() *Project_Locale {
	if m != nil {
		return m.PreferredLocale
	}
	return nil
}

func (m *Project) GetExternalIpCheckUri() []string {
	if m != nil {
		return m.ExternalIpCheckUri
	}
	return nil
}

func (m *Project) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "Project"))
	}
	m.Name = fv
}

func (m *Project) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "Project"))
	}
	m.Metadata = fv
}

func (m *Project) SetMultiRegionPolicy(fv *policy.Policy) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MultiRegionPolicy", "Project"))
	}
	m.MultiRegionPolicy = fv
}

func (m *Project) SetProbeStatusExposureMode(fv Project_ProbeStatusExposureMode) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ProbeStatusExposureMode", "Project"))
	}
	m.ProbeStatusExposureMode = fv
}

func (m *Project) SetSlackWebhookUrl(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SlackWebhookUrl", "Project"))
	}
	m.SlackWebhookUrl = fv
}

func (m *Project) SetTeamsWebhookUrl(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TeamsWebhookUrl", "Project"))
	}
	m.TeamsWebhookUrl = fv
}

func (m *Project) SetInternetSummaryTargetGroup(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "InternetSummaryTargetGroup", "Project"))
	}
	m.InternetSummaryTargetGroup = fv
}

func (m *Project) SetPreferredLocale(fv *Project_Locale) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreferredLocale", "Project"))
	}
	m.PreferredLocale = fv
}

func (m *Project) SetExternalIpCheckUri(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ExternalIpCheckUri", "Project"))
	}
	m.ExternalIpCheckUri = fv
}

type Project_Locale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// ISO 639-3 3 character language code
	LangugageCode string `protobuf:"bytes,1,opt,name=langugage_code,json=langugageCode,proto3" json:"langugage_code,omitempty" firestore:"langugageCode"`
}

func (m *Project_Locale) Reset() {
	*m = Project_Locale{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *Project_Locale) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*Project_Locale) ProtoMessage() {}

func (m *Project_Locale) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*Project_Locale) GotenMessage() {}

// Deprecated, Use Project_Locale.ProtoReflect.Descriptor instead.
func (*Project_Locale) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_project_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Project_Locale) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *Project_Locale) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *Project_Locale) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *Project_Locale) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *Project_Locale) GetLangugageCode() string {
	if m != nil {
		return m.LangugageCode
	}
	return ""
}

func (m *Project_Locale) SetLangugageCode(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LangugageCode", "Project_Locale"))
	}
	m.LangugageCode = fv
}

var watchdog_proto_v1alpha2_project_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_project_proto_rawDesc = []byte{
	0x0a, 0x25, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x40, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x95, 0x08, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x23,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xb2, 0xda,
	0x21, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x57, 0x0a, 0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x08, 0xca, 0xc6,
	0x27, 0x04, 0x62, 0x02, 0x08, 0x01, 0x52, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x73, 0x0a, 0x1a, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75,
	0x72, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x62, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x17, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6c, 0x61, 0x63, 0x6b,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x57, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x41, 0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x50, 0x0a, 0x10, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x15, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x72, 0x69, 0x1a, 0x3b,
	0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x31, 0x0a, 0x0e, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x67, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xca, 0xc6, 0x27, 0x06, 0x2a, 0x04, 0x12, 0x02, 0x08, 0x03, 0x52, 0x0d, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x67, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x62, 0x0a, 0x17, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75,
	0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x49,
	0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x10, 0x02, 0x3a,
	0xa8, 0x02, 0xea, 0x41, 0x31, 0x0a, 0x1b, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x92, 0xd9, 0x21, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0xb2, 0xdf, 0x21, 0x0a, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x56, 0x69,
	0x65, 0x77, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xca, 0xa3,
	0x22, 0x4d, 0x0a, 0x25, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x11, 0x73, 0x6c, 0x61, 0x63, 0x6b,
	0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x11, 0x74, 0x65,
	0x61, 0x6d, 0x73, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0xc2,
	0x85, 0x2c, 0x67, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1a, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x11, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x5f, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x11, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x5f,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x13, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xa1, 0x03, 0xe8, 0xde, 0x21,
	0x01, 0xd2, 0xff, 0xd0, 0x02, 0x44, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x44, 0x0a, 0x0d,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0xf2, 0x85, 0xd1, 0x02, 0x4c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x64, 0x62, 0x5f, 0x73, 0x79, 0x6e, 0x63,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0xa2, 0x80, 0xd1, 0x02, 0x46, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_project_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_project_proto_rawDescData = watchdog_proto_v1alpha2_project_proto_rawDesc
)

func watchdog_proto_v1alpha2_project_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_project_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_project_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_project_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_project_proto_rawDescData
}

var watchdog_proto_v1alpha2_project_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var watchdog_proto_v1alpha2_project_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var watchdog_proto_v1alpha2_project_proto_goTypes = []interface{}{
	(Project_ProbeStatusExposureMode)(0), // 0: ntt.watchdog.v1alpha2.Project_ProbeStatusExposureMode
	(*Project)(nil),                      // 1: ntt.watchdog.v1alpha2.Project
	(*Project_Locale)(nil),               // 2: ntt.watchdog.v1alpha2.Project.Locale
	(*ntt_meta.Meta)(nil),                // 3: ntt.types.Meta
	(*policy.Policy)(nil),                // 4: ntt.meta.multi_region.Policy
}
var watchdog_proto_v1alpha2_project_proto_depIdxs = []int32{
	3, // 0: ntt.watchdog.v1alpha2.Project.metadata:type_name -> ntt.types.Meta
	4, // 1: ntt.watchdog.v1alpha2.Project.multi_region_policy:type_name -> ntt.meta.multi_region.Policy
	0, // 2: ntt.watchdog.v1alpha2.Project.probe_status_exposure_mode:type_name -> ntt.watchdog.v1alpha2.Project_ProbeStatusExposureMode
	2, // 3: ntt.watchdog.v1alpha2.Project.preferred_locale:type_name -> ntt.watchdog.v1alpha2.Project.Locale
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_project_proto_init() }
func watchdog_proto_v1alpha2_project_proto_init() {
	if watchdog_proto_v1alpha2_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		watchdog_proto_v1alpha2_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project_Locale); i {
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
			RawDescriptor: watchdog_proto_v1alpha2_project_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_project_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_project_proto_depIdxs,
		EnumInfos:         watchdog_proto_v1alpha2_project_proto_enumTypes,
		MessageInfos:      watchdog_proto_v1alpha2_project_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_project_proto = out.File
	watchdog_proto_v1alpha2_project_proto_rawDesc = nil
	watchdog_proto_v1alpha2_project_proto_goTypes = nil
	watchdog_proto_v1alpha2_project_proto_depIdxs = nil
}
