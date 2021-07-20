package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/signin", app.signIn)

	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies/", app.getAllMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:genre_id", app.getAllMoviesByGenre)

	router.HandlerFunc(http.MethodGet, "/v1/genres", app.genreAll)

	router.HandlerFunc(http.MethodPost, "/v1/admin/editmovie", app.editmovie)

	router.HandlerFunc(http.MethodDelete, "/v1/admin/deletemovie/:id", app.deleteMovies)

	return app.enableCORS(router)
}
