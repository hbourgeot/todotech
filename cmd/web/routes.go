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
	mux.HandleFunc("/products", showProducts)
	mux.HandleFunc("/products/cart", getProductforCart)
	mux.HandleFunc("/neworder", newOrder)
	mux.HandleFunc("/admin", login)
	mux.HandleFunc("/admin/panel/", adminCRUD)
	mux.HandleFunc("/admin/panel/ins", adminCreate)
	mux.HandleFunc("/admin/panel/upda", adminUpdate)
	mux.HandleFunc("/admin/panel/del", adminDelete)
	mux.HandleFunc("/submit", loginAdm)

	return mux
}
