// Code generated by protoc-gen-goten-object
// File: watchdog/proto/v1alpha/admin_area.proto
// DO NOT EDIT!!!

package admin_area

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
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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
	_ = &latlng.LatLng{}
)

type BBox_FieldMask struct {
	Paths []BBox_FieldPath
}

func FullBBox_FieldMask() *BBox_FieldMask {
	res := &BBox_FieldMask{}
	res.Paths = append(res.Paths, &BBox_FieldTerminalPath{selector: BBox_FieldPathSelectorSouthWest})
	res.Paths = append(res.Paths, &BBox_FieldTerminalPath{selector: BBox_FieldPathSelectorNorthEast})
	return res
}

func (fieldMask *BBox_FieldMask) String() string {
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
func (fieldMask *BBox_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *BBox_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseBBox_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *BBox_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*BBox_FieldTerminalPath); ok {
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

func (fieldMask *BBox_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseBBox_FieldPath(raw)
	})
}

func (fieldMask *BBox_FieldMask) ProtoMessage() {}

func (fieldMask *BBox_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *BBox_FieldMask) Subtract(other *BBox_FieldMask) *BBox_FieldMask {
	result := &BBox_FieldMask{}
	removedSelectors := make([]bool, 2)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *BBox_FieldTerminalPath:
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

func (fieldMask *BBox_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*BBox_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *BBox_FieldMask) FilterInputFields() *BBox_FieldMask {
	result := &BBox_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *BBox_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *BBox_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]BBox_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseBBox_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask BBox_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *BBox_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *BBox_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask BBox_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *BBox_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *BBox_FieldMask) AppendPath(path BBox_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *BBox_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(BBox_FieldPath))
}

func (fieldMask *BBox_FieldMask) GetPaths() []BBox_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *BBox_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *BBox_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseBBox_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *BBox_FieldMask) Set(target, source *BBox) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *BBox_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*BBox), source.(*BBox))
}

func (fieldMask *BBox_FieldMask) Project(source *BBox) *BBox {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &BBox{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *BBox_FieldTerminalPath:
			switch tp.selector {
			case BBox_FieldPathSelectorSouthWest:
				result.SouthWest = source.SouthWest
			case BBox_FieldPathSelectorNorthEast:
				result.NorthEast = source.NorthEast
			}
		}
	}
	return result
}

func (fieldMask *BBox_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*BBox))
}

func (fieldMask *BBox_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type Geometry_FieldMask struct {
	Paths []Geometry_FieldPath
}

func FullGeometry_FieldMask() *Geometry_FieldMask {
	res := &Geometry_FieldMask{}
	res.Paths = append(res.Paths, &Geometry_FieldTerminalPath{selector: Geometry_FieldPathSelectorCenter})
	res.Paths = append(res.Paths, &Geometry_FieldTerminalPath{selector: Geometry_FieldPathSelectorBbox})
	return res
}

func (fieldMask *Geometry_FieldMask) String() string {
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
func (fieldMask *Geometry_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *Geometry_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseGeometry_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *Geometry_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 2)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*Geometry_FieldTerminalPath); ok {
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

func (fieldMask *Geometry_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseGeometry_FieldPath(raw)
	})
}

func (fieldMask *Geometry_FieldMask) ProtoMessage() {}

func (fieldMask *Geometry_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *Geometry_FieldMask) Subtract(other *Geometry_FieldMask) *Geometry_FieldMask {
	result := &Geometry_FieldMask{}
	removedSelectors := make([]bool, 2)
	otherSubMasks := map[Geometry_FieldPathSelector]gotenobject.FieldMask{
		Geometry_FieldPathSelectorBbox: &BBox_FieldMask{},
	}
	mySubMasks := map[Geometry_FieldPathSelector]gotenobject.FieldMask{
		Geometry_FieldPathSelectorBbox: &BBox_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *Geometry_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *Geometry_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*Geometry_FieldTerminalPath); ok {
					switch tp.selector {
					case Geometry_FieldPathSelectorBbox:
						mySubMasks[Geometry_FieldPathSelectorBbox] = FullBBox_FieldMask()
					}
				} else if tp, ok := path.(*Geometry_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &Geometry_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *Geometry_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*Geometry_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *Geometry_FieldMask) FilterInputFields() *Geometry_FieldMask {
	result := &Geometry_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *Geometry_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *Geometry_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]Geometry_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseGeometry_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask Geometry_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *Geometry_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Geometry_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask Geometry_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *Geometry_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *Geometry_FieldMask) AppendPath(path Geometry_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *Geometry_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(Geometry_FieldPath))
}

