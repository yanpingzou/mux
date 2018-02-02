package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// CORSMiddleware injects CORS headers to each request
// when it's configured.
type CORSMiddleware struct {
	defaultHeaders string
}

// NewCORSMiddleware creates a new CORSMiddleware with default headers.
func NewCORSMiddleware(d string) CORSMiddleware {
	return CORSMiddleware{defaultHeaders: d}
}

// WrapHandler returns a new handler function wrapping the previous one in the request chain.
func (c CORSMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// If "api-cors-header" is not given, but "api-enable-cors" is true, we set cors to "*"
	// otherwise, all head values will be passed to HTTP handler
	corsHeaders := c.defaultHeaders
	if corsHeaders == "" {
		corsHeaders = "*"
	}

	logrus.Debugf("CORS header is enabled and set to: %s", corsHeaders)
	w.Header().Add("Access-Control-Allow-Origin", corsHeaders)
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Registry-Auth")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD, POST, DELETE, PUT, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	next(w, r)
}
