// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/hop_monitor_custom.proto
// DO NOT EDIT!!!

package hop_monitor_client

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probing_target "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target"
)

// ensure the imports are used
var (
	_ = new(fmt.Stringer)
	_ = new(sort.Interface)

	_ = new(proto.Message)
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldPath)
)

// make sure we're using proto imports
var (
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target.ProbingTarget{}
)

func (o *RunHopMonitorRequest) GotenObjectExt() {}

func (o *RunHopMonitorRequest) MakeFullFieldMask() *RunHopMonitorRequest_FieldMask {
	return FullRunHopMonitorRequest_FieldMask()
}

func (o *RunHopMonitorRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunHopMonitorRequest_FieldMask()
}

func (o *RunHopMonitorRequest) MakeDiffFieldMask(other *RunHopMonitorRequest) *RunHopMonitorRequest_FieldMask {
	if o == nil && other == nil {
		return &RunHopMonitorRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunHopMonitorRequest_FieldMask()
	}

	res := &RunHopMonitorRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorName})
	}
	if o.GetSource() != other.GetSource() {
		res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorSource})
	}
	if o.GetTarget().String() != other.GetTarget().String() {
		res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorTarget})
	}
	if o.GetDestination() != other.GetDestination() {
		res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorDestination})
	}
	if o.GetSizeBytes() != other.GetSizeBytes() {
		res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorSizeBytes})
	}
	if o.GetAttempts() != other.GetAttempts() {
		res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorAttempts})
	}
	if o.GetMode() != other.GetMode() {
		res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorMode})
	}
	return res
}

func (o *RunHopMonitorRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunHopMonitorRequest))
}

func (o *RunHopMonitorRequest) Clone() *RunHopMonitorRequest {
	if o == nil {
		return nil
	}
	result := &RunHopMonitorRequest{}
	if o.Name == nil {
		result.Name = nil
	} else if data, err := o.Name.ProtoString(); err != nil {
		panic(err)
	} else {
		result.Name = &probe.Reference{}
		if err := result.Name.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	result.Source = o.Source
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
	result.Destination = o.Destination
	result.SizeBytes = o.SizeBytes
	result.Attempts = o.Attempts
	result.Mode = o.Mode
	return result
}

func (o *RunHopMonitorRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunHopMonitorRequest) Merge(source *RunHopMonitorRequest) {
	if source.GetName() != nil {
		if data, err := source.GetName().ProtoString(); err != nil {
			panic(err)
		} else {
			o.Name = &probe.Reference{}
			if err := o.Name.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.Name = nil
	}
	o.Source = source.GetSource()
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
	o.Destination = source.GetDestination()
	o.SizeBytes = source.GetSizeBytes()
	o.Attempts = source.GetAttempts()
	o.Mode = source.GetMode()
}

func (o *RunHopMonitorRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunHopMonitorRequest))
}

func (o *RunHopMonitorResponse) GotenObjectExt() {}

func (o *RunHopMonitorResponse) MakeFullFieldMask() *RunHopMonitorResponse_FieldMask {
	return FullRunHopMonitorResponse_FieldMask()
}

func (o *RunHopMonitorResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunHopMonitorResponse_FieldMask()
}

func (o *RunHopMonitorResponse) MakeDiffFieldMask(other *RunHopMonitorResponse) *RunHopMonitorResponse_FieldMask {
	if o == nil && other == nil {
		return &RunHopMonitorResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunHopMonitorResponse_FieldMask()
	}

	res := &RunHopMonitorResponse_FieldMask{}

	if len(o.GetPaths()) == len(other.GetPaths()) {
		for i, lValue := range o.GetPaths() {
			rValue := other.GetPaths()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorPaths})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorPaths})
	}

	if len(o.GetHopStats()) == len(other.GetHopStats()) {
		for i, lValue := range o.GetHopStats() {
			rValue := other.GetHopStats()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorHopStats})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorHopStats})
	}

	if len(o.GetHopInfo()) == len(other.GetHopInfo()) {
		for i, lValue := range o.GetHopInfo() {
			rValue := other.GetHopInfo()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorHopInfo})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorHopInfo})
	}
	if o.GetIpVersion() != other.GetIpVersion() {
		res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorIpVersion})
	}
	return res
}

func (o *RunHopMonitorResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunHopMonitorResponse))
}

func (o *RunHopMonitorResponse) Clone() *RunHopMonitorResponse {
	if o == nil {
		return nil
	}
	result := &RunHopMonitorResponse{}
	result.Paths = make([]*common.Path, len(o.Paths))
	for i, sourceValue := range o.Paths {
		result.Paths[i] = sourceValue.Clone()
	}
	result.HopStats = map[string]*common.HopStat{}
	for key, sourceValue := range o.HopStats {
		result.HopStats[key] = sourceValue.Clone()
	}
	result.HopInfo = map[string]*common.HopInfo{}
	for key, sourceValue := range o.HopInfo {
		result.HopInfo[key] = sourceValue.Clone()
	}
	result.IpVersion = o.IpVersion
	return result
}

func (o *RunHopMonitorResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunHopMonitorResponse) Merge(source *RunHopMonitorResponse) {
	for _, sourceValue := range source.GetPaths() {
		exists := false
		for _, currentValue := range o.Paths {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.Path
			if sourceValue != nil {
				newDstElement = new(common.Path)
				newDstElement.Merge(sourceValue)
			}
			o.Paths = append(o.Paths, newDstElement)
		}
	}

	if source.GetHopStats() != nil {
		if o.HopStats == nil {
			o.HopStats = make(map[string]*common.HopStat, len(source.GetHopStats()))
		}
		for key, sourceValue := range source.GetHopStats() {
			if sourceValue != nil {
				if o.HopStats[key] == nil {
					o.HopStats[key] = new(common.HopStat)
				}
				o.HopStats[key].Merge(sourceValue)
			}
		}
	}
	if source.GetHopInfo() != nil {
		if o.HopInfo == nil {
			o.HopInfo = make(map[string]*common.HopInfo, len(source.GetHopInfo()))
		}
		for key, sourceValue := range source.GetHopInfo() {
			if sourceValue != nil {
				if o.HopInfo[key] == nil {
					o.HopInfo[key] = new(common.HopInfo)
				}
				o.HopInfo[key].Merge(sourceValue)
			}
		}
	}
	o.IpVersion = source.GetIpVersion()
}

func (o *RunHopMonitorResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunHopMonitorResponse))
}
