package cm16

import (
	"net/http"
	"path"
	"strings"

	"github.com/rightscale/rsc/rsapi"
)

const (
	// Used by rsc to display command line help
	ApiName = "RightScale CM API 1.6"
)

// Data structure that holds parsed command line values
var commandValues rsapi.ActionCommands

// Register all commands with kinpin application
func RegisterCommands(registrar *rsapi.Registrar) {
	commandValues = rsapi.ActionCommands{}
	registrar.RegisterActionCommands(ApiName, GenMetadata, commandValues)
}

// Parse and run command
func (a *Api) RunCommand(cmd string) (*http.Response, error) {
	parsed, err := a.ParseCommand(cmd, "/api", commandValues)
	if err != nil {
		return nil, err
	}
	href := parsed.Uri
	if !strings.HasPrefix(href, "/api") {
		href = path.Join("/api", href)
	}
	req, err := a.BuildHTTPRequest("GET", href, "1.6", parsed.QueryParams, nil)
	if err != nil {
		return nil, err
	}
	return a.PerformRequest(req)
}

// Show command help
func (a *Api) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, "/api", commandValues)
}

// Show command hrefs
func (a *Api) ShowApiActions(cmd string) error {
	return a.ShowActions(cmd, "/api", commandValues)
}
