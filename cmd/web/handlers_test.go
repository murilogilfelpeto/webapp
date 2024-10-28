package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_handlers(t *testing.T) {
	tests := []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{name: "Home", url: "/", expectedStatusCode: http.StatusOK},
		{name: "Not Found", url: "/dummy", expectedStatusCode: http.StatusNotFound},
	}

	var app application
	routes := app.routes()

	//create a test server
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	pathToTemplates = "./../../templates/"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ts.Client().Get(ts.URL + tt.url)
			if err != nil {
				t.Fatalf("could not send GET request: %v", err)
			}

			defer closeResponseBody(t, resp)

			if resp.StatusCode != tt.expectedStatusCode {
				t.Errorf("want %d; got %d", tt.expectedStatusCode, resp.StatusCode)
			}
		})
	}
}

func closeResponseBody(t *testing.T, resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		t.Errorf("could not close response body: %v", err)
	}
}
