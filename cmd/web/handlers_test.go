package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestAppHome(t *testing.T) {
	testCases := []struct {
		name         string
		putInSession string
		expectedHTML string
	}{
		{
			name:         "First Visit - No session",
			putInSession: "",
			expectedHTML: "<small>From Session: </small>",
		},
		{
			name:         "Second Visit - With session",
			putInSession: "test",
			expectedHTML: "<small>From Session: test</small>",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req = addContextAndSessionToRequest(req, app)
			_ = app.Session.Destroy(req.Context())

			if len(test.putInSession) > 0 {
				app.Session.Put(req.Context(), "test", test.putInSession)
			}

			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(app.Home)
			handler.ServeHTTP(rec, req)

			if !strings.Contains(rec.Body.String(), test.expectedHTML) {
				t.Errorf("want %s; got %s", test.expectedHTML, rec.Body.String())
			}
		})
	}
}

func closeResponseBody(t *testing.T, resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		t.Errorf("could not close response body: %v", err)
	}
}

func getContext(req *http.Request) context.Context {
	ctx := context.WithValue(req.Context(), contextUserKey, "unknown")
	return ctx
}

func addContextAndSessionToRequest(req *http.Request, app application) *http.Request {
	req = req.WithContext(getContext(req))

	ctx, _ := app.Session.Load(req.Context(), req.Header.Get("X-Session"))

	return req.WithContext(ctx)
}
