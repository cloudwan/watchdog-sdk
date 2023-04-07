// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/network_info_custom.proto
// DO NOT EDIT!!!

package network_info_client

// proto imports
import (
	ntt_memo "github.com/cloudwan/edgelq-sdk/common/types/memo"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
	devices_device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device"
	devices_project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
	iam_attestation_domain "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/attestation_domain"
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	meta_service "github.com/cloudwan/edgelq-sdk/meta/resources/v1alpha2/service"
	admin_area "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/admin_area"
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probe_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// make sure we're using proto imports
var (
	_ = &ntt_memo.Memo{}
	_ = &ntt_meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
	_ = &devices_device.Device{}
	_ = &devices_project.Project{}
	_ = &iam_attestation_domain.AttestationDomain{}
	_ = &iam_iam_common.PCR{}
	_ = &iam_organization.Organization{}
	_ = &iam_project.Project{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &meta_service.Service{}
	_ = &duration.Duration{}
	_ = &field_mask.FieldMask{}
	_ = &timestamp.Timestamp{}
	_ = &wrappers.DoubleValue{}
	_ = &latlng.LatLng{}
	_ = &admin_area.BBox{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probe_group.ProbeGroup{}
	_ = &project.Project{}
)

type GetNetworkInfoRequestFieldPathBuilder struct{}

func NewGetNetworkInfoRequestFieldPathBuilder() GetNetworkInfoRequestFieldPathBuilder {
	return GetNetworkInfoRequestFieldPathBuilder{}
}
func (GetNetworkInfoRequestFieldPathBuilder) Name() GetNetworkInfoRequestPathSelectorName {
	return GetNetworkInfoRequestPathSelectorName{}
}
func (GetNetworkInfoRequestFieldPathBuilder) ResponseFornat() GetNetworkInfoRequestPathSelectorResponseFornat {
	return GetNetworkInfoRequestPathSelectorResponseFornat{}
}

type GetNetworkInfoRequestPathSelectorName struct{}

func (GetNetworkInfoRequestPathSelectorName) FieldPath() *GetNetworkInfoRequest_FieldTerminalPath {
	return &GetNetworkInfoRequest_FieldTerminalPath{selector: GetNetworkInfoRequest_FieldPathSelectorName}
}

func (s GetNetworkInfoRequestPathSelectorName) WithValue(value *probe.Reference) *GetNetworkInfoRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetNetworkInfoRequest_FieldTerminalPathValue)
}

func (s GetNetworkInfoRequestPathSelectorName) WithArrayOfValues(values []*probe.Reference) *GetNetworkInfoRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetNetworkInfoRequest_FieldTerminalPathArrayOfValues)
}

type GetNetworkInfoRequestPathSelectorResponseFornat struct{}

func (GetNetworkInfoRequestPathSelectorResponseFornat) FieldPath() *GetNetworkInfoRequest_FieldTerminalPath {
	return &GetNetworkInfoRequest_FieldTerminalPath{selector: GetNetworkInfoRequest_FieldPathSelectorResponseFornat}
}

func (s GetNetworkInfoRequestPathSelectorResponseFornat) WithValue(value GetNetworkInfoRequest_ResponseFormat) *GetNetworkInfoRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetNetworkInfoRequest_FieldTerminalPathValue)
}

func (s GetNetworkInfoRequestPathSelectorResponseFornat) WithArrayOfValues(values []GetNetworkInfoRequest_ResponseFormat) *GetNetworkInfoRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetNetworkInfoRequest_FieldTerminalPathArrayOfValues)
}

type GetNetworkInfoResponseFieldPathBuilder struct{}

func NewGetNetworkInfoResponseFieldPathBuilder() GetNetworkInfoResponseFieldPathBuilder {
	return GetNetworkInfoResponseFieldPathBuilder{}
}
func (GetNetworkInfoResponseFieldPathBuilder) TextResponse() GetNetworkInfoResponsePathSelectorTextResponse {
	return GetNetworkInfoResponsePathSelectorTextResponse{}
}
func (GetNetworkInfoResponseFieldPathBuilder) JsonResponse() GetNetworkInfoResponsePathSelectorJsonResponse {
	return GetNetworkInfoResponsePathSelectorJsonResponse{}
}

type GetNetworkInfoResponsePathSelectorTextResponse struct{}

func (GetNetworkInfoResponsePathSelectorTextResponse) FieldPath() *GetNetworkInfoResponse_FieldTerminalPath {
	return &GetNetworkInfoResponse_FieldTerminalPath{selector: GetNetworkInfoResponse_FieldPathSelectorTextResponse}
}

func (s GetNetworkInfoResponsePathSelectorTextResponse) WithValue(value string) *GetNetworkInfoResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetNetworkInfoResponse_FieldTerminalPathValue)
}

func (s GetNetworkInfoResponsePathSelectorTextResponse) WithArrayOfValues(values []string) *GetNetworkInfoResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetNetworkInfoResponse_FieldTerminalPathArrayOfValues)
}

