package main

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Hello from rest test!")
	} else {
		fmt.Fprint(w, "Unknown Method")
	}
}

func main() {
	http.HandleFunc("/", handleRoot)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
