// Code generated by protoc-gen-goten-resource
// Resource: Probe
// DO NOT EDIT!!!

package probe

import (
	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	ntt_memo "github.com/cloudwan/edgelq-sdk/common/types/memo"
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	devices_device "github.com/cloudwan/edgelq-sdk/devices/resources/v1alpha2/device"
	iam_service_account "github.com/cloudwan/edgelq-sdk/iam/resources/v1alpha2/service_account"
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// ensure the imports are used
var (
	_ = gotenobject.FieldPath(nil)
	_ = gotenresource.ListQuery(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_memo.Memo{}
	_ = &ntt_meta.Meta{}
	_ = &devices_device.Device{}
	_ = &iam_service_account.ServiceAccount{}
	_ = &timestamp.Timestamp{}
	_ = &common.SoftwareVersion{}
	_ = &probe_group.ProbeGroup{}
	_ = &project.Project{}
)

var (
	descriptor = &Descriptor{
		typeName: gotenresource.NewTypeName(
			"Probe", "Probes", "watchdog.edgelq.com"),
		nameDescriptor: gotenresource.NewNameDescriptor(
			&Probe_FieldTerminalPath{selector: Probe_FieldPathSelectorName},
			"pattern", "probeId",
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

func (d *Descriptor) NewProbe() *Probe {
	return &Probe{}
}

func (d *Descriptor) NewResource() gotenresource.Resource {
	return d.NewProbe()
}

func (d *Descriptor) NewResourceName() gotenresource.Name {
	return NewNameBuilder().Name()
}

func (d *Descriptor) NewProbeName() *Name {
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
func (d *Descriptor) NewProbeCursor() *PagerCursor {
	return &PagerCursor{}
}

func (d *Descriptor) NewResourceCursor() gotenresource.Cursor {
	return d.NewProbeCursor()
}
func (d *Descriptor) NewProbeChange() *ProbeChange {
	return &ProbeChange{}
}

func (d *Descriptor) NewResourceChange() gotenresource.ResourceChange {
	return d.NewProbeChange()
}

func (d *Descriptor) NewProbeQueryResultSnapshot() *QueryResultSnapshot {
	return &QueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultSnapshot() gotenresource.QueryResultSnapshot {
	return d.NewProbeQueryResultSnapshot()
}
func (d *Descriptor) NewProbeQueryResultChange() *QueryResultChange {
	return &QueryResultChange{}
}

func (d *Descriptor) NewSearchQueryResultSnapshot() gotenresource.SearchQueryResultSnapshot {
	return &SearchQueryResultSnapshot{}
}

func (d *Descriptor) NewQueryResultChange() gotenresource.QueryResultChange {
	return d.NewProbeQueryResultChange()
}

func (d *Descriptor) NewProbeList(size, reserved int) ProbeList {
	return make(ProbeList, size, reserved)
}

func (d *Descriptor) NewResourceList(size, reserved int) gotenresource.ResourceList {
	return make(ProbeList, size, reserved)
}
func (d *Descriptor) NewProbeChangeList(size, reserved int) ProbeChangeList {
	return make(ProbeChangeList, size, reserved)
}

func (d *Descriptor) NewResourceChangeList(size, reserved int) gotenresource.ResourceChangeList {
	return make(ProbeChangeList, size, reserved)
}

func (d *Descriptor) NewProbeNameList(size, reserved int) ProbeNameList {
	return make(ProbeNameList, size, reserved)
}

func (d *Descriptor) NewNameList(size, reserved int) gotenresource.NameList {
	return make(ProbeNameList, size, reserved)
}

func (d *Descriptor) NewProbeReferenceList(size, reserved int) ProbeReferenceList {
	return make(ProbeReferenceList, size, reserved)
}

func (d *Descriptor) NewReferenceList(size, reserved int) gotenresource.ReferenceList {
	return make(ProbeReferenceList, size, reserved)
}
func (d *Descriptor) NewProbeParentNameList(size, reserved int) ProbeParentNameList {
	return make(ProbeParentNameList, size, reserved)
}

func (d *Descriptor) NewParentNameList(size, reserved int) gotenresource.ParentNameList {
	return make(ProbeParentNameList, size, reserved)
}
func (d *Descriptor) NewProbeParentReferenceList(size, reserved int) ProbeParentReferenceList {
	return make(ProbeParentReferenceList, size, reserved)
}

func (d *Descriptor) NewParentReferenceList(size, reserved int) gotenresource.ParentReferenceList {
	return make(ProbeParentReferenceList, size, reserved)
}

func (d *Descriptor) NewProbeMap(reserved int) ProbeMap {
	return make(ProbeMap, reserved)
}

func (d *Descriptor) NewResourceMap(reserved int) gotenresource.ResourceMap {
	return make(ProbeMap, reserved)
}
func (d *Descriptor) NewProbeChangeMap(reserved int) ProbeChangeMap {
	return make(ProbeChangeMap, reserved)
}

func (d *Descriptor) NewResourceChangeMap(reserved int) gotenresource.ResourceChangeMap {
	return make(ProbeChangeMap, reserved)
}

func (d *Descriptor) GetResourceTypeName() *gotenresource.TypeName {
	return d.typeName
}

func (d *Descriptor) GetNameDescriptor() *gotenresource.NameDescriptor {
	return d.nameDescriptor
}

func (d *Descriptor) ParseFieldPath(raw string) (gotenobject.FieldPath, error) {
	return ParseProbe_FieldPath(raw)
}

func (d *Descriptor) ParseProbeName(nameStr string) (*Name, error) {
	return ParseName(nameStr)
}

func (d *Descriptor) ParseResourceName(nameStr string) (gotenresource.Name, error) {
	return ParseName(nameStr)
}
