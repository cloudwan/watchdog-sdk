// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/ping_test_custom.proto
// DO NOT EDIT!!!

package ping_test_client

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
	_ = &duration.Duration{}
	_ = &probe.Probe{}
)

func (o *RunPingTestRequest) GotenObjectExt() {}

func (o *RunPingTestRequest) MakeFullFieldMask() *RunPingTestRequest_FieldMask {
	return FullRunPingTestRequest_FieldMask()
}

func (o *RunPingTestRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunPingTestRequest_FieldMask()
}

func (o *RunPingTestRequest) MakeDiffFieldMask(other *RunPingTestRequest) *RunPingTestRequest_FieldMask {
	if o == nil && other == nil {
		return &RunPingTestRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunPingTestRequest_FieldMask()
	}

	res := &RunPingTestRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &RunPingTestRequest_FieldTerminalPath{selector: RunPingTestRequest_FieldPathSelectorName})
	}
	if o.GetSource() != other.GetSource() {
		res.Paths = append(res.Paths, &RunPingTestRequest_FieldTerminalPath{selector: RunPingTestRequest_FieldPathSelectorSource})
	}
	if o.GetDestination() != other.GetDestination() {
		res.Paths = append(res.Paths, &RunPingTestRequest_FieldTerminalPath{selector: RunPingTestRequest_FieldPathSelectorDestination})
	}
	if o.GetSizeBytes() != other.GetSizeBytes() {
		res.Paths = append(res.Paths, &RunPingTestRequest_FieldTerminalPath{selector: RunPingTestRequest_FieldPathSelectorSizeBytes})
	}
	if o.GetCount() != other.GetCount() {
		res.Paths = append(res.Paths, &RunPingTestRequest_FieldTerminalPath{selector: RunPingTestRequest_FieldPathSelectorCount})
	}
	if !proto.Equal(o.GetInterval(), other.GetInterval()) {
		res.Paths = append(res.Paths, &RunPingTestRequest_FieldTerminalPath{selector: RunPingTestRequest_FieldPathSelectorInterval})
	}
	if !proto.Equal(o.GetEchoTimeout(), other.GetEchoTimeout()) {
		res.Paths = append(res.Paths, &RunPingTestRequest_FieldTerminalPath{selector: RunPingTestRequest_FieldPathSelectorEchoTimeout})
	}
	if o.GetDontFragment() != other.GetDontFragment() {
		res.Paths = append(res.Paths, &RunPingTestRequest_FieldTerminalPath{selector: RunPingTestRequest_FieldPathSelectorDontFragment})
	}
	if o.GetTtl() != other.GetTtl() {
		res.Paths = append(res.Paths, &RunPingTestRequest_FieldTerminalPath{selector: RunPingTestRequest_FieldPathSelectorTtl})
	}
	if o.GetTos() != other.GetTos() {
		res.Paths = append(res.Paths, &RunPingTestRequest_FieldTerminalPath{selector: RunPingTestRequest_FieldPathSelectorTos})
	}
	return res
}

func (o *RunPingTestRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunPingTestRequest))
}

func (o *RunPingTestRequest) Clone() *RunPingTestRequest {
	if o == nil {
		return nil
	}
	result := &RunPingTestRequest{}
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
	result.Destination = o.Destination
	result.SizeBytes = o.SizeBytes
	result.Count = o.Count
	result.Interval = proto.Clone(o.Interval).(*duration.Duration)
	result.EchoTimeout = proto.Clone(o.EchoTimeout).(*duration.Duration)
	result.DontFragment = o.DontFragment
	result.Ttl = o.Ttl
	result.Tos = o.Tos
	return result
}

func (o *RunPingTestRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunPingTestRequest) Merge(source *RunPingTestRequest) {
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
	o.Destination = source.GetDestination()
	o.SizeBytes = source.GetSizeBytes()
	o.Count = source.GetCount()
	if source.GetInterval() != nil {
		if o.Interval == nil {
			o.Interval = new(duration.Duration)
		}
		proto.Merge(o.Interval, source.GetInterval())
	}
	if source.GetEchoTimeout() != nil {
		if o.EchoTimeout == nil {
			o.EchoTimeout = new(duration.Duration)
		}
		proto.Merge(o.EchoTimeout, source.GetEchoTimeout())
	}
	o.DontFragment = source.GetDontFragment()
	o.Ttl = source.GetTtl()
	o.Tos = source.GetTos()
}

func (o *RunPingTestRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunPingTestRequest))
}

