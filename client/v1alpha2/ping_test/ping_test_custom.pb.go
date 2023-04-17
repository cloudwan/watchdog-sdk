// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/ping_test_custom.proto
// DO NOT EDIT!!!

package ping_test_client

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
	duration "github.com/golang/protobuf/ptypes/duration"
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
	_ = &duration.Duration{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for method [RunPingTest][ntt.watchdog.v1alpha2.RunPingTest]
// NOTE: copy watchdogagent/api/proto/v1alpha/ping_test_custom.proto and
// increment field numbers Source address to listen packet Skip validating
type RunPingTestRequest struct {
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
	// Destination to send packet
	// IP address and hostname are acceptable
	// As the IPv6 address, a zone index is also acceptable
	Destination string `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty" firestore:"destination"`
	// Byte size of the ICMP payload
	// If unspecified, 100 is picked by default
	//
	// The minimum length is 2 bytes
	// The head 2 bytes body is used to embed a request ID
	// Because ICMP Echo identifier field is not editable
	// in Linux unprivileged ICMP endpoint
	SizeBytes int32 `protobuf:"varint,4,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty" firestore:"sizeBytes"`
	// Number of ICMP Echo requests in the test
	// If unspecified, 3 is picked by default
	Count int32 `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty" firestore:"count"`
	// Interval duration to wait sending each packet
	// If unspecified, 1.0s is picked
	Interval *duration.Duration `protobuf:"bytes,6,opt,name=interval,proto3" json:"interval,omitempty" firestore:"interval"`
	// Timeout duration to wait receiving each packet's reply
	// If unspecified, 1.0s is picked
	EchoTimeout *duration.Duration `protobuf:"bytes,7,opt,name=echo_timeout,json=echoTimeout,proto3" json:"echo_timeout,omitempty" firestore:"echoTimeout"`
	// DF flag in IPv4 header
	// If unspecified or false spcified, skip manipulating packet
	DontFragment bool `protobuf:"varint,8,opt,name=dont_fragment,json=dontFragment,proto3" json:"dont_fragment,omitempty" firestore:"dontFragment"`
	// TTL in IPv4 or hop limit in IPv6
	// If unspecified, skip manipulating packet
	// In the case, the value depends on the running OS
	Ttl int32 `protobuf:"varint,9,opt,name=ttl,proto3" json:"ttl,omitempty" firestore:"ttl"`
	// TOS in IPv4 or traffic class in IPv6
	// If unspecified, skip manipulating packet
	Tos          int32                             `protobuf:"varint,10,opt,name=tos,proto3" json:"tos,omitempty" firestore:"tos"`
	OutputFormat common.OnDemandTestResponseFormat `protobuf:"varint,11,opt,name=output_format,json=outputFormat,proto3,enum=ntt.watchdog.v1alpha2.OnDemandTestResponseFormat" json:"output_format,omitempty" firestore:"outputFormat"`
}

func (m *RunPingTestRequest) Reset() {
	*m = RunPingTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RunPingTestRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RunPingTestRequest) ProtoMessage() {}

func (m *RunPingTestRequest) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RunPingTestRequest) GotenMessage() {}

// Deprecated, Use RunPingTestRequest.ProtoReflect.Descriptor instead.
func (*RunPingTestRequest) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescGZIP(), []int{0}
}

func (m *RunPingTestRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RunPingTestRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RunPingTestRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RunPingTestRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RunPingTestRequest) GetName() *probe.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RunPingTestRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *RunPingTestRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *RunPingTestRequest) GetSizeBytes() int32 {
	if m != nil {
		return m.SizeBytes
	}
	return int32(0)
}

func (m *RunPingTestRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return int32(0)
}

func (m *RunPingTestRequest) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *RunPingTestRequest) GetEchoTimeout() *duration.Duration {
	if m != nil {
		return m.EchoTimeout
	}
	return nil
}

func (m *RunPingTestRequest) GetDontFragment() bool {
	if m != nil {
		return m.DontFragment
	}
	return false
}

func (m *RunPingTestRequest) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return int32(0)
}

func (m *RunPingTestRequest) GetTos() int32 {
	if m != nil {
		return m.Tos
	}
	return int32(0)
}

func (m *RunPingTestRequest) GetOutputFormat() common.OnDemandTestResponseFormat {
	if m != nil {
		return m.OutputFormat
	}
	return common.OnDemandTestResponseFormat_TEXT
}

func (m *RunPingTestRequest) SetName(fv *probe.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RunPingTestRequest"))
	}
	m.Name = fv
}

