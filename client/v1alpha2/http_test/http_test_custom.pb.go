// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/http_test_custom.proto
// DO NOT EDIT!!!

package http_test_client

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

// Request message for method [RunHTTPTest][ntt.watchdog.v1alpha2.RunHTTPTest]
type RunHTTPTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	//  reference of ntt.watchdog.v1alpha2.Probe
	Name *probe.Reference `protobuf:"bytes,1,opt,customtype=Reference,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Target URL to run the http test.
	// http/https scheme is mandatory
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty" firestore:"url"`
	// http headers to use in the http request.
	RequestHeaders map[string]string `protobuf:"bytes,3,rep,name=request_headers,json=requestHeaders,proto3" json:"request_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"requestHeaders"`
	// HTTP request method. Default is GET
	RequestMethod common.HTTPRequestMethod `protobuf:"varint,4,opt,name=request_method,json=requestMethod,proto3,enum=ntt.watchdog.v1alpha2.HTTPRequestMethod" json:"request_method,omitempty" firestore:"requestMethod"`
	// Request timeout duration. Default is 30 seconds.
	Timeout *duration.Duration `protobuf:"bytes,5,opt,name=timeout,proto3" json:"timeout,omitempty" firestore:"timeout"`
	// Request body for POST/PUT
	RequestBody string `protobuf:"bytes,6,opt,name=request_body,json=requestBody,proto3" json:"request_body,omitempty" firestore:"requestBody"`
	// Authentication method BASIC is allowed for http test
	AuthenticationMethod common.AuthenticationMethod `protobuf:"varint,7,opt,name=authentication_method,json=authenticationMethod,proto3,enum=ntt.watchdog.v1alpha2.AuthenticationMethod" json:"authentication_method,omitempty" firestore:"authenticationMethod"`
	// Username for basic auth
	Username string `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty" firestore:"username"`
	// Password for basic auth
	Password string `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty" firestore:"password"`
	// Source address to use for the outbound packet.
	// If unspecified, agent will bind to :: or 0.0.0.0 by default and
	// the operating system stack chooses the proper address.
	// The IP version is chosen according to the version of destination address
	// This is intended for advanced debug, it is recommended to leave this empty.
	SourceIp string `protobuf:"bytes,11,opt,name=source_ip,json=sourceIp,proto3" json:"source_ip,omitempty" firestore:"sourceIp"`
	// Default is Text format. Json is for internal use only
	OutputFormat common.OnDemandTestResponseFormat `protobuf:"varint,12,opt,name=output_format,json=outputFormat,proto3,enum=ntt.watchdog.v1alpha2.OnDemandTestResponseFormat" json:"output_format,omitempty" firestore:"outputFormat"`
}

func (m *RunHTTPTestRequest) Reset() {
	*m = RunHTTPTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RunHTTPTestRequest) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RunHTTPTestRequest) ProtoMessage() {}

func (m *RunHTTPTestRequest) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RunHTTPTestRequest) GotenMessage() {}

// Deprecated, Use RunHTTPTestRequest.ProtoReflect.Descriptor instead.
func (*RunHTTPTestRequest) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_http_test_custom_proto_rawDescGZIP(), []int{0}
}

func (m *RunHTTPTestRequest) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RunHTTPTestRequest) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RunHTTPTestRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RunHTTPTestRequest) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RunHTTPTestRequest) GetName() *probe.Reference {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RunHTTPTestRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RunHTTPTestRequest) GetRequestHeaders() map[string]string {
	if m != nil {
		return m.RequestHeaders
	}
	return nil
}

func (m *RunHTTPTestRequest) GetRequestMethod() common.HTTPRequestMethod {
	if m != nil {
		return m.RequestMethod
	}
	return common.HTTPRequestMethod_GET
}

func (m *RunHTTPTestRequest) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RunHTTPTestRequest) GetRequestBody() string {
	if m != nil {
		return m.RequestBody
	}
	return ""
}

func (m *RunHTTPTestRequest) GetAuthenticationMethod() common.AuthenticationMethod {
	if m != nil {
		return m.AuthenticationMethod
	}
	return common.AuthenticationMethod_NO_AUTH
}

func (m *RunHTTPTestRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RunHTTPTestRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RunHTTPTestRequest) GetSourceIp() string {
	if m != nil {
		return m.SourceIp
	}
	return ""
}

func (m *RunHTTPTestRequest) GetOutputFormat() common.OnDemandTestResponseFormat {
	if m != nil {
		return m.OutputFormat
	}
	return common.OnDemandTestResponseFormat_TEXT
}

