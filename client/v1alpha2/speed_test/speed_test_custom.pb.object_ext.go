// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/speed_test_custom.proto
// DO NOT EDIT!!!

package speed_test_client

import (
	"fmt"
	"sort"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gotenobject "github.com/cloudwan/goten-sdk/runtime/object"
)

// proto imports
import (
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probing_session "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_session"
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
	_ = &probe.Probe{}
	_ = &probing_session.ProbingSession{}
	_ = &probing_target.ProbingTarget{}
)

func (o *RunSpeedTestRequest) GotenObjectExt() {}

func (o *RunSpeedTestRequest) MakeFullFieldMask() *RunSpeedTestRequest_FieldMask {
	return FullRunSpeedTestRequest_FieldMask()
}

func (o *RunSpeedTestRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunSpeedTestRequest_FieldMask()
}

func (o *RunSpeedTestRequest) MakeDiffFieldMask(other *RunSpeedTestRequest) *RunSpeedTestRequest_FieldMask {
	if o == nil && other == nil {
		return &RunSpeedTestRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunSpeedTestRequest_FieldMask()
	}

	res := &RunSpeedTestRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &RunSpeedTestRequest_FieldTerminalPath{selector: RunSpeedTestRequest_FieldPathSelectorName})
	}
	if o.GetDirection() != other.GetDirection() {
		res.Paths = append(res.Paths, &RunSpeedTestRequest_FieldTerminalPath{selector: RunSpeedTestRequest_FieldPathSelectorDirection})
	}
	if o.GetProbingSession().String() != other.GetProbingSession().String() {
		res.Paths = append(res.Paths, &RunSpeedTestRequest_FieldTerminalPath{selector: RunSpeedTestRequest_FieldPathSelectorProbingSession})
	}
	if o.GetProbingTarget().String() != other.GetProbingTarget().String() {
		res.Paths = append(res.Paths, &RunSpeedTestRequest_FieldTerminalPath{selector: RunSpeedTestRequest_FieldPathSelectorProbingTarget})
	}
	return res
}

func (o *RunSpeedTestRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunSpeedTestRequest))
}

func (o *RunSpeedTestRequest) Clone() *RunSpeedTestRequest {
	if o == nil {
		return nil
	}
	result := &RunSpeedTestRequest{}
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
	result.Direction = o.Direction
	if o.ProbingSession == nil {
		result.ProbingSession = nil
	} else if data, err := o.ProbingSession.ProtoString(); err != nil {
		panic(err)
	} else {
		result.ProbingSession = &probing_session.Reference{}
		if err := result.ProbingSession.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	if o.ProbingTarget == nil {
		result.ProbingTarget = nil
	} else if data, err := o.ProbingTarget.ProtoString(); err != nil {
		panic(err)
	} else {
		result.ProbingTarget = &probing_target.Reference{}
		if err := result.ProbingTarget.ParseProtoString(data); err != nil {
			panic(err)
		}
	}
	return result
}

func (o *RunSpeedTestRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunSpeedTestRequest) Merge(source *RunSpeedTestRequest) {
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
	o.Direction = source.GetDirection()
	if source.GetProbingSession() != nil {
		if data, err := source.GetProbingSession().ProtoString(); err != nil {
			panic(err)
		} else {
			o.ProbingSession = &probing_session.Reference{}
			if err := o.ProbingSession.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.ProbingSession = nil
	}
	if source.GetProbingTarget() != nil {
		if data, err := source.GetProbingTarget().ProtoString(); err != nil {
			panic(err)
		} else {
			o.ProbingTarget = &probing_target.Reference{}
			if err := o.ProbingTarget.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.ProbingTarget = nil
	}
}

func (o *RunSpeedTestRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunSpeedTestRequest))
}

func (o *RunSpeedTestResponse) GotenObjectExt() {}

func (o *RunSpeedTestResponse) MakeFullFieldMask() *RunSpeedTestResponse_FieldMask {
	return FullRunSpeedTestResponse_FieldMask()
}

func (o *RunSpeedTestResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunSpeedTestResponse_FieldMask()
}

func (o *RunSpeedTestResponse) MakeDiffFieldMask(other *RunSpeedTestResponse) *RunSpeedTestResponse_FieldMask {
	if o == nil && other == nil {
		return &RunSpeedTestResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunSpeedTestResponse_FieldMask()
	}

	res := &RunSpeedTestResponse_FieldMask{}
	if o.GetDirection() != other.GetDirection() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorDirection})
	}
	if o.GetSpeedMbps() != other.GetSpeedMbps() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorSpeedMbps})
	}
	if o.GetServerName() != other.GetServerName() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorServerName})
	}
	if o.GetServerIp() != other.GetServerIp() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorServerIp})
	}
	if o.GetServerLatency() != other.GetServerLatency() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorServerLatency})
	}
	if o.GetServerRetransmitPercentage() != other.GetServerRetransmitPercentage() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorServerRetransmitPercentage})
	}
	if o.GetProbingSession().String() != other.GetProbingSession().String() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorProbingSession})
	}
	if o.GetTarget().String() != other.GetTarget().String() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorTarget})
	}
	if o.GetLocalIp() != other.GetLocalIp() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorLocalIp})
	}
	if o.GetLocalInterface() != other.GetLocalInterface() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorLocalInterface})
	}
	if o.GetBytesTransferred() != other.GetBytesTransferred() {
		res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorBytesTransferred})
	}
	return res
}