func (m *RunPingTestRequest) SetSource(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Source", "RunPingTestRequest"))
	}
	m.Source = fv
}

func (m *RunPingTestRequest) SetDestination(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Destination", "RunPingTestRequest"))
	}
	m.Destination = fv
}

func (m *RunPingTestRequest) SetSizeBytes(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SizeBytes", "RunPingTestRequest"))
	}
	m.SizeBytes = fv
}

func (m *RunPingTestRequest) SetCount(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Count", "RunPingTestRequest"))
	}
	m.Count = fv
}

func (m *RunPingTestRequest) SetInterval(fv *duration.Duration) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Interval", "RunPingTestRequest"))
	}
	m.Interval = fv
}

func (m *RunPingTestRequest) SetEchoTimeout(fv *duration.Duration) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "EchoTimeout", "RunPingTestRequest"))
	}
	m.EchoTimeout = fv
}

func (m *RunPingTestRequest) SetDontFragment(fv bool) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "DontFragment", "RunPingTestRequest"))
	}
	m.DontFragment = fv
}

func (m *RunPingTestRequest) SetTtl(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Ttl", "RunPingTestRequest"))
	}
	m.Ttl = fv
}

func (m *RunPingTestRequest) SetTos(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Tos", "RunPingTestRequest"))
	}
	m.Tos = fv
}

func (m *RunPingTestRequest) SetOutputFormat(fv common.OnDemandTestResponseFormat) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OutputFormat", "RunPingTestRequest"))
	}
	m.OutputFormat = fv
}

// Response message for method [RunPingTest][ntt.watchdog.v1alpha2.RunPingTest]
type RunPingTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	JsonResponse  *RunPingTestResponse_JsonResponse `protobuf:"bytes,1,opt,name=json_response,json=jsonResponse,proto3" json:"json_response,omitempty" firestore:"jsonResponse"`
	TextResponse  string                            `protobuf:"bytes,2,opt,name=text_response,json=textResponse,proto3" json:"text_response,omitempty" firestore:"textResponse"`
}

func (m *RunPingTestResponse) Reset() {
	*m = RunPingTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RunPingTestResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RunPingTestResponse) ProtoMessage() {}

func (m *RunPingTestResponse) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RunPingTestResponse) GotenMessage() {}

// Deprecated, Use RunPingTestResponse.ProtoReflect.Descriptor instead.
func (*RunPingTestResponse) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescGZIP(), []int{1}
}

func (m *RunPingTestResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RunPingTestResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RunPingTestResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RunPingTestResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RunPingTestResponse) GetJsonResponse() *RunPingTestResponse_JsonResponse {
	if m != nil {
		return m.JsonResponse
	}
	return nil
}

func (m *RunPingTestResponse) GetTextResponse() string {
	if m != nil {
		return m.TextResponse
	}
	return ""
}

func (m *RunPingTestResponse) SetJsonResponse(fv *RunPingTestResponse_JsonResponse) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "JsonResponse", "RunPingTestResponse"))
	}
	m.JsonResponse = fv
}

func (m *RunPingTestResponse) SetTextResponse(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TextResponse", "RunPingTestResponse"))
	}
	m.TextResponse = fv
}

type RunPingTestResponse_JsonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// IP address
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" firestore:"from"`
	// Byte size of the ICMP payload
	SizeBytes int32 `protobuf:"varint,2,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty" firestore:"sizeBytes"`
	// Sequence number field
	SequenceNumber int32              `protobuf:"varint,3,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty" firestore:"sequenceNumber"`
	Ttl            int32              `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty" firestore:"ttl"`
	Rtt            *duration.Duration `protobuf:"bytes,5,opt,name=rtt,proto3" json:"rtt,omitempty" firestore:"rtt"`
	// Error message
	Error   string                                         `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty" firestore:"error"`
	Summary *RunPingTestResponse_JsonResponse_SummaryStats `protobuf:"bytes,7,opt,name=summary,proto3" json:"summary,omitempty" firestore:"summary"`
}

func (m *RunPingTestResponse_JsonResponse) Reset() {
	*m = RunPingTestResponse_JsonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RunPingTestResponse_JsonResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RunPingTestResponse_JsonResponse) ProtoMessage() {}

func (m *RunPingTestResponse_JsonResponse) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RunPingTestResponse_JsonResponse) GotenMessage() {}

// Deprecated, Use RunPingTestResponse_JsonResponse.ProtoReflect.Descriptor instead.
func (*RunPingTestResponse_JsonResponse) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescGZIP(), []int{1, 0}
}

