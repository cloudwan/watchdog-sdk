// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/internet_quality_rating.proto
// DO NOT EDIT!!!

package internet_quality_rating

// proto imports
import (
	ntt_memo "github.com/cloudwan/edgelq-sdk/common/types/memo"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
	devices_device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device"
	devices_project "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/project"
	iam_attestation_domain "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/attestation_domain"
	iam_iam_common "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/common"
	iam_condition "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/condition"
	iam_organization "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/organization"
	iam_permission "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/permission"
	iam_project "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/project"
	iam_role "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/role"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	iam_user "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/user"
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
	_ = &iam_iam_common.Actor{}
	_ = &iam_condition.Condition{}
	_ = &iam_organization.Organization{}
	_ = &iam_permission.Permission{}
	_ = &iam_project.Project{}
	_ = &iam_role.Role{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &iam_user.User{}
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

type InternetQualityRatingFieldPathBuilder struct{}

func NewInternetQualityRatingFieldPathBuilder() InternetQualityRatingFieldPathBuilder {
	return InternetQualityRatingFieldPathBuilder{}
}
func (InternetQualityRatingFieldPathBuilder) Name() InternetQualityRatingPathSelectorName {
	return InternetQualityRatingPathSelectorName{}
}
func (InternetQualityRatingFieldPathBuilder) Rating() InternetQualityRatingPathSelectorRating {
	return InternetQualityRatingPathSelectorRating{}
}
func (InternetQualityRatingFieldPathBuilder) Timestamp() InternetQualityRatingPathSelectorTimestamp {
	return InternetQualityRatingPathSelectorTimestamp{}
}
func (InternetQualityRatingFieldPathBuilder) Metadata() InternetQualityRatingPathSelectorMetadata {
	return InternetQualityRatingPathSelectorMetadata{}
}

type InternetQualityRatingPathSelectorName struct{}

func (InternetQualityRatingPathSelectorName) FieldPath() *InternetQualityRating_FieldTerminalPath {
	return &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorName}
}

func (s InternetQualityRatingPathSelectorName) WithValue(value *Name) *InternetQualityRating_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldTerminalPathValue)
}

func (s InternetQualityRatingPathSelectorName) WithArrayOfValues(values []*Name) *InternetQualityRating_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldTerminalPathArrayOfValues)
}

type InternetQualityRatingPathSelectorRating struct{}

func (InternetQualityRatingPathSelectorRating) FieldPath() *InternetQualityRating_FieldTerminalPath {
	return &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorRating}
}

func (s InternetQualityRatingPathSelectorRating) WithValue(value InternetQualityRating_Rating) *InternetQualityRating_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldTerminalPathValue)
}

func (s InternetQualityRatingPathSelectorRating) WithArrayOfValues(values []InternetQualityRating_Rating) *InternetQualityRating_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldTerminalPathArrayOfValues)
}

type InternetQualityRatingPathSelectorTimestamp struct{}

func (InternetQualityRatingPathSelectorTimestamp) FieldPath() *InternetQualityRating_FieldTerminalPath {
	return &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorTimestamp}
}

func (s InternetQualityRatingPathSelectorTimestamp) WithValue(value *timestamp.Timestamp) *InternetQualityRating_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldTerminalPathValue)
}

func (s InternetQualityRatingPathSelectorTimestamp) WithArrayOfValues(values []*timestamp.Timestamp) *InternetQualityRating_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldTerminalPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadata struct{}

func (InternetQualityRatingPathSelectorMetadata) FieldPath() *InternetQualityRating_FieldTerminalPath {
	return &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorMetadata}
}

func (s InternetQualityRatingPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *InternetQualityRating_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldTerminalPathValue)
}

func (s InternetQualityRatingPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *InternetQualityRating_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldTerminalPathArrayOfValues)
}

func (InternetQualityRatingPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{selector: InternetQualityRating_FieldPathSelectorMetadata, subPath: subPath}
}

