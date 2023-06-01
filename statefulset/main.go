package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

var (
	HTTPRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "statefulset_http_requests_total",
			Help: "Number of the http requests received since the server started",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(HTTPRequests)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		code := 200
		switch path {
		case "/test":
			w.WriteHeader(200)
			w.Write([]byte("OK"))
		case "/metrics":
			promhttp.Handler().ServeHTTP(w, r)
		default:
			w.WriteHeader(404)
			w.Write([]byte("Not Found"))
		}
		HTTPRequests.WithLabelValues(strconv.Itoa(code)).Inc()
	})
	http.ListenAndServe(":8081", nil)
}
