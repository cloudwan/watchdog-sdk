// Code generated by protoc-gen-goten-validate
// File: watchdog/proto/v1alpha2/shared_token.proto
// DO NOT EDIT!!!

package shared_token

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
	ntt_meta "github.com/cloudwan/edgelq-sdk/common/types/meta"
	common "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/common"
	probe_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
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
	_ = &ntt_meta.Meta{}
	_ = &common.SoftwareVersion{}
	_ = &probe_group.ProbeGroup{}
	_ = &project.Project{}
)

func (obj *SharedToken) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.ProbeTemplate).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("SharedToken", "probeTemplate", obj.ProbeTemplate, "nested object validation failed", err)
		}
	}
	if len(obj.DeduplicationFieldPaths) > 1 {
		values := make(map[string]struct{})
		for _, v := range obj.DeduplicationFieldPaths {
			if _, ok := values[v]; ok {
				return gotenvalidate.NewValidationError("SharedToken", "deduplicationFieldPaths", obj.DeduplicationFieldPaths, "field must contain unique items", nil)
			}
			values[v] = struct{}{}
		}
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("SharedToken", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SharedToken_ProbeTemplate) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbeTemplate", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Spec).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbeTemplate", "spec", obj.Spec, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SharedToken_ProbeTemplate_Meta) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if len(obj.Tags) > 32 {
		return gotenvalidate.NewValidationError("Meta", "tags", obj.Tags, "field must have at most 32 items", nil)
	}
	if len(obj.Tags) > 1 {
		values := make(map[string]struct{})
		for _, v := range obj.Tags {
			if _, ok := values[v]; ok {
				return gotenvalidate.NewValidationError("Meta", "tags", obj.Tags, "field must contain unique items", nil)
			}
			values[v] = struct{}{}
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *SharedToken_ProbeTemplate_Spec) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.PrimaryLocation).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "primaryLocation", obj.PrimaryLocation, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.LocationDiscovery).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "locationDiscovery", obj.LocationDiscovery, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ContactInfo).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "contactInfo", obj.ContactInfo, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
