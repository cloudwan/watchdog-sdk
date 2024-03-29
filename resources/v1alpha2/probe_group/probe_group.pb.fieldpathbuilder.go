// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/probe_group.proto
// DO NOT EDIT!!!

package probe_group

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

type ProbeGroupFieldPathBuilder struct{}

func NewProbeGroupFieldPathBuilder() ProbeGroupFieldPathBuilder {
	return ProbeGroupFieldPathBuilder{}
}
func (ProbeGroupFieldPathBuilder) Name() ProbeGroupPathSelectorName {
	return ProbeGroupPathSelectorName{}
}
func (ProbeGroupFieldPathBuilder) DisplayName() ProbeGroupPathSelectorDisplayName {
	return ProbeGroupPathSelectorDisplayName{}
}
func (ProbeGroupFieldPathBuilder) State() ProbeGroupPathSelectorState {
	return ProbeGroupPathSelectorState{}
}
func (ProbeGroupFieldPathBuilder) Metadata() ProbeGroupPathSelectorMetadata {
	return ProbeGroupPathSelectorMetadata{}
}

type ProbeGroupPathSelectorName struct{}

func (ProbeGroupPathSelectorName) FieldPath() *ProbeGroup_FieldTerminalPath {
	return &ProbeGroup_FieldTerminalPath{selector: ProbeGroup_FieldPathSelectorName}
}

func (s ProbeGroupPathSelectorName) WithValue(value *Name) *ProbeGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldTerminalPathValue)
}

func (s ProbeGroupPathSelectorName) WithArrayOfValues(values []*Name) *ProbeGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldTerminalPathArrayOfValues)
}

type ProbeGroupPathSelectorDisplayName struct{}

func (ProbeGroupPathSelectorDisplayName) FieldPath() *ProbeGroup_FieldTerminalPath {
	return &ProbeGroup_FieldTerminalPath{selector: ProbeGroup_FieldPathSelectorDisplayName}
}

func (s ProbeGroupPathSelectorDisplayName) WithValue(value string) *ProbeGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldTerminalPathValue)
}

func (s ProbeGroupPathSelectorDisplayName) WithArrayOfValues(values []string) *ProbeGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldTerminalPathArrayOfValues)
}

type ProbeGroupPathSelectorState struct{}

func (ProbeGroupPathSelectorState) FieldPath() *ProbeGroup_FieldTerminalPath {
	return &ProbeGroup_FieldTerminalPath{selector: ProbeGroup_FieldPathSelectorState}
}

func (s ProbeGroupPathSelectorState) WithValue(value *ProbeGroup_State) *ProbeGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldTerminalPathValue)
}

func (s ProbeGroupPathSelectorState) WithArrayOfValues(values []*ProbeGroup_State) *ProbeGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldTerminalPathArrayOfValues)
}

func (ProbeGroupPathSelectorState) WithSubPath(subPath ProbeGroupState_FieldPath) *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{selector: ProbeGroup_FieldPathSelectorState, subPath: subPath}
}

func (s ProbeGroupPathSelectorState) WithSubValue(subPathValue ProbeGroupState_FieldPathValue) *ProbeGroup_FieldSubPathValue {
	return &ProbeGroup_FieldSubPathValue{ProbeGroup_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ProbeGroupPathSelectorState) WithSubArrayOfValues(subPathArrayOfValues ProbeGroupState_FieldPathArrayOfValues) *ProbeGroup_FieldSubPathArrayOfValues {
	return &ProbeGroup_FieldSubPathArrayOfValues{ProbeGroup_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ProbeGroupPathSelectorState) WithSubArrayItemValue(subPathArrayItemValue ProbeGroupState_FieldPathArrayItemValue) *ProbeGroup_FieldSubPathArrayItemValue {
	return &ProbeGroup_FieldSubPathArrayItemValue{ProbeGroup_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (ProbeGroupPathSelectorState) TargetCount() ProbeGroupPathSelectorStateTargetCount {
	return ProbeGroupPathSelectorStateTargetCount{}
}

func (ProbeGroupPathSelectorState) RegionalTargetCounts() ProbeGroupPathSelectorStateRegionalTargetCounts {
	return ProbeGroupPathSelectorStateRegionalTargetCounts{}
}

type ProbeGroupPathSelectorStateTargetCount struct{}

func (ProbeGroupPathSelectorStateTargetCount) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorState,
		subPath:  NewProbeGroupStateFieldPathBuilder().TargetCount().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorStateTargetCount) WithValue(value int64) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorStateTargetCount) WithArrayOfValues(values []int64) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorStateRegionalTargetCounts struct{}

func (ProbeGroupPathSelectorStateRegionalTargetCounts) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorState,
		subPath:  NewProbeGroupStateFieldPathBuilder().RegionalTargetCounts().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorStateRegionalTargetCounts) WithValue(value map[string]int64) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorStateRegionalTargetCounts) WithArrayOfValues(values []map[string]int64) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

