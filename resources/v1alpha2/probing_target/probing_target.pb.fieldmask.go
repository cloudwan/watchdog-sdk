// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/probing_target.proto
// DO NOT EDIT!!!

package probing_target

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
	probing_target_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
)

// ensure the imports are used
var (
	_ = json.Marshaler(nil)
	_ = strings.Builder{}

	_ = firestorepb.Value{}
	_ = codes.NotFound
	_ = status.Status{}
	_ = proto.Message(nil)
	_ = preflect.Message(nil)
	_ = fieldmaskpb.FieldMask{}

	_ = gotenobject.FieldMask(nil)
)

// make sure we're using proto imports
var (
	_ = &ntt_meta.Meta{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &project.Project{}
)

type ProbingTarget_FieldMask struct {
	Paths []ProbingTarget_FieldPath
}

func FullProbingTarget_FieldMask() *ProbingTarget_FieldMask {
	res := &ProbingTarget_FieldMask{}
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorName})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorGroup})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorMode})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorIpVersion})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorAddress})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorCategory})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorLocationType})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorLocation})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorProbingConstraint})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorDefaultProbingSettings})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorHttpProbingConfig})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorAgent})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorAddresses})
	res.Paths = append(res.Paths, &ProbingTarget_FieldTerminalPath{selector: ProbingTarget_FieldPathSelectorTargetType})
	return res
}

func (fieldMask *ProbingTarget_FieldMask) String() string {
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
func (fieldMask *ProbingTarget_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProbingTarget_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProbingTarget_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProbingTarget_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 16)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProbingTarget_FieldTerminalPath); ok {
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

func (fieldMask *ProbingTarget_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProbingTarget_FieldPath(raw)
	})
}

func (fieldMask *ProbingTarget_FieldMask) ProtoMessage() {}

func (fieldMask *ProbingTarget_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProbingTarget_FieldMask) Subtract(other *ProbingTarget_FieldMask) *ProbingTarget_FieldMask {
	result := &ProbingTarget_FieldMask{}
	removedSelectors := make([]bool, 16)
	otherSubMasks := map[ProbingTarget_FieldPathSelector]gotenobject.FieldMask{
		ProbingTarget_FieldPathSelectorMetadata:               &ntt_meta.Meta_FieldMask{},
		ProbingTarget_FieldPathSelectorLocation:               &common.Location_FieldMask{},
		ProbingTarget_FieldPathSelectorProbingConstraint:      &common.ProbingConstraint_FieldMask{},
		ProbingTarget_FieldPathSelectorDefaultProbingSettings: &common.ProbingSettings_FieldMask{},
		ProbingTarget_FieldPathSelectorHttpProbingConfig:      &common.HTTPProbingConfig_FieldMask{},
	}
	mySubMasks := map[ProbingTarget_FieldPathSelector]gotenobject.FieldMask{
		ProbingTarget_FieldPathSelectorMetadata:               &ntt_meta.Meta_FieldMask{},
		ProbingTarget_FieldPathSelectorLocation:               &common.Location_FieldMask{},
		ProbingTarget_FieldPathSelectorProbingConstraint:      &common.ProbingConstraint_FieldMask{},
		ProbingTarget_FieldPathSelectorDefaultProbingSettings: &common.ProbingSettings_FieldMask{},
		ProbingTarget_FieldPathSelectorHttpProbingConfig:      &common.HTTPProbingConfig_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProbingTarget_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ProbingTarget_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ProbingTarget_FieldTerminalPath); ok {
					switch tp.selector {
					case ProbingTarget_FieldPathSelectorMetadata:
						mySubMasks[ProbingTarget_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					case ProbingTarget_FieldPathSelectorLocation:
						mySubMasks[ProbingTarget_FieldPathSelectorLocation] = common.FullLocation_FieldMask()
					case ProbingTarget_FieldPathSelectorProbingConstraint:
						mySubMasks[ProbingTarget_FieldPathSelectorProbingConstraint] = common.FullProbingConstraint_FieldMask()
					case ProbingTarget_FieldPathSelectorDefaultProbingSettings:
						mySubMasks[ProbingTarget_FieldPathSelectorDefaultProbingSettings] = common.FullProbingSettings_FieldMask()
					case ProbingTarget_FieldPathSelectorHttpProbingConfig:
						mySubMasks[ProbingTarget_FieldPathSelectorHttpProbingConfig] = common.FullHTTPProbingConfig_FieldMask()
					}
				} else if tp, ok := path.(*ProbingTarget_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ProbingTarget_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ProbingTarget_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProbingTarget_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProbingTarget_FieldMask) FilterInputFields() *ProbingTarget_FieldMask {
	result := &ProbingTarget_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case ProbingTarget_FieldPathSelectorMetadata:
			if _, ok := path.(*ProbingTarget_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProbingTarget_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*ProbingTarget_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProbingTarget_FieldSubPath{selector: ProbingTarget_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProbingTarget_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProbingTarget_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProbingTarget_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProbingTarget_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProbingTarget_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProbingTarget_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingTarget_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProbingTarget_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProbingTarget_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingTarget_FieldMask) AppendPath(path ProbingTarget_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProbingTarget_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProbingTarget_FieldPath))
}

func (fieldMask *ProbingTarget_FieldMask) GetPaths() []ProbingTarget_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProbingTarget_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProbingTarget_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProbingTarget_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProbingTarget_FieldMask) Set(target, source *ProbingTarget) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProbingTarget_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProbingTarget), source.(*ProbingTarget))
}

