package router

import (
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"printer-api/models"
	"printer-api/middleware"
	"printer-api/api"
)

func InitRouter(config models.Configuration) http.Handler {
	r := mux.NewRouter()
	router := r.PathPrefix(config.Server.APIPathPrefix).Subrouter()

	cors := handlers.CORS(
		handlers.AllowedOrigins(strings.Split(config.Server.AllowedOrigins, ",")),
		handlers.AllowedMethods(strings.Split(config.Server.AllowedMethods, ",")),
		handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"}),
	)

	middleware.InitLogger(config)
	RegisterAPIRoutes(router, config)
	return cors(router)
}

func RegisterAPIRoutes(router *mux.Router, config models.Configuration) {
	api.RegisterAggregatesRoutes(router, config)
	api.RegisterStatsRoutes(router, config)
}
