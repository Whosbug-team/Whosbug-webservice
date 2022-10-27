package middlewear

import (
	"net/url"
	"strconv"
	"webService_Refactoring/modules"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMonitor() gin.HandlerFunc {
	return func(context *gin.Context) {
		purl, _ := url.Parse(context.Request.RequestURI)

		timer := prometheus.NewTimer(modules.HttpDuration.WithLabelValues(purl.Path))
		defer timer.ObserveDuration()
		modules.TotalRequset.With(prometheus.Labels{
			"method": context.Request.Method,
			"path":   purl.Path,
		}).Inc()
		modules.ResponseStatus.With(prometheus.Labels{
			"status": strconv.Itoa(context.Writer.Status()),
		}).Inc()
		modules.ApiAllRequset.WithLabelValues(
			context.Request.Proto,
			strconv.FormatInt((context.Request.ContentLength), 10),
			context.Request.Host,
		)
		modules.ResponseSize.WithLabelValues(strconv.Itoa(context.Writer.Size()))
	}
}
