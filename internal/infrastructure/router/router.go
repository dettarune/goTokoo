package router

import (
	"net/http"

	"github.com/dettarune/goTokoo/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func NewRouter(appConfig *config.BootstrapConfig) *chi.Mux {
	r := chi.NewRouter()
	
	r.Use(middleware.RequestID)	
	r.Use(middleware.Recoverer)		
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}) 
			

	return r
}