func (ProbeGroupPathSelectorStateRegionalTargetCounts) WithKey(key string) ProbeGroupMapPathSelectorStateRegionalTargetCounts {
	return ProbeGroupMapPathSelectorStateRegionalTargetCounts{key: key}
}

type ProbeGroupMapPathSelectorStateRegionalTargetCounts struct {
	key string
}

func (s ProbeGroupMapPathSelectorStateRegionalTargetCounts) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorState,
		subPath:  NewProbeGroupStateFieldPathBuilder().RegionalTargetCounts().WithKey(s.key).FieldPath(),
	}
}

func (s ProbeGroupMapPathSelectorStateRegionalTargetCounts) WithValue(value int64) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupMapPathSelectorStateRegionalTargetCounts) WithArrayOfValues(values []int64) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadata struct{}

func (ProbeGroupPathSelectorMetadata) FieldPath() *ProbeGroup_FieldTerminalPath {
	return &ProbeGroup_FieldTerminalPath{selector: ProbeGroup_FieldPathSelectorMetadata}
}

func (s ProbeGroupPathSelectorMetadata) WithValue(value *ntt_meta.Meta) *ProbeGroup_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldTerminalPathValue)
}

func (s ProbeGroupPathSelectorMetadata) WithArrayOfValues(values []*ntt_meta.Meta) *ProbeGroup_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldTerminalPathArrayOfValues)
}

func (ProbeGroupPathSelectorMetadata) WithSubPath(subPath ntt_meta.Meta_FieldPath) *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{selector: ProbeGroup_FieldPathSelectorMetadata, subPath: subPath}
}

func (s ProbeGroupPathSelectorMetadata) WithSubValue(subPathValue ntt_meta.Meta_FieldPathValue) *ProbeGroup_FieldSubPathValue {
	return &ProbeGroup_FieldSubPathValue{ProbeGroup_FieldPath: s.WithSubPath(subPathValue), subPathValue: subPathValue}
}

func (s ProbeGroupPathSelectorMetadata) WithSubArrayOfValues(subPathArrayOfValues ntt_meta.Meta_FieldPathArrayOfValues) *ProbeGroup_FieldSubPathArrayOfValues {
	return &ProbeGroup_FieldSubPathArrayOfValues{ProbeGroup_FieldPath: s.WithSubPath(subPathArrayOfValues), subPathArrayOfValues: subPathArrayOfValues}
}

func (s ProbeGroupPathSelectorMetadata) WithSubArrayItemValue(subPathArrayItemValue ntt_meta.Meta_FieldPathArrayItemValue) *ProbeGroup_FieldSubPathArrayItemValue {
	return &ProbeGroup_FieldSubPathArrayItemValue{ProbeGroup_FieldPath: s.WithSubPath(subPathArrayItemValue), subPathItemValue: subPathArrayItemValue}
}

func (ProbeGroupPathSelectorMetadata) CreateTime() ProbeGroupPathSelectorMetadataCreateTime {
	return ProbeGroupPathSelectorMetadataCreateTime{}
}

func (ProbeGroupPathSelectorMetadata) UpdateTime() ProbeGroupPathSelectorMetadataUpdateTime {
	return ProbeGroupPathSelectorMetadataUpdateTime{}
}

func (ProbeGroupPathSelectorMetadata) DeleteTime() ProbeGroupPathSelectorMetadataDeleteTime {
	return ProbeGroupPathSelectorMetadataDeleteTime{}
}

