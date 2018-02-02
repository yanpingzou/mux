package httputils

import "net/http"
import "golang.org/x/net/context"

// APIFunc is an adapter to allow the use of ordinary functions
// Any function that has the appropriate signature can be registered as an API endpoint.
type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error
