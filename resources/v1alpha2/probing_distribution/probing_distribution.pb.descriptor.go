// Code generated by protoc-gen-goten-resource
// Resource: ProbingDistribution
// DO NOT EDIT!!!

package probing_distribution

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probing_target "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target.ProbingTarget{}
	_ = &project.Project{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"ProbingDistribution", "ProbingDistributions", "watchdog.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&ProbingDistribution_FieldTerminalPath{selector: ProbingDistribution_FieldPathSelectorName},
			"pattern", "probingDistributionId",
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

func (d *Descriptor) NewProbingDistribution() *ProbingDistribution {
	return &ProbingDistribution{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewProbingDistribution()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewProbingDistributionName() *Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewGetQuery() gotenresource.GetQuery {
	return &GetQuery{}
}

func (d *Descriptor) NewListQuery() gotenresource.ListQuery {
	return &ListQuery{}
}

func (d *Descriptor) NewSearchQuery() gotenresource.SearchQuery {
	return &SearchQuery{}
}

func (d *Descriptor) NewWatchQuery() gotenresource.WatchQuery {
	return &WatchQuery{}
}
func (d *Descriptor) NewProbingDistributionCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewProbingDistributionCursor()
}
func (d *Descriptor) NewProbingDistributionChange() *ProbingDistributionChange {
	return &ProbingDistributionChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewProbingDistributionChange()
}

func (d *Descriptor) NewProbingDistributionQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewProbingDistributionQueryResultSnapshot()
}
func (d *Descriptor) NewProbingDistributionQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return &SearchQueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewProbingDistributionQueryResultChange()
}

func (d *Descriptor) NewProbingDistributionList(size, reserved int) ProbingDistributionList {
	return make(ProbingDistributionList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(ProbingDistributionList, size, reserved)
}
func (d *Descriptor) NewProbingDistributionChangeList(size, reserved int) ProbingDistributionChangeList {
	return make(ProbingDistributionChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(ProbingDistributionChangeList, size, reserved)
}

func (d *Descriptor) NewProbingDistributionNameList(size, reserved int) ProbingDistributionNameList {
	return make(ProbingDistributionNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(ProbingDistributionNameList, size, reserved)
}

func (d *Descriptor) NewProbingDistributionReferenceList(size, reserved int) ProbingDistributionReferenceList {
	return make(ProbingDistributionReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(ProbingDistributionReferenceList, size, reserved)
}
func (d *Descriptor) NewProbingDistributionParentNameList(size, reserved int) ProbingDistributionParentNameList {
	return make(ProbingDistributionParentNameList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(ProbingDistributionParentNameList, size, reserved)
}
func (d *Descriptor) NewProbingDistributionParentReferenceList(size, reserved int) ProbingDistributionParentReferenceList {
	return make(ProbingDistributionParentReferenceList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(ProbingDistributionParentReferenceList, size, reserved)
}

func (d *Descriptor) NewProbingDistributionMap(reserved int) ProbingDistributionMap {
	return make(ProbingDistributionMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(ProbingDistributionMap, reserved)
}
func (d *Descriptor) NewProbingDistributionChangeMap(reserved int) ProbingDistributionChangeMap {
	return make(ProbingDistributionChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(ProbingDistributionChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseProbingDistribution_FieldPath(raw)
}

func (d *Descriptor) ParseProbingDistributionName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}