func (ProbeGroupPathSelectorMetadata) Uuid() ProbeGroupPathSelectorMetadataUuid {
	return ProbeGroupPathSelectorMetadataUuid{}
}

func (ProbeGroupPathSelectorMetadata) Tags() ProbeGroupPathSelectorMetadataTags {
	return ProbeGroupPathSelectorMetadataTags{}
}

func (ProbeGroupPathSelectorMetadata) Labels() ProbeGroupPathSelectorMetadataLabels {
	return ProbeGroupPathSelectorMetadataLabels{}
}

func (ProbeGroupPathSelectorMetadata) Annotations() ProbeGroupPathSelectorMetadataAnnotations {
	return ProbeGroupPathSelectorMetadataAnnotations{}
}

func (ProbeGroupPathSelectorMetadata) Generation() ProbeGroupPathSelectorMetadataGeneration {
	return ProbeGroupPathSelectorMetadataGeneration{}
}

func (ProbeGroupPathSelectorMetadata) ResourceVersion() ProbeGroupPathSelectorMetadataResourceVersion {
	return ProbeGroupPathSelectorMetadataResourceVersion{}
}

func (ProbeGroupPathSelectorMetadata) OwnerReferences() ProbeGroupPathSelectorMetadataOwnerReferences {
	return ProbeGroupPathSelectorMetadataOwnerReferences{}
}

func (ProbeGroupPathSelectorMetadata) Shards() ProbeGroupPathSelectorMetadataShards {
	return ProbeGroupPathSelectorMetadataShards{}
}

func (ProbeGroupPathSelectorMetadata) Syncing() ProbeGroupPathSelectorMetadataSyncing {
	return ProbeGroupPathSelectorMetadataSyncing{}
}

func (ProbeGroupPathSelectorMetadata) Lifecycle() ProbeGroupPathSelectorMetadataLifecycle {
	return ProbeGroupPathSelectorMetadataLifecycle{}
}

type ProbeGroupPathSelectorMetadataCreateTime struct{}

func (ProbeGroupPathSelectorMetadataCreateTime) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().CreateTime().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataCreateTime) WithValue(value *timestamp.Timestamp) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataCreateTime) WithArrayOfValues(values []*timestamp.Timestamp) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataUpdateTime struct{}

func (ProbeGroupPathSelectorMetadataUpdateTime) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().UpdateTime().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataUpdateTime) WithValue(value *timestamp.Timestamp) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataUpdateTime) WithArrayOfValues(values []*timestamp.Timestamp) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataDeleteTime struct{}

func (ProbeGroupPathSelectorMetadataDeleteTime) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().DeleteTime().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataDeleteTime) WithValue(value *timestamp.Timestamp) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataDeleteTime) WithArrayOfValues(values []*timestamp.Timestamp) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataUuid struct{}

func (ProbeGroupPathSelectorMetadataUuid) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Uuid().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataUuid) WithValue(value string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataUuid) WithArrayOfValues(values []string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataTags struct{}

func (ProbeGroupPathSelectorMetadataTags) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Tags().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataTags) WithValue(value []string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataTags) WithArrayOfValues(values [][]string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

func (s ProbeGroupPathSelectorMetadataTags) WithItemValue(value string) *ProbeGroup_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ProbeGroup_FieldSubPathArrayItemValue)
}

type ProbeGroupPathSelectorMetadataLabels struct{}

func (ProbeGroupPathSelectorMetadataLabels) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataLabels) WithValue(value map[string]string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataLabels) WithArrayOfValues(values []map[string]string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

func (ProbeGroupPathSelectorMetadataLabels) WithKey(key string) ProbeGroupMapPathSelectorMetadataLabels {
	return ProbeGroupMapPathSelectorMetadataLabels{key: key}
}

type ProbeGroupMapPathSelectorMetadataLabels struct {
	key string
}

func (s ProbeGroupMapPathSelectorMetadataLabels) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Labels().WithKey(s.key).FieldPath(),
	}
}

