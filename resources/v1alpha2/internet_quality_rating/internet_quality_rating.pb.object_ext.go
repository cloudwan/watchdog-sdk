// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/internet_quality_rating.proto
// DO NOT EDIT!!!

package internet_quality_rating

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
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &probe.Probe{}
)

func (o *InternetQualityRating) GotenObjectExt() {}

func (o *InternetQualityRating) MakeFullFieldMask() *InternetQualityRating_FieldMask {
	return FullInternetQualityRating_FieldMask()
}

func (o *InternetQualityRating) MakeRawFullFieldMask() gotenobject.FieldMask {
	return FullInternetQualityRating_FieldMask()
}

func (o *InternetQualityRating) MakeDiffFieldMask(other *InternetQualityRating) *InternetQualityRating_FieldMask {
	if o == nil && other == nil {
		return &InternetQualityRating_FieldMask{}
	}
	if o == nil || other == nil {
		return FullInternetQualityRating_FieldMask()
	}

	res := &InternetQualityRating_FieldMask{}
	if o.GetName().String() != other.GetName().String() {
		res.Paths = append(res.Paths, &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorName})
	}
	if o.GetRating() != other.GetRating() {
		res.Paths = append(res.Paths, &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorRating})
	}
	if !proto.Equal(o.GetTimestamp(), other.GetTimestamp()) {
		res.Paths = append(res.Paths, &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorTimestamp})
	}
	{
		subMask := o.GetMetadata().MakeDiffFieldMask(other.GetMetadata())
		if subMask.IsFull() {
			res.Paths = append(res.Paths, &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorMetadata})
		} else {
			for _, subpath := range subMask.Paths {
				res.Paths = append(res.Paths, &InternetQualityRating_FieldSubPath{selector: InternetQualityRating_FieldPathSelectorMetadata, subPath: subpath})
			}
		}
	}
	return res
}

func (o *InternetQualityRating) MakeRawDiffFieldMask(other gotenobject.GotenObjectExt) gotenobject.FieldMask {
	return o.MakeDiffFieldMask(other.(*InternetQualityRating))
}

func (o *InternetQualityRating) Clone() *InternetQualityRating {
	if o == nil {
		return nil
	}
	result := &InternetQualityRating{}
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
	result.Rating = o.Rating
	result.Timestamp = proto.Clone(o.Timestamp).(*timestamp.Timestamp)
	result.Metadata = o.Metadata.Clone()
	return result
}

func (o *InternetQualityRating) CloneRaw() gotenobject.GotenObjectExt {
	return o.Clone()
}

func (o *InternetQualityRating) Merge(source *InternetQualityRating) {
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
	o.Rating = source.GetRating()
	if source.GetTimestamp() != nil {
		if o.Timestamp == nil {
			o.Timestamp = new(timestamp.Timestamp)
		}
		proto.Merge(o.Timestamp, source.GetTimestamp())
	}
	if source.GetMetadata() != nil {
		if o.Metadata == nil {
			o.Metadata = new(ntt_meta.Meta)
		}
		o.Metadata.Merge(source.GetMetadata())
	}
}

func (o *InternetQualityRating) MergeRaw(source gotenobject.GotenObjectExt) {
	o.Merge(source.(*InternetQualityRating))
}