func (m *RunHTTPTestRequest) SetName(fv *probe.Reference) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "RunHTTPTestRequest"))
	}
	m.Name = fv
}

func (m *RunHTTPTestRequest) SetUrl(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Url", "RunHTTPTestRequest"))
	}
	m.Url = fv
}

func (m *RunHTTPTestRequest) SetRequestHeaders(fv map[string]string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RequestHeaders", "RunHTTPTestRequest"))
	}
	m.RequestHeaders = fv
}

func (m *RunHTTPTestRequest) SetRequestMethod(fv common.HTTPRequestMethod) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RequestMethod", "RunHTTPTestRequest"))
	}
	m.RequestMethod = fv
}

func (m *RunHTTPTestRequest) SetTimeout(fv *duration.Duration) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Timeout", "RunHTTPTestRequest"))
	}
	m.Timeout = fv
}

func (m *RunHTTPTestRequest) SetRequestBody(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "RequestBody", "RunHTTPTestRequest"))
	}
	m.RequestBody = fv
}

func (m *RunHTTPTestRequest) SetAuthenticationMethod(fv common.AuthenticationMethod) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AuthenticationMethod", "RunHTTPTestRequest"))
	}
	m.AuthenticationMethod = fv
}

func (m *RunHTTPTestRequest) SetUsername(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Username", "RunHTTPTestRequest"))
	}
	m.Username = fv
}

func (m *RunHTTPTestRequest) SetPassword(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Password", "RunHTTPTestRequest"))
	}
	m.Password = fv
}

func (m *RunHTTPTestRequest) SetSourceIp(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "SourceIp", "RunHTTPTestRequest"))
	}
	m.SourceIp = fv
}

func (m *RunHTTPTestRequest) SetOutputFormat(fv common.OnDemandTestResponseFormat) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "OutputFormat", "RunHTTPTestRequest"))
	}
	m.OutputFormat = fv
}

// Response message for method [RunHTTPTest][ntt.watchdog.v1alpha2.RunHTTPTest]
type RunHTTPTestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Console type text response
	TextResponse string `protobuf:"bytes,1,opt,name=text_response,json=textResponse,proto3" json:"text_response,omitempty" firestore:"textResponse"`
	// Json format is not preferred for the ondemand tests
	JsonResponse *RunHTTPTestResponse_JsonResponse `protobuf:"bytes,2,opt,name=json_response,json=jsonResponse,proto3" json:"json_response,omitempty" firestore:"jsonResponse"`
}

func (m *RunHTTPTestResponse) Reset() {
	*m = RunHTTPTestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RunHTTPTestResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RunHTTPTestResponse) ProtoMessage() {}

func (m *RunHTTPTestResponse) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RunHTTPTestResponse) GotenMessage() {}

// Deprecated, Use RunHTTPTestResponse.ProtoReflect.Descriptor instead.
func (*RunHTTPTestResponse) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_http_test_custom_proto_rawDescGZIP(), []int{1}
}

func (m *RunHTTPTestResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RunHTTPTestResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RunHTTPTestResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RunHTTPTestResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RunHTTPTestResponse) GetTextResponse() string {
	if m != nil {
		return m.TextResponse
	}
	return ""
}

func (m *RunHTTPTestResponse) GetJsonResponse() *RunHTTPTestResponse_JsonResponse {
	if m != nil {
		return m.JsonResponse
	}
	return nil
}

func (m *RunHTTPTestResponse) SetTextResponse(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "TextResponse", "RunHTTPTestResponse"))
	}
	m.TextResponse = fv
}

func (m *RunHTTPTestResponse) SetJsonResponse(fv *RunHTTPTestResponse_JsonResponse) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "JsonResponse", "RunHTTPTestResponse"))
	}
	m.JsonResponse = fv
}

// Json format is not preferred for the ondemand tests
type RunHTTPTestResponse_JsonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// HTTP Response code
	ResponseCode int32 `protobuf:"varint,1,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty" firestore:"responseCode"`
	// HTTP Response Headers
	ResponseHeaders []*RunHTTPTestResponse_JsonResponse_Header `protobuf:"bytes,2,rep,name=response_headers,json=responseHeaders,proto3" json:"response_headers,omitempty" firestore:"responseHeaders"`
	// HTTP Response body
	ResponseBody []byte `protobuf:"bytes,3,opt,name=response_body,json=responseBody,proto3" json:"response_body,omitempty" firestore:"responseBody"`
}