func (s ProbeGroupMapPathSelectorMetadataLabels) WithValue(value string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupMapPathSelectorMetadataLabels) WithArrayOfValues(values []string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataAnnotations struct{}

func (ProbeGroupPathSelectorMetadataAnnotations) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataAnnotations) WithValue(value map[string]string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataAnnotations) WithArrayOfValues(values []map[string]string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

func (ProbeGroupPathSelectorMetadataAnnotations) WithKey(key string) ProbeGroupMapPathSelectorMetadataAnnotations {
	return ProbeGroupMapPathSelectorMetadataAnnotations{key: key}
}

type ProbeGroupMapPathSelectorMetadataAnnotations struct {
	key string
}

func (s ProbeGroupMapPathSelectorMetadataAnnotations) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Annotations().WithKey(s.key).FieldPath(),
	}
}

func (s ProbeGroupMapPathSelectorMetadataAnnotations) WithValue(value string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupMapPathSelectorMetadataAnnotations) WithArrayOfValues(values []string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataGeneration struct{}

func (ProbeGroupPathSelectorMetadataGeneration) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Generation().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataGeneration) WithValue(value int64) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataGeneration) WithArrayOfValues(values []int64) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataResourceVersion struct{}

func (ProbeGroupPathSelectorMetadataResourceVersion) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().ResourceVersion().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataResourceVersion) WithValue(value string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataResourceVersion) WithArrayOfValues(values []string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataOwnerReferences struct{}

func (ProbeGroupPathSelectorMetadataOwnerReferences) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataOwnerReferences) WithValue(value []*ntt_meta.OwnerReference) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataOwnerReferences) WithArrayOfValues(values [][]*ntt_meta.OwnerReference) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

func (s ProbeGroupPathSelectorMetadataOwnerReferences) WithItemValue(value *ntt_meta.OwnerReference) *ProbeGroup_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ProbeGroup_FieldSubPathArrayItemValue)
}

func (ProbeGroupPathSelectorMetadataOwnerReferences) Kind() ProbeGroupPathSelectorMetadataOwnerReferencesKind {
	return ProbeGroupPathSelectorMetadataOwnerReferencesKind{}
}

func (ProbeGroupPathSelectorMetadataOwnerReferences) Version() ProbeGroupPathSelectorMetadataOwnerReferencesVersion {
	return ProbeGroupPathSelectorMetadataOwnerReferencesVersion{}
}

func (ProbeGroupPathSelectorMetadataOwnerReferences) Name() ProbeGroupPathSelectorMetadataOwnerReferencesName {
	return ProbeGroupPathSelectorMetadataOwnerReferencesName{}
}

func (ProbeGroupPathSelectorMetadataOwnerReferences) Region() ProbeGroupPathSelectorMetadataOwnerReferencesRegion {
	return ProbeGroupPathSelectorMetadataOwnerReferencesRegion{}
}

func (ProbeGroupPathSelectorMetadataOwnerReferences) Controller() ProbeGroupPathSelectorMetadataOwnerReferencesController {
	return ProbeGroupPathSelectorMetadataOwnerReferencesController{}
}

func (ProbeGroupPathSelectorMetadataOwnerReferences) BlockOwnerDeletion() ProbeGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion {
	return ProbeGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion{}
}

func (ProbeGroupPathSelectorMetadataOwnerReferences) RequiresOwnerReference() ProbeGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference {
	return ProbeGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference{}
}

type ProbeGroupPathSelectorMetadataOwnerReferencesKind struct{}

func (ProbeGroupPathSelectorMetadataOwnerReferencesKind) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Kind().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesKind) WithValue(value string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesKind) WithArrayOfValues(values []string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataOwnerReferencesVersion struct{}

func (ProbeGroupPathSelectorMetadataOwnerReferencesVersion) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Version().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesVersion) WithValue(value string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesVersion) WithArrayOfValues(values []string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataOwnerReferencesName struct{}

func (ProbeGroupPathSelectorMetadataOwnerReferencesName) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Name().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesName) WithValue(value string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesName) WithArrayOfValues(values []string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataOwnerReferencesRegion struct{}

func (ProbeGroupPathSelectorMetadataOwnerReferencesRegion) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Region().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesRegion) WithValue(value string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesRegion) WithArrayOfValues(values []string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataOwnerReferencesController struct{}

func (ProbeGroupPathSelectorMetadataOwnerReferencesController) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().Controller().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesController) WithValue(value bool) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesController) WithArrayOfValues(values []bool) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion struct{}