type GetNetworkInfoResponsePathSelectorJsonResponse struct{}

func (GetNetworkInfoResponsePathSelectorJsonResponse) FieldPath() *GetNetworkInfoResponse_FieldTerminalPath {
	return &GetNetworkInfoResponse_FieldTerminalPath{selector: GetNetworkInfoResponse_FieldPathSelectorJsonResponse}
}

func (s GetNetworkInfoResponsePathSelectorJsonResponse) WithValue(value *GetNetworkInfoResponse_JsonResponse) *GetNetworkInfoResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetNetworkInfoResponse_FieldTerminalPathValue)
}

func (s GetNetworkInfoResponsePathSelectorJsonResponse) WithArrayOfValues(values []*GetNetworkInfoResponse_JsonResponse) *GetNetworkInfoResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetNetworkInfoResponse_FieldTerminalPathArrayOfValues)
}

func (GetNetworkInfoResponsePathSelectorJsonResponse) WithSubPath(subPath GetNetworkInfoResponseJsonResponse_FieldPath) *GetNetworkInfoResponse_FieldSubPath {
	return &GetNetworkInfoResponse_FieldSubPath{selector: GetNetworkInfoResponse_FieldPathSelectorJsonResponse, subPath: subPath}
}

func (s GetNetworkInfoResponsePathSelectorJsonResponse) WithSubValue(subPathValue GetNetworkInfoResponseJsonResponse_FieldPathValue) *GetNetworkInfoResponse_FieldSubPathValue {
	return &GetNetworkInfoResponse_FieldSubPathValue{GetNetworkInfoResponse_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s GetNetworkInfoResponsePathSelectorJsonResponse) WithSubArrayOfValues(subPathArrayOfValues GetNetworkInfoResponseJsonResponse_FieldPathArrayOfValues) *GetNetworkInfoResponse_FieldSubPathArrayOfValues {
	return &GetNetworkInfoResponse_FieldSubPathArrayOfValues{GetNetworkInfoResponse_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s GetNetworkInfoResponsePathSelectorJsonResponse) WithSubArrayItemValue(subPathArrayItemValue GetNetworkInfoResponseJsonResponse_FieldPathArrayItemValue) *GetNetworkInfoResponse_FieldSubPathArrayItemValue {
	return &GetNetworkInfoResponse_FieldSubPathArrayItemValue{GetNetworkInfoResponse_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (GetNetworkInfoResponsePathSelectorJsonResponse) Interfaces() GetNetworkInfoResponsePathSelectorJsonResponseInterfaces {
	return GetNetworkInfoResponsePathSelectorJsonResponseInterfaces{}
}

func (GetNetworkInfoResponsePathSelectorJsonResponse) Routes() GetNetworkInfoResponsePathSelectorJsonResponseRoutes {
	return GetNetworkInfoResponsePathSelectorJsonResponseRoutes{}
}

func (GetNetworkInfoResponsePathSelectorJsonResponse) WifiInfo() GetNetworkInfoResponsePathSelectorJsonResponseWifiInfo {
	return GetNetworkInfoResponsePathSelectorJsonResponseWifiInfo{}
}

type GetNetworkInfoResponsePathSelectorJsonResponseInterfaces struct{}

func (GetNetworkInfoResponsePathSelectorJsonResponseInterfaces) FieldPath() *GetNetworkInfoResponse_FieldSubPath {
	return &GetNetworkInfoResponse_FieldSubPath{
		selector: GetNetworkInfoResponse_FieldPathSelectorJsonResponse,
		subPath:  NewGetNetworkInfoResponseJsonResponseFieldPathBuilder().Interfaces().FieldPath(),
	}
}

func (s GetNetworkInfoResponsePathSelectorJsonResponseInterfaces) WithValue(value string) *GetNetworkInfoResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GetNetworkInfoResponse_FieldSubPathValue)
}

func (s GetNetworkInfoResponsePathSelectorJsonResponseInterfaces) WithArrayOfValues(values []string) *GetNetworkInfoResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetNetworkInfoResponse_FieldSubPathArrayOfValues)
}

type GetNetworkInfoResponsePathSelectorJsonResponseRoutes struct{}

func (GetNetworkInfoResponsePathSelectorJsonResponseRoutes) FieldPath() *GetNetworkInfoResponse_FieldSubPath {
	return &GetNetworkInfoResponse_FieldSubPath{
		selector: GetNetworkInfoResponse_FieldPathSelectorJsonResponse,
		subPath:  NewGetNetworkInfoResponseJsonResponseFieldPathBuilder().Routes().FieldPath(),
	}
}

func (s GetNetworkInfoResponsePathSelectorJsonResponseRoutes) WithValue(value string) *GetNetworkInfoResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GetNetworkInfoResponse_FieldSubPathValue)
}

func (s GetNetworkInfoResponsePathSelectorJsonResponseRoutes) WithArrayOfValues(values []string) *GetNetworkInfoResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetNetworkInfoResponse_FieldSubPathArrayOfValues)
}

