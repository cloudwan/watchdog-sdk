// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/shared_token_custom.proto
// DO NOT EDIT!!!

package shared_token_client

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
	api "github.com/cloudwan/edgelq-sdk/common/api"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	shared_token "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/shared_token"
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
	_ = &api.Account{}
	_ = &probe.Probe{}
	_ = &shared_token.SharedToken{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method
// [RegisterProbe][ntt.watchdog.v1alpha2.RegisterProbe]
type RegisterProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  reference of ntt.watchdog.v1alpha2.SharedToken
	Name *shared_token.Reference `protobuf:"bytes,1,opt,customtype=Reference,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Key value pairs to be used to substitute placeholders in template values.
	Values map[string]string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"values"`
	// The current probe status.
	Status *probe.Probe_Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty" firestore:"status"`
}

func (m *RegisterProbeRequest) Reset() {
	*m = RegisterProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_shared_token_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RegisterProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RegisterProbeRequest) ProtoMessage() {}

func (m *RegisterProbeRequest) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_shared_token_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RegisterProbeRequest) GotenMessage() {}

// Deprecated, Use RegisterProbeRequest.ProtoReflect.Descriptor instead.
func (*RegisterProbeRequest) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_shared_token_custom_proto_rawDescGZIP(), []int{0}
}

func (m *RegisterProbeRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RegisterProbeRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RegisterProbeRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RegisterProbeRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RegisterProbeRequest) GetName() *shared_token.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RegisterProbeRequest) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *RegisterProbeRequest) GetStatus() *probe.Probe_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *RegisterProbeRequest) SetName(fv *shared_token.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RegisterProbeRequest"))
	}
	m.Name = fv
}

func (m *RegisterProbeRequest) SetValues(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Values", "RegisterProbeRequest"))
	}
	m.Values = fv
}

func (m *RegisterProbeRequest) SetStatus(fv *probe.Probe_Status) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Status", "RegisterProbeRequest"))
	}
	m.Status = fv
}

// Response message for method
// [RegisterProbe][ntt.watchdog.v1alpha2.RegisterProbe]
type RegisterProbeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// The service account generated for the request.
	ServiceAccount *api.ServiceAccount `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty" firestore:"serviceAccount"`
	// The name of probe generated for the request.
	Probe *probe.Reference `protobuf:"bytes,2,opt,customtype=Reference,name=probe,proto3" json:"probe,omitempty" firestore:"probe"`
}

func (m *RegisterProbeResponse) Reset() {
	*m = RegisterProbeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_shared_token_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RegisterProbeResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RegisterProbeResponse) ProtoMessage() {}

func (m *RegisterProbeResponse) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_shared_token_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RegisterProbeResponse) GotenMessage() {}

// Deprecated, Use RegisterProbeResponse.ProtoReflect.Descriptor instead.
func (*RegisterProbeResponse) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_shared_token_custom_proto_rawDescGZIP(), []int{1}
}

func (m *RegisterProbeResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RegisterProbeResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RegisterProbeResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RegisterProbeResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RegisterProbeResponse) GetServiceAccount() *api.ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

func (m *RegisterProbeResponse) GetProbe() *probe.Reference {
	if m != nil {
		return m.Probe
	}
	return nil
}

func (m *RegisterProbeResponse) SetServiceAccount(fv *api.ServiceAccount) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ServiceAccount", "RegisterProbeResponse"))
	}
	m.ServiceAccount = fv
}

func (m *RegisterProbeResponse) SetProbe(fv *probe.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Probe", "RegisterProbeResponse"))
	}
	m.Probe = fv
}

var watchdog_proto_v1alpha2_shared_token_custom_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_shared_token_custom_proto_rawDesc = []byte{
	0x0a, 0x31, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xb2, 0xda, 0x21, 0x0f, 0x12,
	0x0d, 0x0a, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0xba, 0x9d,
	0x22, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x3a, 0x0a, 0xc2, 0x85, 0x2c, 0x06, 0x32, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa8, 0x01,
	0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09, 0x12, 0x07,
	0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x3a, 0x28,
	0xc2, 0x85, 0x2c, 0x24, 0x1a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x3a, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x42, 0xd4, 0x01, 0xe8, 0xde, 0x21, 0x01, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x16, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x3b, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x44, 0x0a, 0x0d, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_shared_token_custom_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_shared_token_custom_proto_rawDescData = watchdog_proto_v1alpha2_shared_token_custom_proto_rawDesc
)

func watchdog_proto_v1alpha2_shared_token_custom_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_shared_token_custom_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_shared_token_custom_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_shared_token_custom_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_shared_token_custom_proto_rawDescData
}

var watchdog_proto_v1alpha2_shared_token_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var watchdog_proto_v1alpha2_shared_token_custom_proto_goTypes = []interface{}{
	(*RegisterProbeRequest)(nil),  // 0: ntt.watchdog.v1alpha2.RegisterProbeRequest
	(*RegisterProbeResponse)(nil), // 1: ntt.watchdog.v1alpha2.RegisterProbeResponse
	nil,                           // 2: ntt.watchdog.v1alpha2.RegisterProbeRequest.ValuesEntry
	(*probe.Probe_Status)(nil),    // 3: ntt.watchdog.v1alpha2.Probe.Status
	(*api.ServiceAccount)(nil),    // 4: ntt.api.ServiceAccount
}
var watchdog_proto_v1alpha2_shared_token_custom_proto_depIdxs = []int32{
	2, // 0: ntt.watchdog.v1alpha2.RegisterProbeRequest.values:type_name -> ntt.watchdog.v1alpha2.RegisterProbeRequest.ValuesEntry
	3, // 1: ntt.watchdog.v1alpha2.RegisterProbeRequest.status:type_name -> ntt.watchdog.v1alpha2.Probe.Status
	4, // 2: ntt.watchdog.v1alpha2.RegisterProbeResponse.service_account:type_name -> ntt.api.ServiceAccount
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_shared_token_custom_proto_init() }
func watchdog_proto_v1alpha2_shared_token_custom_proto_init() {
	if watchdog_proto_v1alpha2_shared_token_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_shared_token_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterProbeRequest); i {
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
		watchdog_proto_v1alpha2_shared_token_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterProbeResponse); i {
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
			RawDescriptor: watchdog_proto_v1alpha2_shared_token_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_shared_token_custom_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_shared_token_custom_proto_depIdxs,
		MessageInfos:      watchdog_proto_v1alpha2_shared_token_custom_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_shared_token_custom_proto = out.File
	watchdog_proto_v1alpha2_shared_token_custom_proto_rawDesc = nil
	watchdog_proto_v1alpha2_shared_token_custom_proto_goTypes = nil
	watchdog_proto_v1alpha2_shared_token_custom_proto_depIdxs = nil
}
