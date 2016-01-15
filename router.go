package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	*mux.Router
}

var trac bool

func New() *Router {
	r := Router{mux.NewRouter()}
	return &r
}

func (r *Router) SetTrac(b bool) {
	trac = b
}

func (r *Router) AddFunc(path string, method string, f func(*Context)) *mux.Route {
	return r.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		context := &Context{w, req}
		if trac {
			context.Track()
		}
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
}


