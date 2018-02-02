package middleware

import (
	"efk/src/server/httputils"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// LoggerMiddleware record every request.
type LoggerMiddleware struct{}

// LoggerDefaultDateFormat is the
// format used for date by the
// default Logger instance.
var LoggerDefaultDateFormat = "2006-01-02 15:04:05.000"

// NewLoggerMiddleware create a new LoggerMiddleware.
func NewLoggerMiddleware() LoggerMiddleware {
	return LoggerMiddleware{}
}

// WrapHandler returns a new handler function wrapping the previous one in the request chain.
func (l LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(w, r)
	end := time.Now()

	rw := w.(httputils.ResponseWriter)

	remoteIP := r.Header.Get("X-Real-IP")
	if remoteIP == "" {
		remoteIP = strings.Split(r.RemoteAddr, ":")[0]
	}

	fmt.Printf("%v  %3d %12s %15s %5s %s\n",
		end.Format(LoggerDefaultDateFormat),
		rw.Status(),
		time.Since(start),
		remoteIP,
		r.Method,
		r.URL.String(),
	)
}
