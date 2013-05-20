package web

import (
	"regexp"
)

type Route struct {
	method   string
	pattern  string
	Callback func(*Response, *Request)
	params   map[string]string
}

func (r *Route) matches(uri string) bool {
	reg := regexp.MustCompile(":([^/]+)")
	// replace ":name" with "(?P<name>)"
	replace := reg.ReplaceAllString(r.pattern, "(?P<$1>[^/]+)")
	reg = regexp.MustCompile(replace + "(/{0,1})$") // Allow leading slash
	// Does the route match?
	if reg.MatchString(uri) {
		// Find all parameters
		params := reg.FindAllStringSubmatch(uri, -1)
		paramValues := params[0]
		paramNames := reg.SubexpNames()

		// Put them in a [string] string map
		paramMap := map[string]string{}
		for i := 1; i < len(paramValues); i++ {
			paramMap[paramNames[i]] = paramValues[i]
		}

		r.params = paramMap

		return true
	}
	return false
}
