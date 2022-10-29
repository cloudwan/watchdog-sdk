// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/internet_quality_rating.proto
// DO NOT EDIT!!!

package internet_quality_rating

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	_ = &ntt_meta.Meta{}
	_ = &timestamp.Timestamp{}
	_ = &probe.Probe{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InternetQualityRating_Rating int32

const (
	InternetQualityRating_RATING_UNSPECIFIED InternetQualityRating_Rating = 0
	InternetQualityRating_BAD                InternetQualityRating_Rating = 1
	InternetQualityRating_POOR               InternetQualityRating_Rating = 2
	InternetQualityRating_FAIR               InternetQualityRating_Rating = 3
	InternetQualityRating_GOOD               InternetQualityRating_Rating = 4
	InternetQualityRating_EXCELLENT          InternetQualityRating_Rating = 5
)

var (
	InternetQualityRating_Rating_name = map[int32]string{
		0: "RATING_UNSPECIFIED",
		1: "BAD",
		2: "POOR",
		3: "FAIR",
		4: "GOOD",
		5: "EXCELLENT",
	}

	InternetQualityRating_Rating_value = map[string]int32{
		"RATING_UNSPECIFIED": 0,
		"BAD":                1,
		"POOR":               2,
		"FAIR":               3,
		"GOOD":               4,
		"EXCELLENT":          5,
	}
)

func (x InternetQualityRating_Rating) Enum() *InternetQualityRating_Rating {
	p := new(InternetQualityRating_Rating)
	*p = x
	return p
}

func (x InternetQualityRating_Rating) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), preflect.EnumNumber(x))
}

func (InternetQualityRating_Rating) Descriptor() preflect.EnumDescriptor {
	return watchdog_proto_v1alpha2_internet_quality_rating_proto_enumTypes[0].Descriptor()
}

func (InternetQualityRating_Rating) Type() preflect.EnumType {
	return &watchdog_proto_v1alpha2_internet_quality_rating_proto_enumTypes[0]
}

func (x InternetQualityRating_Rating) Number() preflect.EnumNumber {
	return preflect.EnumNumber(x)
}

// Deprecated, Use InternetQualityRating_Rating.ProtoReflect.Descriptor instead.
func (InternetQualityRating_Rating) EnumDescriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDescGZIP(), []int{0, 0}
}

// InternetQualityRating Resource
type InternetQualityRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	// Name of InternetQualityRating
	Name      *Name                        `protobuf:"bytes,1,opt,customtype=Name,name=name,proto3" json:"name,omitempty" firestore:"name"`
	Rating    InternetQualityRating_Rating `protobuf:"varint,2,opt,name=rating,proto3,enum=ntt.watchdog.v1alpha2.InternetQualityRating_Rating" json:"rating,omitempty" firestore:"rating"`
	Timestamp *timestamp.Timestamp         `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty" firestore:"timestamp"`
	// Metadata
	Metadata *ntt_meta.Meta `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty" firestore:"metadata"`
}

func (m *InternetQualityRating) Reset() {
	*m = InternetQualityRating{}
	if protoimpl.UnsafeEnabled {
		mi := &watchdog_proto_v1alpha2_internet_quality_rating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		ms.StoreMessageInfo(mi)
	}
}

func (m *InternetQualityRating) String() string {
	return protoimpl.X.MessageStringOf(m)
}

func (*InternetQualityRating) ProtoMessage() {}

func (m *InternetQualityRating) ProtoReflect() preflect.Message {
	mi := &watchdog_proto_v1alpha2_internet_quality_rating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && m != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(m))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(m)
}

func (*InternetQualityRating) GotenMessage() {}

// Deprecated, Use InternetQualityRating.ProtoReflect.Descriptor instead.
func (*InternetQualityRating) Descriptor() ([]byte, []int) {
	return watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDescGZIP(), []int{0}
}

func (m *InternetQualityRating) Unmarshal(b []byte) error {
	return proto.Unmarshal(b, m)
}

func (m *InternetQualityRating) Marshal() ([]byte, error) {
	return proto.Marshal(m)
}

func (m *InternetQualityRating) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{}.Marshal(m)
}

func (m *InternetQualityRating) UnmarshalJSON(data []byte) error {
	return protojson.Unmarshal(data, m)
}

func (m *InternetQualityRating) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *InternetQualityRating) GetRating() InternetQualityRating_Rating {
	if m != nil {
		return m.Rating
	}
	return InternetQualityRating_RATING_UNSPECIFIED
}

