package routes

import (
	"net/http"

	//"github.com/coderminer/restful/auth"
	"os"

	"github.com/gorilla/mux"
	"github.com/serviceComputing1/server/auth"
	"github.com/serviceComputing1/server/service"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("GET", "/people/", service.GetPeople, auth.TokenMiddleware)
	register("GET", "/people/{id}", service.GetPerson, auth.TokenMiddleware)
	register("GET", "/api", service.GetAllApi, nil)
	register("GET", "/", service.GetIndex, nil)
}

func NewRouter() *mux.Router {

	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	router := mux.NewRouter()

	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/dist")))
	//router.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./dist/css/")))

	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