func (m *RunHTTPTestResponse_JsonResponse) Reset() {
	*m = RunHTTPTestResponse_JsonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RunHTTPTestResponse_JsonResponse) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RunHTTPTestResponse_JsonResponse) ProtoMessage() {}

func (m *RunHTTPTestResponse_JsonResponse) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RunHTTPTestResponse_JsonResponse) GotenMessage() {}

// Deprecated, Use RunHTTPTestResponse_JsonResponse.ProtoReflect.Descriptor instead.
func (*RunHTTPTestResponse_JsonResponse) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_http_test_custom_proto_rawDescGZIP(), []int{1, 0}
}

func (m *RunHTTPTestResponse_JsonResponse) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RunHTTPTestResponse_JsonResponse) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RunHTTPTestResponse_JsonResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RunHTTPTestResponse_JsonResponse) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RunHTTPTestResponse_JsonResponse) GetResponseCode() int32 {
	if m != nil {
		return m.ResponseCode
	}
	return int32(0)
}

func (m *RunHTTPTestResponse_JsonResponse) GetResponseHeaders() []*RunHTTPTestResponse_JsonResponse_Header {
	if m != nil {
		return m.ResponseHeaders
	}
	return nil
}

func (m *RunHTTPTestResponse_JsonResponse) GetResponseBody() []byte {
	if m != nil {
		return m.ResponseBody
	}
	return nil
}

func (m *RunHTTPTestResponse_JsonResponse) SetResponseCode(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ResponseCode", "RunHTTPTestResponse_JsonResponse"))
	}
	m.ResponseCode = fv
}

func (m *RunHTTPTestResponse_JsonResponse) SetResponseHeaders(fv []*RunHTTPTestResponse_JsonResponse_Header) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ResponseHeaders", "RunHTTPTestResponse_JsonResponse"))
	}
	m.ResponseHeaders = fv
}

func (m *RunHTTPTestResponse_JsonResponse) SetResponseBody(fv []byte) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ResponseBody", "RunHTTPTestResponse_JsonResponse"))
	}
	m.ResponseBody = fv
}

// HTTP Response Header
type RunHTTPTestResponse_JsonResponse_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Key           string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" firestore:"key"`
	Values        []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" firestore:"values"`
}

func (m *RunHTTPTestResponse_JsonResponse_Header) Reset() {
	*m = RunHTTPTestResponse_JsonResponse_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *RunHTTPTestResponse_JsonResponse_Header) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*RunHTTPTestResponse_JsonResponse_Header) ProtoMessage() {}

func (m *RunHTTPTestResponse_JsonResponse_Header) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*RunHTTPTestResponse_JsonResponse_Header) GotenMessage() {}

// Deprecated, Use RunHTTPTestResponse_JsonResponse_Header.ProtoReflect.Descriptor instead.
func (*RunHTTPTestResponse_JsonResponse_Header) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_http_test_custom_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (m *RunHTTPTestResponse_JsonResponse_Header) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *RunHTTPTestResponse_JsonResponse_Header) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *RunHTTPTestResponse_JsonResponse_Header) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *RunHTTPTestResponse_JsonResponse_Header) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *RunHTTPTestResponse_JsonResponse_Header) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RunHTTPTestResponse_JsonResponse_Header) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *RunHTTPTestResponse_JsonResponse_Header) SetKey(fv string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Key", "RunHTTPTestResponse_JsonResponse_Header"))
	}
	m.Key = fv
}

func (m *RunHTTPTestResponse_JsonResponse_Header) SetValues(fv []string) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Values", "RunHTTPTestResponse_JsonResponse_Header"))
	}
	m.Values = fv
}

var watchdog_proto_v1alpha2_http_test_custom_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_http_test_custom_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74,
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
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x05,
	0x0a, 0x12, 0x52, 0x75, 0x6e, 0x48, 0x54, 0x54, 0x50, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x09, 0x12, 0x07, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0xba, 0x9d, 0x22, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x66, 0x0a,
	0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52,
	0x75, 0x6e, 0x48, 0x54, 0x54, 0x50, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x4f, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x60,
	0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x14, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x70, 0x12, 0x56, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x4f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52,
	0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x1a, 0x41, 0x0a,
	0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x3a, 0x10, 0xc2, 0x85, 0x2c, 0x0c, 0x32, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x22, 0x92, 0x03, 0x0a, 0x13, 0x52, 0x75, 0x6e,
	0x48, 0x54, 0x54, 0x50, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x75, 0x6e, 0x48, 0x54, 0x54, 0x50, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0xf7, 0x01, 0x0a, 0x0c, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x69, 0x0a, 0x10, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x75, 0x6e, 0x48,
	0x54, 0x54, 0x50, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x32, 0x0a, 0x06, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0xcb, 0x01,
	0xe8, 0xde, 0x21, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x42, 0x13, 0x48, 0x54, 0x54, 0x50, 0x54, 0x65, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x3b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x44, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_http_test_custom_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_http_test_custom_proto_rawDescData = watchdog_proto_v1alpha2_http_test_custom_proto_rawDesc
)

