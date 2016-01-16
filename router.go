package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// @version 1
type Router struct {
	*mux.Router
}

var track bool

func SetTrac(b bool) {
	track = b
}

func New() *Router {
	r := Router{mux.NewRouter().StrictSlash(true)}
	return &r
}

func (r *Router) AddFunc(path string, method string, f func(*Context)) *mux.Route {
	return r.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		context := &Context{w, req}
		f(context)
		if track {
			Logger(context)
		}
	}).Methods(method)
}

/*func (r *Router) Group(prefix string, routes ...*mux.Route) {
	for _, route := range routes {
		r = r.PathPrefix(prefix).Subrouter()
		r.
	}
}*/

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


/**
 * @version 2
 */

/*type Route struct {
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
}*/

func Logger(ctx *Context	) {
	start := time.Now()

	CLog("[SUCC] ========@@ $ @@[ %s ]@@ $ @@========\n", ctx.IP())
	CLog("[TRAC] @@ 方法 @@: # %s #\n", ctx.Method())
	CLog("[TRAC] @@ 地址 @@: # %s #\n", ctx.Uri())
	CLog("[TRAC] @@ 用时 @@: ( %s )\n", time.Since(start))
	println("")
}


