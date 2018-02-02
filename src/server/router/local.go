package router

import (
	"efk/src/server/httputils"
)

// localRoute defines an individual API route to connect
// with the server. It implements Route.
type localRoute struct {
	method  string
	path    string
	handler httputils.APIFunc
}

// Path returns the subpath where the route responds to.
func (l localRoute) Path() string {
	return l.path
}

// Method returns the http method that the route responds to.
func (l localRoute) Method() string {
	return l.method
}

// Handler returns the APIFunc to let the server wrap it in middlewares.
func (l localRoute) Handler() httputils.APIFunc {
	return l.handler
}

// NewRoute initializes a new local route for the router.
func NewRoute(method, path string, handler httputils.APIFunc) Route {
	r := localRoute{method, path, handler}

	return r
}

// NewGetRoute initializes a new route with the http method GET.
func NewGetRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("GET", path, handler)
}

// NewPostRoute initializes a new route with the http method POST.
func NewPostRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("POST", path, handler)
}

// NewPutRoute initializes a new route with the http method PUT.
func NewPutRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("PUT", path, handler)
}

// NewDeleteRoute initializes a new route with the http method DELETE.
func NewDeleteRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("DELETE", path, handler)
}

// NewOptionsRoute initializes a new route with the http method OPTIONS.
func NewOptionsRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("OPTIONS", path, handler)
}

// NewHeadRoute initializes a new route with the http method HEAD.
func NewHeadRoute(path string, handler httputils.APIFunc) Route {
	return NewRoute("HEAD", path, handler)
}