func (o *RunPingTestResponse) GotenObjectExt() {}

func (o *RunPingTestResponse) MakeFullFieldMask() *RunPingTestResponse_FieldMask {
	return FullRunPingTestResponse_FieldMask()
}

func (o *RunPingTestResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunPingTestResponse_FieldMask()
}

func (o *RunPingTestResponse) MakeDiffFieldMask(other *RunPingTestResponse) *RunPingTestResponse_FieldMask {
	if o == nil && other == nil {
		return &RunPingTestResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunPingTestResponse_FieldMask()
	}

	res := &RunPingTestResponse_FieldMask{}
	if o.GetFrom() != other.GetFrom() {
		res.Paths = append(res.Paths, &RunPingTestResponse_FieldTerminalPath{selector: RunPingTestResponse_FieldPathSelectorFrom})
	}
	if o.GetSizeBytes() != other.GetSizeBytes() {
		res.Paths = append(res.Paths, &RunPingTestResponse_FieldTerminalPath{selector: RunPingTestResponse_FieldPathSelectorSizeBytes})
	}
	if o.GetSequenceNumber() != other.GetSequenceNumber() {
		res.Paths = append(res.Paths, &RunPingTestResponse_FieldTerminalPath{selector: RunPingTestResponse_FieldPathSelectorSequenceNumber})
	}
	if o.GetTtl() != other.GetTtl() {
		res.Paths = append(res.Paths, &RunPingTestResponse_FieldTerminalPath{selector: RunPingTestResponse_FieldPathSelectorTtl})
	}
	if !proto.Equal(o.GetRtt(), other.GetRtt()) {
		res.Paths = append(res.Paths, &RunPingTestResponse_FieldTerminalPath{selector: RunPingTestResponse_FieldPathSelectorRtt})
	}
	if o.GetError() != other.GetError() {
		res.Paths = append(res.Paths, &RunPingTestResponse_FieldTerminalPath{selector: RunPingTestResponse_FieldPathSelectorError})
	}
	{
		subMask := o.GetSummary().MakeDiffFieldMask(other.GetSummary())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &RunPingTestResponse_FieldTerminalPath{selector: RunPingTestResponse_FieldPathSelectorSummary})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &RunPingTestResponse_FieldSubPath{selector: RunPingTestResponse_FieldPathSelectorSummary, subPath: subpath})
			}
		}
	}
	return res
}

func (o *RunPingTestResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunPingTestResponse))
}

func (o *RunPingTestResponse) Clone() *RunPingTestResponse {
	if o == nil {
		return nil
	}
	result := &RunPingTestResponse{}
	result.From = o.From
	result.SizeBytes = o.SizeBytes
	result.SequenceNumber = o.SequenceNumber
	result.Ttl = o.Ttl
	result.Rtt = proto.Clone(o.Rtt).(*duration.Duration)
	result.Error = o.Error
	result.Summary = o.Summary.Clone()
	return result
}

func (o *RunPingTestResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunPingTestResponse) Merge(source *RunPingTestResponse) {
	o.From = source.GetFrom()
	o.SizeBytes = source.GetSizeBytes()
	o.SequenceNumber = source.GetSequenceNumber()
	o.Ttl = source.GetTtl()
	if source.GetRtt() != nil {
		if o.Rtt == nil {
			o.Rtt = new(duration.Duration)
		}
		proto.Merge(o.Rtt, source.GetRtt())
	}
	o.Error = source.GetError()
	if source.GetSummary() != nil {
		if o.Summary == nil {
			o.Summary = new(RunPingTestResponse_SummaryStats)
		}
		o.Summary.Merge(source.GetSummary())
	}
}

func (o *RunPingTestResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunPingTestResponse))
}

func (o *RunPingTestResponse_SummaryStats) GotenObjectExt() {}

func (o *RunPingTestResponse_SummaryStats) MakeFullFieldMask() *RunPingTestResponse_SummaryStats_FieldMask {
	return FullRunPingTestResponse_SummaryStats_FieldMask()
}

func (o *RunPingTestResponse_SummaryStats) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunPingTestResponse_SummaryStats_FieldMask()
}

