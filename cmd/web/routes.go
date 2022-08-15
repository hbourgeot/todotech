package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	directory := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", directory))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/not-found", notFound)
	mux.HandleFunc("/admin", login)
	mux.HandleFunc("/admin/panel/", adminCRUD)
	mux.HandleFunc("/submit", loginAdm)

	return mux
}
