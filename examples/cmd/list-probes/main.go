package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"github.com/cloudwan/goten-sdk/runtime/api/view"
	gotenresource "github.com/cloudwan/goten-sdk/runtime/resource"

	cprobe "github.com/cloudwan/watchdog-sdk/client/v1alpha2/probe"
	rprobe "github.com/cloudwan/watchdog-sdk/resources/v1alpha2/probe"

	"github.com/cloudwan/edgelq-sdk/examples/utils"
)

/* This is simple ilustation of how to construct a client and send a request.
 *
 * Based on provided parameters (endpoint, credentials, projectID) it lists Probe resources
 * from Watchdog service.
 *
 * In order to use execute example, you need either user account (prepare access token, like one from web browser),
 * or credentials for service account (JSON format, as defined as in edgelq-sdk/common/api/credentials.proto -
 * ServiceAccount). Of course, user or service account will need permission to list probes in project you
 * select, so this is pre-requirement.
 */

func main() {
	ctx := context.Background()

	// ------------------------- Get params and just verify they were provided -------------------------
	endpoint := pflag.StringP("endpoint", "e", "", "API endpoint (for example, watchdog.stg01b.edgelq.com:443)")
	projectId := pflag.StringP("project", "p", "", "Name of the project to list probes from")
	accessToken := pflag.StringP("access-token", "a", "", "Active token for your user")
	credsFile := pflag.StringP("credentials", "c", "", "Path to service account credentials file")
	pflag.Parse()
	if *projectId == "" {
		pflag.Usage()
		panic(fmt.Errorf("no project ID was provided"))
	}
	if *endpoint == "" {
		pflag.Usage()
		panic(fmt.Errorf("no endpoint was provided"))
	}
	if *accessToken == "" && *credsFile == "" {
		pflag.Usage()
		panic(fmt.Errorf("access token OR credentials file are necessary in order to execute this example"))
	}

	// ------------------------- Create GRPC connection -------------------------
	grpcConn := utils.Dial(ctx, *endpoint, *accessToken, *credsFile)
	probeClient := cprobe.NewProbeServiceClient(grpcConn)

	// ------------------------- Lets begin! -------------------------
	listProbesResp, err := probeClient.ListProbes(ctx, &cprobe.ListProbesRequest{
		Parent: rprobe.NewNameBuilder().SetProjectId(*projectId).SetRegionId(gotenresource.WildcardId).Parent(),
		View:   view.View_BASIC,
	})
	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("failed to list probes: %s\n", err))
		os.Exit(1)
	}

	for _, probe := range listProbesResp.GetProbes() {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("received probe %s\n\n", probe))
	}
}
