// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/pcap_custom.proto
// DO NOT EDIT!!!

package pcap_client

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

type ReportPcapRequestFieldPathBuilder struct{}

func NewReportPcapRequestFieldPathBuilder() ReportPcapRequestFieldPathBuilder {
	return ReportPcapRequestFieldPathBuilder{}
}
func (ReportPcapRequestFieldPathBuilder) Name() ReportPcapRequestPathSelectorName {
	return ReportPcapRequestPathSelectorName{}
}
func (ReportPcapRequestFieldPathBuilder) Filter() ReportPcapRequestPathSelectorFilter {
	return ReportPcapRequestPathSelectorFilter{}
}
func (ReportPcapRequestFieldPathBuilder) StartTime() ReportPcapRequestPathSelectorStartTime {
	return ReportPcapRequestPathSelectorStartTime{}
}
func (ReportPcapRequestFieldPathBuilder) EndTime() ReportPcapRequestPathSelectorEndTime {
	return ReportPcapRequestPathSelectorEndTime{}
}
func (ReportPcapRequestFieldPathBuilder) PcapBytes() ReportPcapRequestPathSelectorPcapBytes {
	return ReportPcapRequestPathSelectorPcapBytes{}
}

type ReportPcapRequestPathSelectorName struct{}

func (ReportPcapRequestPathSelectorName) FieldPath() *ReportPcapRequest_FieldTerminalPath {
	return &ReportPcapRequest_FieldTerminalPath{selector: ReportPcapRequest_FieldPathSelectorName}
}

func (s ReportPcapRequestPathSelectorName) WithValue(value *probe.Reference) *ReportPcapRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ReportPcapRequest_FieldTerminalPathValue)
}

func (s ReportPcapRequestPathSelectorName) WithArrayOfValues(values []*probe.Reference) *ReportPcapRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ReportPcapRequest_FieldTerminalPathArrayOfValues)
}

type ReportPcapRequestPathSelectorFilter struct{}

func (ReportPcapRequestPathSelectorFilter) FieldPath() *ReportPcapRequest_FieldTerminalPath {
	return &ReportPcapRequest_FieldTerminalPath{selector: ReportPcapRequest_FieldPathSelectorFilter}
}

func (s ReportPcapRequestPathSelectorFilter) WithValue(value string) *ReportPcapRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ReportPcapRequest_FieldTerminalPathValue)
}

func (s ReportPcapRequestPathSelectorFilter) WithArrayOfValues(values []string) *ReportPcapRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ReportPcapRequest_FieldTerminalPathArrayOfValues)
}

type ReportPcapRequestPathSelectorStartTime struct{}

func (ReportPcapRequestPathSelectorStartTime) FieldPath() *ReportPcapRequest_FieldTerminalPath {
	return &ReportPcapRequest_FieldTerminalPath{selector: ReportPcapRequest_FieldPathSelectorStartTime}
}

func (s ReportPcapRequestPathSelectorStartTime) WithValue(value *timestamp.Timestamp) *ReportPcapRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ReportPcapRequest_FieldTerminalPathValue)
}

func (s ReportPcapRequestPathSelectorStartTime) WithArrayOfValues(values []*timestamp.Timestamp) *ReportPcapRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ReportPcapRequest_FieldTerminalPathArrayOfValues)
}

type ReportPcapRequestPathSelectorEndTime struct{}

func (ReportPcapRequestPathSelectorEndTime) FieldPath() *ReportPcapRequest_FieldTerminalPath {
	return &ReportPcapRequest_FieldTerminalPath{selector: ReportPcapRequest_FieldPathSelectorEndTime}
}

func (s ReportPcapRequestPathSelectorEndTime) WithValue(value *timestamp.Timestamp) *ReportPcapRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ReportPcapRequest_FieldTerminalPathValue)
}

func (s ReportPcapRequestPathSelectorEndTime) WithArrayOfValues(values []*timestamp.Timestamp) *ReportPcapRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ReportPcapRequest_FieldTerminalPathArrayOfValues)
}

type ReportPcapRequestPathSelectorPcapBytes struct{}

func (ReportPcapRequestPathSelectorPcapBytes) FieldPath() *ReportPcapRequest_FieldTerminalPath {
	return &ReportPcapRequest_FieldTerminalPath{selector: ReportPcapRequest_FieldPathSelectorPcapBytes}
}

