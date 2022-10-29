// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/hop_monitor_custom.proto
// DO NOT EDIT!!!

package hop_monitor_client

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
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probing_target "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target"
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
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target.ProbingTarget{}
)

type RunHopMonitorRequest_FieldMask struct {
	Paths []RunHopMonitorRequest_FieldPath
}

func FullRunHopMonitorRequest_FieldMask() *RunHopMonitorRequest_FieldMask {
	res := &RunHopMonitorRequest_FieldMask{}
	res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorName})
	res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorSource})
	res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorTarget})
	res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorDestination})
	res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorSizeBytes})
	res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorAttempts})
	res.Paths = append(res.Paths, &RunHopMonitorRequest_FieldTerminalPath{selector: RunHopMonitorRequest_FieldPathSelectorMode})
	return res
}

func (fieldMask *RunHopMonitorRequest_FieldMask) String() string {
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
func (fieldMask *RunHopMonitorRequest_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RunHopMonitorRequest_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRunHopMonitorRequest_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RunHopMonitorRequest_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 7)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RunHopMonitorRequest_FieldTerminalPath); ok {
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

func (fieldMask *RunHopMonitorRequest_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRunHopMonitorRequest_FieldPath(raw)
	})
}

func (fieldMask *RunHopMonitorRequest_FieldMask) ProtoMessage() {}

func (fieldMask *RunHopMonitorRequest_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RunHopMonitorRequest_FieldMask) Subtract(other *RunHopMonitorRequest_FieldMask) *RunHopMonitorRequest_FieldMask {
	result := &RunHopMonitorRequest_FieldMask{}
	removedSelectors := make([]bool, 7)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RunHopMonitorRequest_FieldTerminalPath:
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

func (fieldMask *RunHopMonitorRequest_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RunHopMonitorRequest_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RunHopMonitorRequest_FieldMask) FilterInputFields() *RunHopMonitorRequest_FieldMask {
	result := &RunHopMonitorRequest_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RunHopMonitorRequest_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RunHopMonitorRequest_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RunHopMonitorRequest_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRunHopMonitorRequest_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RunHopMonitorRequest_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RunHopMonitorRequest_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RunHopMonitorRequest_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RunHopMonitorRequest_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RunHopMonitorRequest_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RunHopMonitorRequest_FieldMask) AppendPath(path RunHopMonitorRequest_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RunHopMonitorRequest_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RunHopMonitorRequest_FieldPath))
}

func (fieldMask *RunHopMonitorRequest_FieldMask) GetPaths() []RunHopMonitorRequest_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RunHopMonitorRequest_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RunHopMonitorRequest_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRunHopMonitorRequest_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RunHopMonitorRequest_FieldMask) Set(target, source *RunHopMonitorRequest) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RunHopMonitorRequest_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RunHopMonitorRequest), source.(*RunHopMonitorRequest))
}

func (fieldMask *RunHopMonitorRequest_FieldMask) Project(source *RunHopMonitorRequest) *RunHopMonitorRequest {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RunHopMonitorRequest{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RunHopMonitorRequest_FieldTerminalPath:
			switch tp.selector {
			case RunHopMonitorRequest_FieldPathSelectorName:
				result.Name = source.Name
			case RunHopMonitorRequest_FieldPathSelectorSource:
				result.Source = source.Source
			case RunHopMonitorRequest_FieldPathSelectorTarget:
				result.Target = source.Target
			case RunHopMonitorRequest_FieldPathSelectorDestination:
				result.Destination = source.Destination
			case RunHopMonitorRequest_FieldPathSelectorSizeBytes:
				result.SizeBytes = source.SizeBytes
			case RunHopMonitorRequest_FieldPathSelectorAttempts:
				result.Attempts = source.Attempts
			case RunHopMonitorRequest_FieldPathSelectorMode:
				result.Mode = source.Mode
			}
		}
	}
	return result
}

func (fieldMask *RunHopMonitorRequest_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RunHopMonitorRequest))
}

func (fieldMask *RunHopMonitorRequest_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type RunHopMonitorResponse_FieldMask struct {
	Paths []RunHopMonitorResponse_FieldPath
}

func FullRunHopMonitorResponse_FieldMask() *RunHopMonitorResponse_FieldMask {
	res := &RunHopMonitorResponse_FieldMask{}
	res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorPaths})
	res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorHopStats})
	res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorHopInfo})
	res.Paths = append(res.Paths, &RunHopMonitorResponse_FieldTerminalPath{selector: RunHopMonitorResponse_FieldPathSelectorIpVersion})
	return res
}

func (fieldMask *RunHopMonitorResponse_FieldMask) String() string {
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
func (fieldMask *RunHopMonitorResponse_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RunHopMonitorResponse_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRunHopMonitorResponse_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RunHopMonitorResponse_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 4)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RunHopMonitorResponse_FieldTerminalPath); ok {
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

func (fieldMask *RunHopMonitorResponse_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRunHopMonitorResponse_FieldPath(raw)
	})
}

func (fieldMask *RunHopMonitorResponse_FieldMask) ProtoMessage() {}

