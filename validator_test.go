package jwtvalidator_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Dimoniq/jwtvalidator"
)

func TestDemo(t *testing.T) {
	cfg := jwtvalidator.CreateConfig()
	// cfg.Headers["X-Host"] = "[[.Host]]"
	// cfg.Headers["X-Method"] = "[[.Method]]"
	// cfg.Headers["X-URL"] = "[[.URL]]"
	// cfg.Headers["X-URL"] = "[[.URL]]"
	// cfg.Headers["X-Demo"] = "test"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := jwtvalidator.New(ctx, next, cfg, "jwtvalidator")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

	assertHeader(t, req, "X-DIMA-TEST-HEADER", "HELLO WORLD")
}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
