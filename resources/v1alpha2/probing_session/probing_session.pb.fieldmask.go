// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/probing_session.proto
// DO NOT EDIT!!!

package probing_session

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
	_ = &duration.Duration{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_distribution.ProbingDistribution{}
	_ = &probing_target.ProbingTarget{}
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &project.Project{}
)

type ProbingSession_FieldMask struct {
	Paths []ProbingSession_FieldPath
}

func FullProbingSession_FieldMask() *ProbingSession_FieldMask {
	res := &ProbingSession_FieldMask{}
	res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorName})
	res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorSpec})
	res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorProbingDistribution})
	res.Paths = append(res.Paths, &ProbingSession_FieldTerminalPath{selector: ProbingSession_FieldPathSelectorStatus})
	return res
}

func (fieldMask *ProbingSession_FieldMask) String() string {
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
func (fieldMask *ProbingSession_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProbingSession_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProbingSession_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProbingSession_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 6)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProbingSession_FieldTerminalPath); ok {
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

func (fieldMask *ProbingSession_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProbingSession_FieldPath(raw)
	})
}

func (fieldMask *ProbingSession_FieldMask) ProtoMessage() {}

func (fieldMask *ProbingSession_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProbingSession_FieldMask) Subtract(other *ProbingSession_FieldMask) *ProbingSession_FieldMask {
	result := &ProbingSession_FieldMask{}
	removedSelectors := make([]bool, 6)
	otherSubMasks := map[ProbingSession_FieldPathSelector]gotenobject.FieldMask{
		ProbingSession_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
		ProbingSession_FieldPathSelectorSpec:     &ProbingSession_Spec_FieldMask{},
		ProbingSession_FieldPathSelectorStatus:   &ProbingSession_Status_FieldMask{},
	}
	mySubMasks := map[ProbingSession_FieldPathSelector]gotenobject.FieldMask{
		ProbingSession_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
		ProbingSession_FieldPathSelectorSpec:     &ProbingSession_Spec_FieldMask{},
		ProbingSession_FieldPathSelectorStatus:   &ProbingSession_Status_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProbingSession_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ProbingSession_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ProbingSession_FieldTerminalPath); ok {
					switch tp.selector {
					case ProbingSession_FieldPathSelectorMetadata:
						mySubMasks[ProbingSession_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					case ProbingSession_FieldPathSelectorSpec:
						mySubMasks[ProbingSession_FieldPathSelectorSpec] = FullProbingSession_Spec_FieldMask()
					case ProbingSession_FieldPathSelectorStatus:
						mySubMasks[ProbingSession_FieldPathSelectorStatus] = FullProbingSession_Status_FieldMask()
					}
				} else if tp, ok := path.(*ProbingSession_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ProbingSession_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ProbingSession_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProbingSession_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProbingSession_FieldMask) FilterInputFields() *ProbingSession_FieldMask {
	result := &ProbingSession_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case ProbingSession_FieldPathSelectorMetadata:
			if _, ok := path.(*ProbingSession_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProbingSession_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*ProbingSession_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProbingSession_FieldSubPath{selector: ProbingSession_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProbingSession_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProbingSession_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProbingSession_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProbingSession_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProbingSession_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProbingSession_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingSession_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProbingSession_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProbingSession_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingSession_FieldMask) AppendPath(path ProbingSession_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProbingSession_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProbingSession_FieldPath))
}

func (fieldMask *ProbingSession_FieldMask) GetPaths() []ProbingSession_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProbingSession_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProbingSession_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProbingSession_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProbingSession_FieldMask) Set(target, source *ProbingSession) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProbingSession_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProbingSession), source.(*ProbingSession))
}

