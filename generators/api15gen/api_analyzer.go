package main

import (
	"fmt"
	"sort"
	"strings"

	"bitbucket.org/pkg/inflect"
)

// The analyzer struct holds the analysis results
type ApiAnalyzer struct {
	// Raw resources as defined in API json metadata
	rawResources map[string]interface{}
	// Attribute type mappings defined in attributes.json
	attributeTypes map[string]string
	// Temporary analysis construct
	// Holds all types indexed by name, multiple actions could generate types with the same
	// name. Keep them all here then make names unique as needed once we gathered all of them.
	rawTypes map[string][]*ObjectDataType
}

// The api descriptor struct contains the results of the analyzer Analyze() method.
// This includes all the API resources, actions and types.
type ApiDescriptor struct {
	// Resources indexed by name
	Resources map[string]*Resource
	// Resource names ordered alphabetically
	ResourceNames []string
	// Types used by resource actions indexed by name
	Types map[string]*ObjectDataType
	// Type names ordered alphabetically
	TypeNames []string
}

// Factory method for API analyzer
func NewApiAnalyzer(resources map[string]interface{}, attributeTypes map[string]string) *ApiAnalyzer {
	return &ApiAnalyzer{
		rawResources:   resources,
		attributeTypes: attributeTypes,
		rawTypes:       make(map[string][]*ObjectDataType),
	}
}

// Analyze iterate through all resources and initializes the Resources and ParamTypes fields of
// the ApiAnalyzer struct accordingly.
func (a *ApiAnalyzer) Analyze() *ApiDescriptor {
	var descriptor = &ApiDescriptor{
		Resources: make(map[string]*Resource),
		Types:     make(map[string]*ObjectDataType),
	}
	var rawResourceNames = make([]string, len(a.rawResources))
	var idx = 0
	for n, _ := range a.rawResources {
		rawResourceNames[idx] = n
		idx += 1
	}
	sort.Strings(rawResourceNames)
	for _, name := range rawResourceNames {
		var resource = a.rawResources[name]
		a.AnalyzeResource(name, resource, descriptor)
	}
	descriptor.FinalizeTypeNames(a.rawTypes)
	return descriptor
}

