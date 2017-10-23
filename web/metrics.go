package web

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const handlerTag = "handler"

var inFlightGauges = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "in_flight_requests",
	Help: "Number of requests currently being served by handler.",
}, []string{handlerTag})

var requestTimeOpts = prometheus.HistogramOpts{
	Name:    "request_duration",
	Help:    "A histogram of request times, in milliseconds.",
	Buckets: stdTimeBuckets,
}

var requestSizeOpts = prometheus.HistogramOpts{
	Name:    "request_size",
	Help:    "A histogram of request sizes, in bytes.",
	Buckets: stdSizeBuckets,
}

var responseSizeOpts = prometheus.HistogramOpts{
	Name:    "response_size",
	Help:    "A histogram of response sizes, in bytes.",
	Buckets: stdSizeBuckets,
}

var stdSizeBuckets = []float64{100, 1024, 16 * 1024, 128 * 1024, 1024 * 1024}
var stdTimeBuckets = []float64{1, 5, 10, 20, 50, 100, 500, 1000, 2000}

func init() {
	prometheus.MustRegister(inFlightGauges)
}

func requestMetrics(h http.Handler) http.Handler {
	var name = currentRoute.Name
	inflight := inFlightGauges.WithLabelValues(name)
	cl := prometheus.Labels{"handler": name}

	requestTimeOpts.ConstLabels = cl
	reqTimeVec := prometheus.NewHistogramVec(requestTimeOpts, []string{"code"})
	prometheus.Register(reqTimeVec)

	requestSizeOpts.ConstLabels = cl
	reqSizeVec := prometheus.NewHistogramVec(requestSizeOpts, []string{})
	prometheus.Register(reqSizeVec)

	responseSizeOpts.ConstLabels = cl
	respSizeVec := prometheus.NewHistogramVec(responseSizeOpts, []string{})
	prometheus.Register(respSizeVec)

	// inside out
	h = promhttp.InstrumentHandlerRequestSize(reqSizeVec, h)
	h = promhttp.InstrumentHandlerResponseSize(respSizeVec, h)
	h = promhttp.InstrumentHandlerDuration(reqTimeVec, h)
	h = promhttp.InstrumentHandlerInFlight(inflight, h)

	return h
}
