// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha/speed_test_service.proto
// DO NOT EDIT!!!

package speed_test_client

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
import ()

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
var ()

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var watchdog_proto_v1alpha_speed_test_service_proto preflect.FileDescriptor

var watchdog_proto_v1alpha_speed_test_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67,
	0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2a, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64,
	0x67, 0x65, 0x6c, 0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xd3, 0x02, 0x0a, 0x10, 0x53, 0x70, 0x65, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8c, 0x02, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52,
	0x75, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x70, 0x65,
	0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa2,
	0x01, 0x82, 0xdb, 0x21, 0x15, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x0c, 0x72, 0x75,
	0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x8a, 0xdb, 0x21, 0x1d, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x52,
	0x75, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x32, 0x22, 0x30, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72, 0x75, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x54,
	0x65, 0x73, 0x74, 0x92, 0x97, 0x22, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x72,
	0x75, 0x6e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0xc8, 0xd5, 0x22, 0x01, 0xd0,
	0xd5, 0x22, 0x02, 0x30, 0x01, 0x1a, 0x30, 0xca, 0x41, 0x13, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64,
	0x6f, 0x67, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x17,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x64, 0x67,
	0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0xc4, 0x03, 0xe8, 0xde, 0x21, 0x01, 0x82, 0xff,
	0xd0, 0x02, 0x4b, 0x0a, 0x11, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x8a, 0xff,
	0xd0, 0x02, 0x4b, 0x0a, 0x11, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x92, 0x8c,
	0xd1, 0x02, 0x53, 0x0a, 0x15, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x70, 0x65, 0x65,
	0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x42, 0x15, 0x53, 0x70, 0x65, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00, 0x5a, 0x48, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61,
	0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x3b, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x43, 0x0a, 0x0d, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha_speed_test_service_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha_speed_test_service_proto_rawDescData = watchdog_proto_v1alpha_speed_test_service_proto_rawDesc
)

func watchdog_proto_v1alpha_speed_test_service_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha_speed_test_service_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha_speed_test_service_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha_speed_test_service_proto_rawDescData)
	})
	return watchdog_proto_v1alpha_speed_test_service_proto_rawDescData
}

var watchdog_proto_v1alpha_speed_test_service_proto_goTypes = []interface{}{
	(*RunSpeedTestRequest)(nil),  // 0: ntt.watchdog.v1alpha.RunSpeedTestRequest
	(*RunSpeedTestResponse)(nil), // 1: ntt.watchdog.v1alpha.RunSpeedTestResponse
}
var watchdog_proto_v1alpha_speed_test_service_proto_depIdxs = []int32{
	0, // 0: ntt.watchdog.v1alpha.SpeedTestService.RunSpeedTest:input_type -> ntt.watchdog.v1alpha.RunSpeedTestRequest
	1, // 1: ntt.watchdog.v1alpha.SpeedTestService.RunSpeedTest:output_type -> ntt.watchdog.v1alpha.RunSpeedTestResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha_speed_test_service_proto_init() }
func watchdog_proto_v1alpha_speed_test_service_proto_init() {
	if watchdog_proto_v1alpha_speed_test_service_proto != nil {
		return
	}

	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: watchdog_proto_v1alpha_speed_test_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           watchdog_proto_v1alpha_speed_test_service_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha_speed_test_service_proto_depIdxs,
	}.Build()
	watchdog_proto_v1alpha_speed_test_service_proto = out.File
	watchdog_proto_v1alpha_speed_test_service_proto_rawDesc = nil
	watchdog_proto_v1alpha_speed_test_service_proto_goTypes = nil
	watchdog_proto_v1alpha_speed_test_service_proto_depIdxs = nil
}