func (m *InternetQualityRating) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *InternetQualityRating) GetMetadata() *ntt_meta.Meta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *InternetQualityRating) SetName(fv *Name) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Name", "InternetQualityRating"))
	}
	m.Name = fv
}

func (m *InternetQualityRating) SetRating(fv InternetQualityRating_Rating) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Rating", "InternetQualityRating"))
	}
	m.Rating = fv
}

func (m *InternetQualityRating) SetTimestamp(fv *timestamp.Timestamp) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Timestamp", "InternetQualityRating"))
	}
	m.Timestamp = fv
}

func (m *InternetQualityRating) SetMetadata(fv *ntt_meta.Meta) {
	if m == nil {
		panic(fmt.Errorf("can't set %s on nil %s", "Metadata", "InternetQualityRating"))
	}
	m.Metadata = fv
}

var watchdog_proto_v1alpha2_internet_quality_rating_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDesc = []byte{
	0x0a, 0x35, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x05, 0x0a, 0x15, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x31, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x1d, 0xb2, 0xda, 0x21, 0x19, 0x0a, 0x17, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x06, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x42,
	0x41, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x08,
	0x0a, 0x04, 0x46, 0x41, 0x49, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x47, 0x4f, 0x4f, 0x44,
	0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x4e, 0x54, 0x10,
	0x05, 0x3a, 0xd2, 0x02, 0xea, 0x41, 0x90, 0x01, 0x0a, 0x29, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64,
	0x6f, 0x67, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x63, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x7d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x7d, 0x92, 0xd9, 0x21, 0x3f, 0x0a, 0x16, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x51, 0x75,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x05, 0x50, 0x72,
	0x6f, 0x62, 0x65, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x05, 0xda, 0x94, 0x23, 0x50, 0x12,
	0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x46, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x3e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x7d, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2d, 0x2f, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x73, 0x2f, 0x2d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x51,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x2d, 0xc2,
	0x85, 0x2c, 0x23, 0x22, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x2a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0xbe, 0x03, 0xe8, 0xde, 0x21, 0x01, 0xd2, 0xff, 0xd0,
	0x02, 0x64, 0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x42, 0x1a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x51, 0x75,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x5f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x3b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0xd2, 0x84, 0xd1, 0x02, 0x44, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x80, 0xd1, 0x02,
	0x66, 0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDescData = watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDesc
)

func watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDescData
}

var watchdog_proto_v1alpha2_internet_quality_rating_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var watchdog_proto_v1alpha2_internet_quality_rating_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var watchdog_proto_v1alpha2_internet_quality_rating_proto_goTypes = []interface{}{
	(InternetQualityRating_Rating)(0), // 0: ntt.watchdog.v1alpha2.InternetQualityRating_Rating
	(*InternetQualityRating)(nil),     // 1: ntt.watchdog.v1alpha2.InternetQualityRating
	(*timestamp.Timestamp)(nil),       // 2: google.protobuf.Timestamp
	(*ntt_meta.Meta)(nil),             // 3: ntt.types.Meta
}
var watchdog_proto_v1alpha2_internet_quality_rating_proto_depIdxs = []int32{
	0, // 0: ntt.watchdog.v1alpha2.InternetQualityRating.rating:type_name -> ntt.watchdog.v1alpha2.InternetQualityRating_Rating
	2, // 1: ntt.watchdog.v1alpha2.InternetQualityRating.timestamp:type_name -> google.protobuf.Timestamp
	3, // 2: ntt.watchdog.v1alpha2.InternetQualityRating.metadata:type_name -> ntt.types.Meta
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_internet_quality_rating_proto_init() }
func watchdog_proto_v1alpha2_internet_quality_rating_proto_init() {
	if watchdog_proto_v1alpha2_internet_quality_rating_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {

		watchdog_proto_v1alpha2_internet_quality_rating_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternetQualityRating); i {
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

	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           watchdog_proto_v1alpha2_internet_quality_rating_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_internet_quality_rating_proto_depIdxs,
		EnumInfos:         watchdog_proto_v1alpha2_internet_quality_rating_proto_enumTypes,
		MessageInfos:      watchdog_proto_v1alpha2_internet_quality_rating_proto_msgTypes,
	}.Build()
	watchdog_proto_v1alpha2_internet_quality_rating_proto = out.File
	watchdog_proto_v1alpha2_internet_quality_rating_proto_rawDesc = nil
	watchdog_proto_v1alpha2_internet_quality_rating_proto_goTypes = nil
	watchdog_proto_v1alpha2_internet_quality_rating_proto_depIdxs = nil
}
