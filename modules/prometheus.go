package modules

import (
	"github.com/prometheus/client_golang/prometheus"
)

var ResponseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

var TotalRequset = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"method", "path"},
)

var HttpDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	},
	[]string{"path"},
)

var ApiAllRequset = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_all_requset",
	},
	[]string{"proto", "contentLength", "host"},
)

var ResponseSize = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_size",
	},
	[]string{"size"},
)

func init() {
	prometheus.MustRegister(TotalRequset)
	prometheus.MustRegister(ResponseStatus)
	prometheus.MustRegister(ApiAllRequset)
	prometheus.MustRegister(ResponseSize)
	prometheus.MustRegister(HttpDuration)
}
