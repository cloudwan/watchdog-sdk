// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha2/speed_test_custom.proto
// DO NOT EDIT!!!

package speed_test_client

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
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probing_session "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_session"
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
	_ = &probe.Probe{}
	_ = &probing_session.ProbingSession{}
	_ = &probing_target.ProbingTarget{}
)

type RunSpeedTestRequest_FieldMask struct {
	Paths []RunSpeedTestRequest_FieldPath
}

func FullRunSpeedTestRequest_FieldMask() *RunSpeedTestRequest_FieldMask {
	res := &RunSpeedTestRequest_FieldMask{}
	res.Paths = append(res.Paths, &RunSpeedTestRequest_FieldTerminalPath{selector: RunSpeedTestRequest_FieldPathSelectorName})
	res.Paths = append(res.Paths, &RunSpeedTestRequest_FieldTerminalPath{selector: RunSpeedTestRequest_FieldPathSelectorDirection})
	res.Paths = append(res.Paths, &RunSpeedTestRequest_FieldTerminalPath{selector: RunSpeedTestRequest_FieldPathSelectorProbingSession})
	res.Paths = append(res.Paths, &RunSpeedTestRequest_FieldTerminalPath{selector: RunSpeedTestRequest_FieldPathSelectorProbingTarget})
	return res
}

func (fieldMask *RunSpeedTestRequest_FieldMask) String() string {
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
func (fieldMask *RunSpeedTestRequest_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RunSpeedTestRequest_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRunSpeedTestRequest_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RunSpeedTestRequest_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 4)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RunSpeedTestRequest_FieldTerminalPath); ok {
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

func (fieldMask *RunSpeedTestRequest_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRunSpeedTestRequest_FieldPath(raw)
	})
}

func (fieldMask *RunSpeedTestRequest_FieldMask) ProtoMessage() {}

func (fieldMask *RunSpeedTestRequest_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RunSpeedTestRequest_FieldMask) Subtract(other *RunSpeedTestRequest_FieldMask) *RunSpeedTestRequest_FieldMask {
	result := &RunSpeedTestRequest_FieldMask{}
	removedSelectors := make([]bool, 4)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RunSpeedTestRequest_FieldTerminalPath:
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

func (fieldMask *RunSpeedTestRequest_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RunSpeedTestRequest_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RunSpeedTestRequest_FieldMask) FilterInputFields() *RunSpeedTestRequest_FieldMask {
	result := &RunSpeedTestRequest_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RunSpeedTestRequest_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RunSpeedTestRequest_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RunSpeedTestRequest_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRunSpeedTestRequest_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RunSpeedTestRequest_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RunSpeedTestRequest_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RunSpeedTestRequest_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RunSpeedTestRequest_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RunSpeedTestRequest_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RunSpeedTestRequest_FieldMask) AppendPath(path RunSpeedTestRequest_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RunSpeedTestRequest_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RunSpeedTestRequest_FieldPath))
}

func (fieldMask *RunSpeedTestRequest_FieldMask) GetPaths() []RunSpeedTestRequest_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RunSpeedTestRequest_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RunSpeedTestRequest_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRunSpeedTestRequest_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RunSpeedTestRequest_FieldMask) Set(target, source *RunSpeedTestRequest) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RunSpeedTestRequest_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RunSpeedTestRequest), source.(*RunSpeedTestRequest))
}

func (fieldMask *RunSpeedTestRequest_FieldMask) Project(source *RunSpeedTestRequest) *RunSpeedTestRequest {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RunSpeedTestRequest{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RunSpeedTestRequest_FieldTerminalPath:
			switch tp.selector {
			case RunSpeedTestRequest_FieldPathSelectorName:
				result.Name = source.Name
			case RunSpeedTestRequest_FieldPathSelectorDirection:
				result.Direction = source.Direction
			case RunSpeedTestRequest_FieldPathSelectorProbingSession:
				result.ProbingSession = source.ProbingSession
			case RunSpeedTestRequest_FieldPathSelectorProbingTarget:
				result.ProbingTarget = source.ProbingTarget
			}
		}
	}
	return result
}

