package jwtvalidator

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// Demo a Demo plugin.
type JwtValidator struct {
	next     http.Handler
	headers  map[string]string
	name     string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// if len(config.Headers) == 0 {
	// 	return nil, fmt.Errorf("headers cannot be empty")
	// }

	return &JwtValidator{
		// headers:  config.Headers,
		next:     next,
		name:     name,
	}, nil
}

func (a *JwtValidator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	req.Header.Set("X-DIMA-TEST-HEADER", "HELLO WORLD")

	a.next.ServeHTTP(rw, req)
}