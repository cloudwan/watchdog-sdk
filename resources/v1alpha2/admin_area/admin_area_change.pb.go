// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/admin_area_change.proto
// DO NOT EDIT!!!

package admin_area

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

// AdminAreaChange is used by Watch notifications Responses to describe change
// of single AdminArea One of Added, Modified, Removed
type AdminAreaChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// AdminArea change
	//
	// Types that are valid to be assigned to ChangeType:
	//	*AdminAreaChange_Added_
	//	*AdminAreaChange_Modified_
	//	*AdminAreaChange_Current_
	//	*AdminAreaChange_Removed_
	ChangeType isAdminAreaChange_ChangeType `protobuf_oneof:"change_type"`
}

func (m *AdminAreaChange) Reset() {
	*m = AdminAreaChange{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AdminAreaChange) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AdminAreaChange) ProtoMessage() {}

func (m *AdminAreaChange) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AdminAreaChange) GotenMessage() {}

// Deprecated, Use AdminAreaChange.ProtoReflect.Descriptor instead.
func (*AdminAreaChange) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_admin_area_change_proto_rawDescGZIP(), []int{0}
}

func (m *AdminAreaChange) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AdminAreaChange) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AdminAreaChange) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AdminAreaChange) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

type isAdminAreaChange_ChangeType interface {
	isAdminAreaChange_ChangeType()
}

type AdminAreaChange_Added_ struct {
	// Added is returned when watched document is added, either created or
	// enters Query view
	Added *AdminAreaChange_Added `protobuf:"bytes,1,opt,name=added,proto3,oneof" firestore:"added"`
}
type AdminAreaChange_Modified_ struct {
	// Modified is returned when watched document is modified
	Modified *AdminAreaChange_Modified `protobuf:"bytes,2,opt,name=modified,proto3,oneof" firestore:"modified"`
}
type AdminAreaChange_Current_ struct {
	// Current is returned in stateless watch when document enters query view or
	// is modified within.
	Current *AdminAreaChange_Current `protobuf:"bytes,4,opt,name=current,proto3,oneof" firestore:"current"`
}
type AdminAreaChange_Removed_ struct {
	// Removed is returned when AdminArea is deleted or leaves Query view
	Removed *AdminAreaChange_Removed `protobuf:"bytes,3,opt,name=removed,proto3,oneof" firestore:"removed"`
}

func (*AdminAreaChange_Added_) isAdminAreaChange_ChangeType()    {}
func (*AdminAreaChange_Modified_) isAdminAreaChange_ChangeType() {}
func (*AdminAreaChange_Current_) isAdminAreaChange_ChangeType()  {}
func (*AdminAreaChange_Removed_) isAdminAreaChange_ChangeType()  {}
func (m *AdminAreaChange) GetChangeType() isAdminAreaChange_ChangeType {
	if m != nil {
		return m.ChangeType
	}
	return nil
}
func (m *AdminAreaChange) GetAdded() *AdminAreaChange_Added {
	if x, ok := m.GetChangeType().(*AdminAreaChange_Added_); ok {
		return x.Added
	}
	return nil
}
func (m *AdminAreaChange) GetModified() *AdminAreaChange_Modified {
	if x, ok := m.GetChangeType().(*AdminAreaChange_Modified_); ok {
		return x.Modified
	}
	return nil
}
func (m *AdminAreaChange) GetCurrent() *AdminAreaChange_Current {
	if x, ok := m.GetChangeType().(*AdminAreaChange_Current_); ok {
		return x.Current
	}
	return nil
}
func (m *AdminAreaChange) GetRemoved() *AdminAreaChange_Removed {
	if x, ok := m.GetChangeType().(*AdminAreaChange_Removed_); ok {
		return x.Removed
	}
	return nil
}
func (m *AdminAreaChange) SetChangeType(ofv isAdminAreaChange_ChangeType) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "isAdminAreaChange_ChangeType", "AdminAreaChange"))
	}
	m.ChangeType = ofv
}
func (m *AdminAreaChange) SetAdded(fv *AdminAreaChange_Added) {
	m.SetChangeType(&AdminAreaChange_Added_{Added: fv})
}
func (m *AdminAreaChange) SetModified(fv *AdminAreaChange_Modified) {
	m.SetChangeType(&AdminAreaChange_Modified_{Modified: fv})
}
func (m *AdminAreaChange) SetCurrent(fv *AdminAreaChange_Current) {
	m.SetChangeType(&AdminAreaChange_Current_{Current: fv})
}
func (m *AdminAreaChange) SetRemoved(fv *AdminAreaChange_Removed) {
	m.SetChangeType(&AdminAreaChange_Removed_{Removed: fv})
}

