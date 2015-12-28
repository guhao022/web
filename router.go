package mux

import (
	ctx "github.com/num5/context"
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	*mux.Router
	group string
}

func New() *Router {
	r := Router{mux.NewRouter(),""}
	return &r
}

func (r *Router) AddFunc(path string, method string, f func(*ctx.Context)) *mux.Route {
	var s *mux.Router
	if len(r.group) > 0 {
		s = r.PathPrefix(r.group).Subrouter()
	}
	s = r.Router

	return s.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		context := ctx.Context{w, req}
		f(&context)
	}).Methods(method)
}

func (r *Router) Group(pattern string, mroute ...*mux.Route) error {
	r.group = pattern
	for _, mr := range mroute {
		err := mr.GetError()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Router) Get(path string, f func(*ctx.Context)) *mux.Route {
	return r.AddFunc(path, "GET", f)
}

func (r *Router) Post(path string, f func(*ctx.Context)) *mux.Route {
	return r.AddFunc(path, "POST", f)
}

func (r *Router) Delete(path string, f func(*ctx.Context)) *mux.Route {
	return r.AddFunc(path, "DELETE", f)
}

func (r *Router) Put(path string, f func(*ctx.Context)) *mux.Route {
	return r.AddFunc(path, "PUT", f)
}


