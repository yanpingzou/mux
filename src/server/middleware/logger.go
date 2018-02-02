package middleware

import (
	"efk/src/server/httputils"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

// LoggerMiddleware record every request.
type LoggerMiddleware struct{}

// LoggerDefaultFormat is the format
// logged used by the default Logger instance.
var LoggerDefaultFormat = "{{.StartTime}} | {{.Status}} | \t {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}} \n"

// LoggerDefaultDateFormat is the
// format used for date by the
// default Logger instance.
var LoggerDefaultDateFormat = "2006-01-02 15:04:05.000"

// NewLoggerMiddleware create a new LoggerMiddleware.
func NewLoggerMiddleware() LoggerMiddleware {
	return LoggerMiddleware{}
}

// WrapHandler returns a new handler function wrapping the previous one in the request chain.
func (l LoggerMiddleware) WrapHandler(handler func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error) func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		start := time.Now()
		err := handler(ctx, w, r, vars)
		end := time.Now()
		remoteIP := r.Header.Get("X-Real-Ip")
		if remoteIP == "" {
			remoteIP = strings.Split(r.RemoteAddr, ":")[0]
		}
		rw := w.(httputils.ResponseWriter)
		log.Infof("%v %5d %12s %15s %6s %s\n", end.Format(LoggerDefaultDateFormat), rw.Status(), time.Since(start), remoteIP, r.Method, r.URL.String())
		return err
	}
}
