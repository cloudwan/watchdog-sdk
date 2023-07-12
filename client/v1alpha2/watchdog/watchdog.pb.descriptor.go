// Code generated by protoc-gen-goten-client
// Service: Watchdog
// DO NOT EDIT!!!

package watchdog_client

import (
	gotenclient "github.com/cloudwan/goten-sdk/runtime/client"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"
)

// proto imports
import (
	activation_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/activation"
	admin_area_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/admin_area"
	agent_log_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/agent_log"
	agent_software_version_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/agent_software_version"
	dns_query_test_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/dns_query_test"
	geo_resolver_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/geo_resolver"
	hop_monitor_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/hop_monitor"
	hop_report_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/hop_report"
	http_metrics_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/http_metrics"
	http_test_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/http_test"
	internet_quality_rating_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/internet_quality_rating"
	network_info_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/network_info"
	pcap_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/pcap"
	ping_test_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/ping_test"
	probe_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probe"
	probe_group_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probe_group"
	probe_hardware_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probe_hardware"
	probing_config_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probing_config"
	probing_distribution_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probing_distribution"
	probing_session_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probing_session"
	probing_target_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probing_target"
	probing_target_group_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probing_target_group"
	project_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/project"
	quality_profile_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/quality_profile"
	shared_token_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/shared_token"
	speed_test_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/speed_test"
	tag_client "github.com/cloudwan/watchdog-sdk/client/v1alpha2/tag"
	admin_area "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/admin_area"
	agent_software_version "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/agent_software_version"
	internet_quality_rating "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/internet_quality_rating"
	probe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"
	probe_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe_group"
	probing_config "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_config"
	probing_distribution "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_distribution"
	probing_session "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_session"
	probing_target "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target"
	probing_target_group "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probing_target_group"
	project "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/project"
	quality_profile "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/quality_profile"
	shared_token "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/shared_token"
	tag "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/tag"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = new(gotenclient.MethodDescriptor)
	_ = gotenresource.WildcardId
)

// make sure we're using proto imports
var (
	_ = &admin_area.BBox{}
	_ = &admin_area_client.GetAdminAreaRequest{}
	_ = &agent_software_version.AgentSoftwareVersion{}
	_ = &agent_software_version_client.GetAgentSoftwareVersionRequest{}
	_ = &internet_quality_rating.InternetQualityRating{}
	_ = &internet_quality_rating_client.GetInternetQualityRatingRequest{}
	_ = &probe.Probe{}
	_ = &probe_group.ProbeGroup{}
	_ = &probe_group_client.GetProbeGroupRequest{}
	_ = &probe_client.GetProbeRequest{}
	_ = &probing_config.ProbingConfig{}
	_ = &probing_config_client.GetProbingConfigRequest{}
	_ = &probing_distribution.ProbingDistribution{}
	_ = &probing_distribution_client.GetProbingDistributionRequest{}
	_ = &probing_session.ProbingSession{}
	_ = &probing_session_client.GetProbingSessionRequest{}
	_ = &probing_target.ProbingTarget{}
	_ = &probing_target_group.ProbingTargetGroup{}
	_ = &probing_target_group_client.GetProbingTargetGroupRequest{}
	_ = &probing_target_client.GetProbingTargetRequest{}
	_ = &project.Project{}
	_ = &project_client.GetProjectRequest{}
	_ = &quality_profile.Profile{}
	_ = &quality_profile_client.GetQualityProfileRequest{}
	_ = &shared_token.SharedToken{}
	_ = &shared_token_client.GetSharedTokenRequest{}
	_ = &tag.Tag{}
	_ = &tag_client.GetTagRequest{}
)

var (
	descriptorInitialized bool
	watchdogDescriptor    *WatchdogDescriptor
)

type WatchdogDescriptor struct{}

func (d *WatchdogDescriptor) GetServiceName() string {
	return "watchdog"
}

func (d *WatchdogDescriptor) GetServiceDomain() string {
	return "watchdog.edgelq.com"
}

func (d *WatchdogDescriptor) GetVersion() string {
	return "v1alpha2"
}

func (d *WatchdogDescriptor) GetNextVersion() string {

	return ""
}

func (d *WatchdogDescriptor) AllResourceDescriptors() []gotenresource.Descriptor {
	return []gotenresource.Descriptor{
		admin_area.GetDescriptor(),
		agent_software_version.GetDescriptor(),
		internet_quality_rating.GetDescriptor(),
		probe.GetDescriptor(),
		probe_group.GetDescriptor(),
		probing_config.GetDescriptor(),
		probing_distribution.GetDescriptor(),
		probing_session.GetDescriptor(),
		probing_target.GetDescriptor(),
		probing_target_group.GetDescriptor(),
		project.GetDescriptor(),
		quality_profile.GetDescriptor(),
		shared_token.GetDescriptor(),
		tag.GetDescriptor(),
	}
}

func (d *WatchdogDescriptor) AllApiDescriptors() []gotenclient.ApiDescriptor {
	return []gotenclient.ApiDescriptor{
		activation_client.GetActivationServiceDescriptor(),
		admin_area_client.GetAdminAreaServiceDescriptor(),
		agent_log_client.GetAgentLogServiceDescriptor(),
		agent_software_version_client.GetAgentSoftwareVersionServiceDescriptor(),
		dns_query_test_client.GetDNSQueryTestServiceDescriptor(),
		geo_resolver_client.GetGeoResolverServiceDescriptor(),
		http_metrics_client.GetHTTPMetricsServiceDescriptor(),
		http_test_client.GetHTTPTestServiceDescriptor(),
		hop_monitor_client.GetHopMonitorServiceDescriptor(),
		hop_report_client.GetHopReportServiceDescriptor(),
		internet_quality_rating_client.GetInternetQualityRatingServiceDescriptor(),
		network_info_client.GetNetworkInfoServiceDescriptor(),
		pcap_client.GetPcapServiceDescriptor(),
		ping_test_client.GetPingTestServiceDescriptor(),
		probe_group_client.GetProbeGroupServiceDescriptor(),
		probe_hardware_client.GetProbeHardwareServiceDescriptor(),
		probe_client.GetProbeServiceDescriptor(),
		probing_config_client.GetProbingConfigServiceDescriptor(),
		probing_distribution_client.GetProbingDistributionServiceDescriptor(),
		probing_session_client.GetProbingSessionServiceDescriptor(),
		probing_target_group_client.GetProbingTargetGroupServiceDescriptor(),
		probing_target_client.GetProbingTargetServiceDescriptor(),
		project_client.GetProjectServiceDescriptor(),
		quality_profile_client.GetQualityProfileServiceDescriptor(),
		shared_token_client.GetSharedTokenServiceDescriptor(),
		speed_test_client.GetSpeedTestServiceDescriptor(),
		tag_client.GetTagServiceDescriptor(),
	}
}

func (d *WatchdogDescriptor) AllImportedServiceInfos() []gotenclient.ServiceImportInfo {
	return []gotenclient.ServiceImportInfo{
		{
			Domain:  "devices.edgelq.com",
			Version: "v1alpha2",
		},
		{
			Domain:  "iam.edgelq.com",
			Version: "v1alpha2",
		},
	}
}

func GetWatchdogDescriptor() *WatchdogDescriptor {
	return watchdogDescriptor
}

func initDescriptor() {
	watchdogDescriptor = &WatchdogDescriptor{}
	gotenclient.GetRegistry().RegisterSvcDescriptor(watchdogDescriptor)
}

func init() {
	if !descriptorInitialized {
		initDescriptor()
		descriptorInitialized = true
	}
}
