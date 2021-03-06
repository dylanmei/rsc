//************************************************************************//
//                     RightScale API client
//
// Generated with:
// $ praxisgen -metadata=ss/ssc/restful_doc -output=ss/ssc -pkg=ssc -target=1.0 -client=Api
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package ssc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// API Version
const APIVersion = "1.0"

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

/******  AccountPreference ******/

// The AccountPreference resource stores preferences that apply account-wide, such as UI customization settings and other settings.
// The Self-Service portal uses some of these preferences in the portal itself, and this resource allows you to extend the settings
// to use in your own integration.
type AccountPreference struct {
	CreatedBy  *User             `json:"created_by,omitempty"`
	GroupName  string            `json:"group_name,omitempty"`
	Href       string            `json:"href,omitempty"`
	Kind       string            `json:"kind,omitempty"`
	Name       string            `json:"name,omitempty"`
	Timestamps *TimestampsStruct `json:"timestamps,omitempty"`
	Value      string            `json:"value,omitempty"`
}

// Locator returns a locator for the given resource
func (r *AccountPreference) Locator(api *Api) *AccountPreferenceLocator {
	return api.AccountPreferenceLocator(r.Href)
}

//===== Locator

// AccountPreferenceLocator exposes the AccountPreference resource actions.
type AccountPreferenceLocator struct {
	Href
	api *Api
}

// AccountPreferenceLocator builds a locator from the given href.
func (api *Api) AccountPreferenceLocator(href string) *AccountPreferenceLocator {
	return &AccountPreferenceLocator{Href(href), api}
}

//===== Actions

