package mux

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
	return r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		context := &Context{w, r}
		f(context)
	}).Methods(method)
}

func (r *Router) Prefix(prefix string, route ...mux.Route) {
	for _, mr := range route {
		mr.PathPrefix(prefix)
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


