//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=rl10/docs/api -output=rl10 -pkg=rl10 -target=unversioned -client=Api
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rl10 // import "gopkg.in/rightscale/rsc.v3/rl10"

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/rightscale/rsc.v3/metadata"
	"gopkg.in/rightscale/rsc.v3/rsapi"
)

// An Href contains the relative path to a resource or resource collection,
// e.g. "/api/servers/123" or "/api/servers".
type Href string

// ActionPath computes the path to the given resource action. For example given the href
// "/api/servers/123" calling ActionPath with resource "servers" and action "clone" returns the path
// "/api/servers/123/clone" and verb POST.
// The algorithm consists of extracting the variables from the href by looking up a matching
// pattern from the resource metadata. The variables are then substituted in the action path.
// If there are more than one pattern that match the href then the algorithm picks the one that can
// substitute the most variables.
func (r *Href) ActionPath(rName, aName string) (*metadata.ActionPath, error) {
	res, ok := GenMetadata[rName]
	if !ok {
		return nil, fmt.Errorf("No resource with name '%s'", rName)
	}
	var action *metadata.Action
	for _, a := range res.Actions {
		if a.Name == aName {
			action = a
			break
		}
	}
	if action == nil {
		return nil, fmt.Errorf("No action with name '%s' on %s", aName, rName)
	}
	vars, err := res.ExtractVariables(string(*r))
	if err != nil {
		return nil, err
	}
	return action.Url(vars)
}

/******  DebugCookbookPath ******/

// Manipulate debug cookbook directory location
type DebugCookbookPath struct {
}

//===== Locator

// DebugCookbookPathLocator exposes the DebugCookbookPath resource actions.
type DebugCookbookPathLocator struct {
	Href
	api *Api
}

// DebugCookbookPathLocator builds a locator from the given href.
func (api *Api) DebugCookbookPathLocator(href string) *DebugCookbookPathLocator {
	return &DebugCookbookPathLocator{Href(href), api}
}

//===== Actions

// GET /rll/debug/cookbook
//
// Retrieve debug cookbook directory location
func (loc *DebugCookbookPathLocator) Show() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("DebugCookbookPath", "show")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

// PUT /rll/debug/cookbook
//
// Set debug cookbook directory location
func (loc *DebugCookbookPathLocator) Update(path string) (string, error) {
	var res string
	if path == "" {
		return res, fmt.Errorf("path is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"path": path,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("DebugCookbookPath", "update")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

// DELETE /rll/debug/cookbook
//
// Remove debug cookbook directory location
func (loc *DebugCookbookPathLocator) Delete() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("DebugCookbookPath", "delete")
	if err != nil {
		return err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

/******  Env ******/

// Manipulate global script environment variables
type Env struct {
}

//===== Locator

// EnvLocator exposes the Env resource actions.
type EnvLocator struct {
	Href
	api *Api
}

// EnvLocator builds a locator from the given href.
func (api *Api) EnvLocator(href string) *EnvLocator {
	return &EnvLocator{Href(href), api}
}

//===== Actions

// GET /rll/env
//
// Retrieve all environment variables
func (loc *EnvLocator) Index() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Env", "index")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

// GET /rll/env/:name
//
// Retrieve environment variable value
func (loc *EnvLocator) Show() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Env", "show")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

// PUT /rll/env/:name
//
// Set environment variable value
func (loc *EnvLocator) Update(payload string) (string, error) {
	var res string
	if payload == "" {
		return res, fmt.Errorf("payload is required")
	}
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	payloadParams = rsapi.ApiParams{
		"payload": payload,
	}
	uri, err := loc.ActionPath("Env", "update")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

// DELETE /rll/env/:name
//
// Delete environment variable
func (loc *EnvLocator) Delete() error {
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Env", "delete")
	if err != nil {
		return err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	return nil
}

/******  Proc ******/

// List of process variables, such as version, identity, and protocol_version
type Proc struct {
}

//===== Locator

// ProcLocator exposes the Proc resource actions.
type ProcLocator struct {
	Href
	api *Api
}

// ProcLocator builds a locator from the given href.
func (api *Api) ProcLocator(href string) *ProcLocator {
	return &ProcLocator{Href(href), api}
}

//===== Actions

// GET /rll/proc
//
// List all process variables
func (loc *ProcLocator) Index() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Proc", "index")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

// GET /rll/proc/:name
//
// Retrieve process variable value
func (loc *ProcLocator) Show() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Proc", "show")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

/******  Rl10 ******/

// Miscellaneous RightLink 10 local requests
type Rl10 struct {
}

//===== Locator

// Rl10Locator exposes the Rl10 resource actions.
type Rl10Locator struct {
	Href
	api *Api
}

// Rl10Locator builds a locator from the given href.
func (api *Api) Rl10Locator(href string) *Rl10Locator {
	return &Rl10Locator{Href(href), api}
}

//===== Actions

// POST /rll/upgrade
//
// Relaunch the RightLink process using a specified binary
func (loc *Rl10Locator) Upgrade(exec string) (string, error) {
	var res string
	if exec == "" {
		return res, fmt.Errorf("exec is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"exec": exec,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Rl10", "upgrade")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

// POST /rll/run/recipe
//
// Run git-based scripts (as recipes) synchronously
func (loc *Rl10Locator) RunRecipe(recipe string, options rsapi.ApiParams) (string, error) {
	var res string
	if recipe == "" {
		return res, fmt.Errorf("recipe is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"recipe": recipe,
	}
	var argumentsOpt = options["arguments"]
	if argumentsOpt != nil {
		queryParams["arguments"] = argumentsOpt
	}
	var jsonOpt = options["json"]
	if jsonOpt != nil {
		queryParams["json"] = jsonOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Rl10", "run_recipe")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

// POST /rll/run/right_script
//
// Run RightScripts synchronously
func (loc *Rl10Locator) RunRightScript(options rsapi.ApiParams) (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{}
	var argumentsOpt = options["arguments"]
	if argumentsOpt != nil {
		queryParams["arguments"] = argumentsOpt
	}
	var rightScriptOpt = options["right_script"]
	if rightScriptOpt != nil {
		queryParams["right_script"] = rightScriptOpt
	}
	var rightScriptIdOpt = options["right_script_id"]
	if rightScriptIdOpt != nil {
		queryParams["right_script_id"] = rightScriptIdOpt
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("Rl10", "run_right_script")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

/******  TSS ******/

// Manipulate the TSS proxy
type TSS struct {
}

//===== Locator

// TSSLocator exposes the TSS resource actions.
type TSSLocator struct {
	Href
	api *Api
}

// TSSLocator builds a locator from the given href.
func (api *Api) TSSLocator(href string) *TSSLocator {
	return &TSSLocator{Href(href), api}
}

//===== Actions

// GET /rll/tss/hostname
//
// Get the TSS hostname to proxy
func (loc *TSSLocator) GetHostname() (string, error) {
	var res string
	var queryParams rsapi.ApiParams
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("TSS", "get_hostname")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

// PUT /rll/tss/hostname
//
// Set the TSS hostname to proxy
func (loc *TSSLocator) PutHostname(hostname string) (string, error) {
	var res string
	if hostname == "" {
		return res, fmt.Errorf("hostname is required")
	}
	var queryParams rsapi.ApiParams
	queryParams = rsapi.ApiParams{
		"hostname": hostname,
	}
	var payloadParams rsapi.ApiParams
	uri, err := loc.ActionPath("TSS", "put_hostname")
	if err != nil {
		return res, err
	}
	resp, err := loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return res, fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	res = string(respBody)
	return res, err
}

/****** Parameter Data Types ******/
