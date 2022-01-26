// Code generated by protoc-gen-goten-validate
// File: watchdog/proto/v1alpha2/probing_config.proto
// DO NOT EDIT!!!

package probing_config

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
	probing_target_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
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
	_ = &ntt_meta.Meta{}
	_ = &duration.Duration{}
	_ = &common.SoftwareVersion{}
	_ = &probe.Probe{}
	_ = &probing_target.ProbingTarget{}
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &project.Project{}
)

func (obj *ProbingConfig) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingConfig", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Spec).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingConfig", "spec", obj.Spec, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.Status).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingConfig", "status", obj.Status, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ProbingConfig_Spec) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Speedtest).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Spec", "speedtest", obj.Speedtest, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ProbingConfig_Status) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ProbingConfig_Spec_SpeedTest) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ProbingConfig_Spec_SpeedTest_Server) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Location).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("Server", "location", obj.Location, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
func (obj *ProbingSession) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Location).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingSession", "location", obj.Location, "nested object validation failed", err)
		}
	}
	if !(obj.Tos <= 255) {
		return gotenvalidate.NewValidationError("ProbingSession", "tos", obj.Tos, "field must be less or equal to 255", nil)
	}
	if subobj, ok := interface{}(obj.PathProbing).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingSession", "pathProbing", obj.PathProbing, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.SpeeedtestSettings).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingSession", "speeedtestSettings", obj.SpeeedtestSettings, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.HttpProbingConfig).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingSession", "httpProbingConfig", obj.HttpProbingConfig, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ProxyConfiguration).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingSession", "proxyConfiguration", obj.ProxyConfiguration, "nested object validation failed", err)
		}
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
