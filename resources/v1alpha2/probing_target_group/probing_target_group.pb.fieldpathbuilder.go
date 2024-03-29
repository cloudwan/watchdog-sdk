// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/probing_target_group.proto
// DO NOT EDIT!!!

package probing_target_group

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	multi_region_policy "github.com/cloudwan/edgelq-sdk/common/types/multi_region_policy"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &multi_region_policy.MultiRegionPolicy{}
	_ = &timestamp.Timestamp{}
	_ = &project.Project{}
)

type ProbingTargetGroupFieldPathBuilder struct{}

func NewProbingTargetGroupFieldPathBuilder() ProbingTargetGroupFieldPathBuilder {
	return ProbingTargetGroupFieldPathBuilder{}
}
func (ProbingTargetGroupFieldPathBuilder) Name() ProbingTargetGroupPathSelectorName {
	return ProbingTargetGroupPathSelectorName{}
}
func (ProbingTargetGroupFieldPathBuilder) DisplayName() ProbingTargetGroupPathSelectorDisplayName {
	return ProbingTargetGroupPathSelectorDisplayName{}
}
func (ProbingTargetGroupFieldPathBuilder) State() ProbingTargetGroupPathSelectorState {
	return ProbingTargetGroupPathSelectorState{}
}
func (ProbingTargetGroupFieldPathBuilder) Metadata() ProbingTargetGroupPathSelectorMetadata {
	return ProbingTargetGroupPathSelectorMetadata{}
}

type ProbingTargetGroupPathSelectorName struct{}

func (ProbingTargetGroupPathSelectorName) FieldPath() *ProbingTargetGroup_FieldTerminalPath {
	return &ProbingTargetGroup_FieldTerminalPath{selector: ProbingTargetGroup_FieldPathSelectorName}
}

func (s ProbingTargetGroupPathSelectorName) WithValue(value *Name) *ProbingTargetGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldTerminalPathValue)
}

func (s ProbingTargetGroupPathSelectorName) WithArrayOfValues(values []*Name) *ProbingTargetGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldTerminalPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorDisplayName struct{}

func (ProbingTargetGroupPathSelectorDisplayName) FieldPath() *ProbingTargetGroup_FieldTerminalPath {
	return &ProbingTargetGroup_FieldTerminalPath{selector: ProbingTargetGroup_FieldPathSelectorDisplayName}
}

func (s ProbingTargetGroupPathSelectorDisplayName) WithValue(value string) *ProbingTargetGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldTerminalPathValue)
}

func (s ProbingTargetGroupPathSelectorDisplayName) WithArrayOfValues(values []string) *ProbingTargetGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldTerminalPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorState struct{}

func (ProbingTargetGroupPathSelectorState) FieldPath() *ProbingTargetGroup_FieldTerminalPath {
	return &ProbingTargetGroup_FieldTerminalPath{selector: ProbingTargetGroup_FieldPathSelectorState}
}

func (s ProbingTargetGroupPathSelectorState) WithValue(value *ProbingTargetGroup_State) *ProbingTargetGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldTerminalPathValue)
}

func (s ProbingTargetGroupPathSelectorState) WithArrayOfValues(values []*ProbingTargetGroup_State) *ProbingTargetGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldTerminalPathArrayOfValues)
}

func (ProbingTargetGroupPathSelectorState) WithSubPath(subPath ProbingTargetGroupState_FieldPath) *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{selector: ProbingTargetGroup_FieldPathSelectorState, subPath: subPath}
}

func (s ProbingTargetGroupPathSelectorState) WithSubValue(subPathValue ProbingTargetGroupState_FieldPathValue) *ProbingTargetGroup_FieldSubPathValue {
	return &ProbingTargetGroup_FieldSubPathValue{ProbingTargetGroup_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ProbingTargetGroupPathSelectorState) WithSubArrayOfValues(subPathArrayOfValues ProbingTargetGroupState_FieldPathArrayOfValues) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return &ProbingTargetGroup_FieldSubPathArrayOfValues{ProbingTargetGroup_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ProbingTargetGroupPathSelectorState) WithSubArrayItemValue(subPathArrayItemValue ProbingTargetGroupState_FieldPathArrayItemValue) *ProbingTargetGroup_FieldSubPathArrayItemValue {
	return &ProbingTargetGroup_FieldSubPathArrayItemValue{ProbingTargetGroup_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (ProbingTargetGroupPathSelectorState) TargetCount() ProbingTargetGroupPathSelectorStateTargetCount {
	return ProbingTargetGroupPathSelectorStateTargetCount{}
}

type ProbingTargetGroupPathSelectorStateTargetCount struct{}

func (ProbingTargetGroupPathSelectorStateTargetCount) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorState,
		subPath:  NewProbingTargetGroupStateFieldPathBuilder().TargetCount().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorStateTargetCount) WithValue(value int64) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorStateTargetCount) WithArrayOfValues(values []int64) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadata struct{}

func (ProbingTargetGroupPathSelectorMetadata) FieldPath() *ProbingTargetGroup_FieldTerminalPath {
	return &ProbingTargetGroup_FieldTerminalPath{selector: ProbingTargetGroup_FieldPathSelectorMetadata}
}

func (s ProbingTargetGroupPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *ProbingTargetGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldTerminalPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *ProbingTargetGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldTerminalPathArrayOfValues)
}

