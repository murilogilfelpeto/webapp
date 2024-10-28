package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
	"testing"
)

func Test_application_routes(t *testing.T) {
	var endpoints = []struct {
		uri    string
		method string
	}{
		{"/", "GET"},
		{"/login", "POST"},
		{"/static/*", "GET"},
	}

	var app application
	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, endpoint := range endpoints {
		// check if the endpoint exists
		if !routeExists(endpoint.uri, endpoint.method, chiRoutes) {
			t.Errorf("endpoint %s %s not found", endpoint.method, endpoint.uri)
		}
	}
}

func routeExists(testUri, testMethod string, routes chi.Routes) bool {
	found := false

	_ = chi.Walk(routes, func(method string, uri string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(uri, testUri) {
			found = true
		}
		return nil
	})

	return found
}
