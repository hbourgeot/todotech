package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"todotech/internal/crud"
)

type templateData struct {
	Products []*crud.Products
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func home(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.URL.Path != "/" {
		http.Redirect(w, r, "/not-found", http.StatusSeeOther)
		return
	}

	temp := template.Must(template.ParseFiles("/go/src/todotech/ui/index.go.html"))
	err := temp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("/go/src/todotech/ui/404.html"))
	err := temp.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("/go/src/todotech/ui/login.html"))
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
		loginErrors["username"] = "This field cannot be more than 100 character"
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

	products, err := crud.GetAllProducts()
	if err != nil {
		log.Fatal(err)
		return
	}

	data := &templateData{
		Products: products,
	}

	temp := template.Must(template.ParseFiles("/go/src/todotech/ui/panel.gohtml"))
	err = temp.Execute(w, data)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func adminCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	insertErrors := make(map[string]string)

	code, err := strconv.Atoi(r.PostForm.Get("insert-code"))
	if err != nil {
		insertErrors["code"] = "The Code must be a number"
	}

	imageUrl := r.PostForm.Get("insert-imageurl")
	price, err := strconv.ParseFloat(r.PostForm.Get("insert-price"), 64)
	if err != nil {
		insertErrors["price"] = "The price must be a number with decimals"
	}

	name := r.PostForm.Get("insert-name")
	cat := r.PostForm.Get("insert-cat")

	if code == 0 {
		insertErrors["code"] = "The code must be different than 0"
	}

	if price == 0 {
		insertErrors["price"] = "The product cannot be free"
	}

	if strings.TrimSpace(imageUrl) == "" {
		insertErrors["image"] = "Please enter the image-url"
	} else if len(strings.TrimSpace(imageUrl)) < 10 {
		insertErrors["image"] = "Please enter the complete URL to the image"
	}

	if strings.TrimSpace(name) == "" {
		insertErrors["name"] = "Please enter the product name"
	}
	if strings.TrimSpace(cat) == "" {
		insertErrors["cat"] = "Please enter the category"
	}

	if len(insertErrors) > 0 {
		fmt.Fprint(w, insertErrors)
		return
	}

	err = crud.InsertProducts(code, name, cat, imageUrl, price)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	http.Redirect(w, r, "/admin/panel?login=true", http.StatusSeeOther)
}

func getProductforCart(w http.ResponseWriter, r *http.Request) {
	var code int
	err := json.NewDecoder(r.Body).Decode(&code)
	if err != nil {
		fmt.Println(err)
		return
	}

	product, err := crud.GetProductByCode(code)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(product)
}

func adminUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	insertErrors := make(map[string]string)

	code, err := strconv.Atoi(r.PostForm.Get("update-code"))
	if err != nil {
		insertErrors["code"] = "The Code must be a number"
	}

	colToUpdate := r.PostForm.Get("col-tomodify")
	colToUpdate = strings.ToLower(colToUpdate)
	newValue := r.PostForm.Get("new-val")

	if code == 0 {
		insertErrors["code"] = "The code must be different than 0"
	}

	if strings.TrimSpace(colToUpdate) == "" || len(strings.TrimSpace(colToUpdate)) == 0 {
		insertErrors["image"] = "Please enter the complete column to update"
	}

	if strings.TrimSpace(newValue) == "" || len(strings.TrimSpace(newValue)) < 3 {
		insertErrors["image"] = "Please enter the complete column to update"
	}

	if len(insertErrors) > 0 {
		fmt.Fprint(w, insertErrors)
		return
	}
	switch colToUpdate {
	case "code":
		newCode, err := strconv.Atoi(newValue)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		err = crud.UpdateProducts(colToUpdate, newCode, code)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		break
	case "name", "cat", "category", "image", "image-url", "url", "imageurl", "img":
		err = crud.UpdateProducts(colToUpdate, newValue, code)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		break
	case "price":
		newPrice, err := strconv.ParseFloat(newValue, 64)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		err = crud.UpdateProducts(colToUpdate, newPrice, code)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		break
	}

	http.Redirect(w, r, "/admin/panel?login=true", http.StatusSeeOther)
}

func adminDelete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		return
	}

	deleteErrors := make(map[string]string)

	code, err := strconv.Atoi(r.PostForm.Get("delete-code"))
	if err != nil {
		deleteErrors["code"] = "The Code must be a number"
	}

	if code == 0 {
		deleteErrors["code"] = "The code must be different than 0"
	}

	if len(deleteErrors) > 0 {
		fmt.Fprint(w, deleteErrors)
		return
	}

	err = crud.DeleteProducts(code)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	http.Redirect(w, r, "/admin/panel?login=true", http.StatusSeeOther)
}

func showProducts(w http.ResponseWriter, r *http.Request) {
	products, err := crud.GetAllProducts()
	if err != nil {
		log.Fatal(err)
		return
	}

	data := &templateData{
		Products: products,
	}

	temp := template.Must(template.ParseFiles("go/src/todotech/ui/products.gohtml"))
	err = temp.Execute(w, data)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func newOrder(w http.ResponseWriter, r *http.Request) {
	var order crud.JSONorder
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		fmt.Println(err)
		return
	}
	var cart []int
	for i := range order.Cart {
		product, err := strconv.Atoi(order.Cart[i])
		if err != nil {
			fmt.Println(err)
			return
		}

		cart = append(cart, product)
	}
	err = crud.CreateClient(order.UserName, order.Name, order.Country)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)
	}

	err = crud.GenerateOrder(order.UserName, cart)
	if err != nil {
		fmt.Println(err)
		return
	}
}