func (s InternetQualityRatingPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *InternetQualityRating_FieldSubPathValue {
	return &InternetQualityRating_FieldSubPathValue{InternetQualityRating_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s InternetQualityRatingPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *InternetQualityRating_FieldSubPathArrayOfValues {
	return &InternetQualityRating_FieldSubPathArrayOfValues{InternetQualityRating_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s InternetQualityRatingPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *InternetQualityRating_FieldSubPathArrayItemValue {
	return &InternetQualityRating_FieldSubPathArrayItemValue{InternetQualityRating_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (InternetQualityRatingPathSelectorMetadata) CreateTime() InternetQualityRatingPathSelectorMetadataCreateTime {
	return InternetQualityRatingPathSelectorMetadataCreateTime{}
}

func (InternetQualityRatingPathSelectorMetadata) UpdateTime() InternetQualityRatingPathSelectorMetadataUpdateTime {
	return InternetQualityRatingPathSelectorMetadataUpdateTime{}
}

func (InternetQualityRatingPathSelectorMetadata) Uuid() InternetQualityRatingPathSelectorMetadataUuid {
	return InternetQualityRatingPathSelectorMetadataUuid{}
}

func (InternetQualityRatingPathSelectorMetadata) Tags() InternetQualityRatingPathSelectorMetadataTags {
	return InternetQualityRatingPathSelectorMetadataTags{}
}

func (InternetQualityRatingPathSelectorMetadata) Labels() InternetQualityRatingPathSelectorMetadataLabels {
	return InternetQualityRatingPathSelectorMetadataLabels{}
}

func (InternetQualityRatingPathSelectorMetadata) Annotations() InternetQualityRatingPathSelectorMetadataAnnotations {
	return InternetQualityRatingPathSelectorMetadataAnnotations{}
}

func (InternetQualityRatingPathSelectorMetadata) Generation() InternetQualityRatingPathSelectorMetadataGeneration {
	return InternetQualityRatingPathSelectorMetadataGeneration{}
}

func (InternetQualityRatingPathSelectorMetadata) ResourceVersion() InternetQualityRatingPathSelectorMetadataResourceVersion {
	return InternetQualityRatingPathSelectorMetadataResourceVersion{}
}

func (InternetQualityRatingPathSelectorMetadata) OwnerReferences() InternetQualityRatingPathSelectorMetadataOwnerReferences {
	return InternetQualityRatingPathSelectorMetadataOwnerReferences{}
}

func (InternetQualityRatingPathSelectorMetadata) Shards() InternetQualityRatingPathSelectorMetadataShards {
	return InternetQualityRatingPathSelectorMetadataShards{}
}

func (InternetQualityRatingPathSelectorMetadata) Syncing() InternetQualityRatingPathSelectorMetadataSyncing {
	return InternetQualityRatingPathSelectorMetadataSyncing{}
}

type InternetQualityRatingPathSelectorMetadataCreateTime struct{}

func (InternetQualityRatingPathSelectorMetadataCreateTime) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataUpdateTime struct{}

func (InternetQualityRatingPathSelectorMetadataUpdateTime) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataUuid struct{}

func (InternetQualityRatingPathSelectorMetadataUuid) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataUuid) WithValue(value string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataUuid) WithArrayOfValues(values []string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataTags struct{}

func (InternetQualityRatingPathSelectorMetadataTags) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataTags) WithValue(value []string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

func (s InternetQualityRatingPathSelectorMetadataTags) WithItemValue(value string) *InternetQualityRating_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*InternetQualityRating_FieldSubPathArrayItemValue)
}

type InternetQualityRatingPathSelectorMetadataLabels struct{}

func (InternetQualityRatingPathSelectorMetadataLabels) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataLabels) WithValue(value map[string]string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

func (InternetQualityRatingPathSelectorMetadataLabels) WithKey(key string) InternetQualityRatingMapPathSelectorMetadataLabels {
	return InternetQualityRatingMapPathSelectorMetadataLabels{key: key}
}

type InternetQualityRatingMapPathSelectorMetadataLabels struct {
	key string
}

func (s InternetQualityRatingMapPathSelectorMetadataLabels) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s InternetQualityRatingMapPathSelectorMetadataLabels) WithValue(value string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataAnnotations struct{}

func (InternetQualityRatingPathSelectorMetadataAnnotations) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataAnnotations) WithValue(value map[string]string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

func (InternetQualityRatingPathSelectorMetadataAnnotations) WithKey(key string) InternetQualityRatingMapPathSelectorMetadataAnnotations {
	return InternetQualityRatingMapPathSelectorMetadataAnnotations{key: key}
}

type InternetQualityRatingMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s InternetQualityRatingMapPathSelectorMetadataAnnotations) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s InternetQualityRatingMapPathSelectorMetadataAnnotations) WithValue(value string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataGeneration struct{}

func (InternetQualityRatingPathSelectorMetadataGeneration) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataGeneration) WithValue(value int64) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataResourceVersion struct{}

func (InternetQualityRatingPathSelectorMetadataResourceVersion) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataResourceVersion) WithValue(value string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataOwnerReferences struct{}

func (InternetQualityRatingPathSelectorMetadataOwnerReferences) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *InternetQualityRating_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*InternetQualityRating_FieldSubPathArrayItemValue)
}

func (InternetQualityRatingPathSelectorMetadataOwnerReferences) ApiVersion() InternetQualityRatingPathSelectorMetadataOwnerReferencesApiVersion {
	return InternetQualityRatingPathSelectorMetadataOwnerReferencesApiVersion{}
}

func (InternetQualityRatingPathSelectorMetadataOwnerReferences) Kind() InternetQualityRatingPathSelectorMetadataOwnerReferencesKind {
	return InternetQualityRatingPathSelectorMetadataOwnerReferencesKind{}
}

func (InternetQualityRatingPathSelectorMetadataOwnerReferences) Name() InternetQualityRatingPathSelectorMetadataOwnerReferencesName {
	return InternetQualityRatingPathSelectorMetadataOwnerReferencesName{}
}

func (InternetQualityRatingPathSelectorMetadataOwnerReferences) Uid() InternetQualityRatingPathSelectorMetadataOwnerReferencesUid {
	return InternetQualityRatingPathSelectorMetadataOwnerReferencesUid{}
}

func (InternetQualityRatingPathSelectorMetadataOwnerReferences) Controller() InternetQualityRatingPathSelectorMetadataOwnerReferencesController {
	return InternetQualityRatingPathSelectorMetadataOwnerReferencesController{}
}

func (InternetQualityRatingPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() InternetQualityRatingPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return InternetQualityRatingPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

type InternetQualityRatingPathSelectorMetadataOwnerReferencesApiVersion struct{}

func (InternetQualityRatingPathSelectorMetadataOwnerReferencesApiVersion) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().ApiVersion().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesApiVersion) WithValue(value string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesApiVersion) WithArrayOfValues(values []string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataOwnerReferencesKind struct{}

func (InternetQualityRatingPathSelectorMetadataOwnerReferencesKind) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataOwnerReferencesName struct{}

func (InternetQualityRatingPathSelectorMetadataOwnerReferencesName) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesName) WithValue(value string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataOwnerReferencesUid struct{}

func (InternetQualityRatingPathSelectorMetadataOwnerReferencesUid) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Uid().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesUid) WithValue(value string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesUid) WithArrayOfValues(values []string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataOwnerReferencesController struct{}

func (InternetQualityRatingPathSelectorMetadataOwnerReferencesController) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (InternetQualityRatingPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataShards struct{}

func (InternetQualityRatingPathSelectorMetadataShards) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataShards) WithValue(value map[string]int64) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

func (InternetQualityRatingPathSelectorMetadataShards) WithKey(key string) InternetQualityRatingMapPathSelectorMetadataShards {
	return InternetQualityRatingMapPathSelectorMetadataShards{key: key}
}

type InternetQualityRatingMapPathSelectorMetadataShards struct {
	key string
}

func (s InternetQualityRatingMapPathSelectorMetadataShards) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s InternetQualityRatingMapPathSelectorMetadataShards) WithValue(value int64) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataSyncing struct{}

func (InternetQualityRatingPathSelectorMetadataSyncing) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataSyncing) WithValue(value *ntt_meta.SyncingMeta) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataSyncing) WithArrayOfValues(values []*ntt_meta.SyncingMeta) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

