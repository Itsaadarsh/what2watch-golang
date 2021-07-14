package main

import (
	"backend/models"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Print(errors.New("invalid ID parameter"))
		app.errorJSON(w, err)
		return
	}

	movie := models.Movie{
		ID:          id,
		Title:       "John Wick",
		Description: "Action Movie",
		Year:        2021,
		ReleaseDate: time.Date(2021, 01, 01, 01, 0, 0, 0, time.Local),
		Runtime:     120,
		Rating:      5,
		MPAARating:  "PG-16",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.logger.Print(err)
	}
}
func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request) {}
