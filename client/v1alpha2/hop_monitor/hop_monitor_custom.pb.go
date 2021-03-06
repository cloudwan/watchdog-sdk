// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/hop_monitor_custom.proto
// DO NOT EDIT!!!

package hop_monitor_client

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
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probing_target "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target"
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
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target.ProbingTarget{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [RunHopMonitor][ntt.watchdog.v1alpha2.RunHopMonitor]
type RunHopMonitorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  reference of ntt.watchdog.v1alpha2.Probe
	Name *probe.Reference `protobuf:"bytes,1,opt,customtype=Reference,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Source address to listen packet
	// Skip validating goten's ip format
	// in order to accept a zone index for IPv6 link local address
	// ex. fe80::3%eth0
	//
	// If unspecified, agent will pick :: or 0.0.0.0 by default
	// The IP version is chosen according to the version of destination address
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty" firestore:"source"`
	// Target or Destination is required as user input
	Target *probing_target.Reference `protobuf:"bytes,3,opt,customtype=Reference,name=target,proto3" json:"target,omitempty" firestore:"target"`
	// Destination to send packet
	// IP address and hostname are acceptable
	// As the IPv6 address, a zone index is also acceptable
	// Destination or Target is required from user
	Destination string `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty" firestore:"destination"`
	// Byte size of the payload
	// If unspecified, 100 is picked by default
	// The minimum length includes space for timestamp and a tracker
	SizeBytes int32 `protobuf:"varint,5,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty" firestore:"sizeBytes"`
	// Number of attempts to discover unique paths
	Attempts int32              `protobuf:"varint,6,opt,name=attempts,proto3" json:"attempts,omitempty" firestore:"attempts"`
	Mode     common.ProbingMode `protobuf:"varint,7,opt,name=mode,proto3,enum=ntt.watchdog.v1alpha2.ProbingMode" json:"mode,omitempty" firestore:"mode"`
}

func (m *RunHopMonitorRequest) Reset() {
	*m = RunHopMonitorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_hop_monitor_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RunHopMonitorRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RunHopMonitorRequest) ProtoMessage() {}

func (m *RunHopMonitorRequest) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_hop_monitor_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RunHopMonitorRequest) GotenMessage() {}

// Deprecated, Use RunHopMonitorRequest.ProtoReflect.Descriptor instead.
func (*RunHopMonitorRequest) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDescGZIP(), []int{0}
}

func (m *RunHopMonitorRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RunHopMonitorRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RunHopMonitorRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RunHopMonitorRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RunHopMonitorRequest) GetName() *probe.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RunHopMonitorRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *RunHopMonitorRequest) GetTarget() *probing_target.Reference {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *RunHopMonitorRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *RunHopMonitorRequest) GetSizeBytes() int32 {
	if m != nil {
		return m.SizeBytes
	}
	return int32(0)
}

func (m *RunHopMonitorRequest) GetAttempts() int32 {
	if m != nil {
		return m.Attempts
	}
	return int32(0)
}

func (m *RunHopMonitorRequest) GetMode() common.ProbingMode {
	if m != nil {
		return m.Mode
	}
	return common.ProbingMode_PROBING_MODE_UNSPECIFIED
}

func (m *RunHopMonitorRequest) SetName(fv *probe.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RunHopMonitorRequest"))
	}
	m.Name = fv
}

func (m *RunHopMonitorRequest) SetSource(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Source", "RunHopMonitorRequest"))
	}
	m.Source = fv
}

func (m *RunHopMonitorRequest) SetTarget(fv *probing_target.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Target", "RunHopMonitorRequest"))
	}
	m.Target = fv
}

func (m *RunHopMonitorRequest) SetDestination(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Destination", "RunHopMonitorRequest"))
	}
	m.Destination = fv
}

func (m *RunHopMonitorRequest) SetSizeBytes(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SizeBytes", "RunHopMonitorRequest"))
	}
	m.SizeBytes = fv
}

func (m *RunHopMonitorRequest) SetAttempts(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Attempts", "RunHopMonitorRequest"))
	}
	m.Attempts = fv
}

func (m *RunHopMonitorRequest) SetMode(fv common.ProbingMode) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Mode", "RunHopMonitorRequest"))
	}
	m.Mode = fv
}

