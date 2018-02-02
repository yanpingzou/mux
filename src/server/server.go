package server

import (
	"efk/src/server/httputils"
	"efk/src/server/middleware"
	"efk/src/server/router"
	"efk/src/server/router/efks"
	"net/http"

	"golang.org/x/net/context"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// versionMatcher defines a variable matcher to be parsed by the router
// when a request is about to be served.
const versionMatcher = "/api/v1"

type ukey string

// UAStringKey is used as key type for user-agent string in net/context struct
const UAStringKey ukey = "upstream-user-agent"

// Server contains instance details for the server
type Server struct {
	addr        string
	m           *mux.Router
	routers     []router.Router
	middlewares []middleware.Middleware
}

// New returns a new instance of the server based on the specified configuration.
// It allocates resources which will be needed for ServeAPI(ports, unix-sockets).
func New() *Server {
	return &Server{}
}

// accept sets a listener the server accepts connections into.
func (s *Server) accept(addr string) {
	s.addr = addr
}

//serve starts listening for inbound requests.
func (s *Server) serve() {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.m,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}

// CreateMux initialize the main routers that the server use.
func (s *Server) createMux() *mux.Router {
	m := mux.NewRouter()
	for _, router := range s.routers {
		for _, route := range router.Routes() {
			f := s.makeHTTPHandler(route.Handler())
			log.Debugf("Registering %s, %s", route.Method(), route.Path())

			m.PathPrefix(versionMatcher).Path(route.Path()).Methods(route.Method()).Handler(f)
		}
	}

	notFoundHandler := httputils.MakeErrorHandler(httputils.PageNotFoundError{})
	m.HandleFunc("/{path:.*}", notFoundHandler)
	m.NotFoundHandler = notFoundHandler

	return m
}

func (s *Server) makeHTTPHandler(handler httputils.APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Define the context that we'll pass around to share info
		ctx := context.WithValue(context.Background(), UAStringKey, r.Header.Get("User-Agent"))
		// Define middleware
		handlerFunc := s.handlerWithGlobalMiddlewares(handler)

		vars := mux.Vars(r) // only url parameters
		if vars == nil {
			vars = make(map[string]string)
		}

		w = httputils.NewResponseWriter(w)

		if err := handlerFunc(ctx, w, r, vars); err != nil {
			statusCode := httputils.GetHTTPErrorStatusCode(err)
			if statusCode >= 500 {
				log.Errorf("Handler for %s %s returned error: %v", r.Method, r.URL.Path, err)
			}
			httputils.MakeErrorHandler(err)(w, r)
		}
	}
}

// useMiddleware appends a new middleware to the request chain.
// This needs to be called before the API routes are configured.
func (s *Server) useMiddleware(m ...middleware.Middleware) {
	s.middlewares = append(s.middlewares, m...)
}

// InitRouter initializes the list of routers for the server.
// This method also enables the Go profiler if enableProfiler is true.
func (s *Server) InitRouter(routers ...router.Router) {
	s.routers = append(s.routers, routers...)
	s.m = s.createMux()
}

// Init create a web server
func Init() {
	s := New()
	s.accept(":8090")
	s.InitRouter(
		efks.NewRouter(),
	)
	s.useMiddleware(
		middleware.NewCORSMiddleware("*"),
		middleware.NewRecoverMiddleware(),
		middleware.NewLoggerMiddleware(),
	)
	s.serve()
}