func (fieldMask *Geometry_FieldMask) GetPaths() []Geometry_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *Geometry_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *Geometry_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseGeometry_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *Geometry_FieldMask) Set(target, source *Geometry) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *Geometry_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*Geometry), source.(*Geometry))
}

func (fieldMask *Geometry_FieldMask) Project(source *Geometry) *Geometry {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &Geometry{}
	bboxMask := &BBox_FieldMask{}
	wholeBboxAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *Geometry_FieldTerminalPath:
			switch tp.selector {
			case Geometry_FieldPathSelectorCenter:
				result.Center = source.Center
			case Geometry_FieldPathSelectorBbox:
				result.Bbox = source.Bbox
				wholeBboxAccepted = true
			}
		case *Geometry_FieldSubPath:
			switch tp.selector {
			case Geometry_FieldPathSelectorBbox:
				bboxMask.AppendPath(tp.subPath.(BBox_FieldPath))
			}
		}
	}
	if wholeBboxAccepted == false && len(bboxMask.Paths) > 0 {
		result.Bbox = bboxMask.Project(source.GetBbox())
	}
	return result
}

func (fieldMask *Geometry_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*Geometry))
}

func (fieldMask *Geometry_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type AdminHierarchy_FieldMask struct {
	Paths []AdminHierarchy_FieldPath
}

func FullAdminHierarchy_FieldMask() *AdminHierarchy_FieldMask {
	res := &AdminHierarchy_FieldMask{}
	res.Paths = append(res.Paths, &AdminHierarchy_FieldTerminalPath{selector: AdminHierarchy_FieldPathSelectorContinent})
	res.Paths = append(res.Paths, &AdminHierarchy_FieldTerminalPath{selector: AdminHierarchy_FieldPathSelectorCountry})
	res.Paths = append(res.Paths, &AdminHierarchy_FieldTerminalPath{selector: AdminHierarchy_FieldPathSelectorAdmin1})
	res.Paths = append(res.Paths, &AdminHierarchy_FieldTerminalPath{selector: AdminHierarchy_FieldPathSelectorAdmin2})
	res.Paths = append(res.Paths, &AdminHierarchy_FieldTerminalPath{selector: AdminHierarchy_FieldPathSelectorAdmin3})
	res.Paths = append(res.Paths, &AdminHierarchy_FieldTerminalPath{selector: AdminHierarchy_FieldPathSelectorAdmin4})
	res.Paths = append(res.Paths, &AdminHierarchy_FieldTerminalPath{selector: AdminHierarchy_FieldPathSelectorAdmin5})
	return res
}

func (fieldMask *AdminHierarchy_FieldMask) String() string {
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
func (fieldMask *AdminHierarchy_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *AdminHierarchy_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseAdminHierarchy_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *AdminHierarchy_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 7)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*AdminHierarchy_FieldTerminalPath); ok {
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

func (fieldMask *AdminHierarchy_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseAdminHierarchy_FieldPath(raw)
	})
}

func (fieldMask *AdminHierarchy_FieldMask) ProtoMessage() {}

func (fieldMask *AdminHierarchy_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *AdminHierarchy_FieldMask) Subtract(other *AdminHierarchy_FieldMask) *AdminHierarchy_FieldMask {
	result := &AdminHierarchy_FieldMask{}
	removedSelectors := make([]bool, 7)

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *AdminHierarchy_FieldTerminalPath:
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

func (fieldMask *AdminHierarchy_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*AdminHierarchy_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *AdminHierarchy_FieldMask) FilterInputFields() *AdminHierarchy_FieldMask {
	result := &AdminHierarchy_FieldMask{}
	result.Paths = append(result.Paths, fieldMask.Paths...)
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *AdminHierarchy_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *AdminHierarchy_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]AdminHierarchy_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseAdminHierarchy_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask AdminHierarchy_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *AdminHierarchy_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *AdminHierarchy_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask AdminHierarchy_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *AdminHierarchy_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *AdminHierarchy_FieldMask) AppendPath(path AdminHierarchy_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *AdminHierarchy_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(AdminHierarchy_FieldPath))
}