// GET /accounts/:account_id/account_preferences
//
// List the AccountPreferences for this account.
func (loc *AccountPreferenceLocator) Index(options rsapi.ApiParams) ([]*AccountPreference, error) {
	var res []*AccountPreference
	var params rsapi.ApiParams
	params = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var groupOpt = options["group"]
	if groupOpt != nil {
		params["group"] = groupOpt
	}
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("AccountPreference", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /accounts/:account_id/account_preferences/:name
//
// Get details for a particular AccountPreference
func (loc *AccountPreferenceLocator) Show() (*AccountPreference, error) {
	var res *AccountPreference
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("AccountPreference", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /accounts/:account_id/account_preferences
//
// Create a new AccountPreference or update an existing AccountPreference with the new value
func (loc *AccountPreferenceLocator) Create(groupName string, name string, value string) (*AccountPreferenceLocator, error) {
	var res *AccountPreferenceLocator
	if groupName == "" {
		return res, fmt.Errorf("groupName is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	if value == "" {
		return res, fmt.Errorf("value is required")
	}
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	p = rsapi.ApiParams{
		"group_name": groupName,
		"name":       name,
		"value":      value,
	}
	uri, err := loc.ActionPath("AccountPreference", "create")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &AccountPreferenceLocator{Href(location), loc.api}, nil
	}
}

// DELETE /accounts/:account_id/account_preferences/:name
//
// Delete an AccountPreference
func (loc *AccountPreferenceLocator) Delete() error {
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("AccountPreference", "delete")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

/******  Application ******/

// An Application is an element in the Catalog that can be launched by users. Applications are generally created by uploading CAT
// files to the Designer and publishing them to the Catalog, though they can also be created via API calls to the Catalog directly without
// going through Designer. If an Application was created from Designer through the publish action, it contains a link back to the Template
// resource in Designer.
// In the Self-Service portal, an Application is equivalent to an item in the Catalog. Most users have access to these Application resources
// and can launch them to create Executions in the Manager application.
type Application struct {
	CompiledCat        string            `json:"compiled_cat,omitempty"`
	CreatedBy          *User             `json:"created_by,omitempty"`
	Href               string            `json:"href,omitempty"`
	Id                 string            `json:"id,omitempty"`
	Kind               string            `json:"kind,omitempty"`
	LongDescription    string            `json:"long_description,omitempty"`
	Name               string            `json:"name,omitempty"`
	Parameters         []*Parameter      `json:"parameters,omitempty"`
	RequiredParameters []string          `json:"required_parameters,omitempty"`
	ScheduleRequired   bool              `json:"schedule_required,omitempty"`
	Schedules          []*Schedule       `json:"schedules,omitempty"`
	ShortDescription   string            `json:"short_description,omitempty"`
	TemplateInfo       *TemplateInfo     `json:"template_info,omitempty"`
	Timestamps         *TimestampsStruct `json:"timestamps,omitempty"`
}

// Locator returns a locator for the given resource
func (r *Application) Locator(api *Api) *ApplicationLocator {
	return api.ApplicationLocator(r.Href)
}

//===== Locator

// ApplicationLocator exposes the Application resource actions.
type ApplicationLocator struct {
	Href
	api *Api
}

// ApplicationLocator builds a locator from the given href.
func (api *Api) ApplicationLocator(href string) *ApplicationLocator {
	return &ApplicationLocator{Href(href), api}
}

//===== Actions

// GET /catalogs/:catalog_id/applications
//
// List the Applications available in the specified Catalog.
func (loc *ApplicationLocator) Index(options rsapi.ApiParams) ([]*Application, error) {
	var res []*Application
	var params rsapi.ApiParams
	params = rsapi.ApiParams{}
	var idsOpt = options["ids"]
	if idsOpt != nil {
		params["ids[]"] = idsOpt
	}
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("Application", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /catalogs/:catalog_id/applications/:id
//
// Show detailed information about a given Application.
func (loc *ApplicationLocator) Show(options rsapi.ApiParams) (*Application, error) {
	var res *Application
	var params rsapi.ApiParams
	params = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("Application", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /catalogs/:catalog_id/applications
//
// Create a new Application in the Catalog.
func (loc *ApplicationLocator) Create(compiledCat *CompiledCAT, name string, shortDescription string, options rsapi.ApiParams) (*ApplicationLocator, error) {
	var res *ApplicationLocator
	if compiledCat == nil {
		return res, fmt.Errorf("compiledCat is required")
	}
	if name == "" {
		return res, fmt.Errorf("name is required")
	}
	if shortDescription == "" {
		return res, fmt.Errorf("shortDescription is required")
	}
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	p = rsapi.ApiParams{
		"compiled_cat":      compiledCat,
		"name":              name,
		"short_description": shortDescription,
	}
	var longDescriptionOpt = options["long_description"]
	if longDescriptionOpt != nil {
		p["long_description"] = longDescriptionOpt
	}
	var scheduleRequiredOpt = options["schedule_required"]
	if scheduleRequiredOpt != nil {
		p["schedule_required"] = scheduleRequiredOpt
	}
	var schedulesOpt = options["schedules"]
	if schedulesOpt != nil {
		p["schedules"] = schedulesOpt
	}
	var templateHrefOpt = options["template_href"]
	if templateHrefOpt != nil {
		p["template_href"] = templateHrefOpt
	}
	uri, err := loc.ActionPath("Application", "create")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &ApplicationLocator{Href(location), loc.api}, nil
	}
}

// PUT /catalogs/:catalog_id/applications/:id
//
// Update the content of an existing Application.
func (loc *ApplicationLocator) Update(options rsapi.ApiParams) error {
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	p = rsapi.ApiParams{}
	var compiledCatOpt = options["compiled_cat"]
	if compiledCatOpt != nil {
		p["compiled_cat"] = compiledCatOpt
	}
	var longDescriptionOpt = options["long_description"]
	if longDescriptionOpt != nil {
		p["long_description"] = longDescriptionOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var scheduleRequiredOpt = options["schedule_required"]
	if scheduleRequiredOpt != nil {
		p["schedule_required"] = scheduleRequiredOpt
	}
	var schedulesOpt = options["schedules"]
	if schedulesOpt != nil {
		p["schedules"] = schedulesOpt
	}
	var shortDescriptionOpt = options["short_description"]
	if shortDescriptionOpt != nil {
		p["short_description"] = shortDescriptionOpt
	}
	var templateHrefOpt = options["template_href"]
	if templateHrefOpt != nil {
		p["template_href"] = templateHrefOpt
	}
	uri, err := loc.ActionPath("Application", "update")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

// PUT /catalogs/:catalog_id/applications
//
// Update the content of multiple Applications.
func (loc *ApplicationLocator) MultiUpdate(id string, options rsapi.ApiParams) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	p = rsapi.ApiParams{
		"id": id,
	}
	var compiledCatOpt = options["compiled_cat"]
	if compiledCatOpt != nil {
		p["compiled_cat"] = compiledCatOpt
	}
	var longDescriptionOpt = options["long_description"]
	if longDescriptionOpt != nil {
		p["long_description"] = longDescriptionOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var scheduleRequiredOpt = options["schedule_required"]
	if scheduleRequiredOpt != nil {
		p["schedule_required"] = scheduleRequiredOpt
	}
	var schedulesOpt = options["schedules"]
	if schedulesOpt != nil {
		p["schedules"] = schedulesOpt
	}
	var shortDescriptionOpt = options["short_description"]
	if shortDescriptionOpt != nil {
		p["short_description"] = shortDescriptionOpt
	}
	var templateHrefOpt = options["template_href"]
	if templateHrefOpt != nil {
		p["template_href"] = templateHrefOpt
	}
	uri, err := loc.ActionPath("Application", "multi_update")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

// DELETE /catalogs/:catalog_id/applications/:id
//
// Delete an Application from the Catalog
func (loc *ApplicationLocator) Delete() error {
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("Application", "delete")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

// DELETE /catalogs/:catalog_id/applications
//
// Delete multiple Applications from the Catalog
func (loc *ApplicationLocator) MultiDelete(ids []string) error {
	if len(ids) == 0 {
		return fmt.Errorf("ids is required")
	}
	var params rsapi.ApiParams
	params = rsapi.ApiParams{
		"ids[]": ids,
	}
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("Application", "multi_delete")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

// GET /catalogs/:catalog_id/applications/:id/download
//
// Download the underlying CAT source of an Application.
func (loc *ApplicationLocator) Download(apiVersion string) error {
	if apiVersion == "" {
		return fmt.Errorf("apiVersion is required")
	}
	var params rsapi.ApiParams
	params = rsapi.ApiParams{
		"api_version": apiVersion,
	}
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("Application", "download")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

// POST /catalogs/:catalog_id/applications/:id/actions/launch
//
// Launches an Application by creating an Execution with ScheduledActions as needed to match the optional Schedule provided.
func (loc *ApplicationLocator) Launch(options rsapi.ApiParams) error {
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	p = rsapi.ApiParams{}
	var descriptionOpt = options["description"]
	if descriptionOpt != nil {
		p["description"] = descriptionOpt
	}
	var endDateOpt = options["end_date"]
	if endDateOpt != nil {
		p["end_date"] = endDateOpt
	}
	var nameOpt = options["name"]
	if nameOpt != nil {
		p["name"] = nameOpt
	}
	var options_Opt = options["options"]
	if options_Opt != nil {
		p["options"] = options_Opt
	}
	var scheduleNameOpt = options["schedule_name"]
	if scheduleNameOpt != nil {
		p["schedule_name"] = scheduleNameOpt
	}
	uri, err := loc.ActionPath("Application", "launch")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

/******  NotificationRule ******/

// A notification rule describes which notification should be created
// when events occur in the system. Events may be generated when an
// execution status changes or when an operation fails for example.
// A rule has a source which can be a specific resource or a group of
// resources (described via a link-like syntax), a target which
// corresponds to a user (for now) and a minimum severity used to filter
// out events with lower severities.
type NotificationRule struct {
	AccountId   string            `json:"account_id,omitempty"`
	Href        string            `json:"href,omitempty"`
	Id          string            `json:"id,omitempty"`
	Kind        string            `json:"kind,omitempty"`
	MinSeverity string            `json:"min_severity,omitempty"`
	Priority    int               `json:"priority,omitempty"`
	Source      string            `json:"source,omitempty"`
	Target      string            `json:"target,omitempty"`
	Timestamps  *TimestampsStruct `json:"timestamps,omitempty"`
}

// Locator returns a locator for the given resource
func (r *NotificationRule) Locator(api *Api) *NotificationRuleLocator {
	return api.NotificationRuleLocator(r.Href)
}

//===== Locator

// NotificationRuleLocator exposes the NotificationRule resource actions.
type NotificationRuleLocator struct {
	Href
	api *Api
}

// NotificationRuleLocator builds a locator from the given href.
func (api *Api) NotificationRuleLocator(href string) *NotificationRuleLocator {
	return &NotificationRuleLocator{Href(href), api}
}

//===== Actions

// GET /accounts/:account_id/notification_rules
//
// List all notification rules, potentially filtering by a collection of resources.
func (loc *NotificationRuleLocator) Index(source string, options rsapi.ApiParams) ([]*NotificationRule, error) {
	var res []*NotificationRule
	if source == "" {
		return res, fmt.Errorf("source is required")
	}
	var params rsapi.ApiParams
	params = rsapi.ApiParams{
		"source": source,
	}
	var targetsOpt = options["targets"]
	if targetsOpt != nil {
		params["targets"] = targetsOpt
	}
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("NotificationRule", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /accounts/:account_id/notification_rules
//
// Create one notification rule for a specific target and source.
// The source must be unique in the scope of target and account.
func (loc *NotificationRuleLocator) Create(minSeverity string, source string, target string, options rsapi.ApiParams) (*NotificationRuleLocator, error) {
	var res *NotificationRuleLocator
	if minSeverity == "" {
		return res, fmt.Errorf("minSeverity is required")
	}
	if source == "" {
		return res, fmt.Errorf("source is required")
	}
	if target == "" {
		return res, fmt.Errorf("target is required")
	}
	var params rsapi.ApiParams
	params = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.ApiParams
	p = rsapi.ApiParams{
		"min_severity": minSeverity,
		"source":       source,
		"target":       target,
	}
	uri, err := loc.ActionPath("NotificationRule", "create")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &NotificationRuleLocator{Href(location), loc.api}, nil
	}
}

// PATCH /accounts/:account_id/notification_rules/:id
//
// Change min severity of existing rule
func (loc *NotificationRuleLocator) Patch(minSeverity string) error {
	if minSeverity == "" {
		return fmt.Errorf("minSeverity is required")
	}
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	p = rsapi.ApiParams{
		"min_severity": minSeverity,
	}
	uri, err := loc.ActionPath("NotificationRule", "patch")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

// GET /accounts/:account_id/notification_rules/:id
//
// Show one notification rule.
func (loc *NotificationRuleLocator) Show() (*NotificationRule, error) {
	var res *NotificationRule
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("NotificationRule", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// DELETE /accounts/:account_id/notification_rules/:id
//
// Delete one notification rule.
func (loc *NotificationRuleLocator) Delete() error {
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("NotificationRule", "delete")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

// DELETE /accounts/:account_id/notification_rules
//
// Delete one or more notification rules by id or source and target.
func (loc *NotificationRuleLocator) MultiDelete(options rsapi.ApiParams) error {
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	p = rsapi.ApiParams{}
	var idOpt = options["id"]
	if idOpt != nil {
		p["id"] = idOpt
	}
	var sourceOpt = options["source"]
	if sourceOpt != nil {
		p["source"] = sourceOpt
	}
	var targetOpt = options["target"]
	if targetOpt != nil {
		p["target"] = targetOpt
	}
	uri, err := loc.ActionPath("NotificationRule", "multi_delete")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

/******  UserPreference ******/

// The UserPreference resource stores preferences on a per user basis, such as default notification preference.
// The Self-Service portal uses these preferences in the portal.
type UserPreference struct {
	CreatedBy          *User               `json:"created_by,omitempty"`
	Href               string              `json:"href,omitempty"`
	Id                 string              `json:"id,omitempty"`
	Kind               string              `json:"kind,omitempty"`
	Timestamps         *TimestampsStruct   `json:"timestamps,omitempty"`
	UserId             int                 `json:"user_id,omitempty"`
	UserPreferenceInfo *UserPreferenceInfo `json:"user_preference_info,omitempty"`
	Value              string              `json:"value,omitempty"`
}

// Locator returns a locator for the given resource
func (r *UserPreference) Locator(api *Api) *UserPreferenceLocator {
	return api.UserPreferenceLocator(r.Href)
}

//===== Locator

// UserPreferenceLocator exposes the UserPreference resource actions.
type UserPreferenceLocator struct {
	Href
	api *Api
}

// UserPreferenceLocator builds a locator from the given href.
func (api *Api) UserPreferenceLocator(href string) *UserPreferenceLocator {
	return &UserPreferenceLocator{Href(href), api}
}

//===== Actions

// GET /accounts/:account_id/user_preferences
//
// List the UserPreference for users in this account.
// Only administrators and infrastructure users may request the preferences of other users.
// Users who are not members of the admin role need to specify a filter with their ID in order to retrieve their preferences.
func (loc *UserPreferenceLocator) Index(options rsapi.ApiParams) ([]*UserPreference, error) {
	var res []*UserPreference
	var params rsapi.ApiParams
	params = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("UserPreference", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /accounts/:account_id/user_preferences/:id
//
// Get details for a particular UserPreference
func (loc *UserPreferenceLocator) Show(options rsapi.ApiParams) (*UserPreference, error) {
	var res *UserPreference
	var params rsapi.ApiParams
	params = rsapi.ApiParams{}
	var viewOpt = options["view"]
	if viewOpt != nil {
		params["view"] = viewOpt
	}
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("UserPreference", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// POST /accounts/:account_id/user_preferences
//
// Create a new UserPreference.
// Multiple resources can be created at once with a multipart request.
// Values are validated with the corresponding UserPreferenceInfo.
func (loc *UserPreferenceLocator) Create(userId string, userPreferenceInfoId string, value interface{}) (*UserPreferenceLocator, error) {
	var res *UserPreferenceLocator
	if userId == "" {
		return res, fmt.Errorf("userId is required")
	}
	if userPreferenceInfoId == "" {
		return res, fmt.Errorf("userPreferenceInfoId is required")
	}
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	p = rsapi.ApiParams{
		"user_id":                 userId,
		"user_preference_info_id": userPreferenceInfoId,
		"value":                   value,
	}
	uri, err := loc.ActionPath("UserPreference", "create")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &UserPreferenceLocator{Href(location), loc.api}, nil
	}
}

// PATCH /accounts/:account_id/user_preferences/:id
//
// Update the value of a UserPreference.
// Multiple values may be updated using a multipart request.
// Values are validated with the corresponding UserPreferenceInfo.
func (loc *UserPreferenceLocator) Update(value interface{}, options rsapi.ApiParams) error {
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	p = rsapi.ApiParams{
		"value": value,
	}
	var idOpt = options["id"]
	if idOpt != nil {
		p["id"] = idOpt
	}
	uri, err := loc.ActionPath("UserPreference", "update")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

// DELETE /accounts/:account_id/user_preferences/:id
//
// Delete a UserPreference
func (loc *UserPreferenceLocator) Delete() error {
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("UserPreference", "delete")
	if err != nil {
		return err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return err
	}
	resp, err := loc.api.PerformRequest(req)
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

/******  UserPreferenceInfo ******/

// The UserPreferenceInfo resource defines the available user preferences supported by the system.
// It is also used to validate values saved in UserPreference.
type UserPreferenceInfo struct {
	Category        string            `json:"category,omitempty"`
	DefaultValue    string            `json:"default_value,omitempty"`
	DisplayName     string            `json:"display_name,omitempty"`
	HelpText        string            `json:"help_text,omitempty"`
	Href            string            `json:"href,omitempty"`
	Id              string            `json:"id,omitempty"`
	Kind            string            `json:"kind,omitempty"`
	Name            string            `json:"name,omitempty"`
	ValueConstraint []string          `json:"value_constraint,omitempty"`
	ValueRange      *ValueRangeStruct `json:"value_range,omitempty"`
	ValueType       string            `json:"value_type,omitempty"`
}

// Locator returns a locator for the given resource
func (r *UserPreferenceInfo) Locator(api *Api) *UserPreferenceInfoLocator {
	return api.UserPreferenceInfoLocator(r.Href)
}

//===== Locator

// UserPreferenceInfoLocator exposes the UserPreferenceInfo resource actions.
type UserPreferenceInfoLocator struct {
	Href
	api *Api
}

// UserPreferenceInfoLocator builds a locator from the given href.
func (api *Api) UserPreferenceInfoLocator(href string) *UserPreferenceInfoLocator {
	return &UserPreferenceInfoLocator{Href(href), api}
}

//===== Actions

// GET /accounts/:account_id/user_preference_infos
//
// List the UserPreferenceInfo.
func (loc *UserPreferenceInfoLocator) Index(options rsapi.ApiParams) ([]*UserPreferenceInfo, error) {
	var res []*UserPreferenceInfo
	var params rsapi.ApiParams
	params = rsapi.ApiParams{}
	var filterOpt = options["filter"]
	if filterOpt != nil {
		params["filter[]"] = filterOpt
	}
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("UserPreferenceInfo", "index")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	err = json.Unmarshal(respBody, &res)
	return res, err
}

// GET /accounts/:account_id/user_preference_infos/:id
//
// Get details for a particular UserPreferenceInfo
func (loc *UserPreferenceInfoLocator) Show() (*UserPreferenceInfo, error) {
	var res *UserPreferenceInfo
	var params rsapi.ApiParams
	var p rsapi.ApiParams
	uri, err := loc.ActionPath("UserPreferenceInfo", "show")
	if err != nil {
		return res, err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HttpMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return res, err
	}
	resp, err := loc.api.PerformRequest(req)
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
	err = json.Unmarshal(respBody, &res)
	return res, err
}

/****** Parameter Data Types ******/

type CompiledCAT struct {
	Conditions         map[string]interface{} `json:"conditions,omitempty"`
	Definitions        map[string]interface{} `json:"definitions,omitempty"`
	LongDescription    string                 `json:"long_description,omitempty"`
	Mappings           map[string]interface{} `json:"mappings,omitempty"`
	Name               string                 `json:"name,omitempty"`
	Namespaces         []*Namespace           `json:"namespaces,omitempty"`
	Operations         map[string]interface{} `json:"operations,omitempty"`
	Outputs            map[string]interface{} `json:"outputs,omitempty"`
	Parameters         map[string]interface{} `json:"parameters,omitempty"`
	RequiredParameters []string               `json:"required_parameters,omitempty"`
	Resources          map[string]interface{} `json:"resources,omitempty"`
	RsCaVer            int                    `json:"rs_ca_ver,omitempty"`
	ShortDescription   string                 `json:"short_description,omitempty"`
	Source             string                 `json:"source,omitempty"`
}

type ConfigurationOption struct {
	Name  string      `json:"name,omitempty"`
	Type_ string      `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Namespace struct {
	Name    string            `json:"name,omitempty"`
	Service *NamespaceService `json:"service,omitempty"`
	Types   []*NamespaceType  `json:"types,omitempty"`
}

type NamespaceService struct {
	Auth    string                 `json:"auth,omitempty"`
	Headers map[string]interface{} `json:"headers,omitempty"`
	Host    string                 `json:"host,omitempty"`
	Path    string                 `json:"path,omitempty"`
}

type NamespaceType struct {
	Delete    string                 `json:"delete,omitempty"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Name      string                 `json:"name,omitempty"`
	Path      string                 `json:"path,omitempty"`
	Provision string                 `json:"provision,omitempty"`
}

type Parameter struct {
	Default     interface{}                 `json:"default,omitempty"`
	Description string                      `json:"description,omitempty"`
	Name        string                      `json:"name,omitempty"`
	Operations  []interface{}               `json:"operations,omitempty"`
	Type_       string                      `json:"type,omitempty"`
	Ui          *ParametersUiStruct         `json:"ui,omitempty"`
	Validation  *ParametersValidationStruct `json:"validation,omitempty"`
}

type ParametersUiStruct struct {
	Category string `json:"category,omitempty"`
	Index    int    `json:"index,omitempty"`
	Label    string `json:"label,omitempty"`
}

type ParametersValidationStruct struct {
	AllowedPattern        string        `json:"allowed_pattern,omitempty"`
	AllowedValues         []interface{} `json:"allowed_values,omitempty"`
	ConstraintDescription string        `json:"constraint_description,omitempty"`
	MaxLength             int           `json:"max_length,omitempty"`
	MaxValue              int           `json:"max_value,omitempty"`
	MinLength             int           `json:"min_length,omitempty"`
	MinValue              int           `json:"min_value,omitempty"`
	NoEcho                bool          `json:"no_echo,omitempty"`
}

type Recurrence struct {
	Hour   int    `json:"hour,omitempty"`
	Minute int    `json:"minute,omitempty"`
	Rule   string `json:"rule,omitempty"`
}

type Schedule struct {
	CreatedFrom     string      `json:"created_from,omitempty"`
	Description     string      `json:"description,omitempty"`
	Name            string      `json:"name,omitempty"`
	StartRecurrence *Recurrence `json:"start_recurrence,omitempty"`
	StopRecurrence  *Recurrence `json:"stop_recurrence,omitempty"`
}

type TemplateInfo struct {
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}

type TimestampsStruct struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type User struct {
	Email string `json:"email,omitempty"`
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
}

type UserPreferenceInfoParam struct {
	Category        string                              `json:"category,omitempty"`
	DefaultValue    string                              `json:"default_value,omitempty"`
	DisplayName     string                              `json:"display_name,omitempty"`
	HelpText        string                              `json:"help_text,omitempty"`
	Href            string                              `json:"href,omitempty"`
	Id              string                              `json:"id,omitempty"`
	Kind            string                              `json:"kind,omitempty"`
	Name            string                              `json:"name,omitempty"`
	ValueConstraint []string                            `json:"value_constraint,omitempty"`
	ValueRange      *UserPreferenceInfoValueRangeStruct `json:"value_range,omitempty"`
	ValueType       string                              `json:"value_type,omitempty"`
}

type UserPreferenceInfoValueRangeStruct struct {
	Max int `json:"max,omitempty"`
	Min int `json:"min,omitempty"`
}

type ValueRangeStruct struct {
	Max int `json:"max,omitempty"`
	Min int `json:"min,omitempty"`
}
