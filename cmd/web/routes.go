package main

import (
	"fmt"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}