package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// (Rest API) /metrics → Prometheus (scrape) → Grafana (consulta e visualiza)
var HttpReq = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http-req-total",
	},
	[]string{"path", "method"},
)

func SetProme(server *gin.Engine) {
	prometheus.MustRegister(HttpReq)

	server.Use(func(ctx *gin.Context) {
		path := ctx.FullPath()
		method := ctx.Request.Method
		ctx.Next()
		HttpReq.WithLabelValues(path, method).Inc()
	})

	server.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
