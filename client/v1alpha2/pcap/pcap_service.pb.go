// Code generated by protoc-gen-goten-go
// File: watchdog/proto/v1alpha2/pcap_service.proto
// DO NOT EDIT!!!

package pcap_client

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
	empty "github.com/golang/protobuf/ptypes/empty"
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
	_ = &empty.Empty{}
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var watchdog_proto_v1alpha2_pcap_service_proto preflect.FileDescriptor

var watchdog_proto_v1alpha2_pcap_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x63, 0x61, 0x70, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6e, 0x74,
	0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f,
	0x74, 0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x74,
	0x65, 0x6e, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63,
	0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x74, 0x65,
	0x6e, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x67, 0x6f, 0x74, 0x65, 0x6e, 0x2f,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x77, 0x61,
	0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x63, 0x61, 0x70, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x69, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x64, 0x67, 0x65, 0x6c,
	0x71, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x34, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70,
	0x72, 0x6f, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe7, 0x0a, 0x0a, 0x0b, 0x50, 0x63, 0x61, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe1, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x50, 0x63, 0x61, 0x70, 0x12, 0x28, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x50, 0x63, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x90, 0x01, 0x82, 0xdb, 0x21, 0x1b, 0x0a, 0x05,
	0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x63, 0x61,
	0x70, 0x2a, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xa2, 0xdc, 0x21, 0x02, 0x10, 0x01, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x3b, 0x22, 0x39, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f, 0x62,
	0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x63, 0x61, 0x70,
	0x92, 0x97, 0x22, 0x22, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x50, 0x63, 0x61, 0x70, 0xc8, 0xd5, 0x22, 0x00, 0x12, 0xa7, 0x01, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x12, 0x25, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4d, 0x82, 0xdb, 0x21, 0x09, 0x12, 0x07, 0x67, 0x65, 0x74,
	0x50, 0x63, 0x61, 0x70, 0xa2, 0xdc, 0x21, 0x02, 0x08, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x3a, 0x67, 0x65, 0x74, 0x50,
	0x63, 0x61, 0x70, 0x92, 0x97, 0x22, 0x19, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x10,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x67, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70,
	0xc8, 0xd5, 0x22, 0x00, 0x12, 0xa5, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70,
	0x46, 0x69, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x2e,
	0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x46, 0x69, 0x6c,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x63, 0x61,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xae, 0x01, 0x82, 0xdb, 0x21, 0x25,
	0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x14, 0x67, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70,
	0x46, 0x69, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x2a, 0x06, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0xa2, 0xdc, 0x21, 0x02, 0x08, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x45, 0x22, 0x43, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x2a,
	0x7d, 0x3a, 0x67, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x92, 0x97, 0x22, 0x2c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x73, 0x2e, 0x67, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0xc8, 0xd5, 0x22, 0x00, 0x30, 0x01, 0x12, 0xc8, 0x02, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x36, 0x2e, 0x6e, 0x74, 0x74, 0x2e,
	0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x37, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x63, 0x61,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xba, 0x01, 0x82, 0xdb, 0x21,
	0x29, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x12, 0x18, 0x67, 0x65, 0x74, 0x50, 0x63, 0x61,
	0x70, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x2a, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xa2, 0xdc, 0x21, 0x02, 0x08, 0x01,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x49, 0x22, 0x47, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x67, 0x65, 0x74, 0x50, 0x63, 0x61, 0x70, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x92,
	0x97, 0x22, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x67, 0x65, 0x74, 0x50, 0x63,
	0x61, 0x70, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0xc8, 0xd5, 0x22, 0x00, 0x12, 0xa4, 0x02, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x63, 0x61, 0x70, 0x49, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x30,
	0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x63, 0x61, 0x70,
	0x49, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x63,
	0x61, 0x70, 0x49, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xa8, 0x01, 0x82, 0xdb, 0x21, 0x23, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x62,
	0x65, 0x12, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x63, 0x61, 0x70, 0x49, 0x73, 0x52, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2a, 0x06, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xa2, 0xdc, 0x21,
	0x02, 0x08, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43, 0x22, 0x41, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x63, 0x61, 0x70, 0x49, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x92, 0x97, 0x22, 0x2a,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x63, 0x61,
	0x70, 0x49, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0xc8, 0xd5, 0x22, 0x00, 0x1a, 0x30,
	0xca, 0x41, 0x13, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2e, 0x65, 0x64, 0x67, 0x65,
	0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x17, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x6c, 0x71, 0x2e, 0x63, 0x6f, 0x6d,
	0x42, 0xc8, 0x02, 0xe8, 0xde, 0x21, 0x01, 0x82, 0xff, 0xd0, 0x02, 0x40, 0x0a, 0x0b, 0x70, 0x63,
	0x61, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x63, 0x61, 0x70, 0x8a, 0xff, 0xd0, 0x02,
	0x40, 0x0a, 0x0b, 0x70, 0x63, 0x61, 0x70, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f, 0x67, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x70, 0x63, 0x61,
	0x70, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x74, 0x74, 0x2e, 0x77, 0x61, 0x74, 0x63, 0x68,
	0x64, 0x6f, 0x67, 0x2e, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42,
	0x10, 0x50, 0x63, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x00, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x64, 0x6f,
	0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2f, 0x70, 0x63, 0x61, 0x70, 0x3b, 0x70, 0x63, 0x61, 0x70, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0xd2, 0x84, 0xd1, 0x02, 0x44, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x77, 0x61, 0x6e, 0x2f, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	watchdog_proto_v1alpha2_pcap_service_proto_rawDescOnce sync.Once
	watchdog_proto_v1alpha2_pcap_service_proto_rawDescData = watchdog_proto_v1alpha2_pcap_service_proto_rawDesc
)