func (o *RunPingTestResponse_SummaryStats) MakeDiffFieldMask(other *RunPingTestResponse_SummaryStats) *RunPingTestResponse_SummaryStats_FieldMask {
	if o == nil && other == nil {
		return &RunPingTestResponse_SummaryStats_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunPingTestResponse_SummaryStats_FieldMask()
	}

	res := &RunPingTestResponse_SummaryStats_FieldMask{}
	if !proto.Equal(o.GetMinRtt(), other.GetMinRtt()) {
		res.Paths = append(res.Paths, &RunPingTestResponseSummaryStats_FieldTerminalPath{selector: RunPingTestResponseSummaryStats_FieldPathSelectorMinRtt})
	}
	if !proto.Equal(o.GetAvgRtt(), other.GetAvgRtt()) {
		res.Paths = append(res.Paths, &RunPingTestResponseSummaryStats_FieldTerminalPath{selector: RunPingTestResponseSummaryStats_FieldPathSelectorAvgRtt})
	}
	if !proto.Equal(o.GetMaxRtt(), other.GetMaxRtt()) {
		res.Paths = append(res.Paths, &RunPingTestResponseSummaryStats_FieldTerminalPath{selector: RunPingTestResponseSummaryStats_FieldPathSelectorMaxRtt})
	}
	if !proto.Equal(o.GetStddevRtt(), other.GetStddevRtt()) {
		res.Paths = append(res.Paths, &RunPingTestResponseSummaryStats_FieldTerminalPath{selector: RunPingTestResponseSummaryStats_FieldPathSelectorStddevRtt})
	}
	if o.GetTransmittedCounter() != other.GetTransmittedCounter() {
		res.Paths = append(res.Paths, &RunPingTestResponseSummaryStats_FieldTerminalPath{selector: RunPingTestResponseSummaryStats_FieldPathSelectorTransmittedCounter})
	}
	if o.GetReceivedCounter() != other.GetReceivedCounter() {
		res.Paths = append(res.Paths, &RunPingTestResponseSummaryStats_FieldTerminalPath{selector: RunPingTestResponseSummaryStats_FieldPathSelectorReceivedCounter})
	}
	if o.GetLossRatio() != other.GetLossRatio() {
		res.Paths = append(res.Paths, &RunPingTestResponseSummaryStats_FieldTerminalPath{selector: RunPingTestResponseSummaryStats_FieldPathSelectorLossRatio})
	}
	return res
}

func (o *RunPingTestResponse_SummaryStats) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunPingTestResponse_SummaryStats))
}

func (o *RunPingTestResponse_SummaryStats) Clone() *RunPingTestResponse_SummaryStats {
	if o == nil {
		return nil
	}
	result := &RunPingTestResponse_SummaryStats{}
	result.MinRtt = proto.Clone(o.MinRtt).(*duration.Duration)
	result.AvgRtt = proto.Clone(o.AvgRtt).(*duration.Duration)
	result.MaxRtt = proto.Clone(o.MaxRtt).(*duration.Duration)
	result.StddevRtt = proto.Clone(o.StddevRtt).(*duration.Duration)
	result.TransmittedCounter = o.TransmittedCounter
	result.ReceivedCounter = o.ReceivedCounter
	result.LossRatio = o.LossRatio
	return result
}

func (o *RunPingTestResponse_SummaryStats) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunPingTestResponse_SummaryStats) Merge(source *RunPingTestResponse_SummaryStats) {
	if source.GetMinRtt() != nil {
		if o.MinRtt == nil {
			o.MinRtt = new(duration.Duration)
		}
		proto.Merge(o.MinRtt, source.GetMinRtt())
	}
	if source.GetAvgRtt() != nil {
		if o.AvgRtt == nil {
			o.AvgRtt = new(duration.Duration)
		}
		proto.Merge(o.AvgRtt, source.GetAvgRtt())
	}
	if source.GetMaxRtt() != nil {
		if o.MaxRtt == nil {
			o.MaxRtt = new(duration.Duration)
		}
		proto.Merge(o.MaxRtt, source.GetMaxRtt())
	}
	if source.GetStddevRtt() != nil {
		if o.StddevRtt == nil {
			o.StddevRtt = new(duration.Duration)
		}
		proto.Merge(o.StddevRtt, source.GetStddevRtt())
	}
	o.TransmittedCounter = source.GetTransmittedCounter()
	o.ReceivedCounter = source.GetReceivedCounter()
	o.LossRatio = source.GetLossRatio()
}

func (o *RunPingTestResponse_SummaryStats) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunPingTestResponse_SummaryStats))
}

func (o *RunPingTestRequestToProbe) GotenObjectExt() {}

func (o *RunPingTestRequestToProbe) MakeFullFieldMask() *RunPingTestRequestToProbe_FieldMask {
	return FullRunPingTestRequestToProbe_FieldMask()
}