func (fieldMask *ProbingSession_FieldMask) Project(source *ProbingSession) *ProbingSession {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProbingSession{}
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	specMask := &ProbingSession_Spec_FieldMask{}
	wholeSpecAccepted := false
	statusMask := &ProbingSession_Status_FieldMask{}
	wholeStatusAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProbingSession_FieldTerminalPath:
			switch tp.selector {
			case ProbingSession_FieldPathSelectorName:
				result.Name = source.Name
			case ProbingSession_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case ProbingSession_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case ProbingSession_FieldPathSelectorSpec:
				result.Spec = source.Spec
				wholeSpecAccepted = true
			case ProbingSession_FieldPathSelectorProbingDistribution:
				result.ProbingDistribution = source.ProbingDistribution
			case ProbingSession_FieldPathSelectorStatus:
				result.Status = source.Status
				wholeStatusAccepted = true
			}
		case *ProbingSession_FieldSubPath:
			switch tp.selector {
			case ProbingSession_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			case ProbingSession_FieldPathSelectorSpec:
				specMask.AppendPath(tp.subPath.(ProbingSessionSpec_FieldPath))
			case ProbingSession_FieldPathSelectorStatus:
				statusMask.AppendPath(tp.subPath.(ProbingSessionStatus_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeSpecAccepted == false && len(specMask.Paths) > 0 {
		result.Spec = specMask.Project(source.GetSpec())
	}
	if wholeStatusAccepted == false && len(statusMask.Paths) > 0 {
		result.Status = statusMask.Project(source.GetStatus())
	}
	return result
}

func (fieldMask *ProbingSession_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProbingSession))
}

func (fieldMask *ProbingSession_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type ProbingSession_Spec_FieldMask struct {
	Paths []ProbingSessionSpec_FieldPath
}

func FullProbingSession_Spec_FieldMask() *ProbingSession_Spec_FieldMask {
	res := &ProbingSession_Spec_FieldMask{}
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorProbe})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorTarget})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorTargetGroup})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorMode})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorIpVersion})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorAddress})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorAddresses})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorPort})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorType})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorInterval})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorTos})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorPathProbing})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorSpeedtestSettings})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorHttpProbingConfig})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorProxyConfiguration})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorLocationType})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorLocation})
	res.Paths = append(res.Paths, &ProbingSessionSpec_FieldTerminalPath{selector: ProbingSessionSpec_FieldPathSelectorEnablePcap})
	return res
}

func (fieldMask *ProbingSession_Spec_FieldMask) String() string {
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
func (fieldMask *ProbingSession_Spec_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProbingSession_Spec_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProbingSessionSpec_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProbingSession_Spec_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 18)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProbingSessionSpec_FieldTerminalPath); ok {
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

func (fieldMask *ProbingSession_Spec_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProbingSessionSpec_FieldPath(raw)
	})
}

func (fieldMask *ProbingSession_Spec_FieldMask) ProtoMessage() {}

func (fieldMask *ProbingSession_Spec_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProbingSession_Spec_FieldMask) Subtract(other *ProbingSession_Spec_FieldMask) *ProbingSession_Spec_FieldMask {
	result := &ProbingSession_Spec_FieldMask{}
	removedSelectors := make([]bool, 18)
	otherSubMasks := map[ProbingSessionSpec_FieldPathSelector]gotenobject.FieldMask{
		ProbingSessionSpec_FieldPathSelectorPathProbing:        &common.PathProbe_FieldMask{},
		ProbingSessionSpec_FieldPathSelectorSpeedtestSettings:  &common.SpeedTestSettings_FieldMask{},
		ProbingSessionSpec_FieldPathSelectorHttpProbingConfig:  &common.HTTPProbingConfig_FieldMask{},
		ProbingSessionSpec_FieldPathSelectorProxyConfiguration: &common.ProxyConfiguration_FieldMask{},
		ProbingSessionSpec_FieldPathSelectorLocation:           &common.Location_FieldMask{},
	}
	mySubMasks := map[ProbingSessionSpec_FieldPathSelector]gotenobject.FieldMask{
		ProbingSessionSpec_FieldPathSelectorPathProbing:        &common.PathProbe_FieldMask{},
		ProbingSessionSpec_FieldPathSelectorSpeedtestSettings:  &common.SpeedTestSettings_FieldMask{},
		ProbingSessionSpec_FieldPathSelectorHttpProbingConfig:  &common.HTTPProbingConfig_FieldMask{},
		ProbingSessionSpec_FieldPathSelectorProxyConfiguration: &common.ProxyConfiguration_FieldMask{},
		ProbingSessionSpec_FieldPathSelectorLocation:           &common.Location_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProbingSessionSpec_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ProbingSessionSpec_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ProbingSessionSpec_FieldTerminalPath); ok {
					switch tp.selector {
					case ProbingSessionSpec_FieldPathSelectorPathProbing:
						mySubMasks[ProbingSessionSpec_FieldPathSelectorPathProbing] = common.FullPathProbe_FieldMask()
					case ProbingSessionSpec_FieldPathSelectorSpeedtestSettings:
						mySubMasks[ProbingSessionSpec_FieldPathSelectorSpeedtestSettings] = common.FullSpeedTestSettings_FieldMask()
					case ProbingSessionSpec_FieldPathSelectorHttpProbingConfig:
						mySubMasks[ProbingSessionSpec_FieldPathSelectorHttpProbingConfig] = common.FullHTTPProbingConfig_FieldMask()
					case ProbingSessionSpec_FieldPathSelectorProxyConfiguration:
						mySubMasks[ProbingSessionSpec_FieldPathSelectorProxyConfiguration] = common.FullProxyConfiguration_FieldMask()
					case ProbingSessionSpec_FieldPathSelectorLocation:
						mySubMasks[ProbingSessionSpec_FieldPathSelectorLocation] = common.FullLocation_FieldMask()
					}
				} else if tp, ok := path.(*ProbingSessionSpec_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ProbingSessionSpec_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ProbingSession_Spec_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProbingSession_Spec_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProbingSession_Spec_FieldMask) FilterInputFields() *ProbingSession_Spec_FieldMask {
	result := &ProbingSession_Spec_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProbingSession_Spec_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProbingSession_Spec_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProbingSessionSpec_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProbingSessionSpec_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProbingSession_Spec_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProbingSession_Spec_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingSession_Spec_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProbingSession_Spec_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProbingSession_Spec_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingSession_Spec_FieldMask) AppendPath(path ProbingSessionSpec_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProbingSession_Spec_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProbingSessionSpec_FieldPath))
}