func (m *RunPingTestResponse_JsonResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RunPingTestResponse_JsonResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RunPingTestResponse_JsonResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RunPingTestResponse_JsonResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RunPingTestResponse_JsonResponse) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *RunPingTestResponse_JsonResponse) GetSizeBytes() int32 {
	if m != nil {
		return m.SizeBytes
	}
	return int32(0)
}

func (m *RunPingTestResponse_JsonResponse) GetSequenceNumber() int32 {
	if m != nil {
		return m.SequenceNumber
	}
	return int32(0)
}

func (m *RunPingTestResponse_JsonResponse) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return int32(0)
}

func (m *RunPingTestResponse_JsonResponse) GetRtt() *duration.Duration {
	if m != nil {
		return m.Rtt
	}
	return nil
}

func (m *RunPingTestResponse_JsonResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RunPingTestResponse_JsonResponse) GetSummary() *RunPingTestResponse_JsonResponse_SummaryStats {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *RunPingTestResponse_JsonResponse) SetFrom(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "From", "RunPingTestResponse_JsonResponse"))
	}
	m.From = fv
}

func (m *RunPingTestResponse_JsonResponse) SetSizeBytes(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SizeBytes", "RunPingTestResponse_JsonResponse"))
	}
	m.SizeBytes = fv
}

func (m *RunPingTestResponse_JsonResponse) SetSequenceNumber(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SequenceNumber", "RunPingTestResponse_JsonResponse"))
	}
	m.SequenceNumber = fv
}

func (m *RunPingTestResponse_JsonResponse) SetTtl(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Ttl", "RunPingTestResponse_JsonResponse"))
	}
	m.Ttl = fv
}

func (m *RunPingTestResponse_JsonResponse) SetRtt(fv *duration.Duration) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Rtt", "RunPingTestResponse_JsonResponse"))
	}
	m.Rtt = fv
}

func (m *RunPingTestResponse_JsonResponse) SetError(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Error", "RunPingTestResponse_JsonResponse"))
	}
	m.Error = fv
}

func (m *RunPingTestResponse_JsonResponse) SetSummary(fv *RunPingTestResponse_JsonResponse_SummaryStats) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Summary", "RunPingTestResponse_JsonResponse"))
	}
	m.Summary = fv
}

// Summary statistics accompanies each response
// It summarizes the results at that time
type RunPingTestResponse_JsonResponse_SummaryStats struct {
	state              protoimpl.MessageState
	sizeCache          protoimpl.SizeCache
	unknownFields      protoimpl.UnknownFields
	MinRtt             *duration.Duration `protobuf:"bytes,1,opt,name=min_rtt,json=minRtt,proto3" json:"min_rtt,omitempty" firestore:"minRtt"`
	AvgRtt             *duration.Duration `protobuf:"bytes,2,opt,name=avg_rtt,json=avgRtt,proto3" json:"avg_rtt,omitempty" firestore:"avgRtt"`
	MaxRtt             *duration.Duration `protobuf:"bytes,3,opt,name=max_rtt,json=maxRtt,proto3" json:"max_rtt,omitempty" firestore:"maxRtt"`
	StddevRtt          *duration.Duration `protobuf:"bytes,4,opt,name=stddev_rtt,json=stddevRtt,proto3" json:"stddev_rtt,omitempty" firestore:"stddevRtt"`
	TransmittedCounter int32              `protobuf:"varint,5,opt,name=transmitted_counter,json=transmittedCounter,proto3" json:"transmitted_counter,omitempty" firestore:"transmittedCounter"`
	ReceivedCounter    int32              `protobuf:"varint,6,opt,name=received_counter,json=receivedCounter,proto3" json:"received_counter,omitempty" firestore:"receivedCounter"`
	// Percent
	LossRatio float64 `protobuf:"fixed64,7,opt,name=loss_ratio,json=lossRatio,proto3" json:"loss_ratio,omitempty" firestore:"lossRatio"`
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) Reset() {
	*m = RunPingTestResponse_JsonResponse_SummaryStats{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RunPingTestResponse_JsonResponse_SummaryStats) ProtoMessage() {}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RunPingTestResponse_JsonResponse_SummaryStats) GotenMessage() {}

