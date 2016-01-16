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

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandleFunc 	func(*Context)
}

var track bool

func SetTrac(b bool) {
	track = b
}

func Register(routes []*Route) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		router.
		Methods(route.Method).
		Name(route.Name).
		HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx := &Context{w, req}
			route.HandleFunc(ctx)
			if track {
				Logger(ctx, route.Name)
			}
		})


	}

	return router
}

func Logger(ctx *Context, name string) {
		start := time.Now()

		CLog(
			"< %s >\t# %s #\t< %s >\t[ %s ]\n",
			ctx.Method(),
			ctx.Uri(),
			name,
			time.Since(start),
		)
}