// AnalyzeResource analyzes the given resource and updates the Resources and ParamTypes analyzer
// fields accordingly
func (a *ApiAnalyzer) AnalyzeResource(name string, resource interface{}, descriptor *ApiDescriptor) {
	var res = resource.(map[string]interface{})

	// Compute description
	var description string
	if d, ok := res["description"].(string); ok {
		description = d
	}

	// Compute attributes
	var attributes []*Attribute
	var atts map[string]interface{}
	if m, ok := res["media_type"].(map[string]interface{}); ok {
		atts = m["attributes"].(map[string]interface{})
		attributes = make([]*Attribute, len(atts))
		for idx, n := range sortedKeys(atts) {
			attributes[idx] = &Attribute{attributeName(n), n, a.attributeTypes[n]}
		}
	} else {
		attributes = []*Attribute{}
	}

	// Compute hrefs
	var baseHref, resourceHref string
	var methods = res["methods"].(map[string]interface{})
	if index, ok := methods["index"].(map[string]interface{}); ok {
		_, baseHref = parseRoute("", index["route"].(string))
	} else if show, ok := methods["show"].(map[string]interface{}); ok {
		_, baseHref = parseRoute("", show["route"].(string))
		var index = strings.LastIndex(baseHref, "/")
		baseHref = baseHref[:index]
	}
	if baseHref != "" {
		resourceHref = baseHref + "/:id"
	}

	// Compute actions
	var actionNames = sortedKeys(methods)
	var resourceActions = []*Action{}
	var collectionActions = []*Action{}
	for _, actionName := range actionNames {
		var m = methods[actionName]
		var meth = m.(map[string]interface{})
		var params map[string]interface{}
		if p, ok := meth["parameters"]; ok {
			params = p.(map[string]interface{})
		}
		var description = "No description provided for " + actionName + "."
		if d, _ := meth["description"]; d != nil {
			description = d.(string)
		}
		var isResourceAction bool
		httpMethod, path := parseRoute(fmt.Sprintf("%s#%s", name, actionName),
			meth["route"].(string))
		if len(baseHref) > 0 {
			if strings.HasPrefix(path, resourceHref) {
				path = path[len(resourceHref):]
				isResourceAction = true
			} else if strings.HasSuffix(path, baseHref) {
				path = path[len(baseHref):]
			}
		}
		var contentType string
		if c, ok := meth["content_type"].(string); ok {
			contentType = c
		}
		var paramAnalyzer = NewAnalyzer(params)
		paramAnalyzer.Analyze()

		// Record new parameter types
		var paramTypeNames = make([]string, len(paramAnalyzer.ParamTypes))
		var idx = 0
		for n, _ := range paramAnalyzer.ParamTypes {
			paramTypeNames[idx] = n
			idx += 1
		}
		sort.Strings(paramTypeNames)
		for _, name := range paramTypeNames {
			var pType = paramAnalyzer.ParamTypes[name]
			if _, ok := a.rawTypes[name]; ok {
				a.rawTypes[name] = append(a.rawTypes[name], pType)
			} else {
				a.rawTypes[name] = []*ObjectDataType{pType}
			}
		}

		// Update description with parameter descriptions
		var mandatory = []string{}
		var optional = []string{}
		for _, p := range append(paramAnalyzer.QueryParams, paramAnalyzer.PayloadParams...) {
			if p.Mandatory {
				if p.Description != "" {
					var desc = fmt.Sprintf("%s: %s", p.VarName, p.Description)
					mandatory = append(mandatory, desc)
				}
			} else {
				var desc = p.VarName
				if p.Description != "" {
					desc += ": " + p.Description
				}
				optional = append(optional, desc)
			}
		}
		sort.Strings(mandatory)
		sort.Strings(optional)
		if len(mandatory) > 0 {
			description += "\n\t" + strings.Join(mandatory, "\n\t")
		}
		if len(optional) > 0 {
			description += "\n-- Optional parameters:\n\t" +
				strings.Join(optional, "\n\t")
		}

		// Record action
		var action = Action{
			Name:          actionName,
			MethodName:    methodName(actionName, name),
			Description:   removeBlankLines(description),
			HttpMethod:    httpMethod,
			Path:          path,
			PayloadParams: paramAnalyzer.PayloadParams,
			QueryParams:   paramAnalyzer.QueryParams,
			Return:        parseReturn(actionName, name, contentType),
		}
		if isResourceAction {
			resourceActions = append(resourceActions, &action)
		} else {
			collectionActions = append(collectionActions, &action)
		}
	}

	// We're done!
	name = inflect.Singularize(name)
	descriptor.Resources[name] = &Resource{
		Name:              name,
		CollectionName:    inflect.Pluralize(name),
		Description:       removeBlankLines(description),
		BaseHref:          baseHref,
		ResourceActions:   resourceActions,
		CollectionActions: collectionActions,
		Attributes:        attributes,
	}
}

// Go through all the types generated by the analyzer and generate unique names
func (d *ApiDescriptor) FinalizeTypeNames(rawTypes map[string][]*ObjectDataType) {

	// 1. Make sure data type names don't clash with resource names
	var rawTypeNames = make([]string, len(rawTypes))
	var idx = 0
	for n, _ := range rawTypes {
		rawTypeNames[idx] = n
		idx += 1
	}
	sort.Strings(rawTypeNames)
	for _, tn := range rawTypeNames {
		var types = rawTypes[tn]
		for rn, _ := range d.Resources {
			if tn == rn {
				var oldTn = tn
				if strings.HasSuffix(tn, "Param") {
					tn = fmt.Sprintf("%s2", tn)
				} else {
					tn = fmt.Sprintf("%sParam", tn)
				}
				for _, ty := range types {
					ty.Name = tn
				}
				rawTypes[tn] = types
				delete(rawTypes, oldTn)
			}
		}
	}

	// 2. Make data type names unique
	idx = 0
	for n, _ := range rawTypes {
		rawTypeNames[idx] = n
		idx += 1
	}
	sort.Strings(rawTypeNames)
	for _, tn := range rawTypeNames {
		var types = rawTypes[tn]
		var first = types[0]
		d.Types[tn] = first
		if len(types) > 1 {
			for i, ty := range types[1:] {
				var found = false
				for j := 0; j < i; j++ {
					if ty.IsEquivalent(types[j]) {
						found = true
						break
					}
				}
				if !found {
					var newName = d.uniqueTypeName(tn)
					ty.Name = newName
					d.Types[newName] = ty
				}
			}
		}
	}

	// 3. Finally initialize .ResourceNames and .TypeNames
	idx = 0
	var resourceNames = make([]string, len(d.Resources))
	for n, _ := range d.Resources {
		resourceNames[idx] = n
		idx += 1
	}
	sort.Strings(resourceNames)
	d.ResourceNames = resourceNames

	var typeNames = make([]string, len(d.Types))
	idx = 0
	for tn, _ := range d.Types {
		typeNames[idx] = tn
		idx += 1
	}
	sort.Strings(typeNames)
	d.TypeNames = typeNames
}

