package router

import (
	"github.com/dettarune/goTokoo/internal/delivery/http"
	"github.com/go-chi/chi/v5"
)

type RouteConfig struct {
	App               *chi.Mux
	UserController    *http.UserController
}

func (c *RouteConfig) SetupRoute() {
	c.UserRoute()
}

func (c *RouteConfig) UserRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	// c.App.Post("/api/users/_login", c.UserController.Login)
}