func (o *RunSpeedTestResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunSpeedTestResponse))
}

func (o *RunSpeedTestResponse) Clone() *RunSpeedTestResponse {
	if o == nil {
		return nil
	}
	result := &RunSpeedTestResponse{}
	result.Direction = o.Direction
	result.SpeedMbps = o.SpeedMbps
	result.ServerName = o.ServerName
	result.ServerIp = o.ServerIp
	result.ServerLatency = o.ServerLatency
	result.ServerRetransmitPercentage = o.ServerRetransmitPercentage
	if o.ProbingSession == nil {
		result.ProbingSession = nil
	} else if data, err := o.ProbingSession.ProtoString(); err != nil {
		panic(err)
	} else {
		result.ProbingSession = &probing_session.Reference{}
		if err := result.ProbingSession.ParseProtoString(data); err != nil {
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
	result.LocalIp = o.LocalIp
	result.LocalInterface = o.LocalInterface
	result.BytesTransferred = o.BytesTransferred
	return result
}

func (o *RunSpeedTestResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunSpeedTestResponse) Merge(source *RunSpeedTestResponse) {
	o.Direction = source.GetDirection()
	o.SpeedMbps = source.GetSpeedMbps()
	o.ServerName = source.GetServerName()
	o.ServerIp = source.GetServerIp()
	o.ServerLatency = source.GetServerLatency()
	o.ServerRetransmitPercentage = source.GetServerRetransmitPercentage()
	if source.GetProbingSession() != nil {
		if data, err := source.GetProbingSession().ProtoString(); err != nil {
			panic(err)
		} else {
			o.ProbingSession = &probing_session.Reference{}
			if err := o.ProbingSession.ParseProtoString(data); err != nil {
				panic(err)
			}
		}
	} else {
		o.ProbingSession = nil
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
	o.LocalIp = source.GetLocalIp()
	o.LocalInterface = source.GetLocalInterface()
	o.BytesTransferred = source.GetBytesTransferred()
}

func (o *RunSpeedTestResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunSpeedTestResponse))
}

func (o *RunSpeedTestRequestToProbe) GotenObjectExt() {}

func (o *RunSpeedTestRequestToProbe) MakeFullFieldMask() *RunSpeedTestRequestToProbe_FieldMask {
	return FullRunSpeedTestRequestToProbe_FieldMask()
}

func (o *RunSpeedTestRequestToProbe) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunSpeedTestRequestToProbe_FieldMask()
}

func (o *RunSpeedTestRequestToProbe) MakeDiffFieldMask(other *RunSpeedTestRequestToProbe) *RunSpeedTestRequestToProbe_FieldMask {
	if o == nil && other == nil {
		return &RunSpeedTestRequestToProbe_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunSpeedTestRequestToProbe_FieldMask()
	}

	res := &RunSpeedTestRequestToProbe_FieldMask{}
	if o.GetDirection() != other.GetDirection() {
		res.Paths = append(res.Paths, &RunSpeedTestRequestToProbe_FieldTerminalPath{selector: RunSpeedTestRequestToProbe_FieldPathSelectorDirection})
	}
	if o.GetTargetName() != other.GetTargetName() {
		res.Paths = append(res.Paths, &RunSpeedTestRequestToProbe_FieldTerminalPath{selector: RunSpeedTestRequestToProbe_FieldPathSelectorTargetName})
	}
	if o.GetProbingSession() != other.GetProbingSession() {
		res.Paths = append(res.Paths, &RunSpeedTestRequestToProbe_FieldTerminalPath{selector: RunSpeedTestRequestToProbe_FieldPathSelectorProbingSession})
	}
	return res
}

func (o *RunSpeedTestRequestToProbe) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunSpeedTestRequestToProbe))
}

func (o *RunSpeedTestRequestToProbe) Clone() *RunSpeedTestRequestToProbe {
	if o == nil {
		return nil
	}
	result := &RunSpeedTestRequestToProbe{}
	result.Direction = o.Direction
	result.TargetName = o.TargetName
	result.ProbingSession = o.ProbingSession
	return result
}

func (o *RunSpeedTestRequestToProbe) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunSpeedTestRequestToProbe) Merge(source *RunSpeedTestRequestToProbe) {
	o.Direction = source.GetDirection()
	o.TargetName = source.GetTargetName()
	o.ProbingSession = source.GetProbingSession()
}

func (o *RunSpeedTestRequestToProbe) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunSpeedTestRequestToProbe))
}
