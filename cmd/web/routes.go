package main

import (
	"net/http"

	"github.com/gdalelio/bookings/pkg/config"
	"github.com/gdalelio/bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	/*  mux := pat.New()

	    mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	    mux.Get("/about", http.HandlerFunc(handlers.Repo.About)) */

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//write to consle when page is hit
	//mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/check-availabilty", handlers.Repo.Availability)
	mux.Get("/reservations", handlers.Repo.Reservation)
	mux.Get("/contact", handlers.Repo.Contact)

	


	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
