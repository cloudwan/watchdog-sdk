// Code generated by protoc-gen-goten-resource
// Resource: ProbingConfig
// DO NOT EDIT!!!

package probing_config

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
	duration "github.com/golang/protobuf/ptypes/duration"
)

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &duration.Duration{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target.ProbingTarget{}
	_ = &project.Project{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"ProbingConfig", "ProbingConfigs", "watchdog.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&ProbingConfig_FieldTerminalPath{selector: ProbingConfig_FieldPathSelectorName},
			"pattern", "probingConfigId",
			[]string{"projectId", "regionId"},
			[]gotenresource.NamePattern{NamePattern_Project_Region}),
	}
)

type Descriptor struct {
	nameDescriptor *gotenresource.NameDescriptor
	typeName       *gotenresource.TypeName
}

func GetDescriptor() *Descriptor {
	return descriptor
}

func (d *Descriptor) NewProbingConfig() *ProbingConfig {
	return &ProbingConfig{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewProbingConfig()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewProbingConfigName() *Name {
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
func (d *Descriptor) NewProbingConfigCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewProbingConfigCursor()
}
func (d *Descriptor) NewProbingConfigChange() *ProbingConfigChange {
	return &ProbingConfigChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewProbingConfigChange()
}

func (d *Descriptor) NewProbingConfigQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewProbingConfigQueryResultSnapshot()
}
func (d *Descriptor) NewProbingConfigQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return nil
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewProbingConfigQueryResultChange()
}

func (d *Descriptor) NewProbingConfigList(size, reserved int) ProbingConfigList {
	return make(ProbingConfigList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(ProbingConfigList, size, reserved)
}
func (d *Descriptor) NewProbingConfigChangeList(size, reserved int) ProbingConfigChangeList {
	return make(ProbingConfigChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(ProbingConfigChangeList, size, reserved)
}

func (d *Descriptor) NewProbingConfigNameList(size, reserved int) ProbingConfigNameList {
	return make(ProbingConfigNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(ProbingConfigNameList, size, reserved)
}

func (d *Descriptor) NewProbingConfigReferenceList(size, reserved int) ProbingConfigReferenceList {
	return make(ProbingConfigReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(ProbingConfigReferenceList, size, reserved)
}
func (d *Descriptor) NewProbingConfigParentNameList(size, reserved int) ProbingConfigParentNameList {
	return make(ProbingConfigParentNameList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(ProbingConfigParentNameList, size, reserved)
}
func (d *Descriptor) NewProbingConfigParentReferenceList(size, reserved int) ProbingConfigParentReferenceList {
	return make(ProbingConfigParentReferenceList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(ProbingConfigParentReferenceList, size, reserved)
}

func (d *Descriptor) NewProbingConfigMap(reserved int) ProbingConfigMap {
	return make(ProbingConfigMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(ProbingConfigMap, reserved)
}
func (d *Descriptor) NewProbingConfigChangeMap(reserved int) ProbingConfigChangeMap {
	return make(ProbingConfigChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(ProbingConfigChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseProbingConfig_FieldPath(raw)
}

func (d *Descriptor) ParseProbingConfigName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}
