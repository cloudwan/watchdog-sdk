// Code generated by protoc-gen-goten-resource
// Resource: ProbingSession
// DO NOT EDIT!!!

package probing_session

import (
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/common/types/traits"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwan/goten-sdk/runtime/goten"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
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
	_ = codes.NotFound
	_ = fmt.Stringer(nil)
	_ = proto.Message(nil)
	_ = status.Status{}
	_ = url.URL{}
	_ = strings.Builder{}

	_ = goten.GotenMessage(nil)
	_ = gotenresource.ListQuery(nil)
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

var parentRegexPath_Project_Region = regexp.MustCompile("^projects/(?P<project_id>-|[\\w][\\w.-]{0,127})/regions/(?P<region_id>-|[a-zA-Z0-9-]{1,128})$")

type ParentName struct {
	NamePattern
	ProjectId string `firestore:"projectId"`
	RegionId  string `firestore:"regionId"`
}

func ParseParentName(name string) (*ParentName, error) {
	var matches []string
	if matches = parentRegexPath_Project_Region.FindStringSubmatch(name); matches != nil {
		return NewNameBuilder().
			SetProjectId(matches[1]).
			SetRegionId(matches[2]).
			Parent(), nil
	}

	return nil, status.Errorf(codes.InvalidArgument, "unable to parse '%s' as ProbingSession parent name", name)
}

func MustParseParentName(name string) *ParentName {
	result, err := ParseParentName(name)
	if err != nil {
		panic(err)
	}
	return result
}

func (name *ParentName) SetFromSegments(segments gotenresource.NameSegments) error {
	if len(segments) == 2 && segments[0].CollectionLowerJson == "projects" && segments[1].CollectionLowerJson == "regions" {
		name.Pattern = NamePattern_Project_Region
		name.ProjectId = segments[0].Id
		name.RegionId = segments[1].Id
		return nil
	}
	return status.Errorf(codes.InvalidArgument, "unable to use segments %s to form ProbingSession parent name", segments)
}

func (name *ParentName) GetProjectName() *project.Name {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Project_Region:
		return project.NewNameBuilder().
			SetId(name.ProjectId).
			Name()
	default:
		return nil
	}
}

func (name *ParentName) IsSpecified() bool {
	if name == nil || name.Pattern == "" {
		return false
	}
	switch name.Pattern {
	case NamePattern_Project_Region:
		return name.ProjectId != "" && name.RegionId != ""
	}
	return false
}

func (name *ParentName) IsFullyQualified() bool {
	if name == nil || name.Pattern == "" {
		return false
	}

	switch name.Pattern {
	case NamePattern_Project_Region:
		return name.ProjectId != "" && name.ProjectId != gotenresource.WildcardId && name.RegionId != "" && name.RegionId != gotenresource.WildcardId
	}

	return false
}

func (name *ParentName) FullyQualifiedName() (string, error) {
	if !name.IsFullyQualified() {
		return "", status.Errorf(codes.InvalidArgument, "Parent name for ProbingSession is not fully qualified")
	}
	return fmt.Sprintf("//watchdog.edgelq.com/%s", name.String()), nil
}

func (name *ParentName) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (name *ParentName) GetPattern() gotenresource.NamePattern {
	if name == nil {
		return ""
	}
	return name.Pattern
}

func (name *ParentName) GetIdParts() map[string]string {
	if name != nil {
		return map[string]string{
			"projectId": name.ProjectId,
			"regionId":  name.RegionId,
		}
	}
	return map[string]string{
		"projectId": "",
		"regionId":  "",
	}
}

func (name *ParentName) GetSegments() gotenresource.NameSegments {
	if name == nil {
		return nil
	}

	switch name.Pattern {
	case NamePattern_Project_Region:
		return gotenresource.NameSegments{
			gotenresource.NameSegment{
				CollectionLowerJson: "projects",
				Id:                  name.ProjectId,
			},
			gotenresource.NameSegment{
				CollectionLowerJson: "regions",
				Id:                  name.RegionId,
			},
		}
	}
	return nil
}

func (name *ParentName) GetIParentName() gotenresource.Name {
	return nil
}

func (name *ParentName) GetIUnderlyingParentName() gotenresource.Name {
	return nil
}

func (name *ParentName) String() string {
	if name == nil {
		return "<nil>"
	}

	if valueStr, err := name.ProtoString(); err != nil {
		panic(err)
	} else {
		return valueStr
	}
}