func (InternetQualityRatingPathSelectorMetadataSyncing) OwningRegion() InternetQualityRatingPathSelectorMetadataSyncingOwningRegion {
	return InternetQualityRatingPathSelectorMetadataSyncingOwningRegion{}
}

func (InternetQualityRatingPathSelectorMetadataSyncing) Regions() InternetQualityRatingPathSelectorMetadataSyncingRegions {
	return InternetQualityRatingPathSelectorMetadataSyncingRegions{}
}

type InternetQualityRatingPathSelectorMetadataSyncingOwningRegion struct{}

func (InternetQualityRatingPathSelectorMetadataSyncingOwningRegion) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

type InternetQualityRatingPathSelectorMetadataSyncingRegions struct{}

func (InternetQualityRatingPathSelectorMetadataSyncingRegions) FieldPath() *InternetQualityRating_FieldSubPath {
	return &InternetQualityRating_FieldSubPath{
		selector: InternetQualityRating_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s InternetQualityRatingPathSelectorMetadataSyncingRegions) WithValue(value []string) *InternetQualityRating_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*InternetQualityRating_FieldSubPathValue)
}

func (s InternetQualityRatingPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *InternetQualityRating_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*InternetQualityRating_FieldSubPathArrayOfValues)
}

func (s InternetQualityRatingPathSelectorMetadataSyncingRegions) WithItemValue(value string) *InternetQualityRating_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*InternetQualityRating_FieldSubPathArrayItemValue)
}
