// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/dns_query_test_custom.proto
// DO NOT EDIT!!!

package dns_query_test_client

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
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
)

func (o *RunDNSQueryTestRequest) GotenObjectExt() {}

func (o *RunDNSQueryTestRequest) MakeFullFieldMask() *RunDNSQueryTestRequest_FieldMask {
	return FullRunDNSQueryTestRequest_FieldMask()
}

func (o *RunDNSQueryTestRequest) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunDNSQueryTestRequest_FieldMask()
}

func (o *RunDNSQueryTestRequest) MakeDiffFieldMask(other *RunDNSQueryTestRequest) *RunDNSQueryTestRequest_FieldMask {
	if o == nil && other == nil {
		return &RunDNSQueryTestRequest_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunDNSQueryTestRequest_FieldMask()
	}

	res := &RunDNSQueryTestRequest_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequest_FieldTerminalPath{selector: RunDNSQueryTestRequest_FieldPathSelectorName})
	}
	{
		subMask := o.GetQuery().MakeDiffFieldMask(other.GetQuery())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &RunDNSQueryTestRequest_FieldTerminalPath{selector: RunDNSQueryTestRequest_FieldPathSelectorQuery})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &RunDNSQueryTestRequest_FieldSubPath{selector: RunDNSQueryTestRequest_FieldPathSelectorQuery, subPath: subpath})
			}
		}
	}
	if o.GetServer() != other.GetServer() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequest_FieldTerminalPath{selector: RunDNSQueryTestRequest_FieldPathSelectorServer})
	}
	if o.GetPort() != other.GetPort() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequest_FieldTerminalPath{selector: RunDNSQueryTestRequest_FieldPathSelectorPort})
	}
	if o.GetTcp() != other.GetTcp() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequest_FieldTerminalPath{selector: RunDNSQueryTestRequest_FieldPathSelectorTcp})
	}
	if o.GetNoRecursionDesired() != other.GetNoRecursionDesired() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequest_FieldTerminalPath{selector: RunDNSQueryTestRequest_FieldPathSelectorNoRecursionDesired})
	}
	if o.GetReverse() != other.GetReverse() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequest_FieldTerminalPath{selector: RunDNSQueryTestRequest_FieldPathSelectorReverse})
	}
	return res
}

func (o *RunDNSQueryTestRequest) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunDNSQueryTestRequest))
}

func (o *RunDNSQueryTestRequest) Clone() *RunDNSQueryTestRequest {
	if o == nil {
		return nil
	}
	result := &RunDNSQueryTestRequest{}
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
	result.Query = o.Query.Clone()
	result.Server = o.Server
	result.Port = o.Port
	result.Tcp = o.Tcp
	result.NoRecursionDesired = o.NoRecursionDesired
	result.Reverse = o.Reverse
	return result
}

func (o *RunDNSQueryTestRequest) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunDNSQueryTestRequest) Merge(source *RunDNSQueryTestRequest) {
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
	if source.GetQuery() != nil {
		if o.Query == nil {
			o.Query = new(common.DNSQuery)
		}
		o.Query.Merge(source.GetQuery())
	}
	o.Server = source.GetServer()
	o.Port = source.GetPort()
	o.Tcp = source.GetTcp()
	o.NoRecursionDesired = source.GetNoRecursionDesired()
	o.Reverse = source.GetReverse()
}

func (o *RunDNSQueryTestRequest) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunDNSQueryTestRequest))
}

func (o *RunDNSQueryTestResponse) GotenObjectExt() {}

func (o *RunDNSQueryTestResponse) MakeFullFieldMask() *RunDNSQueryTestResponse_FieldMask {
	return FullRunDNSQueryTestResponse_FieldMask()
}

func (o *RunDNSQueryTestResponse) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunDNSQueryTestResponse_FieldMask()
}