func (ProbeGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().BlockOwnerDeletion().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithValue(value bool) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesBlockOwnerDeletion) WithArrayOfValues(values []bool) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference struct{}

func (ProbeGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().OwnerReferences().RequiresOwnerReference().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithValue(value bool) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataOwnerReferencesRequiresOwnerReference) WithArrayOfValues(values []bool) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataShards struct{}

func (ProbeGroupPathSelectorMetadataShards) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataShards) WithValue(value map[string]int64) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataShards) WithArrayOfValues(values []map[string]int64) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

func (ProbeGroupPathSelectorMetadataShards) WithKey(key string) ProbeGroupMapPathSelectorMetadataShards {
	return ProbeGroupMapPathSelectorMetadataShards{key: key}
}

type ProbeGroupMapPathSelectorMetadataShards struct {
	key string
}

func (s ProbeGroupMapPathSelectorMetadataShards) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Shards().WithKey(s.key).FieldPath(),
	}
}

func (s ProbeGroupMapPathSelectorMetadataShards) WithValue(value int64) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupMapPathSelectorMetadataShards) WithArrayOfValues(values []int64) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataSyncing struct{}

func (ProbeGroupPathSelectorMetadataSyncing) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataSyncing) WithValue(value *ntt_meta.SyncingMeta) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataSyncing) WithArrayOfValues(values []*ntt_meta.SyncingMeta) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

func (ProbeGroupPathSelectorMetadataSyncing) OwningRegion() ProbeGroupPathSelectorMetadataSyncingOwningRegion {
	return ProbeGroupPathSelectorMetadataSyncingOwningRegion{}
}

func (ProbeGroupPathSelectorMetadataSyncing) Regions() ProbeGroupPathSelectorMetadataSyncingRegions {
	return ProbeGroupPathSelectorMetadataSyncingRegions{}
}

type ProbeGroupPathSelectorMetadataSyncingOwningRegion struct{}

func (ProbeGroupPathSelectorMetadataSyncingOwningRegion) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().OwningRegion().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataSyncingOwningRegion) WithValue(value string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataSyncingOwningRegion) WithArrayOfValues(values []string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataSyncingRegions struct{}

func (ProbeGroupPathSelectorMetadataSyncingRegions) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Syncing().Regions().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataSyncingRegions) WithValue(value []string) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataSyncingRegions) WithArrayOfValues(values [][]string) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

func (s ProbeGroupPathSelectorMetadataSyncingRegions) WithItemValue(value string) *ProbeGroup_FieldSubPathArrayItemValue {
	return s.FieldPath().WithIArrayItemValue(value).(*ProbeGroup_FieldSubPathArrayItemValue)
}

type ProbeGroupPathSelectorMetadataLifecycle struct{}

func (ProbeGroupPathSelectorMetadataLifecycle) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataLifecycle) WithValue(value *ntt_meta.Lifecycle) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataLifecycle) WithArrayOfValues(values []*ntt_meta.Lifecycle) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

func (ProbeGroupPathSelectorMetadataLifecycle) State() ProbeGroupPathSelectorMetadataLifecycleState {
	return ProbeGroupPathSelectorMetadataLifecycleState{}
}

func (ProbeGroupPathSelectorMetadataLifecycle) BlockDeletion() ProbeGroupPathSelectorMetadataLifecycleBlockDeletion {
	return ProbeGroupPathSelectorMetadataLifecycleBlockDeletion{}
}

type ProbeGroupPathSelectorMetadataLifecycleState struct{}

func (ProbeGroupPathSelectorMetadataLifecycleState) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().State().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataLifecycleState) WithValue(value ntt_meta.Lifecycle_State) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataLifecycleState) WithArrayOfValues(values []ntt_meta.Lifecycle_State) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupPathSelectorMetadataLifecycleBlockDeletion struct{}

