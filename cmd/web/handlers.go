package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/not-found", http.StatusSeeOther)
		return
	}

	temp := template.Must(template.ParseFiles("./ui/index.go.html"))
	err := temp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("./ui/404.html"))
	err := temp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("./ui/login.html"))
	err := temp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func loginAdm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	user, pass := r.PostForm.Get("admin-user"), r.PostForm.Get("admin-pass")
	loginErrors := make(map[string]string)

	if strings.TrimSpace(user) == "" {
		loginErrors["username"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(user) > 100 {
		loginErrors["username"] = "This field cannot be more than 100 characters long"
	}

	if strings.TrimSpace(pass) == "" {
		loginErrors["password"] = "This field cannot be blank"
	}

	if len(loginErrors) > 0 {
		fmt.Fprint(w, loginErrors)
		return
	}

	http.Redirect(w, r, "/admin/panel?login=true", http.StatusSeeOther)
}

func adminCRUD(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Get("login") != "true" {

		http.Redirect(w, r, "/not-found", http.StatusSeeOther)
		return
	}

	temp := template.Must(template.ParseFiles("./ui/panel.html"))
	err := temp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
