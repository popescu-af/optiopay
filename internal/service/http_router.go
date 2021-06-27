package service

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/popescu-af/saas-y/pkg/log"
)

// A PathDefinition groups an HTTP method on a path with its handler function.
type PathDefinition struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// Paths represents a collection of path definitions.
type Paths []PathDefinition

// NewRouter creates a new router for the given paths.
func NewRouter(paths Paths) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, p := range paths {
		if p.Method == "WS" {
			router.
				Path(p.Path).
				Handler(apiLogger(p.Handler))
		} else {
			router.
				Methods(p.Method).
				Path(p.Path).
				Handler(apiLogger(p.Handler))
		}
	}
	return router
}

func apiLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.InfoCtx("serving", log.Context{"method": r.Method, "path": r.RequestURI})
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.InfoCtx("served", log.Context{"method": r.Method, "path": r.RequestURI, "duration": time.Since(start).String()})
	})
}