func (fieldMask *ProbingSession_Spec_FieldMask) GetPaths() []ProbingSessionSpec_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProbingSession_Spec_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProbingSession_Spec_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProbingSessionSpec_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProbingSession_Spec_FieldMask) Set(target, source *ProbingSession_Spec) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProbingSession_Spec_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProbingSession_Spec), source.(*ProbingSession_Spec))
}

func (fieldMask *ProbingSession_Spec_FieldMask) Project(source *ProbingSession_Spec) *ProbingSession_Spec {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProbingSession_Spec{}
	pathProbingMask := &common.PathProbe_FieldMask{}
	wholePathProbingAccepted := false
	speedtestSettingsMask := &common.SpeedTestSettings_FieldMask{}
	wholeSpeedtestSettingsAccepted := false
	httpProbingConfigMask := &common.HTTPProbingConfig_FieldMask{}
	wholeHttpProbingConfigAccepted := false
	proxyConfigurationMask := &common.ProxyConfiguration_FieldMask{}
	wholeProxyConfigurationAccepted := false
	locationMask := &common.Location_FieldMask{}
	wholeLocationAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProbingSessionSpec_FieldTerminalPath:
			switch tp.selector {
			case ProbingSessionSpec_FieldPathSelectorProbe:
				result.Probe = source.Probe
			case ProbingSessionSpec_FieldPathSelectorTarget:
				result.Target = source.Target
			case ProbingSessionSpec_FieldPathSelectorTargetGroup:
				result.TargetGroup = source.TargetGroup
			case ProbingSessionSpec_FieldPathSelectorMode:
				result.Mode = source.Mode
			case ProbingSessionSpec_FieldPathSelectorIpVersion:
				result.IpVersion = source.IpVersion
			case ProbingSessionSpec_FieldPathSelectorAddress:
				result.Address = source.Address
			case ProbingSessionSpec_FieldPathSelectorAddresses:
				result.Addresses = source.Addresses
			case ProbingSessionSpec_FieldPathSelectorPort:
				result.Port = source.Port
			case ProbingSessionSpec_FieldPathSelectorType:
				result.Type = source.Type
			case ProbingSessionSpec_FieldPathSelectorInterval:
				result.Interval = source.Interval
			case ProbingSessionSpec_FieldPathSelectorTos:
				result.Tos = source.Tos
			case ProbingSessionSpec_FieldPathSelectorPathProbing:
				result.PathProbing = source.PathProbing
				wholePathProbingAccepted = true
			case ProbingSessionSpec_FieldPathSelectorSpeedtestSettings:
				result.SpeedtestSettings = source.SpeedtestSettings
				wholeSpeedtestSettingsAccepted = true
			case ProbingSessionSpec_FieldPathSelectorHttpProbingConfig:
				result.HttpProbingConfig = source.HttpProbingConfig
				wholeHttpProbingConfigAccepted = true
			case ProbingSessionSpec_FieldPathSelectorProxyConfiguration:
				result.ProxyConfiguration = source.ProxyConfiguration
				wholeProxyConfigurationAccepted = true
			case ProbingSessionSpec_FieldPathSelectorLocationType:
				result.LocationType = source.LocationType
			case ProbingSessionSpec_FieldPathSelectorLocation:
				result.Location = source.Location
				wholeLocationAccepted = true
			case ProbingSessionSpec_FieldPathSelectorEnablePcap:
				result.EnablePcap = source.EnablePcap
			}
		case *ProbingSessionSpec_FieldSubPath:
			switch tp.selector {
			case ProbingSessionSpec_FieldPathSelectorPathProbing:
				pathProbingMask.AppendPath(tp.subPath.(common.PathProbe_FieldPath))
			case ProbingSessionSpec_FieldPathSelectorSpeedtestSettings:
				speedtestSettingsMask.AppendPath(tp.subPath.(common.SpeedTestSettings_FieldPath))
			case ProbingSessionSpec_FieldPathSelectorHttpProbingConfig:
				httpProbingConfigMask.AppendPath(tp.subPath.(common.HTTPProbingConfig_FieldPath))
			case ProbingSessionSpec_FieldPathSelectorProxyConfiguration:
				proxyConfigurationMask.AppendPath(tp.subPath.(common.ProxyConfiguration_FieldPath))
			case ProbingSessionSpec_FieldPathSelectorLocation:
				locationMask.AppendPath(tp.subPath.(common.Location_FieldPath))
			}
		}
	}
	if wholePathProbingAccepted == false && len(pathProbingMask.Paths) > 0 {
		result.PathProbing = pathProbingMask.Project(source.GetPathProbing())
	}
	if wholeSpeedtestSettingsAccepted == false && len(speedtestSettingsMask.Paths) > 0 {
		result.SpeedtestSettings = speedtestSettingsMask.Project(source.GetSpeedtestSettings())
	}
	if wholeHttpProbingConfigAccepted == false && len(httpProbingConfigMask.Paths) > 0 {
		result.HttpProbingConfig = httpProbingConfigMask.Project(source.GetHttpProbingConfig())
	}
	if wholeProxyConfigurationAccepted == false && len(proxyConfigurationMask.Paths) > 0 {
		result.ProxyConfiguration = proxyConfigurationMask.Project(source.GetProxyConfiguration())
	}
	if wholeLocationAccepted == false && len(locationMask.Paths) > 0 {
		result.Location = locationMask.Project(source.GetLocation())
	}
	return result
}