func (ProbingTargetGroupPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{selector: ProbingTargetGroup_FieldPathSelectorMetadata, subPath: subPath}
}

func (s ProbingTargetGroupPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *ProbingTargetGroup_FieldSubPathValue {
	return &ProbingTargetGroup_FieldSubPathValue{ProbingTargetGroup_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ProbingTargetGroupPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return &ProbingTargetGroup_FieldSubPathArrayOfValues{ProbingTargetGroup_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ProbingTargetGroupPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *ProbingTargetGroup_FieldSubPathArrayItemValue {
	return &ProbingTargetGroup_FieldSubPathArrayItemValue{ProbingTargetGroup_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (ProbingTargetGroupPathSelectorMetadata) CreateTime() ProbingTargetGroupPathSelectorMetadataCreateTime {
	return ProbingTargetGroupPathSelectorMetadataCreateTime{}
}

func (ProbingTargetGroupPathSelectorMetadata) UpdateTime() ProbingTargetGroupPathSelectorMetadataUpdateTime {
	return ProbingTargetGroupPathSelectorMetadataUpdateTime{}
}

func (ProbingTargetGroupPathSelectorMetadata) DeleteTime() ProbingTargetGroupPathSelectorMetadataDeleteTime {
	return ProbingTargetGroupPathSelectorMetadataDeleteTime{}
}

func (ProbingTargetGroupPathSelectorMetadata) Uuid() ProbingTargetGroupPathSelectorMetadataUuid {
	return ProbingTargetGroupPathSelectorMetadataUuid{}
}

func (ProbingTargetGroupPathSelectorMetadata) Tags() ProbingTargetGroupPathSelectorMetadataTags {
	return ProbingTargetGroupPathSelectorMetadataTags{}
}

func (ProbingTargetGroupPathSelectorMetadata) Labels() ProbingTargetGroupPathSelectorMetadataLabels {
	return ProbingTargetGroupPathSelectorMetadataLabels{}
}

func (ProbingTargetGroupPathSelectorMetadata) Annotations() ProbingTargetGroupPathSelectorMetadataAnnotations {
	return ProbingTargetGroupPathSelectorMetadataAnnotations{}
}

func (ProbingTargetGroupPathSelectorMetadata) Generation() ProbingTargetGroupPathSelectorMetadataGeneration {
	return ProbingTargetGroupPathSelectorMetadataGeneration{}
}

func (ProbingTargetGroupPathSelectorMetadata) ResourceVersion() ProbingTargetGroupPathSelectorMetadataResourceVersion {
	return ProbingTargetGroupPathSelectorMetadataResourceVersion{}
}

func (ProbingTargetGroupPathSelectorMetadata) OwnerReferences() ProbingTargetGroupPathSelectorMetadataOwnerReferences {
	return ProbingTargetGroupPathSelectorMetadataOwnerReferences{}
}

func (ProbingTargetGroupPathSelectorMetadata) Shards() ProbingTargetGroupPathSelectorMetadataShards {
	return ProbingTargetGroupPathSelectorMetadataShards{}
}

func (ProbingTargetGroupPathSelectorMetadata) Syncing() ProbingTargetGroupPathSelectorMetadataSyncing {
	return ProbingTargetGroupPathSelectorMetadataSyncing{}
}

func (ProbingTargetGroupPathSelectorMetadata) Lifecycle() ProbingTargetGroupPathSelectorMetadataLifecycle {
	return ProbingTargetGroupPathSelectorMetadataLifecycle{}
}

type ProbingTargetGroupPathSelectorMetadataCreateTime struct{}

func (ProbingTargetGroupPathSelectorMetadataCreateTime) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataUpdateTime struct{}

func (ProbingTargetGroupPathSelectorMetadataUpdateTime) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataDeleteTime struct{}

func (ProbingTargetGroupPathSelectorMetadataDeleteTime) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataDeleteTime) WithValue(value *timestamp.Timestamp) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamp.Timestamp) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataUuid struct{}

func (ProbingTargetGroupPathSelectorMetadataUuid) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataUuid) WithValue(value string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataUuid) WithArrayOfValues(values []string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataTags struct{}

func (ProbingTargetGroupPathSelectorMetadataTags) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataTags) WithValue(value []string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

func (s ProbingTargetGroupPathSelectorMetadataTags) WithItemValue(value string) *ProbingTargetGroup_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ProbingTargetGroup_FieldSubPathArrayItemValue)
}

type ProbingTargetGroupPathSelectorMetadataLabels struct{}

func (ProbingTargetGroupPathSelectorMetadataLabels) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataLabels) WithValue(value map[string]string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

func (ProbingTargetGroupPathSelectorMetadataLabels) WithKey(key string) ProbingTargetGroupMapPathSelectorMetadataLabels {
	return ProbingTargetGroupMapPathSelectorMetadataLabels{key: key}
}

type ProbingTargetGroupMapPathSelectorMetadataLabels struct {
	key string
}

func (s ProbingTargetGroupMapPathSelectorMetadataLabels) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s ProbingTargetGroupMapPathSelectorMetadataLabels) WithValue(value string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataAnnotations struct{}

func (ProbingTargetGroupPathSelectorMetadataAnnotations) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataAnnotations) WithValue(value map[string]string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

func (ProbingTargetGroupPathSelectorMetadataAnnotations) WithKey(key string) ProbingTargetGroupMapPathSelectorMetadataAnnotations {
	return ProbingTargetGroupMapPathSelectorMetadataAnnotations{key: key}
}

type ProbingTargetGroupMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s ProbingTargetGroupMapPathSelectorMetadataAnnotations) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s ProbingTargetGroupMapPathSelectorMetadataAnnotations) WithValue(value string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataGeneration struct{}

func (ProbingTargetGroupPathSelectorMetadataGeneration) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataGeneration) WithValue(value int64) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataResourceVersion struct{}