func (s ReportPcapRequestPathSelectorPcapBytes) WithValue(value []byte) *ReportPcapRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ReportPcapRequest_FieldTerminalPathValue)
}

func (s ReportPcapRequestPathSelectorPcapBytes) WithArrayOfValues(values [][]byte) *ReportPcapRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ReportPcapRequest_FieldTerminalPathArrayOfValues)
}

type GetPcapRequestFieldPathBuilder struct{}

func NewGetPcapRequestFieldPathBuilder() GetPcapRequestFieldPathBuilder {
	return GetPcapRequestFieldPathBuilder{}
}
func (GetPcapRequestFieldPathBuilder) Probe() GetPcapRequestPathSelectorProbe {
	return GetPcapRequestPathSelectorProbe{}
}
func (GetPcapRequestFieldPathBuilder) Interval() GetPcapRequestPathSelectorInterval {
	return GetPcapRequestPathSelectorInterval{}
}
func (GetPcapRequestFieldPathBuilder) PageSize() GetPcapRequestPathSelectorPageSize {
	return GetPcapRequestPathSelectorPageSize{}
}
func (GetPcapRequestFieldPathBuilder) PageToken() GetPcapRequestPathSelectorPageToken {
	return GetPcapRequestPathSelectorPageToken{}
}

type GetPcapRequestPathSelectorProbe struct{}

func (GetPcapRequestPathSelectorProbe) FieldPath() *GetPcapRequest_FieldTerminalPath {
	return &GetPcapRequest_FieldTerminalPath{selector: GetPcapRequest_FieldPathSelectorProbe}
}

func (s GetPcapRequestPathSelectorProbe) WithValue(value *probe.Reference) *GetPcapRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapRequest_FieldTerminalPathValue)
}

func (s GetPcapRequestPathSelectorProbe) WithArrayOfValues(values []*probe.Reference) *GetPcapRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapRequest_FieldTerminalPathArrayOfValues)
}

type GetPcapRequestPathSelectorInterval struct{}

func (GetPcapRequestPathSelectorInterval) FieldPath() *GetPcapRequest_FieldTerminalPath {
	return &GetPcapRequest_FieldTerminalPath{selector: GetPcapRequest_FieldPathSelectorInterval}
}

func (s GetPcapRequestPathSelectorInterval) WithValue(value *common.TimeInterval) *GetPcapRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapRequest_FieldTerminalPathValue)
}

func (s GetPcapRequestPathSelectorInterval) WithArrayOfValues(values []*common.TimeInterval) *GetPcapRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapRequest_FieldTerminalPathArrayOfValues)
}

func (GetPcapRequestPathSelectorInterval) WithSubPath(subPath common.TimeInterval_FieldPath) *GetPcapRequest_FieldSubPath {
	return &GetPcapRequest_FieldSubPath{selector: GetPcapRequest_FieldPathSelectorInterval, subPath: subPath}
}

func (s GetPcapRequestPathSelectorInterval) WithSubValue(subPathValue common.TimeInterval_FieldPathValue) *GetPcapRequest_FieldSubPathValue {
	return &GetPcapRequest_FieldSubPathValue{GetPcapRequest_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s GetPcapRequestPathSelectorInterval) WithSubArrayOfValues(subPathArrayOfValues common.TimeInterval_FieldPathArrayOfValues) *GetPcapRequest_FieldSubPathArrayOfValues {
	return &GetPcapRequest_FieldSubPathArrayOfValues{GetPcapRequest_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s GetPcapRequestPathSelectorInterval) WithSubArrayItemValue(subPathArrayItemValue common.TimeInterval_FieldPathArrayItemValue) *GetPcapRequest_FieldSubPathArrayItemValue {
	return &GetPcapRequest_FieldSubPathArrayItemValue{GetPcapRequest_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (GetPcapRequestPathSelectorInterval) EndTime() GetPcapRequestPathSelectorIntervalEndTime {
	return GetPcapRequestPathSelectorIntervalEndTime{}
}

func (GetPcapRequestPathSelectorInterval) StartTime() GetPcapRequestPathSelectorIntervalStartTime {
	return GetPcapRequestPathSelectorIntervalStartTime{}
}

type GetPcapRequestPathSelectorIntervalEndTime struct{}

func (GetPcapRequestPathSelectorIntervalEndTime) FieldPath() *GetPcapRequest_FieldSubPath {
	return &GetPcapRequest_FieldSubPath{
		selector: GetPcapRequest_FieldPathSelectorInterval,
		subPath:  common.NewTimeIntervalFieldPathBuilder().EndTime().FieldPath(),
	}
}

func (s GetPcapRequestPathSelectorIntervalEndTime) WithValue(value *timestamp.Timestamp) *GetPcapRequest_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapRequest_FieldSubPathValue)
}

func (s GetPcapRequestPathSelectorIntervalEndTime) WithArrayOfValues(values []*timestamp.Timestamp) *GetPcapRequest_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapRequest_FieldSubPathArrayOfValues)
}

