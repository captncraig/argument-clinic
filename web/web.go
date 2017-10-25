package web

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/captncraig/argument-clinic/data"
	"github.com/captncraig/argument-clinic/errors"
	"github.com/captncraig/argument-clinic/models"
)

type server struct {
	db data.DataAccess
}

// Listen runs the web server. It will not return anything until the http server terminates with an error
func Listen(addr string, db data.DataAccess) error {
	routeTable := mux.NewRouter()
	srv := &server{db: db}
	for _, route := range srv.routes() {
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

func (s *server) routes() []routeDef {
	return []routeDef{
		{
			Name:    "CreateComment",
			Method:  http.MethodPost,
			Route:   "/api/comments",
			Handler: h(s.createComment),
		},
		{
			Name:    "ScrapeMetrics",
			Method:  http.MethodGet,
			Route:   "/metrics",
			Handler: promhttp.Handler(),
		},
	}
}

func (s *server) createComment(r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)
	req := &models.CreateCommentRequest{}
	if err := dec.Decode(req); err != nil {
		return nil, errors.ErrDecodingJSON(err)
	}
	if req.Name == "" {
		req.Name = "Anonymous"
	}
	id, err := s.db.CreateComment(r.Context(), req)
	if err != nil {
		return nil, err
	}
	return &models.Comment{
		ID:   id,
		Name: req.Name,
		Text: req.Text,
		Date: time.Now().UTC(),
	}, nil
}