func (ProbingTargetGroupPathSelectorMetadataResourceVersion) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataResourceVersion) WithValue(value string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataOwnerReferences struct{}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferences) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *ProbingTargetGroup_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ProbingTargetGroup_FieldSubPathArrayItemValue)
}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferences) Kind() ProbingTargetGroupPathSelectorMetadataOwnerReferencesKind {
	return ProbingTargetGroupPathSelectorMetadataOwnerReferencesKind{}
}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferences) Version() ProbingTargetGroupPathSelectorMetadataOwnerReferencesVersion {
	return ProbingTargetGroupPathSelectorMetadataOwnerReferencesVersion{}
}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferences) Name() ProbingTargetGroupPathSelectorMetadataOwnerReferencesName {
	return ProbingTargetGroupPathSelectorMetadataOwnerReferencesName{}
}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferences) Region() ProbingTargetGroupPathSelectorMetadataOwnerReferencesRegion {
	return ProbingTargetGroupPathSelectorMetadataOwnerReferencesRegion{}
}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferences) Controller() ProbingTargetGroupPathSelectorMetadataOwnerReferencesController {
	return ProbingTargetGroupPathSelectorMetadataOwnerReferencesController{}
}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() ProbingTargetGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return ProbingTargetGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferences) RequiresOwnerReference() ProbingTargetGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return ProbingTargetGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type ProbingTargetGroupPathSelectorMetadataOwnerReferencesKind struct{}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferencesKind) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataOwnerReferencesVersion struct{}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferencesVersion) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataOwnerReferencesName struct{}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferencesName) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesName) WithValue(value string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataOwnerReferencesRegion struct{}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferencesRegion) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataOwnerReferencesController struct{}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferencesController) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (ProbingTargetGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataShards struct{}

func (ProbingTargetGroupPathSelectorMetadataShards) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataShards) WithValue(value map[string]int64) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

func (ProbingTargetGroupPathSelectorMetadataShards) WithKey(key string) ProbingTargetGroupMapPathSelectorMetadataShards {
	return ProbingTargetGroupMapPathSelectorMetadataShards{key: key}
}