type GetPcapRequestPathSelectorIntervalStartTime struct{}

func (GetPcapRequestPathSelectorIntervalStartTime) FieldPath() *GetPcapRequest_FieldSubPath {
	return &GetPcapRequest_FieldSubPath{
		selector: GetPcapRequest_FieldPathSelectorInterval,
		subPath:  common.NewTimeIntervalFieldPathBuilder().StartTime().FieldPath(),
	}
}

func (s GetPcapRequestPathSelectorIntervalStartTime) WithValue(value *timestamp.Timestamp) *GetPcapRequest_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapRequest_FieldSubPathValue)
}

func (s GetPcapRequestPathSelectorIntervalStartTime) WithArrayOfValues(values []*timestamp.Timestamp) *GetPcapRequest_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapRequest_FieldSubPathArrayOfValues)
}

type GetPcapRequestPathSelectorPageSize struct{}

func (GetPcapRequestPathSelectorPageSize) FieldPath() *GetPcapRequest_FieldTerminalPath {
	return &GetPcapRequest_FieldTerminalPath{selector: GetPcapRequest_FieldPathSelectorPageSize}
}

func (s GetPcapRequestPathSelectorPageSize) WithValue(value int32) *GetPcapRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapRequest_FieldTerminalPathValue)
}

func (s GetPcapRequestPathSelectorPageSize) WithArrayOfValues(values []int32) *GetPcapRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapRequest_FieldTerminalPathArrayOfValues)
}

type GetPcapRequestPathSelectorPageToken struct{}

func (GetPcapRequestPathSelectorPageToken) FieldPath() *GetPcapRequest_FieldTerminalPath {
	return &GetPcapRequest_FieldTerminalPath{selector: GetPcapRequest_FieldPathSelectorPageToken}
}

func (s GetPcapRequestPathSelectorPageToken) WithValue(value string) *GetPcapRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapRequest_FieldTerminalPathValue)
}

func (s GetPcapRequestPathSelectorPageToken) WithArrayOfValues(values []string) *GetPcapRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapRequest_FieldTerminalPathArrayOfValues)
}

type GetPcapResponseFieldPathBuilder struct{}

func NewGetPcapResponseFieldPathBuilder() GetPcapResponseFieldPathBuilder {
	return GetPcapResponseFieldPathBuilder{}
}
func (GetPcapResponseFieldPathBuilder) PcapBytes() GetPcapResponsePathSelectorPcapBytes {
	return GetPcapResponsePathSelectorPcapBytes{}
}

type GetPcapResponsePathSelectorPcapBytes struct{}

func (GetPcapResponsePathSelectorPcapBytes) FieldPath() *GetPcapResponse_FieldTerminalPath {
	return &GetPcapResponse_FieldTerminalPath{selector: GetPcapResponse_FieldPathSelectorPcapBytes}
}

func (s GetPcapResponsePathSelectorPcapBytes) WithValue(value []byte) *GetPcapResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapResponse_FieldTerminalPathValue)
}

func (s GetPcapResponsePathSelectorPcapBytes) WithArrayOfValues(values [][]byte) *GetPcapResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapResponse_FieldTerminalPathArrayOfValues)
}

type GetPcapFileFromAgentRequestFieldPathBuilder struct{}

func NewGetPcapFileFromAgentRequestFieldPathBuilder() GetPcapFileFromAgentRequestFieldPathBuilder {
	return GetPcapFileFromAgentRequestFieldPathBuilder{}
}
func (GetPcapFileFromAgentRequestFieldPathBuilder) Name() GetPcapFileFromAgentRequestPathSelectorName {
	return GetPcapFileFromAgentRequestPathSelectorName{}
}

type GetPcapFileFromAgentRequestPathSelectorName struct{}

