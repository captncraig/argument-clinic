package web

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const handlerTag = "handler"

var inFlightGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name: "in_flight_requests",
	Help: "A gauge of requests currently being served by the wrapped handler.",
}, []string{handlerTag})

var reqTimes = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "request_duration",
		Help:    "A histogram of latencies for requests, in milliseconds.",
		Buckets: []float64{1, 2, 3, 4, 5, 10, 15, 25, 50, 100, 200},
	},
	[]string{"handler"},
)

func init() {
	prometheus.MustRegister(inFlightGauge, reqTimes)
}

func requestMetrics(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rn := getRoute(r)
		inflightG := inFlightGauge.WithLabelValues(rn)
		inflightG.Inc()
		defer inflightG.Dec()
		timerH := reqTimes.WithLabelValues(rn)
		start := time.Now()
		defer timerH.Observe(float64(time.Now().Sub(start) / time.Millisecond))
		h.ServeHTTP(w, r)
	})
}