func (ProbeGroupPathSelectorMetadataLifecycleBlockDeletion) FieldPath() *ProbeGroup_FieldSubPath {
	return &ProbeGroup_FieldSubPath{
		selector: ProbeGroup_FieldPathSelectorMetadata,
		subPath:  ntt_meta.NewMetaFieldPathBuilder().Lifecycle().BlockDeletion().FieldPath(),
	}
}

func (s ProbeGroupPathSelectorMetadataLifecycleBlockDeletion) WithValue(value bool) *ProbeGroup_FieldSubPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroup_FieldSubPathValue)
}

func (s ProbeGroupPathSelectorMetadataLifecycleBlockDeletion) WithArrayOfValues(values []bool) *ProbeGroup_FieldSubPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroup_FieldSubPathArrayOfValues)
}

type ProbeGroupStateFieldPathBuilder struct{}

func NewProbeGroupStateFieldPathBuilder() ProbeGroupStateFieldPathBuilder {
	return ProbeGroupStateFieldPathBuilder{}
}
func (ProbeGroupStateFieldPathBuilder) TargetCount() ProbeGroup_StatePathSelectorTargetCount {
	return ProbeGroup_StatePathSelectorTargetCount{}
}
func (ProbeGroupStateFieldPathBuilder) RegionalTargetCounts() ProbeGroup_StatePathSelectorRegionalTargetCounts {
	return ProbeGroup_StatePathSelectorRegionalTargetCounts{}
}

type ProbeGroup_StatePathSelectorTargetCount struct{}

func (ProbeGroup_StatePathSelectorTargetCount) FieldPath() *ProbeGroupState_FieldTerminalPath {
	return &ProbeGroupState_FieldTerminalPath{selector: ProbeGroupState_FieldPathSelectorTargetCount}
}

func (s ProbeGroup_StatePathSelectorTargetCount) WithValue(value int64) *ProbeGroupState_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroupState_FieldTerminalPathValue)
}

func (s ProbeGroup_StatePathSelectorTargetCount) WithArrayOfValues(values []int64) *ProbeGroupState_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroupState_FieldTerminalPathArrayOfValues)
}

type ProbeGroup_StatePathSelectorRegionalTargetCounts struct{}

func (ProbeGroup_StatePathSelectorRegionalTargetCounts) FieldPath() *ProbeGroupState_FieldTerminalPath {
	return &ProbeGroupState_FieldTerminalPath{selector: ProbeGroupState_FieldPathSelectorRegionalTargetCounts}
}

func (s ProbeGroup_StatePathSelectorRegionalTargetCounts) WithValue(value map[string]int64) *ProbeGroupState_FieldTerminalPathValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroupState_FieldTerminalPathValue)
}

func (s ProbeGroup_StatePathSelectorRegionalTargetCounts) WithArrayOfValues(values []map[string]int64) *ProbeGroupState_FieldTerminalPathArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroupState_FieldTerminalPathArrayOfValues)
}

func (ProbeGroup_StatePathSelectorRegionalTargetCounts) WithKey(key string) ProbeGroup_StateMapPathSelectorRegionalTargetCounts {
	return ProbeGroup_StateMapPathSelectorRegionalTargetCounts{key: key}
}

type ProbeGroup_StateMapPathSelectorRegionalTargetCounts struct {
	key string
}

func (s ProbeGroup_StateMapPathSelectorRegionalTargetCounts) FieldPath() *ProbeGroupState_FieldPathMap {
	return &ProbeGroupState_FieldPathMap{selector: ProbeGroupState_FieldPathSelectorRegionalTargetCounts, key: s.key}
}

func (s ProbeGroup_StateMapPathSelectorRegionalTargetCounts) WithValue(value int64) *ProbeGroupState_FieldPathMapValue {
	return s.FieldPath().WithIValue(value).(*ProbeGroupState_FieldPathMapValue)
}

func (s ProbeGroup_StateMapPathSelectorRegionalTargetCounts) WithArrayOfValues(values []int64) *ProbeGroupState_FieldPathMapArrayOfValues {
	return s.FieldPath().WithIArrayOfValues(values).(*ProbeGroupState_FieldPathMapArrayOfValues)
}
