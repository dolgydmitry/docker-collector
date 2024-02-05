package server

import (
	"docker-collector/pkg/routes"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RunServer(host string) {
	log.Info().Msgf("Let's run this server in host: %v\n", host)
	router := mux.NewRouter()
	routes.RegisterBasicRoute(router)
	http.ListenAndServe(host, router)
}

func RunServerApp(serverAddress string) {
	log.Info().Msgf("Let's run this server on the host: %v\n", serverAddress)
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatal().Msgf("server cannot run: %s", err)
	}
}