func (fieldMask *AdminHierarchy_FieldMask) GetPaths() []AdminHierarchy_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *AdminHierarchy_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *AdminHierarchy_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseAdminHierarchy_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *AdminHierarchy_FieldMask) Set(target, source *AdminHierarchy) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *AdminHierarchy_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*AdminHierarchy), source.(*AdminHierarchy))
}

func (fieldMask *AdminHierarchy_FieldMask) Project(source *AdminHierarchy) *AdminHierarchy {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &AdminHierarchy{}

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *AdminHierarchy_FieldTerminalPath:
			switch tp.selector {
			case AdminHierarchy_FieldPathSelectorContinent:
				result.Continent = source.Continent
			case AdminHierarchy_FieldPathSelectorCountry:
				result.Country = source.Country
			case AdminHierarchy_FieldPathSelectorAdmin1:
				result.Admin1 = source.Admin1
			case AdminHierarchy_FieldPathSelectorAdmin2:
				result.Admin2 = source.Admin2
			case AdminHierarchy_FieldPathSelectorAdmin3:
				result.Admin3 = source.Admin3
			case AdminHierarchy_FieldPathSelectorAdmin4:
				result.Admin4 = source.Admin4
			case AdminHierarchy_FieldPathSelectorAdmin5:
				result.Admin5 = source.Admin5
			}
		}
	}
	return result
}

func (fieldMask *AdminHierarchy_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*AdminHierarchy))
}

func (fieldMask *AdminHierarchy_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}

type AdminArea_FieldMask struct {
	Paths []AdminArea_FieldPath
}

func FullAdminArea_FieldMask() *AdminArea_FieldMask {
	res := &AdminArea_FieldMask{}
	res.Paths = append(res.Paths, &AdminArea_FieldTerminalPath{selector: AdminArea_FieldPathSelectorName})
	res.Paths = append(res.Paths, &AdminArea_FieldTerminalPath{selector: AdminArea_FieldPathSelectorDisplayName})
	res.Paths = append(res.Paths, &AdminArea_FieldTerminalPath{selector: AdminArea_FieldPathSelectorHierarchy})
	res.Paths = append(res.Paths, &AdminArea_FieldTerminalPath{selector: AdminArea_FieldPathSelectorType})
	res.Paths = append(res.Paths, &AdminArea_FieldTerminalPath{selector: AdminArea_FieldPathSelectorLabelGeometry})
	res.Paths = append(res.Paths, &AdminArea_FieldTerminalPath{selector: AdminArea_FieldPathSelectorGeometry})
	res.Paths = append(res.Paths, &AdminArea_FieldTerminalPath{selector: AdminArea_FieldPathSelectorMetadata})
	res.Paths = append(res.Paths, &AdminArea_FieldTerminalPath{selector: AdminArea_FieldPathSelectorLocalLanguageNames})
	return res
}

