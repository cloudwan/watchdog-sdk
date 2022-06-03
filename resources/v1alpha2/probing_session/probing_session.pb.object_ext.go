// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/probing_session.proto
// DO NOT EDIT!!!

package probing_session

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
	probing_distribution "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_distribution"
	probing_target "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target"
	probing_target_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
	duration "github.com/golang/protobuf/ptypes/duration"
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
	_ = &duration.Duration{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_distribution.ProbingDistribution{}
	_ = &probing_target.ProbingTarget{}
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &project.Project{}
)

func (o *ProbingSession) GotenObjectExt() {}

func (o *ProbingSession) MakeFullFieldMask() *ProbingSession_FieldMask {
	return FullProbingSession_FieldMask()
}

func (o *ProbingSession) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProbingSession_FieldMask()
}

func (o *ProbingSession) MakeDiffFieldMask(other *ProbingSession) *ProbingSession_FieldMask {
	if o == nil && other == nil {
		return &ProbingSession_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProbingSession_FieldMask()
	}

	res := &ProbingSession_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorName})
	}
	if o.GetDisplayName() != other.GetDisplayName() {
		res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorDisplayName})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingSession_FieldSubPath{selector: ProbingSession_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetSpec().MakeDiffFieldMask(other.GetSpec())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorSpec})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingSession_FieldSubPath{selector: ProbingSession_FieldPathSelectorSpec, subPath: subpath})
			}
		}
	}
	if o.GetProbingDistribution().String() != other.GetProbingDistribution().String() {
		res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorProbingDistribution})
	}
	{
		subMask := o.GetStatus().MakeDiffFieldMask(other.GetStatus())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorStatus})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingSession_FieldSubPath{selector: ProbingSession_FieldPathSelectorStatus, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ProbingSession) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProbingSession))
}

func (o *ProbingSession) Clone() *ProbingSession {
	if o == nil {
		return nil
	}
	result := &ProbingSession{}
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
	result.Spec = o.Spec.Clone()
	if o.ProbingDistribution == nil {
		result.ProbingDistribution = nil
	} else if data, err := o.ProbingDistribution.ProtoString(); err != nil {
		panic(err)
	} else {
		result.ProbingDistribution = &probing_distribution.Reference{}
		if err := result.ProbingDistribution.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Status = o.Status.Clone()
	return result
}

func (o *ProbingSession) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProbingSession) Merge(source *ProbingSession) {
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
	if source.GetSpec() != nil {
		if o.Spec == nil {
			o.Spec = new(ProbingSession_Spec)
		}
		o.Spec.Merge(source.GetSpec())
	}
	if source.GetProbingDistribution() != nil {
		if data, err := source.GetProbingDistribution().ProtoString(); err != nil {
			panic(err)
		} else {
			o.ProbingDistribution = &probing_distribution.Reference{}
			if err := o.ProbingDistribution.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.ProbingDistribution = nil
	}
	if source.GetStatus() != nil {
		if o.Status == nil {
			o.Status = new(ProbingSession_Status)
		}
		o.Status.Merge(source.GetStatus())
	}
}

func (o *ProbingSession) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProbingSession))
}

func (o *ProbingSession_Spec) GotenObjectExt() {}

func (o *ProbingSession_Spec) MakeFullFieldMask() *ProbingSession_Spec_FieldMask {
	return FullProbingSession_Spec_FieldMask()
}

func (o *ProbingSession_Spec) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProbingSession_Spec_FieldMask()
}

