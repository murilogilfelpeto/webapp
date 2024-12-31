package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	metricMiddleware := NewPatternMiddleware("webapp")

	// register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(metricMiddleware)
	mux.Use(app.addIpToContext)
	mux.Use(app.Session.LoadAndSave)

	// register routes
	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)

	mux.Handle("/metrics/prometheus", promhttp.Handler())

	// static assets
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
