package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodPost, "/v1/graphql", app.moviesGraphql)

	router.HandlerFunc(http.MethodPost, "/v1/signin", app.signIn)

	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies/", app.getAllMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:genre_id", app.getAllMoviesByGenre)

	router.HandlerFunc(http.MethodGet, "/v1/genres", app.genreAll)

	router.POST("/v1/admin/editmovie", app.wrap(secure.ThenFunc(app.editmovie)))
	router.DELETE("/v1/admin/deletemovie/:id", app.wrap(secure.ThenFunc(app.deleteMovies)))

	return app.enableCORS(router)
}
