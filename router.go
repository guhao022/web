package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
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

func Logger(ctx *Context) {
	start := time.Now()

	CLog("[SUCC] ========@@ $ @@[ %s ]@@ $ @@========\n", ctx.IP())
	CLog("[TRAC] @@ 方法 @@: # %s #\n", ctx.Method())
	CLog("[TRAC] @@ 地址 @@: # %s #\n", ctx.Uri())
	CLog("[TRAC] @@ 用时 @@: ( %s )\n", time.Since(start))
	println("")
}


/**
 * @version 2
 */

/*type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

var track bool

func SetTrac(b bool) {
	track = b
}

func Register(routes []*Route) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandleFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}*/

/*func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		CLog("[SUCC] ========@@ $ @@[ %s ]@@ $ @@========\n", IP(r))
		CLog("[TRAC] @@ 方法 @@: # %s #\n", r.Method)
		CLog("[TRAC] @@ 路由 @@: # %s #\n", name)
		CLog("[TRAC] @@ 地址 @@: # %s #\n", r.RequestURI)
		CLog("[TRAC] @@ 用时 @@: ( %s )\n", time.Since(start))
		println("")
	})

}*/

func IP(r *http.Request) string {
	ipstr := r.Header.Get("X-Forwarded-For")
	ips := strings.Split(ipstr, ",")
	if len(ips) > 0 && ips[0] != "" {
		rip := strings.Split(ips[0], ":")
		return rip[0]
	}
	ip := strings.Split(r.RemoteAddr, ":")
	if len(ip) > 0 {
		if ip[0] != "[" {
			return ip[0]
		}
	}
	return "127.0.0.1"
}