func (name *ParentName) DescendsFrom(ancestor string) bool {
	if name == nil {
		return false
	}

	switch name.Pattern {
	case NamePattern_Project_Region:
		return ancestor == "projects" || ancestor == "regions"
	}

	return false
}

func (name *ParentName) AsReference() *ParentReference {
	return &ParentReference{ParentName: *name}
}

func (name *ParentName) AsRawReference() gotenresource.Reference {
	return name.AsReference()
}

// implement methods required by protobuf-go library for string-struct conversion

func (name *ParentName) ProtoString() (string, error) {
	if name == nil {
		return "", nil
	}
	switch name.Pattern {
	case NamePattern_Project_Region:
		return "projects/" + name.ProjectId + "/regions/" + name.RegionId, nil
	}
	return "", nil
}

func (name *ParentName) ParseProtoString(data string) error {
	parsed, err := ParseParentName(data)
	if err != nil {
		return err
	}
	*name = *parsed
	return nil
}

// GotenEqual returns true if other is of same type and paths are equal (implements goten.Equaler interface)
func (name *ParentName) GotenEqual(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*ParentName)
	if !ok {
		other2, ok := other.(ParentName)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return name == nil
	} else if name == nil {
		return false
	}
	if name.ProjectId != other1.ProjectId {
		return false
	}
	if name.RegionId != other1.RegionId {
		return false
	}
	if name.Pattern != other1.Pattern {
		return false
	}

	return true
}

// Matches is same as GotenEqual, but also will accept "other" if name is wildcard.
func (name *ParentName) Matches(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*ParentName)
	if !ok {
		other2, ok := other.(ParentName)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return name == nil
	} else if name == nil {
		return false
	}

	if name.Pattern != other1.Pattern {
		return false
	}
	switch name.Pattern {
	case NamePattern_Project_Region:
		if name.ProjectId != other1.ProjectId &&
			name.ProjectId != gotenresource.WildcardId {
			return false
		}
		if name.RegionId != other1.RegionId &&
			name.RegionId != gotenresource.WildcardId {
			return false
		}
	}

	return true
}

// implement CustomTypeCliValue method
func (name *ParentName) SetFromCliFlag(raw string) error {
	parsedName, err := ParseParentName(raw)
	if err != nil {
		return err
	}
	*name = *parsedName
	return nil
}

type ParentReference struct {
	ParentName
	project *project.Project
}

func MakeParentReference(name *ParentName) (*ParentReference, error) {
	return &ParentReference{
		ParentName: *name,
	}, nil
}

func ParseParentReference(name string) (*ParentReference, error) {
	parsedName, err := ParseParentName(name)
	if err != nil {
		return nil, err
	}
	return MakeParentReference(parsedName)
}

func MustParseParentReference(name string) *ParentReference {
	result, err := ParseParentReference(name)
	if err != nil {
		panic(err)
	}
	return result
}
func (ref *ParentReference) GetProjectReference() *project.Reference {
	if ref == nil {
		return nil
	}

	switch ref.Pattern {
	case NamePattern_Project_Region:
		return project.NewNameBuilder().
			SetId(ref.ProjectId).
			Reference()
	default:
		return nil
	}
}

func (ref *ParentReference) GetUnderlyingReference() gotenresource.Reference {
	if ref == nil {
		return nil
	}
	projectRef := ref.GetProjectReference()
	if projectRef != nil {
		return projectRef
	}

	return nil
}

func (ref *ParentReference) ResolveRaw(res gotenresource.Resource) error {
	switch typedRes := res.(type) {
	case *project.Project:
		if name := ref.GetProjectName(); name == nil {
			return status.Errorf(codes.InvalidArgument, "cannot set Project as parent of ProbingSession, because pattern does not match")
		}
		ref.project = typedRes
		return nil
	default:
		return status.Errorf(codes.Internal, "Invalid parent type for ProbingSession, got %s", reflect.TypeOf(res).Elem().Name())
	}
}

func (ref *ParentReference) Resolved() bool {
	if name := ref.GetProjectName(); name != nil {
		return ref.project != nil
	}
	return true
}

func (ref *ParentReference) ClearCached() {
	ref.project = nil
}

func (ref *ParentReference) GetProject() *project.Project {
	if ref == nil {
		return nil
	}
	return ref.project
}

func (ref *ParentReference) GetRawResource() gotenresource.Resource {
	if name := ref.ParentName.GetProjectName(); name != nil {
		return ref.project
	}
	return nil
}