func (GetPcapFileFromAgentRequestPathSelectorName) FieldPath() *GetPcapFileFromAgentRequest_FieldTerminalPath {
	return &GetPcapFileFromAgentRequest_FieldTerminalPath{selector: GetPcapFileFromAgentRequest_FieldPathSelectorName}
}

func (s GetPcapFileFromAgentRequestPathSelectorName) WithValue(value *probe.Reference) *GetPcapFileFromAgentRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapFileFromAgentRequest_FieldTerminalPathValue)
}

func (s GetPcapFileFromAgentRequestPathSelectorName) WithArrayOfValues(values []*probe.Reference) *GetPcapFileFromAgentRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapFileFromAgentRequest_FieldTerminalPathArrayOfValues)
}

type GetPcapFileInfoFromAgentRequestFieldPathBuilder struct{}

func NewGetPcapFileInfoFromAgentRequestFieldPathBuilder() GetPcapFileInfoFromAgentRequestFieldPathBuilder {
	return GetPcapFileInfoFromAgentRequestFieldPathBuilder{}
}
func (GetPcapFileInfoFromAgentRequestFieldPathBuilder) Name() GetPcapFileInfoFromAgentRequestPathSelectorName {
	return GetPcapFileInfoFromAgentRequestPathSelectorName{}
}

type GetPcapFileInfoFromAgentRequestPathSelectorName struct{}

func (GetPcapFileInfoFromAgentRequestPathSelectorName) FieldPath() *GetPcapFileInfoFromAgentRequest_FieldTerminalPath {
	return &GetPcapFileInfoFromAgentRequest_FieldTerminalPath{selector: GetPcapFileInfoFromAgentRequest_FieldPathSelectorName}
}

func (s GetPcapFileInfoFromAgentRequestPathSelectorName) WithValue(value *probe.Reference) *GetPcapFileInfoFromAgentRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapFileInfoFromAgentRequest_FieldTerminalPathValue)
}

func (s GetPcapFileInfoFromAgentRequestPathSelectorName) WithArrayOfValues(values []*probe.Reference) *GetPcapFileInfoFromAgentRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapFileInfoFromAgentRequest_FieldTerminalPathArrayOfValues)
}

type GetPcapFileInfoFromAgentResponseFieldPathBuilder struct{}

func NewGetPcapFileInfoFromAgentResponseFieldPathBuilder() GetPcapFileInfoFromAgentResponseFieldPathBuilder {
	return GetPcapFileInfoFromAgentResponseFieldPathBuilder{}
}
func (GetPcapFileInfoFromAgentResponseFieldPathBuilder) StartTime() GetPcapFileInfoFromAgentResponsePathSelectorStartTime {
	return GetPcapFileInfoFromAgentResponsePathSelectorStartTime{}
}
func (GetPcapFileInfoFromAgentResponseFieldPathBuilder) EndTime() GetPcapFileInfoFromAgentResponsePathSelectorEndTime {
	return GetPcapFileInfoFromAgentResponsePathSelectorEndTime{}
}
func (GetPcapFileInfoFromAgentResponseFieldPathBuilder) SizeBytes() GetPcapFileInfoFromAgentResponsePathSelectorSizeBytes {
	return GetPcapFileInfoFromAgentResponsePathSelectorSizeBytes{}
}

type GetPcapFileInfoFromAgentResponsePathSelectorStartTime struct{}

func (GetPcapFileInfoFromAgentResponsePathSelectorStartTime) FieldPath() *GetPcapFileInfoFromAgentResponse_FieldTerminalPath {
	return &GetPcapFileInfoFromAgentResponse_FieldTerminalPath{selector: GetPcapFileInfoFromAgentResponse_FieldPathSelectorStartTime}
}

func (s GetPcapFileInfoFromAgentResponsePathSelectorStartTime) WithValue(value *timestamp.Timestamp) *GetPcapFileInfoFromAgentResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapFileInfoFromAgentResponse_FieldTerminalPathValue)
}

func (s GetPcapFileInfoFromAgentResponsePathSelectorStartTime) WithArrayOfValues(values []*timestamp.Timestamp) *GetPcapFileInfoFromAgentResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapFileInfoFromAgentResponse_FieldTerminalPathArrayOfValues)
}

type GetPcapFileInfoFromAgentResponsePathSelectorEndTime struct{}

