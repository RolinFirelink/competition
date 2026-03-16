package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

// NewRouter wires routes for public and admin surfaces.
func NewRouter(svc *Service) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// health
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// auth
	r.Post("/auth/login", svc.LoginUser)
	r.Post("/auth/register", svc.RegisterUser)

	// public data
	r.Get("/banners", svc.ListBanners)
	r.Get("/students", svc.PublicStudents)
	r.Get("/students/{id}/routes", svc.ListGrowthRoutes)
	r.Post("/students/{id}/routes", svc.CreateGrowthRoute)
	r.Post("/routes/{rid}/reply", svc.ReplyRoute)
	r.Get("/events", svc.ListEvents)
	r.Get("/events/{id}", svc.GetEvent)

	// admin-only
	r.Route("/admin", func(r chi.Router) {
		r.Post("/login", svc.LoginAdmin)
		r.Group(func(r chi.Router) {
			r.Use(svc.AdminGuard)
			r.Get("/stats", svc.AdminStats)
			r.Get("/students", svc.ListStudents)
			r.Post("/students/{id}/penalty", svc.PenalizeStudent)
		})
	})

	return cors.AllowAll().Handler(r)
}
