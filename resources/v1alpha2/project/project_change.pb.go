// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/project_change.proto
// DO NOT EDIT!!!

package project

import (
	"fmt"
	"reflect"
	"sync"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	preflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// proto imports
import (
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = reflect.Method{}
	_ = sync.Once{}

	_ = protojson.MarshalOptions{}
	_ = proto.MarshalOptions{}
	_ = preflect.Value{}
	_ = protoimpl.DescBuilder{}
)

// make sure we're using proto imports
var (
	_ = &field_mask.FieldMask{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ProjectChange is used by Watch notifications Responses to describe change of
// single Project One of Added, Modified, Removed
type ProjectChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Project change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*ProjectChange_Added_
	//	*ProjectChange_Modified_
	//	*ProjectChange_Current_
	//	*ProjectChange_Removed_
	ChangeType isProjectChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *ProjectChange) Reset() {
	*m = ProjectChange{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_project_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProjectChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProjectChange) ProtoMessage() {}

func (m *ProjectChange) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_project_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProjectChange) GotenMessage() {}

// Deprecated, Use ProjectChange.ProtoReflect.Descriptor instead.
func (*ProjectChange) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_project_change_proto_rawDescGZIP(), []int{0}
}

func (m *ProjectChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProjectChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProjectChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProjectChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isProjectChange_ChangeType interface {
	isProjectChange_ChangeType()
}

type ProjectChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *ProjectChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type ProjectChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *ProjectChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type ProjectChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *ProjectChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type ProjectChange_Removed_ struct {
	// Removed is returned when Project is deleted or leaves Query view
	Removed *ProjectChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*ProjectChange_Added_) isProjectChange_ChangeType()    {}
func (*ProjectChange_Modified_) isProjectChange_ChangeType() {}
func (*ProjectChange_Current_) isProjectChange_ChangeType()  {}
func (*ProjectChange_Removed_) isProjectChange_ChangeType()  {}
func (m *ProjectChange) GetChangeType() isProjectChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *ProjectChange) GetAdded() *ProjectChange_Added {
	if x, ok := m.GetChangeType().(*ProjectChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *ProjectChange) GetModified() *ProjectChange_Modified {
	if x, ok := m.GetChangeType().(*ProjectChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *ProjectChange) GetCurrent() *ProjectChange_Current {
	if x, ok := m.GetChangeType().(*ProjectChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *ProjectChange) GetRemoved() *ProjectChange_Removed {
	if x, ok := m.GetChangeType().(*ProjectChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *ProjectChange) SetChangeType(ofv isProjectChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isProjectChange_ChangeType", "ProjectChange"))
	}
	m.ChangeType = ofv
}
func (m *ProjectChange) SetAdded(fv *ProjectChange_Added) {
	m.SetChangeType(&ProjectChange_Added_{Added: fv})
}
func (m *ProjectChange) SetModified(fv *ProjectChange_Modified) {
	m.SetChangeType(&ProjectChange_Modified_{Modified: fv})
}
func (m *ProjectChange) SetCurrent(fv *ProjectChange_Current) {
	m.SetChangeType(&ProjectChange_Current_{Current: fv})
}
func (m *ProjectChange) SetRemoved(fv *ProjectChange_Removed) {
	m.SetChangeType(&ProjectChange_Removed_{Removed: fv})
}

// Project has been added to query view
type ProjectChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Project       *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty" firestore:"project"`
	// Integer describing index of added Project in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ProjectChange_Added) Reset() {
	*m = ProjectChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_project_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProjectChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProjectChange_Added) ProtoMessage() {}

func (m *ProjectChange_Added) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_project_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProjectChange_Added) GotenMessage() {}

// Deprecated, Use ProjectChange_Added.ProtoReflect.Descriptor instead.
func (*ProjectChange_Added) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_project_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *ProjectChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProjectChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProjectChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProjectChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProjectChange_Added) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *ProjectChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ProjectChange_Added) SetProject(fv *Project) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Project", "ProjectChange_Added"))
	}
	m.Project = fv
}

func (m *ProjectChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ProjectChange_Added"))
	}
	m.ViewIndex = fv
}

// Project changed some of it's fields - contains either full document or
// masked change
type ProjectChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified Project
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of Project or masked difference, depending on mask_changes
	// instrumentation of issued [WatchProjectRequest] or [WatchProjectsRequest]
	Project *Project `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty" firestore:"project"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *Project_FieldMask `protobuf:"bytes,3,opt,customtype=Project_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified Project.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying Project new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ProjectChange_Modified) Reset() {
	*m = ProjectChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_project_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProjectChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProjectChange_Modified) ProtoMessage() {}

func (m *ProjectChange_Modified) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_project_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProjectChange_Modified) GotenMessage() {}

// Deprecated, Use ProjectChange_Modified.ProtoReflect.Descriptor instead.
func (*ProjectChange_Modified) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_project_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *ProjectChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProjectChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProjectChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProjectChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProjectChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ProjectChange_Modified) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *ProjectChange_Modified) GetFieldMask() *Project_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *ProjectChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *ProjectChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ProjectChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ProjectChange_Modified"))
	}
	m.Name = fv
}

func (m *ProjectChange_Modified) SetProject(fv *Project) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Project", "ProjectChange_Modified"))
	}
	m.Project = fv
}

func (m *ProjectChange_Modified) SetFieldMask(fv *Project_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "ProjectChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *ProjectChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "ProjectChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *ProjectChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ProjectChange_Modified"))
	}
	m.ViewIndex = fv
}

// Project has been added or modified in a query view. Version used for
// stateless watching
type ProjectChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Project       *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty" firestore:"project"`
}