type ProbingTargetGroupMapPathSelectorMetadataShards struct {
	key string
}

func (s ProbingTargetGroupMapPathSelectorMetadataShards) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s ProbingTargetGroupMapPathSelectorMetadataShards) WithValue(value int64) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataSyncing struct{}

func (ProbingTargetGroupPathSelectorMetadataSyncing) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataSyncing) WithValue(value *ntt_meta.SyncingMeta) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataSyncing) WithArrayOfValues(values []*ntt_meta.SyncingMeta) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

func (ProbingTargetGroupPathSelectorMetadataSyncing) OwningRegion() ProbingTargetGroupPathSelectorMetadataSyncingOwningRegion {
	return ProbingTargetGroupPathSelectorMetadataSyncingOwningRegion{}
}

func (ProbingTargetGroupPathSelectorMetadataSyncing) Regions() ProbingTargetGroupPathSelectorMetadataSyncingRegions {
	return ProbingTargetGroupPathSelectorMetadataSyncingRegions{}
}

type ProbingTargetGroupPathSelectorMetadataSyncingOwningRegion struct{}

func (ProbingTargetGroupPathSelectorMetadataSyncingOwningRegion) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataSyncingRegions struct{}

func (ProbingTargetGroupPathSelectorMetadataSyncingRegions) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataSyncingRegions) WithValue(value []string) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

func (s ProbingTargetGroupPathSelectorMetadataSyncingRegions) WithItemValue(value string) *ProbingTargetGroup_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ProbingTargetGroup_FieldSubPathArrayItemValue)
}

type ProbingTargetGroupPathSelectorMetadataLifecycle struct{}

func (ProbingTargetGroupPathSelectorMetadataLifecycle) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataLifecycle) WithValue(value *ntt_meta.Lifecycle) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataLifecycle) WithArrayOfValues(values []*ntt_meta.Lifecycle) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

func (ProbingTargetGroupPathSelectorMetadataLifecycle) State() ProbingTargetGroupPathSelectorMetadataLifecycleState {
	return ProbingTargetGroupPathSelectorMetadataLifecycleState{}
}

func (ProbingTargetGroupPathSelectorMetadataLifecycle) BlockDeletion() ProbingTargetGroupPathSelectorMetadataLifecycleBlockDeletion {
	return ProbingTargetGroupPathSelectorMetadataLifecycleBlockDeletion{}
}

type ProbingTargetGroupPathSelectorMetadataLifecycleState struct{}

func (ProbingTargetGroupPathSelectorMetadataLifecycleState) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataLifecycleState) WithValue(value ntt_meta.Lifecycle_State) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataLifecycleState) WithArrayOfValues(values []ntt_meta.Lifecycle_State) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupPathSelectorMetadataLifecycleBlockDeletion struct{}

func (ProbingTargetGroupPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *ProbingTargetGroup_FieldSubPath {
	return &ProbingTargetGroup_FieldSubPath{
		selector: ProbingTargetGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s ProbingTargetGroupPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *ProbingTargetGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroup_FieldSubPathValue)
}

func (s ProbingTargetGroupPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *ProbingTargetGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroup_FieldSubPathArrayOfValues)
}

type ProbingTargetGroupStateFieldPathBuilder struct{}

func NewProbingTargetGroupStateFieldPathBuilder() ProbingTargetGroupStateFieldPathBuilder {
	return ProbingTargetGroupStateFieldPathBuilder{}
}
func (ProbingTargetGroupStateFieldPathBuilder) TargetCount() ProbingTargetGroup_StatePathSelectorTargetCount {
	return ProbingTargetGroup_StatePathSelectorTargetCount{}
}

type ProbingTargetGroup_StatePathSelectorTargetCount struct{}

func (ProbingTargetGroup_StatePathSelectorTargetCount) FieldPath() *ProbingTargetGroupState_FieldTerminalPath {
	return &ProbingTargetGroupState_FieldTerminalPath{selector: ProbingTargetGroupState_FieldPathSelectorTargetCount}
}

func (s ProbingTargetGroup_StatePathSelectorTargetCount) WithValue(value int64) *ProbingTargetGroupState_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbingTargetGroupState_FieldTerminalPathValue)
}

func (s ProbingTargetGroup_StatePathSelectorTargetCount) WithArrayOfValues(values []int64) *ProbingTargetGroupState_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbingTargetGroupState_FieldTerminalPathArrayOfValues)
}
