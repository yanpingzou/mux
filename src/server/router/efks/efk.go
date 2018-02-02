package efks

import (
	"efk/src/server/httputils"
	"efk/src/server/router"
	"fmt"
	"net/http"

	"golang.org/x/net/context"
)

// NewRouter initializes a new efk router.
func NewRouter() router.Router {
	r := &efkRouter{}
	r.initRoutes()
	return r
}

// efkRouter is a router to talk with the efk controller.
type efkRouter struct {
	routes []router.Route
}

// Routes returns the available routes to the efk controller
func (r *efkRouter) Routes() []router.Route {
	return r.routes
}

func (r *efkRouter) initRoutes() {
	r.routes = []router.Route{
		router.NewGetRoute("/hello/{name}", hello),
		router.NewGetRoute("/hero", hero),
		router.NewGetRoute("/herr", herr),
	}
}

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	responseData := fmt.Sprintf("hello, %s", vars["name"])
	return httputils.MakeResponse(w, http.StatusOK, responseData)
}

func hero(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	a := []string{"a", "b"}
	fmt.Println(a[3]) // panic
	return httputils.MakeResponse(w, http.StatusOK, a)
}

func herr(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	err := httputils.ForbiddenError{}
	return httputils.MakeErrResponse(w, err)
}
