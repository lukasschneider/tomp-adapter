package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/operator", func(r chi.Router) {
		r.Get("/ping", s.HelloWorldHandler)
		r.Get("/meta", s.HelloWorldHandler)
		r.Get("/stations", s.HelloWorldHandler)
		r.Get("/available-assets", s.HelloWorldHandler)
		r.Get("/alerts", s.HelloWorldHandler)
		r.Get("/operating-calender", s.HelloWorldHandler)
		r.Get("/operating-hours", s.HelloWorldHandler)
		r.Get("/information", s.HelloWorldHandler)
		r.Get("/pricing-plans", s.HelloWorldHandler)
		r.Get("/regions", s.HelloWorldHandler)
	})

	r.Route("/planning", func(r chi.Router) {
		r.Post("/inquiries", s.HelloWorldHandler)
		r.Post("/offers", s.HelloWorldHandler)
	})

	r.Route("/bookings", func(r chi.Router) {
		r.Post("/", s.HelloWorldHandler)
		r.Post("/one-stop", s.HelloWorldHandler)
		r.Post("/{id}/events", s.HelloWorldHandler)
		r.Get("/{id}", s.HelloWorldHandler)
		r.Get("/", s.HelloWorldHandler)
		r.Put("/{id}", s.HelloWorldHandler)
		r.Delete("/{id}/subscriptions", s.HelloWorldHandler)
		r.Post("/{id}/subscriptions", s.HelloWorldHandler)
		r.Post("/{id}/notifications", s.HelloWorldHandler)
		r.Get("/{id}/notifications", s.HelloWorldHandler)
	})

	r.Route("/legs", func(r chi.Router) {
		r.Get("/{id}/available-assets", s.HelloWorldHandler)
		r.Get("/{id}", s.HelloWorldHandler)
		r.Put("/{id}", s.HelloWorldHandler)
		r.Post("/{id}/ancillaries/{category}/{number}", s.HelloWorldHandler)
		r.Delete("/{id}/ancillaries/{category}/{number}", s.HelloWorldHandler)
		r.Post("/{id}/events", s.HelloWorldHandler)
		r.Get("/{id}/progress", s.HelloWorldHandler)
		r.Post("/{id}/progress", s.HelloWorldHandler)
		r.Post("/{id}/confirmation", s.HelloWorldHandler)
	})

	r.Route("/payment", func(r chi.Router) {
		r.Get("/journal-entry", s.HelloWorldHandler)
		r.Get("/{id}/claim-extra-costs", s.HelloWorldHandler)
	})

	r.Route("/support", func(r chi.Router) {
		r.Get("/", s.HelloWorldHandler)
		r.Get("/{id}/status", s.HelloWorldHandler)
	})

	r.Route("/customers", func(r chi.Router) {
		r.Post("/", s.HelloWorldHandler)
		r.Patch("/{id}", s.HelloWorldHandler)
		r.Get("/{id}", s.HelloWorldHandler)
		r.Delete("/{id}", s.HelloWorldHandler)
	})

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