type GetNetworkInfoResponsePathSelectorJsonResponseWifiInfo struct{}

func (GetNetworkInfoResponsePathSelectorJsonResponseWifiInfo) FieldPath() *GetNetworkInfoResponse_FieldSubPath {
	return &GetNetworkInfoResponse_FieldSubPath{
		selector: GetNetworkInfoResponse_FieldPathSelectorJsonResponse,
		subPath:  NewGetNetworkInfoResponseJsonResponseFieldPathBuilder().WifiInfo().FieldPath(),
	}
}

func (s GetNetworkInfoResponsePathSelectorJsonResponseWifiInfo) WithValue(value string) *GetNetworkInfoResponse_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GetNetworkInfoResponse_FieldSubPathValue)
}

func (s GetNetworkInfoResponsePathSelectorJsonResponseWifiInfo) WithArrayOfValues(values []string) *GetNetworkInfoResponse_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetNetworkInfoResponse_FieldSubPathArrayOfValues)
}

type GetNetworkInfoResponseJsonResponseFieldPathBuilder struct{}

func NewGetNetworkInfoResponseJsonResponseFieldPathBuilder() GetNetworkInfoResponseJsonResponseFieldPathBuilder {
	return GetNetworkInfoResponseJsonResponseFieldPathBuilder{}
}
func (GetNetworkInfoResponseJsonResponseFieldPathBuilder) Interfaces() GetNetworkInfoResponse_JsonResponsePathSelectorInterfaces {
	return GetNetworkInfoResponse_JsonResponsePathSelectorInterfaces{}
}
func (GetNetworkInfoResponseJsonResponseFieldPathBuilder) Routes() GetNetworkInfoResponse_JsonResponsePathSelectorRoutes {
	return GetNetworkInfoResponse_JsonResponsePathSelectorRoutes{}
}
func (GetNetworkInfoResponseJsonResponseFieldPathBuilder) WifiInfo() GetNetworkInfoResponse_JsonResponsePathSelectorWifiInfo {
	return GetNetworkInfoResponse_JsonResponsePathSelectorWifiInfo{}
}

type GetNetworkInfoResponse_JsonResponsePathSelectorInterfaces struct{}

func (GetNetworkInfoResponse_JsonResponsePathSelectorInterfaces) FieldPath() *GetNetworkInfoResponseJsonResponse_FieldTerminalPath {
	return &GetNetworkInfoResponseJsonResponse_FieldTerminalPath{selector: GetNetworkInfoResponseJsonResponse_FieldPathSelectorInterfaces}
}

func (s GetNetworkInfoResponse_JsonResponsePathSelectorInterfaces) WithValue(value string) *GetNetworkInfoResponseJsonResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetNetworkInfoResponseJsonResponse_FieldTerminalPathValue)
}

func (s GetNetworkInfoResponse_JsonResponsePathSelectorInterfaces) WithArrayOfValues(values []string) *GetNetworkInfoResponseJsonResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetNetworkInfoResponseJsonResponse_FieldTerminalPathArrayOfValues)
}

type GetNetworkInfoResponse_JsonResponsePathSelectorRoutes struct{}

func (GetNetworkInfoResponse_JsonResponsePathSelectorRoutes) FieldPath() *GetNetworkInfoResponseJsonResponse_FieldTerminalPath {
	return &GetNetworkInfoResponseJsonResponse_FieldTerminalPath{selector: GetNetworkInfoResponseJsonResponse_FieldPathSelectorRoutes}
}

func (s GetNetworkInfoResponse_JsonResponsePathSelectorRoutes) WithValue(value string) *GetNetworkInfoResponseJsonResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetNetworkInfoResponseJsonResponse_FieldTerminalPathValue)
}

func (s GetNetworkInfoResponse_JsonResponsePathSelectorRoutes) WithArrayOfValues(values []string) *GetNetworkInfoResponseJsonResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetNetworkInfoResponseJsonResponse_FieldTerminalPathArrayOfValues)
}

type GetNetworkInfoResponse_JsonResponsePathSelectorWifiInfo struct{}

func (GetNetworkInfoResponse_JsonResponsePathSelectorWifiInfo) FieldPath() *GetNetworkInfoResponseJsonResponse_FieldTerminalPath {
	return &GetNetworkInfoResponseJsonResponse_FieldTerminalPath{selector: GetNetworkInfoResponseJsonResponse_FieldPathSelectorWifiInfo}
}

func (s GetNetworkInfoResponse_JsonResponsePathSelectorWifiInfo) WithValue(value string) *GetNetworkInfoResponseJsonResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetNetworkInfoResponseJsonResponse_FieldTerminalPathValue)
}

func (s GetNetworkInfoResponse_JsonResponsePathSelectorWifiInfo) WithArrayOfValues(values []string) *GetNetworkInfoResponseJsonResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetNetworkInfoResponseJsonResponse_FieldTerminalPathArrayOfValues)
}