// Deprecated, Use RunPingTestResponse_JsonResponse_SummaryStats.ProtoReflect.Descriptor instead.
func (*RunPingTestResponse_JsonResponse_SummaryStats) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) GetMinRtt() *duration.Duration {
	if m != nil {
		return m.MinRtt
	}
	return nil
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) GetAvgRtt() *duration.Duration {
	if m != nil {
		return m.AvgRtt
	}
	return nil
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) GetMaxRtt() *duration.Duration {
	if m != nil {
		return m.MaxRtt
	}
	return nil
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) GetStddevRtt() *duration.Duration {
	if m != nil {
		return m.StddevRtt
	}
	return nil
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) GetTransmittedCounter() int32 {
	if m != nil {
		return m.TransmittedCounter
	}
	return int32(0)
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) GetReceivedCounter() int32 {
	if m != nil {
		return m.ReceivedCounter
	}
	return int32(0)
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) GetLossRatio() float64 {
	if m != nil {
		return m.LossRatio
	}
	return float64(0)
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) SetMinRtt(fv *duration.Duration) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MinRtt", "RunPingTestResponse_JsonResponse_SummaryStats"))
	}
	m.MinRtt = fv
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) SetAvgRtt(fv *duration.Duration) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AvgRtt", "RunPingTestResponse_JsonResponse_SummaryStats"))
	}
	m.AvgRtt = fv
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) SetMaxRtt(fv *duration.Duration) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "MaxRtt", "RunPingTestResponse_JsonResponse_SummaryStats"))
	}
	m.MaxRtt = fv
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) SetStddevRtt(fv *duration.Duration) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "StddevRtt", "RunPingTestResponse_JsonResponse_SummaryStats"))
	}
	m.StddevRtt = fv
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) SetTransmittedCounter(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TransmittedCounter", "RunPingTestResponse_JsonResponse_SummaryStats"))
	}
	m.TransmittedCounter = fv
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) SetReceivedCounter(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ReceivedCounter", "RunPingTestResponse_JsonResponse_SummaryStats"))
	}
	m.ReceivedCounter = fv
}

func (m *RunPingTestResponse_JsonResponse_SummaryStats) SetLossRatio(fv float64) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "LossRatio", "RunPingTestResponse_JsonResponse_SummaryStats"))
	}
	m.LossRatio = fv
}

var watchdog_proto_v1alpha2_ping_test_custom_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_ping_test_custom_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64,
	0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x04,
	0x0a, 0x12, 0x52, 0x75, 0x6e, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x09, 0x12, 0x07, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0xba, 0x9d, 0x22, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xca, 0xc6, 0x27, 0x06, 0x12,
	0x04, 0x2a, 0x02, 0x08, 0x02, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x2a, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x14, 0xca, 0xc6, 0x27, 0x06, 0x12, 0x04, 0x2a, 0x02, 0x08, 0x01, 0xca, 0xc6, 0x27, 0x06, 0x12,
	0x04, 0x1a, 0x02, 0x08, 0x64, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0xca, 0xc6, 0x27, 0x09, 0x5a,
	0x07, 0x32, 0x05, 0x10, 0x80, 0xc2, 0xd7, 0x2f, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x4b, 0x0a, 0x0c, 0x65, 0x63, 0x68, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0d, 0xca, 0xc6, 0x27, 0x09, 0x5a, 0x07, 0x32, 0x05, 0x10, 0x80, 0xc2,
	0xd7, 0x2f, 0x52, 0x0b, 0x65, 0x63, 0x68, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x6e, 0x74, 0x5f, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x6f, 0x6e, 0x74, 0x46, 0x72, 0x61, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x15, 0xca, 0xc6, 0x27, 0x06, 0x12, 0x04, 0x2a, 0x02, 0x08, 0x01, 0xca, 0xc6, 0x27,
	0x07, 0x12, 0x05, 0x1a, 0x03, 0x08, 0xff, 0x01, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x27, 0x0a,
	0x03, 0x74, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x42, 0x15, 0xca, 0xc6, 0x27, 0x06,
	0x12, 0x04, 0x2a, 0x02, 0x08, 0x00, 0xca, 0xc6, 0x27, 0x07, 0x12, 0x05, 0x1a, 0x03, 0x08, 0xff,
	0x01, 0x52, 0x03, 0x74, 0x6f, 0x73, 0x12, 0x56, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x52, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3a, 0x10,
	0xc2, 0x85, 0x2c, 0x0c, 0x32, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x9c, 0x06, 0x0a, 0x13, 0x52, 0x75, 0x6e, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x6a, 0x73, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x75, 0x6e, 0x50, 0x69, 0x6e, 0x67, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x81, 0x05, 0x0a, 0x0c,
	0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x2b, 0x0a, 0x03, 0x72, 0x74,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x03, 0x72, 0x74, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5e, 0x0a,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x75, 0x6e, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x1a, 0xdf, 0x02,
	0x0a, 0x0c, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x32,
	0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x74, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x52,
	0x74, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x61, 0x76, 0x67, 0x5f, 0x72, 0x74, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x61, 0x76, 0x67, 0x52, 0x74, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x74,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x52, 0x74, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x73, 0x74,
	0x64, 0x64, 0x65, 0x76, 0x5f, 0x72, 0x74, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x73, 0x74, 0x64, 0x64, 0x65,
	0x76, 0x52, 0x74, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x73, 0x73, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x42,
	0xcb, 0x01, 0xe8, 0xde, 0x21, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x42, 0x13, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x65, 0x73, 0x74, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x47, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e,
	0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x3b, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x44, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescData = watchdog_proto_v1alpha2_ping_test_custom_proto_rawDesc
)

func watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_ping_test_custom_proto_rawDescData
}

var watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var watchdog_proto_v1alpha2_ping_test_custom_proto_goTypes = []interface{}{
	(*RunPingTestRequest)(nil),                            // 0: ntt.watchdog.v1alpha2.RunPingTestRequest
	(*RunPingTestResponse)(nil),                           // 1: ntt.watchdog.v1alpha2.RunPingTestResponse
	(*RunPingTestResponse_JsonResponse)(nil),              // 2: ntt.watchdog.v1alpha2.RunPingTestResponse.JsonResponse
	(*RunPingTestResponse_JsonResponse_SummaryStats)(nil), // 3: ntt.watchdog.v1alpha2.RunPingTestResponse.JsonResponse.SummaryStats
	(*duration.Duration)(nil),                             // 4: google.protobuf.Duration
	(common.OnDemandTestResponseFormat)(0),                // 5: ntt.watchdog.v1alpha2.OnDemandTestResponseFormat
}
var watchdog_proto_v1alpha2_ping_test_custom_proto_depIdxs = []int32{
	4,  // 0: ntt.watchdog.v1alpha2.RunPingTestRequest.interval:type_name -> google.protobuf.Duration
	4,  // 1: ntt.watchdog.v1alpha2.RunPingTestRequest.echo_timeout:type_name -> google.protobuf.Duration
	5,  // 2: ntt.watchdog.v1alpha2.RunPingTestRequest.output_format:type_name -> ntt.watchdog.v1alpha2.OnDemandTestResponseFormat
	2,  // 3: ntt.watchdog.v1alpha2.RunPingTestResponse.json_response:type_name -> ntt.watchdog.v1alpha2.RunPingTestResponse.JsonResponse
	4,  // 4: ntt.watchdog.v1alpha2.RunPingTestResponse.JsonResponse.rtt:type_name -> google.protobuf.Duration
	3,  // 5: ntt.watchdog.v1alpha2.RunPingTestResponse.JsonResponse.summary:type_name -> ntt.watchdog.v1alpha2.RunPingTestResponse.JsonResponse.SummaryStats
	4,  // 6: ntt.watchdog.v1alpha2.RunPingTestResponse.JsonResponse.SummaryStats.min_rtt:type_name -> google.protobuf.Duration
	4,  // 7: ntt.watchdog.v1alpha2.RunPingTestResponse.JsonResponse.SummaryStats.avg_rtt:type_name -> google.protobuf.Duration
	4,  // 8: ntt.watchdog.v1alpha2.RunPingTestResponse.JsonResponse.SummaryStats.max_rtt:type_name -> google.protobuf.Duration
	4,  // 9: ntt.watchdog.v1alpha2.RunPingTestResponse.JsonResponse.SummaryStats.stddev_rtt:type_name -> google.protobuf.Duration
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_ping_test_custom_proto_init() }
func watchdog_proto_v1alpha2_ping_test_custom_proto_init() {
	if watchdog_proto_v1alpha2_ping_test_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunPingTestRequest); i {
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
		watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunPingTestResponse); i {
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
		watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunPingTestResponse_JsonResponse); i {
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
		watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunPingTestResponse_JsonResponse_SummaryStats); i {
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
			RawDescriptor: watchdog_proto_v1alpha2_ping_test_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_ping_test_custom_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_ping_test_custom_proto_depIdxs,
		MessageInfos:      watchdog_proto_v1alpha2_ping_test_custom_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_ping_test_custom_proto = out.File
	watchdog_proto_v1alpha2_ping_test_custom_proto_rawDesc = nil
	watchdog_proto_v1alpha2_ping_test_custom_proto_goTypes = nil
	watchdog_proto_v1alpha2_ping_test_custom_proto_depIdxs = nil
}
