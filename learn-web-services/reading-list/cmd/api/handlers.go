package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"reading_list/internal/data"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	err := app.writeJson(w, http.StatusOK, envelope{"healthcheck": data})

	if err == nil {
		return
	}

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		app.getBooks(w)
		return
	}

	if r.Method == http.MethodPost {
		var input struct {
			Title     string   `json:"title"`
			Published int      `json:"published"`
			Pages     int      `json:"pages"`
			Genres    []string `json:"genres"`
			Rating    float64  `json:"rating"`
		}

		err := app.readJson(w, r, &input)

		if err != nil {
			app.writeBadRequest(w)
			return
		}

		fmt.Fprintf(w, "%v\n", input)
		return
	}

	app.writeMethodNotAllowed(w)
}

func (app *application) getUpdateDeleteCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		app.writeBadRequest(w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		app.getBook(w, idInt)
		return
	case http.MethodPut:
		app.updateBook(w, idInt, r)
		return
	case http.MethodDelete:
		app.deleteBook(w, idInt)
		return
	default:
		app.writeMethodNotAllowed(w)
	}
}

func (app *application) getBooks(w http.ResponseWriter) {
	books := []data.Book{
		{
			ID:        1,
			CreatedAt: time.Now(),
			Title:     "Gareth - How I learned go!",
			Pubished:  2024,
			Pages:     50,
			Genres:    []string{"Biography", "Factual", "Self Help"},
			Rating:    4.4,
			Version:   1,
		},

		{
			ID:        2,
			CreatedAt: time.Now(),
			Title:     "Battletech - The First Succession War",
			Pubished:  2020,
			Pages:     500,
			Genres:    []string{"Fiction", "Sci-Fi"},
			Rating:    4.4,
			Version:   1,
		},
	}

	err := app.writeJson(w, http.StatusOK, envelope{"books": books})
	if err == nil {
		return
	}

	app.writeInternalServerError(w)
}

func (app *application) getBook(w http.ResponseWriter, id int64) {
	book := data.Book{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Gareth - How I learned go!",
		Pubished:  2024,
		Pages:     50,
		Genres:    []string{"Biography", "Factual", "Self Help"},
		Rating:    4.4,
		Version:   1,
	}

	err := app.writeJson(w, http.StatusOK, envelope{"book": book})
	if err != nil {
		return
	}

	app.writeInternalServerError(w)
}

func (app *application) updateBook(w http.ResponseWriter, id int64, r *http.Request) {
	var input struct {
		Title     *string  `json:"title"`
		Published *int     `json:"published"`
		Pages     *int     `json:"pages"`
		Genres    []string `json:"genres"`
		Rating    *float32 `json:"rating"`
	}
	err := app.readJson(w, r, &input)
	if err != nil {
		app.writeBadRequest(w)
		return
	}

	book := data.Book{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Some book to update!",
		Pubished:  2024,
		Pages:     50,
		Genres:    []string{"Biography", "Factual", "Self Help"},
		Rating:    4.4,
		Version:   1,
	}

	if input.Title != nil {
		book.Title = *input.Title
	}

	if input.Published != nil {
		book.Pubished = *input.Published
	}

	if input.Pages != nil {
		book.Pages = *input.Pages
	}

	if len(input.Genres) > 0 {
		book.Genres = input.Genres
	}

	if input.Rating != nil {
		book.Rating = *input.Rating
	}

	fmt.Fprintf(w, "%v\n", book)
}

func (app *application) deleteBook(w http.ResponseWriter, id int64) {
	fmt.Fprintf(w, "Delete book with id of %d", id)
}
