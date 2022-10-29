// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/internet_quality_rating.proto
// DO NOT EDIT!!!

package internet_quality_rating

import (
	"encoding/json"
	"strings"

	firestorepb "google.golang.org/genproto/googleapis/firestore/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
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
	_ = new(json.Marshaler)
	_ = strings.Builder{}

	_ = firestorepb.Value{}
	_ = codes.NotFound
	_ = status.Status{}
	_ = new(proto.Message)
	_ = new(preflect.Message)
	_ = fieldmaskpb.FieldMask{}

	_ = new(gotenobject.FieldMask)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &probe.Probe{}
)

type InternetQualityRating_FieldMask struct {
	Paths []InternetQualityRating_FieldPath
}

func FullInternetQualityRating_FieldMask() *InternetQualityRating_FieldMask {
	res := &InternetQualityRating_FieldMask{}
	res.Paths = append(res.Paths, &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorName})
	res.Paths = append(res.Paths, &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorRating})
	res.Paths = append(res.Paths, &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorTimestamp})
	res.Paths = append(res.Paths, &InternetQualityRating_FieldTerminalPath{selector: InternetQualityRating_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *InternetQualityRating_FieldMask) String() string {
	if fieldMask == nil {
		return "<nil>"
	}
	pathsStr := make([]string, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		pathsStr = append(pathsStr, path.String())
	}
	return strings.Join(pathsStr, ", ")
}

// firestore encoding/decoding integration
func (fieldMask *InternetQualityRating_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
	if fieldMask == nil {
		return &firestorepb.Value{ValueType: &firestorepb.Value_NullValue{}}, nil
	}
	arrayValues := make([]*firestorepb.Value, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.GetPaths() {
		arrayValues = append(arrayValues, &firestorepb.Value{ValueType: &firestorepb.Value_StringValue{StringValue: path.String()}})
	}
	return &firestorepb.Value{
		ValueType: &firestorepb.Value_ArrayValue{ArrayValue: &firestorepb.ArrayValue{Values: arrayValues}},
	}, nil
}

func (fieldMask *InternetQualityRating_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseInternetQualityRating_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *InternetQualityRating_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 4)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*InternetQualityRating_FieldTerminalPath); ok {
			presentSelectors[int(asFinal.selector)] = true
		}
	}
	for _, flag := range presentSelectors {
		if !flag {
			return false
		}
	}
	return true
}

func (fieldMask *InternetQualityRating_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseInternetQualityRating_FieldPath(raw)
	})
}

func (fieldMask *InternetQualityRating_FieldMask) ProtoMessage() {}

func (fieldMask *InternetQualityRating_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *InternetQualityRating_FieldMask) Subtract(other *InternetQualityRating_FieldMask) *InternetQualityRating_FieldMask {
	result := &InternetQualityRating_FieldMask{}
	removedSelectors := make([]bool, 4)
	otherSubMasks := map[InternetQualityRating_FieldPathSelector]gotenobject.FieldMask{
		InternetQualityRating_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[InternetQualityRating_FieldPathSelector]gotenobject.FieldMask{
		InternetQualityRating_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *InternetQualityRating_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *InternetQualityRating_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*InternetQualityRating_FieldTerminalPath); ok {
					switch tp.selector {
					case InternetQualityRating_FieldPathSelectorMetadata:
						mySubMasks[InternetQualityRating_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*InternetQualityRating_FieldSubPath); ok {
					mySubMasks[tp.selector].AppendRawPath(tp.subPath)
				}
			} else {
				result.Paths = append(result.Paths, path)
			}
		}
	}
	for selector, mySubMask := range mySubMasks {
		if mySubMask.PathsCount() > 0 {
			for _, allowedPath := range mySubMask.SubtractRaw(otherSubMasks[selector]).GetRawPaths() {
				result.Paths = append(result.Paths, &InternetQualityRating_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *InternetQualityRating_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*InternetQualityRating_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *InternetQualityRating_FieldMask) FilterInputFields() *InternetQualityRating_FieldMask {
	result := &InternetQualityRating_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case InternetQualityRating_FieldPathSelectorMetadata:
			if _, ok := path.(*InternetQualityRating_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &InternetQualityRating_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*InternetQualityRating_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &InternetQualityRating_FieldSubPath{selector: InternetQualityRating_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *InternetQualityRating_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *InternetQualityRating_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]InternetQualityRating_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseInternetQualityRating_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask InternetQualityRating_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *InternetQualityRating_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *InternetQualityRating_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask InternetQualityRating_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *InternetQualityRating_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *InternetQualityRating_FieldMask) AppendPath(path InternetQualityRating_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *InternetQualityRating_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(InternetQualityRating_FieldPath))
}

func (fieldMask *InternetQualityRating_FieldMask) GetPaths() []InternetQualityRating_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *InternetQualityRating_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *InternetQualityRating_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseInternetQualityRating_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *InternetQualityRating_FieldMask) Set(target, source *InternetQualityRating) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *InternetQualityRating_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*InternetQualityRating), source.(*InternetQualityRating))
}

func (fieldMask *InternetQualityRating_FieldMask) Project(source *InternetQualityRating) *InternetQualityRating {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &InternetQualityRating{}
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *InternetQualityRating_FieldTerminalPath:
			switch tp.selector {
			case InternetQualityRating_FieldPathSelectorName:
				result.Name = source.Name
			case InternetQualityRating_FieldPathSelectorRating:
				result.Rating = source.Rating
			case InternetQualityRating_FieldPathSelectorTimestamp:
				result.Timestamp = source.Timestamp
			case InternetQualityRating_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *InternetQualityRating_FieldSubPath:
			switch tp.selector {
			case InternetQualityRating_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *InternetQualityRating_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*InternetQualityRating))
}

func (fieldMask *InternetQualityRating_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