func (fieldMask *ProbingTarget_FieldMask) Project(source *ProbingTarget) *ProbingTarget {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProbingTarget{}
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	locationMask := &common.Location_FieldMask{}
	wholeLocationAccepted := false
	probingConstraintMask := &common.ProbingConstraint_FieldMask{}
	wholeProbingConstraintAccepted := false
	defaultProbingSettingsMask := &common.ProbingSettings_FieldMask{}
	wholeDefaultProbingSettingsAccepted := false
	httpProbingConfigMask := &common.HTTPProbingConfig_FieldMask{}
	wholeHttpProbingConfigAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProbingTarget_FieldTerminalPath:
			switch tp.selector {
			case ProbingTarget_FieldPathSelectorName:
				result.Name = source.Name
			case ProbingTarget_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case ProbingTarget_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case ProbingTarget_FieldPathSelectorGroup:
				result.Group = source.Group
			case ProbingTarget_FieldPathSelectorMode:
				result.Mode = source.Mode
			case ProbingTarget_FieldPathSelectorIpVersion:
				result.IpVersion = source.IpVersion
			case ProbingTarget_FieldPathSelectorAddress:
				result.Address = source.Address
			case ProbingTarget_FieldPathSelectorCategory:
				result.Category = source.Category
			case ProbingTarget_FieldPathSelectorLocationType:
				result.LocationType = source.LocationType
			case ProbingTarget_FieldPathSelectorLocation:
				result.Location = source.Location
				wholeLocationAccepted = true
			case ProbingTarget_FieldPathSelectorProbingConstraint:
				result.ProbingConstraint = source.ProbingConstraint
				wholeProbingConstraintAccepted = true
			case ProbingTarget_FieldPathSelectorDefaultProbingSettings:
				result.DefaultProbingSettings = source.DefaultProbingSettings
				wholeDefaultProbingSettingsAccepted = true
			case ProbingTarget_FieldPathSelectorHttpProbingConfig:
				result.HttpProbingConfig = source.HttpProbingConfig
				wholeHttpProbingConfigAccepted = true
			case ProbingTarget_FieldPathSelectorAgent:
				result.Agent = source.Agent
			case ProbingTarget_FieldPathSelectorAddresses:
				result.Addresses = source.Addresses
			case ProbingTarget_FieldPathSelectorTargetType:
				result.TargetType = source.TargetType
			}
		case *ProbingTarget_FieldSubPath:
			switch tp.selector {
			case ProbingTarget_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			case ProbingTarget_FieldPathSelectorLocation:
				locationMask.AppendPath(tp.subPath.(common.Location_FieldPath))
			case ProbingTarget_FieldPathSelectorProbingConstraint:
				probingConstraintMask.AppendPath(tp.subPath.(common.ProbingConstraint_FieldPath))
			case ProbingTarget_FieldPathSelectorDefaultProbingSettings:
				defaultProbingSettingsMask.AppendPath(tp.subPath.(common.ProbingSettings_FieldPath))
			case ProbingTarget_FieldPathSelectorHttpProbingConfig:
				httpProbingConfigMask.AppendPath(tp.subPath.(common.HTTPProbingConfig_FieldPath))
			}
		}
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeLocationAccepted == false && len(locationMask.Paths) > 0 {
		result.Location = locationMask.Project(source.GetLocation())
	}
	if wholeProbingConstraintAccepted == false && len(probingConstraintMask.Paths) > 0 {
		result.ProbingConstraint = probingConstraintMask.Project(source.GetProbingConstraint())
	}
	if wholeDefaultProbingSettingsAccepted == false && len(defaultProbingSettingsMask.Paths) > 0 {
		result.DefaultProbingSettings = defaultProbingSettingsMask.Project(source.GetDefaultProbingSettings())
	}
	if wholeHttpProbingConfigAccepted == false && len(httpProbingConfigMask.Paths) > 0 {
		result.HttpProbingConfig = httpProbingConfigMask.Project(source.GetHttpProbingConfig())
	}
	return result
}

func (fieldMask *ProbingTarget_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProbingTarget))
}

func (fieldMask *ProbingTarget_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
