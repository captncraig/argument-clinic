package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bugsnag/bugsnag-go"
	"github.com/captncraig/argument-clinic/errors"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// slightly modified handler type.
// A function that takes a request and returns some data and/or an error.
// type of returned data matters:
// int - simple status code
// templateToRender - html template and context
// apiError - error message
// anything else - json serialized to response
type myHandler func(r *http.Request) (interface{}, error)

// a few special request types we can handle specially

// a template with context that we can render
type templateToRender struct {
	tpl *template.Template
	ctx interface{}
}

// the final type that gets sent to the client for all api requests
type apiResponse struct {
	Success bool
	Message string
	Data    interface{} `json:",omitempty"`
}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := m(r)
	if err != nil {
		if _, ok := err.(*errors.ApiError); !ok {
			log.Printf("PPPP %T", err)
			panic(err)
		} else {
			data = err
		}
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
	case *errors.ApiError:
		serveError(w, r, x)
	default:
		serveJSON(w, r, x)
	}
}

func serveError(w http.ResponseWriter, r *http.Request, err *errors.ApiError) {
	resp := &apiResponse{
		Success: false,
		Message: err.Message,
	}
	bugsnagLog(err, r)
	resp.serve(w, r, err.Code)
}

func serveJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	resp := &apiResponse{
		Success: true,
		Message: "ok",
		Data:    data,
	}
	fmt.Println("BBBBB", resp)
	resp.serve(w, r, http.StatusOK)
}

func (ar *apiResponse) serve(w http.ResponseWriter, r *http.Request, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	j := json.NewEncoder(w)
	if strings.Contains(r.Host, "localhost") {
		j.SetIndent("", "  ")
	}
	if err := j.Encode(ar); err != nil {
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

func getRoute(r *http.Request) string {
	route := mux.CurrentRoute(r)
	if route != nil {
		rn := route.GetName()
		if rn != "" {
			return rn
		}
	}
	// TODO: log this to bugsnag
	return "unknown"
}

func init() {
	if bsnagKey := os.Getenv("BUGSNAG_KEY"); bsnagKey != "" {
		bugsnag.Configure(bugsnag.Configuration{
			APIKey: bsnagKey,
			//Logger:       nil,
			PanicHandler: func() {},
		})
		middlewareChain = middlewareChain.Append(bugsnagMiddleware)
		bugsnag.OnBeforeNotify(fillEventData)
		bugsnagLog = func(err *errors.ApiError, r *http.Request) {
			bugsnag.Notify(err, r)
		}
		bugsnagNotifyDeploy()
	}
}

// this just needs to catch unhandled panics. Should be none, so these are the most severe.
func bugsnagMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement this!
		h.ServeHTTP(w, r)
	})
}

func fillEventData(ev *bugsnag.Event, cfg *bugsnag.Configuration) error {
	apiErr, ok := ev.Error.Err.(*errors.ApiError)
	if !ok {
		return nil
	}
	ev.ErrorClass = apiErr.ErrType
	ev.Message = apiErr.Orig.Error()
	return nil
}

var bugsnagLog = func(*errors.ApiError, *http.Request) {}

func bugsnagNotifyDeploy() {
	data := map[string]string{
		"apiKey":     os.Getenv("BUGSNAG_KEY"),
		"repository": "https://github.com/captncraig/argument-clinic",
		"revision":   os.Getenv("APP_VERSION"),
		"appVersion": "1.2.3",
		"branch":     "go",
	}
	dat, _ := json.Marshal(data)
	fmt.Println(string(dat))
	resp, err := http.Post("https://notify.bugsnag.com/deploy", "application/json", bytes.NewReader(dat))
	fmt.Println(err, resp.Status)
}