func (o *ProbingSession_Spec) MakeDiffFieldMask(other *ProbingSession_Spec) *ProbingSession_Spec_FieldMask {
	if o == nil && other == nil {
		return &ProbingSession_Spec_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProbingSession_Spec_FieldMask()
	}

	res := &ProbingSession_Spec_FieldMask{}
	if o.GetProbe().String() != other.GetProbe().String() {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorProbe})
	}
	if o.GetTarget().String() != other.GetTarget().String() {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorTarget})
	}
	if o.GetTargetGroup().String() != other.GetTargetGroup().String() {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorTargetGroup})
	}
	if o.GetMode() != other.GetMode() {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorMode})
	}
	if o.GetIpVersion() != other.GetIpVersion() {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorIpVersion})
	}
	if o.GetAddress() != other.GetAddress() {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorAddress})
	}

	if len(o.GetAddresses()) == len(other.GetAddresses()) {
		for i, lValue := range o.GetAddresses() {
			rValue := other.GetAddresses()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorAddresses})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorAddresses})
	}
	if o.GetPort() != other.GetPort() {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorPort})
	}
	if o.GetType() != other.GetType() {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorType})
	}
	if !proto.Equal(o.GetInterval(), other.GetInterval()) {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorInterval})
	}
	if o.GetTos() != other.GetTos() {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorTos})
	}
	{
		subMask := o.GetPathProbing().MakeDiffFieldMask(other.GetPathProbing())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorPathProbing})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingSessionSpec_FieldSubPath{selector: ProbingSessionSpec_FieldPathSelectorPathProbing, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetSpeedtestSettings().MakeDiffFieldMask(other.GetSpeedtestSettings())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorSpeedtestSettings})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingSessionSpec_FieldSubPath{selector: ProbingSessionSpec_FieldPathSelectorSpeedtestSettings, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetHttpProbingConfig().MakeDiffFieldMask(other.GetHttpProbingConfig())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorHttpProbingConfig})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingSessionSpec_FieldSubPath{selector: ProbingSessionSpec_FieldPathSelectorHttpProbingConfig, subPath: subpath})
			}
		}
	}
	{
		subMask := o.GetProxyConfiguration().MakeDiffFieldMask(other.GetProxyConfiguration())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorProxyConfiguration})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingSessionSpec_FieldSubPath{selector: ProbingSessionSpec_FieldPathSelectorProxyConfiguration, subPath: subpath})
			}
		}
	}
	if o.GetLocationType() != other.GetLocationType() {
		res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorLocationType})
	}
	{
		subMask := o.GetLocation().MakeDiffFieldMask(other.GetLocation())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorLocation})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &ProbingSessionSpec_FieldSubPath{selector: ProbingSessionSpec_FieldPathSelectorLocation, subPath: subpath})
			}
		}
	}
	return res
}

func (o *ProbingSession_Spec) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProbingSession_Spec))
}

