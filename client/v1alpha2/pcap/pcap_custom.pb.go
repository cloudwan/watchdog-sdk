// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/pcap_custom.proto
// DO NOT EDIT!!!

package pcap_client

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
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &timestamp.Timestamp{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method [ReportPcap][ntt.watchdog.v1alpha2.ReportPcap]
type ReportPcapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  reference of ntt.watchdog.v1alpha2.Probe
	Name *probe.Reference `protobuf:"bytes,1,opt,customtype=Reference,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// The filter that was applied for this capture
	Filter    string               `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty" firestore:"filter"`
	StartTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" firestore:"startTime"`
	EndTime   *timestamp.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" firestore:"endTime"`
	PcapBytes []byte               `protobuf:"bytes,5,opt,name=pcap_bytes,json=pcapBytes,proto3" json:"pcap_bytes,omitempty" firestore:"pcapBytes"`
}

func (m *ReportPcapRequest) Reset() {
	*m = ReportPcapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ReportPcapRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ReportPcapRequest) ProtoMessage() {}

func (m *ReportPcapRequest) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ReportPcapRequest) GotenMessage() {}

// Deprecated, Use ReportPcapRequest.ProtoReflect.Descriptor instead.
func (*ReportPcapRequest) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_pcap_custom_proto_rawDescGZIP(), []int{0}
}

func (m *ReportPcapRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ReportPcapRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ReportPcapRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ReportPcapRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ReportPcapRequest) GetName() *probe.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ReportPcapRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ReportPcapRequest) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ReportPcapRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ReportPcapRequest) GetPcapBytes() []byte {
	if m != nil {
		return m.PcapBytes
	}
	return nil
}

func (m *ReportPcapRequest) SetName(fv *probe.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ReportPcapRequest"))
	}
	m.Name = fv
}

func (m *ReportPcapRequest) SetFilter(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Filter", "ReportPcapRequest"))
	}
	m.Filter = fv
}

func (m *ReportPcapRequest) SetStartTime(fv *timestamp.Timestamp) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "StartTime", "ReportPcapRequest"))
	}
	m.StartTime = fv
}

func (m *ReportPcapRequest) SetEndTime(fv *timestamp.Timestamp) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EndTime", "ReportPcapRequest"))
	}
	m.EndTime = fv
}

func (m *ReportPcapRequest) SetPcapBytes(fv []byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PcapBytes", "ReportPcapRequest"))
	}
	m.PcapBytes = fv
}

// Request message for method [GetPcap][ntt.watchdog.v1alpha2.GetPcap]
type GetPcapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  reference of ntt.watchdog.v1alpha2.Probe
	Probe *probe.Reference `protobuf:"bytes,1,opt,customtype=Reference,name=probe,proto3" json:"probe,omitempty" firestore:"probe"`
	// Interval for which the pcap is required
	Interval *common.TimeInterval `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty" firestore:"interval"`
	// A positive number that is the maximum number of results to return. If
	// `page_size` is empty or more than 100,000 results, the effective
	// `page_size` is 100,000 results.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" firestore:"pageSize"`
	// If this field is not empty then it must contain the `nextPageToken` value
	// returned by a previous call to this method.  Using this field causes the
	// method to return additional results from the previous method call.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty" firestore:"pageToken"`
}

func (m *GetPcapRequest) Reset() {
	*m = GetPcapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GetPcapRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GetPcapRequest) ProtoMessage() {}

func (m *GetPcapRequest) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GetPcapRequest) GotenMessage() {}

// Deprecated, Use GetPcapRequest.ProtoReflect.Descriptor instead.
func (*GetPcapRequest) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_pcap_custom_proto_rawDescGZIP(), []int{1}
}

func (m *GetPcapRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GetPcapRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GetPcapRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GetPcapRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GetPcapRequest) GetProbe() *probe.Reference {
	if m != nil {
		return m.Probe
	}
	return nil
}

func (m *GetPcapRequest) GetInterval() *common.TimeInterval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *GetPcapRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return int32(0)
}

func (m *GetPcapRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *GetPcapRequest) SetProbe(fv *probe.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Probe", "GetPcapRequest"))
	}
	m.Probe = fv
}

func (m *GetPcapRequest) SetInterval(fv *common.TimeInterval) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Interval", "GetPcapRequest"))
	}
	m.Interval = fv
}

func (m *GetPcapRequest) SetPageSize(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageSize", "GetPcapRequest"))
	}
	m.PageSize = fv
}

func (m *GetPcapRequest) SetPageToken(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PageToken", "GetPcapRequest"))
	}
	m.PageToken = fv
}

// Response message for method [GetPcap][ntt.watchdog.v1alpha2.GetPcap]
type GetPcapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	PcapBytes     []byte `protobuf:"bytes,1,opt,name=pcap_bytes,json=pcapBytes,proto3" json:"pcap_bytes,omitempty" firestore:"pcapBytes"`
}

func (m *GetPcapResponse) Reset() {
	*m = GetPcapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GetPcapResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GetPcapResponse) ProtoMessage() {}

func (m *GetPcapResponse) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GetPcapResponse) GotenMessage() {}

// Deprecated, Use GetPcapResponse.ProtoReflect.Descriptor instead.
func (*GetPcapResponse) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_pcap_custom_proto_rawDescGZIP(), []int{2}
}

func (m *GetPcapResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GetPcapResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GetPcapResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GetPcapResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GetPcapResponse) GetPcapBytes() []byte {
	if m != nil {
		return m.PcapBytes
	}
	return nil
}

