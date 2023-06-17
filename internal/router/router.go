package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type Router struct {
	*chi.Mux
	//bot controller
}

func NewRouter() *Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("public"))).ServeHTTP(w, r)
	})

	return &Router{
		Mux: r,
	}
}

//func (r *Router) BotAPI() {
//	r.Get("/status", r.botController)
//	r.Post("/bot/reload", r.botController)
//	r.Post("/bot/server/update", r.botController)
//}

func (r *Router) Swagger() {
	r.Get("/swagger", swaggerUI)
}
