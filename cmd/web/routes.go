package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {

	mux := chi.NewRouter()

	// mux.Use(middleware.RequestID)
	// mux.Use(middleware.RealIP)
	// mux.Use(middleware.Recoverer)

	// if a.debug {
	// 	mux.Use(middleware.Logger)
	// }

	// mux.Get("/comments", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Comments"))
	// })

	test := func(w http.ResponseWriter, r *http.Request ){
		path := r.URL.Path
		statement:= " ->> this path is yet to be writen"
		w.Write([]byte(path + statement))
	}

// routes
	mux.Get("/", test)	
mux.Get("/comments/{post}", test)
mux.Post("/comments/{post}", test)

mux.Get("/login",test)
mux.Post("/login",test)

mux.Get("/signup",test)
mux.Post("/signup",test)

// auth required routes
mux.Get("/logout",test)

mux.Get("/vote",test)
mux.Get("/submit",test)
mux.Post("/submit",test)


	return mux
}
