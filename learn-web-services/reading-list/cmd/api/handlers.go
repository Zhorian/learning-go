package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

	err := app.writeJson(w, http.StatusOK, envelope{"healthcheck": data}, nil)

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
		app.createBook(w, r)
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
	books, err := app.models.Books.GetAll()
	if err != nil {
		app.writeInternalServerError(w)
		return
	}

	err = app.writeJson(w, http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		app.writeInternalServerError(w)
	}
}

func (app *application) getBook(w http.ResponseWriter, id int64) {
	book, err := app.models.Books.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	err = app.writeJson(w, http.StatusOK, envelope{"book": book}, nil)
	if err == nil {
		return
	}

	app.writeInternalServerError(w)
}

func (app *application) createBook(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title     string   `json:"title"`
		Published int      `json:"published"`
		Pages     int      `json:"pages"`
		Genres    []string `json:"genres"`
		Rating    float32  `json:"rating"`
	}

	err := app.readJson(w, r, &input)

	if err != nil {
		app.writeBadRequest(w)
		return
	}

	book := &data.Book{
		Title:    input.Title,
		Pubished: input.Published,
		Pages:    input.Pages,
		Genres:   input.Genres,
		Rating:   input.Rating,
	}

	err = app.models.Books.Insert(book)
	if err != nil {
		app.writeInternalServerError(w)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/books/%d", book.ID))

	err = app.writeJson(w, http.StatusCreated, envelope{"book": book}, headers)
	if err != nil {
		app.writeInternalServerError(w)
		return
	}
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

	book, err := app.models.Books.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
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

	err = app.models.Books.Update(book)
	if err != nil {
		app.writeInternalServerError(w)
		return
	}

	err = app.writeJson(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.writeInternalServerError(w)
		return
	}
}

func (app *application) deleteBook(w http.ResponseWriter, id int64) {
	err := app.models.Books.Delete(id)

	if err != nil {
		switch {
		case errors.Is(err, errors.New("record not found")):
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