func (ref *ParentReference) IsFullyQualified() bool {
	if ref == nil {
		return false
	}
	return ref.ParentName.IsFullyQualified()
}

func (ref *ParentReference) IsSpecified() bool {
	if ref == nil {
		return false
	}
	return ref.ParentName.IsSpecified()
}

func (ref *ParentReference) GetResourceDescriptor() gotenresource.Descriptor {
	return descriptor
}

func (ref *ParentReference) GetPattern() gotenresource.NamePattern {
	if ref == nil {
		return ""
	}
	return ref.Pattern
}

func (ref *ParentReference) GetIdParts() map[string]string {
	if ref != nil {
		return ref.ParentName.GetIdParts()
	}
	return map[string]string{
		"projectId": "",
		"regionId":  "",
	}
}

func (ref *ParentReference) GetSegments() gotenresource.NameSegments {
	if ref != nil {
		return ref.ParentName.GetSegments()
	}
	return nil
}

func (ref *ParentReference) GetIParentName() gotenresource.Name {
	return nil
}

func (ref *ParentReference) GetIUnderlyingParentName() gotenresource.Name {
	return nil
}

func (ref *ParentReference) String() string {
	if ref == nil {
		return "<nil>"
	}
	return ref.ParentName.String()
}

// implement methods required by protobuf-go library for string-struct conversion

func (ref *ParentReference) ProtoString() (string, error) {
	if ref == nil {
		return "", nil
	}
	return ref.ParentName.ProtoString()
}

func (ref *ParentReference) ParseProtoString(data string) error {
	parsed, err := ParseParentReference(data)
	if err != nil {
		return err
	}
	*ref = *parsed
	return nil
}

// GotenEqual returns true if other is of same type and paths are equal (implements goten.Equaler interface)
func (ref *ParentReference) GotenEqual(other interface{}) bool {
	if other == nil {
		return ref == nil
	}
	other1, ok := other.(*ParentReference)
	if !ok {
		other2, ok := other.(ParentReference)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return ref == nil
	} else if ref == nil {
		return false
	}
	if ref.project != other1.project {
		return false
	}

	return ref.ParentName.GotenEqual(other1.ParentName)
}

// Matches is same as GotenEqual, but also will accept "other" if name is wildcard.
func (name *ParentReference) Matches(other interface{}) bool {
	if other == nil {
		return name == nil
	}
	other1, ok := other.(*ParentReference)
	if !ok {
		other2, ok := other.(ParentReference)
		if ok {
			other1 = &other2
		} else {
			return false
		}
	}
	if other1 == nil {
		return name == nil
	} else if name == nil {
		return false
	}
	return name.ParentName.Matches(&other1.ParentName)
}

// Google CEL integration Type
var celParentReference = types.NewTypeValue("ParentReference", traits.ReceiverType)

func (name *ParentReference) TypeName() string {
	return ".ntt.watchdog.v1alpha2.ProbingSession.ParentReference"
}

func (name *ParentReference) HasTrait(trait int) bool {
	return trait == traits.ReceiverType
}

func (name *ParentReference) Equal(other ref.Val) ref.Val {
	return types.Bool(false)
}

func (name *ParentReference) Value() interface{} {
	return name
}

func (name *ParentReference) Match(pattern ref.Val) ref.Val {
	return types.Bool(true)
}

func (name *ParentReference) Receive(function string, overload string, args []ref.Val) ref.Val {
	switch function {
	case "satisfies":
		rhsName, err := ParseParentReference(string(args[0].(types.String)))
		if err != nil {
			return types.ValOrErr(celParentReference, "function %s name parse error: %s", function, err)
		}
		return types.Bool(rhsName.Matches(name))
	default:
		return types.ValOrErr(celParentReference, "no such function - %s", function)
	}
}

func (name *ParentReference) ConvertToNative(typeDesc reflect.Type) (interface{}, error) {
	panic("not required")
}

func (name *ParentReference) ConvertToType(typeVal ref.Type) ref.Val {
	switch typeVal.TypeName() {
	case types.StringType.TypeName():
		return types.String(name.String())
	default:
		panic(fmt.Errorf("unable to convert %s to CEL type %s", "ParentReference", typeVal.TypeName()))
	}
}

func (name *ParentReference) Type() ref.Type {
	return celParentReference
}

// implement CustomTypeCliValue method
func (ref *ParentReference) SetFromCliFlag(raw string) error {
	parsedRef, err := ParseParentReference(raw)
	if err != nil {
		return err
	}
	*ref = *parsedRef
	return nil
}