func (o *ProbingSession_Spec) Clone() *ProbingSession_Spec {
	if o == nil {
		return nil
	}
	result := &ProbingSession_Spec{}
	if o.Probe == nil {
		result.Probe = nil
	} else if data, err := o.Probe.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Probe = &probe.Reference{}
		if err := result.Probe.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.Target == nil {
		result.Target = nil
	} else if data, err := o.Target.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Target = &probing_target.Reference{}
		if err := result.Target.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.TargetGroup == nil {
		result.TargetGroup = nil
	} else if data, err := o.TargetGroup.ProtoString(); err != nil {
		panic(err)
	} else {
		result.TargetGroup = &probing_target_group.Reference{}
		if err := result.TargetGroup.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Mode = o.Mode
	result.IpVersion = o.IpVersion
	result.Address = o.Address
	result.Addresses = make([]string, len(o.Addresses))
	for i, sourceValue := range o.Addresses {
		result.Addresses[i] = sourceValue
	}
	result.Port = o.Port
	result.Type = o.Type
	result.Interval = proto.Clone(o.Interval).(*duration.Duration)
	result.Tos = o.Tos
	result.PathProbing = o.PathProbing.Clone()
	result.SpeedtestSettings = o.SpeedtestSettings.Clone()
	result.HttpProbingConfig = o.HttpProbingConfig.Clone()
	result.ProxyConfiguration = o.ProxyConfiguration.Clone()
	result.LocationType = o.LocationType
	result.Location = o.Location.Clone()
	return result
}

func (o *ProbingSession_Spec) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProbingSession_Spec) Merge(source *ProbingSession_Spec) {
	if source.GetProbe() != nil {
		if data, err := source.GetProbe().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Probe = &probe.Reference{}
			if err := o.Probe.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Probe = nil
	}
	if source.GetTarget() != nil {
		if data, err := source.GetTarget().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Target = &probing_target.Reference{}
			if err := o.Target.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Target = nil
	}
	if source.GetTargetGroup() != nil {
		if data, err := source.GetTargetGroup().ProtoString(); err != nil {
			panic(err)
		} else {
			o.TargetGroup = &probing_target_group.Reference{}
			if err := o.TargetGroup.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.TargetGroup = nil
	}
	o.Mode = source.GetMode()
	o.IpVersion = source.GetIpVersion()
	o.Address = source.GetAddress()
	for _, sourceValue := range source.GetAddresses() {
		exists := false
		for _, currentValue := range o.Addresses {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Addresses = append(o.Addresses, newDstElement)
		}
	}

	o.Port = source.GetPort()
	o.Type = source.GetType()
	if source.GetInterval() != nil {
		if o.Interval == nil {
			o.Interval = new(duration.Duration)
		}
		proto.Merge(o.Interval, source.GetInterval())
	}
	o.Tos = source.GetTos()
	if source.GetPathProbing() != nil {
		if o.PathProbing == nil {
			o.PathProbing = new(common.PathProbe)
		}
		o.PathProbing.Merge(source.GetPathProbing())
	}
	if source.GetSpeedtestSettings() != nil {
		if o.SpeedtestSettings == nil {
			o.SpeedtestSettings = new(common.SpeedTestSettings)
		}
		o.SpeedtestSettings.Merge(source.GetSpeedtestSettings())
	}
	if source.GetHttpProbingConfig() != nil {
		if o.HttpProbingConfig == nil {
			o.HttpProbingConfig = new(common.HTTPProbingConfig)
		}
		o.HttpProbingConfig.Merge(source.GetHttpProbingConfig())
	}
	if source.GetProxyConfiguration() != nil {
		if o.ProxyConfiguration == nil {
			o.ProxyConfiguration = new(common.ProxyConfiguration)
		}
		o.ProxyConfiguration.Merge(source.GetProxyConfiguration())
	}
	o.LocationType = source.GetLocationType()
	if source.GetLocation() != nil {
		if o.Location == nil {
			o.Location = new(common.Location)
		}
		o.Location.Merge(source.GetLocation())
	}
}

func (o *ProbingSession_Spec) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProbingSession_Spec))
}

func (o *ProbingSession_Status) GotenObjectExt() {}

func (o *ProbingSession_Status) MakeFullFieldMask() *ProbingSession_Status_FieldMask {
	return FullProbingSession_Status_FieldMask()
}

func (o *ProbingSession_Status) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullProbingSession_Status_FieldMask()
}

func (o *ProbingSession_Status) MakeDiffFieldMask(other *ProbingSession_Status) *ProbingSession_Status_FieldMask {
	if o == nil && other == nil {
		return &ProbingSession_Status_FieldMask{}
	}
	if o == nil || other == nil {
		return FullProbingSession_Status_FieldMask()
	}

	res := &ProbingSession_Status_FieldMask{}
	return res
}

func (o *ProbingSession_Status) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*ProbingSession_Status))
}

func (o *ProbingSession_Status) Clone() *ProbingSession_Status {
	if o == nil {
		return nil
	}
	result := &ProbingSession_Status{}
	return result
}

func (o *ProbingSession_Status) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *ProbingSession_Status) Merge(source *ProbingSession_Status) {
}

func (o *ProbingSession_Status) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*ProbingSession_Status))
}
