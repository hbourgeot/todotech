package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("./ui/index.go.html"))
	err := temp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {

}
