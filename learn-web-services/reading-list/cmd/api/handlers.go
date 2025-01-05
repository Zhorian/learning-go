package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "environment: %s\n", version)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Display a list of books on the reading list")
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

func (app *application) getBook(w http.ResponseWriter, id int64) {
	fmt.Fprintf(w, "Display book with id of %d", id)
}

func (app *application) updateBook(w http.ResponseWriter, id int64) {
	fmt.Fprintf(w, "Update book with id of %d", id)
}

func (app *application) deleteBook(w http.ResponseWriter, id int64) {
	fmt.Fprintf(w, "Delete book with id of %d", id)
}
