package web

import (
	"context"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/justinas/alice"
	"github.com/oxtoacart/bpool"
)

// slightly modified handler type.
// A function that takes a request and returns some data and/or an error.
// type of returned data matters:
// int - simple status code
// templateToRender - html template and context
// anything else - json serialized to response
type myHandler func(r *http.Request) (interface{}, error)

// a few special request types we can handle specially

// a template with context that we can render
type templateToRender struct {
	tpl *template.Template
	ctx interface{}
}

// an application error
type apiError struct {
	code    int
	message string
}

func (a apiError) Error() string {
	return a.message
}

// the final type that gets sent to the client for all api requests
type apiResponse struct {
	Success bool
	Message string
	Data    interface{}
}

var pool = bpool.NewBufferPool(500)

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := m(r)
	if err != nil {
		if ae, ok := err.(apiError); ok {
			serveJSON(w, r, ae.code, ae.message, nil)
			return
		}
		panic(err)
	}
	// no data -> empty 200
	if data == nil {
		return
	}
	switch x := data.(type) {
	case int:
		w.WriteHeader(x)
	case templateToRender:
		panic("template rendering not implemented yet")
	default:
		serveJSON(w, r, http.StatusOK, "ok", x)
	}
}

func serveJSON(w http.ResponseWriter, r *http.Request, statusCode int, message string, data interface{}) {
	resp := &apiResponse{
		Success: statusCode < 300,
		Message: message,
		Data:    data,
	}
	buf := pool.Get()
	defer pool.Put(buf)
	j := json.NewEncoder(buf)
	if strings.Contains(r.Host, "localhost") {
		j.SetIndent("", "  ")
	}
	if err := j.Encode(resp); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := io.Copy(w, buf); err != nil {
		panic(err)
	}
}

// base middleware chain for everything
var middlewareChain = alice.New(
	// handle panics at top
	panicRecovery,
	// record request counts / times
	requestMetrics,
)

// h builds a final handler from base chain plus additional ones
func h(m myHandler, middlewares ...alice.Constructor) http.Handler {
	if len(middlewares) == 0 {
		return m
	}
	return alice.New(middlewares...).Then(m)
}

func ctx(r *http.Request, key contextKey, value interface{}) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), key, value))
}

func panicRecovery(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				// todo: serve better page
				// todo: log much better
				log.Println(err)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

type contextKey uint16

const (
	ctxRoute contextKey = iota
)

func getRoute(r *http.Request) string {
	route := mux.CurrentRoute(r)
	if route != nil {
		rn := route.GetName()
		if rn != "" {
			return rn
		}
	}
	log.Println("Unknown route name!", r.URL.Path)
	return "unknown"
}
