package httputils

import (
	"fmt"
	"net/http"

	"efk/src/errdefs"

	"github.com/sirupsen/logrus"
)

// GetHTTPErrorStatusCode retrieves status code from error message.
func GetHTTPErrorStatusCode(err error) int {
	if err == nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("unexpected HTTP error handling")
		return http.StatusInternalServerError
	}

	var statusCode int

	switch {
	case errdefs.IsNotFound(err):
		statusCode = http.StatusNotFound
	case errdefs.IsInvalidParameter(err):
		statusCode = http.StatusBadRequest
	case errdefs.IsConflict(err):
		statusCode = http.StatusConflict
	case errdefs.IsUnauthorized(err):
		statusCode = http.StatusUnauthorized
	case errdefs.IsUnavailable(err):
		statusCode = http.StatusServiceUnavailable
	case errdefs.IsForbidden(err):
		statusCode = http.StatusForbidden
	case errdefs.IsNotModified(err):
		statusCode = http.StatusNotModified
	case errdefs.IsNotImplemented(err):
		statusCode = http.StatusNotImplemented
	case errdefs.IsSystem(err) || errdefs.IsUnknown(err):
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusInternalServerError
		logrus.WithFields(logrus.Fields{
			"module":     "api",
			"error_type": fmt.Sprintf("%T", err),
		}).Debugf("FIXME: Got an API for which error does not match any expected type!!!: %+v", err)
	}

	return statusCode
}

// MakeErrorHandler makes an HTTP handler that decodes a Docker error and
// returns it in the response.
func MakeErrorHandler(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		MakeErrResponse(w, err)
	}
}

// PageNotFoundError means that requst not found.
type PageNotFoundError struct{}

func (PageNotFoundError) Error() string {
	return "Page not found."
}

// NotFound implement ErrNotFound
func (PageNotFoundError) NotFound() {}

// SystemError system errors.
type SystemError struct {
	Message string
}

func (s SystemError) Error() string {
	return s.Message
}

// ErrSystem implement ErrSystem
func (SystemError) ErrSystem() {}

// ForbiddenError defines Forbidden Error
type ForbiddenError struct{}

func (ForbiddenError) Error() string {
	return "Forbidden"
}

// Forbidden implement ErrForbidden
func (ForbiddenError) Forbidden() {}