func watchdog_proto_v1alpha2_pcap_service_proto_rawDescGZIP() []byte {
	watchdog_proto_v1alpha2_pcap_service_proto_rawDescOnce.Do(func() {
		watchdog_proto_v1alpha2_pcap_service_proto_rawDescData = protoimpl.X.CompressGZIP(watchdog_proto_v1alpha2_pcap_service_proto_rawDescData)
	})
	return watchdog_proto_v1alpha2_pcap_service_proto_rawDescData
}

var watchdog_proto_v1alpha2_pcap_service_proto_goTypes = []interface{}{
	(*ReportPcapRequest)(nil),                // 0: ntt.watchdog.v1alpha2.ReportPcapRequest
	(*GetPcapRequest)(nil),                   // 1: ntt.watchdog.v1alpha2.GetPcapRequest
	(*GetPcapFileFromAgentRequest)(nil),      // 2: ntt.watchdog.v1alpha2.GetPcapFileFromAgentRequest
	(*GetPcapFileInfoFromAgentRequest)(nil),  // 3: ntt.watchdog.v1alpha2.GetPcapFileInfoFromAgentRequest
	(*CheckPcapIsRunningRequest)(nil),        // 4: ntt.watchdog.v1alpha2.CheckPcapIsRunningRequest
	(*empty.Empty)(nil),                      // 5: google.protobuf.Empty
	(*GetPcapResponse)(nil),                  // 6: ntt.watchdog.v1alpha2.GetPcapResponse
	(*GetPcapFileInfoFromAgentResponse)(nil), // 7: ntt.watchdog.v1alpha2.GetPcapFileInfoFromAgentResponse
	(*CheckPcapIsRunningResponse)(nil),       // 8: ntt.watchdog.v1alpha2.CheckPcapIsRunningResponse
}
var watchdog_proto_v1alpha2_pcap_service_proto_depIdxs = []int32{
	0, // 0: ntt.watchdog.v1alpha2.PcapService.ReportPcap:input_type -> ntt.watchdog.v1alpha2.ReportPcapRequest
	1, // 1: ntt.watchdog.v1alpha2.PcapService.GetPcap:input_type -> ntt.watchdog.v1alpha2.GetPcapRequest
	2, // 2: ntt.watchdog.v1alpha2.PcapService.GetPcapFileFromAgent:input_type -> ntt.watchdog.v1alpha2.GetPcapFileFromAgentRequest
	3, // 3: ntt.watchdog.v1alpha2.PcapService.GetPcapFileInfoFromAgent:input_type -> ntt.watchdog.v1alpha2.GetPcapFileInfoFromAgentRequest
	4, // 4: ntt.watchdog.v1alpha2.PcapService.CheckPcapIsRunning:input_type -> ntt.watchdog.v1alpha2.CheckPcapIsRunningRequest
	5, // 5: ntt.watchdog.v1alpha2.PcapService.ReportPcap:output_type -> google.protobuf.Empty
	6, // 6: ntt.watchdog.v1alpha2.PcapService.GetPcap:output_type -> ntt.watchdog.v1alpha2.GetPcapResponse
	6, // 7: ntt.watchdog.v1alpha2.PcapService.GetPcapFileFromAgent:output_type -> ntt.watchdog.v1alpha2.GetPcapResponse
	7, // 8: ntt.watchdog.v1alpha2.PcapService.GetPcapFileInfoFromAgent:output_type -> ntt.watchdog.v1alpha2.GetPcapFileInfoFromAgentResponse
	8, // 9: ntt.watchdog.v1alpha2.PcapService.CheckPcapIsRunning:output_type -> ntt.watchdog.v1alpha2.CheckPcapIsRunningResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { watchdog_proto_v1alpha2_pcap_service_proto_init() }
func watchdog_proto_v1alpha2_pcap_service_proto_init() {
	if watchdog_proto_v1alpha2_pcap_service_proto != nil {
		return
	}

	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: watchdog_proto_v1alpha2_pcap_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           watchdog_proto_v1alpha2_pcap_service_proto_goTypes,
		DependencyIndexes: watchdog_proto_v1alpha2_pcap_service_proto_depIdxs,
	}.Build()
	watchdog_proto_v1alpha2_pcap_service_proto = out.File
	watchdog_proto_v1alpha2_pcap_service_proto_rawDesc = nil
	watchdog_proto_v1alpha2_pcap_service_proto_goTypes = nil
	watchdog_proto_v1alpha2_pcap_service_proto_depIdxs = nil
}
