// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/network_info_custom.proto
// DO NOT EDIT!!!

package network_info_client

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
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
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
	_ = &probe.Probe{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetNetworkInfoRequest_ResponseFormat int32

const (
	GetNetworkInfoRequest_JSON GetNetworkInfoRequest_ResponseFormat = 0
	GetNetworkInfoRequest_TEXT GetNetworkInfoRequest_ResponseFormat = 1
)

var (
	GetNetworkInfoRequest_ResponseFormat_name = map[int32]string{
		0: "JSON",
		1: "TEXT",
	}

	GetNetworkInfoRequest_ResponseFormat_value = map[string]int32{
		"JSON": 0,
		"TEXT": 1,
	}
)

func (x GetNetworkInfoRequest_ResponseFormat) Enum() *GetNetworkInfoRequest_ResponseFormat {
	p := new(GetNetworkInfoRequest_ResponseFormat)
	*p = x
	return p
}

func (x GetNetworkInfoRequest_ResponseFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (GetNetworkInfoRequest_ResponseFormat) Descriptor() preflect.EnumDescriptor {
	return watchdog_proto_v1alpha2_network_info_custom_proto_enumTypes[0].Descriptor()
}

func (GetNetworkInfoRequest_ResponseFormat) Type() preflect.EnumType {
	return &watchdog_proto_v1alpha2_network_info_custom_proto_enumTypes[0]
}

func (x GetNetworkInfoRequest_ResponseFormat) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use GetNetworkInfoRequest_ResponseFormat.ProtoReflect.Descriptor instead.
func (GetNetworkInfoRequest_ResponseFormat) EnumDescriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_network_info_custom_proto_rawDescGZIP(), []int{0, 0}
}

// Request message for method
// [GetNetworkInfo][ntt.watchdog.v1alpha2.GetNetworkInfo]
type GetNetworkInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  reference of ntt.watchdog.v1alpha2.Probe
	Name           *probe.Reference                     `protobuf:"bytes,1,opt,customtype=Reference,name=name,proto3" json:"name,omitempty" firestore:"name"`
	ResponseFornat GetNetworkInfoRequest_ResponseFormat `protobuf:"varint,2,opt,name=response_fornat,json=responseFornat,proto3,enum=ntt.watchdog.v1alpha2.GetNetworkInfoRequest_ResponseFormat" json:"response_fornat,omitempty" firestore:"responseFornat"`
}

func (m *GetNetworkInfoRequest) Reset() {
	*m = GetNetworkInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GetNetworkInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GetNetworkInfoRequest) ProtoMessage() {}

func (m *GetNetworkInfoRequest) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GetNetworkInfoRequest) GotenMessage() {}

// Deprecated, Use GetNetworkInfoRequest.ProtoReflect.Descriptor instead.
func (*GetNetworkInfoRequest) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_network_info_custom_proto_rawDescGZIP(), []int{0}
}

func (m *GetNetworkInfoRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GetNetworkInfoRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GetNetworkInfoRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GetNetworkInfoRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GetNetworkInfoRequest) GetName() *probe.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GetNetworkInfoRequest) GetResponseFornat() GetNetworkInfoRequest_ResponseFormat {
	if m != nil {
		return m.ResponseFornat
	}
	return GetNetworkInfoRequest_JSON
}

func (m *GetNetworkInfoRequest) SetName(fv *probe.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "GetNetworkInfoRequest"))
	}
	m.Name = fv
}

func (m *GetNetworkInfoRequest) SetResponseFornat(fv GetNetworkInfoRequest_ResponseFormat) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ResponseFornat", "GetNetworkInfoRequest"))
	}
	m.ResponseFornat = fv
}

// Response message for method
// [GetNetworkInfo][ntt.watchdog.v1alpha2.GetNetworkInfo]
type GetNetworkInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	TextResponse  string                               `protobuf:"bytes,1,opt,name=text_response,json=textResponse,proto3" json:"text_response,omitempty" firestore:"textResponse"`
	JsonResponse  *GetNetworkInfoResponse_JsonResponse `protobuf:"bytes,2,opt,name=json_response,json=jsonResponse,proto3" json:"json_response,omitempty" firestore:"jsonResponse"`
}

func (m *GetNetworkInfoResponse) Reset() {
	*m = GetNetworkInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GetNetworkInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GetNetworkInfoResponse) ProtoMessage() {}

func (m *GetNetworkInfoResponse) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GetNetworkInfoResponse) GotenMessage() {}

// Deprecated, Use GetNetworkInfoResponse.ProtoReflect.Descriptor instead.
func (*GetNetworkInfoResponse) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_network_info_custom_proto_rawDescGZIP(), []int{1}
}

func (m *GetNetworkInfoResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GetNetworkInfoResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GetNetworkInfoResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GetNetworkInfoResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GetNetworkInfoResponse) GetTextResponse() string {
	if m != nil {
		return m.TextResponse
	}
	return ""
}

func (m *GetNetworkInfoResponse) GetJsonResponse() *GetNetworkInfoResponse_JsonResponse {
	if m != nil {
		return m.JsonResponse
	}
	return nil
}

func (m *GetNetworkInfoResponse) SetTextResponse(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TextResponse", "GetNetworkInfoResponse"))
	}
	m.TextResponse = fv
}

func (m *GetNetworkInfoResponse) SetJsonResponse(fv *GetNetworkInfoResponse_JsonResponse) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "JsonResponse", "GetNetworkInfoResponse"))
	}
	m.JsonResponse = fv
}

