package routes

import (
	"docker-collector/pkg/controllers"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var reg = controllers.GetRegistry()

var RegisterBasicRoute = func(router *mux.Router) {
	router.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
}
