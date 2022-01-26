// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha/probing_target_group.proto
// DO NOT EDIT!!!

package probing_target_group

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
	_ = &project.Project{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ProbingTargetGroup Resource
type ProbingTargetGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of ProbingTargetGroup
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Human-readable name
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" firestore:"displayName"`
	// state of group
	State *ProbingTargetGroup_State `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty" firestore:"state"`
	// metadata is additional information of the probing target group.
	Metadata *ntt_meta.Meta `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *ProbingTargetGroup) Reset() {
	*m = ProbingTargetGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha_probing_target_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProbingTargetGroup) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProbingTargetGroup) ProtoMessage() {}

func (m *ProbingTargetGroup) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha_probing_target_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProbingTargetGroup) GotenMessage() {}

// Deprecated, Use ProbingTargetGroup.ProtoReflect.Descriptor instead.
func (*ProbingTargetGroup) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha_probing_target_group_proto_rawDescGZIP(), []int{0}
}

func (m *ProbingTargetGroup) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProbingTargetGroup) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProbingTargetGroup) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProbingTargetGroup) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProbingTargetGroup) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ProbingTargetGroup) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ProbingTargetGroup) GetState() *ProbingTargetGroup_State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *ProbingTargetGroup) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ProbingTargetGroup) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ProbingTargetGroup"))
	}
	m.Name = fv
}

func (m *ProbingTargetGroup) SetDisplayName(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DisplayName", "ProbingTargetGroup"))
	}
	m.DisplayName = fv
}

func (m *ProbingTargetGroup) SetState(fv *ProbingTargetGroup_State) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "State", "ProbingTargetGroup"))
	}
	m.State = fv
}

func (m *ProbingTargetGroup) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "ProbingTargetGroup"))
	}
	m.Metadata = fv
}

// Group State
type ProbingTargetGroup_State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// total count of [ProbingTargets][ntt.watchdog.v1alpha.ProbingTarget]
	// associated with this group
	TargetCount int64 `protobuf:"varint,1,opt,name=target_count,json=targetCount,proto3" json:"target_count,omitempty" firestore:"targetCount"`
}

func (m *ProbingTargetGroup_State) Reset() {
	*m = ProbingTargetGroup_State{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha_probing_target_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProbingTargetGroup_State) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProbingTargetGroup_State) ProtoMessage() {}

func (m *ProbingTargetGroup_State) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha_probing_target_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProbingTargetGroup_State) GotenMessage() {}

// Deprecated, Use ProbingTargetGroup_State.ProtoReflect.Descriptor instead.
func (*ProbingTargetGroup_State) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha_probing_target_group_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ProbingTargetGroup_State) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProbingTargetGroup_State) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProbingTargetGroup_State) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProbingTargetGroup_State) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProbingTargetGroup_State) GetTargetCount() int64 {
	if m != nil {
		return m.TargetCount
	}
	return int64(0)
}

func (m *ProbingTargetGroup_State) SetTargetCount(fv int64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TargetCount", "ProbingTargetGroup_State"))
	}
	m.TargetCount = fv
}

var watchdog_proto_v1alpha_probing_target_group_proto preflect.FileDescriptor

var watchdog_proto_v1alpha_probing_target_group_proto_rawDesc = []byte{
	0x0a, 0x31, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67,
	0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xfd, 0x03, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xb2, 0xda, 0x21, 0x16, 0x0a, 0x14, 0x0a, 0x12, 0x50,
	0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x2b, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x2a, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0xf4, 0x01, 0xea, 0x41, 0x67, 0x0a,
	0x26, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x62,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x7d, 0x92, 0xd9, 0x21, 0x39, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x62,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12,
	0x13, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x1a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0xfa, 0xde, 0x21, 0x14, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0xda, 0x94, 0x23, 0x08, 0x12,
	0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xc2, 0x85, 0x2c, 0x25, 0x22, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x42, 0x90, 0x04, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0, 0x02, 0x5d, 0x0a, 0x1a, 0x70, 0x72,
	0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x92, 0x8c, 0xd1, 0x02, 0x67, 0x0a, 0x1f,
	0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x42, 0x17, 0x50, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77,
	0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x72,
	0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x3b, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0xd2, 0x84, 0xd1, 0x02, 0x43, 0x0a, 0x0d, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xa2,
	0x80, 0xd1, 0x02, 0x5f, 0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x62, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha_probing_target_group_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha_probing_target_group_proto_rawDescData = watchdog_proto_v1alpha_probing_target_group_proto_rawDesc
)

func watchdog_proto_v1alpha_probing_target_group_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha_probing_target_group_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha_probing_target_group_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha_probing_target_group_proto_rawDescData)
	})
	return watchdog_proto_v1alpha_probing_target_group_proto_rawDescData
}

var watchdog_proto_v1alpha_probing_target_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var watchdog_proto_v1alpha_probing_target_group_proto_goTypes = []interface{}{
	(*ProbingTargetGroup)(nil),       // 0: ntt.watchdog.v1alpha.ProbingTargetGroup
	(*ProbingTargetGroup_State)(nil), // 1: ntt.watchdog.v1alpha.ProbingTargetGroup.State
	(*ntt_meta.Meta)(nil),            // 2: ntt.types.Meta
}
var watchdog_proto_v1alpha_probing_target_group_proto_depIdxs = []int32{
	1, // 0: ntt.watchdog.v1alpha.ProbingTargetGroup.state:type_name -> ntt.watchdog.v1alpha.ProbingTargetGroup.State
	2, // 1: ntt.watchdog.v1alpha.ProbingTargetGroup.metadata:type_name -> ntt.types.Meta
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha_probing_target_group_proto_init() }
func watchdog_proto_v1alpha_probing_target_group_proto_init() {
	if watchdog_proto_v1alpha_probing_target_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha_probing_target_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbingTargetGroup); i {
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
		watchdog_proto_v1alpha_probing_target_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbingTargetGroup_State); i {
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
			RawDescriptor: watchdog_proto_v1alpha_probing_target_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha_probing_target_group_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha_probing_target_group_proto_depIdxs,
		MessageInfos:      watchdog_proto_v1alpha_probing_target_group_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha_probing_target_group_proto = out.File
	watchdog_proto_v1alpha_probing_target_group_proto_rawDesc = nil
	watchdog_proto_v1alpha_probing_target_group_proto_goTypes = nil
	watchdog_proto_v1alpha_probing_target_group_proto_depIdxs = nil
}