func (m *ProjectChange_Current) Reset() {
	*m = ProjectChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_project_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProjectChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProjectChange_Current) ProtoMessage() {}

func (m *ProjectChange_Current) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_project_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProjectChange_Current) GotenMessage() {}

// Deprecated, Use ProjectChange_Current.ProtoReflect.Descriptor instead.
func (*ProjectChange_Current) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_project_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *ProjectChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProjectChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProjectChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProjectChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProjectChange_Current) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *ProjectChange_Current) SetProject(fv *Project) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Project", "ProjectChange_Current"))
	}
	m.Project = fv
}

// Removed is returned when Project is deleted or leaves Query view
type ProjectChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed Project index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *ProjectChange_Removed) Reset() {
	*m = ProjectChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_project_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *ProjectChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*ProjectChange_Removed) ProtoMessage() {}

func (m *ProjectChange_Removed) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_project_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*ProjectChange_Removed) GotenMessage() {}

// Deprecated, Use ProjectChange_Removed.ProtoReflect.Descriptor instead.
func (*ProjectChange_Removed) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_project_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *ProjectChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *ProjectChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *ProjectChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *ProjectChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *ProjectChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ProjectChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *ProjectChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "ProjectChange_Removed"))
	}
	m.Name = fv
}

func (m *ProjectChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "ProjectChange_Removed"))
	}
	m.ViewIndex = fv
}

var watchdog_proto_v1alpha2_project_change_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_project_change_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x06, 0x0a, 0x0d, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x42, 0x0a, 0x05,
	0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x2e, 0x41, 0x64, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x4b, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x48, 0x0a,
	0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x48, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x1a, 0x60, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x1a, 0x84, 0x02, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f,
	0xb2, 0xda, 0x21, 0x0b, 0x0a, 0x09, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x4a, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42,
	0x0f, 0xb2, 0xda, 0x21, 0x0b, 0x32, 0x09, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x13, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x43, 0x0a, 0x07, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a,
	0x4d, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xb2, 0xda, 0x21, 0x0b, 0x0a, 0x09,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x0d,
	0x9a, 0xd9, 0x21, 0x09, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x0d, 0x0a,
	0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x79, 0xe8, 0xde,
	0x21, 0x00, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x42, 0x12, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3b,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_project_change_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_project_change_proto_rawDescData = watchdog_proto_v1alpha2_project_change_proto_rawDesc
)

func watchdog_proto_v1alpha2_project_change_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_project_change_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_project_change_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_project_change_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_project_change_proto_rawDescData
}

var watchdog_proto_v1alpha2_project_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var watchdog_proto_v1alpha2_project_change_proto_goTypes = []interface{}{
	(*ProjectChange)(nil),          // 0: ntt.watchdog.v1alpha2.ProjectChange
	(*ProjectChange_Added)(nil),    // 1: ntt.watchdog.v1alpha2.ProjectChange.Added
	(*ProjectChange_Modified)(nil), // 2: ntt.watchdog.v1alpha2.ProjectChange.Modified
	(*ProjectChange_Current)(nil),  // 3: ntt.watchdog.v1alpha2.ProjectChange.Current
	(*ProjectChange_Removed)(nil),  // 4: ntt.watchdog.v1alpha2.ProjectChange.Removed
	(*Project)(nil),                // 5: ntt.watchdog.v1alpha2.Project
	(*Project_FieldMask)(nil),      // 6: ntt.watchdog.v1alpha2.Project_FieldMask
}
var watchdog_proto_v1alpha2_project_change_proto_depIdxs = []int32{
	1, // 0: ntt.watchdog.v1alpha2.ProjectChange.added:type_name -> ntt.watchdog.v1alpha2.ProjectChange.Added
	2, // 1: ntt.watchdog.v1alpha2.ProjectChange.modified:type_name -> ntt.watchdog.v1alpha2.ProjectChange.Modified
	3, // 2: ntt.watchdog.v1alpha2.ProjectChange.current:type_name -> ntt.watchdog.v1alpha2.ProjectChange.Current
	4, // 3: ntt.watchdog.v1alpha2.ProjectChange.removed:type_name -> ntt.watchdog.v1alpha2.ProjectChange.Removed
	5, // 4: ntt.watchdog.v1alpha2.ProjectChange.Added.project:type_name -> ntt.watchdog.v1alpha2.Project
	5, // 5: ntt.watchdog.v1alpha2.ProjectChange.Modified.project:type_name -> ntt.watchdog.v1alpha2.Project
	6, // 6: ntt.watchdog.v1alpha2.ProjectChange.Modified.field_mask:type_name -> ntt.watchdog.v1alpha2.Project_FieldMask
	5, // 7: ntt.watchdog.v1alpha2.ProjectChange.Current.project:type_name -> ntt.watchdog.v1alpha2.Project
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_project_change_proto_init() }
func watchdog_proto_v1alpha2_project_change_proto_init() {
	if watchdog_proto_v1alpha2_project_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_project_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectChange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		watchdog_proto_v1alpha2_project_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectChange_Added); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		watchdog_proto_v1alpha2_project_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectChange_Modified); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		watchdog_proto_v1alpha2_project_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectChange_Current); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		watchdog_proto_v1alpha2_project_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectChange_Removed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}

	watchdog_proto_v1alpha2_project_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ProjectChange_Added_)(nil),
		(*ProjectChange_Modified_)(nil),
		(*ProjectChange_Current_)(nil),
		(*ProjectChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: watchdog_proto_v1alpha2_project_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_project_change_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_project_change_proto_depIdxs,
		MessageInfos:      watchdog_proto_v1alpha2_project_change_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_project_change_proto = out.File
	watchdog_proto_v1alpha2_project_change_proto_rawDesc = nil
	watchdog_proto_v1alpha2_project_change_proto_goTypes = nil
	watchdog_proto_v1alpha2_project_change_proto_depIdxs = nil
}
