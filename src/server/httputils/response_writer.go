package httputils

import "net/http"

// ResponseWriter is a wrapper around http.ResponseWriter that provides extra information about
// the response. It is recommended that middleware handlers use this construct to wrap a responsewriter
// if the functionality calls for it.
type ResponseWriter interface {
	http.ResponseWriter
	// Status returns the status code of the response or 0 if the response has
	// not been written
	Status() int
}

// NewResponseWriter creates a ResponseWriter that wraps an http.ResponseWriter
func NewResponseWriter(w http.ResponseWriter) ResponseWriter {
	nrw := &responseWriter{
		ResponseWriter: w,
	}

	return nrw
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(s int) {
	rw.status = s
	rw.ResponseWriter.WriteHeader(s)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	return rw.ResponseWriter.Write(b)
}
