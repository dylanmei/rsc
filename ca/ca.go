package ca

import (
	"regexp"
	"strings"

	"github.com/rightscale/rsc/ca/cac"
	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// Metadata synthetized from all CA APIs metadata
var GenMetadata map[string]*metadata.Resource

// CA 1.0 common client to all cloud analytics APIs
type Api struct {
	*rsapi.Api
}

// Build client from command line
func FromCommandLine(cmdLine *cmd.CommandLine) (*Api, error) {
	api, err := rsapi.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	setupMetadata()
	api.Host = apiHostFromLogin(cmdLine.Host)
	api.Metadata = GenMetadata
	return &Api{api}, nil
}

// New returns a CA API client.
func New(h string, a rsapi.Authenticator) *Api {
	api := rsapi.New(h, a)
	setupMetadata()
	api.Metadata = GenMetadata
	return &Api{Api: api}
}

func apiHostFromLogin(host string) string {
	integration, _ := regexp.MatchString("^cobalt", host)
	staging, _ := regexp.MatchString("^moo", host)

	prefix := ""

	switch {
	case integration:
		prefix = "ca1-analytics-499"
	case staging:
		prefix = "moo-analytics"
	default:
		prefix = "analytics"
	}

	urlElems := strings.Split(host, ".")
	urlElems[0] = prefix
	return strings.Join(urlElems, ".")
}

// Initialize GenMetadata from each CA API generated metadata
func setupMetadata() {
	GenMetadata = map[string]*metadata.Resource{}
	for n, r := range cac.GenMetadata {
		GenMetadata[n] = r
	}
}
