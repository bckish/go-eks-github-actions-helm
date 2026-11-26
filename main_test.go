package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoints(t *testing.T) {
	h := handler("test-1")
	for _, tc := range []struct{ path, key, value string }{
		{"/healthz", "status", "ok"}, {"/readyz", "status", "ready"}, {"/version", "version", "test-1"},
	} {
		t.Run(tc.path, func(t *testing.T) {
			w := httptest.NewRecorder()
			h.ServeHTTP(w, httptest.NewRequest(http.MethodGet, tc.path, nil))
			var result map[string]string
			if w.Code != http.StatusOK || w.Header().Get("Content-Type") != "application/json" || json.Unmarshal(w.Body.Bytes(), &result) != nil || result[tc.key] != tc.value {
				t.Fatalf("unexpected response %d: %s", w.Code, w.Body.String())
			}
		})
	}
	for _, tc := range []struct {
		method, path string
		code         int
	}{{"POST", "/healthz", 405}, {"GET", "/missing", 404}} {
		w := httptest.NewRecorder()
		h.ServeHTTP(w, httptest.NewRequest(tc.method, tc.path, nil))
		if w.Code != tc.code {
			t.Fatalf("%s %s: got %d", tc.method, tc.path, w.Code)
		}
	}
}