func (fieldMask *ProbingSession_Spec_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProbingSession_Spec))
}

func (fieldMask *ProbingSession_Spec_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type ProbingSession_Status_FieldMask struct {
	Paths []ProbingSessionStatus_FieldPath
}

func FullProbingSession_Status_FieldMask() *ProbingSession_Status_FieldMask {
	res := &ProbingSession_Status_FieldMask{}
	return res
}

func (fieldMask *ProbingSession_Status_FieldMask) String() string {
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
func (fieldMask *ProbingSession_Status_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProbingSession_Status_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProbingSessionStatus_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProbingSession_Status_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 0)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProbingSessionStatus_FieldTerminalPath); ok {
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

func (fieldMask *ProbingSession_Status_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProbingSessionStatus_FieldPath(raw)
	})
}

func (fieldMask *ProbingSession_Status_FieldMask) ProtoMessage() {}

func (fieldMask *ProbingSession_Status_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProbingSession_Status_FieldMask) Subtract(other *ProbingSession_Status_FieldMask) *ProbingSession_Status_FieldMask {
	result := &ProbingSession_Status_FieldMask{}
	removedSelectors := make([]bool, 0)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProbingSessionStatus_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			result.Paths = append(result.Paths, path)
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ProbingSession_Status_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProbingSession_Status_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProbingSession_Status_FieldMask) FilterInputFields() *ProbingSession_Status_FieldMask {
	result := &ProbingSession_Status_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProbingSession_Status_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProbingSession_Status_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProbingSessionStatus_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProbingSessionStatus_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProbingSession_Status_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProbingSession_Status_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingSession_Status_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProbingSession_Status_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProbingSession_Status_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingSession_Status_FieldMask) AppendPath(path ProbingSessionStatus_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProbingSession_Status_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProbingSessionStatus_FieldPath))
}

func (fieldMask *ProbingSession_Status_FieldMask) GetPaths() []ProbingSessionStatus_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProbingSession_Status_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProbingSession_Status_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProbingSessionStatus_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProbingSession_Status_FieldMask) Set(target, source *ProbingSession_Status) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProbingSession_Status_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProbingSession_Status), source.(*ProbingSession_Status))
}

func (fieldMask *ProbingSession_Status_FieldMask) Project(source *ProbingSession_Status) *ProbingSession_Status {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProbingSession_Status{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProbingSessionStatus_FieldTerminalPath:
			switch tp.selector {
			}
		}
	}
	return result
}

func (fieldMask *ProbingSession_Status_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProbingSession_Status))
}

func (fieldMask *ProbingSession_Status_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
