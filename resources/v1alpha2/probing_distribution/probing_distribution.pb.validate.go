// Code generated by protoc-gen-goten-validate
// File: watchdog/proto/v1alpha2/probing_distribution.proto
// DO NOT EDIT!!!

package probing_distribution

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
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probing_target "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target"
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
	_ = &probe.Probe{}
	_ = &probing_target.ProbingTarget{}
	_ = &project.Project{}
)

func (obj *ProbingDistribution) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingDistribution", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if obj.Spec == nil {
		return gotenvalidate.NewValidationError("ProbingDistribution", "spec", obj.Spec, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Spec).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingDistribution", "spec", obj.Spec, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Status).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingDistribution", "status", obj.Status, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ProbingDistribution_Spec) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if obj.ProbeSelector == nil {
		return gotenvalidate.NewValidationError("Spec", "probeSelector", obj.ProbeSelector, "field is required", nil)
	}
	if obj.TargetSelector == nil {
		return gotenvalidate.NewValidationError("Spec", "targetSelector", obj.TargetSelector, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.Constraint).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "constraint", obj.Constraint, "nested object validation failed", err)
		}
	}
	if obj.ProbingSettings == nil {
		return gotenvalidate.NewValidationError("Spec", "probingSettings", obj.ProbingSettings, "field is required", nil)
	}
	if subobj, ok := interface{}(obj.ProbingSettings).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "probingSettings", obj.ProbingSettings, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ProbingDistribution_Status) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ProbingDistribution_Status_Regional) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
