package ca

import (
	"net/http"

	"github.com/rightscale/rsc/rsapi"
)

const (
	// Used by rsc to display command line help
	ApiName = "RightScale Cloud Analytics APIs"
)

// Data structure that holds parsed command line values
var commandValues rsapi.ActionCommands

// Register all commands with kinpin application
func RegisterCommands(registrar rsapi.ApiCommandRegistrar) {
	commandValues = rsapi.ActionCommands{}
	setupMetadata()
	registrar.RegisterActionCommands(ApiName, GenMetadata, commandValues)
}

// Parse and run command
func (a *Api) RunCommand(cmd string) (*http.Response, error) {
	parsed, err := a.ParseCommand(cmd, "", commandValues)
	if err != nil {
		return nil, err
	}
	req, err := a.BuildHTTPRequest(parsed.HttpMethod, parsed.Uri, "1.0", parsed.QueryParams, parsed.PayloadParams)
	if err != nil {
		return nil, err
	}
	return a.PerformRequest(req)
}

// Show command help
func (a *Api) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, "", commandValues)
}

// Show command actions
func (a *Api) ShowApiActions(cmd string) error {
	return a.ShowActions(cmd, "", commandValues)
}