// Response message for method
// [RunHopMonitor][ntt.watchdog.v1alpha2.RunHopMonitor]
type RunHopMonitorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Paths         []*common.Path             `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty" firestore:"paths"`
	HopStats      map[string]*common.HopStat `protobuf:"bytes,2,rep,name=hop_stats,json=hopStats,proto3" json:"hop_stats,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"hopStats"`
	HopInfo       map[string]*common.HopInfo `protobuf:"bytes,3,rep,name=hop_info,json=hopInfo,proto3" json:"hop_info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"hopInfo"`
	IpVersion     common.IpVersion           `protobuf:"varint,4,opt,name=ip_version,json=ipVersion,proto3,enum=ntt.watchdog.v1alpha2.IpVersion" json:"ip_version,omitempty" firestore:"ipVersion"`
}

func (m *RunHopMonitorResponse) Reset() {
	*m = RunHopMonitorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_hop_monitor_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RunHopMonitorResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RunHopMonitorResponse) ProtoMessage() {}

func (m *RunHopMonitorResponse) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_hop_monitor_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RunHopMonitorResponse) GotenMessage() {}

// Deprecated, Use RunHopMonitorResponse.ProtoReflect.Descriptor instead.
func (*RunHopMonitorResponse) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDescGZIP(), []int{1}
}

func (m *RunHopMonitorResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RunHopMonitorResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RunHopMonitorResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RunHopMonitorResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RunHopMonitorResponse) GetPaths() []*common.Path {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *RunHopMonitorResponse) GetHopStats() map[string]*common.HopStat {
	if m != nil {
		return m.HopStats
	}
	return nil
}

func (m *RunHopMonitorResponse) GetHopInfo() map[string]*common.HopInfo {
	if m != nil {
		return m.HopInfo
	}
	return nil
}

func (m *RunHopMonitorResponse) GetIpVersion() common.IpVersion {
	if m != nil {
		return m.IpVersion
	}
	return common.IpVersion_IP_VERSION_UNSPECIFIED
}

func (m *RunHopMonitorResponse) SetPaths(fv []*common.Path) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Paths", "RunHopMonitorResponse"))
	}
	m.Paths = fv
}

func (m *RunHopMonitorResponse) SetHopStats(fv map[string]*common.HopStat) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "HopStats", "RunHopMonitorResponse"))
	}
	m.HopStats = fv
}

func (m *RunHopMonitorResponse) SetHopInfo(fv map[string]*common.HopInfo) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "HopInfo", "RunHopMonitorResponse"))
	}
	m.HopInfo = fv
}

func (m *RunHopMonitorResponse) SetIpVersion(fv common.IpVersion) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "IpVersion", "RunHopMonitorResponse"))
	}
	m.IpVersion = fv
}

var watchdog_proto_v1alpha2_hop_monitor_custom_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDesc = []byte{
	0x0a, 0x30, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x68, 0x6f, 0x70, 0x5f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x02, 0x0a, 0x14, 0x52, 0x75,
	0x6e, 0x48, 0x6f, 0x70, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x11, 0xb2, 0xda, 0x21, 0x09, 0x12, 0x07, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0xba,
	0x9d, 0x22, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x2d, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x15, 0xb2, 0xda, 0x21, 0x11, 0x12, 0x0f, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x62, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xca, 0xc6, 0x27, 0x06, 0x12, 0x04, 0x2a, 0x02,
	0x08, 0x14, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x30, 0x0a,
	0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x14, 0xca, 0xc6, 0x27, 0x06, 0x12, 0x04, 0x2a, 0x02, 0x08, 0x00, 0xca, 0xc6, 0x27, 0x06, 0x12,
	0x04, 0x1a, 0x02, 0x08, 0x64, 0x52, 0x08, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x73, 0x12,
	0x36, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x3a, 0x10, 0xc2, 0x85, 0x2c, 0x0c, 0x32, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x3a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf3, 0x03, 0x0a, 0x15, 0x52, 0x75,
	0x6e, 0x48, 0x6f, 0x70, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52,
	0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12, 0x57, 0x0a, 0x09, 0x68, 0x6f, 0x70, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x52, 0x75, 0x6e, 0x48, 0x6f, 0x70, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x68, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x54, 0x0a, 0x08, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x39, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x75, 0x6e, 0x48, 0x6f, 0x70,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x48, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x6f,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x49, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x69, 0x70, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x5b, 0x0a, 0x0d, 0x48, 0x6f, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x48, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x5a, 0x0a, 0x0c, 0x48, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x48, 0x6f, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0xd1, 0x01, 0xe8, 0xde, 0x21, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x42, 0x15, 0x48, 0x6f, 0x70, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x4b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x68, 0x6f, 0x70, 0x5f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x3b, 0x68, 0x6f, 0x70, 0x5f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x44, 0x0a,
	0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDescData = watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDesc
)

func watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDescData
}

var watchdog_proto_v1alpha2_hop_monitor_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var watchdog_proto_v1alpha2_hop_monitor_custom_proto_goTypes = []interface{}{
	(*RunHopMonitorRequest)(nil),  // 0: ntt.watchdog.v1alpha2.RunHopMonitorRequest
	(*RunHopMonitorResponse)(nil), // 1: ntt.watchdog.v1alpha2.RunHopMonitorResponse
	nil,                           // 2: ntt.watchdog.v1alpha2.RunHopMonitorResponse.HopStatsEntry
	nil,                           // 3: ntt.watchdog.v1alpha2.RunHopMonitorResponse.HopInfoEntry
	(common.ProbingMode)(0),       // 4: ntt.watchdog.v1alpha2.ProbingMode
	(*common.Path)(nil),           // 5: ntt.watchdog.v1alpha2.Path
	(*common.HopStat)(nil),        // 6: ntt.watchdog.v1alpha2.HopStat
	(*common.HopInfo)(nil),        // 7: ntt.watchdog.v1alpha2.HopInfo
	(common.IpVersion)(0),         // 8: ntt.watchdog.v1alpha2.IpVersion
}
var watchdog_proto_v1alpha2_hop_monitor_custom_proto_depIdxs = []int32{
	4, // 0: ntt.watchdog.v1alpha2.RunHopMonitorRequest.mode:type_name -> ntt.watchdog.v1alpha2.ProbingMode
	5, // 1: ntt.watchdog.v1alpha2.RunHopMonitorResponse.paths:type_name -> ntt.watchdog.v1alpha2.Path
	2, // 2: ntt.watchdog.v1alpha2.RunHopMonitorResponse.hop_stats:type_name -> ntt.watchdog.v1alpha2.RunHopMonitorResponse.HopStatsEntry
	3, // 3: ntt.watchdog.v1alpha2.RunHopMonitorResponse.hop_info:type_name -> ntt.watchdog.v1alpha2.RunHopMonitorResponse.HopInfoEntry
	8, // 4: ntt.watchdog.v1alpha2.RunHopMonitorResponse.ip_version:type_name -> ntt.watchdog.v1alpha2.IpVersion
	6, // 5: ntt.watchdog.v1alpha2.RunHopMonitorResponse.HopStatsEntry.value:type_name -> ntt.watchdog.v1alpha2.HopStat
	7, // 6: ntt.watchdog.v1alpha2.RunHopMonitorResponse.HopInfoEntry.value:type_name -> ntt.watchdog.v1alpha2.HopInfo
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_hop_monitor_custom_proto_init() }
func watchdog_proto_v1alpha2_hop_monitor_custom_proto_init() {
	if watchdog_proto_v1alpha2_hop_monitor_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_hop_monitor_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunHopMonitorRequest); i {
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
		watchdog_proto_v1alpha2_hop_monitor_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunHopMonitorResponse); i {
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
			RawDescriptor: watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_hop_monitor_custom_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_hop_monitor_custom_proto_depIdxs,
		MessageInfos:      watchdog_proto_v1alpha2_hop_monitor_custom_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_hop_monitor_custom_proto = out.File
	watchdog_proto_v1alpha2_hop_monitor_custom_proto_rawDesc = nil
	watchdog_proto_v1alpha2_hop_monitor_custom_proto_goTypes = nil
	watchdog_proto_v1alpha2_hop_monitor_custom_proto_depIdxs = nil
}