func (m *GetPcapResponse) SetPcapBytes(fv []byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PcapBytes", "GetPcapResponse"))
	}
	m.PcapBytes = fv
}

type GetPcapFileFromAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  reference of ntt.watchdog.v1alpha2.Probe
	Name *probe.Reference `protobuf:"bytes,1,opt,customtype=Reference,name=name,proto3" json:"name,omitempty" firestore:"name"`
}

func (m *GetPcapFileFromAgentRequest) Reset() {
	*m = GetPcapFileFromAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GetPcapFileFromAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GetPcapFileFromAgentRequest) ProtoMessage() {}

func (m *GetPcapFileFromAgentRequest) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GetPcapFileFromAgentRequest) GotenMessage() {}

// Deprecated, Use GetPcapFileFromAgentRequest.ProtoReflect.Descriptor instead.
func (*GetPcapFileFromAgentRequest) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_pcap_custom_proto_rawDescGZIP(), []int{3}
}

func (m *GetPcapFileFromAgentRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GetPcapFileFromAgentRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GetPcapFileFromAgentRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GetPcapFileFromAgentRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GetPcapFileFromAgentRequest) GetName() *probe.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GetPcapFileFromAgentRequest) SetName(fv *probe.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "GetPcapFileFromAgentRequest"))
	}
	m.Name = fv
}

var watchdog_proto_v1alpha2_pcap_custom_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_pcap_custom_proto_rawDesc = []byte{
	0x0a, 0x29, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x63, 0x61, 0x70, 0x5f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65,
	0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x11, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x50, 0x63, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2,
	0xda, 0x21, 0x09, 0x12, 0x07, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0xba, 0x9d, 0x22, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x39,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x63, 0x61, 0x70, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x63, 0x61, 0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x3a,
	0x10, 0xc2, 0x85, 0x2c, 0x0c, 0x32, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0xe0, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09, 0x12, 0x07, 0x0a, 0x05, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x49, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x42, 0x08, 0xca, 0xc6, 0x27, 0x04, 0x62, 0x02, 0x08, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0e, 0xca, 0xc6, 0x27, 0x0a, 0x12, 0x08, 0x1a,
	0x04, 0x08, 0xa0, 0x8d, 0x06, 0x2a, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x3a, 0x12, 0xc2, 0x85, 0x2c, 0x0e, 0x32, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x3a, 0x05, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x63, 0x61, 0x70, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x63, 0x61,
	0x70, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x63, 0x61,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0d, 0xb2, 0xda, 0x21, 0x09, 0x12, 0x07, 0x0a, 0x05, 0x50, 0x72, 0x6f,
	0x62, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x10, 0xc2, 0x85, 0x2c, 0x0c, 0x32, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xbd, 0x01, 0xe8, 0xde, 0x21,
	0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42,
	0x0f, 0x50, 0x63, 0x61, 0x70, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x00, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x70, 0x63, 0x61, 0x70, 0x3b, 0x70, 0x63, 0x61, 0x70, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0xd2, 0x84, 0xd1, 0x02, 0x44, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_pcap_custom_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_pcap_custom_proto_rawDescData = watchdog_proto_v1alpha2_pcap_custom_proto_rawDesc
)

func watchdog_proto_v1alpha2_pcap_custom_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_pcap_custom_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_pcap_custom_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_pcap_custom_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_pcap_custom_proto_rawDescData
}

var watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var watchdog_proto_v1alpha2_pcap_custom_proto_goTypes = []interface{}{
	(*ReportPcapRequest)(nil),           // 0: ntt.watchdog.v1alpha2.ReportPcapRequest
	(*GetPcapRequest)(nil),              // 1: ntt.watchdog.v1alpha2.GetPcapRequest
	(*GetPcapResponse)(nil),             // 2: ntt.watchdog.v1alpha2.GetPcapResponse
	(*GetPcapFileFromAgentRequest)(nil), // 3: ntt.watchdog.v1alpha2.GetPcapFileFromAgentRequest
	(*timestamp.Timestamp)(nil),         // 4: google.protobuf.Timestamp
	(*common.TimeInterval)(nil),         // 5: ntt.watchdog.v1alpha2.TimeInterval
}
var watchdog_proto_v1alpha2_pcap_custom_proto_depIdxs = []int32{
	4, // 0: ntt.watchdog.v1alpha2.ReportPcapRequest.start_time:type_name -> google.protobuf.Timestamp
	4, // 1: ntt.watchdog.v1alpha2.ReportPcapRequest.end_time:type_name -> google.protobuf.Timestamp
	5, // 2: ntt.watchdog.v1alpha2.GetPcapRequest.interval:type_name -> ntt.watchdog.v1alpha2.TimeInterval
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_pcap_custom_proto_init() }
func watchdog_proto_v1alpha2_pcap_custom_proto_init() {
	if watchdog_proto_v1alpha2_pcap_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportPcapRequest); i {
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
		watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPcapRequest); i {
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
		watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPcapResponse); i {
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
		watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPcapFileFromAgentRequest); i {
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
			RawDescriptor: watchdog_proto_v1alpha2_pcap_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_pcap_custom_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_pcap_custom_proto_depIdxs,
		MessageInfos:      watchdog_proto_v1alpha2_pcap_custom_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_pcap_custom_proto = out.File
	watchdog_proto_v1alpha2_pcap_custom_proto_rawDesc = nil
	watchdog_proto_v1alpha2_pcap_custom_proto_goTypes = nil
	watchdog_proto_v1alpha2_pcap_custom_proto_depIdxs = nil
}
