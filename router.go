package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	*mux.Router
}

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

func (r *Router) Group(prefix string, routes ...*mux.Route) {
	if len(routes) > 0 {
		for _, mr := range routes {
			mr.PathPrefix(prefix).Subrouter()
		}
	}
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


