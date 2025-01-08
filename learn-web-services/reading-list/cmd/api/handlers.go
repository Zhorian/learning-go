package main

import (
	"encoding/json"
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

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		app.getBooks(w)
		return
	}

	if r.Method == http.MethodPost {
		fmt.Fprintf(w, "Add a book on the reading list")
		return
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (app *application) getUpdateDeleteCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/v1/books/"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		app.getBook(w, idInt)
		return
	case http.MethodPut:
		app.updateBook(w, idInt)
		return
	case http.MethodDelete:
		app.deleteBook(w, idInt)
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
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

	js, err := json.Marshal(books)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) updateBook(w http.ResponseWriter, id int64) {
	fmt.Fprintf(w, "Update book with id of %d", id)
}

func (app *application) deleteBook(w http.ResponseWriter, id int64) {
	fmt.Fprintf(w, "Delete book with id of %d", id)
}
