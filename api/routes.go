package api

import (
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []route{
	route{
		"createRepo",
		"POST",
		"/repositories",
		CreateRepo,
	},
	route{
		"listRepos",
		"GET",
		"/repositories",
		ListRepos,
	},
	route{
		"addRPM",
		"POST",
		"/repositories/{name}/packages",
		AddRPM,
	},
	route{
		"listRPMs",
		"GET",
		"/repositories/{name}/packages",
		ListPackages,
	},
}