func (fieldMask *AdminArea_FieldMask) String() string {
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
func (fieldMask *AdminArea_FieldMask) EncodeFirestore() (*firestorepb.Value, error) {
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

func (fieldMask *AdminArea_FieldMask) DecodeFirestore(fpbv *firestorepb.Value) error {
	for _, value := range fpbv.GetArrayValue().GetValues() {
		parsedPath, err := ParseAdminArea_FieldPath(value.GetStringValue())
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, parsedPath)
	}
	return nil
}

func (fieldMask *AdminArea_FieldMask) IsFull() bool {
	if fieldMask == nil {
		return false
	}
	presentSelectors := make([]bool, 8)
	for _, path := range fieldMask.Paths {
		if asFinal, ok := path.(*AdminArea_FieldTerminalPath); ok {
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

func (fieldMask *AdminArea_FieldMask) ProtoReflect() preflect.Message {
	return gotenobject.MakeFieldMaskReflection(fieldMask, func(raw string) (gotenobject.FieldPath, error) {
		return ParseAdminArea_FieldPath(raw)
	})
}

func (fieldMask *AdminArea_FieldMask) ProtoMessage() {}

func (fieldMask *AdminArea_FieldMask) Reset() {
	if fieldMask != nil {
		fieldMask.Paths = nil
	}
}

func (fieldMask *AdminArea_FieldMask) Subtract(other *AdminArea_FieldMask) *AdminArea_FieldMask {
	result := &AdminArea_FieldMask{}
	removedSelectors := make([]bool, 8)
	otherSubMasks := map[AdminArea_FieldPathSelector]gotenobject.FieldMask{
		AdminArea_FieldPathSelectorHierarchy:     &AdminHierarchy_FieldMask{},
		AdminArea_FieldPathSelectorLabelGeometry: &Geometry_FieldMask{},
		AdminArea_FieldPathSelectorGeometry:      &Geometry_FieldMask{},
		AdminArea_FieldPathSelectorMetadata:      &ntt_meta.Meta_FieldMask{},
	}
	mySubMasks := map[AdminArea_FieldPathSelector]gotenobject.FieldMask{
		AdminArea_FieldPathSelectorHierarchy:     &AdminHierarchy_FieldMask{},
		AdminArea_FieldPathSelectorLabelGeometry: &Geometry_FieldMask{},
		AdminArea_FieldPathSelectorGeometry:      &Geometry_FieldMask{},
		AdminArea_FieldPathSelectorMetadata:      &ntt_meta.Meta_FieldMask{},
	}

	for _, path := range other.GetPaths() {
		switch tp := path.(type) {
		case *AdminArea_FieldTerminalPath:
			removedSelectors[int(tp.selector)] = true
		case *AdminArea_FieldSubPath:
			otherSubMasks[tp.selector].AppendRawPath(tp.subPath)
		}
	}
	for _, path := range fieldMask.GetPaths() {
		if !removedSelectors[int(path.Selector())] {
			if otherSubMask := otherSubMasks[path.Selector()]; otherSubMask != nil && otherSubMask.PathsCount() > 0 {
				if tp, ok := path.(*AdminArea_FieldTerminalPath); ok {
					switch tp.selector {
					case AdminArea_FieldPathSelectorHierarchy:
						mySubMasks[AdminArea_FieldPathSelectorHierarchy] = FullAdminHierarchy_FieldMask()
					case AdminArea_FieldPathSelectorLabelGeometry:
						mySubMasks[AdminArea_FieldPathSelectorLabelGeometry] = FullGeometry_FieldMask()
					case AdminArea_FieldPathSelectorGeometry:
						mySubMasks[AdminArea_FieldPathSelectorGeometry] = FullGeometry_FieldMask()
					case AdminArea_FieldPathSelectorMetadata:
						mySubMasks[AdminArea_FieldPathSelectorMetadata] = ntt_meta.FullMeta_FieldMask()
					}
				} else if tp, ok := path.(*AdminArea_FieldSubPath); ok {
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
				result.Paths = append(result.Paths, &AdminArea_FieldSubPath{selector: selector, subPath: allowedPath})
			}
		}
	}

	if len(result.Paths) == 0 {
		return nil
	}
	return result
}

func (fieldMask *AdminArea_FieldMask) SubtractRaw(other gotenobject.FieldMask) gotenobject.FieldMask {
	return fieldMask.Subtract(other.(*AdminArea_FieldMask))
}

// FilterInputFields generates copy of field paths with output_only field paths removed
func (fieldMask *AdminArea_FieldMask) FilterInputFields() *AdminArea_FieldMask {
	result := &AdminArea_FieldMask{}
	for _, path := range fieldMask.Paths {
		switch path.Selector() {
		case AdminArea_FieldPathSelectorMetadata:
			if _, ok := path.(*AdminArea_FieldTerminalPath); ok {
				for _, subpath := range ntt_meta.FullMeta_FieldMask().FilterInputFields().Paths {
					result.Paths = append(result.Paths, &AdminArea_FieldSubPath{selector: path.Selector(), subPath: subpath})
				}
			} else if sub, ok := path.(*AdminArea_FieldSubPath); ok {
				selectedMask := &ntt_meta.Meta_FieldMask{
					Paths: []ntt_meta.Meta_FieldPath{sub.subPath.(ntt_meta.Meta_FieldPath)},
				}
				for _, allowedPath := range selectedMask.FilterInputFields().Paths {
					result.Paths = append(result.Paths, &AdminArea_FieldSubPath{selector: AdminArea_FieldPathSelectorMetadata, subPath: allowedPath})
				}
			}
		default:
			result.Paths = append(result.Paths, path)
		}
	}
	return result
}

// ToFieldMask is used for proto conversions
func (fieldMask *AdminArea_FieldMask) ToProtoFieldMask() *fieldmaskpb.FieldMask {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	for _, path := range fieldMask.Paths {
		protoFieldMask.Paths = append(protoFieldMask.Paths, path.String())
	}
	return protoFieldMask
}

func (fieldMask *AdminArea_FieldMask) FromProtoFieldMask(protoFieldMask *fieldmaskpb.FieldMask) error {
	if fieldMask == nil {
		return status.Error(codes.Internal, "target field mask is nil")
	}
	fieldMask.Paths = make([]AdminArea_FieldPath, 0, len(protoFieldMask.Paths))
	for _, strPath := range protoFieldMask.Paths {
		path, err := ParseAdminArea_FieldPath(strPath)
		if err != nil {
			return err
		}
		fieldMask.Paths = append(fieldMask.Paths, path)
	}
	return nil
}

// implement methods required by customType
func (fieldMask AdminArea_FieldMask) Marshal() ([]byte, error) {
	protoFieldMask := fieldMask.ToProtoFieldMask()
	return proto.Marshal(protoFieldMask)
}

func (fieldMask *AdminArea_FieldMask) Unmarshal(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := proto.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *AdminArea_FieldMask) Size() int {
	return proto.Size(fieldMask.ToProtoFieldMask())
}

func (fieldMask AdminArea_FieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(fieldMask.ToProtoFieldMask())
}

func (fieldMask *AdminArea_FieldMask) UnmarshalJSON(data []byte) error {
	protoFieldMask := &fieldmaskpb.FieldMask{}
	if err := json.Unmarshal(data, protoFieldMask); err != nil {
		return err
	}
	if err := fieldMask.FromProtoFieldMask(protoFieldMask); err != nil {
		return err
	}
	return nil
}

func (fieldMask *AdminArea_FieldMask) AppendPath(path AdminArea_FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path)
}

func (fieldMask *AdminArea_FieldMask) AppendRawPath(path gotenobject.FieldPath) {
	fieldMask.Paths = append(fieldMask.Paths, path.(AdminArea_FieldPath))
}

func (fieldMask *AdminArea_FieldMask) GetPaths() []AdminArea_FieldPath {
	if fieldMask == nil {
		return nil
	}
	return fieldMask.Paths
}

func (fieldMask *AdminArea_FieldMask) GetRawPaths() []gotenobject.FieldPath {
	if fieldMask == nil {
		return nil
	}
	rawPaths := make([]gotenobject.FieldPath, 0, len(fieldMask.Paths))
	for _, path := range fieldMask.Paths {
		rawPaths = append(rawPaths, path)
	}
	return rawPaths
}

func (fieldMask *AdminArea_FieldMask) SetFromCliFlag(raw string) error {
	path, err := ParseAdminArea_FieldPath(raw)
	if err != nil {
		return err
	}
	fieldMask.Paths = append(fieldMask.Paths, path)
	return nil
}

func (fieldMask *AdminArea_FieldMask) Set(target, source *AdminArea) {
	for _, path := range fieldMask.Paths {
		val, _ := path.GetSingle(source)
		// if val is nil, then field does not exist in source, skip
		// otherwise, process (can still reflect.ValueOf(val).IsNil!)
		if val != nil {
			path.WithIValue(val).SetTo(&target)
		}
	}
}

func (fieldMask *AdminArea_FieldMask) SetRaw(target, source gotenobject.GotenObjectExt) {
	fieldMask.Set(target.(*AdminArea), source.(*AdminArea))
}

func (fieldMask *AdminArea_FieldMask) Project(source *AdminArea) *AdminArea {
	if source == nil {
		return nil
	}
	if fieldMask == nil {
		return source
	}
	result := &AdminArea{}
	hierarchyMask := &AdminHierarchy_FieldMask{}
	wholeHierarchyAccepted := false
	labelGeometryMask := &Geometry_FieldMask{}
	wholeLabelGeometryAccepted := false
	geometryMask := &Geometry_FieldMask{}
	wholeGeometryAccepted := false
	metadataMask := &ntt_meta.Meta_FieldMask{}
	wholeMetadataAccepted := false
	var localLanguageNamesMapKeys []string
	wholeLocalLanguageNamesAccepted := false

	for _, p := range fieldMask.Paths {
		switch tp := p.(type) {
		case *AdminArea_FieldTerminalPath:
			switch tp.selector {
			case AdminArea_FieldPathSelectorName:
				result.Name = source.Name
			case AdminArea_FieldPathSelectorDisplayName:
				result.DisplayName = source.DisplayName
			case AdminArea_FieldPathSelectorHierarchy:
				result.Hierarchy = source.Hierarchy
				wholeHierarchyAccepted = true
			case AdminArea_FieldPathSelectorType:
				result.Type = source.Type
			case AdminArea_FieldPathSelectorLabelGeometry:
				result.LabelGeometry = source.LabelGeometry
				wholeLabelGeometryAccepted = true
			case AdminArea_FieldPathSelectorGeometry:
				result.Geometry = source.Geometry
				wholeGeometryAccepted = true
			case AdminArea_FieldPathSelectorMetadata:
				result.Metadata = source.Metadata
				wholeMetadataAccepted = true
			case AdminArea_FieldPathSelectorLocalLanguageNames:
				result.LocalLanguageNames = source.LocalLanguageNames
				wholeLocalLanguageNamesAccepted = true
			}
		case *AdminArea_FieldSubPath:
			switch tp.selector {
			case AdminArea_FieldPathSelectorHierarchy:
				hierarchyMask.AppendPath(tp.subPath.(AdminHierarchy_FieldPath))
			case AdminArea_FieldPathSelectorLabelGeometry:
				labelGeometryMask.AppendPath(tp.subPath.(Geometry_FieldPath))
			case AdminArea_FieldPathSelectorGeometry:
				geometryMask.AppendPath(tp.subPath.(Geometry_FieldPath))
			case AdminArea_FieldPathSelectorMetadata:
				metadataMask.AppendPath(tp.subPath.(ntt_meta.Meta_FieldPath))
			}
		case *AdminArea_FieldPathMap:
			switch tp.selector {
			case AdminArea_FieldPathSelectorLocalLanguageNames:
				localLanguageNamesMapKeys = append(localLanguageNamesMapKeys, tp.key)
			}
		}
	}
	if wholeHierarchyAccepted == false && len(hierarchyMask.Paths) > 0 {
		result.Hierarchy = hierarchyMask.Project(source.GetHierarchy())
	}
	if wholeLabelGeometryAccepted == false && len(labelGeometryMask.Paths) > 0 {
		result.LabelGeometry = labelGeometryMask.Project(source.GetLabelGeometry())
	}
	if wholeGeometryAccepted == false && len(geometryMask.Paths) > 0 {
		result.Geometry = geometryMask.Project(source.GetGeometry())
	}
	if wholeMetadataAccepted == false && len(metadataMask.Paths) > 0 {
		result.Metadata = metadataMask.Project(source.GetMetadata())
	}
	if wholeLocalLanguageNamesAccepted == false && len(localLanguageNamesMapKeys) > 0 && source.GetLocalLanguageNames() != nil {
		copiedMap := map[string]string{}
		sourceMap := source.GetLocalLanguageNames()
		for _, key := range localLanguageNamesMapKeys {
			copiedMap[key] = sourceMap[key]
		}
		result.LocalLanguageNames = copiedMap
	}
	return result
}

func (fieldMask *AdminArea_FieldMask) ProjectRaw(source gotenobject.GotenObjectExt) gotenobject.GotenObjectExt {
	return fieldMask.Project(source.(*AdminArea))
}

func (fieldMask *AdminArea_FieldMask) PathsCount() int {
	if fieldMask == nil {
		return 0
	}
	return len(fieldMask.Paths)
}
