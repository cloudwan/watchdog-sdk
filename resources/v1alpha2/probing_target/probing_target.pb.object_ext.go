// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/probing_target.proto
// DO NOT EDIT!!!

package probing_target

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probing_target_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = fmt.Stringer(nil)
	_ = sort.Interface(nil)

	_ = proto.Message(nil)
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldPath(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &project.Project{}
)

func (o *ProbingTarget) GotenObjectExt() {}

func (o *ProbingTarget) MakeFullFieldMask() *ProbingTarget_FieldMask {
	return FullProbingTarget_FieldMask()
}

func (o *ProbingTarget) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProbingTarget_FieldMask()
}

func (o *ProbingTarget) MakeDiffFieldMask(other *ProbingTarget) *ProbingTarget_FieldMask {
	if o == nil && other == nil {
		return &ProbingTarget_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProbingTarget_FieldMask()
	}

	res := &ProbingTarget_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorName})
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorDisplayName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingTarget_FieldSubPath{selector: ProbingTarget_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	if o.GetGroup().String() != other.GetGroup().String() {
		res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorGroup})
	}
	if o.GetMode() != other.GetMode() {
		res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorMode})
	}
	if o.GetIpVersion() != other.GetIpVersion() {
		res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorIpVersion})
	}
	if o.GetAddress() != other.GetAddress() {
		res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorAddress})
	}
	if o.GetCategory() != other.GetCategory() {
		res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorCategory})
	}
	if o.GetLocationType() != other.GetLocationType() {
		res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorLocationType})
	}
	{
		subMask := o.GetLocation().MakeDiffFieldMask(other.GetLocation())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorLocation})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingTarget_FieldSubPath{selector: ProbingTarget_FieldPathSelectorLocation, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetProbingConstraint().MakeDiffFieldMask(other.GetProbingConstraint())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorProbingConstraint})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingTarget_FieldSubPath{selector: ProbingTarget_FieldPathSelectorProbingConstraint, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetDefaultProbingSettings().MakeDiffFieldMask(other.GetDefaultProbingSettings())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorDefaultProbingSettings})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingTarget_FieldSubPath{selector: ProbingTarget_FieldPathSelectorDefaultProbingSettings, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetHttpProbingConfig().MakeDiffFieldMask(other.GetHttpProbingConfig())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorHttpProbingConfig})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingTarget_FieldSubPath{selector: ProbingTarget_FieldPathSelectorHttpProbingConfig, subPath: subpath})
			}
		}
	}
	if o.GetAgent().String() != other.GetAgent().String() {
		res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorAgent})
	}
	return res
}

func (o *ProbingTarget) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProbingTarget))
}

func (o *ProbingTarget) Clone() *ProbingTarget {
	if o == nil {
		return nil
	}
	result := &ProbingTarget{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &Name{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.DisplayName = o.DisplayName
	result.Metadata = o.Metadata.Clone()
	if o.Group == nil {
		result.Group = nil
	} else if data, err := o.Group.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Group = &probing_target_group.Reference{}
		if err := result.Group.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Mode = o.Mode
	result.IpVersion = o.IpVersion
	result.Address = o.Address
	result.Category = o.Category
	result.LocationType = o.LocationType
	result.Location = o.Location.Clone()
	result.ProbingConstraint = o.ProbingConstraint.Clone()
	result.DefaultProbingSettings = o.DefaultProbingSettings.Clone()
	result.HttpProbingConfig = o.HttpProbingConfig.Clone()
	if o.Agent == nil {
		result.Agent = nil
	} else if data, err := o.Agent.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Agent = &probe.Reference{}
		if err := result.Agent.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *ProbingTarget) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProbingTarget) Merge(source *ProbingTarget) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &Name{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	o.DisplayName = source.GetDisplayName()
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(ntt_meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
	if source.GetGroup() != nil {
		if data, err := source.GetGroup().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Group = &probing_target_group.Reference{}
			if err := o.Group.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Group = nil
	}
	o.Mode = source.GetMode()
	o.IpVersion = source.GetIpVersion()
	o.Address = source.GetAddress()
	o.Category = source.GetCategory()
	o.LocationType = source.GetLocationType()
	if source.GetLocation() != nil {
		if o.Location == nil {
			o.Location = new(common.Location)
		}
		o.Location.Merge(source.GetLocation())
	}
	if source.GetProbingConstraint() != nil {
		if o.ProbingConstraint == nil {
			o.ProbingConstraint = new(common.ProbingConstraint)
		}
		o.ProbingConstraint.Merge(source.GetProbingConstraint())
	}
	if source.GetDefaultProbingSettings() != nil {
		if o.DefaultProbingSettings == nil {
			o.DefaultProbingSettings = new(common.ProbingSettings)
		}
		o.DefaultProbingSettings.Merge(source.GetDefaultProbingSettings())
	}
	if source.GetHttpProbingConfig() != nil {
		if o.HttpProbingConfig == nil {
			o.HttpProbingConfig = new(common.HTTPProbingConfig)
		}
		o.HttpProbingConfig.Merge(source.GetHttpProbingConfig())
	}
	if source.GetAgent() != nil {
		if data, err := source.GetAgent().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Agent = &probe.Reference{}
			if err := o.Agent.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Agent = nil
	}
}

func (o *ProbingTarget) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProbingTarget))
}