func (o *RunDNSQueryTestResponse) MakeDiffFieldMask(other *RunDNSQueryTestResponse) *RunDNSQueryTestResponse_FieldMask {
	if o == nil && other == nil {
		return &RunDNSQueryTestResponse_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunDNSQueryTestResponse_FieldMask()
	}

	res := &RunDNSQueryTestResponse_FieldMask{}
	if o.GetId() != other.GetId() {
		res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorId})
	}
	if o.GetRcode() != other.GetRcode() {
		res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorRcode})
	}

	if len(o.GetFlags()) == len(other.GetFlags()) {
		for i, lValue := range o.GetFlags() {
			rValue := other.GetFlags()[i]
			if lValue != rValue {
				res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorFlags})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorFlags})
	}

	if len(o.GetQueries()) == len(other.GetQueries()) {
		for i, lValue := range o.GetQueries() {
			rValue := other.GetQueries()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorQueries})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorQueries})
	}

	if len(o.GetAnswers()) == len(other.GetAnswers()) {
		for i, lValue := range o.GetAnswers() {
			rValue := other.GetAnswers()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorAnswers})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorAnswers})
	}

	if len(o.GetNs()) == len(other.GetNs()) {
		for i, lValue := range o.GetNs() {
			rValue := other.GetNs()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorNs})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorNs})
	}

	if len(o.GetExtras()) == len(other.GetExtras()) {
		for i, lValue := range o.GetExtras() {
			rValue := other.GetExtras()[i]
			if len(lValue.MakeDiffFieldMask(rValue).Paths) > 0 {
				res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorExtras})
				break
			}
		}
	} else {
		res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorExtras})
	}
	if !proto.Equal(o.GetRtt(), other.GetRtt()) {
		res.Paths = append(res.Paths, &RunDNSQueryTestResponse_FieldTerminalPath{selector: RunDNSQueryTestResponse_FieldPathSelectorRtt})
	}
	return res
}

func (o *RunDNSQueryTestResponse) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunDNSQueryTestResponse))
}

func (o *RunDNSQueryTestResponse) Clone() *RunDNSQueryTestResponse {
	if o == nil {
		return nil
	}
	result := &RunDNSQueryTestResponse{}
	result.Id = o.Id
	result.Rcode = o.Rcode
	result.Flags = make([]string, len(o.Flags))
	for i, sourceValue := range o.Flags {
		result.Flags[i] = sourceValue
	}
	result.Queries = make([]*common.DNSQuery, len(o.Queries))
	for i, sourceValue := range o.Queries {
		result.Queries[i] = sourceValue.Clone()
	}
	result.Answers = make([]*common.DNSResourceRecord, len(o.Answers))
	for i, sourceValue := range o.Answers {
		result.Answers[i] = sourceValue.Clone()
	}
	result.Ns = make([]*common.DNSResourceRecord, len(o.Ns))
	for i, sourceValue := range o.Ns {
		result.Ns[i] = sourceValue.Clone()
	}
	result.Extras = make([]*common.DNSResourceRecord, len(o.Extras))
	for i, sourceValue := range o.Extras {
		result.Extras[i] = sourceValue.Clone()
	}
	result.Rtt = proto.Clone(o.Rtt).(*duration.Duration)
	return result
}

func (o *RunDNSQueryTestResponse) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunDNSQueryTestResponse) Merge(source *RunDNSQueryTestResponse) {
	o.Id = source.GetId()
	o.Rcode = source.GetRcode()
	for _, sourceValue := range source.GetFlags() {
		exists := false
		for _, currentValue := range o.Flags {
			if currentValue == sourceValue {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement string
			newDstElement = sourceValue
			o.Flags = append(o.Flags, newDstElement)
		}
	}

	for _, sourceValue := range source.GetQueries() {
		exists := false
		for _, currentValue := range o.Queries {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.DNSQuery
			if sourceValue != nil {
				newDstElement = new(common.DNSQuery)
				newDstElement.Merge(sourceValue)
			}
			o.Queries = append(o.Queries, newDstElement)
		}
	}

	for _, sourceValue := range source.GetAnswers() {
		exists := false
		for _, currentValue := range o.Answers {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.DNSResourceRecord
			if sourceValue != nil {
				newDstElement = new(common.DNSResourceRecord)
				newDstElement.Merge(sourceValue)
			}
			o.Answers = append(o.Answers, newDstElement)
		}
	}

	for _, sourceValue := range source.GetNs() {
		exists := false
		for _, currentValue := range o.Ns {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.DNSResourceRecord
			if sourceValue != nil {
				newDstElement = new(common.DNSResourceRecord)
				newDstElement.Merge(sourceValue)
			}
			o.Ns = append(o.Ns, newDstElement)
		}
	}

	for _, sourceValue := range source.GetExtras() {
		exists := false
		for _, currentValue := range o.Extras {
			if proto.Equal(sourceValue, currentValue) {
				exists = true
				break
			}
		}
		if !exists {
			var newDstElement *common.DNSResourceRecord
			if sourceValue != nil {
				newDstElement = new(common.DNSResourceRecord)
				newDstElement.Merge(sourceValue)
			}
			o.Extras = append(o.Extras, newDstElement)
		}
	}

	if source.GetRtt() != nil {
		if o.Rtt == nil {
			o.Rtt = new(duration.Duration)
		}
		proto.Merge(o.Rtt, source.GetRtt())
	}
}

func (o *RunDNSQueryTestResponse) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunDNSQueryTestResponse))
}