// Build unique type name by appending "next available index" to given prefix
func (d *ApiDescriptor) uniqueTypeName(prefix string) string {
	var u = fmt.Sprintf("%s%d", prefix, 2)
	var taken = false
	for _, tn := range d.TypeNames {
		if tn == u {
			taken = true
			break
		}
	}
	var idx = 3
	for taken {
		u = fmt.Sprintf("%s%d", prefix, idx)
		taken = false
		for _, tn := range d.TypeNames {
			if tn == u {
				taken = true
				break
			}
		}
		if taken {
			idx += 1
		}
	}
	return u
}

/** Helper methods for parsing raw JSON **/

func parseRoute(moniker string, route string) (method string, path string) {
	// :(((( some routes are empty
	switch moniker {
	case "Deployments#servers":
		return "GET", "api/deployments/:id/servers"
	case "ServerArrays#current_instances":
		return "GET", "api/server_arrays/:id/current_instances"
	case "ServerArrays#launch":
		return "POST", "api/server_arrays/:id/launch"
	case "ServerArrays#multi_run_executable":
		return "POST", "api/server_arrays/:id/multi_run_executable"
	case "ServerArrays#multi_terminate":
		return "POST", "api/server_arrays/:id/multi_terminate"
	case "Servers#launch":
		return "POST", "api/servers/:id/launch"
	case "Servers#terminate":
		return "POST", "api/servers/:id/teminate"

	}
	method = strings.TrimRight(route[0:7], " ")
	path = strings.TrimRight(strings.Split(route[8:], "{")[0], " ")

	return
}

// Resources that don't have a media type
var noMediaTypeResources = map[string]bool{
	"HealthCheck":          true,
	"Oauth2":               true,
	"Tag":                  true,
	"UserDatas":            true,
	"ChildAccounts":        true,
	"MonitoringMetricData": true,
	"ImportPreview":        true,
	"Changes":              true,
	"CookbookResolution":   true,
	"ResourceTag":          true,
}

func parseReturn(kind, resName, contentType string) string {
	switch kind {
	case "show":
		return fmt.Sprintf("*%s", resourceType(resName))
	case "index":
		return fmt.Sprintf("[]%s", resourceType(resName))
	case "create":
		if _, ok := noMediaTypeResources[resName]; ok {
			return "map[string]interface{}"
		} else {
			return "Href"
		}
	case "update", "destroy":
		return ""
	default:
		switch {
		case len(contentType) == 0:
			return ""
		case strings.Index(contentType, "application/vnd.rightscale.") == 0:
			if contentType == "application/vnd.rightscale.text" {
				return "string"
			}
			var elems = strings.SplitN(contentType[27:], ";", 2)
			var name = resourceType(inflect.Camelize(elems[0]))
			if len(elems) > 1 && elems[1] == "type=collection" {
				name = "[]" + name
			}
			return name
		default: // Shouldn't be here
			panic("api15gen: Unknown content type " + contentType)
		}
	}

}

// Name of go type for resource with given name
// It should always be the same (camelized) but there are some resources that don't have a media
// type so for these we use a map.
func resourceType(resName string) string {
	if _, ok := noMediaTypeResources[resName]; ok {
		return "map[string]string"
	} else {
		return inflect.Singularize(resName)
	}
}

// Heuristic to create method name from action and resource names
func methodName(action, resource string) string {
	if action != "index" && strings.Index(action, "multi") != 0 {
		resource = inflect.Singularize(resource)
	}
	return inflect.Camelize(action) + inflect.Camelize(resource)
}

// Escape attribute names using go keywords
func attributeName(name string) string {
	return inflect.Camelize(name)
}

// Return keys of given maps sorted
func sortedKeys(m map[string]interface{}) []string {
	var keys = make([]string, len(m))
	var idx = 0
	for k, _ := range m {
		keys[idx] = k
		idx += 1
	}
	sort.Strings(keys)
	return keys
}