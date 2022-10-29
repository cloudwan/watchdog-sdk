// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/probe_group.proto
// DO NOT EDIT!!!

package probe_group

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
	_ = &project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ProbeGroup Resource
type ProbeGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of ProbeGroup
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Display Name of ProbeGroup
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// state of group
	State *ProbeGroup_State `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty" firestore:"state"`
	// metadata is additional information of the probe group.
	Metadata *ntt_meta.Meta `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *ProbeGroup) Reset() {
	*m = ProbeGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_probe_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProbeGroup) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProbeGroup) ProtoMessage() {}

func (m *ProbeGroup) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_probe_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProbeGroup) GotenMessage() {}

// Deprecated, Use ProbeGroup.ProtoReflect.Descriptor instead.
func (*ProbeGroup) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_probe_group_proto_rawDescGZIP(), []int{0}
}

func (m *ProbeGroup) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProbeGroup) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProbeGroup) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProbeGroup) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProbeGroup) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ProbeGroup) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ProbeGroup) GetState() *ProbeGroup_State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *ProbeGroup) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ProbeGroup) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ProbeGroup"))
	}
	m.Name = fv
}

func (m *ProbeGroup) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "ProbeGroup"))
	}
	m.DisplayName = fv
}

func (m *ProbeGroup) SetState(fv *ProbeGroup_State) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "State", "ProbeGroup"))
	}
	m.State = fv
}

func (m *ProbeGroup) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "ProbeGroup"))
	}
	m.Metadata = fv
}

// Group State
type ProbeGroup_State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// total count of [Probes][ntt.watchdog.v1alpha2.Probe] associated with this
	// group
	TargetCount int64 `protobuf:"varint,1,opt,name=target_count,json=targetCount,proto3" json:"target_count,omitempty" firestore:"targetCount"`
	// Count per region
	RegionalTargetCounts map[string]int64 `protobuf:"bytes,2,rep,name=regional_target_counts,json=regionalTargetCounts,proto3" json:"regional_target_counts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3" firestore:"regionalTargetCounts"`
}

func (m *ProbeGroup_State) Reset() {
	*m = ProbeGroup_State{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_probe_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProbeGroup_State) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProbeGroup_State) ProtoMessage() {}

func (m *ProbeGroup_State) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_probe_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProbeGroup_State) GotenMessage() {}

// Deprecated, Use ProbeGroup_State.ProtoReflect.Descriptor instead.
func (*ProbeGroup_State) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_probe_group_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ProbeGroup_State) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProbeGroup_State) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProbeGroup_State) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProbeGroup_State) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProbeGroup_State) GetTargetCount() int64 {
	if m != nil {
		return m.TargetCount
	}
	return int64(0)
}

func (m *ProbeGroup_State) GetRegionalTargetCounts() map[string]int64 {
	if m != nil {
		return m.RegionalTargetCounts
	}
	return nil
}

func (m *ProbeGroup_State) SetTargetCount(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TargetCount", "ProbeGroup_State"))
	}
	m.TargetCount = fv
}

func (m *ProbeGroup_State) SetRegionalTargetCounts(fv map[string]int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RegionalTargetCounts", "ProbeGroup_State"))
	}
	m.RegionalTargetCounts = fv
}

var watchdog_proto_v1alpha2_probe_group_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_probe_group_proto_rawDesc = []byte{
	0x0a, 0x29, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x05, 0x0a, 0x0a,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xb2, 0xda, 0x21, 0x0e, 0x0a, 0x0c,
	0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0xec, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x77,
	0x0a, 0x16, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x41,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x14, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x1a, 0x47, 0x0a, 0x19, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x3a, 0xc0, 0x02, 0xea, 0x41, 0x4e, 0x0a, 0x1e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x7d, 0x92, 0xd9, 0x21, 0x2b, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x1a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x50, 0x05, 0xaa, 0xd9, 0x21, 0x80, 0x01, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x0a, 0x0d, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x74, 0x61, 0x67, 0x73, 0x0a, 0x0f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x10, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x14,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x1a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x14, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0xda, 0x94, 0x23, 0x08, 0x12, 0x06, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x25, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xe2, 0xde, 0x21,
	0x02, 0x08, 0x01, 0x42, 0xeb, 0x02, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x4c, 0x0a,
	0x11, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x0a, 0x1c, 0x63, 0x6f, 0x6d,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0f, 0x50, 0x72, 0x6f, 0x62, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3b, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0xd2, 0x84, 0xd1, 0x02, 0x44, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x80,
	0xd1, 0x02, 0x4e, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_probe_group_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_probe_group_proto_rawDescData = watchdog_proto_v1alpha2_probe_group_proto_rawDesc
)

func watchdog_proto_v1alpha2_probe_group_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_probe_group_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_probe_group_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_probe_group_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_probe_group_proto_rawDescData
}

var watchdog_proto_v1alpha2_probe_group_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var watchdog_proto_v1alpha2_probe_group_proto_goTypes = []interface{}{
	(*ProbeGroup)(nil),       // 0: ntt.watchdog.v1alpha2.ProbeGroup
	(*ProbeGroup_State)(nil), // 1: ntt.watchdog.v1alpha2.ProbeGroup.State
	nil,                      // 2: ntt.watchdog.v1alpha2.ProbeGroup.State.RegionalTargetCountsEntry
	(*ntt_meta.Meta)(nil),    // 3: ntt.types.Meta
}
var watchdog_proto_v1alpha2_probe_group_proto_depIdxs = []int32{
	1, // 0: ntt.watchdog.v1alpha2.ProbeGroup.state:type_name -> ntt.watchdog.v1alpha2.ProbeGroup.State
	3, // 1: ntt.watchdog.v1alpha2.ProbeGroup.metadata:type_name -> ntt.types.Meta
	2, // 2: ntt.watchdog.v1alpha2.ProbeGroup.State.regional_target_counts:type_name -> ntt.watchdog.v1alpha2.ProbeGroup.State.RegionalTargetCountsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_probe_group_proto_init() }
func watchdog_proto_v1alpha2_probe_group_proto_init() {
	if watchdog_proto_v1alpha2_probe_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_probe_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeGroup); i {
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
		watchdog_proto_v1alpha2_probe_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeGroup_State); i {
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
			RawDescriptor: watchdog_proto_v1alpha2_probe_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_probe_group_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_probe_group_proto_depIdxs,
		MessageInfos:      watchdog_proto_v1alpha2_probe_group_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_probe_group_proto = out.File
	watchdog_proto_v1alpha2_probe_group_proto_rawDesc = nil
	watchdog_proto_v1alpha2_probe_group_proto_goTypes = nil
	watchdog_proto_v1alpha2_probe_group_proto_depIdxs = nil
}