func (GetPcapFileInfoFromAgentResponsePathSelectorEndTime) FieldPath() *GetPcapFileInfoFromAgentResponse_FieldTerminalPath {
	return &GetPcapFileInfoFromAgentResponse_FieldTerminalPath{selector: GetPcapFileInfoFromAgentResponse_FieldPathSelectorEndTime}
}

func (s GetPcapFileInfoFromAgentResponsePathSelectorEndTime) WithValue(value *timestamp.Timestamp) *GetPcapFileInfoFromAgentResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapFileInfoFromAgentResponse_FieldTerminalPathValue)
}

func (s GetPcapFileInfoFromAgentResponsePathSelectorEndTime) WithArrayOfValues(values []*timestamp.Timestamp) *GetPcapFileInfoFromAgentResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapFileInfoFromAgentResponse_FieldTerminalPathArrayOfValues)
}

type GetPcapFileInfoFromAgentResponsePathSelectorSizeBytes struct{}

func (GetPcapFileInfoFromAgentResponsePathSelectorSizeBytes) FieldPath() *GetPcapFileInfoFromAgentResponse_FieldTerminalPath {
	return &GetPcapFileInfoFromAgentResponse_FieldTerminalPath{selector: GetPcapFileInfoFromAgentResponse_FieldPathSelectorSizeBytes}
}

func (s GetPcapFileInfoFromAgentResponsePathSelectorSizeBytes) WithValue(value int64) *GetPcapFileInfoFromAgentResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*GetPcapFileInfoFromAgentResponse_FieldTerminalPathValue)
}

func (s GetPcapFileInfoFromAgentResponsePathSelectorSizeBytes) WithArrayOfValues(values []int64) *GetPcapFileInfoFromAgentResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*GetPcapFileInfoFromAgentResponse_FieldTerminalPathArrayOfValues)
}

type CheckPcapIsRunningRequestFieldPathBuilder struct{}

func NewCheckPcapIsRunningRequestFieldPathBuilder() CheckPcapIsRunningRequestFieldPathBuilder {
	return CheckPcapIsRunningRequestFieldPathBuilder{}
}
func (CheckPcapIsRunningRequestFieldPathBuilder) Name() CheckPcapIsRunningRequestPathSelectorName {
	return CheckPcapIsRunningRequestPathSelectorName{}
}

type CheckPcapIsRunningRequestPathSelectorName struct{}

func (CheckPcapIsRunningRequestPathSelectorName) FieldPath() *CheckPcapIsRunningRequest_FieldTerminalPath {
	return &CheckPcapIsRunningRequest_FieldTerminalPath{selector: CheckPcapIsRunningRequest_FieldPathSelectorName}
}

func (s CheckPcapIsRunningRequestPathSelectorName) WithValue(value *probe.Reference) *CheckPcapIsRunningRequest_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*CheckPcapIsRunningRequest_FieldTerminalPathValue)
}

func (s CheckPcapIsRunningRequestPathSelectorName) WithArrayOfValues(values []*probe.Reference) *CheckPcapIsRunningRequest_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*CheckPcapIsRunningRequest_FieldTerminalPathArrayOfValues)
}

type CheckPcapIsRunningResponseFieldPathBuilder struct{}

func NewCheckPcapIsRunningResponseFieldPathBuilder() CheckPcapIsRunningResponseFieldPathBuilder {
	return CheckPcapIsRunningResponseFieldPathBuilder{}
}
func (CheckPcapIsRunningResponseFieldPathBuilder) IsRunning() CheckPcapIsRunningResponsePathSelectorIsRunning {
	return CheckPcapIsRunningResponsePathSelectorIsRunning{}
}

type CheckPcapIsRunningResponsePathSelectorIsRunning struct{}

func (CheckPcapIsRunningResponsePathSelectorIsRunning) FieldPath() *CheckPcapIsRunningResponse_FieldTerminalPath {
	return &CheckPcapIsRunningResponse_FieldTerminalPath{selector: CheckPcapIsRunningResponse_FieldPathSelectorIsRunning}
}

func (s CheckPcapIsRunningResponsePathSelectorIsRunning) WithValue(value bool) *CheckPcapIsRunningResponse_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*CheckPcapIsRunningResponse_FieldTerminalPathValue)
}

func (s CheckPcapIsRunningResponsePathSelectorIsRunning) WithArrayOfValues(values []bool) *CheckPcapIsRunningResponse_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*CheckPcapIsRunningResponse_FieldTerminalPathArrayOfValues)
}