func watchdog_proto_v1alpha2_http_test_custom_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_http_test_custom_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_http_test_custom_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_http_test_custom_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_http_test_custom_proto_rawDescData
}

var watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var watchdog_proto_v1alpha2_http_test_custom_proto_goTypes = []interface{}{
	(*RunHTTPTestRequest)(nil),               // 0: ntt.watchdog.v1alpha2.RunHTTPTestRequest
	(*RunHTTPTestResponse)(nil),              // 1: ntt.watchdog.v1alpha2.RunHTTPTestResponse
	nil,                                      // 2: ntt.watchdog.v1alpha2.RunHTTPTestRequest.RequestHeadersEntry
	(*RunHTTPTestResponse_JsonResponse)(nil), // 3: ntt.watchdog.v1alpha2.RunHTTPTestResponse.JsonResponse
	(*RunHTTPTestResponse_JsonResponse_Header)(nil), // 4: ntt.watchdog.v1alpha2.RunHTTPTestResponse.JsonResponse.Header
	(common.HTTPRequestMethod)(0),                   // 5: ntt.watchdog.v1alpha2.HTTPRequestMethod
	(*duration.Duration)(nil),                       // 6: google.protobuf.Duration
	(common.AuthenticationMethod)(0),                // 7: ntt.watchdog.v1alpha2.AuthenticationMethod
	(common.OnDemandTestResponseFormat)(0),          // 8: ntt.watchdog.v1alpha2.OnDemandTestResponseFormat
}
var watchdog_proto_v1alpha2_http_test_custom_proto_depIdxs = []int32{
	2, // 0: ntt.watchdog.v1alpha2.RunHTTPTestRequest.request_headers:type_name -> ntt.watchdog.v1alpha2.RunHTTPTestRequest.RequestHeadersEntry
	5, // 1: ntt.watchdog.v1alpha2.RunHTTPTestRequest.request_method:type_name -> ntt.watchdog.v1alpha2.HTTPRequestMethod
	6, // 2: ntt.watchdog.v1alpha2.RunHTTPTestRequest.timeout:type_name -> google.protobuf.Duration
	7, // 3: ntt.watchdog.v1alpha2.RunHTTPTestRequest.authentication_method:type_name -> ntt.watchdog.v1alpha2.AuthenticationMethod
	8, // 4: ntt.watchdog.v1alpha2.RunHTTPTestRequest.output_format:type_name -> ntt.watchdog.v1alpha2.OnDemandTestResponseFormat
	3, // 5: ntt.watchdog.v1alpha2.RunHTTPTestResponse.json_response:type_name -> ntt.watchdog.v1alpha2.RunHTTPTestResponse.JsonResponse
	4, // 6: ntt.watchdog.v1alpha2.RunHTTPTestResponse.JsonResponse.response_headers:type_name -> ntt.watchdog.v1alpha2.RunHTTPTestResponse.JsonResponse.Header
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_http_test_custom_proto_init() }
func watchdog_proto_v1alpha2_http_test_custom_proto_init() {
	if watchdog_proto_v1alpha2_http_test_custom_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunHTTPTestRequest); i {
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
		watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunHTTPTestResponse); i {
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
		watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunHTTPTestResponse_JsonResponse); i {
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
		watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunHTTPTestResponse_JsonResponse_Header); i {
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
			RawDescriptor: watchdog_proto_v1alpha2_http_test_custom_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_http_test_custom_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_http_test_custom_proto_depIdxs,
		MessageInfos:      watchdog_proto_v1alpha2_http_test_custom_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_http_test_custom_proto = out.File
	watchdog_proto_v1alpha2_http_test_custom_proto_rawDesc = nil
	watchdog_proto_v1alpha2_http_test_custom_proto_goTypes = nil
	watchdog_proto_v1alpha2_http_test_custom_proto_depIdxs = nil
}
