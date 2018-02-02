package middleware

import (
	"efk/src/server/httputils"

	"net/http"

	log "github.com/sirupsen/logrus"
)

// RecoverMiddleware revover the error..
type RecoverMiddleware struct {
}

// NewRecoverMiddleware create a new NewRecoverMiddleware.
func NewRecoverMiddleware() RecoverMiddleware {
	return RecoverMiddleware{}
}

// WrapHandler returns a new handler function wrapping the previous one in the request chain.
func (l RecoverMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("panic: %+v", err)
			errResponse := httputils.SystemError{
				Message: "System error, please contact the administrator",
			}
			httputils.MakeErrorHandler(errResponse)(w, r) // 500
		}
	}()

	next(w, r)
}
