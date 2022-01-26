// Code generated by protoc-gen-goten-resource
// Resource: QualityProfile
// DO NOT EDIT!!!

package quality_profile

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha/project"
)

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &project.Project{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"QualityProfile", "QualityProfiles", "watchdog.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&QualityProfile_FieldTerminalPath{selector: QualityProfile_FieldPathSelectorName},
			"pattern", "qualityProfileId",
			[]string{"projectId"},
			[]gotenresource.NamePattern{NamePattern_Project}),
	}
)

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewQualityProfile() *QualityProfile {
	return &QualityProfile{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewQualityProfile()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewQualityProfileName() *Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewGetQuery() gotenresource.GetQuery {
	return &GetQuery{}
}

func (d *Descriptor) NewListQuery() gotenresource.ListQuery {
	return &ListQuery{}
}

func (d *Descriptor) NewSearchQuery() gotenresource.SearchQuery {
	return nil
}

func (d *Descriptor) NewWatchQuery() gotenresource.WatchQuery {
	return &WatchQuery{}
}
func (d *Descriptor) NewQualityProfileCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewQualityProfileCursor()
}
func (d *Descriptor) NewQualityProfileChange() *QualityProfileChange {
	return &QualityProfileChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewQualityProfileChange()
}

func (d *Descriptor) NewQualityProfileQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewQualityProfileQueryResultSnapshot()
}
func (d *Descriptor) NewQualityProfileQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewQualityProfileQueryResultChange()
}

func (d *Descriptor) NewQualityProfileList(size, reserved int) QualityProfileList {
	return make(QualityProfileList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(QualityProfileList, size, reserved)
}
func (d *Descriptor) NewQualityProfileChangeList(size, reserved int) QualityProfileChangeList {
	return make(QualityProfileChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(QualityProfileChangeList, size, reserved)
}

func (d *Descriptor) NewQualityProfileNameList(size, reserved int) QualityProfileNameList {
	return make(QualityProfileNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(QualityProfileNameList, size, reserved)
}

func (d *Descriptor) NewQualityProfileReferenceList(size, reserved int) QualityProfileReferenceList {
	return make(QualityProfileReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(QualityProfileReferenceList, size, reserved)
}
func (d *Descriptor) NewQualityProfileParentNameList(size, reserved int) QualityProfileParentNameList {
	return make(QualityProfileParentNameList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(QualityProfileParentNameList, size, reserved)
}
func (d *Descriptor) NewQualityProfileParentReferenceList(size, reserved int) QualityProfileParentReferenceList {
	return make(QualityProfileParentReferenceList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(QualityProfileParentReferenceList, size, reserved)
}

func (d *Descriptor) NewQualityProfileMap(reserved int) QualityProfileMap {
	return make(QualityProfileMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(QualityProfileMap, reserved)
}
func (d *Descriptor) NewQualityProfileChangeMap(reserved int) QualityProfileChangeMap {
	return make(QualityProfileChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(QualityProfileChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseQualityProfile_FieldPath(raw)
}

func (d *Descriptor) ParseQualityProfileName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}