func (fieldMask *RunSpeedTestRequest_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RunSpeedTestRequest))
}

func (fieldMask *RunSpeedTestRequest_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type RunSpeedTestResponse_FieldMask struct {
	Paths []RunSpeedTestResponse_FieldPath
}

func FullRunSpeedTestResponse_FieldMask() *RunSpeedTestResponse_FieldMask {
	res := &RunSpeedTestResponse_FieldMask{}
	res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorDirection})
	res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorSpeedMbps})
	res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorServerName})
	res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorServerIp})
	res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorServerLatency})
	res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorServerRetransmitPercentage})
	res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorProbingSession})
	res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorTarget})
	res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorLocalIp})
	res.Paths = append(res.Paths, &RunSpeedTestResponse_FieldTerminalPath{selector: RunSpeedTestResponse_FieldPathSelectorLocalInterface})
	return res
}

func (fieldMask *RunSpeedTestResponse_FieldMask) String() string {
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
func (fieldMask *RunSpeedTestResponse_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RunSpeedTestResponse_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRunSpeedTestResponse_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RunSpeedTestResponse_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 10)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RunSpeedTestResponse_FieldTerminalPath); ok {
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

func (fieldMask *RunSpeedTestResponse_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRunSpeedTestResponse_FieldPath(raw)
	})
}

func (fieldMask *RunSpeedTestResponse_FieldMask) ProtoMessage() {}

func (fieldMask *RunSpeedTestResponse_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RunSpeedTestResponse_FieldMask) Subtract(other *RunSpeedTestResponse_FieldMask) *RunSpeedTestResponse_FieldMask {
	result := &RunSpeedTestResponse_FieldMask{}
	removedSelectors := make([]bool, 10)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RunSpeedTestResponse_FieldTerminalPath:
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

func (fieldMask *RunSpeedTestResponse_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RunSpeedTestResponse_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RunSpeedTestResponse_FieldMask) FilterInputFields() *RunSpeedTestResponse_FieldMask {
	result := &RunSpeedTestResponse_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RunSpeedTestResponse_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RunSpeedTestResponse_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RunSpeedTestResponse_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRunSpeedTestResponse_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RunSpeedTestResponse_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RunSpeedTestResponse_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RunSpeedTestResponse_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RunSpeedTestResponse_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RunSpeedTestResponse_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RunSpeedTestResponse_FieldMask) AppendPath(path RunSpeedTestResponse_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RunSpeedTestResponse_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RunSpeedTestResponse_FieldPath))
}

func (fieldMask *RunSpeedTestResponse_FieldMask) GetPaths() []RunSpeedTestResponse_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RunSpeedTestResponse_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RunSpeedTestResponse_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRunSpeedTestResponse_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RunSpeedTestResponse_FieldMask) Set(target, source *RunSpeedTestResponse) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RunSpeedTestResponse_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RunSpeedTestResponse), source.(*RunSpeedTestResponse))
}

func (fieldMask *RunSpeedTestResponse_FieldMask) Project(source *RunSpeedTestResponse) *RunSpeedTestResponse {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RunSpeedTestResponse{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RunSpeedTestResponse_FieldTerminalPath:
			switch tp.selector {
			case RunSpeedTestResponse_FieldPathSelectorDirection:
				result.Direction = source.Direction
			case RunSpeedTestResponse_FieldPathSelectorSpeedMbps:
				result.SpeedMbps = source.SpeedMbps
			case RunSpeedTestResponse_FieldPathSelectorServerName:
				result.ServerName = source.ServerName
			case RunSpeedTestResponse_FieldPathSelectorServerIp:
				result.ServerIp = source.ServerIp
			case RunSpeedTestResponse_FieldPathSelectorServerLatency:
				result.ServerLatency = source.ServerLatency
			case RunSpeedTestResponse_FieldPathSelectorServerRetransmitPercentage:
				result.ServerRetransmitPercentage = source.ServerRetransmitPercentage
			case RunSpeedTestResponse_FieldPathSelectorProbingSession:
				result.ProbingSession = source.ProbingSession
			case RunSpeedTestResponse_FieldPathSelectorTarget:
				result.Target = source.Target
			case RunSpeedTestResponse_FieldPathSelectorLocalIp:
				result.LocalIp = source.LocalIp
			case RunSpeedTestResponse_FieldPathSelectorLocalInterface:
				result.LocalInterface = source.LocalInterface
			}
		}
	}
	return result
}