func (o *RunDNSQueryTestRequestToProbe) GotenObjectExt() {}

func (o *RunDNSQueryTestRequestToProbe) MakeFullFieldMask() *RunDNSQueryTestRequestToProbe_FieldMask {
	return FullRunDNSQueryTestRequestToProbe_FieldMask()
}

func (o *RunDNSQueryTestRequestToProbe) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullRunDNSQueryTestRequestToProbe_FieldMask()
}

func (o *RunDNSQueryTestRequestToProbe) MakeDiffFieldMask(other *RunDNSQueryTestRequestToProbe) *RunDNSQueryTestRequestToProbe_FieldMask {
	if o == nil && other == nil {
		return &RunDNSQueryTestRequestToProbe_FieldMask{}
	}
	if o == nil || other == nil {
		return FullRunDNSQueryTestRequestToProbe_FieldMask()
	}

	res := &RunDNSQueryTestRequestToProbe_FieldMask{}
	{
		subMask := o.GetQuery().MakeDiffFieldMask(other.GetQuery())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &RunDNSQueryTestRequestToProbe_FieldTerminalPath{selector: RunDNSQueryTestRequestToProbe_FieldPathSelectorQuery})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &RunDNSQueryTestRequestToProbe_FieldSubPath{selector: RunDNSQueryTestRequestToProbe_FieldPathSelectorQuery, subPath: subpath})
			}
		}
	}
	if o.GetServer() != other.GetServer() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequestToProbe_FieldTerminalPath{selector: RunDNSQueryTestRequestToProbe_FieldPathSelectorServer})
	}
	if o.GetPort() != other.GetPort() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequestToProbe_FieldTerminalPath{selector: RunDNSQueryTestRequestToProbe_FieldPathSelectorPort})
	}
	if o.GetTcp() != other.GetTcp() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequestToProbe_FieldTerminalPath{selector: RunDNSQueryTestRequestToProbe_FieldPathSelectorTcp})
	}
	if o.GetNoRecursionDesired() != other.GetNoRecursionDesired() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequestToProbe_FieldTerminalPath{selector: RunDNSQueryTestRequestToProbe_FieldPathSelectorNoRecursionDesired})
	}
	if o.GetReverse() != other.GetReverse() {
		res.Paths = append(res.Paths, &RunDNSQueryTestRequestToProbe_FieldTerminalPath{selector: RunDNSQueryTestRequestToProbe_FieldPathSelectorReverse})
	}
	return res
}

func (o *RunDNSQueryTestRequestToProbe) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*RunDNSQueryTestRequestToProbe))
}

func (o *RunDNSQueryTestRequestToProbe) Clone() *RunDNSQueryTestRequestToProbe {
	if o == nil {
		return nil
	}
	result := &RunDNSQueryTestRequestToProbe{}
	result.Query = o.Query.Clone()
	result.Server = o.Server
	result.Port = o.Port
	result.Tcp = o.Tcp
	result.NoRecursionDesired = o.NoRecursionDesired
	result.Reverse = o.Reverse
	return result
}

func (o *RunDNSQueryTestRequestToProbe) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *RunDNSQueryTestRequestToProbe) Merge(source *RunDNSQueryTestRequestToProbe) {
	if source.GetQuery() != nil {
		if o.Query == nil {
			o.Query = new(common.DNSQuery)
		}
		o.Query.Merge(source.GetQuery())
	}
	o.Server = source.GetServer()
	o.Port = source.GetPort()
	o.Tcp = source.GetTcp()
	o.NoRecursionDesired = source.GetNoRecursionDesired()
	o.Reverse = source.GetReverse()
}

func (o *RunDNSQueryTestRequestToProbe) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*RunDNSQueryTestRequestToProbe))
}