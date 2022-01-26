// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/probe_group.proto
// DO NOT EDIT!!!

package probe_group

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
	_ = &project.Project{}
)

type ProbeGroup_FieldMask struct {
	Paths []ProbeGroup_FieldPath
}

func FullProbeGroup_FieldMask() *ProbeGroup_FieldMask {
	res := &ProbeGroup_FieldMask{}
	res.Paths = append(res.Paths, &ProbeGroup_FieldTerminalPath{selector: ProbeGroup_FieldPathSelectorName})
	res.Paths = append(res.Paths, &ProbeGroup_FieldTerminalPath{selector: ProbeGroup_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &ProbeGroup_FieldTerminalPath{selector: ProbeGroup_FieldPathSelectorState})
	res.Paths = append(res.Paths, &ProbeGroup_FieldTerminalPath{selector: ProbeGroup_FieldPathSelectorMetadata})
	return res
}

func (fieldMask *ProbeGroup_FieldMask) String() string {
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
func (fieldMask *ProbeGroup_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProbeGroup_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProbeGroup_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProbeGroup_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 4)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProbeGroup_FieldTerminalPath); ok {
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

func (fieldMask *ProbeGroup_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProbeGroup_FieldPath(raw)
	})
}

func (fieldMask *ProbeGroup_FieldMask) ProtoMessage() {}

func (fieldMask *ProbeGroup_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProbeGroup_FieldMask) Subtract(other *ProbeGroup_FieldMask) *ProbeGroup_FieldMask {
	result := &ProbeGroup_FieldMask{}
	removedSelectors := make([]bool, 4)
	otherSubMasks := map[ProbeGroup_FieldPathSelector]gotenobject.FieldMask{
		ProbeGroup_FieldPathSelectorState:    &ProbeGroup_State_FieldMask{},
		ProbeGroup_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[ProbeGroup_FieldPathSelector]gotenobject.FieldMask{
		ProbeGroup_FieldPathSelectorState:    &ProbeGroup_State_FieldMask{},
		ProbeGroup_FieldPathSelectorMetadata: &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProbeGroup_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *ProbeGroup_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*ProbeGroup_FieldTerminalPath); ok {
					switch tp.selector {
					case ProbeGroup_FieldPathSelectorState:
						mySubMasks[ProbeGroup_FieldPathSelectorState] = FullProbeGroup_State_FieldMask()
					case ProbeGroup_FieldPathSelectorMetadata:
						mySubMasks[ProbeGroup_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*ProbeGroup_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &ProbeGroup_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *ProbeGroup_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProbeGroup_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProbeGroup_FieldMask) FilterInputFields() *ProbeGroup_FieldMask {
	result := &ProbeGroup_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case ProbeGroup_FieldPathSelectorMetadata:
			if _, ok := path.(*ProbeGroup_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProbeGroup_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*ProbeGroup_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &ProbeGroup_FieldSubPath{selector: ProbeGroup_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProbeGroup_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProbeGroup_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProbeGroup_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProbeGroup_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProbeGroup_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProbeGroup_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbeGroup_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProbeGroup_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProbeGroup_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbeGroup_FieldMask) AppendPath(path ProbeGroup_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProbeGroup_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProbeGroup_FieldPath))
}

func (fieldMask *ProbeGroup_FieldMask) GetPaths() []ProbeGroup_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProbeGroup_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProbeGroup_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProbeGroup_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProbeGroup_FieldMask) Set(target, source *ProbeGroup) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProbeGroup_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProbeGroup), source.(*ProbeGroup))
}

func (fieldMask *ProbeGroup_FieldMask) Project(source *ProbeGroup) *ProbeGroup {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProbeGroup{}
	stateMask := &ProbeGroup_State_FieldMask{}
	wholeStateAccepted := false
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProbeGroup_FieldTerminalPath:
			switch tp.selector {
			case ProbeGroup_FieldPathSelectorName:
				result.Name = source.Name
			case ProbeGroup_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case ProbeGroup_FieldPathSelectorState:
				result.State = source.State
				wholeStateAccepted = true
			case ProbeGroup_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			}
		case *ProbeGroup_FieldSubPath:
			switch tp.selector {
			case ProbeGroup_FieldPathSelectorState:
				stateMask.AppendPath(tp.subPath.(ProbeGroupState_FieldPath))
			case ProbeGroup_FieldPathSelectorMetadata:
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

func (fieldMask *ProbeGroup_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProbeGroup))
}

