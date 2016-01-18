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
	return &Router{}
}

func (r *Router) AddFunc(path, method string, h http.HandlerFunc) *mux.Route {
	var handler http.Handler
	handler = h
	if track {
		handler = Logger(handler)
	}
	return r.NewRoute().PathPrefix(path).Handler(handler).Methods(method)
}

func (r *Router) Get(path string, h http.HandlerFunc) *mux.Route {
	return r.AddFunc(path, "GET", h)
}

func (r *Router) Post(path string, h http.HandlerFunc) *mux.Route {
	return r.AddFunc(path, "POST", h)
}

func (r *Router) Delete(path string, h http.HandlerFunc) *mux.Route {
	return r.AddFunc(path, "DELETE", h)
}

func (r *Router) Put(path string, h http.HandlerFunc) *mux.Route {
	return r.AddFunc(path, "PUT", h)
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

func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		CLog("[SUCC] ========@@ $ @@[ %s ]@@ $ @@========\n", IP(r))
		CLog("[TRAC] @@ 方法 @@: # %s #\n", r.Method)
		CLog("[TRAC] @@ 地址 @@: # %s #\n", r.RequestURI)
		CLog("[TRAC] @@ 用时 @@: ( %s )\n", time.Since(start))
		println("")
	})

}

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
