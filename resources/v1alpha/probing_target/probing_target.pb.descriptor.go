// Code generated by protoc-gen-goten-resource
// Resource: ProbingTarget
// DO NOT EDIT!!!

package probing_target

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha/probe"
	probing_target_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha/probing_target_group"
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
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &project.Project{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"ProbingTarget", "ProbingTargets", "watchdog.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorName},
			"pattern", "probingTargetId",
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

func (d *Descriptor) NewProbingTarget() *ProbingTarget {
	return &ProbingTarget{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewProbingTarget()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewProbingTargetName() *Name {
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
func (d *Descriptor) NewProbingTargetCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewProbingTargetCursor()
}
func (d *Descriptor) NewProbingTargetChange() *ProbingTargetChange {
	return &ProbingTargetChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewProbingTargetChange()
}

func (d *Descriptor) NewProbingTargetQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewProbingTargetQueryResultSnapshot()
}
func (d *Descriptor) NewProbingTargetQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewProbingTargetQueryResultChange()
}

func (d *Descriptor) NewProbingTargetList(size, reserved int) ProbingTargetList {
	return make(ProbingTargetList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(ProbingTargetList, size, reserved)
}
func (d *Descriptor) NewProbingTargetChangeList(size, reserved int) ProbingTargetChangeList {
	return make(ProbingTargetChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(ProbingTargetChangeList, size, reserved)
}

func (d *Descriptor) NewProbingTargetNameList(size, reserved int) ProbingTargetNameList {
	return make(ProbingTargetNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(ProbingTargetNameList, size, reserved)
}

func (d *Descriptor) NewProbingTargetReferenceList(size, reserved int) ProbingTargetReferenceList {
	return make(ProbingTargetReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(ProbingTargetReferenceList, size, reserved)
}
func (d *Descriptor) NewProbingTargetParentNameList(size, reserved int) ProbingTargetParentNameList {
	return make(ProbingTargetParentNameList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(ProbingTargetParentNameList, size, reserved)
}
func (d *Descriptor) NewProbingTargetParentReferenceList(size, reserved int) ProbingTargetParentReferenceList {
	return make(ProbingTargetParentReferenceList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(ProbingTargetParentReferenceList, size, reserved)
}

func (d *Descriptor) NewProbingTargetMap(reserved int) ProbingTargetMap {
	return make(ProbingTargetMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(ProbingTargetMap, reserved)
}
func (d *Descriptor) NewProbingTargetChangeMap(reserved int) ProbingTargetChangeMap {
	return make(ProbingTargetChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(ProbingTargetChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseProbingTarget_FieldPath(raw)
}

func (d *Descriptor) ParseProbingTargetName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}
