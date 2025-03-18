package routes

import (
	"github.com/JokingLove/market-services/services/rest/service"
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	router  *chi.Mux
	service service.RestService
}

func NewRoutes(r *chi.Mux, service service.RestService) *Routes {
	return &Routes{
		router:  r,
		service: service,
	}
}