func (fieldMask *ProbeGroup_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type ProbeGroup_State_FieldMask struct {
	Paths []ProbeGroupState_FieldPath
}

func FullProbeGroup_State_FieldMask() *ProbeGroup_State_FieldMask {
	res := &ProbeGroup_State_FieldMask{}
	res.Paths = append(res.Paths, &ProbeGroupState_FieldTerminalPath{selector: ProbeGroupState_FieldPathSelectorTargetCount})
	res.Paths = append(res.Paths, &ProbeGroupState_FieldTerminalPath{selector: ProbeGroupState_FieldPathSelectorRegionalTargetCounts})
	return res
}

func (fieldMask *ProbeGroup_State_FieldMask) String() string {
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
func (fieldMask *ProbeGroup_State_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *ProbeGroup_State_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseProbeGroupState_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *ProbeGroup_State_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*ProbeGroupState_FieldTerminalPath); ok {
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

func (fieldMask *ProbeGroup_State_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseProbeGroupState_FieldPath(raw)
	})
}

func (fieldMask *ProbeGroup_State_FieldMask) ProtoMessage() {}

func (fieldMask *ProbeGroup_State_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *ProbeGroup_State_FieldMask) Subtract(other *ProbeGroup_State_FieldMask) *ProbeGroup_State_FieldMask {
	result := &ProbeGroup_State_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *ProbeGroupState_FieldTerminalPath:
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

func (fieldMask *ProbeGroup_State_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*ProbeGroup_State_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *ProbeGroup_State_FieldMask) FilterInputFields() *ProbeGroup_State_FieldMask {
	result := &ProbeGroup_State_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *ProbeGroup_State_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *ProbeGroup_State_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]ProbeGroupState_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseProbeGroupState_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask ProbeGroup_State_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *ProbeGroup_State_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbeGroup_State_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask ProbeGroup_State_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *ProbeGroup_State_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *ProbeGroup_State_FieldMask) AppendPath(path ProbeGroupState_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *ProbeGroup_State_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(ProbeGroupState_FieldPath))
}

func (fieldMask *ProbeGroup_State_FieldMask) GetPaths() []ProbeGroupState_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *ProbeGroup_State_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *ProbeGroup_State_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseProbeGroupState_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *ProbeGroup_State_FieldMask) Set(target, source *ProbeGroup_State) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *ProbeGroup_State_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*ProbeGroup_State), source.(*ProbeGroup_State))
}

func (fieldMask *ProbeGroup_State_FieldMask) Project(source *ProbeGroup_State) *ProbeGroup_State {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &ProbeGroup_State{}
	var regionalTargetCountsMapKeys []string
	wholeRegionalTargetCountsAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *ProbeGroupState_FieldTerminalPath:
			switch tp.selector {
			case ProbeGroupState_FieldPathSelectorTargetCount:
				result.TargetCount = source.TargetCount
			case ProbeGroupState_FieldPathSelectorRegionalTargetCounts:
				result.RegionalTargetCounts = source.RegionalTargetCounts
				wholeRegionalTargetCountsAccepted = true
			}
		case *ProbeGroupState_FieldPathMap:
			switch tp.selector {
			case ProbeGroupState_FieldPathSelectorRegionalTargetCounts:
				regionalTargetCountsMapKeys = append(regionalTargetCountsMapKeys, tp.key)
			}
		}
	}
	if wholeRegionalTargetCountsAccepted == false && len(regionalTargetCountsMapKeys) > 0 && source.GetRegionalTargetCounts() != nil {
		copiedMap := map[string]int64{}
		sourceMap := source.GetRegionalTargetCounts()
		for _, key := range regionalTargetCountsMapKeys {
			copiedMap[key] = sourceMap[key]
		}
		result.RegionalTargetCounts = copiedMap
	}
	return result
}

func (fieldMask *ProbeGroup_State_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*ProbeGroup_State))
}

func (fieldMask *ProbeGroup_State_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