func (fieldMask *RunSpeedTestResponse_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RunSpeedTestResponse))
}

func (fieldMask *RunSpeedTestResponse_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type RunSpeedTestRequestToProbe_FieldMask struct {
	Paths []RunSpeedTestRequestToProbe_FieldPath
}

func FullRunSpeedTestRequestToProbe_FieldMask() *RunSpeedTestRequestToProbe_FieldMask {
	res := &RunSpeedTestRequestToProbe_FieldMask{}
	res.Paths = append(res.Paths, &RunSpeedTestRequestToProbe_FieldTerminalPath{selector: RunSpeedTestRequestToProbe_FieldPathSelectorDirection})
	res.Paths = append(res.Paths, &RunSpeedTestRequestToProbe_FieldTerminalPath{selector: RunSpeedTestRequestToProbe_FieldPathSelectorTargetName})
	res.Paths = append(res.Paths, &RunSpeedTestRequestToProbe_FieldTerminalPath{selector: RunSpeedTestRequestToProbe_FieldPathSelectorProbingSession})
	return res
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) String() string {
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
func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseRunSpeedTestRequestToProbe_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 3)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*RunSpeedTestRequestToProbe_FieldTerminalPath); ok {
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

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseRunSpeedTestRequestToProbe_FieldPath(raw)
	})
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) ProtoMessage() {}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) Subtract(other *RunSpeedTestRequestToProbe_FieldMask) *RunSpeedTestRequestToProbe_FieldMask {
	result := &RunSpeedTestRequestToProbe_FieldMask{}
	removedSelectors := make([]bool, 3)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *RunSpeedTestRequestToProbe_FieldTerminalPath:
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

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*RunSpeedTestRequestToProbe_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) FilterInputFields() *RunSpeedTestRequestToProbe_FieldMask {
	result := &RunSpeedTestRequestToProbe_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]RunSpeedTestRequestToProbe_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseRunSpeedTestRequestToProbe_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask RunSpeedTestRequestToProbe_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask RunSpeedTestRequestToProbe_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) AppendPath(path RunSpeedTestRequestToProbe_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(RunSpeedTestRequestToProbe_FieldPath))
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) GetPaths() []RunSpeedTestRequestToProbe_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseRunSpeedTestRequestToProbe_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) Set(target, source *RunSpeedTestRequestToProbe) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*RunSpeedTestRequestToProbe), source.(*RunSpeedTestRequestToProbe))
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) Project(source *RunSpeedTestRequestToProbe) *RunSpeedTestRequestToProbe {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &RunSpeedTestRequestToProbe{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *RunSpeedTestRequestToProbe_FieldTerminalPath:
			switch tp.selector {
			case RunSpeedTestRequestToProbe_FieldPathSelectorDirection:
				result.Direction = source.Direction
			case RunSpeedTestRequestToProbe_FieldPathSelectorTargetName:
				result.TargetName = source.TargetName
			case RunSpeedTestRequestToProbe_FieldPathSelectorProbingSession:
				result.ProbingSession = source.ProbingSession
			}
		}
	}
	return result
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*RunSpeedTestRequestToProbe))
}

func (fieldMask *RunSpeedTestRequestToProbe_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