// AdminArea has been added to query view
type AdminAreaChange_Added struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	AdminArea     *AdminArea `protobuf:"bytes,1,opt,name=admin_area,json=adminArea,proto3" json:"admin_area,omitempty" firestore:"adminArea"`
	// Integer describing index of added AdminArea in resulting query view.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AdminAreaChange_Added) Reset() {
	*m = AdminAreaChange_Added{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AdminAreaChange_Added) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AdminAreaChange_Added) ProtoMessage() {}

func (m *AdminAreaChange_Added) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AdminAreaChange_Added) GotenMessage() {}

// Deprecated, Use AdminAreaChange_Added.ProtoReflect.Descriptor instead.
func (*AdminAreaChange_Added) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_admin_area_change_proto_rawDescGZIP(), []int{0, 0}
}

func (m *AdminAreaChange_Added) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AdminAreaChange_Added) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AdminAreaChange_Added) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AdminAreaChange_Added) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AdminAreaChange_Added) GetAdminArea() *AdminArea {
	if m != nil {
		return m.AdminArea
	}
	return nil
}

func (m *AdminAreaChange_Added) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AdminAreaChange_Added) SetAdminArea(fv *AdminArea) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AdminArea", "AdminAreaChange_Added"))
	}
	m.AdminArea = fv
}

func (m *AdminAreaChange_Added) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AdminAreaChange_Added"))
	}
	m.ViewIndex = fv
}

// AdminArea changed some of it's fields - contains either full document or
// masked change
type AdminAreaChange_Modified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of modified AdminArea
	Name *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// New version of AdminArea or masked difference, depending on mask_changes
	// instrumentation of issued [WatchAdminAreaRequest] or
	// [WatchAdminAreasRequest]
	AdminArea *AdminArea `protobuf:"bytes,2,opt,name=admin_area,json=adminArea,proto3" json:"admin_area,omitempty" firestore:"adminArea"`
	// Used when mask_changes is set, contains field paths of modified
	// properties.
	FieldMask *AdminArea_FieldMask `protobuf:"bytes,3,opt,customtype=AdminArea_FieldMask,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty" firestore:"fieldMask"`
	// Previous view index specifies previous position of modified AdminArea.
	// When modification doesn't affect sorted order, value will remain
	// identical to [view_index].
	PreviousViewIndex int32 `protobuf:"varint,4,opt,name=previous_view_index,json=previousViewIndex,proto3" json:"previous_view_index,omitempty" firestore:"previousViewIndex"`
	// Integer specifying AdminArea new index in resulting query view.
	ViewIndex int32 `protobuf:"varint,5,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AdminAreaChange_Modified) Reset() {
	*m = AdminAreaChange_Modified{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AdminAreaChange_Modified) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AdminAreaChange_Modified) ProtoMessage() {}

func (m *AdminAreaChange_Modified) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AdminAreaChange_Modified) GotenMessage() {}

// Deprecated, Use AdminAreaChange_Modified.ProtoReflect.Descriptor instead.
func (*AdminAreaChange_Modified) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_admin_area_change_proto_rawDescGZIP(), []int{0, 1}
}

func (m *AdminAreaChange_Modified) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AdminAreaChange_Modified) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AdminAreaChange_Modified) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AdminAreaChange_Modified) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AdminAreaChange_Modified) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AdminAreaChange_Modified) GetAdminArea() *AdminArea {
	if m != nil {
		return m.AdminArea
	}
	return nil
}

func (m *AdminAreaChange_Modified) GetFieldMask() *AdminArea_FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *AdminAreaChange_Modified) GetPreviousViewIndex() int32 {
	if m != nil {
		return m.PreviousViewIndex
	}
	return int32(0)
}

func (m *AdminAreaChange_Modified) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AdminAreaChange_Modified) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AdminAreaChange_Modified"))
	}
	m.Name = fv
}

func (m *AdminAreaChange_Modified) SetAdminArea(fv *AdminArea) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AdminArea", "AdminAreaChange_Modified"))
	}
	m.AdminArea = fv
}

func (m *AdminAreaChange_Modified) SetFieldMask(fv *AdminArea_FieldMask) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "FieldMask", "AdminAreaChange_Modified"))
	}
	m.FieldMask = fv
}

func (m *AdminAreaChange_Modified) SetPreviousViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "PreviousViewIndex", "AdminAreaChange_Modified"))
	}
	m.PreviousViewIndex = fv
}

func (m *AdminAreaChange_Modified) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AdminAreaChange_Modified"))
	}
	m.ViewIndex = fv
}

// AdminArea has been added or modified in a query view. Version used for
// stateless watching
type AdminAreaChange_Current struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	AdminArea     *AdminArea `protobuf:"bytes,1,opt,name=admin_area,json=adminArea,proto3" json:"admin_area,omitempty" firestore:"adminArea"`
}

func (m *AdminAreaChange_Current) Reset() {
	*m = AdminAreaChange_Current{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AdminAreaChange_Current) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AdminAreaChange_Current) ProtoMessage() {}

func (m *AdminAreaChange_Current) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AdminAreaChange_Current) GotenMessage() {}

// Deprecated, Use AdminAreaChange_Current.ProtoReflect.Descriptor instead.
func (*AdminAreaChange_Current) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_admin_area_change_proto_rawDescGZIP(), []int{0, 2}
}

func (m *AdminAreaChange_Current) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AdminAreaChange_Current) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AdminAreaChange_Current) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AdminAreaChange_Current) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AdminAreaChange_Current) GetAdminArea() *AdminArea {
	if m != nil {
		return m.AdminArea
	}
	return nil
}

func (m *AdminAreaChange_Current) SetAdminArea(fv *AdminArea) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "AdminArea", "AdminAreaChange_Current"))
	}
	m.AdminArea = fv
}

// Removed is returned when AdminArea is deleted or leaves Query view
type AdminAreaChange_Removed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Name          *Name `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	// Integer specifying removed AdminArea index. Not populated in stateless
	// watch type.
	ViewIndex int32 `protobuf:"varint,2,opt,name=view_index,json=viewIndex,proto3" json:"view_index,omitempty" firestore:"viewIndex"`
}

