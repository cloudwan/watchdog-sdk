// Code generated by protoc-gen-goten-validate
// File: watchdog/proto/v1alpha2/dns_query_test_custom.proto
// DO NOT EDIT!!!

package dns_query_test_client

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
	duration "github.com/golang/protobuf/ptypes/duration"
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
	_ = &duration.Duration{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
)

func (obj *RunDNSQueryTestRequest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Query).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("RunDNSQueryTestRequest", "query", obj.Query, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *RunDNSQueryTestResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.JsonResponse).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("RunDNSQueryTestResponse", "jsonResponse", obj.JsonResponse, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *RunDNSQueryTestResponse_JsonResponse) GotenValidate() error {
	if obj == nil {
		return nil
	}
	for idx, elem := range obj.Queries {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("JsonResponse", "queries", obj.Queries[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.Answers {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("JsonResponse", "answers", obj.Answers[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.Ns {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("JsonResponse", "ns", obj.Ns[idx], "nested object validation failed", err)
			}
		}
	}
	for idx, elem := range obj.Extras {
		if subobj, ok := interface{}(elem).(gotenvalidate.Validator); ok {
			if err := subobj.GotenValidate(); err != nil {
				return gotenvalidate.NewValidationError("JsonResponse", "extras", obj.Extras[idx], "nested object validation failed", err)
			}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