func (fieldMask *RunHopMonitorResponse_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RunHopMonitorResponse_FieldMask) Subtract(other *RunHopMonitorResponse_FieldMask) *RunHopMonitorResponse_FieldMask {
	result := &RunHopMonitorResponse_FieldMask{}
	removedSelectors := make([]bool, 4)
	otherSubMasks := map[RunHopMonitorResponse_FieldPathSelector]gotenobject.FieldMask{
		RunHopMonitorResponse_FieldPathSelectorPaths: &common.Path_FieldMask{},
	}
	mySubMasks := map[RunHopMonitorResponse_FieldPathSelector]gotenobject.FieldMask{
		RunHopMonitorResponse_FieldPathSelectorPaths: &common.Path_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RunHopMonitorResponse_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *RunHopMonitorResponse_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*RunHopMonitorResponse_FieldTerminalPath); ok {
					switch tp.selector {
					case RunHopMonitorResponse_FieldPathSelectorPaths:
						mySubMasks[RunHopMonitorResponse_FieldPathSelectorPaths] = common.FullPath_FieldMask()
					}
				} else if tp, ok := path.(*RunHopMonitorResponse_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &RunHopMonitorResponse_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *RunHopMonitorResponse_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RunHopMonitorResponse_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RunHopMonitorResponse_FieldMask) FilterInputFields() *RunHopMonitorResponse_FieldMask {
	result := &RunHopMonitorResponse_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RunHopMonitorResponse_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RunHopMonitorResponse_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RunHopMonitorResponse_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRunHopMonitorResponse_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RunHopMonitorResponse_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RunHopMonitorResponse_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RunHopMonitorResponse_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RunHopMonitorResponse_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RunHopMonitorResponse_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RunHopMonitorResponse_FieldMask) AppendPath(path RunHopMonitorResponse_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RunHopMonitorResponse_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RunHopMonitorResponse_FieldPath))
}

func (fieldMask *RunHopMonitorResponse_FieldMask) GetPaths() []RunHopMonitorResponse_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RunHopMonitorResponse_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RunHopMonitorResponse_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRunHopMonitorResponse_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RunHopMonitorResponse_FieldMask) Set(target, source *RunHopMonitorResponse) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RunHopMonitorResponse_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RunHopMonitorResponse), source.(*RunHopMonitorResponse))
}

func (fieldMask *RunHopMonitorResponse_FieldMask) Project(source *RunHopMonitorResponse) *RunHopMonitorResponse {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RunHopMonitorResponse{}
	pathsMask := &common.Path_FieldMask{}
	wholePathsAccepted := false
	var hopStatsMapKeys []string
	wholeHopStatsAccepted := false
	var hopInfoMapKeys []string
	wholeHopInfoAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RunHopMonitorResponse_FieldTerminalPath:
			switch tp.selector {
			case RunHopMonitorResponse_FieldPathSelectorPaths:
				result.Paths = source.Paths
				wholePathsAccepted = true
			case RunHopMonitorResponse_FieldPathSelectorHopStats:
				result.HopStats = source.HopStats
				wholeHopStatsAccepted = true
			case RunHopMonitorResponse_FieldPathSelectorHopInfo:
				result.HopInfo = source.HopInfo
				wholeHopInfoAccepted = true
			case RunHopMonitorResponse_FieldPathSelectorIpVersion:
				result.IpVersion = source.IpVersion
			}
		case *RunHopMonitorResponse_FieldSubPath:
			switch tp.selector {
			case RunHopMonitorResponse_FieldPathSelectorPaths:
				pathsMask.AppendPath(tp.subPath.(common.Path_FieldPath))
			}
		case *RunHopMonitorResponse_FieldPathMap:
			switch tp.selector {
			case RunHopMonitorResponse_FieldPathSelectorHopStats:
				hopStatsMapKeys = append(hopStatsMapKeys, tp.key)
			case RunHopMonitorResponse_FieldPathSelectorHopInfo:
				hopInfoMapKeys = append(hopInfoMapKeys, tp.key)
			}
		}
	}
	if wholePathsAccepted == false && len(pathsMask.Paths) > 0 {
		for _, sourceItem := range source.GetPaths() {
			result.Paths = append(result.Paths, pathsMask.Project(sourceItem))
		}
	}
	if wholeHopStatsAccepted == false && len(hopStatsMapKeys) > 0 && source.GetHopStats() != nil {
		copiedMap := map[string]*common.HopStat{}
		sourceMap := source.GetHopStats()
		for _, key := range hopStatsMapKeys {
			copiedMap[key] = sourceMap[key]
		}
		result.HopStats = copiedMap
	}
	if wholeHopInfoAccepted == false && len(hopInfoMapKeys) > 0 && source.GetHopInfo() != nil {
		copiedMap := map[string]*common.HopInfo{}
		sourceMap := source.GetHopInfo()
		for _, key := range hopInfoMapKeys {
			copiedMap[key] = sourceMap[key]
		}
		result.HopInfo = copiedMap
	}
	return result
}

func (fieldMask *RunHopMonitorResponse_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RunHopMonitorResponse))
}

func (fieldMask *RunHopMonitorResponse_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