func (m *AdminAreaChange_Removed) Reset() {
	*m = AdminAreaChange_Removed{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *AdminAreaChange_Removed) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*AdminAreaChange_Removed) ProtoMessage() {}

func (m *AdminAreaChange_Removed) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*AdminAreaChange_Removed) GotenMessage() {}

// Deprecated, Use AdminAreaChange_Removed.ProtoReflect.Descriptor instead.
func (*AdminAreaChange_Removed) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_admin_area_change_proto_rawDescGZIP(), []int{0, 3}
}

func (m *AdminAreaChange_Removed) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *AdminAreaChange_Removed) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *AdminAreaChange_Removed) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *AdminAreaChange_Removed) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *AdminAreaChange_Removed) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AdminAreaChange_Removed) GetViewIndex() int32 {
	if m != nil {
		return m.ViewIndex
	}
	return int32(0)
}

func (m *AdminAreaChange_Removed) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "AdminAreaChange_Removed"))
	}
	m.Name = fv
}

func (m *AdminAreaChange_Removed) SetViewIndex(fv int32) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "ViewIndex", "AdminAreaChange_Removed"))
	}
	m.ViewIndex = fv
}

var watchdog_proto_v1alpha2_admin_area_change_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_admin_area_change_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x61, 0x72, 0x65, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf6, 0x06, 0x0a, 0x0f, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x41, 0x72, 0x65, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x65,
	0x64, 0x48, 0x00, 0x52, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x12, 0x4d, 0x0a, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x6e,
	0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x00, 0x52,
	0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x07, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6e, 0x74, 0x74,
	0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x4a, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x1a, 0x67, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0a, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61,
	0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x8f, 0x02, 0x0a, 0x08, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f,
	0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x41, 0x72, 0x65, 0x61, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x12,
	0x4c, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42,
	0x11, 0xb2, 0xda, 0x21, 0x0d, 0x32, 0x0b, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72,
	0x65, 0x61, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a,
	0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x56, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a,
	0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x4a, 0x0a, 0x07,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x74,
	0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x09, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x1a, 0x4f, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x11, 0xb2, 0xda, 0x21, 0x0d, 0x0a, 0x0b, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x41, 0x72, 0x65, 0x61, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x76, 0x69, 0x65, 0x77, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x3a, 0x0f, 0x9a, 0xd9, 0x21, 0x0b, 0x0a,
	0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x81, 0x01, 0xe8, 0xde, 0x21, 0x00,
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64,
	0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x14,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x72,
	0x65, 0x61, 0x3b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_admin_area_change_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_admin_area_change_proto_rawDescData = watchdog_proto_v1alpha2_admin_area_change_proto_rawDesc
)

