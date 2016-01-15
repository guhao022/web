package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

/*
// @version 1
type Router struct {
	*mux.Router
}

var trac bool

func New() *Router {
	r := Router{mux.NewRouter()}
	return &r
}

func (r *Router) AddFunc(path string, method string, f func(*Context)) *mux.Route {
	return r.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		context := &Context{w, req}
		f(context)
	}).Methods(method)
}

func (r *Router) Group(prefix string) *Router {
	s := r.PathPrefix(prefix).Subrouter()
	return &Router{s}
}

func (r *Router) Get(path string, f func(*Context)) *mux.Route {
	return r.AddFunc(path, "GET", f)
}

func (r *Router) Post(path string, f func(*Context)) *mux.Route {
	return r.AddFunc(path, "POST", f)
}

func (r *Router) Delete(path string, f func(*Context)) *mux.Route {
	return r.AddFunc(path, "DELETE", f)
}

func (r *Router) Put(path string, f func(*Context)) *mux.Route {
	return r.AddFunc(path, "PUT", f)
}*/


/**
 * @version 2
 */

type Router struct {
	*mux.Router
}

func New() *Router {
	r := Router{mux.NewRouter()}
	return &r
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Func 		func(*Context)
	handlerFunc http.HandlerFunc
}

type Routes []Route

func (r *Router) Register(routes Routes) *mux.Route {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		var ctx *Context
		handler = route.handlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx = &Context{w, req}
		})

		handler = trac(handler, route.Name)

		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)
	}

	return router
}

func trac(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		CLog(
			"%s\t%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}


