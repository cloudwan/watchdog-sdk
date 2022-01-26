// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha/probing_target_group.proto
// DO NOT EDIT!!!

package probing_target_group

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
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha/project"
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
	_ = &project.Project{}
)

type ProbingTargetGroup_FieldMask struct {
	Paths []ProbingTargetGroup_FieldPath
}

func FullProbingTargetGroup_FieldMask() *ProbingTargetGroup_FieldMask {
	res := &ProbingTargetGroup_FieldMask{}
	res.Paths = append(res.Paths, &ProbingTargetGroup_FieldTerminalPath{selector: ProbingTargetGroup_FieldPathSelectorName})
	res.Paths = append(res.Paths, &ProbingTargetGroup_FieldTerminalPath{selector: ProbingTargetGroup_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &ProbingTargetGroup_FieldTerminalPath{selector: ProbingTargetGroup_FieldPathSelectorState})
	res.Paths = append(res.Paths, &ProbingTargetGroup_FieldTerminalPath{selector: ProbingTargetGroup_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *ProbingTargetGroup_FieldMask) String() string {
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
func (fieldMask *ProbingTargetGroup_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProbingTargetGroup_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProbingTargetGroup_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProbingTargetGroup_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 4)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProbingTargetGroup_FieldTerminalPath); ok {
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

func (fieldMask *ProbingTargetGroup_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProbingTargetGroup_FieldPath(raw)
	})
}

func (fieldMask *ProbingTargetGroup_FieldMask) ProtoMessage() {}

func (fieldMask *ProbingTargetGroup_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProbingTargetGroup_FieldMask) Subtract(other *ProbingTargetGroup_FieldMask) *ProbingTargetGroup_FieldMask {
	result := &ProbingTargetGroup_FieldMask{}
	removedSelectors := make([]bool, 4)
	otherSubMasks := map[ProbingTargetGroup_FieldPathSelector]gotenobject.FieldMask{
		ProbingTargetGroup_FieldPathSelectorState:    &ProbingTargetGroup_State_FieldMask{},
		ProbingTargetGroup_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[ProbingTargetGroup_FieldPathSelector]gotenobject.FieldMask{
		ProbingTargetGroup_FieldPathSelectorState:    &ProbingTargetGroup_State_FieldMask{},
		ProbingTargetGroup_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProbingTargetGroup_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ProbingTargetGroup_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ProbingTargetGroup_FieldTerminalPath); ok {
					switch tp.selector {
					case ProbingTargetGroup_FieldPathSelectorState:
						mySubMasks[ProbingTargetGroup_FieldPathSelectorState] = FullProbingTargetGroup_State_FieldMask()
					case ProbingTargetGroup_FieldPathSelectorMetadata:
						mySubMasks[ProbingTargetGroup_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*ProbingTargetGroup_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ProbingTargetGroup_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ProbingTargetGroup_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProbingTargetGroup_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProbingTargetGroup_FieldMask) FilterInputFields() *ProbingTargetGroup_FieldMask {
	result := &ProbingTargetGroup_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case ProbingTargetGroup_FieldPathSelectorMetadata:
			if _, ok := path.(*ProbingTargetGroup_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProbingTargetGroup_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*ProbingTargetGroup_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProbingTargetGroup_FieldSubPath{selector: ProbingTargetGroup_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProbingTargetGroup_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProbingTargetGroup_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProbingTargetGroup_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProbingTargetGroup_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProbingTargetGroup_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProbingTargetGroup_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingTargetGroup_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProbingTargetGroup_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProbingTargetGroup_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingTargetGroup_FieldMask) AppendPath(path ProbingTargetGroup_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProbingTargetGroup_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProbingTargetGroup_FieldPath))
}

func (fieldMask *ProbingTargetGroup_FieldMask) GetPaths() []ProbingTargetGroup_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProbingTargetGroup_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProbingTargetGroup_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProbingTargetGroup_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProbingTargetGroup_FieldMask) Set(target, source *ProbingTargetGroup) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProbingTargetGroup_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProbingTargetGroup), source.(*ProbingTargetGroup))
}

func (fieldMask *ProbingTargetGroup_FieldMask) Project(source *ProbingTargetGroup) *ProbingTargetGroup {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProbingTargetGroup{}
	stateMask := &ProbingTargetGroup_State_FieldMask{}
	wholeStateAccepted := false
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProbingTargetGroup_FieldTerminalPath:
			switch tp.selector {
			case ProbingTargetGroup_FieldPathSelectorName:
				result.Name = source.Name
			case ProbingTargetGroup_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case ProbingTargetGroup_FieldPathSelectorState:
				result.State = source.State
				wholeStateAccepted = true
			case ProbingTargetGroup_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *ProbingTargetGroup_FieldSubPath:
			switch tp.selector {
			case ProbingTargetGroup_FieldPathSelectorState:
				stateMask.AppendPath(tp.subPath.(ProbingTargetGroupState_FieldPath))
			case ProbingTargetGroup_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			}
		}
	}
	if wholeStateAccepted == false && len(stateMask.Paths) > 0 {
		result.State = stateMask.Project(source.GetState())
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	return result
}

func (fieldMask *ProbingTargetGroup_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProbingTargetGroup))
}

func (fieldMask *ProbingTargetGroup_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type ProbingTargetGroup_State_FieldMask struct {
	Paths []ProbingTargetGroupState_FieldPath
}

func FullProbingTargetGroup_State_FieldMask() *ProbingTargetGroup_State_FieldMask {
	res := &ProbingTargetGroup_State_FieldMask{}
	res.Paths = append(res.Paths, &ProbingTargetGroupState_FieldTerminalPath{selector: ProbingTargetGroupState_FieldPathSelectorTargetCount})
	return res
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) String() string {
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
func (fieldMask *ProbingTargetGroup_State_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProbingTargetGroup_State_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProbingTargetGroupState_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 1)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProbingTargetGroupState_FieldTerminalPath); ok {
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

func (fieldMask *ProbingTargetGroup_State_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProbingTargetGroupState_FieldPath(raw)
	})
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) ProtoMessage() {}

func (fieldMask *ProbingTargetGroup_State_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) Subtract(other *ProbingTargetGroup_State_FieldMask) *ProbingTargetGroup_State_FieldMask {
	result := &ProbingTargetGroup_State_FieldMask{}
	removedSelectors := make([]bool, 1)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProbingTargetGroupState_FieldTerminalPath:
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

func (fieldMask *ProbingTargetGroup_State_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProbingTargetGroup_State_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProbingTargetGroup_State_FieldMask) FilterInputFields() *ProbingTargetGroup_State_FieldMask {
	result := &ProbingTargetGroup_State_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProbingTargetGroup_State_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProbingTargetGroupState_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProbingTargetGroupState_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProbingTargetGroup_State_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProbingTargetGroup_State_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) AppendPath(path ProbingTargetGroupState_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProbingTargetGroupState_FieldPath))
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) GetPaths() []ProbingTargetGroupState_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProbingTargetGroupState_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) Set(target, source *ProbingTargetGroup_State) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProbingTargetGroup_State), source.(*ProbingTargetGroup_State))
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) Project(source *ProbingTargetGroup_State) *ProbingTargetGroup_State {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProbingTargetGroup_State{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProbingTargetGroupState_FieldTerminalPath:
			switch tp.selector {
			case ProbingTargetGroupState_FieldPathSelectorTargetCount:
				result.TargetCount = source.TargetCount
			}
		}
	}
	return result
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProbingTargetGroup_State))
}

func (fieldMask *ProbingTargetGroup_State_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
