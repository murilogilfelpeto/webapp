package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_addIpToContext(t *testing.T) {
	tests := []struct {
		name         string
		headerName   string
		headerValue  string
		address      string
		emptyAddress bool
	}{
		{"All empty", "", "", "", false},
		{"Empty Address", "", "", "", true},
		{"Forward Header", "X-Forwarded-For", "192.3.2.1", "", false},
		{"Invalid Address", "", "", "hello:world", false},
	}

	var app application

	// create a dummy handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// make sure the values exist in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error("Expected value in context, got nil")
		}

		// make sure it is a string
		ip, ok := val.(string)
		if !ok {
			t.Error("Expected string in context, got something else")
		}
		t.Log(ip)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlerToTest := app.addIpToContext(nextHandler)
			req := httptest.NewRequest(http.MethodGet, "/", nil)

			if tt.emptyAddress {
				req.RemoteAddr = ""
			}

			if len(tt.headerName) > 0 {
				req.Header.Add(tt.headerName, tt.headerValue)
			}

			if len(tt.address) > 0 {
				req.RemoteAddr = tt.address
			}

			handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
		})
	}
}

func Test_application_ipFromContext(t *testing.T) {
	var app application
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextUserKey, "test")

	val := app.ipFromContext(ctx)

	if val != "test" {
		t.Errorf("Expected test, got %s", val)
	}
}