func watchdog_proto_v1alpha2_admin_area_change_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_admin_area_change_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_admin_area_change_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_admin_area_change_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_admin_area_change_proto_rawDescData
}

var watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var watchdog_proto_v1alpha2_admin_area_change_proto_goTypes = []interface{}{
	(*AdminAreaChange)(nil),          // 0: ntt.watchdog.v1alpha2.AdminAreaChange
	(*AdminAreaChange_Added)(nil),    // 1: ntt.watchdog.v1alpha2.AdminAreaChange.Added
	(*AdminAreaChange_Modified)(nil), // 2: ntt.watchdog.v1alpha2.AdminAreaChange.Modified
	(*AdminAreaChange_Current)(nil),  // 3: ntt.watchdog.v1alpha2.AdminAreaChange.Current
	(*AdminAreaChange_Removed)(nil),  // 4: ntt.watchdog.v1alpha2.AdminAreaChange.Removed
	(*AdminArea)(nil),                // 5: ntt.watchdog.v1alpha2.AdminArea
	(*AdminArea_FieldMask)(nil),      // 6: ntt.watchdog.v1alpha2.AdminArea_FieldMask
}
var watchdog_proto_v1alpha2_admin_area_change_proto_depIdxs = []int32{
	1, // 0: ntt.watchdog.v1alpha2.AdminAreaChange.added:type_name -> ntt.watchdog.v1alpha2.AdminAreaChange.Added
	2, // 1: ntt.watchdog.v1alpha2.AdminAreaChange.modified:type_name -> ntt.watchdog.v1alpha2.AdminAreaChange.Modified
	3, // 2: ntt.watchdog.v1alpha2.AdminAreaChange.current:type_name -> ntt.watchdog.v1alpha2.AdminAreaChange.Current
	4, // 3: ntt.watchdog.v1alpha2.AdminAreaChange.removed:type_name -> ntt.watchdog.v1alpha2.AdminAreaChange.Removed
	5, // 4: ntt.watchdog.v1alpha2.AdminAreaChange.Added.admin_area:type_name -> ntt.watchdog.v1alpha2.AdminArea
	5, // 5: ntt.watchdog.v1alpha2.AdminAreaChange.Modified.admin_area:type_name -> ntt.watchdog.v1alpha2.AdminArea
	6, // 6: ntt.watchdog.v1alpha2.AdminAreaChange.Modified.field_mask:type_name -> ntt.watchdog.v1alpha2.AdminArea_FieldMask
	5, // 7: ntt.watchdog.v1alpha2.AdminAreaChange.Current.admin_area:type_name -> ntt.watchdog.v1alpha2.AdminArea
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_admin_area_change_proto_init() }
func watchdog_proto_v1alpha2_admin_area_change_proto_init() {
	if watchdog_proto_v1alpha2_admin_area_change_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminAreaChange); i {
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
		watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminAreaChange_Added); i {
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
		watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminAreaChange_Modified); i {
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
		watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminAreaChange_Current); i {
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
		watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminAreaChange_Removed); i {
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

	watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AdminAreaChange_Added_)(nil),
		(*AdminAreaChange_Modified_)(nil),
		(*AdminAreaChange_Current_)(nil),
		(*AdminAreaChange_Removed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: watchdog_proto_v1alpha2_admin_area_change_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_admin_area_change_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_admin_area_change_proto_depIdxs,
		MessageInfos:      watchdog_proto_v1alpha2_admin_area_change_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_admin_area_change_proto = out.File
	watchdog_proto_v1alpha2_admin_area_change_proto_rawDesc = nil
	watchdog_proto_v1alpha2_admin_area_change_proto_goTypes = nil
	watchdog_proto_v1alpha2_admin_area_change_proto_depIdxs = nil
}