func (o *RunPingTestRequestToProbe) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunPingTestRequestToProbe_FieldMask()
}

func (o *RunPingTestRequestToProbe) MakeDiffFieldMask(other *RunPingTestRequestToProbe) *RunPingTestRequestToProbe_FieldMask {
	if o == nil && other == nil {
		return &RunPingTestRequestToProbe_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunPingTestRequestToProbe_FieldMask()
	}

	res := &RunPingTestRequestToProbe_FieldMask{}
	if o.GetSource() != other.GetSource() {
		res.Paths = append(res.Paths, &RunPingTestRequestToProbe_FieldTerminalPath{selector: RunPingTestRequestToProbe_FieldPathSelectorSource})
	}
	if o.GetDestination() != other.GetDestination() {
		res.Paths = append(res.Paths, &RunPingTestRequestToProbe_FieldTerminalPath{selector: RunPingTestRequestToProbe_FieldPathSelectorDestination})
	}
	if o.GetSizeBytes() != other.GetSizeBytes() {
		res.Paths = append(res.Paths, &RunPingTestRequestToProbe_FieldTerminalPath{selector: RunPingTestRequestToProbe_FieldPathSelectorSizeBytes})
	}
	if o.GetCount() != other.GetCount() {
		res.Paths = append(res.Paths, &RunPingTestRequestToProbe_FieldTerminalPath{selector: RunPingTestRequestToProbe_FieldPathSelectorCount})
	}
	if !proto.Equal(o.GetInterval(), other.GetInterval()) {
		res.Paths = append(res.Paths, &RunPingTestRequestToProbe_FieldTerminalPath{selector: RunPingTestRequestToProbe_FieldPathSelectorInterval})
	}
	if !proto.Equal(o.GetEchoTimeout(), other.GetEchoTimeout()) {
		res.Paths = append(res.Paths, &RunPingTestRequestToProbe_FieldTerminalPath{selector: RunPingTestRequestToProbe_FieldPathSelectorEchoTimeout})
	}
	if o.GetDontFragment() != other.GetDontFragment() {
		res.Paths = append(res.Paths, &RunPingTestRequestToProbe_FieldTerminalPath{selector: RunPingTestRequestToProbe_FieldPathSelectorDontFragment})
	}
	if o.GetTtl() != other.GetTtl() {
		res.Paths = append(res.Paths, &RunPingTestRequestToProbe_FieldTerminalPath{selector: RunPingTestRequestToProbe_FieldPathSelectorTtl})
	}
	if o.GetTos() != other.GetTos() {
		res.Paths = append(res.Paths, &RunPingTestRequestToProbe_FieldTerminalPath{selector: RunPingTestRequestToProbe_FieldPathSelectorTos})
	}
	return res
}

func (o *RunPingTestRequestToProbe) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunPingTestRequestToProbe))
}

func (o *RunPingTestRequestToProbe) Clone() *RunPingTestRequestToProbe {
	if o == nil {
		return nil
	}
	result := &RunPingTestRequestToProbe{}
	result.Source = o.Source
	result.Destination = o.Destination
	result.SizeBytes = o.SizeBytes
	result.Count = o.Count
	result.Interval = proto.Clone(o.Interval).(*duration.Duration)
	result.EchoTimeout = proto.Clone(o.EchoTimeout).(*duration.Duration)
	result.DontFragment = o.DontFragment
	result.Ttl = o.Ttl
	result.Tos = o.Tos
	return result
}

func (o *RunPingTestRequestToProbe) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunPingTestRequestToProbe) Merge(source *RunPingTestRequestToProbe) {
	o.Source = source.GetSource()
	o.Destination = source.GetDestination()
	o.SizeBytes = source.GetSizeBytes()
	o.Count = source.GetCount()
	if source.GetInterval() != nil {
		if o.Interval == nil {
			o.Interval = new(duration.Duration)
		}
		proto.Merge(o.Interval, source.GetInterval())
	}
	if source.GetEchoTimeout() != nil {
		if o.EchoTimeout == nil {
			o.EchoTimeout = new(duration.Duration)
		}
		proto.Merge(o.EchoTimeout, source.GetEchoTimeout())
	}
	o.DontFragment = source.GetDontFragment()
	o.Ttl = source.GetTtl()
	o.Tos = source.GetTos()
}

func (o *RunPingTestRequestToProbe) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunPingTestRequestToProbe))
}
