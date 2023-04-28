// Code generated by protoc-gen-goten-validate
// File: watchdog/proto/v1alpha2/pcap_custom.proto
// DO NOT EDIT!!!

package pcap_client

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	gotenvalidate "github.com/cloudwan/goten-sdk/runtime/validate"
)

// proto imports
import (
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Errorf
	_ = net.ParseIP
	_ = regexp.Match
	_ = strings.Split
	_ = time.Now
	_ = utf8.RuneCountInString
	_ = url.Parse
	_ = durationpb.Duration{}
	_ = timestamppb.Timestamp{}
	_ = gotenvalidate.NewValidationError
)

// make sure we're using proto imports
var (
	_ = &timestamp.Timestamp{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
)

func (obj *ReportPcapRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *GetPcapRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.Interval == nil {
		return gotenvalidate.NewValidationError("GetPcapRequest", "interval", obj.Interval, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Interval).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("GetPcapRequest", "interval", obj.Interval, "nested object validation failed", err)
		}
	}
	if !(obj.PageSize >= 0 && obj.PageSize <= 100000) {
		return gotenvalidate.NewValidationError("GetPcapRequest", "pageSize", obj.PageSize, "field must be in range [0, 100000]", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *GetPcapResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *GetPcapFileFromAgentRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