type GetNetworkInfoResponse_JsonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Interfaces    string `protobuf:"bytes,1,opt,name=interfaces,proto3" json:"interfaces,omitempty" firestore:"interfaces"`
	Routes        string `protobuf:"bytes,2,opt,name=routes,proto3" json:"routes,omitempty" firestore:"routes"`
	WifiInfo      string `protobuf:"bytes,3,opt,name=wifi_info,json=wifiInfo,proto3" json:"wifi_info,omitempty" firestore:"wifiInfo"`
}

func (m *GetNetworkInfoResponse_JsonResponse) Reset() {
	*m = GetNetworkInfoResponse_JsonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *GetNetworkInfoResponse_JsonResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*GetNetworkInfoResponse_JsonResponse) ProtoMessage() {}

func (m *GetNetworkInfoResponse_JsonResponse) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*GetNetworkInfoResponse_JsonResponse) GotenMessage() {}

// Deprecated, Use GetNetworkInfoResponse_JsonResponse.ProtoReflect.Descriptor instead.
func (*GetNetworkInfoResponse_JsonResponse) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_network_info_custom_proto_rawDescGZIP(), []int{1, 0}
}

func (m *GetNetworkInfoResponse_JsonResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *GetNetworkInfoResponse_JsonResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *GetNetworkInfoResponse_JsonResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *GetNetworkInfoResponse_JsonResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *GetNetworkInfoResponse_JsonResponse) GetInterfaces() string {
	if m != nil {
		return m.Interfaces
	}
	return ""
}

func (m *GetNetworkInfoResponse_JsonResponse) GetRoutes() string {
	if m != nil {
		return m.Routes
	}
	return ""
}

func (m *GetNetworkInfoResponse_JsonResponse) GetWifiInfo() string {
	if m != nil {
		return m.WifiInfo
	}
	return ""
}

func (m *GetNetworkInfoResponse_JsonResponse) SetInterfaces(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Interfaces", "GetNetworkInfoResponse_JsonResponse"))
	}
	m.Interfaces = fv
}

func (m *GetNetworkInfoResponse_JsonResponse) SetRoutes(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Routes", "GetNetworkInfoResponse_JsonResponse"))
	}
	m.Routes = fv
}

func (m *GetNetworkInfoResponse_JsonResponse) SetWifiInfo(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "WifiInfo", "GetNetworkInfoResponse_JsonResponse"))
	}
	m.WifiInfo = fv
}

var watchdog_proto_v1alpha2_network_info_custom_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_network_info_custom_proto_rawDesc = []byte{
	0x0a, 0x31, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x09, 0x12, 0x07, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0xba, 0x9d, 0x22, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x64, 0x0a, 0x0f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6e, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x47, 0x65, 0x74,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x6e, 0x61,
	0x74, 0x22, 0x24, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x01, 0x3a, 0x10, 0xc2, 0x85, 0x2c, 0x0c, 0x32, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x3a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x83, 0x02, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0d, 0x6a, 0x73, 0x6f,
	0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x6a, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x63, 0x0a, 0x0c, 0x4a, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x69, 0x66, 0x69, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x69, 0x66, 0x69, 0x49, 0x6e, 0x66, 0x6f, 0x42,
	0xd4, 0x01, 0xe8, 0xde, 0x21, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x42, 0x16, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x4d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xd2, 0x84, 0xd1,
	0x02, 0x44, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_network_info_custom_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_network_info_custom_proto_rawDescData = watchdog_proto_v1alpha2_network_info_custom_proto_rawDesc
)

func watchdog_proto_v1alpha2_network_info_custom_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_network_info_custom_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_network_info_custom_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_network_info_custom_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_network_info_custom_proto_rawDescData
}

var watchdog_proto_v1alpha2_network_info_custom_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var watchdog_proto_v1alpha2_network_info_custom_proto_goTypes = []interface{}{
	(GetNetworkInfoRequest_ResponseFormat)(0),   // 0: ntt.watchdog.v1alpha2.GetNetworkInfoRequest_ResponseFormat
	(*GetNetworkInfoRequest)(nil),               // 1: ntt.watchdog.v1alpha2.GetNetworkInfoRequest
	(*GetNetworkInfoResponse)(nil),              // 2: ntt.watchdog.v1alpha2.GetNetworkInfoResponse
	(*GetNetworkInfoResponse_JsonResponse)(nil), // 3: ntt.watchdog.v1alpha2.GetNetworkInfoResponse.JsonResponse
}
var watchdog_proto_v1alpha2_network_info_custom_proto_depIdxs = []int32{
	0, // 0: ntt.watchdog.v1alpha2.GetNetworkInfoRequest.response_fornat:type_name -> ntt.watchdog.v1alpha2.GetNetworkInfoRequest_ResponseFormat
	3, // 1: ntt.watchdog.v1alpha2.GetNetworkInfoResponse.json_response:type_name -> ntt.watchdog.v1alpha2.GetNetworkInfoResponse.JsonResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_network_info_custom_proto_init() }
func watchdog_proto_v1alpha2_network_info_custom_proto_init() {
	if watchdog_proto_v1alpha2_network_info_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworkInfoRequest); i {
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
		watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworkInfoResponse); i {
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
		watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworkInfoResponse_JsonResponse); i {
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
			RawDescriptor: watchdog_proto_v1alpha2_network_info_custom_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_network_info_custom_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_network_info_custom_proto_depIdxs,
		EnumInfos:         watchdog_proto_v1alpha2_network_info_custom_proto_enumTypes,
		MessageInfos:      watchdog_proto_v1alpha2_network_info_custom_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_network_info_custom_proto = out.File
	watchdog_proto_v1alpha2_network_info_custom_proto_rawDesc = nil
	watchdog_proto_v1alpha2_network_info_custom_proto_goTypes = nil
	watchdog_proto_v1alpha2_network_info_custom_proto_depIdxs = nil
}
