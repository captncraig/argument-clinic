package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/captncraig/argument-clinic/data"
)

// Listen runs the web server. It will not return anything until the http server terminates with an error
func Listen(addr string, db data.DataAccess) error {
	routeTable := mux.NewRouter()

	for _, route := range routes {
		currentRoute = route
		routeTable.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Route).
			Handler(middlewareChain.Then(route.Handler))
	}

	log.Printf("Listening on %s", addr)
	return http.ListenAndServe(addr, routeTable)
}

type routeDef struct {
	Name    string
	Method  string
	Route   string
	Handler http.Handler
}

// middleware constructors can read and save this at creation time for additional context
var currentRoute routeDef

var routes = []routeDef{
	{
		Name:    "CreateComment",
		Method:  http.MethodGet,
		Route:   "/api/comments",
		Handler: h(createComment),
	},
	{
		Name:    "ScrapeMetrics",
		Method:  http.MethodGet,
		Route:   "/metrics",
		Handler: promhttp.Handler(),
	},
}

func createComment(r *http.Request) (interface{}, error) {
	return nil, fmt.Errorf("AAAAAA")
}
