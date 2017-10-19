package web

import (
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
		routeTable.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Route).
			Handler(middlewareChain.Then(route.Handler))
	}

	log.Printf("Listening on %s", addr)
	return http.ListenAndServe(addr, routeTable)
}

var routes = []struct {
	Name    string
	Method  string
	Route   string
	Handler http.Handler
}{
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
	return nil, apiError{code: 400, message: "need credentials"}
}
