// Code generated by protoc-gen-goten-validate
// File: watchdog/proto/v1alpha2/probing_target.proto
// DO NOT EDIT!!!

package probing_target

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
	probing_target_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target_group"
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
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &project.Project{}
)

func (obj *ProbingTarget) GotenValidate() error {
	if obj == nil {
		return nil
	}
	if subobj, ok := interface{}(obj.Metadata).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingTarget", "metadata", obj.Metadata, "nested object validation failed", err)
		}
	}
	if _, ok := common.ProbingMode_name[int32(obj.Mode)]; !ok {
		return gotenvalidate.NewValidationError("ProbingTarget", "mode", obj.Mode, "field must be a defined enum value", nil)
	}
	if obj.Mode == 0 {
		return gotenvalidate.NewValidationError("ProbingTarget", "mode", obj.Mode, "field must not be equal to any of the following values: 0", nil)
	}
	if _, ok := common.IpVersion_name[int32(obj.IpVersion)]; !ok {
		return gotenvalidate.NewValidationError("ProbingTarget", "ipVersion", obj.IpVersion, "field must be a defined enum value", nil)
	}
	if obj.Address == "" {
		return gotenvalidate.NewValidationError("ProbingTarget", "address", obj.Address, "field is required", nil)
	}
	if _, ok := common.LocationType_name[int32(obj.LocationType)]; !ok {
		return gotenvalidate.NewValidationError("ProbingTarget", "locationType", obj.LocationType, "field must be a defined enum value", nil)
	}
	if subobj, ok := interface{}(obj.Location).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingTarget", "location", obj.Location, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.HttpProbingConfig).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingTarget", "httpProbingConfig", obj.HttpProbingConfig, "nested object validation failed", err)
		}
	}
	if subobj, ok := interface{}(obj.ProxyConfiguration).(gotenvalidate.Validator); ok {
		if err := subobj.GotenValidate(); err != nil {
			return gotenvalidate.NewValidationError("ProbingTarget", "proxyConfiguration", obj.ProxyConfiguration, "nested object validation failed", err)
		}
	}
	if !(obj.UdpPort >= 0 && obj.UdpPort <= 65535) {
		return gotenvalidate.NewValidationError("ProbingTarget", "udpPort", obj.UdpPort, "field must be in range [0, 65535]", nil)
	}
	if !(obj.SpeedtestTcpPort >= 0 && obj.SpeedtestTcpPort <= 65535) {
		return gotenvalidate.NewValidationError("ProbingTarget", "speedtestTcpPort", obj.SpeedtestTcpPort, "field must be in range [0, 65535]", nil)
	}
	if cvobj, ok := interface{}(obj).(gotenvalidate.CustomValidator); ok {
		return cvobj.GotenCustomValidate()
	}
	return nil
}
